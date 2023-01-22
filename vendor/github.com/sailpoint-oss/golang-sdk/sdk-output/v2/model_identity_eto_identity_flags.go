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

// IdentityEtoIdentityFlags struct for IdentityEtoIdentityFlags
type IdentityEtoIdentityFlags struct {
	RESPONSIVE_LAUNCHPAD *string `json:"RESPONSIVE_LAUNCHPAD,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IdentityEtoIdentityFlags IdentityEtoIdentityFlags

// NewIdentityEtoIdentityFlags instantiates a new IdentityEtoIdentityFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdentityEtoIdentityFlags() *IdentityEtoIdentityFlags {
	this := IdentityEtoIdentityFlags{}
	return &this
}

// NewIdentityEtoIdentityFlagsWithDefaults instantiates a new IdentityEtoIdentityFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdentityEtoIdentityFlagsWithDefaults() *IdentityEtoIdentityFlags {
	this := IdentityEtoIdentityFlags{}
	return &this
}

// GetRESPONSIVE_LAUNCHPAD returns the RESPONSIVE_LAUNCHPAD field value if set, zero value otherwise.
func (o *IdentityEtoIdentityFlags) GetRESPONSIVE_LAUNCHPAD() string {
	if o == nil || isNil(o.RESPONSIVE_LAUNCHPAD) {
		var ret string
		return ret
	}
	return *o.RESPONSIVE_LAUNCHPAD
}

// GetRESPONSIVE_LAUNCHPADOk returns a tuple with the RESPONSIVE_LAUNCHPAD field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdentityEtoIdentityFlags) GetRESPONSIVE_LAUNCHPADOk() (*string, bool) {
	if o == nil || isNil(o.RESPONSIVE_LAUNCHPAD) {
		return nil, false
	}
	return o.RESPONSIVE_LAUNCHPAD, true
}

// HasRESPONSIVE_LAUNCHPAD returns a boolean if a field has been set.
func (o *IdentityEtoIdentityFlags) HasRESPONSIVE_LAUNCHPAD() bool {
	if o != nil && !isNil(o.RESPONSIVE_LAUNCHPAD) {
		return true
	}

	return false
}

// SetRESPONSIVE_LAUNCHPAD gets a reference to the given string and assigns it to the RESPONSIVE_LAUNCHPAD field.
func (o *IdentityEtoIdentityFlags) SetRESPONSIVE_LAUNCHPAD(v string) {
	o.RESPONSIVE_LAUNCHPAD = &v
}

func (o IdentityEtoIdentityFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.RESPONSIVE_LAUNCHPAD) {
		toSerialize["RESPONSIVE_LAUNCHPAD"] = o.RESPONSIVE_LAUNCHPAD
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *IdentityEtoIdentityFlags) UnmarshalJSON(bytes []byte) (err error) {
	varIdentityEtoIdentityFlags := _IdentityEtoIdentityFlags{}

	if err = json.Unmarshal(bytes, &varIdentityEtoIdentityFlags); err == nil {
		*o = IdentityEtoIdentityFlags(varIdentityEtoIdentityFlags)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "RESPONSIVE_LAUNCHPAD")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIdentityEtoIdentityFlags struct {
	value *IdentityEtoIdentityFlags
	isSet bool
}

func (v NullableIdentityEtoIdentityFlags) Get() *IdentityEtoIdentityFlags {
	return v.value
}

func (v *NullableIdentityEtoIdentityFlags) Set(val *IdentityEtoIdentityFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableIdentityEtoIdentityFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableIdentityEtoIdentityFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdentityEtoIdentityFlags(val *IdentityEtoIdentityFlags) *NullableIdentityEtoIdentityFlags {
	return &NullableIdentityEtoIdentityFlags{value: val, isSet: true}
}

func (v NullableIdentityEtoIdentityFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdentityEtoIdentityFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


