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

// PublicIdentity Details about a public identity
type PublicIdentity struct {
	// Identity id
	Id *string `json:"id,omitempty"`
	// Human-readable display name of identity.
	Name *string `json:"name,omitempty"`
	// Alternate unique identifier for the identity.
	Alias *string `json:"alias,omitempty"`
	// Email address of identity.
	Email NullableString `json:"email,omitempty"`
	// The lifecycle status for the identity
	Status NullableString `json:"status,omitempty"`
	Manager NullableIdentityReference `json:"manager,omitempty"`
	// The public identity attributes of the identity
	Attributes []IdentityAttribute `json:"attributes,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PublicIdentity PublicIdentity

// NewPublicIdentity instantiates a new PublicIdentity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPublicIdentity() *PublicIdentity {
	this := PublicIdentity{}
	return &this
}

// NewPublicIdentityWithDefaults instantiates a new PublicIdentity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPublicIdentityWithDefaults() *PublicIdentity {
	this := PublicIdentity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *PublicIdentity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *PublicIdentity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *PublicIdentity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PublicIdentity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PublicIdentity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PublicIdentity) SetName(v string) {
	o.Name = &v
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *PublicIdentity) GetAlias() string {
	if o == nil || isNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetAliasOk() (*string, bool) {
	if o == nil || isNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *PublicIdentity) HasAlias() bool {
	if o != nil && !isNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *PublicIdentity) SetAlias(v string) {
	o.Alias = &v
}

// GetEmail returns the Email field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIdentity) GetEmail() string {
	if o == nil || isNil(o.Email.Get()) {
		var ret string
		return ret
	}
	return *o.Email.Get()
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIdentity) GetEmailOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Email.Get(), o.Email.IsSet()
}

// HasEmail returns a boolean if a field has been set.
func (o *PublicIdentity) HasEmail() bool {
	if o != nil && o.Email.IsSet() {
		return true
	}

	return false
}

// SetEmail gets a reference to the given NullableString and assigns it to the Email field.
func (o *PublicIdentity) SetEmail(v string) {
	o.Email.Set(&v)
}
// SetEmailNil sets the value for Email to be an explicit nil
func (o *PublicIdentity) SetEmailNil() {
	o.Email.Set(nil)
}

// UnsetEmail ensures that no value is present for Email, not even an explicit nil
func (o *PublicIdentity) UnsetEmail() {
	o.Email.Unset()
}

// GetStatus returns the Status field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIdentity) GetStatus() string {
	if o == nil || isNil(o.Status.Get()) {
		var ret string
		return ret
	}
	return *o.Status.Get()
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIdentity) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Status.Get(), o.Status.IsSet()
}

// HasStatus returns a boolean if a field has been set.
func (o *PublicIdentity) HasStatus() bool {
	if o != nil && o.Status.IsSet() {
		return true
	}

	return false
}

// SetStatus gets a reference to the given NullableString and assigns it to the Status field.
func (o *PublicIdentity) SetStatus(v string) {
	o.Status.Set(&v)
}
// SetStatusNil sets the value for Status to be an explicit nil
func (o *PublicIdentity) SetStatusNil() {
	o.Status.Set(nil)
}

// UnsetStatus ensures that no value is present for Status, not even an explicit nil
func (o *PublicIdentity) UnsetStatus() {
	o.Status.Unset()
}

// GetManager returns the Manager field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *PublicIdentity) GetManager() IdentityReference {
	if o == nil || isNil(o.Manager.Get()) {
		var ret IdentityReference
		return ret
	}
	return *o.Manager.Get()
}

// GetManagerOk returns a tuple with the Manager field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *PublicIdentity) GetManagerOk() (*IdentityReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Manager.Get(), o.Manager.IsSet()
}

// HasManager returns a boolean if a field has been set.
func (o *PublicIdentity) HasManager() bool {
	if o != nil && o.Manager.IsSet() {
		return true
	}

	return false
}

// SetManager gets a reference to the given NullableIdentityReference and assigns it to the Manager field.
func (o *PublicIdentity) SetManager(v IdentityReference) {
	o.Manager.Set(&v)
}
// SetManagerNil sets the value for Manager to be an explicit nil
func (o *PublicIdentity) SetManagerNil() {
	o.Manager.Set(nil)
}

// UnsetManager ensures that no value is present for Manager, not even an explicit nil
func (o *PublicIdentity) UnsetManager() {
	o.Manager.Unset()
}

// GetAttributes returns the Attributes field value if set, zero value otherwise.
func (o *PublicIdentity) GetAttributes() []IdentityAttribute {
	if o == nil || isNil(o.Attributes) {
		var ret []IdentityAttribute
		return ret
	}
	return o.Attributes
}

// GetAttributesOk returns a tuple with the Attributes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PublicIdentity) GetAttributesOk() ([]IdentityAttribute, bool) {
	if o == nil || isNil(o.Attributes) {
		return nil, false
	}
	return o.Attributes, true
}

// HasAttributes returns a boolean if a field has been set.
func (o *PublicIdentity) HasAttributes() bool {
	if o != nil && !isNil(o.Attributes) {
		return true
	}

	return false
}

// SetAttributes gets a reference to the given []IdentityAttribute and assigns it to the Attributes field.
func (o *PublicIdentity) SetAttributes(v []IdentityAttribute) {
	o.Attributes = v
}

func (o PublicIdentity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if o.Email.IsSet() {
		toSerialize["email"] = o.Email.Get()
	}
	if o.Status.IsSet() {
		toSerialize["status"] = o.Status.Get()
	}
	if o.Manager.IsSet() {
		toSerialize["manager"] = o.Manager.Get()
	}
	if !isNil(o.Attributes) {
		toSerialize["attributes"] = o.Attributes
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *PublicIdentity) UnmarshalJSON(bytes []byte) (err error) {
	varPublicIdentity := _PublicIdentity{}

	if err = json.Unmarshal(bytes, &varPublicIdentity); err == nil {
		*o = PublicIdentity(varPublicIdentity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "alias")
		delete(additionalProperties, "email")
		delete(additionalProperties, "status")
		delete(additionalProperties, "manager")
		delete(additionalProperties, "attributes")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePublicIdentity struct {
	value *PublicIdentity
	isSet bool
}

func (v NullablePublicIdentity) Get() *PublicIdentity {
	return v.value
}

func (v *NullablePublicIdentity) Set(val *PublicIdentity) {
	v.value = val
	v.isSet = true
}

func (v NullablePublicIdentity) IsSet() bool {
	return v.isSet
}

func (v *NullablePublicIdentity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePublicIdentity(val *PublicIdentity) *NullablePublicIdentity {
	return &NullablePublicIdentity{value: val, isSet: true}
}

func (v NullablePublicIdentity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePublicIdentity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


