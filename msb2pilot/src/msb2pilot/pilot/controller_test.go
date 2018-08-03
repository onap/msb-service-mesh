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
