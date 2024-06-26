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

// checks if the StatusSummary type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &StatusSummary{}

// StatusSummary struct for StatusSummary
type StatusSummary struct {
	Ari          *ComponentWithStatus `json:"ari,omitempty"`
	BusConsumer  *ComponentWithStatus `json:"bus_consumer,omitempty"`
	Plugins      *PluginsStatus       `json:"plugins,omitempty"`
	ServiceToken *ComponentWithStatus `json:"service_token,omitempty"`
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

// GetAri returns the Ari field value if set, zero value otherwise.
func (o *StatusSummary) GetAri() ComponentWithStatus {
	if o == nil || IsNil(o.Ari) {
		var ret ComponentWithStatus
		return ret
	}
	return *o.Ari
}

// GetAriOk returns a tuple with the Ari field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSummary) GetAriOk() (*ComponentWithStatus, bool) {
	if o == nil || IsNil(o.Ari) {
		return nil, false
	}
	return o.Ari, true
}

// HasAri returns a boolean if a field has been set.
func (o *StatusSummary) HasAri() bool {
	if o != nil && !IsNil(o.Ari) {
		return true
	}

	return false
}

// SetAri gets a reference to the given ComponentWithStatus and assigns it to the Ari field.
func (o *StatusSummary) SetAri(v ComponentWithStatus) {
	o.Ari = &v
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

// GetPlugins returns the Plugins field value if set, zero value otherwise.
func (o *StatusSummary) GetPlugins() PluginsStatus {
	if o == nil || IsNil(o.Plugins) {
		var ret PluginsStatus
		return ret
	}
	return *o.Plugins
}

// GetPluginsOk returns a tuple with the Plugins field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSummary) GetPluginsOk() (*PluginsStatus, bool) {
	if o == nil || IsNil(o.Plugins) {
		return nil, false
	}
	return o.Plugins, true
}

// HasPlugins returns a boolean if a field has been set.
func (o *StatusSummary) HasPlugins() bool {
	if o != nil && !IsNil(o.Plugins) {
		return true
	}

	return false
}

// SetPlugins gets a reference to the given PluginsStatus and assigns it to the Plugins field.
func (o *StatusSummary) SetPlugins(v PluginsStatus) {
	o.Plugins = &v
}

// GetServiceToken returns the ServiceToken field value if set, zero value otherwise.
func (o *StatusSummary) GetServiceToken() ComponentWithStatus {
	if o == nil || IsNil(o.ServiceToken) {
		var ret ComponentWithStatus
		return ret
	}
	return *o.ServiceToken
}

// GetServiceTokenOk returns a tuple with the ServiceToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *StatusSummary) GetServiceTokenOk() (*ComponentWithStatus, bool) {
	if o == nil || IsNil(o.ServiceToken) {
		return nil, false
	}
	return o.ServiceToken, true
}

// HasServiceToken returns a boolean if a field has been set.
func (o *StatusSummary) HasServiceToken() bool {
	if o != nil && !IsNil(o.ServiceToken) {
		return true
	}

	return false
}

// SetServiceToken gets a reference to the given ComponentWithStatus and assigns it to the ServiceToken field.
func (o *StatusSummary) SetServiceToken(v ComponentWithStatus) {
	o.ServiceToken = &v
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
	if !IsNil(o.Ari) {
		toSerialize["ari"] = o.Ari
	}
	if !IsNil(o.BusConsumer) {
		toSerialize["bus_consumer"] = o.BusConsumer
	}
	if !IsNil(o.Plugins) {
		toSerialize["plugins"] = o.Plugins
	}
	if !IsNil(o.ServiceToken) {
		toSerialize["service_token"] = o.ServiceToken
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
