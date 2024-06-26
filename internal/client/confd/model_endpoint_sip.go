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

// checks if the EndpointSIP type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointSIP{}

// EndpointSIP struct for EndpointSIP
type EndpointSIP struct {
	// A list of PJSIP AOR section options for this endpoint
	AorSectionOptions [][]string `json:"aor_section_options,omitempty"`
	// The ID of the Asterisk onto which this configuration applies.
	AsteriskId *string `json:"asterisk_id,omitempty"`
	// A list of PJSIP auth section options for this endpoint
	AuthSectionOptions [][]string `json:"auth_section_options,omitempty"`
	// A list of PJSIP endpoint section options for this endpoint
	EndpointSectionOptions [][]string `json:"endpoint_section_options,omitempty"`
	// A list of PJSIP identify section options for this endpoint
	IdentifySectionOptions [][]string `json:"identify_section_options,omitempty"`
	// The human readable name for this configuration
	Label *string `json:"label,omitempty"`
	// The name of the PJSIP entity, auto-generated if missing
	Name *string `json:"name,omitempty"`
	// A list of PJSIP auth section options for this endpoint
	OutboundAuthSectionOptions [][]string `json:"outbound_auth_section_options,omitempty"`
	// A list of PJSIP auth section options for this endpoint
	RegistrationOutboundAuthSectionOptions [][]string `json:"registration_outbound_auth_section_options,omitempty"`
	// A list of PJSIP registration section options for this endpoint
	RegistrationSectionOptions [][]string `json:"registration_section_options,omitempty"`
	// A list of templates this configuration will inherit from.  The inheritance only applies to option sections. Not to the name, label and other fields. For a given list of templates [A, B, C] A will be applied over B. C will be applied over (A,B) etc.
	Templates []EndpointSIPTemplatesRelation `json:"templates,omitempty"`
	// The UUID of the tenant
	TenantUuid *string               `json:"tenant_uuid,omitempty"`
	Transport  *SIPTransportRelation `json:"transport,omitempty"`
	// The UUID of this resource
	Uuid *string `json:"uuid,omitempty"`
}

// NewEndpointSIP instantiates a new EndpointSIP object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointSIP() *EndpointSIP {
	this := EndpointSIP{}
	return &this
}

// NewEndpointSIPWithDefaults instantiates a new EndpointSIP object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointSIPWithDefaults() *EndpointSIP {
	this := EndpointSIP{}
	return &this
}

// GetAorSectionOptions returns the AorSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetAorSectionOptions() [][]string {
	if o == nil || IsNil(o.AorSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.AorSectionOptions
}

// GetAorSectionOptionsOk returns a tuple with the AorSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetAorSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.AorSectionOptions) {
		return nil, false
	}
	return o.AorSectionOptions, true
}

// HasAorSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasAorSectionOptions() bool {
	if o != nil && !IsNil(o.AorSectionOptions) {
		return true
	}

	return false
}

// SetAorSectionOptions gets a reference to the given [][]string and assigns it to the AorSectionOptions field.
func (o *EndpointSIP) SetAorSectionOptions(v [][]string) {
	o.AorSectionOptions = v
}

// GetAsteriskId returns the AsteriskId field value if set, zero value otherwise.
func (o *EndpointSIP) GetAsteriskId() string {
	if o == nil || IsNil(o.AsteriskId) {
		var ret string
		return ret
	}
	return *o.AsteriskId
}

// GetAsteriskIdOk returns a tuple with the AsteriskId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetAsteriskIdOk() (*string, bool) {
	if o == nil || IsNil(o.AsteriskId) {
		return nil, false
	}
	return o.AsteriskId, true
}

// HasAsteriskId returns a boolean if a field has been set.
func (o *EndpointSIP) HasAsteriskId() bool {
	if o != nil && !IsNil(o.AsteriskId) {
		return true
	}

	return false
}

// SetAsteriskId gets a reference to the given string and assigns it to the AsteriskId field.
func (o *EndpointSIP) SetAsteriskId(v string) {
	o.AsteriskId = &v
}

// GetAuthSectionOptions returns the AuthSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetAuthSectionOptions() [][]string {
	if o == nil || IsNil(o.AuthSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.AuthSectionOptions
}

// GetAuthSectionOptionsOk returns a tuple with the AuthSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetAuthSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.AuthSectionOptions) {
		return nil, false
	}
	return o.AuthSectionOptions, true
}

// HasAuthSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasAuthSectionOptions() bool {
	if o != nil && !IsNil(o.AuthSectionOptions) {
		return true
	}

	return false
}

