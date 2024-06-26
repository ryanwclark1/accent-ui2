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

// checks if the IncallRelationSchedules type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IncallRelationSchedules{}

// IncallRelationSchedules struct for IncallRelationSchedules
type IncallRelationSchedules struct {
	Schedules []ScheduleRelationBase `json:"schedules,omitempty"`
}

// NewIncallRelationSchedules instantiates a new IncallRelationSchedules object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIncallRelationSchedules() *IncallRelationSchedules {
	this := IncallRelationSchedules{}
	return &this
}

// NewIncallRelationSchedulesWithDefaults instantiates a new IncallRelationSchedules object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIncallRelationSchedulesWithDefaults() *IncallRelationSchedules {
	this := IncallRelationSchedules{}
	return &this
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *IncallRelationSchedules) GetSchedules() []ScheduleRelationBase {
	if o == nil || IsNil(o.Schedules) {
		var ret []ScheduleRelationBase
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IncallRelationSchedules) GetSchedulesOk() ([]ScheduleRelationBase, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *IncallRelationSchedules) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []ScheduleRelationBase and assigns it to the Schedules field.
func (o *IncallRelationSchedules) SetSchedules(v []ScheduleRelationBase) {
	o.Schedules = v
}

func (o IncallRelationSchedules) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IncallRelationSchedules) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	return toSerialize, nil
}

type NullableIncallRelationSchedules struct {
	value *IncallRelationSchedules
	isSet bool
}

func (v NullableIncallRelationSchedules) Get() *IncallRelationSchedules {
	return v.value
}

func (v *NullableIncallRelationSchedules) Set(val *IncallRelationSchedules) {
	v.value = val
	v.isSet = true
}

func (v NullableIncallRelationSchedules) IsSet() bool {
	return v.isSet
}

func (v *NullableIncallRelationSchedules) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIncallRelationSchedules(val *IncallRelationSchedules) *NullableIncallRelationSchedules {
	return &NullableIncallRelationSchedules{value: val, isSet: true}
}

func (v NullableIncallRelationSchedules) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIncallRelationSchedules) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
