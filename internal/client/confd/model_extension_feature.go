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

// checks if the ExtensionFeature type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ExtensionFeature{}

// ExtensionFeature struct for ExtensionFeature
type ExtensionFeature struct {
	Context *string `json:"context,omitempty"`
	// If False the extension feature is disabled
	Enabled *bool  `json:"enabled,omitempty"`
	Exten   string `json:"exten"`
	// The feature of the extension
	Feature *string `json:"feature,omitempty"`
	// Extension UUID
	Uuid *string `json:"uuid,omitempty"`
}

type _ExtensionFeature ExtensionFeature

// NewExtensionFeature instantiates a new ExtensionFeature object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewExtensionFeature(exten string) *ExtensionFeature {
	this := ExtensionFeature{}
	this.Exten = exten
	return &this
}

// NewExtensionFeatureWithDefaults instantiates a new ExtensionFeature object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewExtensionFeatureWithDefaults() *ExtensionFeature {
	this := ExtensionFeature{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *ExtensionFeature) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionFeature) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *ExtensionFeature) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *ExtensionFeature) SetContext(v string) {
	o.Context = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *ExtensionFeature) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionFeature) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *ExtensionFeature) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *ExtensionFeature) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetExten returns the Exten field value
func (o *ExtensionFeature) GetExten() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Exten
}

// GetExtenOk returns a tuple with the Exten field value
// and a boolean to check if the value has been set.
func (o *ExtensionFeature) GetExtenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Exten, true
}

// SetExten sets field value
func (o *ExtensionFeature) SetExten(v string) {
	o.Exten = v
}

// GetFeature returns the Feature field value if set, zero value otherwise.
func (o *ExtensionFeature) GetFeature() string {
	if o == nil || IsNil(o.Feature) {
		var ret string
		return ret
	}
	return *o.Feature
}

// GetFeatureOk returns a tuple with the Feature field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionFeature) GetFeatureOk() (*string, bool) {
	if o == nil || IsNil(o.Feature) {
		return nil, false
	}
	return o.Feature, true
}

// HasFeature returns a boolean if a field has been set.
func (o *ExtensionFeature) HasFeature() bool {
	if o != nil && !IsNil(o.Feature) {
		return true
	}

	return false
}

// SetFeature gets a reference to the given string and assigns it to the Feature field.
func (o *ExtensionFeature) SetFeature(v string) {
	o.Feature = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *ExtensionFeature) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ExtensionFeature) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *ExtensionFeature) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *ExtensionFeature) SetUuid(v string) {
	o.Uuid = &v
}

func (o ExtensionFeature) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ExtensionFeature) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	toSerialize["exten"] = o.Exten
	if !IsNil(o.Feature) {
		toSerialize["feature"] = o.Feature
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

func (o *ExtensionFeature) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"exten",
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

	varExtensionFeature := _ExtensionFeature{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varExtensionFeature)

	if err != nil {
		return err
	}

	*o = ExtensionFeature(varExtensionFeature)

	return err
}

type NullableExtensionFeature struct {
	value *ExtensionFeature
	isSet bool
}

func (v NullableExtensionFeature) Get() *ExtensionFeature {
	return v.value
}

func (v *NullableExtensionFeature) Set(val *ExtensionFeature) {
	v.value = val
	v.isSet = true
}

func (v NullableExtensionFeature) IsSet() bool {
	return v.isSet
}

func (v *NullableExtensionFeature) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtensionFeature(val *ExtensionFeature) *NullableExtensionFeature {
	return &NullableExtensionFeature{value: val, isSet: true}
}

func (v NullableExtensionFeature) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtensionFeature) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
