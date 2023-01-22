/*
SailPoint SaaS API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 2.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointv2sdk

import (
	"encoding/json"
	"time"
)

// AccessRequest struct for AccessRequest
type AccessRequest struct {
	Approvals []Approval `json:"approvals,omitempty"`
	CompletionDate *time.Time `json:"completionDate,omitempty"`
	CompletionStatus *string `json:"completionStatus,omitempty"`
	DateCreated *time.Time `json:"dateCreated,omitempty"`
	Errors []string `json:"errors,omitempty"`
	ExecutionStatus *string `json:"executionStatus,omitempty"`
	Id *string `json:"id,omitempty"`
	LastUpdated *time.Time `json:"lastUpdated,omitempty"`
	Name *string `json:"name,omitempty"`
	Priority *string `json:"priority,omitempty"`
	RequestSummaries []RequestSummary `json:"requestSummaries,omitempty"`
	RequesterId *string `json:"requesterId,omitempty"`
	RequesterDisplayName *string `json:"requesterDisplayName,omitempty"`
	Source *string `json:"source,omitempty"`
	SourceName *string `json:"sourceName,omitempty"`
	TargetName *string `json:"targetName,omitempty"`
	TargetDisplayName *string `json:"targetDisplayName,omitempty"`
	TargetId *string `json:"targetId,omitempty"`
	Type *string `json:"type,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	WorkflowStep *string `json:"workflowStep,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _AccessRequest AccessRequest

// NewAccessRequest instantiates a new AccessRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAccessRequest() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// NewAccessRequestWithDefaults instantiates a new AccessRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAccessRequestWithDefaults() *AccessRequest {
	this := AccessRequest{}
	return &this
}

// GetApprovals returns the Approvals field value if set, zero value otherwise.
func (o *AccessRequest) GetApprovals() []Approval {
	if o == nil || isNil(o.Approvals) {
		var ret []Approval
		return ret
	}
	return o.Approvals
}

// GetApprovalsOk returns a tuple with the Approvals field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetApprovalsOk() ([]Approval, bool) {
	if o == nil || isNil(o.Approvals) {
		return nil, false
	}
	return o.Approvals, true
}

// HasApprovals returns a boolean if a field has been set.
func (o *AccessRequest) HasApprovals() bool {
	if o != nil && !isNil(o.Approvals) {
		return true
	}

	return false
}

// SetApprovals gets a reference to the given []Approval and assigns it to the Approvals field.
func (o *AccessRequest) SetApprovals(v []Approval) {
	o.Approvals = v
}

// GetCompletionDate returns the CompletionDate field value if set, zero value otherwise.
func (o *AccessRequest) GetCompletionDate() time.Time {
	if o == nil || isNil(o.CompletionDate) {
		var ret time.Time
		return ret
	}
	return *o.CompletionDate
}

// GetCompletionDateOk returns a tuple with the CompletionDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetCompletionDateOk() (*time.Time, bool) {
	if o == nil || isNil(o.CompletionDate) {
		return nil, false
	}
	return o.CompletionDate, true
}

// HasCompletionDate returns a boolean if a field has been set.
func (o *AccessRequest) HasCompletionDate() bool {
	if o != nil && !isNil(o.CompletionDate) {
		return true
	}

	return false
}

// SetCompletionDate gets a reference to the given time.Time and assigns it to the CompletionDate field.
func (o *AccessRequest) SetCompletionDate(v time.Time) {
	o.CompletionDate = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise.
func (o *AccessRequest) GetCompletionStatus() string {
	if o == nil || isNil(o.CompletionStatus) {
		var ret string
		return ret
	}
	return *o.CompletionStatus
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetCompletionStatusOk() (*string, bool) {
	if o == nil || isNil(o.CompletionStatus) {
		return nil, false
	}
	return o.CompletionStatus, true
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *AccessRequest) HasCompletionStatus() bool {
	if o != nil && !isNil(o.CompletionStatus) {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given string and assigns it to the CompletionStatus field.
func (o *AccessRequest) SetCompletionStatus(v string) {
	o.CompletionStatus = &v
}

// GetDateCreated returns the DateCreated field value if set, zero value otherwise.
func (o *AccessRequest) GetDateCreated() time.Time {
	if o == nil || isNil(o.DateCreated) {
		var ret time.Time
		return ret
	}
	return *o.DateCreated
}

// GetDateCreatedOk returns a tuple with the DateCreated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetDateCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.DateCreated) {
		return nil, false
	}
	return o.DateCreated, true
}

// HasDateCreated returns a boolean if a field has been set.
func (o *AccessRequest) HasDateCreated() bool {
	if o != nil && !isNil(o.DateCreated) {
		return true
	}

	return false
}

// SetDateCreated gets a reference to the given time.Time and assigns it to the DateCreated field.
func (o *AccessRequest) SetDateCreated(v time.Time) {
	o.DateCreated = &v
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *AccessRequest) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *AccessRequest) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *AccessRequest) SetErrors(v []string) {
	o.Errors = v
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *AccessRequest) GetExecutionStatus() string {
	if o == nil || isNil(o.ExecutionStatus) {
		var ret string
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetExecutionStatusOk() (*string, bool) {
	if o == nil || isNil(o.ExecutionStatus) {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *AccessRequest) HasExecutionStatus() bool {
	if o != nil && !isNil(o.ExecutionStatus) {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given string and assigns it to the ExecutionStatus field.
func (o *AccessRequest) SetExecutionStatus(v string) {
	o.ExecutionStatus = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AccessRequest) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AccessRequest) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *AccessRequest) SetId(v string) {
	o.Id = &v
}

// GetLastUpdated returns the LastUpdated field value if set, zero value otherwise.
func (o *AccessRequest) GetLastUpdated() time.Time {
	if o == nil || isNil(o.LastUpdated) {
		var ret time.Time
		return ret
	}
	return *o.LastUpdated
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.LastUpdated) {
		return nil, false
	}
	return o.LastUpdated, true
}

// HasLastUpdated returns a boolean if a field has been set.
func (o *AccessRequest) HasLastUpdated() bool {
	if o != nil && !isNil(o.LastUpdated) {
		return true
	}

	return false
}

// SetLastUpdated gets a reference to the given time.Time and assigns it to the LastUpdated field.
func (o *AccessRequest) SetLastUpdated(v time.Time) {
	o.LastUpdated = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AccessRequest) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AccessRequest) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AccessRequest) SetName(v string) {
	o.Name = &v
}

// GetPriority returns the Priority field value if set, zero value otherwise.
func (o *AccessRequest) GetPriority() string {
	if o == nil || isNil(o.Priority) {
		var ret string
		return ret
	}
	return *o.Priority
}

// GetPriorityOk returns a tuple with the Priority field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetPriorityOk() (*string, bool) {
	if o == nil || isNil(o.Priority) {
		return nil, false
	}
	return o.Priority, true
}

// HasPriority returns a boolean if a field has been set.
func (o *AccessRequest) HasPriority() bool {
	if o != nil && !isNil(o.Priority) {
		return true
	}

	return false
}

// SetPriority gets a reference to the given string and assigns it to the Priority field.
func (o *AccessRequest) SetPriority(v string) {
	o.Priority = &v
}

// GetRequestSummaries returns the RequestSummaries field value if set, zero value otherwise.
func (o *AccessRequest) GetRequestSummaries() []RequestSummary {
	if o == nil || isNil(o.RequestSummaries) {
		var ret []RequestSummary
		return ret
	}
	return o.RequestSummaries
}

// GetRequestSummariesOk returns a tuple with the RequestSummaries field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequestSummariesOk() ([]RequestSummary, bool) {
	if o == nil || isNil(o.RequestSummaries) {
		return nil, false
	}
	return o.RequestSummaries, true
}

// HasRequestSummaries returns a boolean if a field has been set.
func (o *AccessRequest) HasRequestSummaries() bool {
	if o != nil && !isNil(o.RequestSummaries) {
		return true
	}

	return false
}

// SetRequestSummaries gets a reference to the given []RequestSummary and assigns it to the RequestSummaries field.
func (o *AccessRequest) SetRequestSummaries(v []RequestSummary) {
	o.RequestSummaries = v
}

// GetRequesterId returns the RequesterId field value if set, zero value otherwise.
func (o *AccessRequest) GetRequesterId() string {
	if o == nil || isNil(o.RequesterId) {
		var ret string
		return ret
	}
	return *o.RequesterId
}

// GetRequesterIdOk returns a tuple with the RequesterId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequesterIdOk() (*string, bool) {
	if o == nil || isNil(o.RequesterId) {
		return nil, false
	}
	return o.RequesterId, true
}

// HasRequesterId returns a boolean if a field has been set.
func (o *AccessRequest) HasRequesterId() bool {
	if o != nil && !isNil(o.RequesterId) {
		return true
	}

	return false
}

// SetRequesterId gets a reference to the given string and assigns it to the RequesterId field.
func (o *AccessRequest) SetRequesterId(v string) {
	o.RequesterId = &v
}

// GetRequesterDisplayName returns the RequesterDisplayName field value if set, zero value otherwise.
func (o *AccessRequest) GetRequesterDisplayName() string {
	if o == nil || isNil(o.RequesterDisplayName) {
		var ret string
		return ret
	}
	return *o.RequesterDisplayName
}

// GetRequesterDisplayNameOk returns a tuple with the RequesterDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetRequesterDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.RequesterDisplayName) {
		return nil, false
	}
	return o.RequesterDisplayName, true
}

// HasRequesterDisplayName returns a boolean if a field has been set.
func (o *AccessRequest) HasRequesterDisplayName() bool {
	if o != nil && !isNil(o.RequesterDisplayName) {
		return true
	}

	return false
}

// SetRequesterDisplayName gets a reference to the given string and assigns it to the RequesterDisplayName field.
func (o *AccessRequest) SetRequesterDisplayName(v string) {
	o.RequesterDisplayName = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *AccessRequest) GetSource() string {
	if o == nil || isNil(o.Source) {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetSourceOk() (*string, bool) {
	if o == nil || isNil(o.Source) {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *AccessRequest) HasSource() bool {
	if o != nil && !isNil(o.Source) {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *AccessRequest) SetSource(v string) {
	o.Source = &v
}

// GetSourceName returns the SourceName field value if set, zero value otherwise.
func (o *AccessRequest) GetSourceName() string {
	if o == nil || isNil(o.SourceName) {
		var ret string
		return ret
	}
	return *o.SourceName
}

// GetSourceNameOk returns a tuple with the SourceName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetSourceNameOk() (*string, bool) {
	if o == nil || isNil(o.SourceName) {
		return nil, false
	}
	return o.SourceName, true
}

// HasSourceName returns a boolean if a field has been set.
func (o *AccessRequest) HasSourceName() bool {
	if o != nil && !isNil(o.SourceName) {
		return true
	}

	return false
}

// SetSourceName gets a reference to the given string and assigns it to the SourceName field.
func (o *AccessRequest) SetSourceName(v string) {
	o.SourceName = &v
}

// GetTargetName returns the TargetName field value if set, zero value otherwise.
func (o *AccessRequest) GetTargetName() string {
	if o == nil || isNil(o.TargetName) {
		var ret string
		return ret
	}
	return *o.TargetName
}

// GetTargetNameOk returns a tuple with the TargetName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetTargetNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetName) {
		return nil, false
	}
	return o.TargetName, true
}

// HasTargetName returns a boolean if a field has been set.
func (o *AccessRequest) HasTargetName() bool {
	if o != nil && !isNil(o.TargetName) {
		return true
	}

	return false
}

// SetTargetName gets a reference to the given string and assigns it to the TargetName field.
func (o *AccessRequest) SetTargetName(v string) {
	o.TargetName = &v
}

// GetTargetDisplayName returns the TargetDisplayName field value if set, zero value otherwise.
func (o *AccessRequest) GetTargetDisplayName() string {
	if o == nil || isNil(o.TargetDisplayName) {
		var ret string
		return ret
	}
	return *o.TargetDisplayName
}

// GetTargetDisplayNameOk returns a tuple with the TargetDisplayName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetTargetDisplayNameOk() (*string, bool) {
	if o == nil || isNil(o.TargetDisplayName) {
		return nil, false
	}
	return o.TargetDisplayName, true
}

// HasTargetDisplayName returns a boolean if a field has been set.
func (o *AccessRequest) HasTargetDisplayName() bool {
	if o != nil && !isNil(o.TargetDisplayName) {
		return true
	}

	return false
}

// SetTargetDisplayName gets a reference to the given string and assigns it to the TargetDisplayName field.
func (o *AccessRequest) SetTargetDisplayName(v string) {
	o.TargetDisplayName = &v
}

// GetTargetId returns the TargetId field value if set, zero value otherwise.
func (o *AccessRequest) GetTargetId() string {
	if o == nil || isNil(o.TargetId) {
		var ret string
		return ret
	}
	return *o.TargetId
}

// GetTargetIdOk returns a tuple with the TargetId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetTargetIdOk() (*string, bool) {
	if o == nil || isNil(o.TargetId) {
		return nil, false
	}
	return o.TargetId, true
}

// HasTargetId returns a boolean if a field has been set.
func (o *AccessRequest) HasTargetId() bool {
	if o != nil && !isNil(o.TargetId) {
		return true
	}

	return false
}

// SetTargetId gets a reference to the given string and assigns it to the TargetId field.
func (o *AccessRequest) SetTargetId(v string) {
	o.TargetId = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *AccessRequest) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *AccessRequest) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *AccessRequest) SetType(v string) {
	o.Type = &v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *AccessRequest) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *AccessRequest) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *AccessRequest) SetWarnings(v []string) {
	o.Warnings = v
}

// GetWorkflowStep returns the WorkflowStep field value if set, zero value otherwise.
func (o *AccessRequest) GetWorkflowStep() string {
	if o == nil || isNil(o.WorkflowStep) {
		var ret string
		return ret
	}
	return *o.WorkflowStep
}

// GetWorkflowStepOk returns a tuple with the WorkflowStep field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AccessRequest) GetWorkflowStepOk() (*string, bool) {
	if o == nil || isNil(o.WorkflowStep) {
		return nil, false
	}
	return o.WorkflowStep, true
}

// HasWorkflowStep returns a boolean if a field has been set.
func (o *AccessRequest) HasWorkflowStep() bool {
	if o != nil && !isNil(o.WorkflowStep) {
		return true
	}

	return false
}

// SetWorkflowStep gets a reference to the given string and assigns it to the WorkflowStep field.
func (o *AccessRequest) SetWorkflowStep(v string) {
	o.WorkflowStep = &v
}

func (o AccessRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Approvals) {
		toSerialize["approvals"] = o.Approvals
	}
	if !isNil(o.CompletionDate) {
		toSerialize["completionDate"] = o.CompletionDate
	}
	if !isNil(o.CompletionStatus) {
		toSerialize["completionStatus"] = o.CompletionStatus
	}
	if !isNil(o.DateCreated) {
		toSerialize["dateCreated"] = o.DateCreated
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.ExecutionStatus) {
		toSerialize["executionStatus"] = o.ExecutionStatus
	}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.LastUpdated) {
		toSerialize["lastUpdated"] = o.LastUpdated
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Priority) {
		toSerialize["priority"] = o.Priority
	}
	if !isNil(o.RequestSummaries) {
		toSerialize["requestSummaries"] = o.RequestSummaries
	}
	if !isNil(o.RequesterId) {
		toSerialize["requesterId"] = o.RequesterId
	}
	if !isNil(o.RequesterDisplayName) {
		toSerialize["requesterDisplayName"] = o.RequesterDisplayName
	}
	if !isNil(o.Source) {
		toSerialize["source"] = o.Source
	}
	if !isNil(o.SourceName) {
		toSerialize["sourceName"] = o.SourceName
	}
	if !isNil(o.TargetName) {
		toSerialize["targetName"] = o.TargetName
	}
	if !isNil(o.TargetDisplayName) {
		toSerialize["targetDisplayName"] = o.TargetDisplayName
	}
	if !isNil(o.TargetId) {
		toSerialize["targetId"] = o.TargetId
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.WorkflowStep) {
		toSerialize["workflowStep"] = o.WorkflowStep
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *AccessRequest) UnmarshalJSON(bytes []byte) (err error) {
	varAccessRequest := _AccessRequest{}

	if err = json.Unmarshal(bytes, &varAccessRequest); err == nil {
		*o = AccessRequest(varAccessRequest)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "approvals")
		delete(additionalProperties, "completionDate")
		delete(additionalProperties, "completionStatus")
		delete(additionalProperties, "dateCreated")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "executionStatus")
		delete(additionalProperties, "id")
		delete(additionalProperties, "lastUpdated")
		delete(additionalProperties, "name")
		delete(additionalProperties, "priority")
		delete(additionalProperties, "requestSummaries")
		delete(additionalProperties, "requesterId")
		delete(additionalProperties, "requesterDisplayName")
		delete(additionalProperties, "source")
		delete(additionalProperties, "sourceName")
		delete(additionalProperties, "targetName")
		delete(additionalProperties, "targetDisplayName")
		delete(additionalProperties, "targetId")
		delete(additionalProperties, "type")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "workflowStep")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableAccessRequest struct {
	value *AccessRequest
	isSet bool
}

func (v NullableAccessRequest) Get() *AccessRequest {
	return v.value
}

func (v *NullableAccessRequest) Set(val *AccessRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableAccessRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableAccessRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAccessRequest(val *AccessRequest) *NullableAccessRequest {
	return &NullableAccessRequest{value: val, isSet: true}
}

func (v NullableAccessRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAccessRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


