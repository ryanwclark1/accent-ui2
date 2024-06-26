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

type AdhocConference struct {

	// The ID of the adhoc conference
	ConferenceId string `json:"conference_id,omitempty"`
}

// AssertAdhocConferenceRequired checks if the required fields are not zero-ed
func AssertAdhocConferenceRequired(obj AdhocConference) error {
	return nil
}

// AssertAdhocConferenceConstraints checks if the values respects the defined constraints
func AssertAdhocConferenceConstraints(obj AdhocConference) error {
	return nil
}
