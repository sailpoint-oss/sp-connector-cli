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

// CreateWorkgroupRequest struct for CreateWorkgroupRequest
type CreateWorkgroupRequest struct {
	Name *string `json:"name,omitempty"`
	Description *string `json:"description,omitempty"`
	Owner *CreateWorkgroupRequestOwner `json:"owner,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CreateWorkgroupRequest CreateWorkgroupRequest

// NewCreateWorkgroupRequest instantiates a new CreateWorkgroupRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCreateWorkgroupRequest() *CreateWorkgroupRequest {
	this := CreateWorkgroupRequest{}
	return &this
}

// NewCreateWorkgroupRequestWithDefaults instantiates a new CreateWorkgroupRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCreateWorkgroupRequestWithDefaults() *CreateWorkgroupRequest {
	this := CreateWorkgroupRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CreateWorkgroupRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkgroupRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CreateWorkgroupRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CreateWorkgroupRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *CreateWorkgroupRequest) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkgroupRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *CreateWorkgroupRequest) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *CreateWorkgroupRequest) SetDescription(v string) {
	o.Description = &v
}

// GetOwner returns the Owner field value if set, zero value otherwise.
func (o *CreateWorkgroupRequest) GetOwner() CreateWorkgroupRequestOwner {
	if o == nil || isNil(o.Owner) {
		var ret CreateWorkgroupRequestOwner
		return ret
	}
	return *o.Owner
}

// GetOwnerOk returns a tuple with the Owner field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CreateWorkgroupRequest) GetOwnerOk() (*CreateWorkgroupRequestOwner, bool) {
	if o == nil || isNil(o.Owner) {
		return nil, false
	}
	return o.Owner, true
}

// HasOwner returns a boolean if a field has been set.
func (o *CreateWorkgroupRequest) HasOwner() bool {
	if o != nil && !isNil(o.Owner) {
		return true
	}

	return false
}

// SetOwner gets a reference to the given CreateWorkgroupRequestOwner and assigns it to the Owner field.
func (o *CreateWorkgroupRequest) SetOwner(v CreateWorkgroupRequestOwner) {
	o.Owner = &v
}

func (o CreateWorkgroupRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Owner) {
		toSerialize["owner"] = o.Owner
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CreateWorkgroupRequest) UnmarshalJSON(bytes []byte) (err error) {
	varCreateWorkgroupRequest := _CreateWorkgroupRequest{}

	if err = json.Unmarshal(bytes, &varCreateWorkgroupRequest); err == nil {
		*o = CreateWorkgroupRequest(varCreateWorkgroupRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "owner")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCreateWorkgroupRequest struct {
	value *CreateWorkgroupRequest
	isSet bool
}

func (v NullableCreateWorkgroupRequest) Get() *CreateWorkgroupRequest {
	return v.value
}

func (v *NullableCreateWorkgroupRequest) Set(val *CreateWorkgroupRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCreateWorkgroupRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCreateWorkgroupRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCreateWorkgroupRequest(val *CreateWorkgroupRequest) *NullableCreateWorkgroupRequest {
	return &NullableCreateWorkgroupRequest{value: val, isSet: true}
}

func (v NullableCreateWorkgroupRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCreateWorkgroupRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


