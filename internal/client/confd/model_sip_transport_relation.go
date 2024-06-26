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

// checks if the SIPTransportRelation type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SIPTransportRelation{}

// SIPTransportRelation struct for SIPTransportRelation
type SIPTransportRelation struct {
	Uuid string `json:"uuid"`
}

type _SIPTransportRelation SIPTransportRelation

// NewSIPTransportRelation instantiates a new SIPTransportRelation object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSIPTransportRelation(uuid string) *SIPTransportRelation {
	this := SIPTransportRelation{}
	this.Uuid = uuid
	return &this
}

// NewSIPTransportRelationWithDefaults instantiates a new SIPTransportRelation object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSIPTransportRelationWithDefaults() *SIPTransportRelation {
	this := SIPTransportRelation{}
	return &this
}

// GetUuid returns the Uuid field value
func (o *SIPTransportRelation) GetUuid() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value
// and a boolean to check if the value has been set.
func (o *SIPTransportRelation) GetUuidOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Uuid, true
}

// SetUuid sets field value
func (o *SIPTransportRelation) SetUuid(v string) {
	o.Uuid = v
}

func (o SIPTransportRelation) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SIPTransportRelation) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["uuid"] = o.Uuid
	return toSerialize, nil
}

func (o *SIPTransportRelation) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"uuid",
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

	varSIPTransportRelation := _SIPTransportRelation{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSIPTransportRelation)

	if err != nil {
		return err
	}

	*o = SIPTransportRelation(varSIPTransportRelation)

	return err
}

type NullableSIPTransportRelation struct {
	value *SIPTransportRelation
	isSet bool
}

func (v NullableSIPTransportRelation) Get() *SIPTransportRelation {
	return v.value
}

func (v *NullableSIPTransportRelation) Set(val *SIPTransportRelation) {
	v.value = val
	v.isSet = true
}

func (v NullableSIPTransportRelation) IsSet() bool {
	return v.isSet
}

func (v *NullableSIPTransportRelation) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSIPTransportRelation(val *SIPTransportRelation) *NullableSIPTransportRelation {
	return &NullableSIPTransportRelation{value: val, isSet: true}
}

func (v NullableSIPTransportRelation) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSIPTransportRelation) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
