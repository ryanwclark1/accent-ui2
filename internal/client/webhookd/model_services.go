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

// checks if the Services type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Services{}

// Services struct for Services
type Services struct {
	// The keys are the service names.
	Services map[string]interface{} `json:"services,omitempty"`
}

// NewServices instantiates a new Services object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewServices() *Services {
	this := Services{}
	return &this
}

// NewServicesWithDefaults instantiates a new Services object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewServicesWithDefaults() *Services {
	this := Services{}
	return &this
}

// GetServices returns the Services field value if set, zero value otherwise.
func (o *Services) GetServices() map[string]interface{} {
	if o == nil || IsNil(o.Services) {
		var ret map[string]interface{}
		return ret
	}
	return o.Services
}

// GetServicesOk returns a tuple with the Services field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Services) GetServicesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Services) {
		return map[string]interface{}{}, false
	}
	return o.Services, true
}

// HasServices returns a boolean if a field has been set.
func (o *Services) HasServices() bool {
	if o != nil && !IsNil(o.Services) {
		return true
	}

	return false
}

// SetServices gets a reference to the given map[string]interface{} and assigns it to the Services field.
func (o *Services) SetServices(v map[string]interface{}) {
	o.Services = v
}

func (o Services) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Services) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Services) {
		toSerialize["services"] = o.Services
	}
	return toSerialize, nil
}

type NullableServices struct {
	value *Services
	isSet bool
}

func (v NullableServices) Get() *Services {
	return v.value
}

func (v *NullableServices) Set(val *Services) {
	v.value = val
	v.isSet = true
}

func (v NullableServices) IsSet() bool {
	return v.isSet
}

func (v *NullableServices) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableServices(val *Services) *NullableServices {
	return &NullableServices{value: val, isSet: true}
}

func (v NullableServices) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableServices) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
