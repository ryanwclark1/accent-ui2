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

// checks if the UserRelationCallPickupTargets type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRelationCallPickupTargets{}

// UserRelationCallPickupTargets struct for UserRelationCallPickupTargets
type UserRelationCallPickupTargets struct {
	CallPickupTargetUsers []UserRelationBase `json:"call_pickup_target_users,omitempty"`
}

// NewUserRelationCallPickupTargets instantiates a new UserRelationCallPickupTargets object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRelationCallPickupTargets() *UserRelationCallPickupTargets {
	this := UserRelationCallPickupTargets{}
	return &this
}

// NewUserRelationCallPickupTargetsWithDefaults instantiates a new UserRelationCallPickupTargets object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRelationCallPickupTargetsWithDefaults() *UserRelationCallPickupTargets {
	this := UserRelationCallPickupTargets{}
	return &this
}

// GetCallPickupTargetUsers returns the CallPickupTargetUsers field value if set, zero value otherwise.
func (o *UserRelationCallPickupTargets) GetCallPickupTargetUsers() []UserRelationBase {
	if o == nil || IsNil(o.CallPickupTargetUsers) {
		var ret []UserRelationBase
		return ret
	}
	return o.CallPickupTargetUsers
}

// GetCallPickupTargetUsersOk returns a tuple with the CallPickupTargetUsers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelationCallPickupTargets) GetCallPickupTargetUsersOk() ([]UserRelationBase, bool) {
	if o == nil || IsNil(o.CallPickupTargetUsers) {
		return nil, false
	}
	return o.CallPickupTargetUsers, true
}

// HasCallPickupTargetUsers returns a boolean if a field has been set.
func (o *UserRelationCallPickupTargets) HasCallPickupTargetUsers() bool {
	if o != nil && !IsNil(o.CallPickupTargetUsers) {
		return true
	}

	return false
}

// SetCallPickupTargetUsers gets a reference to the given []UserRelationBase and assigns it to the CallPickupTargetUsers field.
func (o *UserRelationCallPickupTargets) SetCallPickupTargetUsers(v []UserRelationBase) {
	o.CallPickupTargetUsers = v
}

func (o UserRelationCallPickupTargets) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRelationCallPickupTargets) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallPickupTargetUsers) {
		toSerialize["call_pickup_target_users"] = o.CallPickupTargetUsers
	}
	return toSerialize, nil
}

type NullableUserRelationCallPickupTargets struct {
	value *UserRelationCallPickupTargets
	isSet bool
}

func (v NullableUserRelationCallPickupTargets) Get() *UserRelationCallPickupTargets {
	return v.value
}

func (v *NullableUserRelationCallPickupTargets) Set(val *UserRelationCallPickupTargets) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRelationCallPickupTargets) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRelationCallPickupTargets) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRelationCallPickupTargets(val *UserRelationCallPickupTargets) *NullableUserRelationCallPickupTargets {
	return &NullableUserRelationCallPickupTargets{value: val, isSet: true}
}

func (v NullableUserRelationCallPickupTargets) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRelationCallPickupTargets) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
