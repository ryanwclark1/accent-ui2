/*
accent-agentd

Agent service

API version: 1.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package agentd

import (
	"encoding/json"
	"fmt"
)

// StatusValue the model 'StatusValue'
type StatusValue string

// List of StatusValue
const (
	STATUSVALUE_FAIL StatusValue = "fail"
	STATUSVALUE_OK   StatusValue = "ok"
)

// All allowed values of StatusValue enum
var AllowedStatusValueEnumValues = []StatusValue{
	"fail",
	"ok",
}

func (v *StatusValue) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := StatusValue(value)
	for _, existing := range AllowedStatusValueEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid StatusValue", value)
}

// NewStatusValueFromValue returns a pointer to a valid StatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusValueFromValue(v string) (*StatusValue, error) {
	ev := StatusValue(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for StatusValue: valid values are %v", v, AllowedStatusValueEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusValue) IsValid() bool {
	for _, existing := range AllowedStatusValueEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to StatusValue value
func (v StatusValue) Ptr() *StatusValue {
	return &v
}

type NullableStatusValue struct {
	value *StatusValue
	isSet bool
}

func (v NullableStatusValue) Get() *StatusValue {
	return v.value
}

func (v *NullableStatusValue) Set(val *StatusValue) {
	v.value = val
	v.isSet = true
}

func (v NullableStatusValue) IsSet() bool {
	return v.isSet
}

func (v *NullableStatusValue) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatusValue(val *StatusValue) *NullableStatusValue {
	return &NullableStatusValue{value: val, isSet: true}
}

func (v NullableStatusValue) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatusValue) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
