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

import (
	"time"
)

type Meeting struct {
	CreationTime time.Time `json:"creation_time,omitempty"`

	// the external extension to dial to reach this meeting
	Exten string `json:"exten,omitempty"`

	// Format: base64(username:password), same as HTTP Basic Auth.
	GuestSipAuthorization string `json:"guest_sip_authorization,omitempty"`

	// URI to reach this stack (configured by the Ingress HTTP resource)
	IngressHttpUri string `json:"ingress_http_uri,omitempty"`

	Name string `json:"name,omitempty"`

	OwnerUuids []string `json:"owner_uuids,omitempty"`

	// Persistent meetings will not get deleted automatically
	Persistent bool `json:"persistent,omitempty"`

	// when `true`, the `guest_sip_authorization` is always `null`. Instead, clients must request an authorization to access the meeting.
	RequireAuthorization bool `json:"require_authorization,omitempty"`

	Uuid string `json:"uuid,omitempty"`
}

// AssertMeetingRequired checks if the required fields are not zero-ed
func AssertMeetingRequired(obj Meeting) error {
	return nil
}

// AssertMeetingConstraints checks if the values respects the defined constraints
func AssertMeetingConstraints(obj Meeting) error {
	return nil
}