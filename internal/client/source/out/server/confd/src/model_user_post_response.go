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




type UserPostResponse struct {

	Agent UserRelationAgent `json:"agent,omitempty"`

	Auth AuthUserPostAuth `json:"auth,omitempty"`

	// Overwrite all passwords set in call permissions associated to the user
	CallPermissionPassword string `json:"call_permission_password,omitempty"`

	// Record all calls made by this user (deprecated). If set, all `call_record_*_enabled` will be affected
	CallRecordEnabled bool `json:"call_record_enabled,omitempty"`

	// Record all external calls received by this user.
	CallRecordIncomingExternalEnabled bool `json:"call_record_incoming_external_enabled,omitempty"`

	// Record all internal calls received by this user.
	CallRecordIncomingInternalEnabled bool `json:"call_record_incoming_internal_enabled,omitempty"`

	// Record all external calls made by this user
	CallRecordOutgoingExternalEnabled bool `json:"call_record_outgoing_external_enabled,omitempty"`

	// Record all internal calls received by this user
	CallRecordOutgoingInternalEnabled bool `json:"call_record_outgoing_internal_enabled,omitempty"`

	// Authorize call transfers
	CallTransferEnabled bool `json:"call_transfer_enabled,omitempty"`

	// Name that appears on the phone when calling. Formatted as \"Firstname Lastname\" < number >
	CallerId string `json:"caller_id,omitempty"`

	// Additional information about the user
	Description string `json:"description,omitempty"`

	// Authorize to hangup with DTMF
	DtmfHangupEnabled bool `json:"dtmf_hangup_enabled,omitempty"`

	// User's email. Used for directories (i.e. by accent-dird)
	Email string `json:"email,omitempty"`

	// Enable/Disable the user
	Enabled bool `json:"enabled,omitempty"`

	Fallbacks UserFallbacks `json:"fallbacks,omitempty"`

	// User firstname
	Firstname string `json:"firstname,omitempty"`

	Forwards UserForwards `json:"forwards,omitempty"`

	Groups []map[string]interface{} `json:"groups,omitempty"`

	// User ID
	Id int32 `json:"id,omitempty"`

	// User language (e.g. \"en_US\")
	Language string `json:"language,omitempty"`

	// User lastname
	Lastname string `json:"lastname,omitempty"`

	Lines []UserPostRelationLine `json:"lines,omitempty"`

	// Phone number for the user’s mobile device
	MobilePhoneNumber string `json:"mobile_phone_number,omitempty"`

	// Name of the MOH category to use for music on hold
	MusicOnHold string `json:"music_on_hold,omitempty"`

	// Allow user to record a call by pressing *3
	OnlineCallRecordEnabled bool `json:"online_call_record_enabled,omitempty"`

	// Name that appears on the phone when calling
	OutgoingCallerId string `json:"outgoing_caller_id,omitempty"`

	// Password for connecting to the CTI (deprecated)
	Password string `json:"password,omitempty"`

	// Name of the subroutine to execute in asterisk before receiving a call
	PreprocessSubroutine string `json:"preprocess_subroutine,omitempty"`

	// Number of seconds a user's phone will ring before falling back
	RingSeconds int32 `json:"ring_seconds,omitempty"`

	// Number of allowed simultaneous calls on a user's phone
	SimultaneousCalls int32 `json:"simultaneous_calls,omitempty"`

	// Activate presence sharing in the accent client
	SupervisionEnabled bool `json:"supervision_enabled,omitempty"`

	Switchboards []map[string]interface{} `json:"switchboards,omitempty"`

	// The UUID of the tenant
	TenantUuid string `json:"tenant_uuid,omitempty"`

	// User timezone
	Timezone string `json:"timezone,omitempty"`

	// A custom field which purpose is left to the client. If the user has no userfield, then this field is an empty string.
	Userfield string `json:"userfield,omitempty"`

	// Username for connecting to the CTI (deprecated)
	Username string `json:"username,omitempty"`

	// User UUID. This ID is globally unique across multiple Accent instances
	Uuid string `json:"uuid,omitempty"`

	Voicemail Voicemail `json:"voicemail,omitempty"`

	Incalls []UserPostRelationIncallsIncallsInner `json:"incalls,omitempty"`
}

// AssertUserPostResponseRequired checks if the required fields are not zero-ed
func AssertUserPostResponseRequired(obj UserPostResponse) error {
	if err := AssertUserRelationAgentRequired(obj.Agent); err != nil {
		return err
	}
	if err := AssertAuthUserPostAuthRequired(obj.Auth); err != nil {
		return err
	}
	if err := AssertUserFallbacksRequired(obj.Fallbacks); err != nil {
		return err
	}
	if err := AssertUserForwardsRequired(obj.Forwards); err != nil {
		return err
	}
	for _, el := range obj.Groups {
		if err := Assertmap[string]interface{}Required(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Lines {
		if err := AssertUserPostRelationLineRequired(el); err != nil {
			return err
		}
	}
	for _, el := range obj.Switchboards {
		if err := Assertmap[string]interface{}Required(el); err != nil {
			return err
		}
	}
	if err := AssertVoicemailRequired(obj.Voicemail); err != nil {
		return err
	}
	for _, el := range obj.Incalls {
		if err := AssertUserPostRelationIncallsIncallsInnerRequired(el); err != nil {
			return err
		}
	}
	return nil
}

// AssertUserPostResponseConstraints checks if the values respects the defined constraints
func AssertUserPostResponseConstraints(obj UserPostResponse) error {
	return nil
}
