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
	"testing"
)

type T struct {
	F int `yaml:"a,omitempty"`
	B int
}

func TestMarshalYaml(t *testing.T) {

	cases := []struct {
		in   *T
		want string
	}{
		{
			in: &T{
				1, 2,
			},
			want: "a: 1\nb: 2\n",
		}, {
			in: &T{
				B: 2,
			},
			want: "b: 2\n",
		}}

	for _, cas := range cases {
		got, err := MarshalYaml(cas.in)
		if err != nil {
			t.Errorf(err.Error())
		}

		if got != cas.want {
			t.Errorf("MarshalYaml error: want %s, got %s", cas.want, got)
		}

	}

}

func TestUnmarshalYaml(t *testing.T) {
	cases := []struct {
		in   string
		want T
	}{
		{
			in: "a: 1\nb: 2",
			want: T{
				1, 2,
			},
		},
		{
			in: "b: 2\n",
			want: T{
				B: 2,
			},
		},
		{
			in:   "c: 2\n",
			want: T{},
		}}

	for _, cas := range cases {
		got := new(T)
		err := UnmarshalYaml(cas.in, got)

		if err != nil {
			t.Errorf("UnmarshalYaml error: want err, got %v", cas.want)
		}

		if got.F != cas.want.F || got.B != cas.want.B {
			t.Errorf("UnmarshalYaml error: want %v, got %v", cas.want, got)
		}
	}
}
