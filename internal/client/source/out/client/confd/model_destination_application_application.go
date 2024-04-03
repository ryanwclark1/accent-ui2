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

// checks if the DestinationApplicationApplication type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &DestinationApplicationApplication{}

// DestinationApplicationApplication struct for DestinationApplicationApplication
type DestinationApplicationApplication struct {
	CallbackDisa *DestinationApplicationCallbackDISA `json:"callback_disa,omitempty"`
	Custom       *DestinationApplicationCustom       `json:"custom,omitempty"`
	Directory    *DestinationApplicationDirectory    `json:"directory,omitempty"`
	Disa         *DestinationApplicationDISA         `json:"disa,omitempty"`
	FaxToMail    *DestinationApplicationFaxToMail    `json:"fax_to_mail,omitempty"`
	Voicemail    *DestinationApplicationVoicemail    `json:"voicemail,omitempty"`
}

// NewDestinationApplicationApplication instantiates a new DestinationApplicationApplication object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewDestinationApplicationApplication() *DestinationApplicationApplication {
	this := DestinationApplicationApplication{}
	return &this
}

// NewDestinationApplicationApplicationWithDefaults instantiates a new DestinationApplicationApplication object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewDestinationApplicationApplicationWithDefaults() *DestinationApplicationApplication {
	this := DestinationApplicationApplication{}
	return &this
}

// GetCallbackDisa returns the CallbackDisa field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetCallbackDisa() DestinationApplicationCallbackDISA {
	if o == nil || IsNil(o.CallbackDisa) {
		var ret DestinationApplicationCallbackDISA
		return ret
	}
	return *o.CallbackDisa
}

// GetCallbackDisaOk returns a tuple with the CallbackDisa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetCallbackDisaOk() (*DestinationApplicationCallbackDISA, bool) {
	if o == nil || IsNil(o.CallbackDisa) {
		return nil, false
	}
	return o.CallbackDisa, true
}

// HasCallbackDisa returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasCallbackDisa() bool {
	if o != nil && !IsNil(o.CallbackDisa) {
		return true
	}

	return false
}

// SetCallbackDisa gets a reference to the given DestinationApplicationCallbackDISA and assigns it to the CallbackDisa field.
func (o *DestinationApplicationApplication) SetCallbackDisa(v DestinationApplicationCallbackDISA) {
	o.CallbackDisa = &v
}

// GetCustom returns the Custom field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetCustom() DestinationApplicationCustom {
	if o == nil || IsNil(o.Custom) {
		var ret DestinationApplicationCustom
		return ret
	}
	return *o.Custom
}

// GetCustomOk returns a tuple with the Custom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetCustomOk() (*DestinationApplicationCustom, bool) {
	if o == nil || IsNil(o.Custom) {
		return nil, false
	}
	return o.Custom, true
}

// HasCustom returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasCustom() bool {
	if o != nil && !IsNil(o.Custom) {
		return true
	}

	return false
}

// SetCustom gets a reference to the given DestinationApplicationCustom and assigns it to the Custom field.
func (o *DestinationApplicationApplication) SetCustom(v DestinationApplicationCustom) {
	o.Custom = &v
}

// GetDirectory returns the Directory field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetDirectory() DestinationApplicationDirectory {
	if o == nil || IsNil(o.Directory) {
		var ret DestinationApplicationDirectory
		return ret
	}
	return *o.Directory
}

// GetDirectoryOk returns a tuple with the Directory field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetDirectoryOk() (*DestinationApplicationDirectory, bool) {
	if o == nil || IsNil(o.Directory) {
		return nil, false
	}
	return o.Directory, true
}

// HasDirectory returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasDirectory() bool {
	if o != nil && !IsNil(o.Directory) {
		return true
	}

	return false
}

// SetDirectory gets a reference to the given DestinationApplicationDirectory and assigns it to the Directory field.
func (o *DestinationApplicationApplication) SetDirectory(v DestinationApplicationDirectory) {
	o.Directory = &v
}

// GetDisa returns the Disa field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetDisa() DestinationApplicationDISA {
	if o == nil || IsNil(o.Disa) {
		var ret DestinationApplicationDISA
		return ret
	}
	return *o.Disa
}

