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
package models

import (
	"reflect"
	"testing"
)

func TestConvertBaseInfo(t *testing.T) {
	cases := []struct{ in, want string }{
		{
			in:   `{"enable_ssl":"true", "version":"v1","protocol":"REST","publish_port":"28012|28013","url":"/api/itm-pmadaptor/v1","is_manual":"false","visualRange":"0","appversion":"v1.18.20.04"}`,
			want: `v1`,
		},
		{
			in:   `{"enable_ssl":"true", "version":"v2","protocol":"UI", "status":"1", "url":"/api/itm-pmadaptor/v1","is_manual":"false"}`,
			want: `v2`,
		},
		{
			in:   `{"others":"other"}`,
			want: ``,
		},
		{
			in:   ``,
			want: ``,
		},
	}

	for _, cas := range cases {
		got, _ := convertBaseInfo(cas.in)
		if got.Version != cas.want {
			t.Errorf("convertBaseInfo(%s) => got %s, want %s", cas.in, got.AppVersion, cas.want)
		}
	}

}

func TestConvertNameSpace(t *testing.T) {
	cases := []struct{ in, want string }{
		{
			in:   `{"namespace":"test"}`,
			want: `test`,
		},
		{
			in:   `{"namespace":"testwithother", "others":"other"}`,
			want: `testwithother`,
		},
		{
			in:   `{"others":"other"}`,
			want: ``,
		},
		{
			in:   ``,
			want: ``,
		},
	}

	for _, cas := range cases {
		got, _ := convertNameSpace(cas.in)
		if got.NameSpace != cas.want {
			t.Errorf("convertNameSpace(%s) => got %s, want %s", cas.in, got.NameSpace, cas.want)
		}
	}

}

func TestConvertMsbLabel(t *testing.T) {
	cases := []struct{ label, in, want string }{
		{
			label: "base",
			in:    `"base":{"enable_ssl":"true", "version":"v1","protocol":"REST","publish_port":"28012|28013","url":"/api/itm-pmadaptor/v1","is_manual":"false","visualRange":"0","appversion":"v1.18.20.04"}`,
			want:  `*models.BaseInfo`,
		},
		{
			label: "ns",
			in:    `"ns":{"namespace":"test"}`,
			want:  `*models.NameSpace`,
		},
		{
			label: "others",
			in:    `{"others":"other"}`,
			want:  ``,
		},
		{
			label: "",
			in:    ``,
			want:  ``,
		},
	}

	for _, cas := range cases {
		got := convertMsbLabel(cas.label, cas.in)

		if got == nil {
			if cas.want != "" {
				t.Errorf("convertMsbLabel(%s, %s) => got nil, want %s", cas.label, cas.in, cas.want)
			}
		} else {
			if reflect.TypeOf(got).String() != cas.want {
				t.Errorf("convertMsbLabel(%s, %s) => got %v, want %s", cas.label, cas.in, reflect.TypeOf(got), cas.want)
			}
		}
	}

}
