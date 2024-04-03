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

// checks if the DestinationApplicationVoicemail type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationApplicationVoicemail{}

// DestinationApplicationVoicemail struct for DestinationApplicationVoicemail
type DestinationApplicationVoicemail struct {
	// MUST be 'application'
	Type string `json:"type"`
	// MUST be 'voicemail'
	Application string `json:"application"`
	// The context of the application
	Context int32 `json:"context"`
}

type _DestinationApplicationVoicemail DestinationApplicationVoicemail

// NewDestinationApplicationVoicemail instantiates a new DestinationApplicationVoicemail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationApplicationVoicemail(type_ string, application string, context int32) *DestinationApplicationVoicemail {
	this := DestinationApplicationVoicemail{}
	this.Type = type_
	this.Application = application
	this.Context = context
	return &this
}

// NewDestinationApplicationVoicemailWithDefaults instantiates a new DestinationApplicationVoicemail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationApplicationVoicemailWithDefaults() *DestinationApplicationVoicemail {
	this := DestinationApplicationVoicemail{}
	return &this
}

// GetType returns the Type field value
func (o *DestinationApplicationVoicemail) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationApplicationVoicemail) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationApplicationVoicemail) SetType(v string) {
	o.Type = v
}

// GetApplication returns the Application field value
func (o *DestinationApplicationVoicemail) GetApplication() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Application
}

// GetApplicationOk returns a tuple with the Application field value
// and a boolean to check if the value has been set.
func (o *DestinationApplicationVoicemail) GetApplicationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Application, true
}

// SetApplication sets field value
func (o *DestinationApplicationVoicemail) SetApplication(v string) {
	o.Application = v
}

// GetContext returns the Context field value
func (o *DestinationApplicationVoicemail) GetContext() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *DestinationApplicationVoicemail) GetContextOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *DestinationApplicationVoicemail) SetContext(v int32) {
	o.Context = v
}

func (o DestinationApplicationVoicemail) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationApplicationVoicemail) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["type"] = o.Type
	toSerialize["application"] = o.Application
	toSerialize["context"] = o.Context
	return toSerialize, nil
}

func (o *DestinationApplicationVoicemail) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"type",
		"application",
		"context",
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

	varDestinationApplicationVoicemail := _DestinationApplicationVoicemail{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDestinationApplicationVoicemail)

	if err != nil {
		return err
	}

	*o = DestinationApplicationVoicemail(varDestinationApplicationVoicemail)

	return err
}

type NullableDestinationApplicationVoicemail struct {
	value *DestinationApplicationVoicemail
	isSet bool
}

func (v NullableDestinationApplicationVoicemail) Get() *DestinationApplicationVoicemail {
	return v.value
}

func (v *NullableDestinationApplicationVoicemail) Set(val *DestinationApplicationVoicemail) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationApplicationVoicemail) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationApplicationVoicemail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationApplicationVoicemail(val *DestinationApplicationVoicemail) *NullableDestinationApplicationVoicemail {
	return &NullableDestinationApplicationVoicemail{value: val, isSet: true}
}

func (v NullableDestinationApplicationVoicemail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationApplicationVoicemail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
