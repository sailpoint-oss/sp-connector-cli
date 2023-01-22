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

// SodViolationCheckResult The inner object representing the completed SOD Violation check
type SodViolationCheckResult struct {
	Message *ErrorMessageDto `json:"message,omitempty"`
	// Arbitrary key-value pairs. They will never be processed by the IdentityNow system but will be returned on completion of the violation check.
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	ViolationContexts []SodViolationContext `json:"violationContexts,omitempty"`
	// A list of the Policies that were violated
	ViolatedPolicies []BaseReferenceDto `json:"violatedPolicies,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _SodViolationCheckResult SodViolationCheckResult

// NewSodViolationCheckResult instantiates a new SodViolationCheckResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSodViolationCheckResult() *SodViolationCheckResult {
	this := SodViolationCheckResult{}
	return &this
}

// NewSodViolationCheckResultWithDefaults instantiates a new SodViolationCheckResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSodViolationCheckResultWithDefaults() *SodViolationCheckResult {
	this := SodViolationCheckResult{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *SodViolationCheckResult) GetMessage() ErrorMessageDto {
	if o == nil || isNil(o.Message) {
		var ret ErrorMessageDto
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult) GetMessageOk() (*ErrorMessageDto, bool) {
	if o == nil || isNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *SodViolationCheckResult) HasMessage() bool {
	if o != nil && !isNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given ErrorMessageDto and assigns it to the Message field.
func (o *SodViolationCheckResult) SetMessage(v ErrorMessageDto) {
	o.Message = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *SodViolationCheckResult) GetClientMetadata() map[string]string {
	if o == nil || isNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *SodViolationCheckResult) HasClientMetadata() bool {
	if o != nil && !isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *SodViolationCheckResult) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetViolationContexts returns the ViolationContexts field value if set, zero value otherwise.
func (o *SodViolationCheckResult) GetViolationContexts() []SodViolationContext {
	if o == nil || isNil(o.ViolationContexts) {
		var ret []SodViolationContext
		return ret
	}
	return o.ViolationContexts
}

// GetViolationContextsOk returns a tuple with the ViolationContexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult) GetViolationContextsOk() ([]SodViolationContext, bool) {
	if o == nil || isNil(o.ViolationContexts) {
		return nil, false
	}
	return o.ViolationContexts, true
}

// HasViolationContexts returns a boolean if a field has been set.
func (o *SodViolationCheckResult) HasViolationContexts() bool {
	if o != nil && !isNil(o.ViolationContexts) {
		return true
	}

	return false
}

// SetViolationContexts gets a reference to the given []SodViolationContext and assigns it to the ViolationContexts field.
func (o *SodViolationCheckResult) SetViolationContexts(v []SodViolationContext) {
	o.ViolationContexts = v
}

// GetViolatedPolicies returns the ViolatedPolicies field value if set, zero value otherwise.
func (o *SodViolationCheckResult) GetViolatedPolicies() []BaseReferenceDto {
	if o == nil || isNil(o.ViolatedPolicies) {
		var ret []BaseReferenceDto
		return ret
	}
	return o.ViolatedPolicies
}

// GetViolatedPoliciesOk returns a tuple with the ViolatedPolicies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SodViolationCheckResult) GetViolatedPoliciesOk() ([]BaseReferenceDto, bool) {
	if o == nil || isNil(o.ViolatedPolicies) {
		return nil, false
	}
	return o.ViolatedPolicies, true
}

// HasViolatedPolicies returns a boolean if a field has been set.
func (o *SodViolationCheckResult) HasViolatedPolicies() bool {
	if o != nil && !isNil(o.ViolatedPolicies) {
		return true
	}

	return false
}

// SetViolatedPolicies gets a reference to the given []BaseReferenceDto and assigns it to the ViolatedPolicies field.
func (o *SodViolationCheckResult) SetViolatedPolicies(v []BaseReferenceDto) {
	o.ViolatedPolicies = v
}

func (o SodViolationCheckResult) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !isNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if !isNil(o.ViolationContexts) {
		toSerialize["violationContexts"] = o.ViolationContexts
	}
	if !isNil(o.ViolatedPolicies) {
		toSerialize["violatedPolicies"] = o.ViolatedPolicies
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *SodViolationCheckResult) UnmarshalJSON(bytes []byte) (err error) {
	varSodViolationCheckResult := _SodViolationCheckResult{}

	if err = json.Unmarshal(bytes, &varSodViolationCheckResult); err == nil {
		*o = SodViolationCheckResult(varSodViolationCheckResult)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "message")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "violationContexts")
		delete(additionalProperties, "violatedPolicies")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableSodViolationCheckResult struct {
	value *SodViolationCheckResult
	isSet bool
}

func (v NullableSodViolationCheckResult) Get() *SodViolationCheckResult {
	return v.value
}

func (v *NullableSodViolationCheckResult) Set(val *SodViolationCheckResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSodViolationCheckResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSodViolationCheckResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSodViolationCheckResult(val *SodViolationCheckResult) *NullableSodViolationCheckResult {
	return &NullableSodViolationCheckResult{value: val, isSet: true}
}

func (v NullableSodViolationCheckResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSodViolationCheckResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


