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
package controllers

import (
	"encoding/json"
	"msb2pilot/log"
	"msb2pilot/pilot"

	"github.com/astaxie/beego"
)

type RouteRuleController struct {
	beego.Controller
}

func (this *RouteRuleController) Get() {
	log.Log.Debug("routerule controller get method called")

	name := this.GetString("name")
	namespace := this.GetString("namespace")
	log.Log.Debug("name is : " + name)
	log.Log.Debug("namespace is: " + namespace)
	if name != "" { // get routerule detail
		data, _ := pilot.Get("routerules", namespace, name)
		b, _ := json.Marshal(data)
		log.Log.Debug(string(b))

		this.Data["json"] = pilot.ConvertConfig(*data)

	} else { // get routerule name list
		data, _ := pilot.List("routerules", "")
		b, _ := json.Marshal(data)
		log.Log.Debug(string(b))
		rules := make([]interface{}, 0, len(data))
		for _, config := range data {
			rules = append(rules, pilot.ConvertConfig(config))
		}

		this.Data["json"] = rules
	}
	this.ServeJSON()
}
