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

// checks if the CallFilterSurrogateUsers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallFilterSurrogateUsers{}

// CallFilterSurrogateUsers struct for CallFilterSurrogateUsers
type CallFilterSurrogateUsers struct {
	Users []CallFilterSurrogateUser `json:"users,omitempty"`
}

// NewCallFilterSurrogateUsers instantiates a new CallFilterSurrogateUsers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallFilterSurrogateUsers() *CallFilterSurrogateUsers {
	this := CallFilterSurrogateUsers{}
	return &this
}

// NewCallFilterSurrogateUsersWithDefaults instantiates a new CallFilterSurrogateUsers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallFilterSurrogateUsersWithDefaults() *CallFilterSurrogateUsers {
	this := CallFilterSurrogateUsers{}
	return &this
}

// GetUsers returns the Users field value if set, zero value otherwise.
func (o *CallFilterSurrogateUsers) GetUsers() []CallFilterSurrogateUser {
	if o == nil || IsNil(o.Users) {
		var ret []CallFilterSurrogateUser
		return ret
	}
	return o.Users
}

// GetUsersOk returns a tuple with the Users field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUsers) GetUsersOk() ([]CallFilterSurrogateUser, bool) {
	if o == nil || IsNil(o.Users) {
		return nil, false
	}
	return o.Users, true
}

// HasUsers returns a boolean if a field has been set.
func (o *CallFilterSurrogateUsers) HasUsers() bool {
	if o != nil && !IsNil(o.Users) {
		return true
	}

	return false
}

// SetUsers gets a reference to the given []CallFilterSurrogateUser and assigns it to the Users field.
func (o *CallFilterSurrogateUsers) SetUsers(v []CallFilterSurrogateUser) {
	o.Users = v
}

func (o CallFilterSurrogateUsers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallFilterSurrogateUsers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Users) {
		toSerialize["users"] = o.Users
	}
	return toSerialize, nil
}

type NullableCallFilterSurrogateUsers struct {
	value *CallFilterSurrogateUsers
	isSet bool
}

func (v NullableCallFilterSurrogateUsers) Get() *CallFilterSurrogateUsers {
	return v.value
}

func (v *NullableCallFilterSurrogateUsers) Set(val *CallFilterSurrogateUsers) {
	v.value = val
	v.isSet = true
}

func (v NullableCallFilterSurrogateUsers) IsSet() bool {
	return v.isSet
}

func (v *NullableCallFilterSurrogateUsers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallFilterSurrogateUsers(val *CallFilterSurrogateUsers) *NullableCallFilterSurrogateUsers {
	return &NullableCallFilterSurrogateUsers{value: val, isSet: true}
}

func (v NullableCallFilterSurrogateUsers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallFilterSurrogateUsers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
