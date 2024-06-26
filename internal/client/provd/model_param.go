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

// checks if the Param type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Param{}

// Param A configuration parameter
type Param struct {
	Param *ParamObject `json:"param,omitempty"`
}

// NewParam instantiates a new Param object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewParam() *Param {
	this := Param{}
	return &this
}

// NewParamWithDefaults instantiates a new Param object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewParamWithDefaults() *Param {
	this := Param{}
	return &this
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *Param) GetParam() ParamObject {
	if o == nil || IsNil(o.Param) {
		var ret ParamObject
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Param) GetParamOk() (*ParamObject, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *Param) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given ParamObject and assigns it to the Param field.
func (o *Param) SetParam(v ParamObject) {
	o.Param = &v
}

func (o Param) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Param) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return toSerialize, nil
}

type NullableParam struct {
	value *Param
	isSet bool
}

func (v NullableParam) Get() *Param {
	return v.value
}

func (v *NullableParam) Set(val *Param) {
	v.value = val
	v.isSet = true
}

func (v NullableParam) IsSet() bool {
	return v.isSet
}

func (v *NullableParam) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableParam(val *Param) *NullableParam {
	return &NullableParam{value: val, isSet: true}
}

func (v NullableParam) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableParam) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
