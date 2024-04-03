/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"encoding/json"
)

// checks if the UserFuncKeyTemplate type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserFuncKeyTemplate{}

// UserFuncKeyTemplate Association between a User and a FuncKey Template
type UserFuncKeyTemplate struct {
	// FuncKey Template's ID
	TemplateId *int32 `json:"template_id,omitempty"`
	// User's ID
	UserId *int32 `json:"user_id,omitempty"`
}

// NewUserFuncKeyTemplate instantiates a new UserFuncKeyTemplate object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserFuncKeyTemplate() *UserFuncKeyTemplate {
	this := UserFuncKeyTemplate{}
	return &this
}

// NewUserFuncKeyTemplateWithDefaults instantiates a new UserFuncKeyTemplate object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserFuncKeyTemplateWithDefaults() *UserFuncKeyTemplate {
	this := UserFuncKeyTemplate{}
	return &this
}

// GetTemplateId returns the TemplateId field value if set, zero value otherwise.
func (o *UserFuncKeyTemplate) GetTemplateId() int32 {
	if o == nil || IsNil(o.TemplateId) {
		var ret int32
		return ret
	}
	return *o.TemplateId
}

// GetTemplateIdOk returns a tuple with the TemplateId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFuncKeyTemplate) GetTemplateIdOk() (*int32, bool) {
	if o == nil || IsNil(o.TemplateId) {
		return nil, false
	}
	return o.TemplateId, true
}

// HasTemplateId returns a boolean if a field has been set.
func (o *UserFuncKeyTemplate) HasTemplateId() bool {
	if o != nil && !IsNil(o.TemplateId) {
		return true
	}

	return false
}

// SetTemplateId gets a reference to the given int32 and assigns it to the TemplateId field.
func (o *UserFuncKeyTemplate) SetTemplateId(v int32) {
	o.TemplateId = &v
}

// GetUserId returns the UserId field value if set, zero value otherwise.
func (o *UserFuncKeyTemplate) GetUserId() int32 {
	if o == nil || IsNil(o.UserId) {
		var ret int32
		return ret
	}
	return *o.UserId
}

// GetUserIdOk returns a tuple with the UserId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserFuncKeyTemplate) GetUserIdOk() (*int32, bool) {
	if o == nil || IsNil(o.UserId) {
		return nil, false
	}
	return o.UserId, true
}

// HasUserId returns a boolean if a field has been set.
func (o *UserFuncKeyTemplate) HasUserId() bool {
	if o != nil && !IsNil(o.UserId) {
		return true
	}

	return false
}

// SetUserId gets a reference to the given int32 and assigns it to the UserId field.
func (o *UserFuncKeyTemplate) SetUserId(v int32) {
	o.UserId = &v
}

func (o UserFuncKeyTemplate) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserFuncKeyTemplate) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.TemplateId) {
		toSerialize["template_id"] = o.TemplateId
	}
	if !IsNil(o.UserId) {
		toSerialize["user_id"] = o.UserId
	}
	return toSerialize, nil
}

type NullableUserFuncKeyTemplate struct {
	value *UserFuncKeyTemplate
	isSet bool
}

func (v NullableUserFuncKeyTemplate) Get() *UserFuncKeyTemplate {
	return v.value
}

func (v *NullableUserFuncKeyTemplate) Set(val *UserFuncKeyTemplate) {
	v.value = val
	v.isSet = true
}

func (v NullableUserFuncKeyTemplate) IsSet() bool {
	return v.isSet
}

func (v *NullableUserFuncKeyTemplate) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserFuncKeyTemplate(val *UserFuncKeyTemplate) *NullableUserFuncKeyTemplate {
	return &NullableUserFuncKeyTemplate{value: val, isSet: true}
}

func (v NullableUserFuncKeyTemplate) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserFuncKeyTemplate) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
