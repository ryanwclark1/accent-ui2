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

type Outcall struct {

	// The id of the outgoing call
	Id int32 `json:"id,omitempty"`

	// The name of the outcall
	Name string `json:"name,omitempty"`

	Trunks []OutcallRelationTrunk `json:"trunks,omitempty"`

	Extensions []OutcallRelationExtension `json:"extensions,omitempty"`

	Schedules []OutcallRelationSchedule `json:"schedules,omitempty"`

	CallPermissions []CallPermissionRelationBase `json:"call_permissions,omitempty"`

	// Additional information about the outgoing call
	Description string `json:"description,omitempty"`

	Enabled bool `json:"enabled,omitempty"`

	// Use the internal caller id
	InternalCallerId bool `json:"internal_caller_id,omitempty"`

	// Name of the subroutine to execute in asterisk before receiving a call
	PreprocessSubroutine string `json:"preprocess_subroutine,omitempty"`

	// Ringing time before hangup
	RingTime int32 `json:"ring_time,omitempty"`

	// The UUID of the tenant
	TenantUuid string `json:"tenant_uuid,omitempty"`
}

// AssertOutcallRequired checks if the required fields are not zero-ed
func AssertOutcallRequired(obj Outcall) error {
	for _, el := range obj.Trunks {
		if err := AssertOutcallRelationTrunkRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Extensions {
		if err := AssertOutcallRelationExtensionRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Schedules {
		if err := AssertOutcallRelationScheduleRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.CallPermissions {
		if err := AssertCallPermissionRelationBaseRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertOutcallConstraints checks if the values respects the defined constraints
func AssertOutcallConstraints(obj Outcall) error {
	return nil
}
