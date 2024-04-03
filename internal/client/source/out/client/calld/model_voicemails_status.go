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

// checks if the VoicemailsStatus type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoicemailsStatus{}

// VoicemailsStatus struct for VoicemailsStatus
type VoicemailsStatus struct {
	Status     *StatusValue `json:"status,omitempty"`
	CacheItems *int32       `json:"cache_items,omitempty"`
}

// NewVoicemailsStatus instantiates a new VoicemailsStatus object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoicemailsStatus() *VoicemailsStatus {
	this := VoicemailsStatus{}
	return &this
}

// NewVoicemailsStatusWithDefaults instantiates a new VoicemailsStatus object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoicemailsStatusWithDefaults() *VoicemailsStatus {
	this := VoicemailsStatus{}
	return &this
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *VoicemailsStatus) GetStatus() StatusValue {
	if o == nil || IsNil(o.Status) {
		var ret StatusValue
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailsStatus) GetStatusOk() (*StatusValue, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *VoicemailsStatus) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given StatusValue and assigns it to the Status field.
func (o *VoicemailsStatus) SetStatus(v StatusValue) {
	o.Status = &v
}

// GetCacheItems returns the CacheItems field value if set, zero value otherwise.
func (o *VoicemailsStatus) GetCacheItems() int32 {
	if o == nil || IsNil(o.CacheItems) {
		var ret int32
		return ret
	}
	return *o.CacheItems
}

// GetCacheItemsOk returns a tuple with the CacheItems field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailsStatus) GetCacheItemsOk() (*int32, bool) {
	if o == nil || IsNil(o.CacheItems) {
		return nil, false
	}
	return o.CacheItems, true
}

// HasCacheItems returns a boolean if a field has been set.
func (o *VoicemailsStatus) HasCacheItems() bool {
	if o != nil && !IsNil(o.CacheItems) {
		return true
	}

	return false
}

// SetCacheItems gets a reference to the given int32 and assigns it to the CacheItems field.
func (o *VoicemailsStatus) SetCacheItems(v int32) {
	o.CacheItems = &v
}

func (o VoicemailsStatus) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoicemailsStatus) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.CacheItems) {
		toSerialize["cache_items"] = o.CacheItems
	}
	return toSerialize, nil
}

type NullableVoicemailsStatus struct {
	value *VoicemailsStatus
	isSet bool
}

func (v NullableVoicemailsStatus) Get() *VoicemailsStatus {
	return v.value
}

func (v *NullableVoicemailsStatus) Set(val *VoicemailsStatus) {
	v.value = val
	v.isSet = true
}

func (v NullableVoicemailsStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableVoicemailsStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoicemailsStatus(val *VoicemailsStatus) *NullableVoicemailsStatus {
	return &NullableVoicemailsStatus{value: val, isSet: true}
}

func (v NullableVoicemailsStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoicemailsStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
