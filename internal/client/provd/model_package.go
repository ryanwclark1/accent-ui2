/*
accent-provd

Provisioning application REST API

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package provd

import (
	"encoding/json"
)

// checks if the Package type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Package{}

// Package struct for Package
type Package struct {
	Capabilities *map[string]map[string]string `json:"capabilities,omitempty"`
	Description  *string                       `json:"description,omitempty"`
	Dsize        *int32                        `json:"dsize,omitempty"`
	Sha1sum      *string                       `json:"sha1sum,omitempty"`
	Version      *string                       `json:"version,omitempty"`
}

// NewPackage instantiates a new Package object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPackage() *Package {
	this := Package{}
	return &this
}

// NewPackageWithDefaults instantiates a new Package object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPackageWithDefaults() *Package {
	this := Package{}
	return &this
}

// GetCapabilities returns the Capabilities field value if set, zero value otherwise.
func (o *Package) GetCapabilities() map[string]map[string]string {
	if o == nil || IsNil(o.Capabilities) {
		var ret map[string]map[string]string
		return ret
	}
	return *o.Capabilities
}

// GetCapabilitiesOk returns a tuple with the Capabilities field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetCapabilitiesOk() (*map[string]map[string]string, bool) {
	if o == nil || IsNil(o.Capabilities) {
		return nil, false
	}
	return o.Capabilities, true
}

// HasCapabilities returns a boolean if a field has been set.
func (o *Package) HasCapabilities() bool {
	if o != nil && !IsNil(o.Capabilities) {
		return true
	}

	return false
}

// SetCapabilities gets a reference to the given map[string]map[string]string and assigns it to the Capabilities field.
func (o *Package) SetCapabilities(v map[string]map[string]string) {
	o.Capabilities = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Package) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Package) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Package) SetDescription(v string) {
	o.Description = &v
}

// GetDsize returns the Dsize field value if set, zero value otherwise.
func (o *Package) GetDsize() int32 {
	if o == nil || IsNil(o.Dsize) {
		var ret int32
		return ret
	}
	return *o.Dsize
}

// GetDsizeOk returns a tuple with the Dsize field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetDsizeOk() (*int32, bool) {
	if o == nil || IsNil(o.Dsize) {
		return nil, false
	}
	return o.Dsize, true
}

// HasDsize returns a boolean if a field has been set.
func (o *Package) HasDsize() bool {
	if o != nil && !IsNil(o.Dsize) {
		return true
	}

	return false
}

// SetDsize gets a reference to the given int32 and assigns it to the Dsize field.
func (o *Package) SetDsize(v int32) {
	o.Dsize = &v
}

// GetSha1sum returns the Sha1sum field value if set, zero value otherwise.
func (o *Package) GetSha1sum() string {
	if o == nil || IsNil(o.Sha1sum) {
		var ret string
		return ret
	}
	return *o.Sha1sum
}

// GetSha1sumOk returns a tuple with the Sha1sum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetSha1sumOk() (*string, bool) {
	if o == nil || IsNil(o.Sha1sum) {
		return nil, false
	}
	return o.Sha1sum, true
}

// HasSha1sum returns a boolean if a field has been set.
func (o *Package) HasSha1sum() bool {
	if o != nil && !IsNil(o.Sha1sum) {
		return true
	}

	return false
}

// SetSha1sum gets a reference to the given string and assigns it to the Sha1sum field.
func (o *Package) SetSha1sum(v string) {
	o.Sha1sum = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *Package) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Package) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *Package) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *Package) SetVersion(v string) {
	o.Version = &v
}

func (o Package) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Package) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Capabilities) {
		toSerialize["capabilities"] = o.Capabilities
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Dsize) {
		toSerialize["dsize"] = o.Dsize
	}
	if !IsNil(o.Sha1sum) {
		toSerialize["sha1sum"] = o.Sha1sum
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullablePackage struct {
	value *Package
	isSet bool
}

func (v NullablePackage) Get() *Package {
	return v.value
}

func (v *NullablePackage) Set(val *Package) {
	v.value = val
	v.isSet = true
}

func (v NullablePackage) IsSet() bool {
	return v.isSet
}

func (v *NullablePackage) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePackage(val *Package) *NullablePackage {
	return &NullablePackage{value: val, isSet: true}
}

func (v NullablePackage) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePackage) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
