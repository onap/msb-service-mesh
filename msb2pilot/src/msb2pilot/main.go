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
package main

import (
	"fmt"
	"msb2pilot/consul"
	"msb2pilot/log"
	"msb2pilot/models"
	_ "msb2pilot/routers"
	"time"

	"github.com/astaxie/beego"
)

func main() {
	log.Log.Informational("**************** init msb2pilot ************************")
	// start sync msb data
	go syncConsulData()

	beego.Run()
}

func syncConsulData() {
	stop := make(chan struct{})
	monitor := consul.NewConsulMonitor(nil, 20*time.Second, syncMsbData)
	monitor.Start(stop)
}

func syncMsbData(newServices []*models.MsbService) {
	fmt.Println(len(newServices), "services updated", time.Now())
}
