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

// checks if the DestinationExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationExtension{}

// DestinationExtension struct for DestinationExtension
type DestinationExtension struct {
	// Context of the extension
	Context string `json:"context"`
	Exten   string `json:"exten"`
	// MUST be 'extension'
	Type string `json:"type"`
}

type _DestinationExtension DestinationExtension

// NewDestinationExtension instantiates a new DestinationExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationExtension(context string, exten string, type_ string) *DestinationExtension {
	this := DestinationExtension{}
	this.Context = context
	this.Exten = exten
	this.Type = type_
	return &this
}

// NewDestinationExtensionWithDefaults instantiates a new DestinationExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationExtensionWithDefaults() *DestinationExtension {
	this := DestinationExtension{}
	return &this
}

// GetContext returns the Context field value
func (o *DestinationExtension) GetContext() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Context
}

// GetContextOk returns a tuple with the Context field value
// and a boolean to check if the value has been set.
func (o *DestinationExtension) GetContextOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Context, true
}

// SetContext sets field value
func (o *DestinationExtension) SetContext(v string) {
	o.Context = v
}

// GetExten returns the Exten field value
func (o *DestinationExtension) GetExten() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exten
}

// GetExtenOk returns a tuple with the Exten field value
// and a boolean to check if the value has been set.
func (o *DestinationExtension) GetExtenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exten, true
}

// SetExten sets field value
func (o *DestinationExtension) SetExten(v string) {
	o.Exten = v
}

// GetType returns the Type field value
func (o *DestinationExtension) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *DestinationExtension) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *DestinationExtension) SetType(v string) {
	o.Type = v
}

func (o DestinationExtension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["context"] = o.Context
	toSerialize["exten"] = o.Exten
	toSerialize["type"] = o.Type
	return toSerialize, nil
}

func (o *DestinationExtension) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"context",
		"exten",
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

	varDestinationExtension := _DestinationExtension{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varDestinationExtension)

	if err != nil {
		return err
	}

	*o = DestinationExtension(varDestinationExtension)

	return err
}

type NullableDestinationExtension struct {
	value *DestinationExtension
	isSet bool
}

func (v NullableDestinationExtension) Get() *DestinationExtension {
	return v.value
}

func (v *NullableDestinationExtension) Set(val *DestinationExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationExtension(val *DestinationExtension) *NullableDestinationExtension {
	return &NullableDestinationExtension{value: val, isSet: true}
}

func (v NullableDestinationExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
