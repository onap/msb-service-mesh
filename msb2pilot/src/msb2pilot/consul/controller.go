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
	"msb2pilot/util"
	"os"
	"path/filepath"

	"github.com/hashicorp/consul/api"
)

var client *api.Client
var consulAddress string

var (
	cfgFilePath = filepath.Join(util.GetCfgPath(), "consul.yml")
)

const (
	defaultAddress = "http://localhost:8500"
)

func init() {
	consulAddress = getConsulAddress(cfgFilePath)

	conf := api.DefaultConfig()
	conf.Address = consulAddress
	var err error
	client, err = api.NewClient(conf)

	if err != nil {
		log.Log.Error("failed to init consul client", err)
	}
}

func getConsulAddress(path string) string {
	res := os.Getenv(models.EnvConsulAddress)
	if res != "" {
		return res
	}

	cfg, err := loadCfgInfo(path)
	if err != nil {
		log.Log.Error("load consul config info error", err)
		return defaultAddress
	} else {
		if addr, exist := cfg["address"]; exist {
			return addr.(string)
		} else {
			return defaultAddress
		}
	}
}

func loadCfgInfo(path string) (map[interface{}]interface{}, error) {
	log.Log.Informational("consul config path is:" + path)
	cfg, err := util.Read(path)
	if err != nil {
		return nil, err
	}

	result := make(map[interface{}]interface{})
	err = util.UnmarshalYaml(cfg, result)
	if err != nil {
		return nil, err
	}
	return result, nil
}

func GetServices() (map[string][]string, error) {
	data, _, err := client.Catalog().Services(nil)

	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetInstances(serviceName string) ([]*api.CatalogService, error) {
	endpoints, _, err := client.Catalog().Service(serviceName, "", nil)
	if err != nil {
		log.Log.Error("can not get endpoints of ", serviceName)
		return nil, err
	}
	return endpoints, nil
}
