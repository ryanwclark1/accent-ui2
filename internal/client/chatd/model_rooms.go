/*
accent-chatd

Control your message and presence from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chatd

import (
	"encoding/json"
)

// checks if the Rooms type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Rooms{}

// Rooms struct for Rooms
type Rooms struct {
	Filtered *int32 `json:"filtered,omitempty"`
	Items    []Room `json:"items,omitempty"`
	Total    *int32 `json:"total,omitempty"`
}

// NewRooms instantiates a new Rooms object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRooms() *Rooms {
	this := Rooms{}
	return &this
}

// NewRoomsWithDefaults instantiates a new Rooms object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomsWithDefaults() *Rooms {
	this := Rooms{}
	return &this
}

// GetFiltered returns the Filtered field value if set, zero value otherwise.
func (o *Rooms) GetFiltered() int32 {
	if o == nil || IsNil(o.Filtered) {
		var ret int32
		return ret
	}
	return *o.Filtered
}

// GetFilteredOk returns a tuple with the Filtered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rooms) GetFilteredOk() (*int32, bool) {
	if o == nil || IsNil(o.Filtered) {
		return nil, false
	}
	return o.Filtered, true
}

// HasFiltered returns a boolean if a field has been set.
func (o *Rooms) HasFiltered() bool {
	if o != nil && !IsNil(o.Filtered) {
		return true
	}

	return false
}

// SetFiltered gets a reference to the given int32 and assigns it to the Filtered field.
func (o *Rooms) SetFiltered(v int32) {
	o.Filtered = &v
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *Rooms) GetItems() []Room {
	if o == nil || IsNil(o.Items) {
		var ret []Room
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rooms) GetItemsOk() ([]Room, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *Rooms) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Room and assigns it to the Items field.
func (o *Rooms) SetItems(v []Room) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *Rooms) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Rooms) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *Rooms) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *Rooms) SetTotal(v int32) {
	o.Total = &v
}

func (o Rooms) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Rooms) ToMap() (map[string]interface{}, error) {
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

type NullableRooms struct {
	value *Rooms
	isSet bool
}

func (v NullableRooms) Get() *Rooms {
	return v.value
}

func (v *NullableRooms) Set(val *Rooms) {
	v.value = val
	v.isSet = true
}

func (v NullableRooms) IsSet() bool {
	return v.isSet
}

func (v *NullableRooms) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRooms(val *Rooms) *NullableRooms {
	return &NullableRooms{value: val, isSet: true}
}

func (v NullableRooms) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRooms) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
