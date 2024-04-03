/*
accent-auth

Accent's authentication service

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package auth

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserRegister type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRegister{}

// UserRegister struct for UserRegister
type UserRegister struct {
	// The main email address of the new username
	EmailAddress string `json:"email_address"`
	// The user's firstname
	Firstname *string `json:"firstname,omitempty"`
	// The user's lastname
	Lastname *string `json:"lastname,omitempty"`
	// The password of the newly created username
	Password string `json:"password"`
	// The username that will identify that new username
	Username string `json:"username"`
}

type _UserRegister UserRegister

// NewUserRegister instantiates a new UserRegister object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRegister(emailAddress string, password string, username string) *UserRegister {
	this := UserRegister{}
	this.EmailAddress = emailAddress
	this.Password = password
	this.Username = username
	return &this
}

// NewUserRegisterWithDefaults instantiates a new UserRegister object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRegisterWithDefaults() *UserRegister {
	this := UserRegister{}
	return &this
}

// GetEmailAddress returns the EmailAddress field value
func (o *UserRegister) GetEmailAddress() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.EmailAddress
}

// GetEmailAddressOk returns a tuple with the EmailAddress field value
// and a boolean to check if the value has been set.
func (o *UserRegister) GetEmailAddressOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.EmailAddress, true
}

// SetEmailAddress sets field value
func (o *UserRegister) SetEmailAddress(v string) {
	o.EmailAddress = v
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *UserRegister) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegister) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *UserRegister) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *UserRegister) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *UserRegister) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRegister) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *UserRegister) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *UserRegister) SetLastname(v string) {
	o.Lastname = &v
}

// GetPassword returns the Password field value
func (o *UserRegister) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *UserRegister) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *UserRegister) SetPassword(v string) {
	o.Password = v
}

// GetUsername returns the Username field value
func (o *UserRegister) GetUsername() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Username
}

// GetUsernameOk returns a tuple with the Username field value
// and a boolean to check if the value has been set.
func (o *UserRegister) GetUsernameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Username, true
}

// SetUsername sets field value
func (o *UserRegister) SetUsername(v string) {
	o.Username = v
}

func (o UserRegister) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRegister) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["email_address"] = o.EmailAddress
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	toSerialize["password"] = o.Password
	toSerialize["username"] = o.Username
	return toSerialize, nil
}

func (o *UserRegister) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"email_address",
		"password",
		"username",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varUserRegister := _UserRegister{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRegister)

	if err != nil {
		return err
	}

	*o = UserRegister(varUserRegister)

	return err
}

type NullableUserRegister struct {
	value *UserRegister
	isSet bool
}

func (v NullableUserRegister) Get() *UserRegister {
	return v.value
}

func (v *NullableUserRegister) Set(val *UserRegister) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRegister) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRegister) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRegister(val *UserRegister) *NullableUserRegister {
	return &NullableUserRegister{value: val, isSet: true}
}

func (v NullableUserRegister) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRegister) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
