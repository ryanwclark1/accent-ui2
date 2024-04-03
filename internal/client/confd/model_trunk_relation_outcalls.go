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

// checks if the TrunkRelationOutcalls type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &TrunkRelationOutcalls{}

// TrunkRelationOutcalls struct for TrunkRelationOutcalls
type TrunkRelationOutcalls struct {
	Outcalls []OutcallRelationBase `json:"outcalls,omitempty"`
}

// NewTrunkRelationOutcalls instantiates a new TrunkRelationOutcalls object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTrunkRelationOutcalls() *TrunkRelationOutcalls {
	this := TrunkRelationOutcalls{}
	return &this
}

// NewTrunkRelationOutcallsWithDefaults instantiates a new TrunkRelationOutcalls object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTrunkRelationOutcallsWithDefaults() *TrunkRelationOutcalls {
	this := TrunkRelationOutcalls{}
	return &this
}

// GetOutcalls returns the Outcalls field value if set, zero value otherwise.
func (o *TrunkRelationOutcalls) GetOutcalls() []OutcallRelationBase {
	if o == nil || IsNil(o.Outcalls) {
		var ret []OutcallRelationBase
		return ret
	}
	return o.Outcalls
}

// GetOutcallsOk returns a tuple with the Outcalls field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TrunkRelationOutcalls) GetOutcallsOk() ([]OutcallRelationBase, bool) {
	if o == nil || IsNil(o.Outcalls) {
		return nil, false
	}
	return o.Outcalls, true
}

// HasOutcalls returns a boolean if a field has been set.
func (o *TrunkRelationOutcalls) HasOutcalls() bool {
	if o != nil && !IsNil(o.Outcalls) {
		return true
	}

	return false
}

// SetOutcalls gets a reference to the given []OutcallRelationBase and assigns it to the Outcalls field.
func (o *TrunkRelationOutcalls) SetOutcalls(v []OutcallRelationBase) {
	o.Outcalls = v
}

func (o TrunkRelationOutcalls) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o TrunkRelationOutcalls) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Outcalls) {
		toSerialize["outcalls"] = o.Outcalls
	}
	return toSerialize, nil
}

type NullableTrunkRelationOutcalls struct {
	value *TrunkRelationOutcalls
	isSet bool
}

func (v NullableTrunkRelationOutcalls) Get() *TrunkRelationOutcalls {
	return v.value
}

func (v *NullableTrunkRelationOutcalls) Set(val *TrunkRelationOutcalls) {
	v.value = val
	v.isSet = true
}

func (v NullableTrunkRelationOutcalls) IsSet() bool {
	return v.isSet
}

func (v *NullableTrunkRelationOutcalls) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTrunkRelationOutcalls(val *TrunkRelationOutcalls) *NullableTrunkRelationOutcalls {
	return &NullableTrunkRelationOutcalls{value: val, isSet: true}
}

func (v NullableTrunkRelationOutcalls) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTrunkRelationOutcalls) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
