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
	"os"
	"path"
	"runtime"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/test"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestGetPropertyLabels(t *testing.T) {
	t.Parallel()
	ctx := context.Background()

	_, filename, _, _ := runtime.Caller(0)
	goldenPath := path.Join(path.Dir(filename), "get_property_labels")

	testSuite := func(mixer pb.MixerClient, latencyTest bool) {
		for _, c := range []struct {
			goldenFile string
			dcids      []string
		}{
			{
				"property-labels-class.json",
				[]string{"Class"},
			},
			{
				"property-labels-states.json",
				[]string{"geoId/05", "geoId/06"},
			},
		} {
			req := &pb.GetPropertyLabelsRequest{
				Dcids: c.dcids,
			}
			resp, err := mixer.GetPropertyLabels(ctx, req)
			if err != nil {
				t.Errorf("could not GetPropertyLabels: %s", err)
				continue
			}

			if latencyTest {
				continue
			}

			// Here the golden file is not same as the actual API output.
			// The actual payload is not a full serielized protobuf but
			// with the outer level ("data" field) removed. Here is to add that level
			// back so the payload can be unmarshaled.
			payload := "{\"data\":" + resp.GetPayload() + "}"
			var result pb.GetPropertyLabelsResponse
			if err := protojson.Unmarshal([]byte(payload), &result); err != nil {
				t.Errorf("Can not Unmarshal payload")
				continue
			}
			goldenFile := path.Join(goldenPath, c.goldenFile)
			if test.GenerateGolden {
				test.UpdateProtoGolden(&result, goldenPath, c.goldenFile)
				continue
			}
			var expected pb.GetPropertyLabelsResponse
			file, _ := os.ReadFile(goldenFile)
			if err := protojson.Unmarshal(file, &expected); err != nil {
				t.Errorf("Can not Unmarshal golden file %s: %v", goldenFile, err)
				continue
			}
			if diff := cmp.Diff(&result, &expected, protocmp.Transform()); diff != "" {
				t.Errorf("payload got diff: %v", diff)
				continue
			}
		}
	}

	if err := test.TestDriver(
		"GetPropertyLabels", &test.TestOption{}, testSuite); err != nil {
		t.Errorf("TestDriver() = %s", err)
	}
}
