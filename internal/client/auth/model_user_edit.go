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

// checks if the UserEdit type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserEdit{}

// UserEdit struct for UserEdit
type UserEdit struct {
	Enabled *bool `json:"enabled,omitempty"`
	// The user's firstname
	Firstname *string `json:"firstname,omitempty"`
	// The user's lastname
	Lastname *string `json:"lastname,omitempty"`
	Purpose  *string `json:"purpose,omitempty"`
	// The username that will identify that new username
	Username *string `json:"username,omitempty"`
}

// NewUserEdit instantiates a new UserEdit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserEdit() *UserEdit {
	this := UserEdit{}
	var purpose string = "user"
	this.Purpose = &purpose
	return &this
}

// NewUserEditWithDefaults instantiates a new UserEdit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserEditWithDefaults() *UserEdit {
	this := UserEdit{}
	var purpose string = "user"
	this.Purpose = &purpose
	return &this
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *UserEdit) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEdit) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *UserEdit) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *UserEdit) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserEdit) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEdit) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserEdit) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserEdit) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserEdit) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEdit) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserEdit) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserEdit) SetLastname(v string) {
	o.Lastname = &v
}

// GetPurpose returns the Purpose field value if set, zero value otherwise.
func (o *UserEdit) GetPurpose() string {
	if o == nil || IsNil(o.Purpose) {
		var ret string
		return ret
	}
	return *o.Purpose
}

// GetPurposeOk returns a tuple with the Purpose field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEdit) GetPurposeOk() (*string, bool) {
	if o == nil || IsNil(o.Purpose) {
		return nil, false
	}
	return o.Purpose, true
}

// HasPurpose returns a boolean if a field has been set.
func (o *UserEdit) HasPurpose() bool {
	if o != nil && !IsNil(o.Purpose) {
		return true
	}

	return false
}

// SetPurpose gets a reference to the given string and assigns it to the Purpose field.
func (o *UserEdit) SetPurpose(v string) {
	o.Purpose = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *UserEdit) GetUsername() string {
	if o == nil || IsNil(o.Username) {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserEdit) GetUsernameOk() (*string, bool) {
	if o == nil || IsNil(o.Username) {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *UserEdit) HasUsername() bool {
	if o != nil && !IsNil(o.Username) {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *UserEdit) SetUsername(v string) {
	o.Username = &v
}

func (o UserEdit) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserEdit) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Purpose) {
		toSerialize["purpose"] = o.Purpose
	}
	if !IsNil(o.Username) {
		toSerialize["username"] = o.Username
	}
	return toSerialize, nil
}

type NullableUserEdit struct {
	value *UserEdit
	isSet bool
}

func (v NullableUserEdit) Get() *UserEdit {
	return v.value
}

func (v *NullableUserEdit) Set(val *UserEdit) {
	v.value = val
	v.isSet = true
}

func (v NullableUserEdit) IsSet() bool {
	return v.isSet
}

func (v *NullableUserEdit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserEdit(val *UserEdit) *NullableUserEdit {
	return &NullableUserEdit{value: val, isSet: true}
}

func (v NullableUserEdit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserEdit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
