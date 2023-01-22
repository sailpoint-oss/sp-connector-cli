/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
)

// SodViolationContext1ConflictingAccessCriteriaLeftCriteria struct for SodViolationContext1ConflictingAccessCriteriaLeftCriteria
type SodViolationContext1ConflictingAccessCriteriaLeftCriteria struct {
	CriteriaList []SodExemptCriteria1 `json:"criteriaList,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContext1ConflictingAccessCriteriaLeftCriteria SodViolationContext1ConflictingAccessCriteriaLeftCriteria

// NewSodViolationContext1ConflictingAccessCriteriaLeftCriteria instantiates a new SodViolationContext1ConflictingAccessCriteriaLeftCriteria object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContext1ConflictingAccessCriteriaLeftCriteria() *SodViolationContext1ConflictingAccessCriteriaLeftCriteria {
	this := SodViolationContext1ConflictingAccessCriteriaLeftCriteria{}
	return &this
}

// NewSodViolationContext1ConflictingAccessCriteriaLeftCriteriaWithDefaults instantiates a new SodViolationContext1ConflictingAccessCriteriaLeftCriteria object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContext1ConflictingAccessCriteriaLeftCriteriaWithDefaults() *SodViolationContext1ConflictingAccessCriteriaLeftCriteria {
	this := SodViolationContext1ConflictingAccessCriteriaLeftCriteria{}
	return &this
}

// GetCriteriaList returns the CriteriaList field value if set, zero value otherwise.
func (o *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) GetCriteriaList() []SodExemptCriteria1 {
	if o == nil || isNil(o.CriteriaList) {
		var ret []SodExemptCriteria1
		return ret
	}
	return o.CriteriaList
}

// GetCriteriaListOk returns a tuple with the CriteriaList field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) GetCriteriaListOk() ([]SodExemptCriteria1, bool) {
	if o == nil || isNil(o.CriteriaList) {
		return nil, false
	}
	return o.CriteriaList, true
}

// HasCriteriaList returns a boolean if a field has been set.
func (o *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) HasCriteriaList() bool {
	if o != nil && !isNil(o.CriteriaList) {
		return true
	}

	return false
}

// SetCriteriaList gets a reference to the given []SodExemptCriteria1 and assigns it to the CriteriaList field.
func (o *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) SetCriteriaList(v []SodExemptCriteria1) {
	o.CriteriaList = v
}

func (o SodViolationContext1ConflictingAccessCriteriaLeftCriteria) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.CriteriaList) {
		toSerialize["criteriaList"] = o.CriteriaList
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationContext1ConflictingAccessCriteriaLeftCriteria := _SodViolationContext1ConflictingAccessCriteriaLeftCriteria{}

	if err = json.Unmarshal(bytes, &varSodViolationContext1ConflictingAccessCriteriaLeftCriteria); err == nil {
		*o = SodViolationContext1ConflictingAccessCriteriaLeftCriteria(varSodViolationContext1ConflictingAccessCriteriaLeftCriteria)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "criteriaList")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria struct {
	value *SodViolationContext1ConflictingAccessCriteriaLeftCriteria
	isSet bool
}

func (v NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) Get() *SodViolationContext1ConflictingAccessCriteriaLeftCriteria {
	return v.value
}

func (v *NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) Set(val *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria(val *SodViolationContext1ConflictingAccessCriteriaLeftCriteria) *NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria {
	return &NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria{value: val, isSet: true}
}

func (v NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContext1ConflictingAccessCriteriaLeftCriteria) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


