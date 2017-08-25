/**
# Copyright 2015 Google Inc. All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
**/

package main

import (
	"cloud.google.com/go/compute/metadata"
	"testing"
)

func TestGCE(t *testing.T) {
	i := newInstance()
	if !metadata.OnGCE() && i.Error != "Not running on GCE" {
		t.Error("Test not running on GCE, but error does not indicate that fact.")
	}
}
