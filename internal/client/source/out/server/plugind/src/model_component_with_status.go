/*
 * accent-plugind
 *
 * Accent's plugin management service
 *
 * API version: 0.2
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package plugindserver

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
