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

// checks if the OutcallRelationExtension type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &OutcallRelationExtension{}

// OutcallRelationExtension struct for OutcallRelationExtension
type OutcallRelationExtension struct {
	Context *string `json:"context,omitempty"`
	Exten   *string `json:"exten,omitempty"`
	// Extension ID
	Id       *int32  `json:"id,omitempty"`
	CallerId *string `json:"caller_id,omitempty"`
	// The prefix added to the outgoing call when sent on the trunk
	ExternalPrefix *string `json:"external_prefix,omitempty"`
	// The prefix that the user need to dial before the extension
	Prefix *string `json:"prefix,omitempty"`
	// The number of leading digits to strip
	StripDigits *int32 `json:"strip_digits,omitempty"`
}

// NewOutcallRelationExtension instantiates a new OutcallRelationExtension object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutcallRelationExtension() *OutcallRelationExtension {
	this := OutcallRelationExtension{}
	return &this
}

// NewOutcallRelationExtensionWithDefaults instantiates a new OutcallRelationExtension object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutcallRelationExtensionWithDefaults() *OutcallRelationExtension {
	this := OutcallRelationExtension{}
	return &this
}

// GetContext returns the Context field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetContext() string {
	if o == nil || IsNil(o.Context) {
		var ret string
		return ret
	}
	return *o.Context
}

// GetContextOk returns a tuple with the Context field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetContextOk() (*string, bool) {
	if o == nil || IsNil(o.Context) {
		return nil, false
	}
	return o.Context, true
}

// HasContext returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasContext() bool {
	if o != nil && !IsNil(o.Context) {
		return true
	}

	return false
}

// SetContext gets a reference to the given string and assigns it to the Context field.
func (o *OutcallRelationExtension) SetContext(v string) {
	o.Context = &v
}

// GetExten returns the Exten field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetExten() string {
	if o == nil || IsNil(o.Exten) {
		var ret string
		return ret
	}
	return *o.Exten
}

// GetExtenOk returns a tuple with the Exten field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetExtenOk() (*string, bool) {
	if o == nil || IsNil(o.Exten) {
		return nil, false
	}
	return o.Exten, true
}

// HasExten returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasExten() bool {
	if o != nil && !IsNil(o.Exten) {
		return true
	}

	return false
}

// SetExten gets a reference to the given string and assigns it to the Exten field.
func (o *OutcallRelationExtension) SetExten(v string) {
	o.Exten = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *OutcallRelationExtension) SetId(v int32) {
	o.Id = &v
}

// GetCallerId returns the CallerId field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetCallerId() string {
	if o == nil || IsNil(o.CallerId) {
		var ret string
		return ret
	}
	return *o.CallerId
}

// GetCallerIdOk returns a tuple with the CallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetCallerIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallerId) {
		return nil, false
	}
	return o.CallerId, true
}

// HasCallerId returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasCallerId() bool {
	if o != nil && !IsNil(o.CallerId) {
		return true
	}

	return false
}

// SetCallerId gets a reference to the given string and assigns it to the CallerId field.
func (o *OutcallRelationExtension) SetCallerId(v string) {
	o.CallerId = &v
}

// GetExternalPrefix returns the ExternalPrefix field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetExternalPrefix() string {
	if o == nil || IsNil(o.ExternalPrefix) {
		var ret string
		return ret
	}
	return *o.ExternalPrefix
}

// GetExternalPrefixOk returns a tuple with the ExternalPrefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetExternalPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.ExternalPrefix) {
		return nil, false
	}
	return o.ExternalPrefix, true
}

// HasExternalPrefix returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasExternalPrefix() bool {
	if o != nil && !IsNil(o.ExternalPrefix) {
		return true
	}

	return false
}

// SetExternalPrefix gets a reference to the given string and assigns it to the ExternalPrefix field.
func (o *OutcallRelationExtension) SetExternalPrefix(v string) {
	o.ExternalPrefix = &v
}

// GetPrefix returns the Prefix field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetPrefix() string {
	if o == nil || IsNil(o.Prefix) {
		var ret string
		return ret
	}
	return *o.Prefix
}

// GetPrefixOk returns a tuple with the Prefix field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetPrefixOk() (*string, bool) {
	if o == nil || IsNil(o.Prefix) {
		return nil, false
	}
	return o.Prefix, true
}

// HasPrefix returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasPrefix() bool {
	if o != nil && !IsNil(o.Prefix) {
		return true
	}

	return false
}

// SetPrefix gets a reference to the given string and assigns it to the Prefix field.
func (o *OutcallRelationExtension) SetPrefix(v string) {
	o.Prefix = &v
}

// GetStripDigits returns the StripDigits field value if set, zero value otherwise.
func (o *OutcallRelationExtension) GetStripDigits() int32 {
	if o == nil || IsNil(o.StripDigits) {
		var ret int32
		return ret
	}
	return *o.StripDigits
}

// GetStripDigitsOk returns a tuple with the StripDigits field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *OutcallRelationExtension) GetStripDigitsOk() (*int32, bool) {
	if o == nil || IsNil(o.StripDigits) {
		return nil, false
	}
	return o.StripDigits, true
}

// HasStripDigits returns a boolean if a field has been set.
func (o *OutcallRelationExtension) HasStripDigits() bool {
	if o != nil && !IsNil(o.StripDigits) {
		return true
	}

	return false
}

// SetStripDigits gets a reference to the given int32 and assigns it to the StripDigits field.
func (o *OutcallRelationExtension) SetStripDigits(v int32) {
	o.StripDigits = &v
}

func (o OutcallRelationExtension) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o OutcallRelationExtension) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Context) {
		toSerialize["context"] = o.Context
	}
	if !IsNil(o.Exten) {
		toSerialize["exten"] = o.Exten
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.CallerId) {
		toSerialize["caller_id"] = o.CallerId
	}
	if !IsNil(o.ExternalPrefix) {
		toSerialize["external_prefix"] = o.ExternalPrefix
	}
	if !IsNil(o.Prefix) {
		toSerialize["prefix"] = o.Prefix
	}
	if !IsNil(o.StripDigits) {
		toSerialize["strip_digits"] = o.StripDigits
	}
	return toSerialize, nil
}

type NullableOutcallRelationExtension struct {
	value *OutcallRelationExtension
	isSet bool
}

func (v NullableOutcallRelationExtension) Get() *OutcallRelationExtension {
	return v.value
}

func (v *NullableOutcallRelationExtension) Set(val *OutcallRelationExtension) {
	v.value = val
	v.isSet = true
}

func (v NullableOutcallRelationExtension) IsSet() bool {
	return v.isSet
}

func (v *NullableOutcallRelationExtension) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutcallRelationExtension(val *OutcallRelationExtension) *NullableOutcallRelationExtension {
	return &NullableOutcallRelationExtension{value: val, isSet: true}
}

func (v NullableOutcallRelationExtension) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutcallRelationExtension) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
