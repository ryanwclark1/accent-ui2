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

type RelocateList struct {
	Items Relocate `json:"items,omitempty"`
}

// AssertRelocateListRequired checks if the required fields are not zero-ed
func AssertRelocateListRequired(obj RelocateList) error {
	if err := AssertRelocateRequired(obj.Items); err != nil {
		return err
	}
	return nil
}

// AssertRelocateListConstraints checks if the values respects the defined constraints
func AssertRelocateListConstraints(obj RelocateList) error {
	return nil
}
