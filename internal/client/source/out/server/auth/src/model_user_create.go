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

type UserCreate struct {

	// The main email address of the new username
	EmailAddress string `json:"email_address,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	// The user's firstname
	Firstname string `json:"firstname,omitempty"`

	// The user's lastname
	Lastname string `json:"lastname,omitempty"`

	// The password of the newly created username
	Password string `json:"password,omitempty"`

	Purpose string `json:"purpose,omitempty"`

	// The username that will identify that new username
	Username string `json:"username,omitempty"`

	// The user's UUID
	Uuid string `json:"uuid,omitempty"`
}

// AssertUserCreateRequired checks if the required fields are not zero-ed
func AssertUserCreateRequired(obj UserCreate) error {
	return nil
}

// AssertUserCreateConstraints checks if the values respects the defined constraints
func AssertUserCreateConstraints(obj UserCreate) error {
	return nil
}