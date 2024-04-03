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

// checks if the SoundFile type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SoundFile{}

// SoundFile struct for SoundFile
type SoundFile struct {
	// The audio file formats
	Formats []SoundFormat `json:"formats,omitempty"`
	// The name of the file
	Name *string `json:"name,omitempty"`
	// The UUID of the tenant of the file
	TenantUuid *string `json:"tenant_uuid,omitempty"`
}

// NewSoundFile instantiates a new SoundFile object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSoundFile() *SoundFile {
	this := SoundFile{}
	return &this
}

// NewSoundFileWithDefaults instantiates a new SoundFile object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSoundFileWithDefaults() *SoundFile {
	this := SoundFile{}
	return &this
}

// GetFormats returns the Formats field value if set, zero value otherwise.
func (o *SoundFile) GetFormats() []SoundFormat {
	if o == nil || IsNil(o.Formats) {
		var ret []SoundFormat
		return ret
	}
	return o.Formats
}

// GetFormatsOk returns a tuple with the Formats field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundFile) GetFormatsOk() ([]SoundFormat, bool) {
	if o == nil || IsNil(o.Formats) {
		return nil, false
	}
	return o.Formats, true
}

// HasFormats returns a boolean if a field has been set.
func (o *SoundFile) HasFormats() bool {
	if o != nil && !IsNil(o.Formats) {
		return true
	}

	return false
}

// SetFormats gets a reference to the given []SoundFormat and assigns it to the Formats field.
func (o *SoundFile) SetFormats(v []SoundFormat) {
	o.Formats = v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *SoundFile) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundFile) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *SoundFile) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *SoundFile) SetName(v string) {
	o.Name = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *SoundFile) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SoundFile) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *SoundFile) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *SoundFile) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

func (o SoundFile) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SoundFile) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Formats) {
		toSerialize["formats"] = o.Formats
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	return toSerialize, nil
}

type NullableSoundFile struct {
	value *SoundFile
	isSet bool
}

func (v NullableSoundFile) Get() *SoundFile {
	return v.value
}

func (v *NullableSoundFile) Set(val *SoundFile) {
	v.value = val
	v.isSet = true
}

func (v NullableSoundFile) IsSet() bool {
	return v.isSet
}

func (v *NullableSoundFile) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSoundFile(val *SoundFile) *NullableSoundFile {
	return &NullableSoundFile{value: val, isSet: true}
}

func (v NullableSoundFile) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSoundFile) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
