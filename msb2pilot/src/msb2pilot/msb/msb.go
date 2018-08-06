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
package msb

import (
	"encoding/json"
	"io/ioutil"
	"msb2pilot/log"
	"msb2pilot/models"
	"net/http"
	"os"
)

var (
	msbAddr = "http://localhost:9081"
)

func getBaseUrl() string {
	baseUrl := os.Getenv(models.EnvMsbAddress)
	if baseUrl == "" {
		baseUrl = msbAddr
	}

	return baseUrl
}

func GetAllPublishServices() []*models.PublishService {
	url := getBaseUrl() + "/api/msdiscover/v1/publishservicelist"
	res, err := http.Get(url)

	if err != nil {
		log.Log.Error("fail to get public address", url, err)
		return nil
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Log.Error("fail to read response", err)
		return nil
	}

	result := make([]*models.PublishService, 0)
	err = json.Unmarshal(b, &result)
	if err != nil {
		log.Log.Error("fail to unmarshal publish address", err)
		return nil
	}

	return result
}
