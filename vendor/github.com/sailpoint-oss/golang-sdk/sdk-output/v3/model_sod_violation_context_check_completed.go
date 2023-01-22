/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
)

// SodViolationContextCheckCompleted An object referencing a completed SOD violation check
type SodViolationContextCheckCompleted struct {
	// The status of SOD violation check
	State *string `json:"state,omitempty"`
	// The id of the Violation check event
	Uuid *string `json:"uuid,omitempty"`
	ViolationCheckResult *SodViolationCheckResult `json:"violationCheckResult,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationContextCheckCompleted SodViolationContextCheckCompleted

// NewSodViolationContextCheckCompleted instantiates a new SodViolationContextCheckCompleted object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationContextCheckCompleted() *SodViolationContextCheckCompleted {
	this := SodViolationContextCheckCompleted{}
	return &this
}

// NewSodViolationContextCheckCompletedWithDefaults instantiates a new SodViolationContextCheckCompleted object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationContextCheckCompletedWithDefaults() *SodViolationContextCheckCompleted {
	this := SodViolationContextCheckCompleted{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *SodViolationContextCheckCompleted) SetState(v string) {
	o.State = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted) GetUuid() string {
	if o == nil || isNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted) GetUuidOk() (*string, bool) {
	if o == nil || isNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasUuid() bool {
	if o != nil && !isNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SodViolationContextCheckCompleted) SetUuid(v string) {
	o.Uuid = &v
}

// GetViolationCheckResult returns the ViolationCheckResult field value if set, zero value otherwise.
func (o *SodViolationContextCheckCompleted) GetViolationCheckResult() SodViolationCheckResult {
	if o == nil || isNil(o.ViolationCheckResult) {
		var ret SodViolationCheckResult
		return ret
	}
	return *o.ViolationCheckResult
}

// GetViolationCheckResultOk returns a tuple with the ViolationCheckResult field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationContextCheckCompleted) GetViolationCheckResultOk() (*SodViolationCheckResult, bool) {
	if o == nil || isNil(o.ViolationCheckResult) {
		return nil, false
	}
	return o.ViolationCheckResult, true
}

// HasViolationCheckResult returns a boolean if a field has been set.
func (o *SodViolationContextCheckCompleted) HasViolationCheckResult() bool {
	if o != nil && !isNil(o.ViolationCheckResult) {
		return true
	}

	return false
}

// SetViolationCheckResult gets a reference to the given SodViolationCheckResult and assigns it to the ViolationCheckResult field.
func (o *SodViolationContextCheckCompleted) SetViolationCheckResult(v SodViolationCheckResult) {
	o.ViolationCheckResult = &v
}

func (o SodViolationContextCheckCompleted) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !isNil(o.ViolationCheckResult) {
		toSerialize["violationCheckResult"] = o.ViolationCheckResult
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SodViolationContextCheckCompleted) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationContextCheckCompleted := _SodViolationContextCheckCompleted{}

	if err = json.Unmarshal(bytes, &varSodViolationContextCheckCompleted); err == nil {
		*o = SodViolationContextCheckCompleted(varSodViolationContextCheckCompleted)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "uuid")
		delete(additionalProperties, "violationCheckResult")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationContextCheckCompleted struct {
	value *SodViolationContextCheckCompleted
	isSet bool
}

func (v NullableSodViolationContextCheckCompleted) Get() *SodViolationContextCheckCompleted {
	return v.value
}

func (v *NullableSodViolationContextCheckCompleted) Set(val *SodViolationContextCheckCompleted) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationContextCheckCompleted) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationContextCheckCompleted) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationContextCheckCompleted(val *SodViolationContextCheckCompleted) *NullableSodViolationContextCheckCompleted {
	return &NullableSodViolationContextCheckCompleted{value: val, isSet: true}
}

func (v NullableSodViolationContextCheckCompleted) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationContextCheckCompleted) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


