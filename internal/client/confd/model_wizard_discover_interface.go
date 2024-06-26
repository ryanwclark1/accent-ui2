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

// checks if the WizardDiscoverInterface type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WizardDiscoverInterface{}

// WizardDiscoverInterface struct for WizardDiscoverInterface
type WizardDiscoverInterface struct {
	// Interface name (e.g. eth0)
	Interface *string `json:"interface,omitempty"`
	// IPv4 address of the interface
	IpAddress *string `json:"ip_address,omitempty"`
	// Netmask of the IP address
	Netmask *string `json:"netmask,omitempty"`
}

// NewWizardDiscoverInterface instantiates a new WizardDiscoverInterface object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWizardDiscoverInterface() *WizardDiscoverInterface {
	this := WizardDiscoverInterface{}
	return &this
}

// NewWizardDiscoverInterfaceWithDefaults instantiates a new WizardDiscoverInterface object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWizardDiscoverInterfaceWithDefaults() *WizardDiscoverInterface {
	this := WizardDiscoverInterface{}
	return &this
}

// GetInterface returns the Interface field value if set, zero value otherwise.
func (o *WizardDiscoverInterface) GetInterface() string {
	if o == nil || IsNil(o.Interface) {
		var ret string
		return ret
	}
	return *o.Interface
}

// GetInterfaceOk returns a tuple with the Interface field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardDiscoverInterface) GetInterfaceOk() (*string, bool) {
	if o == nil || IsNil(o.Interface) {
		return nil, false
	}
	return o.Interface, true
}

// HasInterface returns a boolean if a field has been set.
func (o *WizardDiscoverInterface) HasInterface() bool {
	if o != nil && !IsNil(o.Interface) {
		return true
	}

	return false
}

// SetInterface gets a reference to the given string and assigns it to the Interface field.
func (o *WizardDiscoverInterface) SetInterface(v string) {
	o.Interface = &v
}

// GetIpAddress returns the IpAddress field value if set, zero value otherwise.
func (o *WizardDiscoverInterface) GetIpAddress() string {
	if o == nil || IsNil(o.IpAddress) {
		var ret string
		return ret
	}
	return *o.IpAddress
}

// GetIpAddressOk returns a tuple with the IpAddress field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardDiscoverInterface) GetIpAddressOk() (*string, bool) {
	if o == nil || IsNil(o.IpAddress) {
		return nil, false
	}
	return o.IpAddress, true
}

// HasIpAddress returns a boolean if a field has been set.
func (o *WizardDiscoverInterface) HasIpAddress() bool {
	if o != nil && !IsNil(o.IpAddress) {
		return true
	}

	return false
}

// SetIpAddress gets a reference to the given string and assigns it to the IpAddress field.
func (o *WizardDiscoverInterface) SetIpAddress(v string) {
	o.IpAddress = &v
}

// GetNetmask returns the Netmask field value if set, zero value otherwise.
func (o *WizardDiscoverInterface) GetNetmask() string {
	if o == nil || IsNil(o.Netmask) {
		var ret string
		return ret
	}
	return *o.Netmask
}

// GetNetmaskOk returns a tuple with the Netmask field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WizardDiscoverInterface) GetNetmaskOk() (*string, bool) {
	if o == nil || IsNil(o.Netmask) {
		return nil, false
	}
	return o.Netmask, true
}

// HasNetmask returns a boolean if a field has been set.
func (o *WizardDiscoverInterface) HasNetmask() bool {
	if o != nil && !IsNil(o.Netmask) {
		return true
	}

	return false
}

// SetNetmask gets a reference to the given string and assigns it to the Netmask field.
func (o *WizardDiscoverInterface) SetNetmask(v string) {
	o.Netmask = &v
}

func (o WizardDiscoverInterface) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WizardDiscoverInterface) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Interface) {
		toSerialize["interface"] = o.Interface
	}
	if !IsNil(o.IpAddress) {
		toSerialize["ip_address"] = o.IpAddress
	}
	if !IsNil(o.Netmask) {
		toSerialize["netmask"] = o.Netmask
	}
	return toSerialize, nil
}

type NullableWizardDiscoverInterface struct {
	value *WizardDiscoverInterface
	isSet bool
}

func (v NullableWizardDiscoverInterface) Get() *WizardDiscoverInterface {
	return v.value
}

func (v *NullableWizardDiscoverInterface) Set(val *WizardDiscoverInterface) {
	v.value = val
	v.isSet = true
}

func (v NullableWizardDiscoverInterface) IsSet() bool {
	return v.isSet
}

func (v *NullableWizardDiscoverInterface) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWizardDiscoverInterface(val *WizardDiscoverInterface) *NullableWizardDiscoverInterface {
	return &NullableWizardDiscoverInterface{value: val, isSet: true}
}

func (v NullableWizardDiscoverInterface) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWizardDiscoverInterface) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
