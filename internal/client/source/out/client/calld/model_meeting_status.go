/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"encoding/json"
)

// checks if the MeetingStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeetingStatus{}

// MeetingStatus struct for MeetingStatus
type MeetingStatus struct {
	// Wether the meeting is full of not
	Full *bool `json:"full,omitempty"`
}

// NewMeetingStatus instantiates a new MeetingStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeetingStatus() *MeetingStatus {
	this := MeetingStatus{}
	return &this
}

// NewMeetingStatusWithDefaults instantiates a new MeetingStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeetingStatusWithDefaults() *MeetingStatus {
	this := MeetingStatus{}
	return &this
}

// GetFull returns the Full field value if set, zero value otherwise.
func (o *MeetingStatus) GetFull() bool {
	if o == nil || IsNil(o.Full) {
		var ret bool
		return ret
	}
	return *o.Full
}

// GetFullOk returns a tuple with the Full field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingStatus) GetFullOk() (*bool, bool) {
	if o == nil || IsNil(o.Full) {
		return nil, false
	}
	return o.Full, true
}

// HasFull returns a boolean if a field has been set.
func (o *MeetingStatus) HasFull() bool {
	if o != nil && !IsNil(o.Full) {
		return true
	}

	return false
}

// SetFull gets a reference to the given bool and assigns it to the Full field.
func (o *MeetingStatus) SetFull(v bool) {
	o.Full = &v
}

func (o MeetingStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeetingStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Full) {
		toSerialize["full"] = o.Full
	}
	return toSerialize, nil
}

type NullableMeetingStatus struct {
	value *MeetingStatus
	isSet bool
}

func (v NullableMeetingStatus) Get() *MeetingStatus {
	return v.value
}

func (v *NullableMeetingStatus) Set(val *MeetingStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableMeetingStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableMeetingStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeetingStatus(val *MeetingStatus) *NullableMeetingStatus {
	return &NullableMeetingStatus{value: val, isSet: true}
}

func (v NullableMeetingStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeetingStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
