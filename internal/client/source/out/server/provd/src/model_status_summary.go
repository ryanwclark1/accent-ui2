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

type StatusSummary struct {
	RestApi ComponentWithStatus `json:"rest_api,omitempty"`
}

// AssertStatusSummaryRequired checks if the required fields are not zero-ed
func AssertStatusSummaryRequired(obj StatusSummary) error {
	if err := AssertComponentWithStatusRequired(obj.RestApi); err != nil {
		return err
	}
	return nil
}

// AssertStatusSummaryConstraints checks if the values respects the defined constraints
func AssertStatusSummaryConstraints(obj StatusSummary) error {
	return nil
}
