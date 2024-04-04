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

// checks if the ScheduleRelationUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleRelationUsers{}

// ScheduleRelationUsers struct for ScheduleRelationUsers
type ScheduleRelationUsers struct {
	Incalls []UserRelationBase `json:"incalls,omitempty"`
}

// NewScheduleRelationUsers instantiates a new ScheduleRelationUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleRelationUsers() *ScheduleRelationUsers {
	this := ScheduleRelationUsers{}
	return &this
}

// NewScheduleRelationUsersWithDefaults instantiates a new ScheduleRelationUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleRelationUsersWithDefaults() *ScheduleRelationUsers {
	this := ScheduleRelationUsers{}
	return &this
}

// GetIncalls returns the Incalls field value if set, zero value otherwise.
func (o *ScheduleRelationUsers) GetIncalls() []UserRelationBase {
	if o == nil || IsNil(o.Incalls) {
		var ret []UserRelationBase
		return ret
	}
	return o.Incalls
}

// GetIncallsOk returns a tuple with the Incalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRelationUsers) GetIncallsOk() ([]UserRelationBase, bool) {
	if o == nil || IsNil(o.Incalls) {
		return nil, false
	}
	return o.Incalls, true
}

// HasIncalls returns a boolean if a field has been set.
func (o *ScheduleRelationUsers) HasIncalls() bool {
	if o != nil && !IsNil(o.Incalls) {
		return true
	}

	return false
}

// SetIncalls gets a reference to the given []UserRelationBase and assigns it to the Incalls field.
func (o *ScheduleRelationUsers) SetIncalls(v []UserRelationBase) {
	o.Incalls = v
}

func (o ScheduleRelationUsers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleRelationUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Incalls) {
		toSerialize["incalls"] = o.Incalls
	}
	return toSerialize, nil
}

type NullableScheduleRelationUsers struct {
	value *ScheduleRelationUsers
	isSet bool
}

func (v NullableScheduleRelationUsers) Get() *ScheduleRelationUsers {
	return v.value
}

func (v *NullableScheduleRelationUsers) Set(val *ScheduleRelationUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleRelationUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleRelationUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleRelationUsers(val *ScheduleRelationUsers) *NullableScheduleRelationUsers {
	return &NullableScheduleRelationUsers{value: val, isSet: true}
}

func (v NullableScheduleRelationUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleRelationUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
