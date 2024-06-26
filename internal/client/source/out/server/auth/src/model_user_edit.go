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

type UserEdit struct {
	Enabled bool `json:"enabled,omitempty"`

	// The user's firstname
	Firstname string `json:"firstname,omitempty"`

	// The user's lastname
	Lastname string `json:"lastname,omitempty"`

	Purpose string `json:"purpose,omitempty"`

	// The username that will identify that new username
	Username string `json:"username,omitempty"`
}

// AssertUserEditRequired checks if the required fields are not zero-ed
func AssertUserEditRequired(obj UserEdit) error {
	return nil
}

// AssertUserEditConstraints checks if the values respects the defined constraints
func AssertUserEditConstraints(obj UserEdit) error {
	return nil
}
