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

type UserCallPermission struct {

	// Call Permission's ID
	CallPermissionId int32 `json:"call_permission_id"`

	// User's ID
	UserId int32 `json:"user_id"`
}

// AssertUserCallPermissionRequired checks if the required fields are not zero-ed
func AssertUserCallPermissionRequired(obj UserCallPermission) error {
	elements := map[string]interface{}{
		"call_permission_id": obj.CallPermissionId,
		"user_id":            obj.UserId,
	}
	for name, el := range elements {
		if isZero := IsZeroValue(el); isZero {
			return &RequiredError{Field: name}
		}
	}

	return nil
}

// AssertUserCallPermissionConstraints checks if the values respects the defined constraints
func AssertUserCallPermissionConstraints(obj UserCallPermission) error {
	return nil
}
