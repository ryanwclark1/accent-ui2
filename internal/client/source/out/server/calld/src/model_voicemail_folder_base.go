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

type VoicemailFolderBase struct {

	// The folder's ID
	Id int32 `json:"id,omitempty"`

	// The folder's name
	Name string `json:"name,omitempty"`

	// The folder's type. When a message if left on a voicemail, it is stored in the folder of type \"new\", unless if it is an urgent message, in which case it is left in the folder of type \"urgent\". When that messages is read, it is moved into the folder of type \"old\". All other folders used the type \"other\".
	Type string `json:"type,omitempty"`
}

// AssertVoicemailFolderBaseRequired checks if the required fields are not zero-ed
func AssertVoicemailFolderBaseRequired(obj VoicemailFolderBase) error {
	return nil
}

// AssertVoicemailFolderBaseConstraints checks if the values respects the defined constraints
func AssertVoicemailFolderBaseConstraints(obj VoicemailFolderBase) error {
	return nil
}