// SetAuthSectionOptions gets a reference to the given [][]string and assigns it to the AuthSectionOptions field.
func (o *EndpointSIP) SetAuthSectionOptions(v [][]string) {
	o.AuthSectionOptions = v
}

// GetEndpointSectionOptions returns the EndpointSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetEndpointSectionOptions() [][]string {
	if o == nil || IsNil(o.EndpointSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.EndpointSectionOptions
}

// GetEndpointSectionOptionsOk returns a tuple with the EndpointSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetEndpointSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.EndpointSectionOptions) {
		return nil, false
	}
	return o.EndpointSectionOptions, true
}

// HasEndpointSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasEndpointSectionOptions() bool {
	if o != nil && !IsNil(o.EndpointSectionOptions) {
		return true
	}

	return false
}

// SetEndpointSectionOptions gets a reference to the given [][]string and assigns it to the EndpointSectionOptions field.
func (o *EndpointSIP) SetEndpointSectionOptions(v [][]string) {
	o.EndpointSectionOptions = v
}

// GetIdentifySectionOptions returns the IdentifySectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetIdentifySectionOptions() [][]string {
	if o == nil || IsNil(o.IdentifySectionOptions) {
		var ret [][]string
		return ret
	}
	return o.IdentifySectionOptions
}

// GetIdentifySectionOptionsOk returns a tuple with the IdentifySectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetIdentifySectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.IdentifySectionOptions) {
		return nil, false
	}
	return o.IdentifySectionOptions, true
}

// HasIdentifySectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasIdentifySectionOptions() bool {
	if o != nil && !IsNil(o.IdentifySectionOptions) {
		return true
	}

	return false
}

// SetIdentifySectionOptions gets a reference to the given [][]string and assigns it to the IdentifySectionOptions field.
func (o *EndpointSIP) SetIdentifySectionOptions(v [][]string) {
	o.IdentifySectionOptions = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EndpointSIP) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EndpointSIP) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EndpointSIP) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointSIP) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointSIP) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointSIP) SetName(v string) {
	o.Name = &v
}

// GetOutboundAuthSectionOptions returns the OutboundAuthSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetOutboundAuthSectionOptions() [][]string {
	if o == nil || IsNil(o.OutboundAuthSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.OutboundAuthSectionOptions
}

// GetOutboundAuthSectionOptionsOk returns a tuple with the OutboundAuthSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetOutboundAuthSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.OutboundAuthSectionOptions) {
		return nil, false
	}
	return o.OutboundAuthSectionOptions, true
}

// HasOutboundAuthSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasOutboundAuthSectionOptions() bool {
	if o != nil && !IsNil(o.OutboundAuthSectionOptions) {
		return true
	}

	return false
}

// SetOutboundAuthSectionOptions gets a reference to the given [][]string and assigns it to the OutboundAuthSectionOptions field.
func (o *EndpointSIP) SetOutboundAuthSectionOptions(v [][]string) {
	o.OutboundAuthSectionOptions = v
}

// GetRegistrationOutboundAuthSectionOptions returns the RegistrationOutboundAuthSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetRegistrationOutboundAuthSectionOptions() [][]string {
	if o == nil || IsNil(o.RegistrationOutboundAuthSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.RegistrationOutboundAuthSectionOptions
}

// GetRegistrationOutboundAuthSectionOptionsOk returns a tuple with the RegistrationOutboundAuthSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetRegistrationOutboundAuthSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.RegistrationOutboundAuthSectionOptions) {
		return nil, false
	}
	return o.RegistrationOutboundAuthSectionOptions, true
}

// HasRegistrationOutboundAuthSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasRegistrationOutboundAuthSectionOptions() bool {
	if o != nil && !IsNil(o.RegistrationOutboundAuthSectionOptions) {
		return true
	}

	return false
}

// SetRegistrationOutboundAuthSectionOptions gets a reference to the given [][]string and assigns it to the RegistrationOutboundAuthSectionOptions field.
func (o *EndpointSIP) SetRegistrationOutboundAuthSectionOptions(v [][]string) {
	o.RegistrationOutboundAuthSectionOptions = v
}

// GetRegistrationSectionOptions returns the RegistrationSectionOptions field value if set, zero value otherwise.
func (o *EndpointSIP) GetRegistrationSectionOptions() [][]string {
	if o == nil || IsNil(o.RegistrationSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.RegistrationSectionOptions
}

// GetRegistrationSectionOptionsOk returns a tuple with the RegistrationSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetRegistrationSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.RegistrationSectionOptions) {
		return nil, false
	}
	return o.RegistrationSectionOptions, true
}

