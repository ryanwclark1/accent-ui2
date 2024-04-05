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

type RegisterIax struct {

	// The id of the register IAX
	Id int32 `json:"id,omitempty"`

	// The password to authenticate to the remote_host
	AuthPassword string `json:"auth_password,omitempty"`

	// The username used by the remote_host for the authentication
	AuthUsername string `json:"auth_username,omitempty"`

	// The callback context to use for the register
	CallbackContext string `json:"callback_context,omitempty"`

	// The callback extension to use for the register
	CallbackExtension string `json:"callback_extension,omitempty"`

	// The register domain
	RemoteHost string `json:"remote_host"`

	// The port of the remote_host
	RemotePort int32 `json:"remote_port,omitempty"`
}

// AssertRegisterIaxRequired checks if the required fields are not zero-ed
func AssertRegisterIaxRequired(obj RegisterIax) error {
	elements := map[string]interface{}{
		"remote_host": obj.RemoteHost,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertRegisterIaxConstraints checks if the values respects the defined constraints
func AssertRegisterIaxConstraints(obj RegisterIax) error {
	return nil
}