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

// checks if the EndpointTrunks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointTrunks{}

// EndpointTrunks struct for EndpointTrunks
type EndpointTrunks struct {
	// The number of trunk endpoint matching the searched terms
	Filtered *int32          `json:"filtered,omitempty"`
	Items    []EndpointTrunk `json:"items,omitempty"`
	// The number of trunk endpoint
	Total *int32 `json:"total,omitempty"`
}

// NewEndpointTrunks instantiates a new EndpointTrunks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointTrunks() *EndpointTrunks {
	this := EndpointTrunks{}
	return &this
}

// NewEndpointTrunksWithDefaults instantiates a new EndpointTrunks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointTrunksWithDefaults() *EndpointTrunks {
	this := EndpointTrunks{}
	return &this
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *EndpointTrunks) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTrunks) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *EndpointTrunks) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *EndpointTrunks) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *EndpointTrunks) GetItems() []EndpointTrunk {
	if o == nil || IsNil(o.Items) {
		var ret []EndpointTrunk
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTrunks) GetItemsOk() ([]EndpointTrunk, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *EndpointTrunks) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []EndpointTrunk and assigns it to the Items field.
func (o *EndpointTrunks) SetItems(v []EndpointTrunk) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *EndpointTrunks) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointTrunks) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *EndpointTrunks) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *EndpointTrunks) SetTotal(v int32) {
	o.Total = &v
}

func (o EndpointTrunks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointTrunks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Filtered) {
		toSerialize["filtered"] = o.Filtered
	}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableEndpointTrunks struct {
	value *EndpointTrunks
	isSet bool
}

func (v NullableEndpointTrunks) Get() *EndpointTrunks {
	return v.value
}

func (v *NullableEndpointTrunks) Set(val *EndpointTrunks) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointTrunks) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointTrunks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointTrunks(val *EndpointTrunks) *NullableEndpointTrunks {
	return &NullableEndpointTrunks{value: val, isSet: true}
}

func (v NullableEndpointTrunks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointTrunks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
