/*
 * accent-confd
 *
 * Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.
 *
 * API version: 1.1
 * Contact: help@accentvoice.io
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package confdserver

type Ha struct {

	// The role of this server in a HA cluster.
	NodeType string `json:"node_type"`

	// The remote IP address of the other machine in the HA cluster. Must be an IP address, a hostname is invalid.
	RemoteAddress string `json:"remote_address,omitempty"`
}

// AssertHaRequired checks if the required fields are not zero-ed
func AssertHaRequired(obj Ha) error {
	elements := map[string]interface{}{
		"node_type": obj.NodeType,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertHaConstraints checks if the values respects the defined constraints
func AssertHaConstraints(obj Ha) error {
	return nil
}
