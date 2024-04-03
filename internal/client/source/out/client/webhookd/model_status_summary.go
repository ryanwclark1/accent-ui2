/*
accent-webhookd

Control your webhooks from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhookd

import (
	"encoding/json"
)

// checks if the StatusSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusSummary{}

// StatusSummary struct for StatusSummary
type StatusSummary struct {
	BusConsumer  *ComponentWithStatus `json:"bus_consumer,omitempty"`
	MasterTenant *ComponentWithStatus `json:"master_tenant,omitempty"`
}

// NewStatusSummary instantiates a new StatusSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewStatusSummary() *StatusSummary {
	this := StatusSummary{}
	return &this
}

// NewStatusSummaryWithDefaults instantiates a new StatusSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewStatusSummaryWithDefaults() *StatusSummary {
	this := StatusSummary{}
	return &this
}

// GetBusConsumer returns the BusConsumer field value if set, zero value otherwise.
func (o *StatusSummary) GetBusConsumer() ComponentWithStatus {
	if o == nil || IsNil(o.BusConsumer) {
		var ret ComponentWithStatus
		return ret
	}
	return *o.BusConsumer
}

// GetBusConsumerOk returns a tuple with the BusConsumer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSummary) GetBusConsumerOk() (*ComponentWithStatus, bool) {
	if o == nil || IsNil(o.BusConsumer) {
		return nil, false
	}
	return o.BusConsumer, true
}

// HasBusConsumer returns a boolean if a field has been set.
func (o *StatusSummary) HasBusConsumer() bool {
	if o != nil && !IsNil(o.BusConsumer) {
		return true
	}

	return false
}

// SetBusConsumer gets a reference to the given ComponentWithStatus and assigns it to the BusConsumer field.
func (o *StatusSummary) SetBusConsumer(v ComponentWithStatus) {
	o.BusConsumer = &v
}

// GetMasterTenant returns the MasterTenant field value if set, zero value otherwise.
func (o *StatusSummary) GetMasterTenant() ComponentWithStatus {
	if o == nil || IsNil(o.MasterTenant) {
		var ret ComponentWithStatus
		return ret
	}
	return *o.MasterTenant
}

// GetMasterTenantOk returns a tuple with the MasterTenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSummary) GetMasterTenantOk() (*ComponentWithStatus, bool) {
	if o == nil || IsNil(o.MasterTenant) {
		return nil, false
	}
	return o.MasterTenant, true
}

// HasMasterTenant returns a boolean if a field has been set.
func (o *StatusSummary) HasMasterTenant() bool {
	if o != nil && !IsNil(o.MasterTenant) {
		return true
	}

	return false
}

// SetMasterTenant gets a reference to the given ComponentWithStatus and assigns it to the MasterTenant field.
func (o *StatusSummary) SetMasterTenant(v ComponentWithStatus) {
	o.MasterTenant = &v
}

func (o StatusSummary) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o StatusSummary) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusConsumer) {
		toSerialize["bus_consumer"] = o.BusConsumer
	}
	if !IsNil(o.MasterTenant) {
		toSerialize["master_tenant"] = o.MasterTenant
	}
	return toSerialize, nil
}

type NullableStatusSummary struct {
	value *StatusSummary
	isSet bool
}

func (v NullableStatusSummary) Get() *StatusSummary {
	return v.value
}

func (v *NullableStatusSummary) Set(val *StatusSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusSummary(val *StatusSummary) *NullableStatusSummary {
	return &NullableStatusSummary{value: val, isSet: true}
}

func (v NullableStatusSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
