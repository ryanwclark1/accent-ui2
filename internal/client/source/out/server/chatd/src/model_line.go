/*
 * accent-chatd
 *
 * Control your message and presence from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package chatdserver

type Line struct {
	Id int32 `json:"id,omitempty"`

	// The current state of the line.
	State string `json:"state,omitempty"`
}

// AssertLineRequired checks if the required fields are not zero-ed
func AssertLineRequired(obj Line) error {
	return nil
}

// AssertLineConstraints checks if the values respects the defined constraints
func AssertLineConstraints(obj Line) error {
	return nil
}