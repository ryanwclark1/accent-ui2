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

// checks if the BackendList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &BackendList{}

// BackendList struct for BackendList
type BackendList struct {
	Data []string `json:"data,omitempty"`
}

// NewBackendList instantiates a new BackendList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackendList() *BackendList {
	this := BackendList{}
	return &this
}

// NewBackendListWithDefaults instantiates a new BackendList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackendListWithDefaults() *BackendList {
	this := BackendList{}
	return &this
}

// GetData returns the Data field value if set, zero value otherwise.
func (o *BackendList) GetData() []string {
	if o == nil || IsNil(o.Data) {
		var ret []string
		return ret
	}
	return o.Data
}

// GetDataOk returns a tuple with the Data field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BackendList) GetDataOk() ([]string, bool) {
	if o == nil || IsNil(o.Data) {
		return nil, false
	}
	return o.Data, true
}

// HasData returns a boolean if a field has been set.
func (o *BackendList) HasData() bool {
	if o != nil && !IsNil(o.Data) {
		return true
	}

	return false
}

// SetData gets a reference to the given []string and assigns it to the Data field.
func (o *BackendList) SetData(v []string) {
	o.Data = v
}

func (o BackendList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o BackendList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Data) {
		toSerialize["data"] = o.Data
	}
	return toSerialize, nil
}

type NullableBackendList struct {
	value *BackendList
	isSet bool
}

func (v NullableBackendList) Get() *BackendList {
	return v.value
}

func (v *NullableBackendList) Set(val *BackendList) {
	v.value = val
	v.isSet = true
}

func (v NullableBackendList) IsSet() bool {
	return v.isSet
}

func (v *NullableBackendList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackendList(val *BackendList) *NullableBackendList {
	return &NullableBackendList{value: val, isSet: true}
}

func (v NullableBackendList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackendList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
