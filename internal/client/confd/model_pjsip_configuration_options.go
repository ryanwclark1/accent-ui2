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

// checks if the PJSIPConfigurationOptions type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PJSIPConfigurationOptions{}

// PJSIPConfigurationOptions struct for PJSIPConfigurationOptions
type PJSIPConfigurationOptions struct {
	// A list of all configuration options for this section
	Aor []PJSIPConfigurationOption `json:"aor,omitempty"`
	// A list of all configuration options for this section
	Auth []PJSIPConfigurationOption `json:"auth,omitempty"`
	// A list of all configuration options for this section
	Contact []PJSIPConfigurationOption `json:"contact,omitempty"`
	// A list of all configuration options for this section
	DomainAlias []PJSIPConfigurationOption `json:"domain_alias,omitempty"`
	// A list of all configuration options for this section
	Endpoint []PJSIPConfigurationOption `json:"endpoint,omitempty"`
	// A list of all configuration options for this section
	Global []PJSIPConfigurationOption `json:"global,omitempty"`
	// A list of all configuration options for this section
	Registration []PJSIPConfigurationOption `json:"registration,omitempty"`
	// A list of all configuration options for this section
	System []PJSIPConfigurationOption `json:"system,omitempty"`
	// A list of all configuration options for this section
	Transport []PJSIPConfigurationOption `json:"transport,omitempty"`
}

// NewPJSIPConfigurationOptions instantiates a new PJSIPConfigurationOptions object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPJSIPConfigurationOptions() *PJSIPConfigurationOptions {
	this := PJSIPConfigurationOptions{}
	return &this
}

// NewPJSIPConfigurationOptionsWithDefaults instantiates a new PJSIPConfigurationOptions object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPJSIPConfigurationOptionsWithDefaults() *PJSIPConfigurationOptions {
	this := PJSIPConfigurationOptions{}
	return &this
}

// GetAor returns the Aor field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetAor() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Aor) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Aor
}

// GetAorOk returns a tuple with the Aor field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetAorOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Aor) {
		return nil, false
	}
	return o.Aor, true
}

// HasAor returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasAor() bool {
	if o != nil && !IsNil(o.Aor) {
		return true
	}

	return false
}

// SetAor gets a reference to the given []PJSIPConfigurationOption and assigns it to the Aor field.
func (o *PJSIPConfigurationOptions) SetAor(v []PJSIPConfigurationOption) {
	o.Aor = v
}

// GetAuth returns the Auth field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetAuth() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Auth) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Auth
}

// GetAuthOk returns a tuple with the Auth field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetAuthOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Auth) {
		return nil, false
	}
	return o.Auth, true
}

// HasAuth returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasAuth() bool {
	if o != nil && !IsNil(o.Auth) {
		return true
	}

	return false
}

// SetAuth gets a reference to the given []PJSIPConfigurationOption and assigns it to the Auth field.
func (o *PJSIPConfigurationOptions) SetAuth(v []PJSIPConfigurationOption) {
	o.Auth = v
}

// GetContact returns the Contact field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetContact() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Contact) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Contact
}

// GetContactOk returns a tuple with the Contact field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetContactOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Contact) {
		return nil, false
	}
	return o.Contact, true
}

// HasContact returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasContact() bool {
	if o != nil && !IsNil(o.Contact) {
		return true
	}

	return false
}

// SetContact gets a reference to the given []PJSIPConfigurationOption and assigns it to the Contact field.
func (o *PJSIPConfigurationOptions) SetContact(v []PJSIPConfigurationOption) {
	o.Contact = v
}

// GetDomainAlias returns the DomainAlias field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetDomainAlias() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.DomainAlias) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.DomainAlias
}

// GetDomainAliasOk returns a tuple with the DomainAlias field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetDomainAliasOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.DomainAlias) {
		return nil, false
	}
	return o.DomainAlias, true
}

// HasDomainAlias returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasDomainAlias() bool {
	if o != nil && !IsNil(o.DomainAlias) {
		return true
	}

	return false
}

// SetDomainAlias gets a reference to the given []PJSIPConfigurationOption and assigns it to the DomainAlias field.
func (o *PJSIPConfigurationOptions) SetDomainAlias(v []PJSIPConfigurationOption) {
	o.DomainAlias = v
}

// GetEndpoint returns the Endpoint field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetEndpoint() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Endpoint) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Endpoint
}

