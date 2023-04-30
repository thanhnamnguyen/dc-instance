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

package propertyvalue

import (
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/google/go-cmp/cmp"
	"google.golang.org/protobuf/testing/protocmp"
)

func TestFilterEntities(t *testing.T) {

	input := []*pb.EntityInfo{
		{
			Dcid:  "dcid1",
			Value: "1",
			Types: []string{"City"},
		},
		{
			Dcid:  "dcid2",
			Value: "2",
			Types: []string{"City"},
		},
		{
			Dcid:  "dcid3",
			Value: "3",
			Types: []string{"County"},
		},
		{
			Dcid:  "dcid4",
			Value: "4",
			Types: []string{"County"},
		},
	}

	for _, c := range []struct {
		typ   string
		limit int
		want  []*pb.EntityInfo
	}{
		{
			"",
			1,
			[]*pb.EntityInfo{
				{
					Dcid:  "dcid1",
					Value: "1",
					Types: []string{"City"},
				},
			},
		},
		{
			"City",
			0,
			[]*pb.EntityInfo{
				{
					Dcid:  "dcid1",
					Value: "1",
					Types: []string{"City"},
				},
				{
					Dcid:  "dcid2",
					Value: "2",
					Types: []string{"City"},
				},
			},
		},
		{
			"City",
			1,
			[]*pb.EntityInfo{
				{
					Dcid:  "dcid1",
					Value: "1",
					Types: []string{"City"},
				},
			},
		},
	} {
		got := filterEntities(input, c.typ, c.limit)
		if diff := cmp.Diff(got, c.want, protocmp.Transform()); diff != "" {
			t.Errorf("trimeNodes() got diff %v", diff)
		}
	}
}
