/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"fmt"
)

// Index Enum representing the currently supported indices. Additional values may be added in the future without notice.
type Index string

// List of Index
const (
	INDEX_ACCESSPROFILES Index = "accessprofiles"
	INDEX_ACCOUNTACTIVITIES Index = "accountactivities"
	INDEX_ENTITLEMENTS Index = "entitlements"
	INDEX_EVENTS Index = "events"
	INDEX_IDENTITIES Index = "identities"
	INDEX_ROLES Index = "roles"
	INDEX_STAR Index = "*"
)

// All allowed values of Index enum
var AllowedIndexEnumValues = []Index{
	"accessprofiles",
	"accountactivities",
	"entitlements",
	"events",
	"identities",
	"roles",
	"*",
}

func (v *Index) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Index(value)
	for _, existing := range AllowedIndexEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Index", value)
}

// NewIndexFromValue returns a pointer to a valid Index
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewIndexFromValue(v string) (*Index, error) {
	ev := Index(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for Index: valid values are %v", v, AllowedIndexEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v Index) IsValid() bool {
	for _, existing := range AllowedIndexEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Index value
func (v Index) Ptr() *Index {
	return &v
}

type NullableIndex struct {
	value *Index
	isSet bool
}

func (v NullableIndex) Get() *Index {
	return v.value
}

func (v *NullableIndex) Set(val *Index) {
	v.value = val
	v.isSet = true
}

func (v NullableIndex) IsSet() bool {
	return v.isSet
}

func (v *NullableIndex) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIndex(val *Index) *NullableIndex {
	return &NullableIndex{value: val, isSet: true}
}

func (v NullableIndex) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIndex) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}

