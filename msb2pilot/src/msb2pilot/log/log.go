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
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

type ConsoleCfg struct {
	Level int `json:"level"`
}

type FileCfg struct {
	FileName string `json:"filename"`
	Level    int    `json:"level"`
	MaxLines int    `json:"maxlines"`
	MaxSize  int    `josn:"maxsize"`
	Daily    bool   `json:"daily"`
	MaxDays  int    `json:"maxdays"`
	Rotate   bool   `json:"rotate"`
}

type Cfg struct {
	Console ConsoleCfg `json:"console"`
	File    FileCfg    `json:"file"`
}

const (
	cfgFileName         = "log.yml"
	defaultConsoleLevel = "Warn"
	defaultFileLevel    = "Info"
)

var (
	Log         *logs.BeeLogger
	loggerLevel = map[string]int{"Emergency": 0, "Alert": 1, "Critical": 2, "Error": 3, "Warn": 4, "Notice": 5, "Info": 6, "Debug": 7}
)

func init() {
	Log = logs.NewLogger()
	Log.EnableFuncCallDepth(true)

	cfg := getDefaultCfg()
	setLogger(logs.AdapterConsole, &cfg.Console)
	checkLogDir(cfg.File.FileName)
	setLogger(logs.AdapterFile, &cfg.File)
}

func setLogger(adapter string, cfg interface{}) bool {
	b, err := json.Marshal(cfg)
	if err != nil {
		fmt.Printf(" cfg json trans error: %s\n", adapter, err.Error())
		return false
	}

	err = Log.SetLogger(adapter, string(b))
	if err != nil {
		fmt.Printf("set %s failed: %s\n", adapter, err.Error())
		return false
	}

	return true

}

func checkLogDir(path string) {
	if path == "" {
		return
	}

	var index int
	pathSep := string(os.PathSeparator)
	if index = strings.LastIndex(path, pathSep); index <= 2 {
		return
	}

	perm, _ := strconv.ParseInt("0660", 8, 64)
	if err := os.MkdirAll(path[0:index], os.FileMode(perm)); err != nil {
		return
	}
}

func getDefaultCfg() *Cfg {
	return &Cfg{
		Console: ConsoleCfg{
			Level: loggerLevel[defaultConsoleLevel],
		},
		File: FileCfg{
			FileName: "msb2pilot.log",
			Level:    loggerLevel[defaultFileLevel],
			MaxLines: 300000,
			MaxSize:  30 * 1024 * 1024,
			Daily:    true,
			MaxDays:  10,
			Rotate:   true,
		},
	}
}
