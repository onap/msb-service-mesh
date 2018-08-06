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
	"testing"
)

func TestCreateRouteRule(t *testing.T) {
	cases := []struct {
		sService, sPath, tService, tPath, want string
	}{
		{ // success demo
			sService: "sservice",
			sPath:    "/",
			tService: "tservice",
			tPath:    "/",
			want: `{
"apiVersion": "config.istio.io/v1alpha2",
"kind": "RouteRule",
"metadata": {
  "name": "msbcustom.tservice"
},
"spec": {
  "destination":{
    "name":"sservice"
  },
  "match":{
    "request":{
      "headers": {
        "uri": {
          "prefix": "/"
        }
      }
    }
  },
  "rewrite": {
    "uri": "/"
  },
  "route":[
    {
      "destination":{
        "name":"tservice"
      }
    }
  ]
}
}

`,
		},
		{ // rule name must consist of lower case alphanuberic charactoers, '-' or '.'. and must start and end with an alphanumberic charactore
			sService: "sservice",
			sPath:    "/",
			tService: "123ABCrule-name.test~!@#$%^&*()_+321",
			tPath:    "/",
			want: `{
"apiVersion": "config.istio.io/v1alpha2",
"kind": "RouteRule",
"metadata": {
  "name": "msbcustom.123rule-name.test321"
},
"spec": {
  "destination":{
    "name":"sservice"
  },
  "match":{
    "request":{
      "headers": {
        "uri": {
          "prefix": "/"
        }
      }
    }
  },
  "rewrite": {
    "uri": "/"
  },
  "route":[
    {
      "destination":{
        "name":"123ABCrule-name.test~!@#$%^&*()_+321"
      }
    }
  ]
}
}

`,
		},
	}

	for _, cas := range cases {
		got := createRouteRule(cas.sService, cas.sPath, cas.tService, cas.tPath)
		if got != cas.want {
			t.Errorf("createRouteRule(%s, %s, %s, %s) => got %s, want %s", cas.sService, cas.sPath, cas.tService, cas.tPath, got, cas.want)
		}
	}
}
