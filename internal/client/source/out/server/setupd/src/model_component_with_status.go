/*
 * accent-setupd
 *
 * Initialize Accent Engine from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package setupdserver

type ComponentWithStatus struct {
	Status StatusValue `json:"status,omitempty"`
}

// AssertComponentWithStatusRequired checks if the required fields are not zero-ed
func AssertComponentWithStatusRequired(obj ComponentWithStatus) error {
	return nil
}

// AssertComponentWithStatusConstraints checks if the values respects the defined constraints
func AssertComponentWithStatusConstraints(obj ComponentWithStatus) error {
	return nil
}
