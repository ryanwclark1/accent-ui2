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

// checks if the FuncKey type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FuncKey{}

// FuncKey Further documentation at https://accentvoice.io/uc-doc/api_sdk/rest_api/confd/func_keys
type FuncKey struct {
	Blf         *bool                  `json:"blf,omitempty"`
	Destination map[string]interface{} `json:"destination,omitempty"`
	Label       *string                `json:"label,omitempty"`
}

// NewFuncKey instantiates a new FuncKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuncKey() *FuncKey {
	this := FuncKey{}
	return &this
}

// NewFuncKeyWithDefaults instantiates a new FuncKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuncKeyWithDefaults() *FuncKey {
	this := FuncKey{}
	return &this
}

// GetBlf returns the Blf field value if set, zero value otherwise.
func (o *FuncKey) GetBlf() bool {
	if o == nil || IsNil(o.Blf) {
		var ret bool
		return ret
	}
	return *o.Blf
}

// GetBlfOk returns a tuple with the Blf field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKey) GetBlfOk() (*bool, bool) {
	if o == nil || IsNil(o.Blf) {
		return nil, false
	}
	return o.Blf, true
}

// HasBlf returns a boolean if a field has been set.
func (o *FuncKey) HasBlf() bool {
	if o != nil && !IsNil(o.Blf) {
		return true
	}

	return false
}

// SetBlf gets a reference to the given bool and assigns it to the Blf field.
func (o *FuncKey) SetBlf(v bool) {
	o.Blf = &v
}

// GetDestination returns the Destination field value if set, zero value otherwise.
func (o *FuncKey) GetDestination() map[string]interface{} {
	if o == nil || IsNil(o.Destination) {
		var ret map[string]interface{}
		return ret
	}
	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKey) GetDestinationOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Destination) {
		return map[string]interface{}{}, false
	}
	return o.Destination, true
}

// HasDestination returns a boolean if a field has been set.
func (o *FuncKey) HasDestination() bool {
	if o != nil && !IsNil(o.Destination) {
		return true
	}

	return false
}

// SetDestination gets a reference to the given map[string]interface{} and assigns it to the Destination field.
func (o *FuncKey) SetDestination(v map[string]interface{}) {
	o.Destination = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FuncKey) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKey) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FuncKey) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FuncKey) SetLabel(v string) {
	o.Label = &v
}

func (o FuncKey) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuncKey) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Blf) {
		toSerialize["blf"] = o.Blf
	}
	if !IsNil(o.Destination) {
		toSerialize["destination"] = o.Destination
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	return toSerialize, nil
}

type NullableFuncKey struct {
	value *FuncKey
	isSet bool
}

func (v NullableFuncKey) Get() *FuncKey {
	return v.value
}

func (v *NullableFuncKey) Set(val *FuncKey) {
	v.value = val
	v.isSet = true
}

func (v NullableFuncKey) IsSet() bool {
	return v.isSet
}

func (v *NullableFuncKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuncKey(val *FuncKey) *NullableFuncKey {
	return &NullableFuncKey{value: val, isSet: true}
}

func (v NullableFuncKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuncKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
