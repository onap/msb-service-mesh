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
package consul

import (
	"msb2pilot/log"
	"msb2pilot/models"
	"time"

	"github.com/hashicorp/consul/api"
)

type Monitor interface {
	Start(<-chan struct{})
}

type ServiceHandler func(newServices []*models.MsbService)

type consulMonitor struct {
	discovery      *api.Client
	period         time.Duration
	serviceHandler ServiceHandler
}

func NewConsulMonitor(client *api.Client, period time.Duration, serviceHandler ServiceHandler) Monitor {
	return &consulMonitor{
		discovery:      client,
		period:         period,
		serviceHandler: serviceHandler,
	}
}

func (this *consulMonitor) Start(stop <-chan struct{}) {
	this.run(stop)
}

func (this *consulMonitor) run(stop <-chan struct{}) {
	ticker := time.NewTicker(this.period)
	for {
		select {
		case <-stop:
			ticker.Stop()
			return
		case <-ticker.C:
			this.updateServiceRecord()
		}
	}

}

func (this *consulMonitor) updateServiceRecord() {
	data, err := GetServices()
	if err != nil {
		log.Log.Error("failed to get services from consul", err)
		return
	}

	newRecords := make([]*models.MsbService, 0, len(data))
	for name := range data {
		endpoints, err := GetInstances(name)
		if err != nil {
			log.Log.Error("failed to get service instance of "+name, err)
			continue
		}
		newRecords = append(newRecords, models.ConvertService(endpoints))
	}

	this.serviceHandler(newRecords)
}
