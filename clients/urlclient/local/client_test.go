/*******************************************************************************
 * Copyright 2020 Dell Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package local

import (
	"testing"
)

var expectedURL = "http://domain.com"

func TestNew(t *testing.T) {
	actualClient := New(expectedURL)

	if actualClient == nil {
		t.Fatal("nil returned from NewLocalClient")
	}
}

func Test_URLPrefix(t *testing.T) {
	urlClient := New(expectedURL)

	actualURL, err := urlClient.Prefix()

	if err != nil {
		t.Fatalf("unexpected error %s", err.Error())
	}

	if actualURL != expectedURL {
		t.Fatalf("expected URL %s, found URL %s", expectedURL, actualURL)
	}
}
