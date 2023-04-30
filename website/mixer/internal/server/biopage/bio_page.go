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

package biopage

import (
	"context"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/store"
	"github.com/datacommonsorg/mixer/internal/store/bigtable"
	"google.golang.org/protobuf/proto"
)

// GetBioPageDataHelper is a wrapper to get bio page data.
func GetBioPageDataHelper(
	ctx context.Context,
	dcid string,
	store *store.Store,
) (*pb.GraphNodes, error) {
	btDataList, err := bigtable.Read(
		ctx,
		store.BtGroup,
		bigtable.BtBioPagePrefix,
		[][]string{{dcid}},
		func(jsonRaw []byte) (interface{}, error) {
			var graph pb.GraphNodes
			if err := proto.Unmarshal(jsonRaw, &graph); err != nil {
				return nil, err
			}
			return &graph, nil
		},
	)
	if err != nil {
		return nil, err
	}
	// btData is orderred by preference. Use the one that has bio data without
	// merging.
	for _, btData := range btDataList {
		for _, row := range btData {
			if row.Parts[0] == dcid {
				return row.Data.(*pb.GraphNodes), nil
			}
		}
	}
	return &pb.GraphNodes{}, nil
}
