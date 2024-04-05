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

import (
	"time"
)

type Call struct {
	AnswerTime time.Time `json:"answer_time,omitempty"`

	Bridges []string `json:"bridges,omitempty"`

	CallId string `json:"call_id,omitempty"`

	CallerIdName string `json:"caller_id_name,omitempty"`

	CallerIdNumber string `json:"caller_id_number,omitempty"`

	ConversationId string `json:"conversation_id,omitempty"`

	CreationTime time.Time `json:"creation_time,omitempty"`

	DialedExtension string `json:"dialed_extension,omitempty"`

	Direction string `json:"direction,omitempty"`

	HangupTime time.Time `json:"hangup_time,omitempty"`

	// This value is only correct when the destination of the call is a user or outgoing call. In other cases, it is always False.
	IsCaller bool `json:"is_caller,omitempty"`

	// If this call has a video track
	IsVideo bool `json:"is_video,omitempty"`

	// Line ID of the endpoint making the call
	LineId int32 `json:"line_id,omitempty"`

	Muted bool `json:"muted,omitempty"`

	OnHold bool `json:"on_hold,omitempty"`

	PeerCallerIdName string `json:"peer_caller_id_name,omitempty"`

	PeerCallerIdNumber string `json:"peer_caller_id_number,omitempty"`

	RecordState string `json:"record_state,omitempty"`

	// Matches the `Call-ID` SIP header of the call. This value can be `null` when not using SIP
	SipCallId string `json:"sip_call_id,omitempty"`

	Status string `json:"status,omitempty"`

	TalkingTo TalkingTo `json:"talking_to,omitempty"`

	UserUuid string `json:"user_uuid,omitempty"`
}

// AssertCallRequired checks if the required fields are not zero-ed
func AssertCallRequired(obj Call) error {
	if err := AssertTalkingToRequired(obj.TalkingTo); err != nil {
		return err
	}
	return nil
}

// AssertCallConstraints checks if the values respects the defined constraints
func AssertCallConstraints(obj Call) error {
	return nil
}
