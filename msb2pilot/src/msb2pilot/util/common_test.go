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
package util

import (
	"os"
	"strings"
	"testing"
)

func TestGetCfgPath(t *testing.T) {
	got := GetCfgPath()
	if !strings.Contains(got, "conf") {
		t.Errorf("GetCfgPath() => got %v, should contains `ocnf`", got)
	}
}

func TestGetGoPath(t *testing.T) {
	oldPaths := os.Getenv("GOPATH")
	cases := []struct {
		in   string
		want []string
	}{
		{ // window
			in: `path1;path2;path3`,
			want: []string{
				`path1`,
				`path2`,
				`path3`,
			},
		},
		{ // linux
			in: `path1:path2:path3`,
			want: []string{
				`path1`,
				`path2`,
				`path3`,
			},
		},
		{ // single Path
			in: `path1`,
			want: []string{
				`path1`,
			},
		},
		{ // single Path
			in: `;`,
			want: []string{
				``, ``,
			},
		},
	}

	for _, cas := range cases {
		os.Setenv("GOPATH", cas.in)
		got := GetGoPath()

		if len(cas.want) != len(got) {
			t.Errorf("GetGoPath() => different size, got %d, want %d, %v, %v", len(got), len(cas.want), got, cas.want)
		}

		for i, item := range cas.want {
			if item != got[i] {
				t.Errorf("GetGoPath() => got %v, want %v", got, cas.want)
				break
			}
		}
	}

	// unset test
	os.Unsetenv("GOPATH")
	got := GetGoPath()
	if len(got) != 0 {
		t.Errorf("GetGoPath() => unset env test got len %d, want 0", len(got))
	}

	os.Setenv("GOPATH", oldPaths)
}

func TestFileExists(t *testing.T) {
	existFile := `common_test.go`
	notExistFile := `common_test.go_11`

	exist := FileExists(existFile)
	if !exist {
		t.Errorf("FileExists(%s) => got false, want true", existFile)
	}

	exist = FileExists(notExistFile)
	if exist {
		t.Errorf("FileExists(%s) => got true, want false", notExistFile)
	}
}
