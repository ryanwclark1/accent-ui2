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

// checks if the PostPasswordReset type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PostPasswordReset{}

// PostPasswordReset struct for PostPasswordReset
type PostPasswordReset struct {
	// The desired password
	Password string `json:"password"`
}

type _PostPasswordReset PostPasswordReset

// NewPostPasswordReset instantiates a new PostPasswordReset object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPostPasswordReset(password string) *PostPasswordReset {
	this := PostPasswordReset{}
	this.Password = password
	return &this
}

// NewPostPasswordResetWithDefaults instantiates a new PostPasswordReset object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPostPasswordResetWithDefaults() *PostPasswordReset {
	this := PostPasswordReset{}
	return &this
}

// GetPassword returns the Password field value
func (o *PostPasswordReset) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *PostPasswordReset) GetPasswordOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *PostPasswordReset) SetPassword(v string) {
	o.Password = v
}

func (o PostPasswordReset) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PostPasswordReset) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["password"] = o.Password
	return toSerialize, nil
}

func (o *PostPasswordReset) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"password",
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

	varPostPasswordReset := _PostPasswordReset{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varPostPasswordReset)

	if err != nil {
		return err
	}

	*o = PostPasswordReset(varPostPasswordReset)

	return err
}

type NullablePostPasswordReset struct {
	value *PostPasswordReset
	isSet bool
}

func (v NullablePostPasswordReset) Get() *PostPasswordReset {
	return v.value
}

func (v *NullablePostPasswordReset) Set(val *PostPasswordReset) {
	v.value = val
	v.isSet = true
}

func (v NullablePostPasswordReset) IsSet() bool {
	return v.isSet
}

func (v *NullablePostPasswordReset) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePostPasswordReset(val *PostPasswordReset) *NullablePostPasswordReset {
	return &NullablePostPasswordReset{value: val, isSet: true}
}

func (v NullablePostPasswordReset) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePostPasswordReset) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
