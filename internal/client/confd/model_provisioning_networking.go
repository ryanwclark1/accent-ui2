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

// checks if the ProvisioningNetworking type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &ProvisioningNetworking{}

// ProvisioningNetworking struct for ProvisioningNetworking
type ProvisioningNetworking struct {
	// The hostname or IP address used for HTTP and TFTP provisioning requests for DHCP integration.
	ProvisionHost *string `json:"provision_host,omitempty"`
	// The base URL on which the provisioning server will be accessible from outside the network.
	ProvisionHttpBaseUrl *string `json:"provision_http_base_url,omitempty"`
	// The port used by the HTTP provisioning server.
	ProvisionHttpPort *int32 `json:"provision_http_port,omitempty"`
}

// NewProvisioningNetworking instantiates a new ProvisioningNetworking object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProvisioningNetworking() *ProvisioningNetworking {
	this := ProvisioningNetworking{}
	return &this
}

// NewProvisioningNetworkingWithDefaults instantiates a new ProvisioningNetworking object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProvisioningNetworkingWithDefaults() *ProvisioningNetworking {
	this := ProvisioningNetworking{}
	return &this
}

// GetProvisionHost returns the ProvisionHost field value if set, zero value otherwise.
func (o *ProvisioningNetworking) GetProvisionHost() string {
	if o == nil || IsNil(o.ProvisionHost) {
		var ret string
		return ret
	}
	return *o.ProvisionHost
}

// GetProvisionHostOk returns a tuple with the ProvisionHost field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningNetworking) GetProvisionHostOk() (*string, bool) {
	if o == nil || IsNil(o.ProvisionHost) {
		return nil, false
	}
	return o.ProvisionHost, true
}

// HasProvisionHost returns a boolean if a field has been set.
func (o *ProvisioningNetworking) HasProvisionHost() bool {
	if o != nil && !IsNil(o.ProvisionHost) {
		return true
	}

	return false
}

// SetProvisionHost gets a reference to the given string and assigns it to the ProvisionHost field.
func (o *ProvisioningNetworking) SetProvisionHost(v string) {
	o.ProvisionHost = &v
}

// GetProvisionHttpBaseUrl returns the ProvisionHttpBaseUrl field value if set, zero value otherwise.
func (o *ProvisioningNetworking) GetProvisionHttpBaseUrl() string {
	if o == nil || IsNil(o.ProvisionHttpBaseUrl) {
		var ret string
		return ret
	}
	return *o.ProvisionHttpBaseUrl
}

// GetProvisionHttpBaseUrlOk returns a tuple with the ProvisionHttpBaseUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningNetworking) GetProvisionHttpBaseUrlOk() (*string, bool) {
	if o == nil || IsNil(o.ProvisionHttpBaseUrl) {
		return nil, false
	}
	return o.ProvisionHttpBaseUrl, true
}

// HasProvisionHttpBaseUrl returns a boolean if a field has been set.
func (o *ProvisioningNetworking) HasProvisionHttpBaseUrl() bool {
	if o != nil && !IsNil(o.ProvisionHttpBaseUrl) {
		return true
	}

	return false
}

// SetProvisionHttpBaseUrl gets a reference to the given string and assigns it to the ProvisionHttpBaseUrl field.
func (o *ProvisioningNetworking) SetProvisionHttpBaseUrl(v string) {
	o.ProvisionHttpBaseUrl = &v
}

// GetProvisionHttpPort returns the ProvisionHttpPort field value if set, zero value otherwise.
func (o *ProvisioningNetworking) GetProvisionHttpPort() int32 {
	if o == nil || IsNil(o.ProvisionHttpPort) {
		var ret int32
		return ret
	}
	return *o.ProvisionHttpPort
}

// GetProvisionHttpPortOk returns a tuple with the ProvisionHttpPort field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProvisioningNetworking) GetProvisionHttpPortOk() (*int32, bool) {
	if o == nil || IsNil(o.ProvisionHttpPort) {
		return nil, false
	}
	return o.ProvisionHttpPort, true
}

// HasProvisionHttpPort returns a boolean if a field has been set.
func (o *ProvisioningNetworking) HasProvisionHttpPort() bool {
	if o != nil && !IsNil(o.ProvisionHttpPort) {
		return true
	}

	return false
}

// SetProvisionHttpPort gets a reference to the given int32 and assigns it to the ProvisionHttpPort field.
func (o *ProvisioningNetworking) SetProvisionHttpPort(v int32) {
	o.ProvisionHttpPort = &v
}

func (o ProvisioningNetworking) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o ProvisioningNetworking) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.ProvisionHost) {
		toSerialize["provision_host"] = o.ProvisionHost
	}
	if !IsNil(o.ProvisionHttpBaseUrl) {
		toSerialize["provision_http_base_url"] = o.ProvisionHttpBaseUrl
	}
	if !IsNil(o.ProvisionHttpPort) {
		toSerialize["provision_http_port"] = o.ProvisionHttpPort
	}
	return toSerialize, nil
}

type NullableProvisioningNetworking struct {
	value *ProvisioningNetworking
	isSet bool
}

func (v NullableProvisioningNetworking) Get() *ProvisioningNetworking {
	return v.value
}

func (v *NullableProvisioningNetworking) Set(val *ProvisioningNetworking) {
	v.value = val
	v.isSet = true
}

func (v NullableProvisioningNetworking) IsSet() bool {
	return v.isSet
}

func (v *NullableProvisioningNetworking) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProvisioningNetworking(val *ProvisioningNetworking) *NullableProvisioningNetworking {
	return &NullableProvisioningNetworking{value: val, isSet: true}
}

func (v NullableProvisioningNetworking) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProvisioningNetworking) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
