/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the DeviceItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DeviceItems{}

// DeviceItems struct for DeviceItems
type DeviceItems struct {
	Items []Device `json:"items,omitempty"`
	Total int32    `json:"total"`
}

type _DeviceItems DeviceItems

// NewDeviceItems instantiates a new DeviceItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDeviceItems(total int32) *DeviceItems {
	this := DeviceItems{}
	this.Total = total
	return &this
}

// NewDeviceItemsWithDefaults instantiates a new DeviceItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDeviceItemsWithDefaults() *DeviceItems {
	this := DeviceItems{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *DeviceItems) GetItems() []Device {
	if o == nil || IsNil(o.Items) {
		var ret []Device
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DeviceItems) GetItemsOk() ([]Device, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *DeviceItems) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Device and assigns it to the Items field.
func (o *DeviceItems) SetItems(v []Device) {
	o.Items = v
}

// GetTotal returns the Total field value
func (o *DeviceItems) GetTotal() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Total
}

// GetTotalOk returns a tuple with the Total field value
// and a boolean to check if the value has been set.
func (o *DeviceItems) GetTotalOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Total, true
}

// SetTotal sets field value
func (o *DeviceItems) SetTotal(v int32) {
	o.Total = v
}

func (o DeviceItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DeviceItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	toSerialize["total"] = o.Total
	return toSerialize, nil
}

func (o *DeviceItems) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"total",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varDeviceItems := _DeviceItems{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDeviceItems)

	if err != nil {
		return err
	}

	*o = DeviceItems(varDeviceItems)

	return err
}

type NullableDeviceItems struct {
	value *DeviceItems
	isSet bool
}

func (v NullableDeviceItems) Get() *DeviceItems {
	return v.value
}

func (v *NullableDeviceItems) Set(val *DeviceItems) {
	v.value = val
	v.isSet = true
}

func (v NullableDeviceItems) IsSet() bool {
	return v.isSet
}

func (v *NullableDeviceItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDeviceItems(val *DeviceItems) *NullableDeviceItems {
	return &NullableDeviceItems{value: val, isSet: true}
}

func (v NullableDeviceItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDeviceItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
