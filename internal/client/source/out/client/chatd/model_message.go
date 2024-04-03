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

// checks if the Message type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Message{}

// Message struct for Message
type Message struct {
	// Alias/nickname of the sender
	Alias *string `json:"alias,omitempty"`
	// The content of the message
	Content *string `json:"content,omitempty"`
	// The date of the message's creation
	CreatedAt *string           `json:"created_at,omitempty"`
	Room      *RoomRelationBase `json:"room,omitempty"`
	// tenant uuid of the sender
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// user uuid of the sender
	UserUuid *string `json:"user_uuid,omitempty"`
	// The UUID of the message
	Uuid *string `json:"uuid,omitempty"`
	// accent uuid of the sender
	AccentUuid *string `json:"accent_uuid,omitempty"`
}

// NewMessage instantiates a new Message object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessage() *Message {
	this := Message{}
	return &this
}

// NewMessageWithDefaults instantiates a new Message object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageWithDefaults() *Message {
	this := Message{}
	return &this
}

// GetAlias returns the Alias field value if set, zero value otherwise.
func (o *Message) GetAlias() string {
	if o == nil || IsNil(o.Alias) {
		var ret string
		return ret
	}
	return *o.Alias
}

// GetAliasOk returns a tuple with the Alias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAliasOk() (*string, bool) {
	if o == nil || IsNil(o.Alias) {
		return nil, false
	}
	return o.Alias, true
}

// HasAlias returns a boolean if a field has been set.
func (o *Message) HasAlias() bool {
	if o != nil && !IsNil(o.Alias) {
		return true
	}

	return false
}

// SetAlias gets a reference to the given string and assigns it to the Alias field.
func (o *Message) SetAlias(v string) {
	o.Alias = &v
}

// GetContent returns the Content field value if set, zero value otherwise.
func (o *Message) GetContent() string {
	if o == nil || IsNil(o.Content) {
		var ret string
		return ret
	}
	return *o.Content
}

// GetContentOk returns a tuple with the Content field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetContentOk() (*string, bool) {
	if o == nil || IsNil(o.Content) {
		return nil, false
	}
	return o.Content, true
}

// HasContent returns a boolean if a field has been set.
func (o *Message) HasContent() bool {
	if o != nil && !IsNil(o.Content) {
		return true
	}

	return false
}

// SetContent gets a reference to the given string and assigns it to the Content field.
func (o *Message) SetContent(v string) {
	o.Content = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *Message) GetCreatedAt() string {
	if o == nil || IsNil(o.CreatedAt) {
		var ret string
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetCreatedAtOk() (*string, bool) {
	if o == nil || IsNil(o.CreatedAt) {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *Message) HasCreatedAt() bool {
	if o != nil && !IsNil(o.CreatedAt) {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given string and assigns it to the CreatedAt field.
func (o *Message) SetCreatedAt(v string) {
	o.CreatedAt = &v
}

// GetRoom returns the Room field value if set, zero value otherwise.
func (o *Message) GetRoom() RoomRelationBase {
	if o == nil || IsNil(o.Room) {
		var ret RoomRelationBase
		return ret
	}
	return *o.Room
}

// GetRoomOk returns a tuple with the Room field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetRoomOk() (*RoomRelationBase, bool) {
	if o == nil || IsNil(o.Room) {
		return nil, false
	}
	return o.Room, true
}

// HasRoom returns a boolean if a field has been set.
func (o *Message) HasRoom() bool {
	if o != nil && !IsNil(o.Room) {
		return true
	}

	return false
}

// SetRoom gets a reference to the given RoomRelationBase and assigns it to the Room field.
func (o *Message) SetRoom(v RoomRelationBase) {
	o.Room = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *Message) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *Message) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *Message) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *Message) GetUserUuid() string {
	if o == nil || IsNil(o.UserUuid) {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetUserUuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserUuid) {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *Message) HasUserUuid() bool {
	if o != nil && !IsNil(o.UserUuid) {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *Message) SetUserUuid(v string) {
	o.UserUuid = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *Message) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *Message) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *Message) SetUuid(v string) {
	o.Uuid = &v
}

// GetAccentUuid returns the AccentUuid field value if set, zero value otherwise.
func (o *Message) GetAccentUuid() string {
	if o == nil || IsNil(o.AccentUuid) {
		var ret string
		return ret
	}
	return *o.AccentUuid
}

// GetAccentUuidOk returns a tuple with the AccentUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Message) GetAccentUuidOk() (*string, bool) {
	if o == nil || IsNil(o.AccentUuid) {
		return nil, false
	}
	return o.AccentUuid, true
}

// HasAccentUuid returns a boolean if a field has been set.
func (o *Message) HasAccentUuid() bool {
	if o != nil && !IsNil(o.AccentUuid) {
		return true
	}

	return false
}

// SetAccentUuid gets a reference to the given string and assigns it to the AccentUuid field.
func (o *Message) SetAccentUuid(v string) {
	o.AccentUuid = &v
}

func (o Message) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Message) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Alias) {
		toSerialize["alias"] = o.Alias
	}
	if !IsNil(o.Content) {
		toSerialize["content"] = o.Content
	}
	if !IsNil(o.CreatedAt) {
		toSerialize["created_at"] = o.CreatedAt
	}
	if !IsNil(o.Room) {
		toSerialize["room"] = o.Room
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.UserUuid) {
		toSerialize["user_uuid"] = o.UserUuid
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.AccentUuid) {
		toSerialize["accent_uuid"] = o.AccentUuid
	}
	return toSerialize, nil
}

type NullableMessage struct {
	value *Message
	isSet bool
}

func (v NullableMessage) Get() *Message {
	return v.value
}

func (v *NullableMessage) Set(val *Message) {
	v.value = val
	v.isSet = true
}

func (v NullableMessage) IsSet() bool {
	return v.isSet
}

func (v *NullableMessage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessage(val *Message) *NullableMessage {
	return &NullableMessage{value: val, isSet: true}
}

func (v NullableMessage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
