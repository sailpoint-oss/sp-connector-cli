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

// BasicAuthConfig Config required if BASIC_AUTH is used.
type BasicAuthConfig struct {
	// The username to authenticate.
	UserName *string `json:"userName,omitempty"`
	// The password to authenticate. On response, this field is set to null as to not return secrets.
	Password NullableString `json:"password,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _BasicAuthConfig BasicAuthConfig

// NewBasicAuthConfig instantiates a new BasicAuthConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBasicAuthConfig() *BasicAuthConfig {
	this := BasicAuthConfig{}
	return &this
}

// NewBasicAuthConfigWithDefaults instantiates a new BasicAuthConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBasicAuthConfigWithDefaults() *BasicAuthConfig {
	this := BasicAuthConfig{}
	return &this
}

// GetUserName returns the UserName field value if set, zero value otherwise.
func (o *BasicAuthConfig) GetUserName() string {
	if o == nil || isNil(o.UserName) {
		var ret string
		return ret
	}
	return *o.UserName
}

// GetUserNameOk returns a tuple with the UserName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BasicAuthConfig) GetUserNameOk() (*string, bool) {
	if o == nil || isNil(o.UserName) {
		return nil, false
	}
	return o.UserName, true
}

// HasUserName returns a boolean if a field has been set.
func (o *BasicAuthConfig) HasUserName() bool {
	if o != nil && !isNil(o.UserName) {
		return true
	}

	return false
}

// SetUserName gets a reference to the given string and assigns it to the UserName field.
func (o *BasicAuthConfig) SetUserName(v string) {
	o.UserName = &v
}

// GetPassword returns the Password field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *BasicAuthConfig) GetPassword() string {
	if o == nil || isNil(o.Password.Get()) {
		var ret string
		return ret
	}
	return *o.Password.Get()
}

// GetPasswordOk returns a tuple with the Password field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BasicAuthConfig) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Password.Get(), o.Password.IsSet()
}

// HasPassword returns a boolean if a field has been set.
func (o *BasicAuthConfig) HasPassword() bool {
	if o != nil && o.Password.IsSet() {
		return true
	}

	return false
}

// SetPassword gets a reference to the given NullableString and assigns it to the Password field.
func (o *BasicAuthConfig) SetPassword(v string) {
	o.Password.Set(&v)
}
// SetPasswordNil sets the value for Password to be an explicit nil
func (o *BasicAuthConfig) SetPasswordNil() {
	o.Password.Set(nil)
}

// UnsetPassword ensures that no value is present for Password, not even an explicit nil
func (o *BasicAuthConfig) UnsetPassword() {
	o.Password.Unset()
}

func (o BasicAuthConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.UserName) {
		toSerialize["userName"] = o.UserName
	}
	if o.Password.IsSet() {
		toSerialize["password"] = o.Password.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *BasicAuthConfig) UnmarshalJSON(bytes []byte) (err error) {
	varBasicAuthConfig := _BasicAuthConfig{}

	if err = json.Unmarshal(bytes, &varBasicAuthConfig); err == nil {
		*o = BasicAuthConfig(varBasicAuthConfig)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "userName")
		delete(additionalProperties, "password")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableBasicAuthConfig struct {
	value *BasicAuthConfig
	isSet bool
}

func (v NullableBasicAuthConfig) Get() *BasicAuthConfig {
	return v.value
}

func (v *NullableBasicAuthConfig) Set(val *BasicAuthConfig) {
	v.value = val
	v.isSet = true
}

func (v NullableBasicAuthConfig) IsSet() bool {
	return v.isSet
}

func (v *NullableBasicAuthConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBasicAuthConfig(val *BasicAuthConfig) *NullableBasicAuthConfig {
	return &NullableBasicAuthConfig{value: val, isSet: true}
}

func (v NullableBasicAuthConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBasicAuthConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


