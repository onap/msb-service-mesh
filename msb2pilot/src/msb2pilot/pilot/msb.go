/**
 * Copyright (c) 2018 ZTE Corporation.
 * All rights reserved. This program and the accompanying materials
 * are made available under the terms of the Eclipse Public License v1.0
 * and the Apache License 2.0 which both accompany this distribution,
 * and are available at http://www.eclipse.org/legal/epl-v10.html
 * and http://www.apache.org/licenses/LICENSE-2.0
 *
 * Contributors:
 *     ZTE - initial Project
 */
package pilot

import (
	"bytes"
	"msb2pilot/log"
	"msb2pilot/models"
	"msb2pilot/msb"
	"os"

	istioModel "istio.io/istio/pilot/pkg/model"
)

var (
	cachedServices []*models.MsbService
)

const (
	defaultVirtualService = "default-apigateway"
)

func SyncMsbData(newServices []*models.MsbService) {
	log.Log.Debug("sync msb rewrite rule to pilot")

	serviceUpdated := isUpdated(cachedServices, newServices)
	if !serviceUpdated { // no service updated
		return
	}
	log.Log.Debug("service updated")

	apiGateway := os.Getenv(models.EnvApiGatewayName)
	publishServices := getPublishServiceMap()
	virtueServiceString := parseServiceToConfig(apiGateway, newServices, publishServices)
	log.Log.Debug(virtueServiceString)
	configs, err := ParseParam(virtueServiceString)

	if err != nil {
		log.Log.Error("param parse error", err)
		return
	}

	updateVirtualService(newServices, configs)
}

func updateVirtualService(newServices []*models.MsbService, configs []istioModel.Config) {
	// if virtualservice exist, then delete it
	config, exist := Get("virtualservice", "default", defaultVirtualService)
	if exist {
		log.Log.Informational("default virtual is: %v", config)
		err := Delete("virtualservice", "default", defaultVirtualService)
		if err != nil {
			log.Log.Debug("failed to delete virture service %v \n", err)
			return
		}
	}

	if len(newServices) == 0 {
		cachedServices = newServices
		return
	}

	fails := Save(OperationCreate, configs)
	if len(fails) != 0 {
		log.Log.Debug("failed to create virture service %v \n", fails)
		return
	} else {
		cachedServices = newServices
	}
}

func isUpdated(oldServices, newServices []*models.MsbService) bool {
	oldServiceMap := toServiceMap(oldServices)
	newServiceMap := toServiceMap(newServices)

	for key, newService := range newServiceMap {
		if oldService, exist := oldServiceMap[key]; exist {
			// service exist: check whether need to update
			if oldService.ModifyIndex != newService.ModifyIndex {
				// service updated
				return true
			}
		} else {
			// old service not exist: add
			return true
		}

		delete(oldServiceMap, key)
	}

	if len(oldServiceMap) != 0 { // some service has been deleted
		return true
	}

	return false
}

func toServiceMap(services []*models.MsbService) map[string]*models.MsbService {
	serviceMap := make(map[string]*models.MsbService)

	for _, service := range services {
		serviceMap[service.ServiceName] = service
	}

	return serviceMap
}

func parseServiceToConfig(host string, services []*models.MsbService, publishServices map[string]*models.PublishService) string {
	httpRoutes := getAllHttpRoute(services, publishServices)

	rule := `{
"apiVersion": "networking.istio.io/v1alpha3",
"kind": "VirtualService",
"metadata": {"name": "` + defaultVirtualService + `"},
"spec": {"hosts":["` + host + `"],"http":[` + httpRoutes + `]}
}`

	return rule
}

func getAllHttpRoute(services []*models.MsbService, publishServices map[string]*models.PublishService) string {
	var buf bytes.Buffer
	hasPre := false
	for _, service := range services {
		if publishService, exist := publishServices[getPublishServiceKey(service)]; exist {

			if service.ConsulLabels.BaseInfo != nil {
				if hasPre {
					buf.WriteString(",")
				}

				rule := createHttpRoute(publishService.PublishUrl, service.ServiceName, service.ConsulLabels.BaseInfo.Url)
				buf.WriteString(rule)

				hasPre = true
			}
		}
	}

	return buf.String()
}

func getPublishServiceKey(svc *models.MsbService) string {
	res := svc.ServiceName

	if svc.ConsulLabels.BaseInfo != nil {
		res += svc.ConsulLabels.BaseInfo.Version
	}

	if svc.ConsulLabels.NameSpace != nil {
		res += svc.ConsulLabels.NameSpace.NameSpace
	}

	return res
}

func getPublishServiceMap() map[string]*models.PublishService {
	publishServices := msb.GetAllPublishServices()

	res := make(map[string]*models.PublishService)

	for _, svc := range publishServices {
		key := svc.ServiceName + svc.Version + svc.NameSpace
		res[key] = svc
	}

	return res
}

//func createRouteRule(sourceService, sourcePath, targetService, targetPath string) string {
//	if sourcePath == "" {
//		sourcePath = "/"
//	}
//	if targetPath == "" {
//		targetPath = "/"
//	}
//	// rule name must consist of lower case alphanuberic charactoers, '-' or '.'. and must start and end with an alphanumberic charactore
//	r := regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*")
//	strs := r.FindAllString(targetService, -1)
//	name := routerulePrefix + strings.Join(strs, "")
//	name = strings.ToLower(name)

//	rule := `{
//"apiVersion": "config.istio.io/v1alpha2",
//"kind": "RouteRule",
//"metadata": {
//  "name": "` + name + `"
//},
//"spec": {
//  "destination":{
//    "name":"` + sourceService + `"
//  },
//  "match":{
//    "request":{
//      "headers": {
//        "uri": {
//          "prefix": "` + sourcePath + `"
//        }
//      }
//    }
//  },
//  "rewrite": {
//    "uri": "` + targetPath + `"
//  },
//  "route":[
//    {
//      "destination":{
//        "name":"` + targetService + `"
//      }
//    }
//  ]
//}
//}

//`
//	return rule
//}

//func createRouteRule(sourceService, sourcePath, targetService, targetPath string) string {
//	if sourcePath == "" {
//		sourcePath = "/"
//	}
//	if targetPath == "" {
//		targetPath = "/"
//	}

//	rule := `
//apiVersion: networking.istio.io/v1alpha3
//kind: VirtualService
//metadata:
//  name: default-apigateway
//spec:
//  hosts:
//  - reviews
//  http:
//  - match:
//    - headers:
//        end-user:
//          exact: jason
//    route:
//    - destination:
//        host: reviews
//  - route:
//    - destination:
//        host: reviews
//	`
//	return rule
//}

func createHttpRoute(sourcePath, targetHost, targetPath string) string {
	//	- match:
	//    - uri:
	//        prefix: /ratings
	//    rewrite:
	//      uri: /v1/bookRatings
	//    route:
	//    - destination:
	//        host: ratings.prod.svc.cluster.local
	//        subset: v1

	if sourcePath == "" {
		sourcePath = "/"
	}
	if targetPath == "" {
		targetPath = "/"
	}

	httpRoute := `{
"match":[{"uri": {"prefix": "` + sourcePath + `"}}],
"rewrite": {"uri": "` + targetPath + `"},
"route": [` + createDestinationWeight(targetHost) + `]
}`

	return httpRoute
}

func createDestinationWeight(targetHost string) string {
	//	destination:
	//     host: reviews.prod.svc.cluster.local
	//     subset: v2
	//  weight: 25

	return `{"destination": {"host": "` + targetHost + `"}}`
}
