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
	"path/filepath"
	"strings"
)

const (
	ConfigPath = "conf"
)

func GetCfgPath() string {
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}

	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	result := filepath.Join(workPath, ConfigPath)

	if !FileExists(result) {
		result = filepath.Join(appPath, ConfigPath)
		if !FileExists(result) {
			goPaths := GetGoPath()
			for _, path := range goPaths {
				result = filepath.Join(path, "src", "msb2pilot", ConfigPath)
				if FileExists(result) {
					return result
				}
			}
			result = "/"
		}
	}

	return result

}

func GetGoPath() []string {
	paths := os.Getenv("GOPATH")
	if strings.Contains(paths, ";") { // windows
		return strings.Split(paths, ";")
	} else if strings.Contains(paths, ":") { // linux
		return strings.Split(paths, ":")
	} else if paths != "" { // only one
		path := make([]string, 1, 1)
		path[0] = paths
		return path
	} else {
		return make([]string, 0, 0)
	}
}

func FileExists(path string) bool {
	if _, err := os.Stat(path); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
