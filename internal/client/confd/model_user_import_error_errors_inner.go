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

// checks if the UserImportErrorErrorsInner type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserImportErrorErrorsInner{}

// UserImportErrorErrorsInner struct for UserImportErrorErrorsInner
type UserImportErrorErrorsInner struct {
	Details *UserImportErrorErrorsInnerDetails `json:"details,omitempty"`
	// Error message
	Message *string `json:"message,omitempty"`
	// Time at which the error occurred
	Timestamp *int32 `json:"timestamp,omitempty"`
}

// NewUserImportErrorErrorsInner instantiates a new UserImportErrorErrorsInner object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserImportErrorErrorsInner() *UserImportErrorErrorsInner {
	this := UserImportErrorErrorsInner{}
	return &this
}

// NewUserImportErrorErrorsInnerWithDefaults instantiates a new UserImportErrorErrorsInner object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserImportErrorErrorsInnerWithDefaults() *UserImportErrorErrorsInner {
	this := UserImportErrorErrorsInner{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *UserImportErrorErrorsInner) GetDetails() UserImportErrorErrorsInnerDetails {
	if o == nil || IsNil(o.Details) {
		var ret UserImportErrorErrorsInnerDetails
		return ret
	}
	return *o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportErrorErrorsInner) GetDetailsOk() (*UserImportErrorErrorsInnerDetails, bool) {
	if o == nil || IsNil(o.Details) {
		return nil, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *UserImportErrorErrorsInner) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given UserImportErrorErrorsInnerDetails and assigns it to the Details field.
func (o *UserImportErrorErrorsInner) SetDetails(v UserImportErrorErrorsInnerDetails) {
	o.Details = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *UserImportErrorErrorsInner) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportErrorErrorsInner) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *UserImportErrorErrorsInner) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *UserImportErrorErrorsInner) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *UserImportErrorErrorsInner) GetTimestamp() int32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret int32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserImportErrorErrorsInner) GetTimestampOk() (*int32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *UserImportErrorErrorsInner) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given int32 and assigns it to the Timestamp field.
func (o *UserImportErrorErrorsInner) SetTimestamp(v int32) {
	o.Timestamp = &v
}

func (o UserImportErrorErrorsInner) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserImportErrorErrorsInner) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	return toSerialize, nil
}

type NullableUserImportErrorErrorsInner struct {
	value *UserImportErrorErrorsInner
	isSet bool
}

func (v NullableUserImportErrorErrorsInner) Get() *UserImportErrorErrorsInner {
	return v.value
}

func (v *NullableUserImportErrorErrorsInner) Set(val *UserImportErrorErrorsInner) {
	v.value = val
	v.isSet = true
}

func (v NullableUserImportErrorErrorsInner) IsSet() bool {
	return v.isSet
}

func (v *NullableUserImportErrorErrorsInner) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserImportErrorErrorsInner(val *UserImportErrorErrorsInner) *NullableUserImportErrorErrorsInner {
	return &NullableUserImportErrorErrorsInner{value: val, isSet: true}
}

func (v NullableUserImportErrorErrorsInner) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserImportErrorErrorsInner) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
