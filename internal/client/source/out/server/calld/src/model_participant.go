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

type Participant struct {

	// Is the participant an admin of the conference?
	Admin bool `json:"admin,omitempty"`

	// The ID of the participant's call
	CallId string `json:"call_id,omitempty"`

	// The participant's name
	CallerIdName string `json:"caller_id_name,omitempty"`

	// The participant's number
	CallerIdNum string `json:"caller_id_num,omitempty"`

	// The participant's ID
	Id string `json:"id,omitempty"`

	// Elapsed seconds since the participant joined the conference
	JoinTime int32 `json:"join_time,omitempty"`

	// The participant's language
	Language string `json:"language,omitempty"`

	// Is the participant muted?
	Muted bool `json:"muted,omitempty"`

	// The UUID of the participant's user. `null` when there is no user.
	UserUuid string `json:"user_uuid,omitempty"`
}

// AssertParticipantRequired checks if the required fields are not zero-ed
func AssertParticipantRequired(obj Participant) error {
	return nil
}

// AssertParticipantConstraints checks if the values respects the defined constraints
func AssertParticipantConstraints(obj Participant) error {
	return nil
}