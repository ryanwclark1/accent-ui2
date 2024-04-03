/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"encoding/json"
)

// checks if the SwitchboardRelationFallbacks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SwitchboardRelationFallbacks{}

// SwitchboardRelationFallbacks struct for SwitchboardRelationFallbacks
type SwitchboardRelationFallbacks struct {
	Fallbacks *SwitchboardFallbacks `json:"fallbacks,omitempty"`
}

// NewSwitchboardRelationFallbacks instantiates a new SwitchboardRelationFallbacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSwitchboardRelationFallbacks() *SwitchboardRelationFallbacks {
	this := SwitchboardRelationFallbacks{}
	return &this
}

// NewSwitchboardRelationFallbacksWithDefaults instantiates a new SwitchboardRelationFallbacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSwitchboardRelationFallbacksWithDefaults() *SwitchboardRelationFallbacks {
	this := SwitchboardRelationFallbacks{}
	return &this
}

// GetFallbacks returns the Fallbacks field value if set, zero value otherwise.
func (o *SwitchboardRelationFallbacks) GetFallbacks() SwitchboardFallbacks {
	if o == nil || IsNil(o.Fallbacks) {
		var ret SwitchboardFallbacks
		return ret
	}
	return *o.Fallbacks
}

// GetFallbacksOk returns a tuple with the Fallbacks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SwitchboardRelationFallbacks) GetFallbacksOk() (*SwitchboardFallbacks, bool) {
	if o == nil || IsNil(o.Fallbacks) {
		return nil, false
	}
	return o.Fallbacks, true
}

// HasFallbacks returns a boolean if a field has been set.
func (o *SwitchboardRelationFallbacks) HasFallbacks() bool {
	if o != nil && !IsNil(o.Fallbacks) {
		return true
	}

	return false
}

// SetFallbacks gets a reference to the given SwitchboardFallbacks and assigns it to the Fallbacks field.
func (o *SwitchboardRelationFallbacks) SetFallbacks(v SwitchboardFallbacks) {
	o.Fallbacks = &v
}

func (o SwitchboardRelationFallbacks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SwitchboardRelationFallbacks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Fallbacks) {
		toSerialize["fallbacks"] = o.Fallbacks
	}
	return toSerialize, nil
}

type NullableSwitchboardRelationFallbacks struct {
	value *SwitchboardRelationFallbacks
	isSet bool
}

func (v NullableSwitchboardRelationFallbacks) Get() *SwitchboardRelationFallbacks {
	return v.value
}

func (v *NullableSwitchboardRelationFallbacks) Set(val *SwitchboardRelationFallbacks) {
	v.value = val
	v.isSet = true
}

func (v NullableSwitchboardRelationFallbacks) IsSet() bool {
	return v.isSet
}

func (v *NullableSwitchboardRelationFallbacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSwitchboardRelationFallbacks(val *SwitchboardRelationFallbacks) *NullableSwitchboardRelationFallbacks {
	return &NullableSwitchboardRelationFallbacks{value: val, isSet: true}
}

func (v NullableSwitchboardRelationFallbacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSwitchboardRelationFallbacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
