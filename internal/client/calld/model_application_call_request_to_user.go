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

// checks if the ApplicationCallRequestToUser type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ApplicationCallRequestToUser{}

// ApplicationCallRequestToUser struct for ApplicationCallRequestToUser
type ApplicationCallRequestToUser struct {
	Autoanswer              *bool   `json:"autoanswer,omitempty"`
	DisplayedCallerIdName   *string `json:"displayed_caller_id_name,omitempty"`
	DisplayedCallerIdNumber *string `json:"displayed_caller_id_number,omitempty"`
	UserUuid                string  `json:"user_uuid"`
	// channel variables that should be assigned on this new channel
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type _ApplicationCallRequestToUser ApplicationCallRequestToUser

// NewApplicationCallRequestToUser instantiates a new ApplicationCallRequestToUser object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewApplicationCallRequestToUser(userUuid string) *ApplicationCallRequestToUser {
	this := ApplicationCallRequestToUser{}
	var autoanswer bool = false
	this.Autoanswer = &autoanswer
	this.UserUuid = userUuid
	return &this
}

// NewApplicationCallRequestToUserWithDefaults instantiates a new ApplicationCallRequestToUser object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewApplicationCallRequestToUserWithDefaults() *ApplicationCallRequestToUser {
	this := ApplicationCallRequestToUser{}
	var autoanswer bool = false
	this.Autoanswer = &autoanswer
	return &this
}

// GetAutoanswer returns the Autoanswer field value if set, zero value otherwise.
func (o *ApplicationCallRequestToUser) GetAutoanswer() bool {
	if o == nil || IsNil(o.Autoanswer) {
		var ret bool
		return ret
	}
	return *o.Autoanswer
}

// GetAutoanswerOk returns a tuple with the Autoanswer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCallRequestToUser) GetAutoanswerOk() (*bool, bool) {
	if o == nil || IsNil(o.Autoanswer) {
		return nil, false
	}
	return o.Autoanswer, true
}

// HasAutoanswer returns a boolean if a field has been set.
func (o *ApplicationCallRequestToUser) HasAutoanswer() bool {
	if o != nil && !IsNil(o.Autoanswer) {
		return true
	}

	return false
}

// SetAutoanswer gets a reference to the given bool and assigns it to the Autoanswer field.
func (o *ApplicationCallRequestToUser) SetAutoanswer(v bool) {
	o.Autoanswer = &v
}

// GetDisplayedCallerIdName returns the DisplayedCallerIdName field value if set, zero value otherwise.
func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdName() string {
	if o == nil || IsNil(o.DisplayedCallerIdName) {
		var ret string
		return ret
	}
	return *o.DisplayedCallerIdName
}

// GetDisplayedCallerIdNameOk returns a tuple with the DisplayedCallerIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayedCallerIdName) {
		return nil, false
	}
	return o.DisplayedCallerIdName, true
}

// HasDisplayedCallerIdName returns a boolean if a field has been set.
func (o *ApplicationCallRequestToUser) HasDisplayedCallerIdName() bool {
	if o != nil && !IsNil(o.DisplayedCallerIdName) {
		return true
	}

	return false
}

// SetDisplayedCallerIdName gets a reference to the given string and assigns it to the DisplayedCallerIdName field.
func (o *ApplicationCallRequestToUser) SetDisplayedCallerIdName(v string) {
	o.DisplayedCallerIdName = &v
}

// GetDisplayedCallerIdNumber returns the DisplayedCallerIdNumber field value if set, zero value otherwise.
func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNumber() string {
	if o == nil || IsNil(o.DisplayedCallerIdNumber) {
		var ret string
		return ret
	}
	return *o.DisplayedCallerIdNumber
}

// GetDisplayedCallerIdNumberOk returns a tuple with the DisplayedCallerIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.DisplayedCallerIdNumber) {
		return nil, false
	}
	return o.DisplayedCallerIdNumber, true
}

// HasDisplayedCallerIdNumber returns a boolean if a field has been set.
func (o *ApplicationCallRequestToUser) HasDisplayedCallerIdNumber() bool {
	if o != nil && !IsNil(o.DisplayedCallerIdNumber) {
		return true
	}

	return false
}

// SetDisplayedCallerIdNumber gets a reference to the given string and assigns it to the DisplayedCallerIdNumber field.
func (o *ApplicationCallRequestToUser) SetDisplayedCallerIdNumber(v string) {
	o.DisplayedCallerIdNumber = &v
}

// GetUserUuid returns the UserUuid field value
func (o *ApplicationCallRequestToUser) GetUserUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value
// and a boolean to check if the value has been set.
func (o *ApplicationCallRequestToUser) GetUserUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.UserUuid, true
}

// SetUserUuid sets field value
func (o *ApplicationCallRequestToUser) SetUserUuid(v string) {
	o.UserUuid = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *ApplicationCallRequestToUser) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ApplicationCallRequestToUser) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *ApplicationCallRequestToUser) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *ApplicationCallRequestToUser) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

func (o ApplicationCallRequestToUser) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ApplicationCallRequestToUser) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Autoanswer) {
		toSerialize["autoanswer"] = o.Autoanswer
	}
	if !IsNil(o.DisplayedCallerIdName) {
		toSerialize["displayed_caller_id_name"] = o.DisplayedCallerIdName
	}
	if !IsNil(o.DisplayedCallerIdNumber) {
		toSerialize["displayed_caller_id_number"] = o.DisplayedCallerIdNumber
	}
	toSerialize["user_uuid"] = o.UserUuid
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *ApplicationCallRequestToUser) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"user_uuid",
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

	varApplicationCallRequestToUser := _ApplicationCallRequestToUser{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varApplicationCallRequestToUser)

	if err != nil {
		return err
	}

	*o = ApplicationCallRequestToUser(varApplicationCallRequestToUser)

	return err
}

type NullableApplicationCallRequestToUser struct {
	value *ApplicationCallRequestToUser
	isSet bool
}

func (v NullableApplicationCallRequestToUser) Get() *ApplicationCallRequestToUser {
	return v.value
}

func (v *NullableApplicationCallRequestToUser) Set(val *ApplicationCallRequestToUser) {
	v.value = val
	v.isSet = true
}

func (v NullableApplicationCallRequestToUser) IsSet() bool {
	return v.isSet
}

func (v *NullableApplicationCallRequestToUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableApplicationCallRequestToUser(val *ApplicationCallRequestToUser) *NullableApplicationCallRequestToUser {
	return &NullableApplicationCallRequestToUser{value: val, isSet: true}
}

func (v NullableApplicationCallRequestToUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableApplicationCallRequestToUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
