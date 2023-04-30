// Copyright 2022 Google LLC
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

func TestVariableGroupInfo(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "variable_group_info")

	testSuite := func(mixer pb.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			node                string
			constrainedEntities []string
			goldenFile          string
		}{
			{
				"dc/g/Person_EducationalAttainment",
				[]string{"country/USA"},
				"school.json",
			},
			{
				"dc/g/Person_EnrollmentLevel-EnrolledInCollegeUndergraduateYears_Race",
				[]string{"country/USA"},
				"school_race.json",
			},
			{
				"dc/g/Demographics",
				[]string{"country/GBR"},
				"demographics_gbr.json",
			},
			// Run this first to test the server cache is not modified.
			{
				"dc/g/Root",
				[]string{"geoId/0649670"},
				"root_mtv.json",
			},
			// Run this first to test the server cache is not modified.
			{
				"dc/g/Root",
				[]string{"geoId/0649670", "country/JPN"},
				"root_mtv_jpn.json",
			},
			{
				"dc/g/Root",
				[]string{},
				"root.json",
			},
			{
				"dc/g/Demographics",
				[]string{},
				"demographics.json",
			},
			{
				"dc/g/Person_CitizenshipStatus-NotAUSCitizen_CorrectionalFacilityOperator-StateOperated&FederallyOperated&PrivatelyOperated",
				[]string{"geoId/0649670"},
				"citizenship.json",
			},
			{
				"g/Feeding_America",
				[]string{"geoId/06"},
				"private.json",
			},
			{
				"dc/g/Economy",
				[]string{"country/ASM"},
				"economy.json",
			},
			{
				"invalid,id",
				[]string{},
				"empty.json",
			},
		} {
			resp, err := mixer.VariableGroupInfo(ctx, &pb.VariableGroupInfoRequest{
				Node:                c.node,
				ConstrainedEntities: c.constrainedEntities,
			})
			if err != nil {
				t.Errorf("VariableGroupInfo() = %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			if test.GenerateGolden {
				test.UpdateProtoGolden(resp, goldenPath, c.goldenFile)
				continue
			}

			var expected pb.VariableGroupInfoResponse
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
		"VariableGroupInfo",
		&test.TestOption{UseCache: true, UseMemdb: true},
		testSuite,
	); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