// HasRegistrationSectionOptions returns a boolean if a field has been set.
func (o *EndpointSIP) HasRegistrationSectionOptions() bool {
	if o != nil && !IsNil(o.RegistrationSectionOptions) {
		return true
	}

	return false
}

// SetRegistrationSectionOptions gets a reference to the given [][]string and assigns it to the RegistrationSectionOptions field.
func (o *EndpointSIP) SetRegistrationSectionOptions(v [][]string) {
	o.RegistrationSectionOptions = v
}

// GetTemplates returns the Templates field value if set, zero value otherwise.
func (o *EndpointSIP) GetTemplates() []EndpointSIPTemplatesRelation {
	if o == nil || IsNil(o.Templates) {
		var ret []EndpointSIPTemplatesRelation
		return ret
	}
	return o.Templates
}

// GetTemplatesOk returns a tuple with the Templates field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetTemplatesOk() ([]EndpointSIPTemplatesRelation, bool) {
	if o == nil || IsNil(o.Templates) {
		return nil, false
	}
	return o.Templates, true
}

// HasTemplates returns a boolean if a field has been set.
func (o *EndpointSIP) HasTemplates() bool {
	if o != nil && !IsNil(o.Templates) {
		return true
	}

	return false
}

// SetTemplates gets a reference to the given []EndpointSIPTemplatesRelation and assigns it to the Templates field.
func (o *EndpointSIP) SetTemplates(v []EndpointSIPTemplatesRelation) {
	o.Templates = v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *EndpointSIP) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *EndpointSIP) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *EndpointSIP) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetTransport returns the Transport field value if set, zero value otherwise.
func (o *EndpointSIP) GetTransport() SIPTransportRelation {
	if o == nil || IsNil(o.Transport) {
		var ret SIPTransportRelation
		return ret
	}
	return *o.Transport
}

// GetTransportOk returns a tuple with the Transport field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetTransportOk() (*SIPTransportRelation, bool) {
	if o == nil || IsNil(o.Transport) {
		return nil, false
	}
	return o.Transport, true
}

// HasTransport returns a boolean if a field has been set.
func (o *EndpointSIP) HasTransport() bool {
	if o != nil && !IsNil(o.Transport) {
		return true
	}

	return false
}

// SetTransport gets a reference to the given SIPTransportRelation and assigns it to the Transport field.
func (o *EndpointSIP) SetTransport(v SIPTransportRelation) {
	o.Transport = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *EndpointSIP) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSIP) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *EndpointSIP) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *EndpointSIP) SetUuid(v string) {
	o.Uuid = &v
}

func (o EndpointSIP) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointSIP) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AorSectionOptions) {
		toSerialize["aor_section_options"] = o.AorSectionOptions
	}
	if !IsNil(o.AsteriskId) {
		toSerialize["asterisk_id"] = o.AsteriskId
	}
	if !IsNil(o.AuthSectionOptions) {
		toSerialize["auth_section_options"] = o.AuthSectionOptions
	}
	if !IsNil(o.EndpointSectionOptions) {
		toSerialize["endpoint_section_options"] = o.EndpointSectionOptions
	}
	if !IsNil(o.IdentifySectionOptions) {
		toSerialize["identify_section_options"] = o.IdentifySectionOptions
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.OutboundAuthSectionOptions) {
		toSerialize["outbound_auth_section_options"] = o.OutboundAuthSectionOptions
	}
	if !IsNil(o.RegistrationOutboundAuthSectionOptions) {
		toSerialize["registration_outbound_auth_section_options"] = o.RegistrationOutboundAuthSectionOptions
	}
	if !IsNil(o.RegistrationSectionOptions) {
		toSerialize["registration_section_options"] = o.RegistrationSectionOptions
	}
	if !IsNil(o.Templates) {
		toSerialize["templates"] = o.Templates
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Transport) {
		toSerialize["transport"] = o.Transport
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableEndpointSIP struct {
	value *EndpointSIP
	isSet bool
}

func (v NullableEndpointSIP) Get() *EndpointSIP {
	return v.value
}

func (v *NullableEndpointSIP) Set(val *EndpointSIP) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointSIP) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointSIP) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointSIP(val *EndpointSIP) *NullableEndpointSIP {
	return &NullableEndpointSIP{value: val, isSet: true}
}

func (v NullableEndpointSIP) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointSIP) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
