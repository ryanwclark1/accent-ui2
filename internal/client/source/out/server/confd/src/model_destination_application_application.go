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

type DestinationApplicationApplication struct {
	CallbackDisa DestinationApplicationCallbackDisa `json:"callback_disa,omitempty"`

	Custom DestinationApplicationCustom `json:"custom,omitempty"`

	Directory DestinationApplicationDirectory `json:"directory,omitempty"`

	Disa DestinationApplicationDisa `json:"disa,omitempty"`

	FaxToMail DestinationApplicationFaxToMail `json:"fax_to_mail,omitempty"`

	Voicemail DestinationApplicationVoicemail `json:"voicemail,omitempty"`
}

// AssertDestinationApplicationApplicationRequired checks if the required fields are not zero-ed
func AssertDestinationApplicationApplicationRequired(obj DestinationApplicationApplication) error {
	if err := AssertDestinationApplicationCallbackDisaRequired(obj.CallbackDisa); err != nil {
		return err
	}
	if err := AssertDestinationApplicationCustomRequired(obj.Custom); err != nil {
		return err
	}
	if err := AssertDestinationApplicationDirectoryRequired(obj.Directory); err != nil {
		return err
	}
	if err := AssertDestinationApplicationDisaRequired(obj.Disa); err != nil {
		return err
	}
	if err := AssertDestinationApplicationFaxToMailRequired(obj.FaxToMail); err != nil {
		return err
	}
	if err := AssertDestinationApplicationVoicemailRequired(obj.Voicemail); err != nil {
		return err
	}
	return nil
}

// AssertDestinationApplicationApplicationConstraints checks if the values respects the defined constraints
func AssertDestinationApplicationApplicationConstraints(obj DestinationApplicationApplication) error {
	return nil
}
