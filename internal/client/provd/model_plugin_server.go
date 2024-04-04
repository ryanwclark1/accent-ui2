/*
accent-provd

Provisioning application REST API

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package provd

import (
	"encoding/json"
)

// checks if the PluginServer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PluginServer{}

// PluginServer struct for PluginServer
type PluginServer struct {
	// The plugins repository URL
	Param *string `json:"param,omitempty"`
}

// NewPluginServer instantiates a new PluginServer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPluginServer() *PluginServer {
	this := PluginServer{}
	return &this
}

// NewPluginServerWithDefaults instantiates a new PluginServer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPluginServerWithDefaults() *PluginServer {
	this := PluginServer{}
	return &this
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *PluginServer) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PluginServer) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *PluginServer) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *PluginServer) SetParam(v string) {
	o.Param = &v
}

func (o PluginServer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PluginServer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return toSerialize, nil
}

type NullablePluginServer struct {
	value *PluginServer
	isSet bool
}

func (v NullablePluginServer) Get() *PluginServer {
	return v.value
}

func (v *NullablePluginServer) Set(val *PluginServer) {
	v.value = val
	v.isSet = true
}

func (v NullablePluginServer) IsSet() bool {
	return v.isSet
}

func (v *NullablePluginServer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePluginServer(val *PluginServer) *NullablePluginServer {
	return &NullablePluginServer{value: val, isSet: true}
}

func (v NullablePluginServer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePluginServer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
