// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package testutil

import (
	"os"
	"testing"

	"github.com/datacommonsorg/mixer/internal/parser/mcf"
	"github.com/datacommonsorg/mixer/internal/translator/types"
)

// DB is the database name.
const DB = "dc_v3"

// ReadTestMapping reads the testing schema mapping files into list of Mapping structs.
func ReadTestMapping(t *testing.T, files []string) []*types.Mapping {
	mappings := []*types.Mapping{}
	for _, f := range files {
		mappingStr, err := os.ReadFile(f)
		if err != nil {
			t.Fatalf("reading test schema mapping file: %s", err)
		}
		mapping, err := mcf.ParseMapping(string(mappingStr), DB)
		if err != nil {
			t.Fatalf("parsing test schema mapping file: %s", err)
		}
		mappings = append(mappings, mapping...)
	}
	return mappings
}
