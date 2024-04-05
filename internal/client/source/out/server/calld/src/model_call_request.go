/*
 * accent-calld
 *
 * Control your calls from a REST API
 *
 * API version: 1.0.0
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package calldserver

type CallRequest struct {
	Destination CallRequestDestination `json:"destination"`

	Source CallRequestSource `json:"source"`

	// Channel variables to set
	Variables map[string]interface{} `json:"variables,omitempty"`
}

// AssertCallRequestRequired checks if the required fields are not zero-ed
func AssertCallRequestRequired(obj CallRequest) error {
	elements := map[string]interface{}{
		"destination": obj.Destination,
		"source":      obj.Source,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	if err := AssertCallRequestDestinationRequired(obj.Destination); err != nil {
		return err
	}
	if err := AssertCallRequestSourceRequired(obj.Source); err != nil {
		return err
	}
	return nil
}

// AssertCallRequestConstraints checks if the values respects the defined constraints
func AssertCallRequestConstraints(obj CallRequest) error {
	return nil
}