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

type PasswordChange struct {

	// The desired password
	NewPassword string `json:"new_password"`

	// The old password
	OldPassword string `json:"old_password"`
}

// AssertPasswordChangeRequired checks if the required fields are not zero-ed
func AssertPasswordChangeRequired(obj PasswordChange) error {
	elements := map[string]interface{}{
		"new_password": obj.NewPassword,
		"old_password": obj.OldPassword,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertPasswordChangeConstraints checks if the values respects the defined constraints
func AssertPasswordChangeConstraints(obj PasswordChange) error {
	return nil
}
