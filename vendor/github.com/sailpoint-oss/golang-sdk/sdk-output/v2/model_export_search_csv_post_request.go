/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
)

// ExportSearchCsvPostRequest struct for ExportSearchCsvPostRequest
type ExportSearchCsvPostRequest struct {
	Match *ExportSearchCsvPostRequestMatch `json:"match,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ExportSearchCsvPostRequest ExportSearchCsvPostRequest

// NewExportSearchCsvPostRequest instantiates a new ExportSearchCsvPostRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExportSearchCsvPostRequest() *ExportSearchCsvPostRequest {
	this := ExportSearchCsvPostRequest{}
	return &this
}

// NewExportSearchCsvPostRequestWithDefaults instantiates a new ExportSearchCsvPostRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExportSearchCsvPostRequestWithDefaults() *ExportSearchCsvPostRequest {
	this := ExportSearchCsvPostRequest{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *ExportSearchCsvPostRequest) GetMatch() ExportSearchCsvPostRequestMatch {
	if o == nil || isNil(o.Match) {
		var ret ExportSearchCsvPostRequestMatch
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExportSearchCsvPostRequest) GetMatchOk() (*ExportSearchCsvPostRequestMatch, bool) {
	if o == nil || isNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *ExportSearchCsvPostRequest) HasMatch() bool {
	if o != nil && !isNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given ExportSearchCsvPostRequestMatch and assigns it to the Match field.
func (o *ExportSearchCsvPostRequest) SetMatch(v ExportSearchCsvPostRequestMatch) {
	o.Match = &v
}

func (o ExportSearchCsvPostRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Match) {
		toSerialize["match"] = o.Match
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ExportSearchCsvPostRequest) UnmarshalJSON(bytes []byte) (err error) {
	varExportSearchCsvPostRequest := _ExportSearchCsvPostRequest{}

	if err = json.Unmarshal(bytes, &varExportSearchCsvPostRequest); err == nil {
		*o = ExportSearchCsvPostRequest(varExportSearchCsvPostRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "match")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableExportSearchCsvPostRequest struct {
	value *ExportSearchCsvPostRequest
	isSet bool
}

func (v NullableExportSearchCsvPostRequest) Get() *ExportSearchCsvPostRequest {
	return v.value
}

func (v *NullableExportSearchCsvPostRequest) Set(val *ExportSearchCsvPostRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableExportSearchCsvPostRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableExportSearchCsvPostRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExportSearchCsvPostRequest(val *ExportSearchCsvPostRequest) *NullableExportSearchCsvPostRequest {
	return &NullableExportSearchCsvPostRequest{value: val, isSet: true}
}

func (v NullableExportSearchCsvPostRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExportSearchCsvPostRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


