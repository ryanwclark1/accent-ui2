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

// checks if the SessionResult type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SessionResult{}

// SessionResult struct for SessionResult
type SessionResult struct {
	Mobile     *bool   `json:"mobile,omitempty"`
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	UserUuid   *string `json:"user_uuid,omitempty"`
	Uuid       *string `json:"uuid,omitempty"`
}

// NewSessionResult instantiates a new SessionResult object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSessionResult() *SessionResult {
	this := SessionResult{}
	return &this
}

// NewSessionResultWithDefaults instantiates a new SessionResult object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSessionResultWithDefaults() *SessionResult {
	this := SessionResult{}
	return &this
}

// GetMobile returns the Mobile field value if set, zero value otherwise.
func (o *SessionResult) GetMobile() bool {
	if o == nil || IsNil(o.Mobile) {
		var ret bool
		return ret
	}
	return *o.Mobile
}

// GetMobileOk returns a tuple with the Mobile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResult) GetMobileOk() (*bool, bool) {
	if o == nil || IsNil(o.Mobile) {
		return nil, false
	}
	return o.Mobile, true
}

// HasMobile returns a boolean if a field has been set.
func (o *SessionResult) HasMobile() bool {
	if o != nil && !IsNil(o.Mobile) {
		return true
	}

	return false
}

// SetMobile gets a reference to the given bool and assigns it to the Mobile field.
func (o *SessionResult) SetMobile(v bool) {
	o.Mobile = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *SessionResult) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResult) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *SessionResult) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *SessionResult) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *SessionResult) GetUserUuid() string {
	if o == nil || IsNil(o.UserUuid) {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResult) GetUserUuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserUuid) {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *SessionResult) HasUserUuid() bool {
	if o != nil && !IsNil(o.UserUuid) {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *SessionResult) SetUserUuid(v string) {
	o.UserUuid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *SessionResult) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SessionResult) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *SessionResult) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *SessionResult) SetUuid(v string) {
	o.Uuid = &v
}

func (o SessionResult) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SessionResult) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Mobile) {
		toSerialize["mobile"] = o.Mobile
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.UserUuid) {
		toSerialize["user_uuid"] = o.UserUuid
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableSessionResult struct {
	value *SessionResult
	isSet bool
}

func (v NullableSessionResult) Get() *SessionResult {
	return v.value
}

func (v *NullableSessionResult) Set(val *SessionResult) {
	v.value = val
	v.isSet = true
}

func (v NullableSessionResult) IsSet() bool {
	return v.isSet
}

func (v *NullableSessionResult) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSessionResult(val *SessionResult) *NullableSessionResult {
	return &NullableSessionResult{value: val, isSet: true}
}

func (v NullableSessionResult) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSessionResult) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
