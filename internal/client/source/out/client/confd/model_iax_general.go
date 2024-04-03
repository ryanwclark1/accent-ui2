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

// checks if the IAXGeneral type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IAXGeneral{}

// IAXGeneral IAX general configuration. This configuration will be in the file 'iax.conf' used by asterisk. Please consult the asterisk documentation for further details on available parameters.
type IAXGeneral struct {
	// General IAX options. These options must be unique and unordered. Otherwise, use `ordered_options`. Option must have the following form:  ``` {   \"options\": {     \"name1\": \"value1\",     ...   } } ```
	Options map[string]interface{} `json:"options,omitempty"`
	// Any options may be repeated as many times and ordered as needed. Ordered options must have the following form:  ``` {   \"ordered_options\": [     [\"name1\", \"value1\"],     [\"name2\", \"value2\"]   ] } ```  The resulting configuration in iax.conf will have the following form:  ``` [general] name1=value1 name2=value2 ```
	OrderedOptions [][]string `json:"ordered_options,omitempty"`
}

// NewIAXGeneral instantiates a new IAXGeneral object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIAXGeneral() *IAXGeneral {
	this := IAXGeneral{}
	return &this
}

// NewIAXGeneralWithDefaults instantiates a new IAXGeneral object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIAXGeneralWithDefaults() *IAXGeneral {
	this := IAXGeneral{}
	return &this
}

// GetOptions returns the Options field value if set, zero value otherwise.
func (o *IAXGeneral) GetOptions() map[string]interface{} {
	if o == nil || IsNil(o.Options) {
		var ret map[string]interface{}
		return ret
	}
	return o.Options
}

// GetOptionsOk returns a tuple with the Options field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAXGeneral) GetOptionsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Options) {
		return map[string]interface{}{}, false
	}
	return o.Options, true
}

// HasOptions returns a boolean if a field has been set.
func (o *IAXGeneral) HasOptions() bool {
	if o != nil && !IsNil(o.Options) {
		return true
	}

	return false
}

// SetOptions gets a reference to the given map[string]interface{} and assigns it to the Options field.
func (o *IAXGeneral) SetOptions(v map[string]interface{}) {
	o.Options = v
}

// GetOrderedOptions returns the OrderedOptions field value if set, zero value otherwise.
func (o *IAXGeneral) GetOrderedOptions() [][]string {
	if o == nil || IsNil(o.OrderedOptions) {
		var ret [][]string
		return ret
	}
	return o.OrderedOptions
}

// GetOrderedOptionsOk returns a tuple with the OrderedOptions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IAXGeneral) GetOrderedOptionsOk() ([][]string, bool) {
	if o == nil || IsNil(o.OrderedOptions) {
		return nil, false
	}
	return o.OrderedOptions, true
}

// HasOrderedOptions returns a boolean if a field has been set.
func (o *IAXGeneral) HasOrderedOptions() bool {
	if o != nil && !IsNil(o.OrderedOptions) {
		return true
	}

	return false
}

// SetOrderedOptions gets a reference to the given [][]string and assigns it to the OrderedOptions field.
func (o *IAXGeneral) SetOrderedOptions(v [][]string) {
	o.OrderedOptions = v
}

func (o IAXGeneral) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IAXGeneral) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Options) {
		toSerialize["options"] = o.Options
	}
	if !IsNil(o.OrderedOptions) {
		toSerialize["ordered_options"] = o.OrderedOptions
	}
	return toSerialize, nil
}

type NullableIAXGeneral struct {
	value *IAXGeneral
	isSet bool
}

func (v NullableIAXGeneral) Get() *IAXGeneral {
	return v.value
}

func (v *NullableIAXGeneral) Set(val *IAXGeneral) {
	v.value = val
	v.isSet = true
}

func (v NullableIAXGeneral) IsSet() bool {
	return v.isSet
}

func (v *NullableIAXGeneral) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIAXGeneral(val *IAXGeneral) *NullableIAXGeneral {
	return &NullableIAXGeneral{value: val, isSet: true}
}

func (v NullableIAXGeneral) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIAXGeneral) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
