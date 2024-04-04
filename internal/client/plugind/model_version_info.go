/*
accent-plugind

Accent's plugin management service

API version: 0.2
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package plugind

import (
	"encoding/json"
)

// checks if the VersionInfo type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &VersionInfo{}

// VersionInfo struct for VersionInfo
type VersionInfo struct {
	// The maximum Accent version with which this plugin works
	MaxAccentVersion *string `json:"max_accent_version,omitempty"`
	// The minimum Accent version with which this plugin works
	MinAccentVersion *string `json:"min_accent_version,omitempty"`
	// An indication wether installing this version would be an upgrade on not. Unstalled plugins are marked as upgradable.
	Upgradable *bool `json:"upgradable,omitempty"`
	// The plugin version
	Version *string `json:"version,omitempty"`
}

// NewVersionInfo instantiates a new VersionInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewVersionInfo() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// NewVersionInfoWithDefaults instantiates a new VersionInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewVersionInfoWithDefaults() *VersionInfo {
	this := VersionInfo{}
	return &this
}

// GetMaxAccentVersion returns the MaxAccentVersion field value if set, zero value otherwise.
func (o *VersionInfo) GetMaxAccentVersion() string {
	if o == nil || IsNil(o.MaxAccentVersion) {
		var ret string
		return ret
	}
	return *o.MaxAccentVersion
}

// GetMaxAccentVersionOk returns a tuple with the MaxAccentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetMaxAccentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MaxAccentVersion) {
		return nil, false
	}
	return o.MaxAccentVersion, true
}

// HasMaxAccentVersion returns a boolean if a field has been set.
func (o *VersionInfo) HasMaxAccentVersion() bool {
	if o != nil && !IsNil(o.MaxAccentVersion) {
		return true
	}

	return false
}

// SetMaxAccentVersion gets a reference to the given string and assigns it to the MaxAccentVersion field.
func (o *VersionInfo) SetMaxAccentVersion(v string) {
	o.MaxAccentVersion = &v
}

// GetMinAccentVersion returns the MinAccentVersion field value if set, zero value otherwise.
func (o *VersionInfo) GetMinAccentVersion() string {
	if o == nil || IsNil(o.MinAccentVersion) {
		var ret string
		return ret
	}
	return *o.MinAccentVersion
}

// GetMinAccentVersionOk returns a tuple with the MinAccentVersion field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetMinAccentVersionOk() (*string, bool) {
	if o == nil || IsNil(o.MinAccentVersion) {
		return nil, false
	}
	return o.MinAccentVersion, true
}

// HasMinAccentVersion returns a boolean if a field has been set.
func (o *VersionInfo) HasMinAccentVersion() bool {
	if o != nil && !IsNil(o.MinAccentVersion) {
		return true
	}

	return false
}

// SetMinAccentVersion gets a reference to the given string and assigns it to the MinAccentVersion field.
func (o *VersionInfo) SetMinAccentVersion(v string) {
	o.MinAccentVersion = &v
}

// GetUpgradable returns the Upgradable field value if set, zero value otherwise.
func (o *VersionInfo) GetUpgradable() bool {
	if o == nil || IsNil(o.Upgradable) {
		var ret bool
		return ret
	}
	return *o.Upgradable
}

// GetUpgradableOk returns a tuple with the Upgradable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetUpgradableOk() (*bool, bool) {
	if o == nil || IsNil(o.Upgradable) {
		return nil, false
	}
	return o.Upgradable, true
}

// HasUpgradable returns a boolean if a field has been set.
func (o *VersionInfo) HasUpgradable() bool {
	if o != nil && !IsNil(o.Upgradable) {
		return true
	}

	return false
}

// SetUpgradable gets a reference to the given bool and assigns it to the Upgradable field.
func (o *VersionInfo) SetUpgradable(v bool) {
	o.Upgradable = &v
}

// GetVersion returns the Version field value if set, zero value otherwise.
func (o *VersionInfo) GetVersion() string {
	if o == nil || IsNil(o.Version) {
		var ret string
		return ret
	}
	return *o.Version
}

// GetVersionOk returns a tuple with the Version field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *VersionInfo) GetVersionOk() (*string, bool) {
	if o == nil || IsNil(o.Version) {
		return nil, false
	}
	return o.Version, true
}

// HasVersion returns a boolean if a field has been set.
func (o *VersionInfo) HasVersion() bool {
	if o != nil && !IsNil(o.Version) {
		return true
	}

	return false
}

// SetVersion gets a reference to the given string and assigns it to the Version field.
func (o *VersionInfo) SetVersion(v string) {
	o.Version = &v
}

func (o VersionInfo) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o VersionInfo) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.MaxAccentVersion) {
		toSerialize["max_accent_version"] = o.MaxAccentVersion
	}
	if !IsNil(o.MinAccentVersion) {
		toSerialize["min_accent_version"] = o.MinAccentVersion
	}
	if !IsNil(o.Upgradable) {
		toSerialize["upgradable"] = o.Upgradable
	}
	if !IsNil(o.Version) {
		toSerialize["version"] = o.Version
	}
	return toSerialize, nil
}

type NullableVersionInfo struct {
	value *VersionInfo
	isSet bool
}

func (v NullableVersionInfo) Get() *VersionInfo {
	return v.value
}

func (v *NullableVersionInfo) Set(val *VersionInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableVersionInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableVersionInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableVersionInfo(val *VersionInfo) *NullableVersionInfo {
	return &NullableVersionInfo{value: val, isSet: true}
}

func (v NullableVersionInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableVersionInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
