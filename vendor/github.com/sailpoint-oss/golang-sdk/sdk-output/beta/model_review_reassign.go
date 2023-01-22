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

// ReviewReassign struct for ReviewReassign
type ReviewReassign struct {
	Reassign []ReassignReference `json:"reassign"`
	// The ID of the identity to which the certification is reassigned
	ReassignTo string `json:"reassignTo"`
	// The reason comment for why the reassign was made
	Reason string `json:"reason"`
	AdditionalProperties map[string]interface{}
}

type _ReviewReassign ReviewReassign

// NewReviewReassign instantiates a new ReviewReassign object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewReassign(reassign []ReassignReference, reassignTo string, reason string) *ReviewReassign {
	this := ReviewReassign{}
	this.Reassign = reassign
	this.ReassignTo = reassignTo
	this.Reason = reason
	return &this
}

// NewReviewReassignWithDefaults instantiates a new ReviewReassign object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewReassignWithDefaults() *ReviewReassign {
	this := ReviewReassign{}
	return &this
}

// GetReassign returns the Reassign field value
func (o *ReviewReassign) GetReassign() []ReassignReference {
	if o == nil {
		var ret []ReassignReference
		return ret
	}

	return o.Reassign
}

// GetReassignOk returns a tuple with the Reassign field value
// and a boolean to check if the value has been set.
func (o *ReviewReassign) GetReassignOk() ([]ReassignReference, bool) {
	if o == nil {
		return nil, false
	}
	return o.Reassign, true
}

// SetReassign sets field value
func (o *ReviewReassign) SetReassign(v []ReassignReference) {
	o.Reassign = v
}

// GetReassignTo returns the ReassignTo field value
func (o *ReviewReassign) GetReassignTo() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ReassignTo
}

// GetReassignToOk returns a tuple with the ReassignTo field value
// and a boolean to check if the value has been set.
func (o *ReviewReassign) GetReassignToOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReassignTo, true
}

// SetReassignTo sets field value
func (o *ReviewReassign) SetReassignTo(v string) {
	o.ReassignTo = v
}

// GetReason returns the Reason field value
func (o *ReviewReassign) GetReason() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Reason
}

// GetReasonOk returns a tuple with the Reason field value
// and a boolean to check if the value has been set.
func (o *ReviewReassign) GetReasonOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Reason, true
}

// SetReason sets field value
func (o *ReviewReassign) SetReason(v string) {
	o.Reason = v
}

func (o ReviewReassign) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["reassign"] = o.Reassign
	}
	if true {
		toSerialize["reassignTo"] = o.ReassignTo
	}
	if true {
		toSerialize["reason"] = o.Reason
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewReassign) UnmarshalJSON(bytes []byte) (err error) {
	varReviewReassign := _ReviewReassign{}

	if err = json.Unmarshal(bytes, &varReviewReassign); err == nil {
		*o = ReviewReassign(varReviewReassign)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "reassign")
		delete(additionalProperties, "reassignTo")
		delete(additionalProperties, "reason")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewReassign struct {
	value *ReviewReassign
	isSet bool
}

func (v NullableReviewReassign) Get() *ReviewReassign {
	return v.value
}

func (v *NullableReviewReassign) Set(val *ReviewReassign) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewReassign) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewReassign) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewReassign(val *ReviewReassign) *NullableReviewReassign {
	return &NullableReviewReassign{value: val, isSet: true}
}

func (v NullableReviewReassign) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewReassign) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


