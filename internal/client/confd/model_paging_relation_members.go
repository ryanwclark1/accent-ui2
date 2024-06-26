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

// checks if the PagingRelationMembers type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PagingRelationMembers{}

// PagingRelationMembers struct for PagingRelationMembers
type PagingRelationMembers struct {
	Members *PagingRelationMemberUsers `json:"members,omitempty"`
}

// NewPagingRelationMembers instantiates a new PagingRelationMembers object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPagingRelationMembers() *PagingRelationMembers {
	this := PagingRelationMembers{}
	return &this
}

// NewPagingRelationMembersWithDefaults instantiates a new PagingRelationMembers object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPagingRelationMembersWithDefaults() *PagingRelationMembers {
	this := PagingRelationMembers{}
	return &this
}

// GetMembers returns the Members field value if set, zero value otherwise.
func (o *PagingRelationMembers) GetMembers() PagingRelationMemberUsers {
	if o == nil || IsNil(o.Members) {
		var ret PagingRelationMemberUsers
		return ret
	}
	return *o.Members
}

// GetMembersOk returns a tuple with the Members field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PagingRelationMembers) GetMembersOk() (*PagingRelationMemberUsers, bool) {
	if o == nil || IsNil(o.Members) {
		return nil, false
	}
	return o.Members, true
}

// HasMembers returns a boolean if a field has been set.
func (o *PagingRelationMembers) HasMembers() bool {
	if o != nil && !IsNil(o.Members) {
		return true
	}

	return false
}

// SetMembers gets a reference to the given PagingRelationMemberUsers and assigns it to the Members field.
func (o *PagingRelationMembers) SetMembers(v PagingRelationMemberUsers) {
	o.Members = &v
}

func (o PagingRelationMembers) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PagingRelationMembers) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Members) {
		toSerialize["members"] = o.Members
	}
	return toSerialize, nil
}

type NullablePagingRelationMembers struct {
	value *PagingRelationMembers
	isSet bool
}

func (v NullablePagingRelationMembers) Get() *PagingRelationMembers {
	return v.value
}

func (v *NullablePagingRelationMembers) Set(val *PagingRelationMembers) {
	v.value = val
	v.isSet = true
}

func (v NullablePagingRelationMembers) IsSet() bool {
	return v.isSet
}

func (v *NullablePagingRelationMembers) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePagingRelationMembers(val *PagingRelationMembers) *NullablePagingRelationMembers {
	return &NullablePagingRelationMembers{value: val, isSet: true}
}

func (v NullablePagingRelationMembers) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePagingRelationMembers) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