// GetDisaOk returns a tuple with the Disa field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetDisaOk() (*DestinationApplicationDISA, bool) {
	if o == nil || IsNil(o.Disa) {
		return nil, false
	}
	return o.Disa, true
}

// HasDisa returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasDisa() bool {
	if o != nil && !IsNil(o.Disa) {
		return true
	}

	return false
}

// SetDisa gets a reference to the given DestinationApplicationDISA and assigns it to the Disa field.
func (o *DestinationApplicationApplication) SetDisa(v DestinationApplicationDISA) {
	o.Disa = &v
}

// GetFaxToMail returns the FaxToMail field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetFaxToMail() DestinationApplicationFaxToMail {
	if o == nil || IsNil(o.FaxToMail) {
		var ret DestinationApplicationFaxToMail
		return ret
	}
	return *o.FaxToMail
}

// GetFaxToMailOk returns a tuple with the FaxToMail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetFaxToMailOk() (*DestinationApplicationFaxToMail, bool) {
	if o == nil || IsNil(o.FaxToMail) {
		return nil, false
	}
	return o.FaxToMail, true
}

// HasFaxToMail returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasFaxToMail() bool {
	if o != nil && !IsNil(o.FaxToMail) {
		return true
	}

	return false
}

// SetFaxToMail gets a reference to the given DestinationApplicationFaxToMail and assigns it to the FaxToMail field.
func (o *DestinationApplicationApplication) SetFaxToMail(v DestinationApplicationFaxToMail) {
	o.FaxToMail = &v
}

// GetVoicemail returns the Voicemail field value if set, zero value otherwise.
func (o *DestinationApplicationApplication) GetVoicemail() DestinationApplicationVoicemail {
	if o == nil || IsNil(o.Voicemail) {
		var ret DestinationApplicationVoicemail
		return ret
	}
	return *o.Voicemail
}

// GetVoicemailOk returns a tuple with the Voicemail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *DestinationApplicationApplication) GetVoicemailOk() (*DestinationApplicationVoicemail, bool) {
	if o == nil || IsNil(o.Voicemail) {
		return nil, false
	}
	return o.Voicemail, true
}

// HasVoicemail returns a boolean if a field has been set.
func (o *DestinationApplicationApplication) HasVoicemail() bool {
	if o != nil && !IsNil(o.Voicemail) {
		return true
	}

	return false
}

// SetVoicemail gets a reference to the given DestinationApplicationVoicemail and assigns it to the Voicemail field.
func (o *DestinationApplicationApplication) SetVoicemail(v DestinationApplicationVoicemail) {
	o.Voicemail = &v
}

func (o DestinationApplicationApplication) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o DestinationApplicationApplication) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.CallbackDisa) {
		toSerialize["callback_disa"] = o.CallbackDisa
	}
	if !IsNil(o.Custom) {
		toSerialize["custom"] = o.Custom
	}
	if !IsNil(o.Directory) {
		toSerialize["directory"] = o.Directory
	}
	if !IsNil(o.Disa) {
		toSerialize["disa"] = o.Disa
	}
	if !IsNil(o.FaxToMail) {
		toSerialize["fax_to_mail"] = o.FaxToMail
	}
	if !IsNil(o.Voicemail) {
		toSerialize["voicemail"] = o.Voicemail
	}
	return toSerialize, nil
}

type NullableDestinationApplicationApplication struct {
	value *DestinationApplicationApplication
	isSet bool
}

func (v NullableDestinationApplicationApplication) Get() *DestinationApplicationApplication {
	return v.value
}

func (v *NullableDestinationApplicationApplication) Set(val *DestinationApplicationApplication) {
	v.value = val
	v.isSet = true
}

func (v NullableDestinationApplicationApplication) IsSet() bool {
	return v.isSet
}

func (v *NullableDestinationApplicationApplication) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableDestinationApplicationApplication(val *DestinationApplicationApplication) *NullableDestinationApplicationApplication {
	return &NullableDestinationApplicationApplication{value: val, isSet: true}
}

func (v NullableDestinationApplicationApplication) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableDestinationApplicationApplication) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
