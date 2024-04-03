/*
accent-auth

Accent's authentication service

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"encoding/json"
)

// checks if the ExternalAuthUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExternalAuthUser{}

// ExternalAuthUser struct for ExternalAuthUser
type ExternalAuthUser struct {
	Uuid *string `json:"uuid,omitempty"`
}

// NewExternalAuthUser instantiates a new ExternalAuthUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExternalAuthUser() *ExternalAuthUser {
	this := ExternalAuthUser{}
	return &this
}

// NewExternalAuthUserWithDefaults instantiates a new ExternalAuthUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExternalAuthUserWithDefaults() *ExternalAuthUser {
	this := ExternalAuthUser{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ExternalAuthUser) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExternalAuthUser) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ExternalAuthUser) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ExternalAuthUser) SetUuid(v string) {
	o.Uuid = &v
}

func (o ExternalAuthUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExternalAuthUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableExternalAuthUser struct {
	value *ExternalAuthUser
	isSet bool
}

func (v NullableExternalAuthUser) Get() *ExternalAuthUser {
	return v.value
}

func (v *NullableExternalAuthUser) Set(val *ExternalAuthUser) {
	v.value = val
	v.isSet = true
}

func (v NullableExternalAuthUser) IsSet() bool {
	return v.isSet
}

func (v *NullableExternalAuthUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExternalAuthUser(val *ExternalAuthUser) *NullableExternalAuthUser {
	return &NullableExternalAuthUser{value: val, isSet: true}
}

func (v NullableExternalAuthUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExternalAuthUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
