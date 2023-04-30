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

package model

// Triple represents a triples entry in the BT triples cache.
type Triple struct {
	SubjectID    string   `json:"subjectId,omitempty"`
	SubjectName  string   `json:"subjectName,omitempty"`
	SubjectTypes []string `json:"subjectTypes,omitempty"`
	Predicate    string   `json:"predicate,omitempty"`
	ObjectID     string   `json:"objectId,omitempty"`
	ObjectName   string   `json:"objectName,omitempty"`
	ObjectValue  string   `json:"objectValue,omitempty"`
	ObjectTypes  []string `json:"objectTypes,omitempty"`
	ProvenanceID string   `json:"provenanceId,omitempty"`
}

// StatisticalVariable contains key info of population and observation.
type StatisticalVariable struct {
	PopType                string            `json:"popType,omitempty"`
	PVs                    map[string]string `json:"pvs,omitempty"`
	MeasuredProp           string            `json:"measuredProp,omitempty"`
	MeasurementMethod      string            `json:"measurementMethod,omitempty"`
	MeasurementDenominator string            `json:"measurementDeonominator,omitempty"`
	MeasurementQualifier   string            `json:"measurementQualifier,omitempty"`
	ScalingFactor          string            `json:"scalingFactor,omitempty"`
	Unit                   string            `json:"unit,omitempty"`
	StatType               string            `json:"statType,omitempty"`
}

// PlacePopInfo contains basic info for a place and a population.
type PlacePopInfo struct {
	PlaceID      string `json:"dcid,omitempty"`
	PopulationID string `json:"population,omitempty"`
}

// SourceSeries represents time series data for a particular source.
type SourceSeries struct {
	ImportName        string             `json:"importName,omitempty"`
	ObservationPeriod string             `json:"observationPeriod,omitempty"`
	MeasurementMethod string             `json:"measurementMethod,omitempty"`
	ScalingFactor     string             `json:"scalingFactor,omitempty"`
	Unit              string             `json:"unit,omitempty"`
	ProvenanceURL     string             `json:"provenanceUrl,omitempty"`
	Val               map[string]float64 `json:"val,omitempty"`
}

// ObsTimeSeries repesents multiple time series data.
type ObsTimeSeries struct {
	PlaceName    string          `json:"placeName,omitempty"`
	PlaceDcid    string          `json:"placeDcid,omitempty"`
	SourceSeries []*SourceSeries `json:"sourceSeries,omitempty"`
}

// GetStatsResponse holds the information
type GetStatsResponse struct {
	Data          map[string]float64 `json:"data,omitempty"`
	PlaceName     string             `json:"placeName,omitempty"`
	ProvenanceURL string             `json:"provenanceUrl,omitempty"`
}

// StatObsProp represents properties for a StatVarObservation.
type StatObsProp struct {
	MeasurementMethod string
	ObservationPeriod string
	Unit              string
	ScalingFactor     string
}
