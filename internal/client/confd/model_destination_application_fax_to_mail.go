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

// checks if the DestinationApplicationFaxToMail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationApplicationFaxToMail{}

// DestinationApplicationFaxToMail struct for DestinationApplicationFaxToMail
type DestinationApplicationFaxToMail struct {
	// MUST be 'application'
	Type string `json:"type"`
	// MUST be 'fax_to_mail'
	Application *string `json:"application,omitempty"`
	// The email of the application
	Email *string `json:"email,omitempty"`
}

type _DestinationApplicationFaxToMail DestinationApplicationFaxToMail

// NewDestinationApplicationFaxToMail instantiates a new DestinationApplicationFaxToMail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationApplicationFaxToMail(type_ string) *DestinationApplicationFaxToMail {
	this := DestinationApplicationFaxToMail{}
	this.Type = type_
	return &this
}

// NewDestinationApplicationFaxToMailWithDefaults instantiates a new DestinationApplicationFaxToMail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationApplicationFaxToMailWithDefaults() *DestinationApplicationFaxToMail {
	this := DestinationApplicationFaxToMail{}
	return &this
}

// GetType returns the Type field value
func (o *DestinationApplicationFaxToMail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationApplicationFaxToMail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationApplicationFaxToMail) SetType(v string) {
	o.Type = v
}

// GetApplication returns the Application field value if set, zero value otherwise.
func (o *DestinationApplicationFaxToMail) GetApplication() string {
	if o == nil || IsNil(o.Application) {
		var ret string
		return ret
	}
	return *o.Application
}

// GetApplicationOk returns a tuple with the Application field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationFaxToMail) GetApplicationOk() (*string, bool) {
	if o == nil || IsNil(o.Application) {
		return nil, false
	}
	return o.Application, true
}

// HasApplication returns a boolean if a field has been set.
func (o *DestinationApplicationFaxToMail) HasApplication() bool {
	if o != nil && !IsNil(o.Application) {
		return true
	}

	return false
}

// SetApplication gets a reference to the given string and assigns it to the Application field.
func (o *DestinationApplicationFaxToMail) SetApplication(v string) {
	o.Application = &v
}

// GetEmail returns the Email field value if set, zero value otherwise.
func (o *DestinationApplicationFaxToMail) GetEmail() string {
	if o == nil || IsNil(o.Email) {
		var ret string
		return ret
	}
	return *o.Email
}

// GetEmailOk returns a tuple with the Email field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationFaxToMail) GetEmailOk() (*string, bool) {
	if o == nil || IsNil(o.Email) {
		return nil, false
	}
	return o.Email, true
}

// HasEmail returns a boolean if a field has been set.
func (o *DestinationApplicationFaxToMail) HasEmail() bool {
	if o != nil && !IsNil(o.Email) {
		return true
	}

	return false
}

// SetEmail gets a reference to the given string and assigns it to the Email field.
func (o *DestinationApplicationFaxToMail) SetEmail(v string) {
	o.Email = &v
}

func (o DestinationApplicationFaxToMail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationApplicationFaxToMail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	if !IsNil(o.Application) {
		toSerialize["application"] = o.Application
	}
	if !IsNil(o.Email) {
		toSerialize["email"] = o.Email
	}
	return toSerialize, nil
}

func (o *DestinationApplicationFaxToMail) UnmarshalJSON(data []byte) (err error) {
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

	varDestinationApplicationFaxToMail := _DestinationApplicationFaxToMail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDestinationApplicationFaxToMail)

	if err != nil {
		return err
	}

	*o = DestinationApplicationFaxToMail(varDestinationApplicationFaxToMail)

	return err
}

type NullableDestinationApplicationFaxToMail struct {
	value *DestinationApplicationFaxToMail
	isSet bool
}

func (v NullableDestinationApplicationFaxToMail) Get() *DestinationApplicationFaxToMail {
	return v.value
}

func (v *NullableDestinationApplicationFaxToMail) Set(val *DestinationApplicationFaxToMail) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationApplicationFaxToMail) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationApplicationFaxToMail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationApplicationFaxToMail(val *DestinationApplicationFaxToMail) *NullableDestinationApplicationFaxToMail {
	return &NullableDestinationApplicationFaxToMail{value: val, isSet: true}
}

func (v NullableDestinationApplicationFaxToMail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationApplicationFaxToMail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
