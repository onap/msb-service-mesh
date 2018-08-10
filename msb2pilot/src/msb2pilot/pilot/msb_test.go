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
	"msb2pilot/models"
	"testing"
)

func TestParseServiceToConfig(t *testing.T) {
	cases := []struct {
		services        []*models.MsbService
		publishServices map[string]*models.PublishService
		want            string
	}{
		{
			services:        []*models.MsbService{},
			publishServices: map[string]*models.PublishService{},
			want: `{
"apiVersion": "networking.istio.io/v1alpha3",
"kind": "VirtualService",
"metadata": {"name": "default-apigateway"},
"spec": {"hosts":["tService"],"http":[]}
}`,
		},
		{
			services: []*models.MsbService{
				&models.MsbService{
					ConsulLabels: &models.ConsulLabels{
						NameSpace: &models.NameSpace{
							NameSpace: "service1namespace",
						},
						BaseInfo: &models.BaseInfo{
							Version: "service1v1",
							Url:     "service1url",
						},
					},
					ServiceName: "service1",
				},
				&models.MsbService{
					ConsulLabels: &models.ConsulLabels{
						NameSpace: &models.NameSpace{
							NameSpace: "service2namespace",
						},
						BaseInfo: &models.BaseInfo{
							Version: "service2v2",
							Url:     "service2url",
						},
					},
					ServiceName: "service2",
				},
			},
			publishServices: map[string]*models.PublishService{
				"service1service1v1service1namespace": &models.PublishService{
					ServiceName: "service1",
					Version:     "service1v1",
					NameSpace:   "service1namespace",
					PublishUrl:  "service1publishurl",
				},
				"service2service2v2service2namespace": &models.PublishService{
					ServiceName: "service2",
					Version:     "service2v2",
					NameSpace:   "service2namespace",
					PublishUrl:  "service2publihurl",
				},
			},
			want: `{
"apiVersion": "networking.istio.io/v1alpha3",
"kind": "VirtualService",
"metadata": {"name": "default-apigateway"},
"spec": {"hosts":["tService"],"http":[{
"match":{"uri": {"prefix": "service1publishurl"}},
"rewrite": {"uri": "service1url"},
"route": [{"destination": {"host": "service1"}}]
},{
"match":{"uri": {"prefix": "service2publihurl"}},
"rewrite": {"uri": "service2url"},
"route": [{"destination": {"host": "service2"}}]
}]}
}`,
		},
	}

	for _, cas := range cases {
		got := parseServiceToConfig("tService", cas.services, cas.publishServices)
		if got != cas.want {
			t.Errorf("parseServiceToConfig() => got %s, want %s", got, cas.want)
		}
	}
}

//func TestCreateHttpRoute(t *testing.T) {
//	cases := []struct {
//		sPath, tService, tPath, want string
//	}{
//		{ // success demo
//			sPath:    "/",
//			tService: "tService",
//			tPath:    "/",
//			want: `{
//"match":{"uri": {"prefix": "/"}},
//"rewrite": {"uri": "/"},
//"route": [{"destination": {"host": "tService"}}]
//}`,
//		},
//	}

//	for _, cas := range cases {
//		got := createHttpRoute(cas.sPath, cas.tService, cas.tPath)
//		if got != cas.want {
//			t.Errorf("createHttpRoute(%s, %s, %s) => got %s, want %s", cas.sPath, cas.tService, cas.tPath, got, cas.want)
//		}
//	}
//}
