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
	"regexp"
	"strings"

	istioModel "istio.io/istio/pilot/pkg/model"
)

var (
	cachedServices []*models.MsbService
)

const (
	routerulePrefix = "msbcustom."
)

func SyncMsbData(newServices []*models.MsbService) {
	if len(cachedServices) == 0 {
		deleteAllMsbRules()
	}
	log.Log.Debug("sync msb rewrite rule to pilot")
	createServices, updateServices, deleteServices := compareServices(cachedServices, newServices)

	saveService(OperationCreate, createServices)
	saveService(OperationUpdate, updateServices)
	saveService(OperationDelete, deleteServices)

	cachedServices = newServices
}

func saveService(operation Operation, services []*models.MsbService) {
	if len(services) == 0 {
		return
	}
	configs, err := parseServiceToConfig(services)
	if err != nil {
		log.Log.Error("param parse error", err)
		return
	}
	fails := Save(operation, configs)
	log.Log.Debug("%d services need to %s, %d fails. \n", len(services), operation, len(fails))
}

func deleteAllMsbRules() {
	log.Log.Informational("delete all msb rules")
	configs, err := List("route-rule", "")

	if err != nil {
		return
	}

	deleteList := msbRuleFilter(configs)

	Save(OperationDelete, deleteList)
}

func msbRuleFilter(configs []istioModel.Config) []istioModel.Config {
	res := make([]istioModel.Config, 0, len(configs))

	for _, config := range configs {
		if strings.HasPrefix(config.Name, routerulePrefix) {
			res = append(res, config)
		}
	}

	return res
}

func compareServices(oldServices, newServices []*models.MsbService) (createServices, updateServices, deleteServices []*models.MsbService) {
	oldServiceMap := toServiceMap(oldServices)
	newServiceMap := toServiceMap(newServices)

	for key, newService := range newServiceMap {
		if oldService, exist := oldServiceMap[key]; exist {
			// service exist: check whether need to update
			if oldService.ModifyIndex != newService.ModifyIndex {
				updateServices = append(updateServices, newService)
			}
		} else {
			// service not exist: add
			createServices = append(createServices, newService)
		}

		delete(oldServiceMap, key)
	}

	for _, service := range oldServiceMap {
		deleteServices = append(deleteServices, service)
	}

	return
}

func toServiceMap(services []*models.MsbService) map[string]*models.MsbService {
	serviceMap := make(map[string]*models.MsbService)

	for _, service := range services {
		serviceMap[service.ServiceName] = service
	}

	return serviceMap
}

func parseServiceToConfig(services []*models.MsbService) ([]istioModel.Config, error) {
	publishServices := getPublishServiceMap()
	apiGateway := os.Getenv(models.EnvApiGatewayName)
	var buf bytes.Buffer
	for _, service := range services {
		if publishService, exist := publishServices[getPublishServiceKey(service)]; exist {

			if service.ConsulLabels.BaseInfo != nil {
				rule := createRouteRule(apiGateway, publishService.PublishUrl, service.ServiceName, service.ConsulLabels.BaseInfo.Url)
				buf.WriteString(rule)
			}
		}
	}
	return ParseParam(buf.String())
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

func createRouteRule(sourceService, sourcePath, targetService, targetPath string) string {
	if sourcePath == "" {
		sourcePath = "/"
	}
	if targetPath == "" {
		targetPath = "/"
	}
	// rule name must consist of lower case alphanuberic charactoers, '-' or '.'. and must start and end with an alphanumberic charactore
	r := regexp.MustCompile("[a-z0-9]([-a-z0-9]*[a-z0-9])?(\\.[a-z0-9]([-a-z0-9]*[a-z0-9])?)*")
	strs := r.FindAllString(targetService, -1)
	name := routerulePrefix + strings.Join(strs, "")
	name = strings.ToLower(name)

	rule := `{
"apiVersion": "config.istio.io/v1alpha2",
"kind": "RouteRule",
"metadata": {
  "name": "` + name + `"
},
"spec": {
  "destination":{
    "name":"` + sourceService + `"
  },
  "match":{
    "request":{
      "headers": {
        "uri": {
          "prefix": "` + sourcePath + `"
        }
      }
    }
  },
  "rewrite": {
    "uri": "` + targetPath + `"
  },
  "route":[
    {
      "destination":{
        "name":"` + targetService + `"
      }
    }
  ]
}
}

`
	return rule
}
