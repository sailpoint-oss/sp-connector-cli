/*
IdentityNow Beta API

Use these APIs to interact with the IdentityNow platform to achieve repeatable, automated processes with greater scalability. These APIs are in beta and are subject to change. We encourage you to join the SailPoint Developer Community forum at https://developer.sailpoint.com/discuss to connect with other developers using our APIs.

API version: 3.1.0-beta
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sailpointbetasdk

import (
	"encoding/json"
	"time"
)

// CancelableAccountActivity struct for CancelableAccountActivity
type CancelableAccountActivity struct {
	// ID of the account activity itself
	Id *string `json:"id,omitempty"`
	Name *string `json:"name,omitempty"`
	Created *time.Time `json:"created,omitempty"`
	Modified *time.Time `json:"modified,omitempty"`
	Completed *time.Time `json:"completed,omitempty"`
	CompletionStatus NullableCompletionStatus `json:"completionStatus,omitempty"`
	Type *string `json:"type,omitempty"`
	RequesterIdentitySummary NullableIdentitySummary `json:"requesterIdentitySummary,omitempty"`
	TargetIdentitySummary NullableIdentitySummary `json:"targetIdentitySummary,omitempty"`
	Errors []string `json:"errors,omitempty"`
	Warnings []string `json:"warnings,omitempty"`
	Items []AccountActivityItem `json:"items,omitempty"`
	ExecutionStatus *ExecutionStatus `json:"executionStatus,omitempty"`
	// Arbitrary key-value pairs, if any were included in the corresponding access request
	ClientMetadata *map[string]string `json:"clientMetadata,omitempty"`
	// Whether the account activity can be canceled before completion
	Cancelable *bool `json:"cancelable,omitempty"`
	CancelComment NullableComment `json:"cancelComment,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _CancelableAccountActivity CancelableAccountActivity

// NewCancelableAccountActivity instantiates a new CancelableAccountActivity object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCancelableAccountActivity() *CancelableAccountActivity {
	this := CancelableAccountActivity{}
	return &this
}

// NewCancelableAccountActivityWithDefaults instantiates a new CancelableAccountActivity object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCancelableAccountActivityWithDefaults() *CancelableAccountActivity {
	this := CancelableAccountActivity{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *CancelableAccountActivity) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *CancelableAccountActivity) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *CancelableAccountActivity) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetModified() time.Time {
	if o == nil || isNil(o.Modified) {
		var ret time.Time
		return ret
	}
	return *o.Modified
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetModifiedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Modified) {
		return nil, false
	}
	return o.Modified, true
}

// HasModified returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasModified() bool {
	if o != nil && !isNil(o.Modified) {
		return true
	}

	return false
}

// SetModified gets a reference to the given time.Time and assigns it to the Modified field.
func (o *CancelableAccountActivity) SetModified(v time.Time) {
	o.Modified = &v
}

// GetCompleted returns the Completed field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetCompleted() time.Time {
	if o == nil || isNil(o.Completed) {
		var ret time.Time
		return ret
	}
	return *o.Completed
}

// GetCompletedOk returns a tuple with the Completed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetCompletedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Completed) {
		return nil, false
	}
	return o.Completed, true
}

// HasCompleted returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasCompleted() bool {
	if o != nil && !isNil(o.Completed) {
		return true
	}

	return false
}

// SetCompleted gets a reference to the given time.Time and assigns it to the Completed field.
func (o *CancelableAccountActivity) SetCompleted(v time.Time) {
	o.Completed = &v
}

// GetCompletionStatus returns the CompletionStatus field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CancelableAccountActivity) GetCompletionStatus() CompletionStatus {
	if o == nil || isNil(o.CompletionStatus.Get()) {
		var ret CompletionStatus
		return ret
	}
	return *o.CompletionStatus.Get()
}

// GetCompletionStatusOk returns a tuple with the CompletionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CancelableAccountActivity) GetCompletionStatusOk() (*CompletionStatus, bool) {
	if o == nil {
		return nil, false
	}
	return o.CompletionStatus.Get(), o.CompletionStatus.IsSet()
}

// HasCompletionStatus returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasCompletionStatus() bool {
	if o != nil && o.CompletionStatus.IsSet() {
		return true
	}

	return false
}

// SetCompletionStatus gets a reference to the given NullableCompletionStatus and assigns it to the CompletionStatus field.
func (o *CancelableAccountActivity) SetCompletionStatus(v CompletionStatus) {
	o.CompletionStatus.Set(&v)
}
// SetCompletionStatusNil sets the value for CompletionStatus to be an explicit nil
func (o *CancelableAccountActivity) SetCompletionStatusNil() {
	o.CompletionStatus.Set(nil)
}

// UnsetCompletionStatus ensures that no value is present for CompletionStatus, not even an explicit nil
func (o *CancelableAccountActivity) UnsetCompletionStatus() {
	o.CompletionStatus.Unset()
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetType() string {
	if o == nil || isNil(o.Type) {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetTypeOk() (*string, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *CancelableAccountActivity) SetType(v string) {
	o.Type = &v
}

// GetRequesterIdentitySummary returns the RequesterIdentitySummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CancelableAccountActivity) GetRequesterIdentitySummary() IdentitySummary {
	if o == nil || isNil(o.RequesterIdentitySummary.Get()) {
		var ret IdentitySummary
		return ret
	}
	return *o.RequesterIdentitySummary.Get()
}

// GetRequesterIdentitySummaryOk returns a tuple with the RequesterIdentitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CancelableAccountActivity) GetRequesterIdentitySummaryOk() (*IdentitySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.RequesterIdentitySummary.Get(), o.RequesterIdentitySummary.IsSet()
}

// HasRequesterIdentitySummary returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasRequesterIdentitySummary() bool {
	if o != nil && o.RequesterIdentitySummary.IsSet() {
		return true
	}

	return false
}

// SetRequesterIdentitySummary gets a reference to the given NullableIdentitySummary and assigns it to the RequesterIdentitySummary field.
func (o *CancelableAccountActivity) SetRequesterIdentitySummary(v IdentitySummary) {
	o.RequesterIdentitySummary.Set(&v)
}
// SetRequesterIdentitySummaryNil sets the value for RequesterIdentitySummary to be an explicit nil
func (o *CancelableAccountActivity) SetRequesterIdentitySummaryNil() {
	o.RequesterIdentitySummary.Set(nil)
}

// UnsetRequesterIdentitySummary ensures that no value is present for RequesterIdentitySummary, not even an explicit nil
func (o *CancelableAccountActivity) UnsetRequesterIdentitySummary() {
	o.RequesterIdentitySummary.Unset()
}

// GetTargetIdentitySummary returns the TargetIdentitySummary field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CancelableAccountActivity) GetTargetIdentitySummary() IdentitySummary {
	if o == nil || isNil(o.TargetIdentitySummary.Get()) {
		var ret IdentitySummary
		return ret
	}
	return *o.TargetIdentitySummary.Get()
}

// GetTargetIdentitySummaryOk returns a tuple with the TargetIdentitySummary field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CancelableAccountActivity) GetTargetIdentitySummaryOk() (*IdentitySummary, bool) {
	if o == nil {
		return nil, false
	}
	return o.TargetIdentitySummary.Get(), o.TargetIdentitySummary.IsSet()
}

// HasTargetIdentitySummary returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasTargetIdentitySummary() bool {
	if o != nil && o.TargetIdentitySummary.IsSet() {
		return true
	}

	return false
}

// SetTargetIdentitySummary gets a reference to the given NullableIdentitySummary and assigns it to the TargetIdentitySummary field.
func (o *CancelableAccountActivity) SetTargetIdentitySummary(v IdentitySummary) {
	o.TargetIdentitySummary.Set(&v)
}
// SetTargetIdentitySummaryNil sets the value for TargetIdentitySummary to be an explicit nil
func (o *CancelableAccountActivity) SetTargetIdentitySummaryNil() {
	o.TargetIdentitySummary.Set(nil)
}

// UnsetTargetIdentitySummary ensures that no value is present for TargetIdentitySummary, not even an explicit nil
func (o *CancelableAccountActivity) UnsetTargetIdentitySummary() {
	o.TargetIdentitySummary.Unset()
}

// GetErrors returns the Errors field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetErrors() []string {
	if o == nil || isNil(o.Errors) {
		var ret []string
		return ret
	}
	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetErrorsOk() ([]string, bool) {
	if o == nil || isNil(o.Errors) {
		return nil, false
	}
	return o.Errors, true
}

// HasErrors returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasErrors() bool {
	if o != nil && !isNil(o.Errors) {
		return true
	}

	return false
}

// SetErrors gets a reference to the given []string and assigns it to the Errors field.
func (o *CancelableAccountActivity) SetErrors(v []string) {
	o.Errors = v
}

// GetWarnings returns the Warnings field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetWarnings() []string {
	if o == nil || isNil(o.Warnings) {
		var ret []string
		return ret
	}
	return o.Warnings
}

// GetWarningsOk returns a tuple with the Warnings field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetWarningsOk() ([]string, bool) {
	if o == nil || isNil(o.Warnings) {
		return nil, false
	}
	return o.Warnings, true
}

// HasWarnings returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasWarnings() bool {
	if o != nil && !isNil(o.Warnings) {
		return true
	}

	return false
}

// SetWarnings gets a reference to the given []string and assigns it to the Warnings field.
func (o *CancelableAccountActivity) SetWarnings(v []string) {
	o.Warnings = v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetItems() []AccountActivityItem {
	if o == nil || isNil(o.Items) {
		var ret []AccountActivityItem
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetItemsOk() ([]AccountActivityItem, bool) {
	if o == nil || isNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasItems() bool {
	if o != nil && !isNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []AccountActivityItem and assigns it to the Items field.
func (o *CancelableAccountActivity) SetItems(v []AccountActivityItem) {
	o.Items = v
}

// GetExecutionStatus returns the ExecutionStatus field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetExecutionStatus() ExecutionStatus {
	if o == nil || isNil(o.ExecutionStatus) {
		var ret ExecutionStatus
		return ret
	}
	return *o.ExecutionStatus
}

// GetExecutionStatusOk returns a tuple with the ExecutionStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetExecutionStatusOk() (*ExecutionStatus, bool) {
	if o == nil || isNil(o.ExecutionStatus) {
		return nil, false
	}
	return o.ExecutionStatus, true
}

// HasExecutionStatus returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasExecutionStatus() bool {
	if o != nil && !isNil(o.ExecutionStatus) {
		return true
	}

	return false
}

// SetExecutionStatus gets a reference to the given ExecutionStatus and assigns it to the ExecutionStatus field.
func (o *CancelableAccountActivity) SetExecutionStatus(v ExecutionStatus) {
	o.ExecutionStatus = &v
}

// GetClientMetadata returns the ClientMetadata field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetClientMetadata() map[string]string {
	if o == nil || isNil(o.ClientMetadata) {
		var ret map[string]string
		return ret
	}
	return *o.ClientMetadata
}

// GetClientMetadataOk returns a tuple with the ClientMetadata field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetClientMetadataOk() (*map[string]string, bool) {
	if o == nil || isNil(o.ClientMetadata) {
		return nil, false
	}
	return o.ClientMetadata, true
}

// HasClientMetadata returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasClientMetadata() bool {
	if o != nil && !isNil(o.ClientMetadata) {
		return true
	}

	return false
}

// SetClientMetadata gets a reference to the given map[string]string and assigns it to the ClientMetadata field.
func (o *CancelableAccountActivity) SetClientMetadata(v map[string]string) {
	o.ClientMetadata = &v
}

// GetCancelable returns the Cancelable field value if set, zero value otherwise.
func (o *CancelableAccountActivity) GetCancelable() bool {
	if o == nil || isNil(o.Cancelable) {
		var ret bool
		return ret
	}
	return *o.Cancelable
}

// GetCancelableOk returns a tuple with the Cancelable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CancelableAccountActivity) GetCancelableOk() (*bool, bool) {
	if o == nil || isNil(o.Cancelable) {
		return nil, false
	}
	return o.Cancelable, true
}

// HasCancelable returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasCancelable() bool {
	if o != nil && !isNil(o.Cancelable) {
		return true
	}

	return false
}

// SetCancelable gets a reference to the given bool and assigns it to the Cancelable field.
func (o *CancelableAccountActivity) SetCancelable(v bool) {
	o.Cancelable = &v
}

// GetCancelComment returns the CancelComment field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *CancelableAccountActivity) GetCancelComment() Comment {
	if o == nil || isNil(o.CancelComment.Get()) {
		var ret Comment
		return ret
	}
	return *o.CancelComment.Get()
}

// GetCancelCommentOk returns a tuple with the CancelComment field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *CancelableAccountActivity) GetCancelCommentOk() (*Comment, bool) {
	if o == nil {
		return nil, false
	}
	return o.CancelComment.Get(), o.CancelComment.IsSet()
}

// HasCancelComment returns a boolean if a field has been set.
func (o *CancelableAccountActivity) HasCancelComment() bool {
	if o != nil && o.CancelComment.IsSet() {
		return true
	}

	return false
}

// SetCancelComment gets a reference to the given NullableComment and assigns it to the CancelComment field.
func (o *CancelableAccountActivity) SetCancelComment(v Comment) {
	o.CancelComment.Set(&v)
}
// SetCancelCommentNil sets the value for CancelComment to be an explicit nil
func (o *CancelableAccountActivity) SetCancelCommentNil() {
	o.CancelComment.Set(nil)
}

// UnsetCancelComment ensures that no value is present for CancelComment, not even an explicit nil
func (o *CancelableAccountActivity) UnsetCancelComment() {
	o.CancelComment.Unset()
}

func (o CancelableAccountActivity) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if !isNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !isNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !isNil(o.Created) {
		toSerialize["created"] = o.Created
	}
	if !isNil(o.Modified) {
		toSerialize["modified"] = o.Modified
	}
	if !isNil(o.Completed) {
		toSerialize["completed"] = o.Completed
	}
	if o.CompletionStatus.IsSet() {
		toSerialize["completionStatus"] = o.CompletionStatus.Get()
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if o.RequesterIdentitySummary.IsSet() {
		toSerialize["requesterIdentitySummary"] = o.RequesterIdentitySummary.Get()
	}
	if o.TargetIdentitySummary.IsSet() {
		toSerialize["targetIdentitySummary"] = o.TargetIdentitySummary.Get()
	}
	if !isNil(o.Errors) {
		toSerialize["errors"] = o.Errors
	}
	if !isNil(o.Warnings) {
		toSerialize["warnings"] = o.Warnings
	}
	if !isNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !isNil(o.ExecutionStatus) {
		toSerialize["executionStatus"] = o.ExecutionStatus
	}
	if !isNil(o.ClientMetadata) {
		toSerialize["clientMetadata"] = o.ClientMetadata
	}
	if !isNil(o.Cancelable) {
		toSerialize["cancelable"] = o.Cancelable
	}
	if o.CancelComment.IsSet() {
		toSerialize["cancelComment"] = o.CancelComment.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *CancelableAccountActivity) UnmarshalJSON(bytes []byte) (err error) {
	varCancelableAccountActivity := _CancelableAccountActivity{}

	if err = json.Unmarshal(bytes, &varCancelableAccountActivity); err == nil {
		*o = CancelableAccountActivity(varCancelableAccountActivity)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "completed")
		delete(additionalProperties, "completionStatus")
		delete(additionalProperties, "type")
		delete(additionalProperties, "requesterIdentitySummary")
		delete(additionalProperties, "targetIdentitySummary")
		delete(additionalProperties, "errors")
		delete(additionalProperties, "warnings")
		delete(additionalProperties, "items")
		delete(additionalProperties, "executionStatus")
		delete(additionalProperties, "clientMetadata")
		delete(additionalProperties, "cancelable")
		delete(additionalProperties, "cancelComment")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableCancelableAccountActivity struct {
	value *CancelableAccountActivity
	isSet bool
}

func (v NullableCancelableAccountActivity) Get() *CancelableAccountActivity {
	return v.value
}

func (v *NullableCancelableAccountActivity) Set(val *CancelableAccountActivity) {
	v.value = val
	v.isSet = true
}

func (v NullableCancelableAccountActivity) IsSet() bool {
	return v.isSet
}

func (v *NullableCancelableAccountActivity) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCancelableAccountActivity(val *CancelableAccountActivity) *NullableCancelableAccountActivity {
	return &NullableCancelableAccountActivity{value: val, isSet: true}
}

func (v NullableCancelableAccountActivity) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCancelableAccountActivity) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


