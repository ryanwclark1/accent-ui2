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

// checks if the Locale type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Locale{}

// Locale struct for Locale
type Locale struct {
	// The current locale. For example: `en_US`
	Param *string `json:"param,omitempty"`
}

// NewLocale instantiates a new Locale object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewLocale() *Locale {
	this := Locale{}
	return &this
}

// NewLocaleWithDefaults instantiates a new Locale object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewLocaleWithDefaults() *Locale {
	this := Locale{}
	return &this
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *Locale) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Locale) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *Locale) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *Locale) SetParam(v string) {
	o.Param = &v
}

func (o Locale) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Locale) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return toSerialize, nil
}

type NullableLocale struct {
	value *Locale
	isSet bool
}

func (v NullableLocale) Get() *Locale {
	return v.value
}

func (v *NullableLocale) Set(val *Locale) {
	v.value = val
	v.isSet = true
}

func (v NullableLocale) IsSet() bool {
	return v.isSet
}

func (v *NullableLocale) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLocale(val *Locale) *NullableLocale {
	return &NullableLocale{value: val, isSet: true}
}

func (v NullableLocale) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLocale) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
