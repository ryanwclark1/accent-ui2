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

type SessionResult struct {
	Mobile bool `json:"mobile,omitempty"`

	TenantUuid string `json:"tenant_uuid,omitempty"`

	UserUuid string `json:"user_uuid,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// AssertSessionResultRequired checks if the required fields are not zero-ed
func AssertSessionResultRequired(obj SessionResult) error {
	return nil
}

// AssertSessionResultConstraints checks if the values respects the defined constraints
func AssertSessionResultConstraints(obj SessionResult) error {
	return nil
}