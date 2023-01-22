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

// ServiceDeskIntegrationTemplateDtoAllOf This is the model for a Service Desk integration template, used to create and edit Service Desk Integrations.
type ServiceDeskIntegrationTemplateDtoAllOf struct {
	// The 'type' property specifies the type of the Service Desk integration template.
	Type string `json:"type"`
	// The 'attributes' property value is a map of attributes available for integrations using this Service Desk integration template.
	Attributes map[string]interface{} `json:"attributes"`
	ProvisioningConfig ProvisioningConfig `json:"provisioningConfig"`
	AdditionalProperties map[string]interface{}
}

type _ServiceDeskIntegrationTemplateDtoAllOf ServiceDeskIntegrationTemplateDtoAllOf

// NewServiceDeskIntegrationTemplateDtoAllOf instantiates a new ServiceDeskIntegrationTemplateDtoAllOf object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServiceDeskIntegrationTemplateDtoAllOf(type_ string, attributes map[string]interface{}, provisioningConfig ProvisioningConfig) *ServiceDeskIntegrationTemplateDtoAllOf {
	this := ServiceDeskIntegrationTemplateDtoAllOf{}
	this.Type = type_
	this.Attributes = attributes
	this.ProvisioningConfig = provisioningConfig
	return &this
}

// NewServiceDeskIntegrationTemplateDtoAllOfWithDefaults instantiates a new ServiceDeskIntegrationTemplateDtoAllOf object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServiceDeskIntegrationTemplateDtoAllOfWithDefaults() *ServiceDeskIntegrationTemplateDtoAllOf {
	this := ServiceDeskIntegrationTemplateDtoAllOf{}
	var type_ string = "Web Service SDIM"
	this.Type = type_
	return &this
}

// GetType returns the Type field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetType(v string) {
	o.Type = v
}

// GetAttributes returns the Attributes field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetAttributes() map[string]interface{} {
	if o == nil {
		var ret map[string]interface{}
		return ret
	}

	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetAttributesOk() (map[string]interface{}, bool) {
	if o == nil {
		return map[string]interface{}{}, false
	}
	return o.Attributes, true
}

// SetAttributes sets field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetAttributes(v map[string]interface{}) {
	o.Attributes = v
}

// GetProvisioningConfig returns the ProvisioningConfig field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetProvisioningConfig() ProvisioningConfig {
	if o == nil {
		var ret ProvisioningConfig
		return ret
	}

	return o.ProvisioningConfig
}

// GetProvisioningConfigOk returns a tuple with the ProvisioningConfig field value
// and a boolean to check if the value has been set.
func (o *ServiceDeskIntegrationTemplateDtoAllOf) GetProvisioningConfigOk() (*ProvisioningConfig, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProvisioningConfig, true
}

// SetProvisioningConfig sets field value
func (o *ServiceDeskIntegrationTemplateDtoAllOf) SetProvisioningConfig(v ProvisioningConfig) {
	o.ProvisioningConfig = v
}

func (o ServiceDeskIntegrationTemplateDtoAllOf) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["attributes"] = o.Attributes
	}
	if true {
		toSerialize["provisioningConfig"] = o.ProvisioningConfig
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ServiceDeskIntegrationTemplateDtoAllOf) UnmarshalJSON(bytes []byte) (err error) {
	varServiceDeskIntegrationTemplateDtoAllOf := _ServiceDeskIntegrationTemplateDtoAllOf{}

	if err = json.Unmarshal(bytes, &varServiceDeskIntegrationTemplateDtoAllOf); err == nil {
		*o = ServiceDeskIntegrationTemplateDtoAllOf(varServiceDeskIntegrationTemplateDtoAllOf)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "type")
		delete(additionalProperties, "attributes")
		delete(additionalProperties, "provisioningConfig")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableServiceDeskIntegrationTemplateDtoAllOf struct {
	value *ServiceDeskIntegrationTemplateDtoAllOf
	isSet bool
}

func (v NullableServiceDeskIntegrationTemplateDtoAllOf) Get() *ServiceDeskIntegrationTemplateDtoAllOf {
	return v.value
}

func (v *NullableServiceDeskIntegrationTemplateDtoAllOf) Set(val *ServiceDeskIntegrationTemplateDtoAllOf) {
	v.value = val
	v.isSet = true
}

func (v NullableServiceDeskIntegrationTemplateDtoAllOf) IsSet() bool {
	return v.isSet
}

func (v *NullableServiceDeskIntegrationTemplateDtoAllOf) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServiceDeskIntegrationTemplateDtoAllOf(val *ServiceDeskIntegrationTemplateDtoAllOf) *NullableServiceDeskIntegrationTemplateDtoAllOf {
	return &NullableServiceDeskIntegrationTemplateDtoAllOf{value: val, isSet: true}
}

func (v NullableServiceDeskIntegrationTemplateDtoAllOf) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServiceDeskIntegrationTemplateDtoAllOf) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


