// Copyright 2021 Google LLC
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
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetStatVarSummary(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "get_statvar_summary")

	testSuite := func(mixer pb.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			svs        []string
			goldenFile string
		}{
			{
				[]string{"Count_CriminalActivities_CombinedCrime"},
				"total_crimes.json",
			},
			{
				[]string{"UnemploymentRate_Person"},
				"unemployment.json",
			},
			{
				[]string{"Count_Person_Female", "Count_Person_Female_AsianAlone", "FertilityRate_Person_Female"},
				"females.json",
			},
			{
				[]string{"IncrementalCount_MedicalConditionIncident_COVID_19_ConfirmedOrProbableCase", "CumulativeCount_MedicalConditionIncident_COVID_19_ConfirmedOrProbableCase"},
				"covid.json",
			},
		} {
			resp, err := mixer.GetStatVarSummary(ctx, &pb.GetStatVarSummaryRequest{
				StatVars: c.svs,
			})
			if err != nil {
				t.Errorf("could not GetStatVarSummary: %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			if test.GenerateGolden {
				test.UpdateProtoGolden(resp, goldenPath, c.goldenFile)
				continue
			}

			var expected pb.GetStatVarSummaryResponse
			if err = test.ReadJSON(goldenPath, c.goldenFile, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file")
				continue
			}

			if diff := cmp.Diff(resp, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("payload got diff: %v", diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"GetStatVarSummary", &test.TestOption{}, testSuite); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
