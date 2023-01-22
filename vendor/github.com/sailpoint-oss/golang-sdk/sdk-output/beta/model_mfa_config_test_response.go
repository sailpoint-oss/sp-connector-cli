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

// MfaConfigTestResponse Response model for configuration test of a given MFA method
type MfaConfigTestResponse struct {
	// The configuration test result.
	State *string `json:"state,omitempty"`
	// The error message to indicate the failure of configuration test.
	Error *string `json:"error,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _MfaConfigTestResponse MfaConfigTestResponse

// NewMfaConfigTestResponse instantiates a new MfaConfigTestResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMfaConfigTestResponse() *MfaConfigTestResponse {
	this := MfaConfigTestResponse{}
	return &this
}

// NewMfaConfigTestResponseWithDefaults instantiates a new MfaConfigTestResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMfaConfigTestResponseWithDefaults() *MfaConfigTestResponse {
	this := MfaConfigTestResponse{}
	return &this
}

// GetState returns the State field value if set, zero value otherwise.
func (o *MfaConfigTestResponse) GetState() string {
	if o == nil || isNil(o.State) {
		var ret string
		return ret
	}
	return *o.State
}

// GetStateOk returns a tuple with the State field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfigTestResponse) GetStateOk() (*string, bool) {
	if o == nil || isNil(o.State) {
		return nil, false
	}
	return o.State, true
}

// HasState returns a boolean if a field has been set.
func (o *MfaConfigTestResponse) HasState() bool {
	if o != nil && !isNil(o.State) {
		return true
	}

	return false
}

// SetState gets a reference to the given string and assigns it to the State field.
func (o *MfaConfigTestResponse) SetState(v string) {
	o.State = &v
}

// GetError returns the Error field value if set, zero value otherwise.
func (o *MfaConfigTestResponse) GetError() string {
	if o == nil || isNil(o.Error) {
		var ret string
		return ret
	}
	return *o.Error
}

// GetErrorOk returns a tuple with the Error field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MfaConfigTestResponse) GetErrorOk() (*string, bool) {
	if o == nil || isNil(o.Error) {
		return nil, false
	}
	return o.Error, true
}

// HasError returns a boolean if a field has been set.
func (o *MfaConfigTestResponse) HasError() bool {
	if o != nil && !isNil(o.Error) {
		return true
	}

	return false
}

// SetError gets a reference to the given string and assigns it to the Error field.
func (o *MfaConfigTestResponse) SetError(v string) {
	o.Error = &v
}

func (o MfaConfigTestResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.State) {
		toSerialize["state"] = o.State
	}
	if !isNil(o.Error) {
		toSerialize["error"] = o.Error
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *MfaConfigTestResponse) UnmarshalJSON(bytes []byte) (err error) {
	varMfaConfigTestResponse := _MfaConfigTestResponse{}

	if err = json.Unmarshal(bytes, &varMfaConfigTestResponse); err == nil {
		*o = MfaConfigTestResponse(varMfaConfigTestResponse)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "state")
		delete(additionalProperties, "error")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableMfaConfigTestResponse struct {
	value *MfaConfigTestResponse
	isSet bool
}

func (v NullableMfaConfigTestResponse) Get() *MfaConfigTestResponse {
	return v.value
}

func (v *NullableMfaConfigTestResponse) Set(val *MfaConfigTestResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMfaConfigTestResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMfaConfigTestResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMfaConfigTestResponse(val *MfaConfigTestResponse) *NullableMfaConfigTestResponse {
	return &NullableMfaConfigTestResponse{value: val, isSet: true}
}

func (v NullableMfaConfigTestResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMfaConfigTestResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


