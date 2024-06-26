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

// checks if the HttpsProxy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &HttpsProxy{}

// HttpsProxy struct for HttpsProxy
type HttpsProxy struct {
	// The proxy for HTTPS requests. Format is `host:port`
	Param *string `json:"param,omitempty"`
}

// NewHttpsProxy instantiates a new HttpsProxy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewHttpsProxy() *HttpsProxy {
	this := HttpsProxy{}
	return &this
}

// NewHttpsProxyWithDefaults instantiates a new HttpsProxy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewHttpsProxyWithDefaults() *HttpsProxy {
	this := HttpsProxy{}
	return &this
}

// GetParam returns the Param field value if set, zero value otherwise.
func (o *HttpsProxy) GetParam() string {
	if o == nil || IsNil(o.Param) {
		var ret string
		return ret
	}
	return *o.Param
}

// GetParamOk returns a tuple with the Param field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *HttpsProxy) GetParamOk() (*string, bool) {
	if o == nil || IsNil(o.Param) {
		return nil, false
	}
	return o.Param, true
}

// HasParam returns a boolean if a field has been set.
func (o *HttpsProxy) HasParam() bool {
	if o != nil && !IsNil(o.Param) {
		return true
	}

	return false
}

// SetParam gets a reference to the given string and assigns it to the Param field.
func (o *HttpsProxy) SetParam(v string) {
	o.Param = &v
}

func (o HttpsProxy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o HttpsProxy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Param) {
		toSerialize["param"] = o.Param
	}
	return toSerialize, nil
}

type NullableHttpsProxy struct {
	value *HttpsProxy
	isSet bool
}

func (v NullableHttpsProxy) Get() *HttpsProxy {
	return v.value
}

func (v *NullableHttpsProxy) Set(val *HttpsProxy) {
	v.value = val
	v.isSet = true
}

func (v NullableHttpsProxy) IsSet() bool {
	return v.isSet
}

func (v *NullableHttpsProxy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableHttpsProxy(val *HttpsProxy) *NullableHttpsProxy {
	return &NullableHttpsProxy{value: val, isSet: true}
}

func (v NullableHttpsProxy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableHttpsProxy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
