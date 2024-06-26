/*
accent-amid

Send AMI actions to Asterisk, providing token based authentication.

API version: 0.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package amid

import (
	"encoding/json"
)

// checks if the CommandResponse type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CommandResponse{}

// CommandResponse struct for CommandResponse
type CommandResponse struct {
	Response []string `json:"response,omitempty"`
}

// NewCommandResponse instantiates a new CommandResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommandResponse() *CommandResponse {
	this := CommandResponse{}
	return &this
}

// NewCommandResponseWithDefaults instantiates a new CommandResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommandResponseWithDefaults() *CommandResponse {
	this := CommandResponse{}
	return &this
}

// GetResponse returns the Response field value if set, zero value otherwise.
func (o *CommandResponse) GetResponse() []string {
	if o == nil || IsNil(o.Response) {
		var ret []string
		return ret
	}
	return o.Response
}

// GetResponseOk returns a tuple with the Response field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommandResponse) GetResponseOk() ([]string, bool) {
	if o == nil || IsNil(o.Response) {
		return nil, false
	}
	return o.Response, true
}

// HasResponse returns a boolean if a field has been set.
func (o *CommandResponse) HasResponse() bool {
	if o != nil && !IsNil(o.Response) {
		return true
	}

	return false
}

// SetResponse gets a reference to the given []string and assigns it to the Response field.
func (o *CommandResponse) SetResponse(v []string) {
	o.Response = v
}

func (o CommandResponse) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CommandResponse) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Response) {
		toSerialize["response"] = o.Response
	}
	return toSerialize, nil
}

type NullableCommandResponse struct {
	value *CommandResponse
	isSet bool
}

func (v NullableCommandResponse) Get() *CommandResponse {
	return v.value
}

func (v *NullableCommandResponse) Set(val *CommandResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCommandResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCommandResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommandResponse(val *CommandResponse) *NullableCommandResponse {
	return &NullableCommandResponse{value: val, isSet: true}
}

func (v NullableCommandResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommandResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
