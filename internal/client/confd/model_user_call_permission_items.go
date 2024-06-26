/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"encoding/json"
)

// checks if the UserCallPermissionItems type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserCallPermissionItems{}

// UserCallPermissionItems struct for UserCallPermissionItems
type UserCallPermissionItems struct {
	Items []UserCallPermission `json:"items,omitempty"`
	Total *int32               `json:"total,omitempty"`
}

// NewUserCallPermissionItems instantiates a new UserCallPermissionItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserCallPermissionItems() *UserCallPermissionItems {
	this := UserCallPermissionItems{}
	return &this
}

// NewUserCallPermissionItemsWithDefaults instantiates a new UserCallPermissionItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserCallPermissionItemsWithDefaults() *UserCallPermissionItems {
	this := UserCallPermissionItems{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *UserCallPermissionItems) GetItems() []UserCallPermission {
	if o == nil || IsNil(o.Items) {
		var ret []UserCallPermission
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCallPermissionItems) GetItemsOk() ([]UserCallPermission, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *UserCallPermissionItems) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []UserCallPermission and assigns it to the Items field.
func (o *UserCallPermissionItems) SetItems(v []UserCallPermission) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *UserCallPermissionItems) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserCallPermissionItems) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *UserCallPermissionItems) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *UserCallPermissionItems) SetTotal(v int32) {
	o.Total = &v
}

func (o UserCallPermissionItems) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserCallPermissionItems) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableUserCallPermissionItems struct {
	value *UserCallPermissionItems
	isSet bool
}

func (v NullableUserCallPermissionItems) Get() *UserCallPermissionItems {
	return v.value
}

func (v *NullableUserCallPermissionItems) Set(val *UserCallPermissionItems) {
	v.value = val
	v.isSet = true
}

func (v NullableUserCallPermissionItems) IsSet() bool {
	return v.isSet
}

func (v *NullableUserCallPermissionItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserCallPermissionItems(val *UserCallPermissionItems) *NullableUserCallPermissionItems {
	return &NullableUserCallPermissionItems{value: val, isSet: true}
}

func (v NullableUserCallPermissionItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserCallPermissionItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
