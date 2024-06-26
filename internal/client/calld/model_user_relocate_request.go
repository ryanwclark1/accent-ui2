/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the UserRelocateRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &UserRelocateRequest{}

// UserRelocateRequest struct for UserRelocateRequest
type UserRelocateRequest struct {
	// Inform the destination phone that it should answer automatically. Limitation: this does not work on SCCP phones.
	AutoAnswer  *bool                `json:"auto_answer,omitempty"`
	Completions []RelocateCompletion `json:"completions,omitempty"`
	// What kind of destination the relocated call should be connected
	Destination string `json:"destination"`
	// Call ID of the relocate initiator. This call must be owned by the authenticated user.
	InitiatorCall string                `json:"initiator_call"`
	Location      *UserRelocateLocation `json:"location,omitempty"`
	// Number of seconds to wait for the recipient to answer
	Timeout *int32 `json:"timeout,omitempty"`
}

type _UserRelocateRequest UserRelocateRequest

// NewUserRelocateRequest instantiates a new UserRelocateRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUserRelocateRequest(destination string, initiatorCall string) *UserRelocateRequest {
	this := UserRelocateRequest{}
	this.Destination = destination
	this.InitiatorCall = initiatorCall
	return &this
}

// NewUserRelocateRequestWithDefaults instantiates a new UserRelocateRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserRelocateRequestWithDefaults() *UserRelocateRequest {
	this := UserRelocateRequest{}
	return &this
}

// GetAutoAnswer returns the AutoAnswer field value if set, zero value otherwise.
func (o *UserRelocateRequest) GetAutoAnswer() bool {
	if o == nil || IsNil(o.AutoAnswer) {
		var ret bool
		return ret
	}
	return *o.AutoAnswer
}

// GetAutoAnswerOk returns a tuple with the AutoAnswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetAutoAnswerOk() (*bool, bool) {
	if o == nil || IsNil(o.AutoAnswer) {
		return nil, false
	}
	return o.AutoAnswer, true
}

// HasAutoAnswer returns a boolean if a field has been set.
func (o *UserRelocateRequest) HasAutoAnswer() bool {
	if o != nil && !IsNil(o.AutoAnswer) {
		return true
	}

	return false
}

// SetAutoAnswer gets a reference to the given bool and assigns it to the AutoAnswer field.
func (o *UserRelocateRequest) SetAutoAnswer(v bool) {
	o.AutoAnswer = &v
}

// GetCompletions returns the Completions field value if set, zero value otherwise.
func (o *UserRelocateRequest) GetCompletions() []RelocateCompletion {
	if o == nil || IsNil(o.Completions) {
		var ret []RelocateCompletion
		return ret
	}
	return o.Completions
}

// GetCompletionsOk returns a tuple with the Completions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetCompletionsOk() ([]RelocateCompletion, bool) {
	if o == nil || IsNil(o.Completions) {
		return nil, false
	}
	return o.Completions, true
}

// HasCompletions returns a boolean if a field has been set.
func (o *UserRelocateRequest) HasCompletions() bool {
	if o != nil && !IsNil(o.Completions) {
		return true
	}

	return false
}

// SetCompletions gets a reference to the given []RelocateCompletion and assigns it to the Completions field.
func (o *UserRelocateRequest) SetCompletions(v []RelocateCompletion) {
	o.Completions = v
}

// GetDestination returns the Destination field value
func (o *UserRelocateRequest) GetDestination() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetDestinationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *UserRelocateRequest) SetDestination(v string) {
	o.Destination = v
}

// GetInitiatorCall returns the InitiatorCall field value
func (o *UserRelocateRequest) GetInitiatorCall() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.InitiatorCall
}

// GetInitiatorCallOk returns a tuple with the InitiatorCall field value
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetInitiatorCallOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InitiatorCall, true
}

// SetInitiatorCall sets field value
func (o *UserRelocateRequest) SetInitiatorCall(v string) {
	o.InitiatorCall = v
}

// GetLocation returns the Location field value if set, zero value otherwise.
func (o *UserRelocateRequest) GetLocation() UserRelocateLocation {
	if o == nil || IsNil(o.Location) {
		var ret UserRelocateLocation
		return ret
	}
	return *o.Location
}

// GetLocationOk returns a tuple with the Location field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetLocationOk() (*UserRelocateLocation, bool) {
	if o == nil || IsNil(o.Location) {
		return nil, false
	}
	return o.Location, true
}

// HasLocation returns a boolean if a field has been set.
func (o *UserRelocateRequest) HasLocation() bool {
	if o != nil && !IsNil(o.Location) {
		return true
	}

	return false
}

// SetLocation gets a reference to the given UserRelocateLocation and assigns it to the Location field.
func (o *UserRelocateRequest) SetLocation(v UserRelocateLocation) {
	o.Location = &v
}

// GetTimeout returns the Timeout field value if set, zero value otherwise.
func (o *UserRelocateRequest) GetTimeout() int32 {
	if o == nil || IsNil(o.Timeout) {
		var ret int32
		return ret
	}
	return *o.Timeout
}

// GetTimeoutOk returns a tuple with the Timeout field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *UserRelocateRequest) GetTimeoutOk() (*int32, bool) {
	if o == nil || IsNil(o.Timeout) {
		return nil, false
	}
	return o.Timeout, true
}

// HasTimeout returns a boolean if a field has been set.
func (o *UserRelocateRequest) HasTimeout() bool {
	if o != nil && !IsNil(o.Timeout) {
		return true
	}

	return false
}

// SetTimeout gets a reference to the given int32 and assigns it to the Timeout field.
func (o *UserRelocateRequest) SetTimeout(v int32) {
	o.Timeout = &v
}

func (o UserRelocateRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o UserRelocateRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AutoAnswer) {
		toSerialize["auto_answer"] = o.AutoAnswer
	}
	if !IsNil(o.Completions) {
		toSerialize["completions"] = o.Completions
	}
	toSerialize["destination"] = o.Destination
	toSerialize["initiator_call"] = o.InitiatorCall
	if !IsNil(o.Location) {
		toSerialize["location"] = o.Location
	}
	if !IsNil(o.Timeout) {
		toSerialize["timeout"] = o.Timeout
	}
	return toSerialize, nil
}

func (o *UserRelocateRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"initiator_call",
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

	varUserRelocateRequest := _UserRelocateRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varUserRelocateRequest)

	if err != nil {
		return err
	}

	*o = UserRelocateRequest(varUserRelocateRequest)

	return err
}

type NullableUserRelocateRequest struct {
	value *UserRelocateRequest
	isSet bool
}

func (v NullableUserRelocateRequest) Get() *UserRelocateRequest {
	return v.value
}

func (v *NullableUserRelocateRequest) Set(val *UserRelocateRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableUserRelocateRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableUserRelocateRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUserRelocateRequest(val *UserRelocateRequest) *NullableUserRelocateRequest {
	return &NullableUserRelocateRequest{value: val, isSet: true}
}

func (v NullableUserRelocateRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUserRelocateRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
