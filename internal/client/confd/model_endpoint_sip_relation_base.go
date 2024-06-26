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

// checks if the EndpointSipRelationBase type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &EndpointSipRelationBase{}

// EndpointSipRelationBase struct for EndpointSipRelationBase
type EndpointSipRelationBase struct {
	// A list of PJSIP auth section options for this endpoint. Only `username` is supported
	AuthSectionOptions [][]string `json:"auth_section_options,omitempty"`
	// The human readable name for this configuration
	Label *string `json:"label,omitempty"`
	// The name of the PJSIP entity, auto-generated if missing
	Name *string `json:"name,omitempty"`
	// The UUID of this resource
	Uuid *string `json:"uuid,omitempty"`
}

// NewEndpointSipRelationBase instantiates a new EndpointSipRelationBase object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEndpointSipRelationBase() *EndpointSipRelationBase {
	this := EndpointSipRelationBase{}
	return &this
}

// NewEndpointSipRelationBaseWithDefaults instantiates a new EndpointSipRelationBase object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEndpointSipRelationBaseWithDefaults() *EndpointSipRelationBase {
	this := EndpointSipRelationBase{}
	return &this
}

// GetAuthSectionOptions returns the AuthSectionOptions field value if set, zero value otherwise.
func (o *EndpointSipRelationBase) GetAuthSectionOptions() [][]string {
	if o == nil || IsNil(o.AuthSectionOptions) {
		var ret [][]string
		return ret
	}
	return o.AuthSectionOptions
}

// GetAuthSectionOptionsOk returns a tuple with the AuthSectionOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSipRelationBase) GetAuthSectionOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.AuthSectionOptions) {
		return nil, false
	}
	return o.AuthSectionOptions, true
}

// HasAuthSectionOptions returns a boolean if a field has been set.
func (o *EndpointSipRelationBase) HasAuthSectionOptions() bool {
	if o != nil && !IsNil(o.AuthSectionOptions) {
		return true
	}

	return false
}

// SetAuthSectionOptions gets a reference to the given [][]string and assigns it to the AuthSectionOptions field.
func (o *EndpointSipRelationBase) SetAuthSectionOptions(v [][]string) {
	o.AuthSectionOptions = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *EndpointSipRelationBase) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSipRelationBase) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *EndpointSipRelationBase) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *EndpointSipRelationBase) SetLabel(v string) {
	o.Label = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *EndpointSipRelationBase) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSipRelationBase) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *EndpointSipRelationBase) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *EndpointSipRelationBase) SetName(v string) {
	o.Name = &v
}

// GetUuid returns the Uuid field value if set, zero value otherwise.
func (o *EndpointSipRelationBase) GetUuid() string {
	if o == nil || IsNil(o.Uuid) {
		var ret string
		return ret
	}
	return *o.Uuid
}

// GetUuidOk returns a tuple with the Uuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *EndpointSipRelationBase) GetUuidOk() (*string, bool) {
	if o == nil || IsNil(o.Uuid) {
		return nil, false
	}
	return o.Uuid, true
}

// HasUuid returns a boolean if a field has been set.
func (o *EndpointSipRelationBase) HasUuid() bool {
	if o != nil && !IsNil(o.Uuid) {
		return true
	}

	return false
}

// SetUuid gets a reference to the given string and assigns it to the Uuid field.
func (o *EndpointSipRelationBase) SetUuid(v string) {
	o.Uuid = &v
}

func (o EndpointSipRelationBase) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o EndpointSipRelationBase) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AuthSectionOptions) {
		toSerialize["auth_section_options"] = o.AuthSectionOptions
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Uuid) {
		toSerialize["uuid"] = o.Uuid
	}
	return toSerialize, nil
}

type NullableEndpointSipRelationBase struct {
	value *EndpointSipRelationBase
	isSet bool
}

func (v NullableEndpointSipRelationBase) Get() *EndpointSipRelationBase {
	return v.value
}

func (v *NullableEndpointSipRelationBase) Set(val *EndpointSipRelationBase) {
	v.value = val
	v.isSet = true
}

func (v NullableEndpointSipRelationBase) IsSet() bool {
	return v.isSet
}

func (v *NullableEndpointSipRelationBase) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEndpointSipRelationBase(val *EndpointSipRelationBase) *NullableEndpointSipRelationBase {
	return &NullableEndpointSipRelationBase{value: val, isSet: true}
}

func (v NullableEndpointSipRelationBase) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEndpointSipRelationBase) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
