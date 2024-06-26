/*
 * accent-amid
 *
 * Send AMI actions to Asterisk, providing token based authentication.
 *
 * API version: 0.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package amidserver

type Command struct {
	Command string `json:"command"`
}

// AssertCommandRequired checks if the required fields are not zero-ed
func AssertCommandRequired(obj Command) error {
	elements := map[string]interface{}{
		"command": obj.Command,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertCommandConstraints checks if the values respects the defined constraints
func AssertCommandConstraints(obj Command) error {
	return nil
}
