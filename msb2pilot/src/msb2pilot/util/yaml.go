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
	"gopkg.in/yaml.v2"
)

/**
* type T struct {
*   F int `yaml:"a, omitempty"`
*   B int
* }
* var t T
* yaml.Marshal(&T{B:2})   /// returns "b: 2 \n"
* yaml.Marshal(&T{F:1})   // returns "a: 1 \nb: 0\n"
 */
func MarshalYaml(in interface{}) (string, error) {
	b, err := yaml.Marshal(in)
	if err != nil {
		return "", err
	} else {
		return string(b), nil
	}
}

/**
* var t T
* yaml.Unmarshal([]byte("a: 1\nb: 2"), &t)
*
 */
func UnmarshalYaml(str string, out interface{}) error {
	return yaml.Unmarshal([]byte(str), out)
}
