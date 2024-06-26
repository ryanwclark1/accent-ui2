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

// checks if the FuncKeyDestinationParameter type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FuncKeyDestinationParameter{}

// FuncKeyDestinationParameter struct for FuncKeyDestinationParameter
type FuncKeyDestinationParameter struct {
	// URL towards a collection of entities representing possible values as a destination
	Collection *string `json:"collection,omitempty"`
	// Parameter name
	Name *string `json:"name,omitempty"`
	// Array of values to choose from for this parameter
	Values []string `json:"values,omitempty"`
}

// NewFuncKeyDestinationParameter instantiates a new FuncKeyDestinationParameter object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFuncKeyDestinationParameter() *FuncKeyDestinationParameter {
	this := FuncKeyDestinationParameter{}
	return &this
}

// NewFuncKeyDestinationParameterWithDefaults instantiates a new FuncKeyDestinationParameter object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFuncKeyDestinationParameterWithDefaults() *FuncKeyDestinationParameter {
	this := FuncKeyDestinationParameter{}
	return &this
}

// GetCollection returns the Collection field value if set, zero value otherwise.
func (o *FuncKeyDestinationParameter) GetCollection() string {
	if o == nil || IsNil(o.Collection) {
		var ret string
		return ret
	}
	return *o.Collection
}

// GetCollectionOk returns a tuple with the Collection field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKeyDestinationParameter) GetCollectionOk() (*string, bool) {
	if o == nil || IsNil(o.Collection) {
		return nil, false
	}
	return o.Collection, true
}

// HasCollection returns a boolean if a field has been set.
func (o *FuncKeyDestinationParameter) HasCollection() bool {
	if o != nil && !IsNil(o.Collection) {
		return true
	}

	return false
}

// SetCollection gets a reference to the given string and assigns it to the Collection field.
func (o *FuncKeyDestinationParameter) SetCollection(v string) {
	o.Collection = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *FuncKeyDestinationParameter) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKeyDestinationParameter) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *FuncKeyDestinationParameter) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *FuncKeyDestinationParameter) SetName(v string) {
	o.Name = &v
}

// GetValues returns the Values field value if set, zero value otherwise.
func (o *FuncKeyDestinationParameter) GetValues() []string {
	if o == nil || IsNil(o.Values) {
		var ret []string
		return ret
	}
	return o.Values
}

// GetValuesOk returns a tuple with the Values field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FuncKeyDestinationParameter) GetValuesOk() ([]string, bool) {
	if o == nil || IsNil(o.Values) {
		return nil, false
	}
	return o.Values, true
}

// HasValues returns a boolean if a field has been set.
func (o *FuncKeyDestinationParameter) HasValues() bool {
	if o != nil && !IsNil(o.Values) {
		return true
	}

	return false
}

// SetValues gets a reference to the given []string and assigns it to the Values field.
func (o *FuncKeyDestinationParameter) SetValues(v []string) {
	o.Values = v
}

func (o FuncKeyDestinationParameter) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FuncKeyDestinationParameter) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Collection) {
		toSerialize["collection"] = o.Collection
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Values) {
		toSerialize["values"] = o.Values
	}
	return toSerialize, nil
}

type NullableFuncKeyDestinationParameter struct {
	value *FuncKeyDestinationParameter
	isSet bool
}

func (v NullableFuncKeyDestinationParameter) Get() *FuncKeyDestinationParameter {
	return v.value
}

func (v *NullableFuncKeyDestinationParameter) Set(val *FuncKeyDestinationParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableFuncKeyDestinationParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableFuncKeyDestinationParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFuncKeyDestinationParameter(val *FuncKeyDestinationParameter) *NullableFuncKeyDestinationParameter {
	return &NullableFuncKeyDestinationParameter{value: val, isSet: true}
}

func (v NullableFuncKeyDestinationParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFuncKeyDestinationParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
