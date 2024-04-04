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

// checks if the DHCPInfoObject type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DHCPInfoObject{}

// DHCPInfoObject struct for DHCPInfoObject
type DHCPInfoObject struct {
	DhcpInfo *DHCPInfo `json:"dhcp_info,omitempty"`
}

// NewDHCPInfoObject instantiates a new DHCPInfoObject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDHCPInfoObject() *DHCPInfoObject {
	this := DHCPInfoObject{}
	return &this
}

// NewDHCPInfoObjectWithDefaults instantiates a new DHCPInfoObject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDHCPInfoObjectWithDefaults() *DHCPInfoObject {
	this := DHCPInfoObject{}
	return &this
}

// GetDhcpInfo returns the DhcpInfo field value if set, zero value otherwise.
func (o *DHCPInfoObject) GetDhcpInfo() DHCPInfo {
	if o == nil || IsNil(o.DhcpInfo) {
		var ret DHCPInfo
		return ret
	}
	return *o.DhcpInfo
}

// GetDhcpInfoOk returns a tuple with the DhcpInfo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DHCPInfoObject) GetDhcpInfoOk() (*DHCPInfo, bool) {
	if o == nil || IsNil(o.DhcpInfo) {
		return nil, false
	}
	return o.DhcpInfo, true
}

// HasDhcpInfo returns a boolean if a field has been set.
func (o *DHCPInfoObject) HasDhcpInfo() bool {
	if o != nil && !IsNil(o.DhcpInfo) {
		return true
	}

	return false
}

// SetDhcpInfo gets a reference to the given DHCPInfo and assigns it to the DhcpInfo field.
func (o *DHCPInfoObject) SetDhcpInfo(v DHCPInfo) {
	o.DhcpInfo = &v
}

func (o DHCPInfoObject) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DHCPInfoObject) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.DhcpInfo) {
		toSerialize["dhcp_info"] = o.DhcpInfo
	}
	return toSerialize, nil
}

type NullableDHCPInfoObject struct {
	value *DHCPInfoObject
	isSet bool
}

func (v NullableDHCPInfoObject) Get() *DHCPInfoObject {
	return v.value
}

func (v *NullableDHCPInfoObject) Set(val *DHCPInfoObject) {
	v.value = val
	v.isSet = true
}

func (v NullableDHCPInfoObject) IsSet() bool {
	return v.isSet
}

func (v *NullableDHCPInfoObject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDHCPInfoObject(val *DHCPInfoObject) *NullableDHCPInfoObject {
	return &NullableDHCPInfoObject{value: val, isSet: true}
}

func (v NullableDHCPInfoObject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDHCPInfoObject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
