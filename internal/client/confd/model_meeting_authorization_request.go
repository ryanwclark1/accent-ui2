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

// checks if the MeetingAuthorizationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &MeetingAuthorizationRequest{}

// MeetingAuthorizationRequest struct for MeetingAuthorizationRequest
type MeetingAuthorizationRequest struct {
	GuestName *string `json:"guest_name,omitempty"`
}

// NewMeetingAuthorizationRequest instantiates a new MeetingAuthorizationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMeetingAuthorizationRequest() *MeetingAuthorizationRequest {
	this := MeetingAuthorizationRequest{}
	return &this
}

// NewMeetingAuthorizationRequestWithDefaults instantiates a new MeetingAuthorizationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMeetingAuthorizationRequestWithDefaults() *MeetingAuthorizationRequest {
	this := MeetingAuthorizationRequest{}
	return &this
}

// GetGuestName returns the GuestName field value if set, zero value otherwise.
func (o *MeetingAuthorizationRequest) GetGuestName() string {
	if o == nil || IsNil(o.GuestName) {
		var ret string
		return ret
	}
	return *o.GuestName
}

// GetGuestNameOk returns a tuple with the GuestName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *MeetingAuthorizationRequest) GetGuestNameOk() (*string, bool) {
	if o == nil || IsNil(o.GuestName) {
		return nil, false
	}
	return o.GuestName, true
}

// HasGuestName returns a boolean if a field has been set.
func (o *MeetingAuthorizationRequest) HasGuestName() bool {
	if o != nil && !IsNil(o.GuestName) {
		return true
	}

	return false
}

// SetGuestName gets a reference to the given string and assigns it to the GuestName field.
func (o *MeetingAuthorizationRequest) SetGuestName(v string) {
	o.GuestName = &v
}

func (o MeetingAuthorizationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o MeetingAuthorizationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.GuestName) {
		toSerialize["guest_name"] = o.GuestName
	}
	return toSerialize, nil
}

type NullableMeetingAuthorizationRequest struct {
	value *MeetingAuthorizationRequest
	isSet bool
}

func (v NullableMeetingAuthorizationRequest) Get() *MeetingAuthorizationRequest {
	return v.value
}

func (v *NullableMeetingAuthorizationRequest) Set(val *MeetingAuthorizationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableMeetingAuthorizationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableMeetingAuthorizationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMeetingAuthorizationRequest(val *MeetingAuthorizationRequest) *NullableMeetingAuthorizationRequest {
	return &NullableMeetingAuthorizationRequest{value: val, isSet: true}
}

func (v NullableMeetingAuthorizationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMeetingAuthorizationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
