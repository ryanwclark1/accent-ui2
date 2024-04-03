/*
accent-amid

Send AMI actions to Asterisk, providing token based authentication.

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amid

import (
	"encoding/json"
)

// checks if the ComponentWithStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ComponentWithStatus{}

// ComponentWithStatus struct for ComponentWithStatus
type ComponentWithStatus struct {
	Status *StatusValue `json:"status,omitempty"`
}

// NewComponentWithStatus instantiates a new ComponentWithStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewComponentWithStatus() *ComponentWithStatus {
	this := ComponentWithStatus{}
	return &this
}

// NewComponentWithStatusWithDefaults instantiates a new ComponentWithStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewComponentWithStatusWithDefaults() *ComponentWithStatus {
	this := ComponentWithStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *ComponentWithStatus) GetStatus() StatusValue {
	if o == nil || IsNil(o.Status) {
		var ret StatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ComponentWithStatus) GetStatusOk() (*StatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *ComponentWithStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusValue and assigns it to the Status field.
func (o *ComponentWithStatus) SetStatus(v StatusValue) {
	o.Status = &v
}

func (o ComponentWithStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ComponentWithStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	return toSerialize, nil
}

type NullableComponentWithStatus struct {
	value *ComponentWithStatus
	isSet bool
}

func (v NullableComponentWithStatus) Get() *ComponentWithStatus {
	return v.value
}

func (v *NullableComponentWithStatus) Set(val *ComponentWithStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableComponentWithStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableComponentWithStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableComponentWithStatus(val *ComponentWithStatus) *NullableComponentWithStatus {
	return &NullableComponentWithStatus{value: val, isSet: true}
}

func (v NullableComponentWithStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableComponentWithStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
