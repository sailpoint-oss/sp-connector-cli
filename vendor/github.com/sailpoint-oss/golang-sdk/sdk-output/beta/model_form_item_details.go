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

// FormItemDetails struct for FormItemDetails
type FormItemDetails struct {
	// Name of the FormItem
	Name *string `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FormItemDetails FormItemDetails

// NewFormItemDetails instantiates a new FormItemDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFormItemDetails() *FormItemDetails {
	this := FormItemDetails{}
	return &this
}

// NewFormItemDetailsWithDefaults instantiates a new FormItemDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFormItemDetailsWithDefaults() *FormItemDetails {
	this := FormItemDetails{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FormItemDetails) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FormItemDetails) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FormItemDetails) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FormItemDetails) SetName(v string) {
	o.Name = &v
}

func (o FormItemDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *FormItemDetails) UnmarshalJSON(bytes []byte) (err error) {
	varFormItemDetails := _FormItemDetails{}

	if err = json.Unmarshal(bytes, &varFormItemDetails); err == nil {
		*o = FormItemDetails(varFormItemDetails)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFormItemDetails struct {
	value *FormItemDetails
	isSet bool
}

func (v NullableFormItemDetails) Get() *FormItemDetails {
	return v.value
}

func (v *NullableFormItemDetails) Set(val *FormItemDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableFormItemDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableFormItemDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFormItemDetails(val *FormItemDetails) *NullableFormItemDetails {
	return &NullableFormItemDetails{value: val, isSet: true}
}

func (v NullableFormItemDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFormItemDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


