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

// ScheduleRecipientsInner struct for ScheduleRecipientsInner
type ScheduleRecipientsInner struct {
	// The type of object being referenced
	Type string `json:"type"`
	// The ID of the referenced object
	Id string `json:"id"`
	AdditionalProperties map[string]interface{}
}

type _ScheduleRecipientsInner ScheduleRecipientsInner

// NewScheduleRecipientsInner instantiates a new ScheduleRecipientsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleRecipientsInner(type_ string, id string) *ScheduleRecipientsInner {
	this := ScheduleRecipientsInner{}
	this.Type = type_
	this.Id = id
	return &this
}

// NewScheduleRecipientsInnerWithDefaults instantiates a new ScheduleRecipientsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleRecipientsInnerWithDefaults() *ScheduleRecipientsInner {
	this := ScheduleRecipientsInner{}
	return &this
}

// GetType returns the Type field value
func (o *ScheduleRecipientsInner) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ScheduleRecipientsInner) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ScheduleRecipientsInner) SetType(v string) {
	o.Type = v
}

// GetId returns the Id field value
func (o *ScheduleRecipientsInner) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ScheduleRecipientsInner) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ScheduleRecipientsInner) SetId(v string) {
	o.Id = v
}

func (o ScheduleRecipientsInner) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["id"] = o.Id
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ScheduleRecipientsInner) UnmarshalJSON(bytes []byte) (err error) {
	varScheduleRecipientsInner := _ScheduleRecipientsInner{}

	if err = json.Unmarshal(bytes, &varScheduleRecipientsInner); err == nil {
		*o = ScheduleRecipientsInner(varScheduleRecipientsInner)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "id")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableScheduleRecipientsInner struct {
	value *ScheduleRecipientsInner
	isSet bool
}

func (v NullableScheduleRecipientsInner) Get() *ScheduleRecipientsInner {
	return v.value
}

func (v *NullableScheduleRecipientsInner) Set(val *ScheduleRecipientsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleRecipientsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleRecipientsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleRecipientsInner(val *ScheduleRecipientsInner) *NullableScheduleRecipientsInner {
	return &NullableScheduleRecipientsInner{value: val, isSet: true}
}

func (v NullableScheduleRecipientsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleRecipientsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


