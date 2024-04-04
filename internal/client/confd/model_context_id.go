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

// checks if the ContextId type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextId{}

// ContextId struct for ContextId
type ContextId struct {
	Id *int32 `json:"id,omitempty"`
}

// NewContextId instantiates a new ContextId object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextId() *ContextId {
	this := ContextId{}
	return &this
}

// NewContextIdWithDefaults instantiates a new ContextId object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextIdWithDefaults() *ContextId {
	this := ContextId{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *ContextId) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextId) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *ContextId) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *ContextId) SetId(v int32) {
	o.Id = &v
}

func (o ContextId) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextId) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableContextId struct {
	value *ContextId
	isSet bool
}

func (v NullableContextId) Get() *ContextId {
	return v.value
}

func (v *NullableContextId) Set(val *ContextId) {
	v.value = val
	v.isSet = true
}

func (v NullableContextId) IsSet() bool {
	return v.isSet
}

func (v *NullableContextId) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextId(val *ContextId) *NullableContextId {
	return &NullableContextId{value: val, isSet: true}
}

func (v NullableContextId) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextId) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
