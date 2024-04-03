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

// checks if the ContextRelationContexts type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ContextRelationContexts{}

// ContextRelationContexts struct for ContextRelationContexts
type ContextRelationContexts struct {
	Contexts []ContextRelationBase `json:"contexts,omitempty"`
}

// NewContextRelationContexts instantiates a new ContextRelationContexts object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewContextRelationContexts() *ContextRelationContexts {
	this := ContextRelationContexts{}
	return &this
}

// NewContextRelationContextsWithDefaults instantiates a new ContextRelationContexts object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewContextRelationContextsWithDefaults() *ContextRelationContexts {
	this := ContextRelationContexts{}
	return &this
}

// GetContexts returns the Contexts field value if set, zero value otherwise.
func (o *ContextRelationContexts) GetContexts() []ContextRelationBase {
	if o == nil || IsNil(o.Contexts) {
		var ret []ContextRelationBase
		return ret
	}
	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ContextRelationContexts) GetContextsOk() ([]ContextRelationBase, bool) {
	if o == nil || IsNil(o.Contexts) {
		return nil, false
	}
	return o.Contexts, true
}

// HasContexts returns a boolean if a field has been set.
func (o *ContextRelationContexts) HasContexts() bool {
	if o != nil && !IsNil(o.Contexts) {
		return true
	}

	return false
}

// SetContexts gets a reference to the given []ContextRelationBase and assigns it to the Contexts field.
func (o *ContextRelationContexts) SetContexts(v []ContextRelationBase) {
	o.Contexts = v
}

func (o ContextRelationContexts) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ContextRelationContexts) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Contexts) {
		toSerialize["contexts"] = o.Contexts
	}
	return toSerialize, nil
}

type NullableContextRelationContexts struct {
	value *ContextRelationContexts
	isSet bool
}

func (v NullableContextRelationContexts) Get() *ContextRelationContexts {
	return v.value
}

func (v *NullableContextRelationContexts) Set(val *ContextRelationContexts) {
	v.value = val
	v.isSet = true
}

func (v NullableContextRelationContexts) IsSet() bool {
	return v.isSet
}

func (v *NullableContextRelationContexts) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableContextRelationContexts(val *ContextRelationContexts) *NullableContextRelationContexts {
	return &NullableContextRelationContexts{value: val, isSet: true}
}

func (v NullableContextRelationContexts) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableContextRelationContexts) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
