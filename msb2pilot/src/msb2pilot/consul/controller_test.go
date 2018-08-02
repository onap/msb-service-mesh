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
	"msb2pilot/models"
	"os"
	"testing"
)

func TestSetConsulAddress(t *testing.T) {
	cases := []struct {
		env, path, want string
	}{
		{
			env:  "testEnv",
			path: "",
			want: `testEnv`,
		},
		{
			env:  "",
			path: cfgFilePath,
			want: `http://127.0.0.1:8500`,
		},
		{
			env:  "testEnvWithPath",
			path: cfgFilePath,
			want: `testEnvWithPath`,
		},
		{
			env:  "",
			path: ``,
			want: `http://localhost:8500`,
		},
		{
			env:  "",
			path: `controller.go`,
			want: `http://localhost:8500`,
		},
	}

	oldEnv := os.Getenv(models.EnvConsulAddress)

	for _, cas := range cases {
		os.Setenv(models.EnvConsulAddress, cas.env)

		res := getConsulAddress(cas.path)
		if res != cas.want {
			t.Errorf("getConsulAddress() => want %s, got %s", cas.want, res)
		}
	}

	os.Setenv(models.EnvConsulAddress, oldEnv)
}

func TestLoadCfgInfo(t *testing.T) {
	cases := []struct {
		path, status string
	}{
		{
			path:   cfgFilePath,
			status: `success`,
		},
		{
			path:   ``,
			status: `path is empty`,
		},
		{
			path:   `controller.go`,
			status: `yaml format error`,
		},
	}

	for _, cas := range cases {
		_, err := loadCfgInfo(cas.path)
		if (cas.status == "success" && err != nil) || (cas.status != "success" && err == nil) {
			t.Errorf("loadCfgInfo() => want %s, got %v", cas.status, err)
		}
	}
}
