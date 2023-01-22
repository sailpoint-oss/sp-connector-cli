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

// TriggerInputIdentityAttributesChanged struct for TriggerInputIdentityAttributesChanged
type TriggerInputIdentityAttributesChanged struct {
	Identity TriggerInputIdentityAttributesChangedIdentity `json:"identity"`
	// A list of one or more identity attributes that changed on the identity.
	Changes []TriggerInputIdentityAttributesChangedChangesInner `json:"changes"`
	AdditionalProperties map[string]interface{}
}

type _TriggerInputIdentityAttributesChanged TriggerInputIdentityAttributesChanged

// NewTriggerInputIdentityAttributesChanged instantiates a new TriggerInputIdentityAttributesChanged object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerInputIdentityAttributesChanged(identity TriggerInputIdentityAttributesChangedIdentity, changes []TriggerInputIdentityAttributesChangedChangesInner) *TriggerInputIdentityAttributesChanged {
	this := TriggerInputIdentityAttributesChanged{}
	this.Identity = identity
	this.Changes = changes
	return &this
}

// NewTriggerInputIdentityAttributesChangedWithDefaults instantiates a new TriggerInputIdentityAttributesChanged object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerInputIdentityAttributesChangedWithDefaults() *TriggerInputIdentityAttributesChanged {
	this := TriggerInputIdentityAttributesChanged{}
	return &this
}

// GetIdentity returns the Identity field value
func (o *TriggerInputIdentityAttributesChanged) GetIdentity() TriggerInputIdentityAttributesChangedIdentity {
	if o == nil {
		var ret TriggerInputIdentityAttributesChangedIdentity
		return ret
	}

	return o.Identity
}

// GetIdentityOk returns a tuple with the Identity field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityAttributesChanged) GetIdentityOk() (*TriggerInputIdentityAttributesChangedIdentity, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Identity, true
}

// SetIdentity sets field value
func (o *TriggerInputIdentityAttributesChanged) SetIdentity(v TriggerInputIdentityAttributesChangedIdentity) {
	o.Identity = v
}

// GetChanges returns the Changes field value
func (o *TriggerInputIdentityAttributesChanged) GetChanges() []TriggerInputIdentityAttributesChangedChangesInner {
	if o == nil {
		var ret []TriggerInputIdentityAttributesChangedChangesInner
		return ret
	}

	return o.Changes
}

// GetChangesOk returns a tuple with the Changes field value
// and a boolean to check if the value has been set.
func (o *TriggerInputIdentityAttributesChanged) GetChangesOk() ([]TriggerInputIdentityAttributesChangedChangesInner, bool) {
	if o == nil {
		return nil, false
	}
	return o.Changes, true
}

// SetChanges sets field value
func (o *TriggerInputIdentityAttributesChanged) SetChanges(v []TriggerInputIdentityAttributesChangedChangesInner) {
	o.Changes = v
}

func (o TriggerInputIdentityAttributesChanged) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["identity"] = o.Identity
	}
	if true {
		toSerialize["changes"] = o.Changes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *TriggerInputIdentityAttributesChanged) UnmarshalJSON(bytes []byte) (err error) {
	varTriggerInputIdentityAttributesChanged := _TriggerInputIdentityAttributesChanged{}

	if err = json.Unmarshal(bytes, &varTriggerInputIdentityAttributesChanged); err == nil {
		*o = TriggerInputIdentityAttributesChanged(varTriggerInputIdentityAttributesChanged)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "identity")
		delete(additionalProperties, "changes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableTriggerInputIdentityAttributesChanged struct {
	value *TriggerInputIdentityAttributesChanged
	isSet bool
}

func (v NullableTriggerInputIdentityAttributesChanged) Get() *TriggerInputIdentityAttributesChanged {
	return v.value
}

func (v *NullableTriggerInputIdentityAttributesChanged) Set(val *TriggerInputIdentityAttributesChanged) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerInputIdentityAttributesChanged) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerInputIdentityAttributesChanged) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerInputIdentityAttributesChanged(val *TriggerInputIdentityAttributesChanged) *NullableTriggerInputIdentityAttributesChanged {
	return &NullableTriggerInputIdentityAttributesChanged{value: val, isSet: true}
}

func (v NullableTriggerInputIdentityAttributesChanged) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerInputIdentityAttributesChanged) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


