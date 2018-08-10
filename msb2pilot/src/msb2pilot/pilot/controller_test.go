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
	"fmt"

	//	"fmt"
	//	"msb2pilot/models"
	//	"os"
	//	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	res, err := List("virtualservice", "default")
	if err != nil {
		t.Errorf("List() => got %v", err)
	} else {
		fmt.Print(res)
	}
}

//func TestParseParam(t *testing.T) {
//	cases := []struct {
//		in string
//	}{
//		//		{
//		//			in: `{
//		//"apiVersion": "networking.istio.io/v1alpha3",
//		//"kind": "VirtualService",
//		//"metadata": {"name": "default-apigateway"},
//		//"spec": {"hosts":["apigateway"],"http":[{
//		//"match":{"uri": {"prefix": "/portaladmin"}},
//		//"rewrite": {"uri": "/portaladmin"},
//		//"route": [{"destination": {"host": "portaladmin"}}]
//		//},{
//		//"match":{"uri": {"prefix": "/pm_mgt/v1"}},
//		//"rewrite": {"uri": "/pm_mgt/v1"},
//		//"route": [{"destination": {"host": "pm_mgt"}}]
//		//}]}
//		//}`,
//		//		},
//		{
//			in: `{
//"apiVersion": "networking.istio.io/v1alpha3",
//"kind": "VirtualService",
//"metadata": {"name": "default-apigateway"},
//"spec": {"destination":{"service":"reviews.service.consul"},"http":[{
//"match":{"uri": {"prefix": "/portaladmin"}},
//"rewrite": {"uri": "/portaladmin"},
//}]}
//}`,
//		},
//		{
//			in: `{
//"apiVersion": "networking.istio.io/v1alpha3",
//"kind": "VirtualService",
//"metadata": {"name": "default-apigateway"},
//"spec": {"hosts":["test"],"http":[]}
//}`,
//		},
//	}

//	for _, cas := range cases {
//		res, err := ParseParam(cas.in)
//		if err != nil {
//			t.Errorf("ParseParam() => got %v", err)
//		} else {
//			fmt.Print(res)
//		}
//	}
//}

func TestCreate(t *testing.T) {
	str := `
		{
			"apiVersion": "networking.istio.io/v1alpha3",
			"kind": "VirtualService",
			"metadata":{
			  "name": "reviews"},
			"spec":{
			  "hosts":["reviews.service.consul"],
			  "http":[{
				"match":[{"uri": {"prefix": "/pm_mgt/v1"}}],
				"rewrite": {"uri": "/portaladmin"},
			 	"route":[{
			    	"destination":{
			        	"host": "reviews.service.consul",
			        	"subset": "v3"
					}}]
				}]
			}
		}
		`

	config, exist := Get("virtualservice", "default", "reviews")
	if exist {
		Delete("virtualservice", "default", "reviews")
	}
	configs, err := ParseParam(str)
	if err != nil {
		t.Errorf("ParseParam() => got %v", err)
	} else {
		fmt.Println(configs)
	}

	res, err := Create(&configs[0])
	if err != nil {
		t.Errorf("Create() => got %v", err)
	} else {
		fmt.Println(res)
	}

	if exist {
		Create(config)
	}
}

//func TestParseParam(t *testing.T) {
//	str := `
//	{
//		"apiVersion": "networking.istio.io/v1alpha3",
//		"kind": "VirtualService",
//		"metadata":{
//		  "name": "reviews"},
//		"spec":{
//		  "hosts":["reviews.service.consul"],
//		  "http":[{
//		 	"route":[{
//		    	"destination":{
//		        	"host": "reviews.service.consul",
//		        	"subset": "v3"
//				}}]
//			}]
//		}
//	}
//	`
//	res, err := ParseParam(str)
//	if err != nil {
//		t.Errorf("ParseParam() => got %v", err)
//	} else {
//		fmt.Println(res)
//	}
//}

//func TestUpdateK8sAddress(t *testing.T) {
//	cases := []struct {
//		path, addr, want, err string
//	}{
//		{
//			path: "k8s.yml222",
//			addr: "filenoteexisttest",
//			want: "",
//			err:  "*os.PathError",
//		},
//		{
//			path: configPath,
//			addr: "",
//			want: "",
//			err:  "",
//		},
//		{
//			path: configPath,
//			addr: "k8stest",
//			want: "k8stest",
//			err:  "",
//		},
//	}

//	oldEnv := os.Getenv(models.EnvK8sAddress)
//	for _, cas := range cases {
//		os.Unsetenv(models.EnvK8sAddress)
//		os.Setenv(models.EnvK8sAddress, cas.addr)

//		got, err := updateK8sAddress(cas.path)
//		if got != cas.want || (err != nil && reflect.TypeOf(err).String() != cas.err) {
//			t.Errorf("updateK8sAddress(%s, %s) => got %s %v, want %s", cas.path, cas.addr, got, reflect.TypeOf(err), cas.want)
//		}
//	}
//	os.Setenv(models.EnvK8sAddress, oldEnv)
//}
