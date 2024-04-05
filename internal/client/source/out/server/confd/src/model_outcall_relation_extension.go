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

type OutcallRelationExtension struct {
	Context string `json:"context,omitempty"`

	Exten string `json:"exten,omitempty"`

	// Extension ID
	Id int32 `json:"id,omitempty"`

	CallerId string `json:"caller_id,omitempty"`

	// The prefix added to the outgoing call when sent on the trunk
	ExternalPrefix string `json:"external_prefix,omitempty"`

	// The prefix that the user need to dial before the extension
	Prefix string `json:"prefix,omitempty"`

	// The number of leading digits to strip
	StripDigits int32 `json:"strip_digits,omitempty"`
}

// AssertOutcallRelationExtensionRequired checks if the required fields are not zero-ed
func AssertOutcallRelationExtensionRequired(obj OutcallRelationExtension) error {
	return nil
}

// AssertOutcallRelationExtensionConstraints checks if the values respects the defined constraints
func AssertOutcallRelationExtensionConstraints(obj OutcallRelationExtension) error {
	return nil
}