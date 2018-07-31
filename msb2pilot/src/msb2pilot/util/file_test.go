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
	"testing"
)

func TestWrite(t *testing.T) {
	cases := []struct {
		path, data, want string
		mode             os.FileMode
	}{
		{
			path: `test.txt`,
			data: `test string`,
			mode: 0666,
			want: `success`,
		},
		{
			path: ``,
			data: `test string`,
			mode: 0666,
			want: `fail`,
		},
	}

	for _, cas := range cases {
		err := Write(cas.path, cas.data, cas.mode)
		if (cas.want == "success" && err != nil) || (cas.want == "fail" && err == nil) {
			t.Errorf("Write() => got %v, want %s", err, cas.want)
		}
	}
}

func TestRead(t *testing.T) {

	cases := []struct {
		path, want string
	}{
		{
			path: `file_test.go`,
			want: `success`,
		},
		{
			path: ``,
			want: `fail`,
		},
	}

	for _, cas := range cases {
		_, err := Read(cas.path)
		if (cas.want == "success" && err != nil) || (cas.want == "fail" && err == nil) {
			t.Errorf("Read() => got %v, want %s", err, cas.want)
		}
	}
}
