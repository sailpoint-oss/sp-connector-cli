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

// AttributeDefinition struct for AttributeDefinition
type AttributeDefinition struct {
	// The name of the attribute.
	Name *string `json:"name,omitempty"`
	Type *AttributeDefinitionType `json:"type,omitempty"`
	Schema *AttributeDefinitionSchema `json:"schema,omitempty"`
	// A human-readable description of the attribute.
	Description *string `json:"description,omitempty"`
	// Flag indicating whether or not the attribute is multi-valued.
	IsMulti *bool `json:"isMulti,omitempty"`
	// Flag indicating whether or not the attribute is an entitlement.
	IsEntitlement *bool `json:"isEntitlement,omitempty"`
	// Flag indicating whether or not the attribute represents a group. This can only be `true` if `isEntitlement` is also `true` **and** there is a schema defined for the attribute.. 
	IsGroup *bool `json:"isGroup,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AttributeDefinition AttributeDefinition

// NewAttributeDefinition instantiates a new AttributeDefinition object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAttributeDefinition() *AttributeDefinition {
	this := AttributeDefinition{}
	return &this
}

// NewAttributeDefinitionWithDefaults instantiates a new AttributeDefinition object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAttributeDefinitionWithDefaults() *AttributeDefinition {
	this := AttributeDefinition{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AttributeDefinition) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AttributeDefinition) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AttributeDefinition) SetName(v string) {
	o.Name = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AttributeDefinition) GetType() AttributeDefinitionType {
	if o == nil || isNil(o.Type) {
		var ret AttributeDefinitionType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetTypeOk() (*AttributeDefinitionType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AttributeDefinition) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given AttributeDefinitionType and assigns it to the Type field.
func (o *AttributeDefinition) SetType(v AttributeDefinitionType) {
	o.Type = &v
}

// GetSchema returns the Schema field value if set, zero value otherwise.
func (o *AttributeDefinition) GetSchema() AttributeDefinitionSchema {
	if o == nil || isNil(o.Schema) {
		var ret AttributeDefinitionSchema
		return ret
	}
	return *o.Schema
}

// GetSchemaOk returns a tuple with the Schema field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetSchemaOk() (*AttributeDefinitionSchema, bool) {
	if o == nil || isNil(o.Schema) {
		return nil, false
	}
	return o.Schema, true
}

// HasSchema returns a boolean if a field has been set.
func (o *AttributeDefinition) HasSchema() bool {
	if o != nil && !isNil(o.Schema) {
		return true
	}

	return false
}

// SetSchema gets a reference to the given AttributeDefinitionSchema and assigns it to the Schema field.
func (o *AttributeDefinition) SetSchema(v AttributeDefinitionSchema) {
	o.Schema = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *AttributeDefinition) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *AttributeDefinition) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *AttributeDefinition) SetDescription(v string) {
	o.Description = &v
}

// GetIsMulti returns the IsMulti field value if set, zero value otherwise.
func (o *AttributeDefinition) GetIsMulti() bool {
	if o == nil || isNil(o.IsMulti) {
		var ret bool
		return ret
	}
	return *o.IsMulti
}

// GetIsMultiOk returns a tuple with the IsMulti field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetIsMultiOk() (*bool, bool) {
	if o == nil || isNil(o.IsMulti) {
		return nil, false
	}
	return o.IsMulti, true
}

// HasIsMulti returns a boolean if a field has been set.
func (o *AttributeDefinition) HasIsMulti() bool {
	if o != nil && !isNil(o.IsMulti) {
		return true
	}

	return false
}

// SetIsMulti gets a reference to the given bool and assigns it to the IsMulti field.
func (o *AttributeDefinition) SetIsMulti(v bool) {
	o.IsMulti = &v
}

// GetIsEntitlement returns the IsEntitlement field value if set, zero value otherwise.
func (o *AttributeDefinition) GetIsEntitlement() bool {
	if o == nil || isNil(o.IsEntitlement) {
		var ret bool
		return ret
	}
	return *o.IsEntitlement
}

// GetIsEntitlementOk returns a tuple with the IsEntitlement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetIsEntitlementOk() (*bool, bool) {
	if o == nil || isNil(o.IsEntitlement) {
		return nil, false
	}
	return o.IsEntitlement, true
}

// HasIsEntitlement returns a boolean if a field has been set.
func (o *AttributeDefinition) HasIsEntitlement() bool {
	if o != nil && !isNil(o.IsEntitlement) {
		return true
	}

	return false
}

// SetIsEntitlement gets a reference to the given bool and assigns it to the IsEntitlement field.
func (o *AttributeDefinition) SetIsEntitlement(v bool) {
	o.IsEntitlement = &v
}

// GetIsGroup returns the IsGroup field value if set, zero value otherwise.
func (o *AttributeDefinition) GetIsGroup() bool {
	if o == nil || isNil(o.IsGroup) {
		var ret bool
		return ret
	}
	return *o.IsGroup
}

// GetIsGroupOk returns a tuple with the IsGroup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AttributeDefinition) GetIsGroupOk() (*bool, bool) {
	if o == nil || isNil(o.IsGroup) {
		return nil, false
	}
	return o.IsGroup, true
}

// HasIsGroup returns a boolean if a field has been set.
func (o *AttributeDefinition) HasIsGroup() bool {
	if o != nil && !isNil(o.IsGroup) {
		return true
	}

	return false
}

// SetIsGroup gets a reference to the given bool and assigns it to the IsGroup field.
func (o *AttributeDefinition) SetIsGroup(v bool) {
	o.IsGroup = &v
}

func (o AttributeDefinition) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Schema) {
		toSerialize["schema"] = o.Schema
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.IsMulti) {
		toSerialize["isMulti"] = o.IsMulti
	}
	if !isNil(o.IsEntitlement) {
		toSerialize["isEntitlement"] = o.IsEntitlement
	}
	if !isNil(o.IsGroup) {
		toSerialize["isGroup"] = o.IsGroup
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AttributeDefinition) UnmarshalJSON(bytes []byte) (err error) {
	varAttributeDefinition := _AttributeDefinition{}

	if err = json.Unmarshal(bytes, &varAttributeDefinition); err == nil {
		*o = AttributeDefinition(varAttributeDefinition)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "type")
		delete(additionalProperties, "schema")
		delete(additionalProperties, "description")
		delete(additionalProperties, "isMulti")
		delete(additionalProperties, "isEntitlement")
		delete(additionalProperties, "isGroup")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAttributeDefinition struct {
	value *AttributeDefinition
	isSet bool
}

func (v NullableAttributeDefinition) Get() *AttributeDefinition {
	return v.value
}

func (v *NullableAttributeDefinition) Set(val *AttributeDefinition) {
	v.value = val
	v.isSet = true
}

func (v NullableAttributeDefinition) IsSet() bool {
	return v.isSet
}

func (v *NullableAttributeDefinition) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAttributeDefinition(val *AttributeDefinition) *NullableAttributeDefinition {
	return &NullableAttributeDefinition{value: val, isSet: true}
}

func (v NullableAttributeDefinition) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAttributeDefinition) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


