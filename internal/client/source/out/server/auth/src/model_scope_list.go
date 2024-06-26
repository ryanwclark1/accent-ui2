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

type ScopeList struct {

	// the scopes and their check result
	Scopes []map[string]interface{} `json:"scopes,omitempty"`
}

// AssertScopeListRequired checks if the required fields are not zero-ed
func AssertScopeListRequired(obj ScopeList) error {
	return nil
}

// AssertScopeListConstraints checks if the values respects the defined constraints
func AssertScopeListConstraints(obj ScopeList) error {
	return nil
}