// GetEndpointOk returns a tuple with the Endpoint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetEndpointOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Endpoint) {
		return nil, false
	}
	return o.Endpoint, true
}

// HasEndpoint returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasEndpoint() bool {
	if o != nil && !IsNil(o.Endpoint) {
		return true
	}

	return false
}

// SetEndpoint gets a reference to the given []PJSIPConfigurationOption and assigns it to the Endpoint field.
func (o *PJSIPConfigurationOptions) SetEndpoint(v []PJSIPConfigurationOption) {
	o.Endpoint = v
}

// GetGlobal returns the Global field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetGlobal() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Global) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Global
}

// GetGlobalOk returns a tuple with the Global field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetGlobalOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Global) {
		return nil, false
	}
	return o.Global, true
}

// HasGlobal returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasGlobal() bool {
	if o != nil && !IsNil(o.Global) {
		return true
	}

	return false
}

// SetGlobal gets a reference to the given []PJSIPConfigurationOption and assigns it to the Global field.
func (o *PJSIPConfigurationOptions) SetGlobal(v []PJSIPConfigurationOption) {
	o.Global = v
}

// GetRegistration returns the Registration field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetRegistration() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Registration) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Registration
}

// GetRegistrationOk returns a tuple with the Registration field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetRegistrationOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Registration) {
		return nil, false
	}
	return o.Registration, true
}

// HasRegistration returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasRegistration() bool {
	if o != nil && !IsNil(o.Registration) {
		return true
	}

	return false
}

// SetRegistration gets a reference to the given []PJSIPConfigurationOption and assigns it to the Registration field.
func (o *PJSIPConfigurationOptions) SetRegistration(v []PJSIPConfigurationOption) {
	o.Registration = v
}

// GetSystem returns the System field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetSystem() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.System) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.System
}

// GetSystemOk returns a tuple with the System field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetSystemOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.System) {
		return nil, false
	}
	return o.System, true
}

// HasSystem returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasSystem() bool {
	if o != nil && !IsNil(o.System) {
		return true
	}

	return false
}

// SetSystem gets a reference to the given []PJSIPConfigurationOption and assigns it to the System field.
func (o *PJSIPConfigurationOptions) SetSystem(v []PJSIPConfigurationOption) {
	o.System = v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *PJSIPConfigurationOptions) GetTransport() []PJSIPConfigurationOption {
	if o == nil || IsNil(o.Transport) {
		var ret []PJSIPConfigurationOption
		return ret
	}
	return o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PJSIPConfigurationOptions) GetTransportOk() ([]PJSIPConfigurationOption, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *PJSIPConfigurationOptions) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given []PJSIPConfigurationOption and assigns it to the Transport field.
func (o *PJSIPConfigurationOptions) SetTransport(v []PJSIPConfigurationOption) {
	o.Transport = v
}

func (o PJSIPConfigurationOptions) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PJSIPConfigurationOptions) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Aor) {
		toSerialize["aor"] = o.Aor
	}
	if !IsNil(o.Auth) {
		toSerialize["auth"] = o.Auth
	}
	if !IsNil(o.Contact) {
		toSerialize["contact"] = o.Contact
	}
	if !IsNil(o.DomainAlias) {
		toSerialize["domain_alias"] = o.DomainAlias
	}
	if !IsNil(o.Endpoint) {
		toSerialize["endpoint"] = o.Endpoint
	}
	if !IsNil(o.Global) {
		toSerialize["global"] = o.Global
	}
	if !IsNil(o.Registration) {
		toSerialize["registration"] = o.Registration
	}
	if !IsNil(o.System) {
		toSerialize["system"] = o.System
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	return toSerialize, nil
}

type NullablePJSIPConfigurationOptions struct {
	value *PJSIPConfigurationOptions
	isSet bool
}

func (v NullablePJSIPConfigurationOptions) Get() *PJSIPConfigurationOptions {
	return v.value
}

func (v *NullablePJSIPConfigurationOptions) Set(val *PJSIPConfigurationOptions) {
	v.value = val
	v.isSet = true
}

func (v NullablePJSIPConfigurationOptions) IsSet() bool {
	return v.isSet
}

func (v *NullablePJSIPConfigurationOptions) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePJSIPConfigurationOptions(val *PJSIPConfigurationOptions) *NullablePJSIPConfigurationOptions {
	return &NullablePJSIPConfigurationOptions{value: val, isSet: true}
}

func (v NullablePJSIPConfigurationOptions) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePJSIPConfigurationOptions) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
