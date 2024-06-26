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

// checks if the WizardConfigured type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardConfigured{}

// WizardConfigured struct for WizardConfigured
type WizardConfigured struct {
	// Whether all services which the wizard depends on are started or not
	Configurable *bool `json:"configurable,omitempty"`
	// A mapping of all dependencies and there statuses
	ConfigurableStatus map[string]interface{} `json:"configurable_status,omitempty"`
	// Whether Accent has already been configured or not
	Configured *bool `json:"configured,omitempty"`
}

// NewWizardConfigured instantiates a new WizardConfigured object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardConfigured() *WizardConfigured {
	this := WizardConfigured{}
	return &this
}

// NewWizardConfiguredWithDefaults instantiates a new WizardConfigured object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardConfiguredWithDefaults() *WizardConfigured {
	this := WizardConfigured{}
	return &this
}

// GetConfigurable returns the Configurable field value if set, zero value otherwise.
func (o *WizardConfigured) GetConfigurable() bool {
	if o == nil || IsNil(o.Configurable) {
		var ret bool
		return ret
	}
	return *o.Configurable
}

// GetConfigurableOk returns a tuple with the Configurable field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardConfigured) GetConfigurableOk() (*bool, bool) {
	if o == nil || IsNil(o.Configurable) {
		return nil, false
	}
	return o.Configurable, true
}

// HasConfigurable returns a boolean if a field has been set.
func (o *WizardConfigured) HasConfigurable() bool {
	if o != nil && !IsNil(o.Configurable) {
		return true
	}

	return false
}

// SetConfigurable gets a reference to the given bool and assigns it to the Configurable field.
func (o *WizardConfigured) SetConfigurable(v bool) {
	o.Configurable = &v
}

// GetConfigurableStatus returns the ConfigurableStatus field value if set, zero value otherwise.
func (o *WizardConfigured) GetConfigurableStatus() map[string]interface{} {
	if o == nil || IsNil(o.ConfigurableStatus) {
		var ret map[string]interface{}
		return ret
	}
	return o.ConfigurableStatus
}

// GetConfigurableStatusOk returns a tuple with the ConfigurableStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardConfigured) GetConfigurableStatusOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.ConfigurableStatus) {
		return map[string]interface{}{}, false
	}
	return o.ConfigurableStatus, true
}

// HasConfigurableStatus returns a boolean if a field has been set.
func (o *WizardConfigured) HasConfigurableStatus() bool {
	if o != nil && !IsNil(o.ConfigurableStatus) {
		return true
	}

	return false
}

// SetConfigurableStatus gets a reference to the given map[string]interface{} and assigns it to the ConfigurableStatus field.
func (o *WizardConfigured) SetConfigurableStatus(v map[string]interface{}) {
	o.ConfigurableStatus = v
}

// GetConfigured returns the Configured field value if set, zero value otherwise.
func (o *WizardConfigured) GetConfigured() bool {
	if o == nil || IsNil(o.Configured) {
		var ret bool
		return ret
	}
	return *o.Configured
}

// GetConfiguredOk returns a tuple with the Configured field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardConfigured) GetConfiguredOk() (*bool, bool) {
	if o == nil || IsNil(o.Configured) {
		return nil, false
	}
	return o.Configured, true
}

// HasConfigured returns a boolean if a field has been set.
func (o *WizardConfigured) HasConfigured() bool {
	if o != nil && !IsNil(o.Configured) {
		return true
	}

	return false
}

// SetConfigured gets a reference to the given bool and assigns it to the Configured field.
func (o *WizardConfigured) SetConfigured(v bool) {
	o.Configured = &v
}

func (o WizardConfigured) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardConfigured) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Configurable) {
		toSerialize["configurable"] = o.Configurable
	}
	if !IsNil(o.ConfigurableStatus) {
		toSerialize["configurable_status"] = o.ConfigurableStatus
	}
	if !IsNil(o.Configured) {
		toSerialize["configured"] = o.Configured
	}
	return toSerialize, nil
}

type NullableWizardConfigured struct {
	value *WizardConfigured
	isSet bool
}

func (v NullableWizardConfigured) Get() *WizardConfigured {
	return v.value
}

func (v *NullableWizardConfigured) Set(val *WizardConfigured) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardConfigured) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardConfigured) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardConfigured(val *WizardConfigured) *NullableWizardConfigured {
	return &NullableWizardConfigured{value: val, isSet: true}
}

func (v NullableWizardConfigured) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardConfigured) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
