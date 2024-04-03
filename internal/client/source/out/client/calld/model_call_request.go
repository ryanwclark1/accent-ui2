/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the CallRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &CallRequest{}

// CallRequest struct for CallRequest
type CallRequest struct {
	Destination CallRequestDestination `json:"destination"`
	Source      CallRequestSource      `json:"source"`
	// Channel variables to set
	Variables map[string]interface{} `json:"variables,omitempty"`
}

type _CallRequest CallRequest

// NewCallRequest instantiates a new CallRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCallRequest(destination CallRequestDestination, source CallRequestSource) *CallRequest {
	this := CallRequest{}
	this.Destination = destination
	this.Source = source
	return &this
}

// NewCallRequestWithDefaults instantiates a new CallRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallRequestWithDefaults() *CallRequest {
	this := CallRequest{}
	return &this
}

// GetDestination returns the Destination field value
func (o *CallRequest) GetDestination() CallRequestDestination {
	if o == nil {
		var ret CallRequestDestination
		return ret
	}

	return o.Destination
}

// GetDestinationOk returns a tuple with the Destination field value
// and a boolean to check if the value has been set.
func (o *CallRequest) GetDestinationOk() (*CallRequestDestination, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Destination, true
}

// SetDestination sets field value
func (o *CallRequest) SetDestination(v CallRequestDestination) {
	o.Destination = v
}

// GetSource returns the Source field value
func (o *CallRequest) GetSource() CallRequestSource {
	if o == nil {
		var ret CallRequestSource
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *CallRequest) GetSourceOk() (*CallRequestSource, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *CallRequest) SetSource(v CallRequestSource) {
	o.Source = v
}

// GetVariables returns the Variables field value if set, zero value otherwise.
func (o *CallRequest) GetVariables() map[string]interface{} {
	if o == nil || IsNil(o.Variables) {
		var ret map[string]interface{}
		return ret
	}
	return o.Variables
}

// GetVariablesOk returns a tuple with the Variables field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CallRequest) GetVariablesOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Variables) {
		return map[string]interface{}{}, false
	}
	return o.Variables, true
}

// HasVariables returns a boolean if a field has been set.
func (o *CallRequest) HasVariables() bool {
	if o != nil && !IsNil(o.Variables) {
		return true
	}

	return false
}

// SetVariables gets a reference to the given map[string]interface{} and assigns it to the Variables field.
func (o *CallRequest) SetVariables(v map[string]interface{}) {
	o.Variables = v
}

func (o CallRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o CallRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["destination"] = o.Destination
	toSerialize["source"] = o.Source
	if !IsNil(o.Variables) {
		toSerialize["variables"] = o.Variables
	}
	return toSerialize, nil
}

func (o *CallRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"destination",
		"source",
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

	varCallRequest := _CallRequest{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varCallRequest)

	if err != nil {
		return err
	}

	*o = CallRequest(varCallRequest)

	return err
}

type NullableCallRequest struct {
	value *CallRequest
	isSet bool
}

func (v NullableCallRequest) Get() *CallRequest {
	return v.value
}

func (v *NullableCallRequest) Set(val *CallRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableCallRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableCallRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCallRequest(val *CallRequest) *NullableCallRequest {
	return &NullableCallRequest{value: val, isSet: true}
}

func (v NullableCallRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCallRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
