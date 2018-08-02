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
package log

import (
	"msb2pilot/util"
	"os"
	"strings"
	"testing"
)

func TestCheckLogDir(t *testing.T) {
	cases := []struct {
		path  string
		exist bool
	}{
		{
			path:  ``,
			exist: false,
		},
		{
			path:  `test.log`,
			exist: false,
		},
		{
			path:  `.` + util.PathSep + `test` + util.PathSep + `test.log`,
			exist: true,
		},
	}

	for _, cas := range cases {
		checkLogDir(cas.path)

		index := strings.LastIndex(cas.path, util.PathSep)
		if cas.exist && !util.FileExists(cas.path[0:index]) {
			t.Errorf("checkLogDir() => dir not exist, want %s", cas.path)
		}
	}

	// clear
	os.RemoveAll("test")
}

func TestLoadCustom(t *testing.T) {
	cases := []struct {
		path string
		want string
	}{
		{
			path: `..` + util.PathSep + "conf" + util.PathSep + cfgFileName,
			want: "success",
		},
		{
			path: ``,
			want: "read file error",
		},
		{
			path: `log_test.go`,
			want: "parse config file error",
		},
	}

	for _, cas := range cases {
		res := loadCustom(cas.path)

		if (res == nil && cas.want == "success") || (res != nil && cas.want != "success") {
			t.Errorf("loadCustom() => want %s, got %v", cas.want, res)
		}
	}
}
