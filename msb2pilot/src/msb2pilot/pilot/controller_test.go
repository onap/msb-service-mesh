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
	"msb2pilot/models"
	"os"
	"reflect"
	"testing"
)

func TestList(t *testing.T) {
	res, err := List("routerules", "default")
	if err != nil {
		t.Errorf("List() => got %v", err)
	} else {
		fmt.Print(res)
	}
}

func TestUpdateK8sAddress(t *testing.T) {
	cases := []struct {
		path, addr, want, err string
	}{
		{
			path: "k8s.yml222",
			addr: "filenoteexisttest",
			want: "",
			err:  "*os.PathError",
		},
		{
			path: configPath,
			addr: "",
			want: "",
			err:  "",
		},
		{
			path: configPath,
			addr: "k8stest",
			want: "k8stest",
			err:  "",
		},
	}

	oldEnv := os.Getenv(models.EnvK8sAddress)
	for _, cas := range cases {
		os.Unsetenv(models.EnvK8sAddress)
		os.Setenv(models.EnvK8sAddress, cas.addr)

		got, err := updateK8sAddress(cas.path)
		if got != cas.want || (err != nil && reflect.TypeOf(err).String() != cas.err) {
			t.Errorf("updateK8sAddress(%s, %s) => got %s %v, want %s", cas.path, cas.addr, got, reflect.TypeOf(err), cas.want)
		}
	}
	os.Setenv(models.EnvK8sAddress, oldEnv)
}
