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
package models

import (
	"encoding/json"
	"msb2pilot/log"
	"strings"

	"github.com/hashicorp/consul/api"
)

func ConvertService(endpoints []*api.CatalogService) *MsbService {
	if len(endpoints) == 0 {
		return nil
	}

	endpoint := endpoints[0]
	service := &MsbService{
		ServiceName:  endpoint.ServiceName,
		ModifyIndex:  endpoint.ModifyIndex,
		ConsulLabels: &ConsulLabels{},
	}

	convertMsbLabels(service.ConsulLabels, endpoint.ServiceTags)
	return service
}

func convertBaseInfo(baseString string) (baseInfo *BaseInfo, err error) {
	baseInfo = new(BaseInfo)
	err = json.Unmarshal([]byte(baseString), baseInfo)

	return
}

func convertNameSpace(ns string) (nameSpace *NameSpace, err error) {
	nameSpace = new(NameSpace)
	err = json.Unmarshal([]byte(ns), nameSpace)
	return
}

func LoadMsbServiceFromMap(services map[string][]string) []*MsbService {
	result := make([]*MsbService, 0, len(services))

	for k, v := range services {
		service := &MsbService{
			ServiceName:  k,
			ConsulLabels: new(ConsulLabels),
		}
		convertMsbLabels(service.ConsulLabels, v)
		result = append(result, service)
	}
	return result
}

func ConsulService2MsbService(consulService *api.CatalogService) *MsbService {
	msbService := &MsbService{
		ServiceName:    consulService.ServiceName,
		ServiceAddress: consulService.ServiceAddress,
		ServicePort:    consulService.ServicePort,
		ConsulLabels:   new(ConsulLabels),
	}

	convertMsbLabels(msbService.ConsulLabels, consulService.ServiceTags)

	return msbService
}

func convertMsbLabel(label, labelstr string) interface{} {
	var result interface{}
	var err error

	labelPrefix := "\"" + label + "\":"

	if strings.HasPrefix(labelstr, labelPrefix) {
		kvp := strings.Split(labelstr, labelPrefix)
		value := kvp[1]

		switch label {
		case "base":
			result, err = convertBaseInfo(value)
		case "ns":
			result, err = convertNameSpace(value)
		}

		if err != nil {
			log.Log.Error("parse msb label error", err)
			return nil
		}

		return result
	}

	return nil
}
func convertMsbLabels(consulLabels *ConsulLabels, labels []string) {
	for _, label := range labels {
		baseInfo := convertMsbLabel("base", label)
		if baseInfo != nil {
			consulLabels.BaseInfo = baseInfo.(*BaseInfo)
		}

		ns := convertMsbLabel("ns", label)
		if ns != nil {
			consulLabels.NameSpace = ns.(*NameSpace)
		}
	}
}
