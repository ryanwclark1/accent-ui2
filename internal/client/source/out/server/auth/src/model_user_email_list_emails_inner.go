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

type UserEmailListEmailsInner struct {
	Address string `json:"address,omitempty"`

	Main bool `json:"main"`
}

// AssertUserEmailListEmailsInnerRequired checks if the required fields are not zero-ed
func AssertUserEmailListEmailsInnerRequired(obj UserEmailListEmailsInner) error {
	elements := map[string]interface{}{
		"main": obj.Main,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUserEmailListEmailsInnerConstraints checks if the values respects the defined constraints
func AssertUserEmailListEmailsInnerConstraints(obj UserEmailListEmailsInner) error {
	return nil
}