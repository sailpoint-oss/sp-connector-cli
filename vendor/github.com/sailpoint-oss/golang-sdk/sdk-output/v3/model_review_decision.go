/*
IdentityNow V3 API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointsdk

import (
	"encoding/json"
	"time"
)

// ReviewDecision struct for ReviewDecision
type ReviewDecision struct {
	// The id of the review decision
	Id string `json:"id"`
	Decision CertificationDecision `json:"decision"`
	// The date at which a user's access should be taken away. Should only be set for `REVOKE` decisions.
	ProposedEndDate *time.Time `json:"proposedEndDate,omitempty"`
	// Indicates whether decision should be marked as part of a larger bulk decision
	Bulk bool `json:"bulk"`
	Recommendation *ReviewRecommendation `json:"recommendation,omitempty"`
	// Comments recorded when the decision was made
	Comments *string `json:"comments,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _ReviewDecision ReviewDecision

// NewReviewDecision instantiates a new ReviewDecision object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewReviewDecision(id string, decision CertificationDecision, bulk bool) *ReviewDecision {
	this := ReviewDecision{}
	this.Id = id
	this.Decision = decision
	this.Bulk = bulk
	return &this
}

// NewReviewDecisionWithDefaults instantiates a new ReviewDecision object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewReviewDecisionWithDefaults() *ReviewDecision {
	this := ReviewDecision{}
	return &this
}

// GetId returns the Id field value
func (o *ReviewDecision) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *ReviewDecision) SetId(v string) {
	o.Id = v
}

// GetDecision returns the Decision field value
func (o *ReviewDecision) GetDecision() CertificationDecision {
	if o == nil {
		var ret CertificationDecision
		return ret
	}

	return o.Decision
}

// GetDecisionOk returns a tuple with the Decision field value
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetDecisionOk() (*CertificationDecision, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Decision, true
}

// SetDecision sets field value
func (o *ReviewDecision) SetDecision(v CertificationDecision) {
	o.Decision = v
}

// GetProposedEndDate returns the ProposedEndDate field value if set, zero value otherwise.
func (o *ReviewDecision) GetProposedEndDate() time.Time {
	if o == nil || isNil(o.ProposedEndDate) {
		var ret time.Time
		return ret
	}
	return *o.ProposedEndDate
}

// GetProposedEndDateOk returns a tuple with the ProposedEndDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetProposedEndDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.ProposedEndDate) {
		return nil, false
	}
	return o.ProposedEndDate, true
}

// HasProposedEndDate returns a boolean if a field has been set.
func (o *ReviewDecision) HasProposedEndDate() bool {
	if o != nil && !isNil(o.ProposedEndDate) {
		return true
	}

	return false
}

// SetProposedEndDate gets a reference to the given time.Time and assigns it to the ProposedEndDate field.
func (o *ReviewDecision) SetProposedEndDate(v time.Time) {
	o.ProposedEndDate = &v
}

// GetBulk returns the Bulk field value
func (o *ReviewDecision) GetBulk() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Bulk
}

// GetBulkOk returns a tuple with the Bulk field value
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetBulkOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Bulk, true
}

// SetBulk sets field value
func (o *ReviewDecision) SetBulk(v bool) {
	o.Bulk = v
}

// GetRecommendation returns the Recommendation field value if set, zero value otherwise.
func (o *ReviewDecision) GetRecommendation() ReviewRecommendation {
	if o == nil || isNil(o.Recommendation) {
		var ret ReviewRecommendation
		return ret
	}
	return *o.Recommendation
}

// GetRecommendationOk returns a tuple with the Recommendation field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetRecommendationOk() (*ReviewRecommendation, bool) {
	if o == nil || isNil(o.Recommendation) {
		return nil, false
	}
	return o.Recommendation, true
}

// HasRecommendation returns a boolean if a field has been set.
func (o *ReviewDecision) HasRecommendation() bool {
	if o != nil && !isNil(o.Recommendation) {
		return true
	}

	return false
}

// SetRecommendation gets a reference to the given ReviewRecommendation and assigns it to the Recommendation field.
func (o *ReviewDecision) SetRecommendation(v ReviewRecommendation) {
	o.Recommendation = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *ReviewDecision) GetComments() string {
	if o == nil || isNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ReviewDecision) GetCommentsOk() (*string, bool) {
	if o == nil || isNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *ReviewDecision) HasComments() bool {
	if o != nil && !isNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *ReviewDecision) SetComments(v string) {
	o.Comments = &v
}

func (o ReviewDecision) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["decision"] = o.Decision
	}
	if !isNil(o.ProposedEndDate) {
		toSerialize["proposedEndDate"] = o.ProposedEndDate
	}
	if true {
		toSerialize["bulk"] = o.Bulk
	}
	if !isNil(o.Recommendation) {
		toSerialize["recommendation"] = o.Recommendation
	}
	if !isNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *ReviewDecision) UnmarshalJSON(bytes []byte) (err error) {
	varReviewDecision := _ReviewDecision{}

	if err = json.Unmarshal(bytes, &varReviewDecision); err == nil {
		*o = ReviewDecision(varReviewDecision)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "decision")
		delete(additionalProperties, "proposedEndDate")
		delete(additionalProperties, "bulk")
		delete(additionalProperties, "recommendation")
		delete(additionalProperties, "comments")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableReviewDecision struct {
	value *ReviewDecision
	isSet bool
}

func (v NullableReviewDecision) Get() *ReviewDecision {
	return v.value
}

func (v *NullableReviewDecision) Set(val *ReviewDecision) {
	v.value = val
	v.isSet = true
}

func (v NullableReviewDecision) IsSet() bool {
	return v.isSet
}

func (v *NullableReviewDecision) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableReviewDecision(val *ReviewDecision) *NullableReviewDecision {
	return &NullableReviewDecision{value: val, isSet: true}
}

func (v NullableReviewDecision) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableReviewDecision) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


