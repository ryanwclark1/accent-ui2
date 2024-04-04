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

// checks if the ConferenceRelationIncalls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConferenceRelationIncalls{}

// ConferenceRelationIncalls struct for ConferenceRelationIncalls
type ConferenceRelationIncalls struct {
	Incalls []ConferenceRelationIncall `json:"incalls,omitempty"`
}

// NewConferenceRelationIncalls instantiates a new ConferenceRelationIncalls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConferenceRelationIncalls() *ConferenceRelationIncalls {
	this := ConferenceRelationIncalls{}
	return &this
}

// NewConferenceRelationIncallsWithDefaults instantiates a new ConferenceRelationIncalls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConferenceRelationIncallsWithDefaults() *ConferenceRelationIncalls {
	this := ConferenceRelationIncalls{}
	return &this
}

// GetIncalls returns the Incalls field value if set, zero value otherwise.
func (o *ConferenceRelationIncalls) GetIncalls() []ConferenceRelationIncall {
	if o == nil || IsNil(o.Incalls) {
		var ret []ConferenceRelationIncall
		return ret
	}
	return o.Incalls
}

// GetIncallsOk returns a tuple with the Incalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConferenceRelationIncalls) GetIncallsOk() ([]ConferenceRelationIncall, bool) {
	if o == nil || IsNil(o.Incalls) {
		return nil, false
	}
	return o.Incalls, true
}

// HasIncalls returns a boolean if a field has been set.
func (o *ConferenceRelationIncalls) HasIncalls() bool {
	if o != nil && !IsNil(o.Incalls) {
		return true
	}

	return false
}

// SetIncalls gets a reference to the given []ConferenceRelationIncall and assigns it to the Incalls field.
func (o *ConferenceRelationIncalls) SetIncalls(v []ConferenceRelationIncall) {
	o.Incalls = v
}

func (o ConferenceRelationIncalls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConferenceRelationIncalls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Incalls) {
		toSerialize["incalls"] = o.Incalls
	}
	return toSerialize, nil
}

type NullableConferenceRelationIncalls struct {
	value *ConferenceRelationIncalls
	isSet bool
}

func (v NullableConferenceRelationIncalls) Get() *ConferenceRelationIncalls {
	return v.value
}

func (v *NullableConferenceRelationIncalls) Set(val *ConferenceRelationIncalls) {
	v.value = val
	v.isSet = true
}

func (v NullableConferenceRelationIncalls) IsSet() bool {
	return v.isSet
}

func (v *NullableConferenceRelationIncalls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConferenceRelationIncalls(val *ConferenceRelationIncalls) *NullableConferenceRelationIncalls {
	return &NullableConferenceRelationIncalls{value: val, isSet: true}
}

func (v NullableConferenceRelationIncalls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConferenceRelationIncalls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
