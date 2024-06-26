/*
accent-chatd

Control your message and presence from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chatd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserMessagePOST type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserMessagePOST{}

// UserMessagePOST struct for UserMessagePOST
type UserMessagePOST struct {
	// Alias/nickname of the sender
	Alias string `json:"alias"`
	// The content of the message
	Content string `json:"content"`
}

type _UserMessagePOST UserMessagePOST

// NewUserMessagePOST instantiates a new UserMessagePOST object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserMessagePOST(alias string, content string) *UserMessagePOST {
	this := UserMessagePOST{}
	this.Alias = alias
	this.Content = content
	return &this
}

// NewUserMessagePOSTWithDefaults instantiates a new UserMessagePOST object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserMessagePOSTWithDefaults() *UserMessagePOST {
	this := UserMessagePOST{}
	return &this
}

// GetAlias returns the Alias field value
func (o *UserMessagePOST) GetAlias() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Alias
}

// GetAliasOk returns a tuple with the Alias field value
// and a boolean to check if the value has been set.
func (o *UserMessagePOST) GetAliasOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Alias, true
}

// SetAlias sets field value
func (o *UserMessagePOST) SetAlias(v string) {
	o.Alias = v
}

// GetContent returns the Content field value
func (o *UserMessagePOST) GetContent() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Content
}

// GetContentOk returns a tuple with the Content field value
// and a boolean to check if the value has been set.
func (o *UserMessagePOST) GetContentOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Content, true
}

// SetContent sets field value
func (o *UserMessagePOST) SetContent(v string) {
	o.Content = v
}

func (o UserMessagePOST) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserMessagePOST) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["alias"] = o.Alias
	toSerialize["content"] = o.Content
	return toSerialize, nil
}

func (o *UserMessagePOST) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"alias",
		"content",
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

	varUserMessagePOST := _UserMessagePOST{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserMessagePOST)

	if err != nil {
		return err
	}

	*o = UserMessagePOST(varUserMessagePOST)

	return err
}

type NullableUserMessagePOST struct {
	value *UserMessagePOST
	isSet bool
}

func (v NullableUserMessagePOST) Get() *UserMessagePOST {
	return v.value
}

func (v *NullableUserMessagePOST) Set(val *UserMessagePOST) {
	v.value = val
	v.isSet = true
}

func (v NullableUserMessagePOST) IsSet() bool {
	return v.isSet
}

func (v *NullableUserMessagePOST) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserMessagePOST(val *UserMessagePOST) *NullableUserMessagePOST {
	return &NullableUserMessagePOST{value: val, isSet: true}
}

func (v NullableUserMessagePOST) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserMessagePOST) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
