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

// checks if the CallFilterSurrogateUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallFilterSurrogateUser{}

// CallFilterSurrogateUser struct for CallFilterSurrogateUser
type CallFilterSurrogateUser struct {
	// User firstname
	Firstname *string `json:"firstname,omitempty"`
	// User lastname
	Lastname *string `json:"lastname,omitempty"`
	// User UUID. This ID is globally unique across multiple Accent instances
	Uuid     *string `json:"uuid,omitempty"`
	Exten    *string `json:"exten,omitempty"`
	MemberId *int32  `json:"member_id,omitempty"`
}

// NewCallFilterSurrogateUser instantiates a new CallFilterSurrogateUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallFilterSurrogateUser() *CallFilterSurrogateUser {
	this := CallFilterSurrogateUser{}
	return &this
}

// NewCallFilterSurrogateUserWithDefaults instantiates a new CallFilterSurrogateUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallFilterSurrogateUserWithDefaults() *CallFilterSurrogateUser {
	this := CallFilterSurrogateUser{}
	return &this
}

// GetFirstname returns the Firstname field value if set, zero value otherwise.
func (o *CallFilterSurrogateUser) GetFirstname() string {
	if o == nil || IsNil(o.Firstname) {
		var ret string
		return ret
	}
	return *o.Firstname
}

// GetFirstnameOk returns a tuple with the Firstname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUser) GetFirstnameOk() (*string, bool) {
	if o == nil || IsNil(o.Firstname) {
		return nil, false
	}
	return o.Firstname, true
}

// HasFirstname returns a boolean if a field has been set.
func (o *CallFilterSurrogateUser) HasFirstname() bool {
	if o != nil && !IsNil(o.Firstname) {
		return true
	}

	return false
}

// SetFirstname gets a reference to the given string and assigns it to the Firstname field.
func (o *CallFilterSurrogateUser) SetFirstname(v string) {
	o.Firstname = &v
}

// GetLastname returns the Lastname field value if set, zero value otherwise.
func (o *CallFilterSurrogateUser) GetLastname() string {
	if o == nil || IsNil(o.Lastname) {
		var ret string
		return ret
	}
	return *o.Lastname
}

// GetLastnameOk returns a tuple with the Lastname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUser) GetLastnameOk() (*string, bool) {
	if o == nil || IsNil(o.Lastname) {
		return nil, false
	}
	return o.Lastname, true
}

// HasLastname returns a boolean if a field has been set.
func (o *CallFilterSurrogateUser) HasLastname() bool {
	if o != nil && !IsNil(o.Lastname) {
		return true
	}

	return false
}

// SetLastname gets a reference to the given string and assigns it to the Lastname field.
func (o *CallFilterSurrogateUser) SetLastname(v string) {
	o.Lastname = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *CallFilterSurrogateUser) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUser) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *CallFilterSurrogateUser) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *CallFilterSurrogateUser) SetUuid(v string) {
	o.Uuid = &v
}

// GetExten returns the Exten field value if set, zero value otherwise.
func (o *CallFilterSurrogateUser) GetExten() string {
	if o == nil || IsNil(o.Exten) {
		var ret string
		return ret
	}
	return *o.Exten
}

// GetExtenOk returns a tuple with the Exten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUser) GetExtenOk() (*string, bool) {
	if o == nil || IsNil(o.Exten) {
		return nil, false
	}
	return o.Exten, true
}

// HasExten returns a boolean if a field has been set.
func (o *CallFilterSurrogateUser) HasExten() bool {
	if o != nil && !IsNil(o.Exten) {
		return true
	}

	return false
}

// SetExten gets a reference to the given string and assigns it to the Exten field.
func (o *CallFilterSurrogateUser) SetExten(v string) {
	o.Exten = &v
}

// GetMemberId returns the MemberId field value if set, zero value otherwise.
func (o *CallFilterSurrogateUser) GetMemberId() int32 {
	if o == nil || IsNil(o.MemberId) {
		var ret int32
		return ret
	}
	return *o.MemberId
}

// GetMemberIdOk returns a tuple with the MemberId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallFilterSurrogateUser) GetMemberIdOk() (*int32, bool) {
	if o == nil || IsNil(o.MemberId) {
		return nil, false
	}
	return o.MemberId, true
}

// HasMemberId returns a boolean if a field has been set.
func (o *CallFilterSurrogateUser) HasMemberId() bool {
	if o != nil && !IsNil(o.MemberId) {
		return true
	}

	return false
}

// SetMemberId gets a reference to the given int32 and assigns it to the MemberId field.
func (o *CallFilterSurrogateUser) SetMemberId(v int32) {
	o.MemberId = &v
}

func (o CallFilterSurrogateUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallFilterSurrogateUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Firstname) {
		toSerialize["firstname"] = o.Firstname
	}
	if !IsNil(o.Lastname) {
		toSerialize["lastname"] = o.Lastname
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	if !IsNil(o.Exten) {
		toSerialize["exten"] = o.Exten
	}
	if !IsNil(o.MemberId) {
		toSerialize["member_id"] = o.MemberId
	}
	return toSerialize, nil
}

type NullableCallFilterSurrogateUser struct {
	value *CallFilterSurrogateUser
	isSet bool
}

func (v NullableCallFilterSurrogateUser) Get() *CallFilterSurrogateUser {
	return v.value
}

func (v *NullableCallFilterSurrogateUser) Set(val *CallFilterSurrogateUser) {
	v.value = val
	v.isSet = true
}

func (v NullableCallFilterSurrogateUser) IsSet() bool {
	return v.isSet
}

func (v *NullableCallFilterSurrogateUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallFilterSurrogateUser(val *CallFilterSurrogateUser) *NullableCallFilterSurrogateUser {
	return &NullableCallFilterSurrogateUser{value: val, isSet: true}
}

func (v NullableCallFilterSurrogateUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallFilterSurrogateUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
