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

// checks if the DestinationHangupBusy type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationHangupBusy{}

// DestinationHangupBusy struct for DestinationHangupBusy
type DestinationHangupBusy struct {
	// MUST be 'hangup'
	Type string `json:"type"`
	// MUST be 'busy'
	Cause *string `json:"cause,omitempty"`
	// The timeout of the hangup
	Timeout *int32 `json:"timeout,omitempty"`
}

type _DestinationHangupBusy DestinationHangupBusy

// NewDestinationHangupBusy instantiates a new DestinationHangupBusy object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationHangupBusy(type_ string) *DestinationHangupBusy {
	this := DestinationHangupBusy{}
	this.Type = type_
	return &this
}

// NewDestinationHangupBusyWithDefaults instantiates a new DestinationHangupBusy object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationHangupBusyWithDefaults() *DestinationHangupBusy {
	this := DestinationHangupBusy{}
	return &this
}

// GetType returns the Type field value
func (o *DestinationHangupBusy) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationHangupBusy) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationHangupBusy) SetType(v string) {
	o.Type = v
}

// GetCause returns the Cause field value if set, zero value otherwise.
func (o *DestinationHangupBusy) GetCause() string {
	if o == nil || IsNil(o.Cause) {
		var ret string
		return ret
	}
	return *o.Cause
}

// GetCauseOk returns a tuple with the Cause field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationHangupBusy) GetCauseOk() (*string, bool) {
	if o == nil || IsNil(o.Cause) {
		return nil, false
	}
	return o.Cause, true
}

// HasCause returns a boolean if a field has been set.
func (o *DestinationHangupBusy) HasCause() bool {
	if o != nil && !IsNil(o.Cause) {
		return true
	}

	return false
}

// SetCause gets a reference to the given string and assigns it to the Cause field.
func (o *DestinationHangupBusy) SetCause(v string) {
	o.Cause = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *DestinationHangupBusy) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationHangupBusy) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *DestinationHangupBusy) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *DestinationHangupBusy) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o DestinationHangupBusy) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationHangupBusy) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Cause) {
		toSerialize["cause"] = o.Cause
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

func (o *DestinationHangupBusy) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
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

	varDestinationHangupBusy := _DestinationHangupBusy{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDestinationHangupBusy)

	if err != nil {
		return err
	}

	*o = DestinationHangupBusy(varDestinationHangupBusy)

	return err
}

type NullableDestinationHangupBusy struct {
	value *DestinationHangupBusy
	isSet bool
}

func (v NullableDestinationHangupBusy) Get() *DestinationHangupBusy {
	return v.value
}

func (v *NullableDestinationHangupBusy) Set(val *DestinationHangupBusy) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationHangupBusy) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationHangupBusy) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationHangupBusy(val *DestinationHangupBusy) *NullableDestinationHangupBusy {
	return &NullableDestinationHangupBusy{value: val, isSet: true}
}

func (v NullableDestinationHangupBusy) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationHangupBusy) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
