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

// checks if the UserRelationGroups type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRelationGroups{}

// UserRelationGroups struct for UserRelationGroups
type UserRelationGroups struct {
	Groups []GroupRelationBase `json:"groups,omitempty"`
}

// NewUserRelationGroups instantiates a new UserRelationGroups object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRelationGroups() *UserRelationGroups {
	this := UserRelationGroups{}
	return &this
}

// NewUserRelationGroupsWithDefaults instantiates a new UserRelationGroups object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRelationGroupsWithDefaults() *UserRelationGroups {
	this := UserRelationGroups{}
	return &this
}

// GetGroups returns the Groups field value if set, zero value otherwise.
func (o *UserRelationGroups) GetGroups() []GroupRelationBase {
	if o == nil || IsNil(o.Groups) {
		var ret []GroupRelationBase
		return ret
	}
	return o.Groups
}

// GetGroupsOk returns a tuple with the Groups field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelationGroups) GetGroupsOk() ([]GroupRelationBase, bool) {
	if o == nil || IsNil(o.Groups) {
		return nil, false
	}
	return o.Groups, true
}

// HasGroups returns a boolean if a field has been set.
func (o *UserRelationGroups) HasGroups() bool {
	if o != nil && !IsNil(o.Groups) {
		return true
	}

	return false
}

// SetGroups gets a reference to the given []GroupRelationBase and assigns it to the Groups field.
func (o *UserRelationGroups) SetGroups(v []GroupRelationBase) {
	o.Groups = v
}

func (o UserRelationGroups) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRelationGroups) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Groups) {
		toSerialize["groups"] = o.Groups
	}
	return toSerialize, nil
}

type NullableUserRelationGroups struct {
	value *UserRelationGroups
	isSet bool
}

func (v NullableUserRelationGroups) Get() *UserRelationGroups {
	return v.value
}

func (v *NullableUserRelationGroups) Set(val *UserRelationGroups) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRelationGroups) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRelationGroups) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRelationGroups(val *UserRelationGroups) *NullableUserRelationGroups {
	return &NullableUserRelationGroups{value: val, isSet: true}
}

func (v NullableUserRelationGroups) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRelationGroups) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
