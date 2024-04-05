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

type UserPostResponse struct {
	Emails []UserEmail `json:"emails,omitempty"`

	Purpose string `json:"purpose,omitempty"`

	Username string `json:"username,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// AssertUserPostResponseRequired checks if the required fields are not zero-ed
func AssertUserPostResponseRequired(obj UserPostResponse) error {
	for _, el := range obj.Emails {
		if err := AssertUserEmailRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUserPostResponseConstraints checks if the values respects the defined constraints
func AssertUserPostResponseConstraints(obj UserPostResponse) error {
	return nil
}
