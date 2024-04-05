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

type ApplicationNodes struct {
	Items []ApplicationNode `json:"items,omitempty"`
}

// AssertApplicationNodesRequired checks if the required fields are not zero-ed
func AssertApplicationNodesRequired(obj ApplicationNodes) error {
	for _, el := range obj.Items {
		if err := AssertApplicationNodeRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertApplicationNodesConstraints checks if the values respects the defined constraints
func AssertApplicationNodesConstraints(obj ApplicationNodes) error {
	return nil
}
