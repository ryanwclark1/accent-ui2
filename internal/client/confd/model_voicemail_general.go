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

// checks if the VoicemailGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VoicemailGeneral{}

// VoicemailGeneral Voicemail general configuration. This configuration will be in the file 'voicemail.conf' used by asterisk. Please consult the asterisk documentation for further details on available parameters.
type VoicemailGeneral struct {
	// General Voicemail options. These options must be unique and unordered. Option must have the following form:  ``` {   \"options\": {     \"name1\": \"value1\",     ...   } } ```
	Options map[string]interface{} `json:"options,omitempty"`
}

// NewVoicemailGeneral instantiates a new VoicemailGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVoicemailGeneral() *VoicemailGeneral {
	this := VoicemailGeneral{}
	return &this
}

// NewVoicemailGeneralWithDefaults instantiates a new VoicemailGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVoicemailGeneralWithDefaults() *VoicemailGeneral {
	this := VoicemailGeneral{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *VoicemailGeneral) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VoicemailGeneral) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *VoicemailGeneral) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *VoicemailGeneral) SetOptions(v map[string]interface{}) {
	o.Options = v
}

func (o VoicemailGeneral) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VoicemailGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	return toSerialize, nil
}

type NullableVoicemailGeneral struct {
	value *VoicemailGeneral
	isSet bool
}

func (v NullableVoicemailGeneral) Get() *VoicemailGeneral {
	return v.value
}

func (v *NullableVoicemailGeneral) Set(val *VoicemailGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableVoicemailGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableVoicemailGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVoicemailGeneral(val *VoicemailGeneral) *NullableVoicemailGeneral {
	return &NullableVoicemailGeneral{value: val, isSet: true}
}

func (v NullableVoicemailGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVoicemailGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
