/*
 * accent-auth
 *
 * Accent's authentication service
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package authserver

type GroupResult struct {
	Name string `json:"name,omitempty"`

	ReadOnly bool `json:"read_only,omitempty"`

	Slug string `json:"slug,omitempty"`

	// *Deprecated* Please use `read_only`
	SystemManaged bool `json:"system_managed,omitempty"`

	TenantUuid string `json:"tenant_uuid,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// AssertGroupResultRequired checks if the required fields are not zero-ed
func AssertGroupResultRequired(obj GroupResult) error {
	return nil
}

// AssertGroupResultConstraints checks if the values respects the defined constraints
func AssertGroupResultConstraints(obj GroupResult) error {
	return nil
}
