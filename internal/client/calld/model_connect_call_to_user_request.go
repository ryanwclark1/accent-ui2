/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"encoding/json"
)

// checks if the ConnectCallToUserRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ConnectCallToUserRequest{}

// ConnectCallToUserRequest struct for ConnectCallToUserRequest
type ConnectCallToUserRequest struct {
	// timeout in seconds for the dial attempt to the targeted user, or null for no timeout(infinite ring time). Omission leads to a default timeout of 30s.
	Timeout *int32 `json:"timeout,omitempty"`
}

// NewConnectCallToUserRequest instantiates a new ConnectCallToUserRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewConnectCallToUserRequest() *ConnectCallToUserRequest {
	this := ConnectCallToUserRequest{}
	return &this
}

// NewConnectCallToUserRequestWithDefaults instantiates a new ConnectCallToUserRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewConnectCallToUserRequestWithDefaults() *ConnectCallToUserRequest {
	this := ConnectCallToUserRequest{}
	return &this
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *ConnectCallToUserRequest) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ConnectCallToUserRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *ConnectCallToUserRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *ConnectCallToUserRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o ConnectCallToUserRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ConnectCallToUserRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

type NullableConnectCallToUserRequest struct {
	value *ConnectCallToUserRequest
	isSet bool
}

func (v NullableConnectCallToUserRequest) Get() *ConnectCallToUserRequest {
	return v.value
}

func (v *NullableConnectCallToUserRequest) Set(val *ConnectCallToUserRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableConnectCallToUserRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableConnectCallToUserRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableConnectCallToUserRequest(val *ConnectCallToUserRequest) *NullableConnectCallToUserRequest {
	return &NullableConnectCallToUserRequest{value: val, isSet: true}
}

func (v NullableConnectCallToUserRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableConnectCallToUserRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
