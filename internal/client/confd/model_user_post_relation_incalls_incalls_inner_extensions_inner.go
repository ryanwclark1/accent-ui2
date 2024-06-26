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

// checks if the UserPostRelationIncallsIncallsInnerExtensionsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserPostRelationIncallsIncallsInnerExtensionsInner{}

// UserPostRelationIncallsIncallsInnerExtensionsInner struct for UserPostRelationIncallsIncallsInnerExtensionsInner
type UserPostRelationIncallsIncallsInnerExtensionsInner struct {
	Context *string `json:"context,omitempty"`
	Exten   *string `json:"exten,omitempty"`
	// Extension ID
	Id *int32 `json:"id,omitempty"`
}

// NewUserPostRelationIncallsIncallsInnerExtensionsInner instantiates a new UserPostRelationIncallsIncallsInnerExtensionsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserPostRelationIncallsIncallsInnerExtensionsInner() *UserPostRelationIncallsIncallsInnerExtensionsInner {
	this := UserPostRelationIncallsIncallsInnerExtensionsInner{}
	return &this
}

// NewUserPostRelationIncallsIncallsInnerExtensionsInnerWithDefaults instantiates a new UserPostRelationIncallsIncallsInnerExtensionsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserPostRelationIncallsIncallsInnerExtensionsInnerWithDefaults() *UserPostRelationIncallsIncallsInnerExtensionsInner {
	this := UserPostRelationIncallsIncallsInnerExtensionsInner{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) SetContext(v string) {
	o.Context = &v
}

// GetExten returns the Exten field value if set, zero value otherwise.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetExten() string {
	if o == nil || IsNil(o.Exten) {
		var ret string
		return ret
	}
	return *o.Exten
}

// GetExtenOk returns a tuple with the Exten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetExtenOk() (*string, bool) {
	if o == nil || IsNil(o.Exten) {
		return nil, false
	}
	return o.Exten, true
}

// HasExten returns a boolean if a field has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) HasExten() bool {
	if o != nil && !IsNil(o.Exten) {
		return true
	}

	return false
}

// SetExten gets a reference to the given string and assigns it to the Exten field.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) SetExten(v string) {
	o.Exten = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *UserPostRelationIncallsIncallsInnerExtensionsInner) SetId(v int32) {
	o.Id = &v
}

func (o UserPostRelationIncallsIncallsInnerExtensionsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserPostRelationIncallsIncallsInnerExtensionsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Exten) {
		toSerialize["exten"] = o.Exten
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	return toSerialize, nil
}

type NullableUserPostRelationIncallsIncallsInnerExtensionsInner struct {
	value *UserPostRelationIncallsIncallsInnerExtensionsInner
	isSet bool
}

func (v NullableUserPostRelationIncallsIncallsInnerExtensionsInner) Get() *UserPostRelationIncallsIncallsInnerExtensionsInner {
	return v.value
}

func (v *NullableUserPostRelationIncallsIncallsInnerExtensionsInner) Set(val *UserPostRelationIncallsIncallsInnerExtensionsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserPostRelationIncallsIncallsInnerExtensionsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserPostRelationIncallsIncallsInnerExtensionsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserPostRelationIncallsIncallsInnerExtensionsInner(val *UserPostRelationIncallsIncallsInnerExtensionsInner) *NullableUserPostRelationIncallsIncallsInnerExtensionsInner {
	return &NullableUserPostRelationIncallsIncallsInnerExtensionsInner{value: val, isSet: true}
}

func (v NullableUserPostRelationIncallsIncallsInnerExtensionsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserPostRelationIncallsIncallsInnerExtensionsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
