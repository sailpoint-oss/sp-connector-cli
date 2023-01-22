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

// WorkflowLibraryFormFields struct for WorkflowLibraryFormFields
type WorkflowLibraryFormFields struct {
	// Describes the form field in the UI
	HelpText *string `json:"helpText,omitempty"`
	// A human readable name for this form field in the UI
	Label *string `json:"label,omitempty"`
	// The name of the input attribute
	Name *string `json:"name,omitempty"`
	// Denotes if this field is a required attribute
	Required *bool `json:"required,omitempty"`
	// The type of the form field
	Type map[string]interface{} `json:"type,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WorkflowLibraryFormFields WorkflowLibraryFormFields

// NewWorkflowLibraryFormFields instantiates a new WorkflowLibraryFormFields object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowLibraryFormFields() *WorkflowLibraryFormFields {
	this := WorkflowLibraryFormFields{}
	return &this
}

// NewWorkflowLibraryFormFieldsWithDefaults instantiates a new WorkflowLibraryFormFields object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowLibraryFormFieldsWithDefaults() *WorkflowLibraryFormFields {
	this := WorkflowLibraryFormFields{}
	return &this
}

// GetHelpText returns the HelpText field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetHelpText() string {
	if o == nil || isNil(o.HelpText) {
		var ret string
		return ret
	}
	return *o.HelpText
}

// GetHelpTextOk returns a tuple with the HelpText field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetHelpTextOk() (*string, bool) {
	if o == nil || isNil(o.HelpText) {
		return nil, false
	}
	return o.HelpText, true
}

// HasHelpText returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasHelpText() bool {
	if o != nil && !isNil(o.HelpText) {
		return true
	}

	return false
}

// SetHelpText gets a reference to the given string and assigns it to the HelpText field.
func (o *WorkflowLibraryFormFields) SetHelpText(v string) {
	o.HelpText = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetLabel() string {
	if o == nil || isNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetLabelOk() (*string, bool) {
	if o == nil || isNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasLabel() bool {
	if o != nil && !isNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *WorkflowLibraryFormFields) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *WorkflowLibraryFormFields) SetName(v string) {
	o.Name = &v
}

// GetRequired returns the Required field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetRequired() bool {
	if o == nil || isNil(o.Required) {
		var ret bool
		return ret
	}
	return *o.Required
}

// GetRequiredOk returns a tuple with the Required field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.Required) {
		return nil, false
	}
	return o.Required, true
}

// HasRequired returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasRequired() bool {
	if o != nil && !isNil(o.Required) {
		return true
	}

	return false
}

// SetRequired gets a reference to the given bool and assigns it to the Required field.
func (o *WorkflowLibraryFormFields) SetRequired(v bool) {
	o.Required = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *WorkflowLibraryFormFields) GetType() map[string]interface{} {
	if o == nil || isNil(o.Type) {
		var ret map[string]interface{}
		return ret
	}
	return o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WorkflowLibraryFormFields) GetTypeOk() (map[string]interface{}, bool) {
	if o == nil || isNil(o.Type) {
		return map[string]interface{}{}, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *WorkflowLibraryFormFields) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given map[string]interface{} and assigns it to the Type field.
func (o *WorkflowLibraryFormFields) SetType(v map[string]interface{}) {
	o.Type = v
}

func (o WorkflowLibraryFormFields) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.HelpText) {
		toSerialize["helpText"] = o.HelpText
	}
	if !isNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Required) {
		toSerialize["required"] = o.Required
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *WorkflowLibraryFormFields) UnmarshalJSON(bytes []byte) (err error) {
	varWorkflowLibraryFormFields := _WorkflowLibraryFormFields{}

	if err = json.Unmarshal(bytes, &varWorkflowLibraryFormFields); err == nil {
		*o = WorkflowLibraryFormFields(varWorkflowLibraryFormFields)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "helpText")
		delete(additionalProperties, "label")
		delete(additionalProperties, "name")
		delete(additionalProperties, "required")
		delete(additionalProperties, "type")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWorkflowLibraryFormFields struct {
	value *WorkflowLibraryFormFields
	isSet bool
}

func (v NullableWorkflowLibraryFormFields) Get() *WorkflowLibraryFormFields {
	return v.value
}

func (v *NullableWorkflowLibraryFormFields) Set(val *WorkflowLibraryFormFields) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowLibraryFormFields) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowLibraryFormFields) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowLibraryFormFields(val *WorkflowLibraryFormFields) *NullableWorkflowLibraryFormFields {
	return &NullableWorkflowLibraryFormFields{value: val, isSet: true}
}

func (v NullableWorkflowLibraryFormFields) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowLibraryFormFields) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


