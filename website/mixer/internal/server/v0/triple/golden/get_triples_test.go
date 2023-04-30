// Copyright 2020 Google LLC
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

package golden

import (
	"context"
	"encoding/json"
	"os"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/server/model"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
)

func TestGetTriples(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "get_triples")

	testSuite := func(mixer pb.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			dcids        []string
			goldenFile   string
			partialMatch bool
			limit        int32
			count        []int
		}{
			{
				[]string{"State", "Country"},
				"place_type.json",
				false,
				-1,
				nil,
			},
			{
				[]string{"zip/00603"},
				"place.json",
				true,
				-1,
				nil,
			},
			{
				[]string{
					"dc/o/w2z8nx9y43k97", // LifeExpectancy_Person_Female
					"dc/o/88cs3xqnmpp55", // Count_Person<CensusPEPSurvey>
					"dc/o/23gt9k7fql176", // Count_Person<dcAggregate/CensusACS5yrSurvey>
					"dc/o/kyv7dxe4s18eh", // Count_Person<>
				},
				"observation.json",
				false,
				-1,
				nil,
			},
			{
				[]string{"Count_Person", "Count_Person_Female"},
				"stat_var.json",
				false,
				-1,
				nil,
			},
			{
				[]string{"City", "County"},
				"limit.json",
				false,
				5,
				[]int{5, 5},
			},
			{
				[]string{"dc/base/HumanReadableStatVars"},
				"limit1.json",
				false,
				1,
				[]int{8, 8},
			},
		} {
			req := &pb.GetTriplesRequest{Dcids: c.dcids}
			if c.limit > 0 {
				req.Limit = c.limit
			}
			resp, err := mixer.GetTriples(ctx, req)
			if err != nil {
				t.Errorf("could not GetTriples: %s", err)
				continue
			}
			var result map[string][]*model.Triple
			if err = json.Unmarshal([]byte(resp.GetPayload()), &result); err != nil {
				t.Errorf("Can not Unmarshal payload, %v", err)
				continue
			}

			if latencyTest {
				continue
			}

			goldenFile := path.Join(goldenPath, c.goldenFile)
			if test.GenerateGolden && c.goldenFile != "" {
				test.UpdateGolden(result, goldenPath, c.goldenFile)
				continue
			}

			if c.limit > 0 {
				for idx, dcid := range c.dcids {
					count := len(result[dcid])
					if count < c.count[idx] {
						t.Errorf(
							"Len of triples for %s expect %d, got %d",
							dcid, c.count[idx], count)
					}
				}
				continue
			}

			var expected map[string][]*model.Triple
			file, _ := os.ReadFile(goldenFile)
			err = json.Unmarshal(file, &expected)
			if err != nil {
				t.Errorf("Can not Unmarshal golden file %s: %v", c.goldenFile, err)
				continue
			}
			if diff := cmp.Diff(result, expected); diff != "" {
				t.Errorf("payload from %s got diff: %v", goldenFile, diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"GetTriples", &test.TestOption{}, testSuite); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
