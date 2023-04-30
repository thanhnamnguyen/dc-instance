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

package server

import (
	"context"
	"testing"

	pb "github.com/datacommonsorg/mixer/internal/proto"
	"github.com/datacommonsorg/mixer/internal/store"
)

func TestNoBigTable(t *testing.T) {
	ctx := context.Background()
	store, err := store.NewStore(nil, nil, nil, "", nil)
	if err != nil {
		t.Fatalf("store.NewStore() = %v", err)
	}
	s := NewMixerServer(store, nil, nil, nil)
	_, err = s.PlacePage(ctx, &pb.PlacePageRequest{
		Node:     "geoId/06",
		Category: "Overview",
	})
	if err.Error() != "rpc error: code = NotFound desc = Bigtable instance is not specified" {
		t.Errorf("Error invalid: %s", err)
	}
}
