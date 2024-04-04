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

// checks if the ScheduleRelationGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ScheduleRelationGroups{}

// ScheduleRelationGroups struct for ScheduleRelationGroups
type ScheduleRelationGroups struct {
	Incalls []GroupRelationBase `json:"incalls,omitempty"`
}

// NewScheduleRelationGroups instantiates a new ScheduleRelationGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewScheduleRelationGroups() *ScheduleRelationGroups {
	this := ScheduleRelationGroups{}
	return &this
}

// NewScheduleRelationGroupsWithDefaults instantiates a new ScheduleRelationGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewScheduleRelationGroupsWithDefaults() *ScheduleRelationGroups {
	this := ScheduleRelationGroups{}
	return &this
}

// GetIncalls returns the Incalls field value if set, zero value otherwise.
func (o *ScheduleRelationGroups) GetIncalls() []GroupRelationBase {
	if o == nil || IsNil(o.Incalls) {
		var ret []GroupRelationBase
		return ret
	}
	return o.Incalls
}

// GetIncallsOk returns a tuple with the Incalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ScheduleRelationGroups) GetIncallsOk() ([]GroupRelationBase, bool) {
	if o == nil || IsNil(o.Incalls) {
		return nil, false
	}
	return o.Incalls, true
}

// HasIncalls returns a boolean if a field has been set.
func (o *ScheduleRelationGroups) HasIncalls() bool {
	if o != nil && !IsNil(o.Incalls) {
		return true
	}

	return false
}

// SetIncalls gets a reference to the given []GroupRelationBase and assigns it to the Incalls field.
func (o *ScheduleRelationGroups) SetIncalls(v []GroupRelationBase) {
	o.Incalls = v
}

func (o ScheduleRelationGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ScheduleRelationGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Incalls) {
		toSerialize["incalls"] = o.Incalls
	}
	return toSerialize, nil
}

type NullableScheduleRelationGroups struct {
	value *ScheduleRelationGroups
	isSet bool
}

func (v NullableScheduleRelationGroups) Get() *ScheduleRelationGroups {
	return v.value
}

func (v *NullableScheduleRelationGroups) Set(val *ScheduleRelationGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableScheduleRelationGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableScheduleRelationGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScheduleRelationGroups(val *ScheduleRelationGroups) *NullableScheduleRelationGroups {
	return &NullableScheduleRelationGroups{value: val, isSet: true}
}

func (v NullableScheduleRelationGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScheduleRelationGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
