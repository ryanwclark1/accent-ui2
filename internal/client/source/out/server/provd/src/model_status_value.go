/*
 * accent-provd
 *
 * Provisioning application REST API
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package provdserver

import (
	"fmt"
)

type StatusValue string

// List of StatusValue
const (
	STATUSVALUE_FAIL StatusValue = "fail"
	STATUSVALUE_OK   StatusValue = "ok"
)

// AllowedStatusValueEnumValues is all the allowed values of StatusValue enum
var AllowedStatusValueEnumValues = []StatusValue{
	"fail",
	"ok",
}

// validStatusValueEnumValue provides a map of StatusValues for fast verification of use input
var validStatusValueEnumValues = map[StatusValue]struct{}{
	"fail": {},
	"ok":   {},
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v StatusValue) IsValid() bool {
	_, ok := validStatusValueEnumValues[v]
	return ok
}

// NewStatusValueFromValue returns a pointer to a valid StatusValue
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewStatusValueFromValue(v string) (StatusValue, error) {
	ev := StatusValue(v)
	if ev.IsValid() {
		return ev, nil
	} else {
		return "", fmt.Errorf("invalid value '%v' for StatusValue: valid values are %v", v, AllowedStatusValueEnumValues)
	}
}

// AssertStatusValueRequired checks if the required fields are not zero-ed
func AssertStatusValueRequired(obj StatusValue) error {
	return nil
}

// AssertStatusValueConstraints checks if the values respects the defined constraints
func AssertStatusValueConstraints(obj StatusValue) error {
	return nil
}