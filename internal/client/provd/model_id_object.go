/*
accent-provd

Provisioning application REST API

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package provd

import (
	"encoding/json"
)

// checks if the IdObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IdObject{}

// IdObject struct for IdObject
type IdObject struct {
	Id *string `json:"id,omitempty"`
}

// NewIdObject instantiates a new IdObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIdObject() *IdObject {
	this := IdObject{}
	return &this
}

// NewIdObjectWithDefaults instantiates a new IdObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIdObjectWithDefaults() *IdObject {
	this := IdObject{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *IdObject) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IdObject) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *IdObject) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *IdObject) SetId(v string) {
	o.Id = &v
}

func (o IdObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IdObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableIdObject struct {
	value *IdObject
	isSet bool
}

func (v NullableIdObject) Get() *IdObject {
	return v.value
}

func (v *NullableIdObject) Set(val *IdObject) {
	v.value = val
	v.isSet = true
}

func (v NullableIdObject) IsSet() bool {
	return v.isSet
}

func (v *NullableIdObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIdObject(val *IdObject) *NullableIdObject {
	return &NullableIdObject{value: val, isSet: true}
}

func (v NullableIdObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIdObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
