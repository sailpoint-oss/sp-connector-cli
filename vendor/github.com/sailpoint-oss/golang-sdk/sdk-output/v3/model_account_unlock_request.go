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

// AccountUnlockRequest Request used for account unlock
type AccountUnlockRequest struct {
	// If set, an external process validates that the user wants to proceed with this request.
	ExternalVerificationId *string `json:"externalVerificationId,omitempty"`
	// If set, the IDN account is unlocked after the workflow completes.
	UnlockIDNAccount *bool `json:"unlockIDNAccount,omitempty"`
	// If set, provisioning updates the account attribute at the source.   This option is used when the account is not synced to ensure the attribute is updated.
	ForceProvisioning *bool `json:"forceProvisioning,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccountUnlockRequest AccountUnlockRequest

// NewAccountUnlockRequest instantiates a new AccountUnlockRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccountUnlockRequest() *AccountUnlockRequest {
	this := AccountUnlockRequest{}
	return &this
}

// NewAccountUnlockRequestWithDefaults instantiates a new AccountUnlockRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccountUnlockRequestWithDefaults() *AccountUnlockRequest {
	this := AccountUnlockRequest{}
	return &this
}

// GetExternalVerificationId returns the ExternalVerificationId field value if set, zero value otherwise.
func (o *AccountUnlockRequest) GetExternalVerificationId() string {
	if o == nil || isNil(o.ExternalVerificationId) {
		var ret string
		return ret
	}
	return *o.ExternalVerificationId
}

// GetExternalVerificationIdOk returns a tuple with the ExternalVerificationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUnlockRequest) GetExternalVerificationIdOk() (*string, bool) {
	if o == nil || isNil(o.ExternalVerificationId) {
		return nil, false
	}
	return o.ExternalVerificationId, true
}

// HasExternalVerificationId returns a boolean if a field has been set.
func (o *AccountUnlockRequest) HasExternalVerificationId() bool {
	if o != nil && !isNil(o.ExternalVerificationId) {
		return true
	}

	return false
}

// SetExternalVerificationId gets a reference to the given string and assigns it to the ExternalVerificationId field.
func (o *AccountUnlockRequest) SetExternalVerificationId(v string) {
	o.ExternalVerificationId = &v
}

// GetUnlockIDNAccount returns the UnlockIDNAccount field value if set, zero value otherwise.
func (o *AccountUnlockRequest) GetUnlockIDNAccount() bool {
	if o == nil || isNil(o.UnlockIDNAccount) {
		var ret bool
		return ret
	}
	return *o.UnlockIDNAccount
}

// GetUnlockIDNAccountOk returns a tuple with the UnlockIDNAccount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUnlockRequest) GetUnlockIDNAccountOk() (*bool, bool) {
	if o == nil || isNil(o.UnlockIDNAccount) {
		return nil, false
	}
	return o.UnlockIDNAccount, true
}

// HasUnlockIDNAccount returns a boolean if a field has been set.
func (o *AccountUnlockRequest) HasUnlockIDNAccount() bool {
	if o != nil && !isNil(o.UnlockIDNAccount) {
		return true
	}

	return false
}

// SetUnlockIDNAccount gets a reference to the given bool and assigns it to the UnlockIDNAccount field.
func (o *AccountUnlockRequest) SetUnlockIDNAccount(v bool) {
	o.UnlockIDNAccount = &v
}

// GetForceProvisioning returns the ForceProvisioning field value if set, zero value otherwise.
func (o *AccountUnlockRequest) GetForceProvisioning() bool {
	if o == nil || isNil(o.ForceProvisioning) {
		var ret bool
		return ret
	}
	return *o.ForceProvisioning
}

// GetForceProvisioningOk returns a tuple with the ForceProvisioning field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccountUnlockRequest) GetForceProvisioningOk() (*bool, bool) {
	if o == nil || isNil(o.ForceProvisioning) {
		return nil, false
	}
	return o.ForceProvisioning, true
}

// HasForceProvisioning returns a boolean if a field has been set.
func (o *AccountUnlockRequest) HasForceProvisioning() bool {
	if o != nil && !isNil(o.ForceProvisioning) {
		return true
	}

	return false
}

// SetForceProvisioning gets a reference to the given bool and assigns it to the ForceProvisioning field.
func (o *AccountUnlockRequest) SetForceProvisioning(v bool) {
	o.ForceProvisioning = &v
}

func (o AccountUnlockRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.ExternalVerificationId) {
		toSerialize["externalVerificationId"] = o.ExternalVerificationId
	}
	if !isNil(o.UnlockIDNAccount) {
		toSerialize["unlockIDNAccount"] = o.UnlockIDNAccount
	}
	if !isNil(o.ForceProvisioning) {
		toSerialize["forceProvisioning"] = o.ForceProvisioning
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccountUnlockRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAccountUnlockRequest := _AccountUnlockRequest{}

	if err = json.Unmarshal(bytes, &varAccountUnlockRequest); err == nil {
		*o = AccountUnlockRequest(varAccountUnlockRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "externalVerificationId")
		delete(additionalProperties, "unlockIDNAccount")
		delete(additionalProperties, "forceProvisioning")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccountUnlockRequest struct {
	value *AccountUnlockRequest
	isSet bool
}

func (v NullableAccountUnlockRequest) Get() *AccountUnlockRequest {
	return v.value
}

func (v *NullableAccountUnlockRequest) Set(val *AccountUnlockRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccountUnlockRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccountUnlockRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccountUnlockRequest(val *AccountUnlockRequest) *NullableAccountUnlockRequest {
	return &NullableAccountUnlockRequest{value: val, isSet: true}
}

func (v NullableAccountUnlockRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccountUnlockRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


