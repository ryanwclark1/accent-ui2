/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

type ConfigPatchItem struct {

	// Patch operation. Supported operations: `replace`.
	Op string `json:"op,omitempty"`

	// JSON path to operate on. Supported paths: `/debug`.
	Path string `json:"path,omitempty"`

	// The new value for the operation. Type of value is dependent of `path`
	Value map[string]interface{} `json:"value,omitempty"`
}

// AssertConfigPatchItemRequired checks if the required fields are not zero-ed
func AssertConfigPatchItemRequired(obj ConfigPatchItem) error {
	return nil
}

// AssertConfigPatchItemConstraints checks if the values respects the defined constraints
func AssertConfigPatchItemConstraints(obj ConfigPatchItem) error {
	return nil
}
