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
	"errors"
	"msb2pilot/log"
	"msb2pilot/models"
	"msb2pilot/util"
	"os"
	"path/filepath"

	"istio.io/istio/pilot/pkg/config/kube/crd"
	"istio.io/istio/pilot/pkg/model"
)

type Operation string

var (
	client     *crd.Client
	configPath = filepath.Join(util.GetCfgPath(), "k8s.yml")
)

func init() {
	updateK8sAddress(configPath)

	var err error
	client, err = crd.NewClient(configPath, model.ConfigDescriptor{
		model.RouteRule,
		model.DestinationPolicy,
		model.DestinationRule,
	}, "")

	if err != nil {
		log.Log.Error("fail to init crd", err)
	}
}

func updateK8sAddress(path string) (string, error) {
	addr := os.Getenv(models.EnvK8sAddress)
	log.Log.Informational("k8s cfg address from env: ", addr)
	if addr == "" {
		return "", nil
	}

	// load cfg file
	cfgstr, err := util.Read(path)
	if err != nil {
		log.Log.Error("file to load k8s config file", err)
		return "", err
	}

	// update address
	cfg := make(map[string]interface{})
	util.UnmarshalYaml(cfgstr, &cfg)
	if clusters, exist := cfg["clusters"]; exist {
		clusterItem := clusters.([]interface{})[0]
		cluster, _ := clusterItem.(map[interface{}]interface{})["cluster"]
		cluster.(map[interface{}]interface{})["server"] = addr
	}

	updatedCfgstr, _ := util.MarshalYaml(cfg)

	err = util.Write(path, updatedCfgstr, 0644)
	if err != nil {
		log.Log.Error("fail to write k8s cfg info to file", err)
	}

	return addr, err
}

func Get(typ, namespace, name string) (*model.Config, bool) {
	proto, err := protoSchema(typ)
	if err != nil {
		log.Log.Informational("get resource error", err)
		return &model.Config{}, false
	}
	return client.Get(proto.Type, name, namespace)
}

func protoSchema(typ string) (model.ProtoSchema, error) {
	for _, desc := range client.ConfigDescriptor() {
		switch typ {
		case crd.ResourceName(desc.Type), crd.ResourceName(desc.Plural):
			return desc, nil
		}
	}
	return model.ProtoSchema{}, errors.New("can not find this kind of resources:[" + typ + "]")
}

func List(typ, namespace string) ([]model.Config, error) {
	proto, err := protoSchema(typ)
	if err != nil {
		return nil, err
	}
	return client.List(proto.Type, namespace)
}
