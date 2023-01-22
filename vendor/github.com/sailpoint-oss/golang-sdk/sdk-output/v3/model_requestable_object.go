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

// RequestableObject struct for RequestableObject
type RequestableObject struct {
	// Id of the requestable object itself
	Id *string `json:"id,omitempty"`
	// Human-readable display name of the requestable object
	Name *string `json:"name,omitempty"`
	// The time when the requestable object was created
	Created *time.Time `json:"created,omitempty"`
	// The time when the requestable object was last modified
	Modified NullableTime `json:"modified,omitempty"`
	// Description of the requestable object.
	Description *string `json:"description,omitempty"`
	Type *RequestableObjectType `json:"type,omitempty"`
	RequestStatus *RequestableObjectRequestStatus `json:"requestStatus,omitempty"`
	// If *requestStatus* is *PENDING*, indicates the id of the associated account activity.
	IdentityRequestId NullableString `json:"identityRequestId,omitempty"`
	OwnerRef NullableIdentityReferenceWithNameAndEmail `json:"ownerRef,omitempty"`
	// Whether the requester must provide comments when requesting the object.
	RequestCommentsRequired *bool `json:"requestCommentsRequired,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RequestableObject RequestableObject

// NewRequestableObject instantiates a new RequestableObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRequestableObject() *RequestableObject {
	this := RequestableObject{}
	return &this
}

// NewRequestableObjectWithDefaults instantiates a new RequestableObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRequestableObjectWithDefaults() *RequestableObject {
	this := RequestableObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *RequestableObject) GetId() string {
	if o == nil || isNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetIdOk() (*string, bool) {
	if o == nil || isNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *RequestableObject) HasId() bool {
	if o != nil && !isNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *RequestableObject) SetId(v string) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *RequestableObject) GetName() string {
	if o == nil || isNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetNameOk() (*string, bool) {
	if o == nil || isNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *RequestableObject) HasName() bool {
	if o != nil && !isNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *RequestableObject) SetName(v string) {
	o.Name = &v
}

// GetCreated returns the Created field value if set, zero value otherwise.
func (o *RequestableObject) GetCreated() time.Time {
	if o == nil || isNil(o.Created) {
		var ret time.Time
		return ret
	}
	return *o.Created
}

// GetCreatedOk returns a tuple with the Created field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetCreatedOk() (*time.Time, bool) {
	if o == nil || isNil(o.Created) {
		return nil, false
	}
	return o.Created, true
}

// HasCreated returns a boolean if a field has been set.
func (o *RequestableObject) HasCreated() bool {
	if o != nil && !isNil(o.Created) {
		return true
	}

	return false
}

// SetCreated gets a reference to the given time.Time and assigns it to the Created field.
func (o *RequestableObject) SetCreated(v time.Time) {
	o.Created = &v
}

// GetModified returns the Modified field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestableObject) GetModified() time.Time {
	if o == nil || isNil(o.Modified.Get()) {
		var ret time.Time
		return ret
	}
	return *o.Modified.Get()
}

// GetModifiedOk returns a tuple with the Modified field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestableObject) GetModifiedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Modified.Get(), o.Modified.IsSet()
}

// HasModified returns a boolean if a field has been set.
func (o *RequestableObject) HasModified() bool {
	if o != nil && o.Modified.IsSet() {
		return true
	}

	return false
}

// SetModified gets a reference to the given NullableTime and assigns it to the Modified field.
func (o *RequestableObject) SetModified(v time.Time) {
	o.Modified.Set(&v)
}
// SetModifiedNil sets the value for Modified to be an explicit nil
func (o *RequestableObject) SetModifiedNil() {
	o.Modified.Set(nil)
}

// UnsetModified ensures that no value is present for Modified, not even an explicit nil
func (o *RequestableObject) UnsetModified() {
	o.Modified.Unset()
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *RequestableObject) GetDescription() string {
	if o == nil || isNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetDescriptionOk() (*string, bool) {
	if o == nil || isNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *RequestableObject) HasDescription() bool {
	if o != nil && !isNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *RequestableObject) SetDescription(v string) {
	o.Description = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *RequestableObject) GetType() RequestableObjectType {
	if o == nil || isNil(o.Type) {
		var ret RequestableObjectType
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetTypeOk() (*RequestableObjectType, bool) {
	if o == nil || isNil(o.Type) {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *RequestableObject) HasType() bool {
	if o != nil && !isNil(o.Type) {
		return true
	}

	return false
}

// SetType gets a reference to the given RequestableObjectType and assigns it to the Type field.
func (o *RequestableObject) SetType(v RequestableObjectType) {
	o.Type = &v
}

// GetRequestStatus returns the RequestStatus field value if set, zero value otherwise.
func (o *RequestableObject) GetRequestStatus() RequestableObjectRequestStatus {
	if o == nil || isNil(o.RequestStatus) {
		var ret RequestableObjectRequestStatus
		return ret
	}
	return *o.RequestStatus
}

// GetRequestStatusOk returns a tuple with the RequestStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetRequestStatusOk() (*RequestableObjectRequestStatus, bool) {
	if o == nil || isNil(o.RequestStatus) {
		return nil, false
	}
	return o.RequestStatus, true
}

// HasRequestStatus returns a boolean if a field has been set.
func (o *RequestableObject) HasRequestStatus() bool {
	if o != nil && !isNil(o.RequestStatus) {
		return true
	}

	return false
}

// SetRequestStatus gets a reference to the given RequestableObjectRequestStatus and assigns it to the RequestStatus field.
func (o *RequestableObject) SetRequestStatus(v RequestableObjectRequestStatus) {
	o.RequestStatus = &v
}

// GetIdentityRequestId returns the IdentityRequestId field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestableObject) GetIdentityRequestId() string {
	if o == nil || isNil(o.IdentityRequestId.Get()) {
		var ret string
		return ret
	}
	return *o.IdentityRequestId.Get()
}

// GetIdentityRequestIdOk returns a tuple with the IdentityRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestableObject) GetIdentityRequestIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.IdentityRequestId.Get(), o.IdentityRequestId.IsSet()
}

// HasIdentityRequestId returns a boolean if a field has been set.
func (o *RequestableObject) HasIdentityRequestId() bool {
	if o != nil && o.IdentityRequestId.IsSet() {
		return true
	}

	return false
}

// SetIdentityRequestId gets a reference to the given NullableString and assigns it to the IdentityRequestId field.
func (o *RequestableObject) SetIdentityRequestId(v string) {
	o.IdentityRequestId.Set(&v)
}
// SetIdentityRequestIdNil sets the value for IdentityRequestId to be an explicit nil
func (o *RequestableObject) SetIdentityRequestIdNil() {
	o.IdentityRequestId.Set(nil)
}

// UnsetIdentityRequestId ensures that no value is present for IdentityRequestId, not even an explicit nil
func (o *RequestableObject) UnsetIdentityRequestId() {
	o.IdentityRequestId.Unset()
}

// GetOwnerRef returns the OwnerRef field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RequestableObject) GetOwnerRef() IdentityReferenceWithNameAndEmail {
	if o == nil || isNil(o.OwnerRef.Get()) {
		var ret IdentityReferenceWithNameAndEmail
		return ret
	}
	return *o.OwnerRef.Get()
}

// GetOwnerRefOk returns a tuple with the OwnerRef field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RequestableObject) GetOwnerRefOk() (*IdentityReferenceWithNameAndEmail, bool) {
	if o == nil {
		return nil, false
	}
	return o.OwnerRef.Get(), o.OwnerRef.IsSet()
}

// HasOwnerRef returns a boolean if a field has been set.
func (o *RequestableObject) HasOwnerRef() bool {
	if o != nil && o.OwnerRef.IsSet() {
		return true
	}

	return false
}

// SetOwnerRef gets a reference to the given NullableIdentityReferenceWithNameAndEmail and assigns it to the OwnerRef field.
func (o *RequestableObject) SetOwnerRef(v IdentityReferenceWithNameAndEmail) {
	o.OwnerRef.Set(&v)
}
// SetOwnerRefNil sets the value for OwnerRef to be an explicit nil
func (o *RequestableObject) SetOwnerRefNil() {
	o.OwnerRef.Set(nil)
}

// UnsetOwnerRef ensures that no value is present for OwnerRef, not even an explicit nil
func (o *RequestableObject) UnsetOwnerRef() {
	o.OwnerRef.Unset()
}

// GetRequestCommentsRequired returns the RequestCommentsRequired field value if set, zero value otherwise.
func (o *RequestableObject) GetRequestCommentsRequired() bool {
	if o == nil || isNil(o.RequestCommentsRequired) {
		var ret bool
		return ret
	}
	return *o.RequestCommentsRequired
}

// GetRequestCommentsRequiredOk returns a tuple with the RequestCommentsRequired field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RequestableObject) GetRequestCommentsRequiredOk() (*bool, bool) {
	if o == nil || isNil(o.RequestCommentsRequired) {
		return nil, false
	}
	return o.RequestCommentsRequired, true
}

// HasRequestCommentsRequired returns a boolean if a field has been set.
func (o *RequestableObject) HasRequestCommentsRequired() bool {
	if o != nil && !isNil(o.RequestCommentsRequired) {
		return true
	}

	return false
}

// SetRequestCommentsRequired gets a reference to the given bool and assigns it to the RequestCommentsRequired field.
func (o *RequestableObject) SetRequestCommentsRequired(v bool) {
	o.RequestCommentsRequired = &v
}

func (o RequestableObject) MarshalJSON() ([]byte, error) {
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
	if o.Modified.IsSet() {
		toSerialize["modified"] = o.Modified.Get()
	}
	if !isNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !isNil(o.Type) {
		toSerialize["type"] = o.Type
	}
	if !isNil(o.RequestStatus) {
		toSerialize["requestStatus"] = o.RequestStatus
	}
	if o.IdentityRequestId.IsSet() {
		toSerialize["identityRequestId"] = o.IdentityRequestId.Get()
	}
	if o.OwnerRef.IsSet() {
		toSerialize["ownerRef"] = o.OwnerRef.Get()
	}
	if !isNil(o.RequestCommentsRequired) {
		toSerialize["requestCommentsRequired"] = o.RequestCommentsRequired
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return json.Marshal(toSerialize)
}

func (o *RequestableObject) UnmarshalJSON(bytes []byte) (err error) {
	varRequestableObject := _RequestableObject{}

	if err = json.Unmarshal(bytes, &varRequestableObject); err == nil {
		*o = RequestableObject(varRequestableObject)
	}

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(bytes, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "name")
		delete(additionalProperties, "created")
		delete(additionalProperties, "modified")
		delete(additionalProperties, "description")
		delete(additionalProperties, "type")
		delete(additionalProperties, "requestStatus")
		delete(additionalProperties, "identityRequestId")
		delete(additionalProperties, "ownerRef")
		delete(additionalProperties, "requestCommentsRequired")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRequestableObject struct {
	value *RequestableObject
	isSet bool
}

func (v NullableRequestableObject) Get() *RequestableObject {
	return v.value
}

func (v *NullableRequestableObject) Set(val *RequestableObject) {
	v.value = val
	v.isSet = true
}

func (v NullableRequestableObject) IsSet() bool {
	return v.isSet
}

func (v *NullableRequestableObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRequestableObject(val *RequestableObject) *NullableRequestableObject {
	return &NullableRequestableObject{value: val, isSet: true}
}

func (v NullableRequestableObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRequestableObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


