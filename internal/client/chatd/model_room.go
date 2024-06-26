/*
accent-chatd

Control your message and presence from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chatd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the Room type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Room{}

// Room struct for Room
type Room struct {
	// The UUID of the room
	Uuid *string `json:"uuid,omitempty"`
	// The name of the room
	Name  *string    `json:"name,omitempty"`
	Users []RoomUser `json:"users"`
}

type _Room Room

// NewRoom instantiates a new Room object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRoom(users []RoomUser) *Room {
	this := Room{}
	this.Users = users
	return &this
}

// NewRoomWithDefaults instantiates a new Room object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRoomWithDefaults() *Room {
	this := Room{}
	return &this
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Room) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Room) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Room) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Room) SetUuid(v string) {
	o.Uuid = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Room) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Room) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Room) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Room) SetName(v string) {
	o.Name = &v
}

// GetUsers returns the Users field value
func (o *Room) GetUsers() []RoomUser {
	if o == nil {
		var ret []RoomUser
		return ret
	}

	return o.Users
}

// GetUsersOk returns a tuple with the Users field value
// and a boolean to check if the value has been set.
func (o *Room) GetUsersOk() ([]RoomUser, bool) {
	if o == nil {
		return nil, false
	}
	return o.Users, true
}

// SetUsers sets field value
func (o *Room) SetUsers(v []RoomUser) {
	o.Users = v
}

func (o Room) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Room) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	toSerialize["users"] = o.Users
	return toSerialize, nil
}

func (o *Room) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"users",
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

	varRoom := _Room{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varRoom)

	if err != nil {
		return err
	}

	*o = Room(varRoom)

	return err
}

type NullableRoom struct {
	value *Room
	isSet bool
}

func (v NullableRoom) Get() *Room {
	return v.value
}

func (v *NullableRoom) Set(val *Room) {
	v.value = val
	v.isSet = true
}

func (v NullableRoom) IsSet() bool {
	return v.isSet
}

func (v *NullableRoom) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRoom(val *Room) *NullableRoom {
	return &NullableRoom{value: val, isSet: true}
}

func (v NullableRoom) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRoom) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
