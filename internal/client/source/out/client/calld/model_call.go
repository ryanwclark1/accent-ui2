/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"encoding/json"
	"time"
)

// checks if the Call type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Call{}

// Call struct for Call
type Call struct {
	AnswerTime      *time.Time `json:"answer_time,omitempty"`
	Bridges         []string   `json:"bridges,omitempty"`
	CallId          *string    `json:"call_id,omitempty"`
	CallerIdName    *string    `json:"caller_id_name,omitempty"`
	CallerIdNumber  *string    `json:"caller_id_number,omitempty"`
	ConversationId  *string    `json:"conversation_id,omitempty"`
	CreationTime    *time.Time `json:"creation_time,omitempty"`
	DialedExtension *string    `json:"dialed_extension,omitempty"`
	Direction       *string    `json:"direction,omitempty"`
	HangupTime      *time.Time `json:"hangup_time,omitempty"`
	// This value is only correct when the destination of the call is a user or outgoing call. In other cases, it is always False.
	IsCaller *bool `json:"is_caller,omitempty"`
	// If this call has a video track
	IsVideo *bool `json:"is_video,omitempty"`
	// Line ID of the endpoint making the call
	LineId             *int32  `json:"line_id,omitempty"`
	Muted              *bool   `json:"muted,omitempty"`
	OnHold             *bool   `json:"on_hold,omitempty"`
	PeerCallerIdName   *string `json:"peer_caller_id_name,omitempty"`
	PeerCallerIdNumber *string `json:"peer_caller_id_number,omitempty"`
	RecordState        *string `json:"record_state,omitempty"`
	// Matches the `Call-ID` SIP header of the call. This value can be `null` when not using SIP
	SipCallId *string    `json:"sip_call_id,omitempty"`
	Status    *string    `json:"status,omitempty"`
	TalkingTo *TalkingTo `json:"talking_to,omitempty"`
	UserUuid  *string    `json:"user_uuid,omitempty"`
}

// NewCall instantiates a new Call object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCall() *Call {
	this := Call{}
	return &this
}

// NewCallWithDefaults instantiates a new Call object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCallWithDefaults() *Call {
	this := Call{}
	return &this
}

// GetAnswerTime returns the AnswerTime field value if set, zero value otherwise.
func (o *Call) GetAnswerTime() time.Time {
	if o == nil || IsNil(o.AnswerTime) {
		var ret time.Time
		return ret
	}
	return *o.AnswerTime
}

// GetAnswerTimeOk returns a tuple with the AnswerTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetAnswerTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.AnswerTime) {
		return nil, false
	}
	return o.AnswerTime, true
}

// HasAnswerTime returns a boolean if a field has been set.
func (o *Call) HasAnswerTime() bool {
	if o != nil && !IsNil(o.AnswerTime) {
		return true
	}

	return false
}

// SetAnswerTime gets a reference to the given time.Time and assigns it to the AnswerTime field.
func (o *Call) SetAnswerTime(v time.Time) {
	o.AnswerTime = &v
}

// GetBridges returns the Bridges field value if set, zero value otherwise.
func (o *Call) GetBridges() []string {
	if o == nil || IsNil(o.Bridges) {
		var ret []string
		return ret
	}
	return o.Bridges
}

// GetBridgesOk returns a tuple with the Bridges field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetBridgesOk() ([]string, bool) {
	if o == nil || IsNil(o.Bridges) {
		return nil, false
	}
	return o.Bridges, true
}

// HasBridges returns a boolean if a field has been set.
func (o *Call) HasBridges() bool {
	if o != nil && !IsNil(o.Bridges) {
		return true
	}

	return false
}

// SetBridges gets a reference to the given []string and assigns it to the Bridges field.
func (o *Call) SetBridges(v []string) {
	o.Bridges = v
}

// GetCallId returns the CallId field value if set, zero value otherwise.
func (o *Call) GetCallId() string {
	if o == nil || IsNil(o.CallId) {
		var ret string
		return ret
	}
	return *o.CallId
}

// GetCallIdOk returns a tuple with the CallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.CallId) {
		return nil, false
	}
	return o.CallId, true
}

// HasCallId returns a boolean if a field has been set.
func (o *Call) HasCallId() bool {
	if o != nil && !IsNil(o.CallId) {
		return true
	}

	return false
}

// SetCallId gets a reference to the given string and assigns it to the CallId field.
func (o *Call) SetCallId(v string) {
	o.CallId = &v
}

// GetCallerIdName returns the CallerIdName field value if set, zero value otherwise.
func (o *Call) GetCallerIdName() string {
	if o == nil || IsNil(o.CallerIdName) {
		var ret string
		return ret
	}
	return *o.CallerIdName
}

// GetCallerIdNameOk returns a tuple with the CallerIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetCallerIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdName) {
		return nil, false
	}
	return o.CallerIdName, true
}

// HasCallerIdName returns a boolean if a field has been set.
func (o *Call) HasCallerIdName() bool {
	if o != nil && !IsNil(o.CallerIdName) {
		return true
	}

	return false
}

// SetCallerIdName gets a reference to the given string and assigns it to the CallerIdName field.
func (o *Call) SetCallerIdName(v string) {
	o.CallerIdName = &v
}

// GetCallerIdNumber returns the CallerIdNumber field value if set, zero value otherwise.
func (o *Call) GetCallerIdNumber() string {
	if o == nil || IsNil(o.CallerIdNumber) {
		var ret string
		return ret
	}
	return *o.CallerIdNumber
}

// GetCallerIdNumberOk returns a tuple with the CallerIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetCallerIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.CallerIdNumber) {
		return nil, false
	}
	return o.CallerIdNumber, true
}

// HasCallerIdNumber returns a boolean if a field has been set.
func (o *Call) HasCallerIdNumber() bool {
	if o != nil && !IsNil(o.CallerIdNumber) {
		return true
	}

	return false
}

// SetCallerIdNumber gets a reference to the given string and assigns it to the CallerIdNumber field.
func (o *Call) SetCallerIdNumber(v string) {
	o.CallerIdNumber = &v
}

// GetConversationId returns the ConversationId field value if set, zero value otherwise.
func (o *Call) GetConversationId() string {
	if o == nil || IsNil(o.ConversationId) {
		var ret string
		return ret
	}
	return *o.ConversationId
}

// GetConversationIdOk returns a tuple with the ConversationId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetConversationIdOk() (*string, bool) {
	if o == nil || IsNil(o.ConversationId) {
		return nil, false
	}
	return o.ConversationId, true
}

// HasConversationId returns a boolean if a field has been set.
func (o *Call) HasConversationId() bool {
	if o != nil && !IsNil(o.ConversationId) {
		return true
	}

	return false
}

// SetConversationId gets a reference to the given string and assigns it to the ConversationId field.
func (o *Call) SetConversationId(v string) {
	o.ConversationId = &v
}

// GetCreationTime returns the CreationTime field value if set, zero value otherwise.
func (o *Call) GetCreationTime() time.Time {
	if o == nil || IsNil(o.CreationTime) {
		var ret time.Time
		return ret
	}
	return *o.CreationTime
}

// GetCreationTimeOk returns a tuple with the CreationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetCreationTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.CreationTime) {
		return nil, false
	}
	return o.CreationTime, true
}

// HasCreationTime returns a boolean if a field has been set.
func (o *Call) HasCreationTime() bool {
	if o != nil && !IsNil(o.CreationTime) {
		return true
	}

	return false
}

// SetCreationTime gets a reference to the given time.Time and assigns it to the CreationTime field.
func (o *Call) SetCreationTime(v time.Time) {
	o.CreationTime = &v
}

// GetDialedExtension returns the DialedExtension field value if set, zero value otherwise.
func (o *Call) GetDialedExtension() string {
	if o == nil || IsNil(o.DialedExtension) {
		var ret string
		return ret
	}
	return *o.DialedExtension
}

// GetDialedExtensionOk returns a tuple with the DialedExtension field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetDialedExtensionOk() (*string, bool) {
	if o == nil || IsNil(o.DialedExtension) {
		return nil, false
	}
	return o.DialedExtension, true
}

// HasDialedExtension returns a boolean if a field has been set.
func (o *Call) HasDialedExtension() bool {
	if o != nil && !IsNil(o.DialedExtension) {
		return true
	}

	return false
}

// SetDialedExtension gets a reference to the given string and assigns it to the DialedExtension field.
func (o *Call) SetDialedExtension(v string) {
	o.DialedExtension = &v
}

// GetDirection returns the Direction field value if set, zero value otherwise.
func (o *Call) GetDirection() string {
	if o == nil || IsNil(o.Direction) {
		var ret string
		return ret
	}
	return *o.Direction
}

// GetDirectionOk returns a tuple with the Direction field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetDirectionOk() (*string, bool) {
	if o == nil || IsNil(o.Direction) {
		return nil, false
	}
	return o.Direction, true
}

// HasDirection returns a boolean if a field has been set.
func (o *Call) HasDirection() bool {
	if o != nil && !IsNil(o.Direction) {
		return true
	}

	return false
}

// SetDirection gets a reference to the given string and assigns it to the Direction field.
func (o *Call) SetDirection(v string) {
	o.Direction = &v
}

// GetHangupTime returns the HangupTime field value if set, zero value otherwise.
func (o *Call) GetHangupTime() time.Time {
	if o == nil || IsNil(o.HangupTime) {
		var ret time.Time
		return ret
	}
	return *o.HangupTime
}

// GetHangupTimeOk returns a tuple with the HangupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetHangupTimeOk() (*time.Time, bool) {
	if o == nil || IsNil(o.HangupTime) {
		return nil, false
	}
	return o.HangupTime, true
}

// HasHangupTime returns a boolean if a field has been set.
func (o *Call) HasHangupTime() bool {
	if o != nil && !IsNil(o.HangupTime) {
		return true
	}

	return false
}

// SetHangupTime gets a reference to the given time.Time and assigns it to the HangupTime field.
func (o *Call) SetHangupTime(v time.Time) {
	o.HangupTime = &v
}

// GetIsCaller returns the IsCaller field value if set, zero value otherwise.
func (o *Call) GetIsCaller() bool {
	if o == nil || IsNil(o.IsCaller) {
		var ret bool
		return ret
	}
	return *o.IsCaller
}

// GetIsCallerOk returns a tuple with the IsCaller field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetIsCallerOk() (*bool, bool) {
	if o == nil || IsNil(o.IsCaller) {
		return nil, false
	}
	return o.IsCaller, true
}

// HasIsCaller returns a boolean if a field has been set.
func (o *Call) HasIsCaller() bool {
	if o != nil && !IsNil(o.IsCaller) {
		return true
	}

	return false
}

// SetIsCaller gets a reference to the given bool and assigns it to the IsCaller field.
func (o *Call) SetIsCaller(v bool) {
	o.IsCaller = &v
}

// GetIsVideo returns the IsVideo field value if set, zero value otherwise.
func (o *Call) GetIsVideo() bool {
	if o == nil || IsNil(o.IsVideo) {
		var ret bool
		return ret
	}
	return *o.IsVideo
}

// GetIsVideoOk returns a tuple with the IsVideo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetIsVideoOk() (*bool, bool) {
	if o == nil || IsNil(o.IsVideo) {
		return nil, false
	}
	return o.IsVideo, true
}

// HasIsVideo returns a boolean if a field has been set.
func (o *Call) HasIsVideo() bool {
	if o != nil && !IsNil(o.IsVideo) {
		return true
	}

	return false
}

// SetIsVideo gets a reference to the given bool and assigns it to the IsVideo field.
func (o *Call) SetIsVideo(v bool) {
	o.IsVideo = &v
}

// GetLineId returns the LineId field value if set, zero value otherwise.
func (o *Call) GetLineId() int32 {
	if o == nil || IsNil(o.LineId) {
		var ret int32
		return ret
	}
	return *o.LineId
}

// GetLineIdOk returns a tuple with the LineId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetLineIdOk() (*int32, bool) {
	if o == nil || IsNil(o.LineId) {
		return nil, false
	}
	return o.LineId, true
}

// HasLineId returns a boolean if a field has been set.
func (o *Call) HasLineId() bool {
	if o != nil && !IsNil(o.LineId) {
		return true
	}

	return false
}

// SetLineId gets a reference to the given int32 and assigns it to the LineId field.
func (o *Call) SetLineId(v int32) {
	o.LineId = &v
}

// GetMuted returns the Muted field value if set, zero value otherwise.
func (o *Call) GetMuted() bool {
	if o == nil || IsNil(o.Muted) {
		var ret bool
		return ret
	}
	return *o.Muted
}

// GetMutedOk returns a tuple with the Muted field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetMutedOk() (*bool, bool) {
	if o == nil || IsNil(o.Muted) {
		return nil, false
	}
	return o.Muted, true
}

// HasMuted returns a boolean if a field has been set.
func (o *Call) HasMuted() bool {
	if o != nil && !IsNil(o.Muted) {
		return true
	}

	return false
}

// SetMuted gets a reference to the given bool and assigns it to the Muted field.
func (o *Call) SetMuted(v bool) {
	o.Muted = &v
}

// GetOnHold returns the OnHold field value if set, zero value otherwise.
func (o *Call) GetOnHold() bool {
	if o == nil || IsNil(o.OnHold) {
		var ret bool
		return ret
	}
	return *o.OnHold
}

// GetOnHoldOk returns a tuple with the OnHold field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetOnHoldOk() (*bool, bool) {
	if o == nil || IsNil(o.OnHold) {
		return nil, false
	}
	return o.OnHold, true
}

// HasOnHold returns a boolean if a field has been set.
func (o *Call) HasOnHold() bool {
	if o != nil && !IsNil(o.OnHold) {
		return true
	}

	return false
}

// SetOnHold gets a reference to the given bool and assigns it to the OnHold field.
func (o *Call) SetOnHold(v bool) {
	o.OnHold = &v
}

// GetPeerCallerIdName returns the PeerCallerIdName field value if set, zero value otherwise.
func (o *Call) GetPeerCallerIdName() string {
	if o == nil || IsNil(o.PeerCallerIdName) {
		var ret string
		return ret
	}
	return *o.PeerCallerIdName
}

// GetPeerCallerIdNameOk returns a tuple with the PeerCallerIdName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetPeerCallerIdNameOk() (*string, bool) {
	if o == nil || IsNil(o.PeerCallerIdName) {
		return nil, false
	}
	return o.PeerCallerIdName, true
}

// HasPeerCallerIdName returns a boolean if a field has been set.
func (o *Call) HasPeerCallerIdName() bool {
	if o != nil && !IsNil(o.PeerCallerIdName) {
		return true
	}

	return false
}

// SetPeerCallerIdName gets a reference to the given string and assigns it to the PeerCallerIdName field.
func (o *Call) SetPeerCallerIdName(v string) {
	o.PeerCallerIdName = &v
}

// GetPeerCallerIdNumber returns the PeerCallerIdNumber field value if set, zero value otherwise.
func (o *Call) GetPeerCallerIdNumber() string {
	if o == nil || IsNil(o.PeerCallerIdNumber) {
		var ret string
		return ret
	}
	return *o.PeerCallerIdNumber
}

// GetPeerCallerIdNumberOk returns a tuple with the PeerCallerIdNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetPeerCallerIdNumberOk() (*string, bool) {
	if o == nil || IsNil(o.PeerCallerIdNumber) {
		return nil, false
	}
	return o.PeerCallerIdNumber, true
}

// HasPeerCallerIdNumber returns a boolean if a field has been set.
func (o *Call) HasPeerCallerIdNumber() bool {
	if o != nil && !IsNil(o.PeerCallerIdNumber) {
		return true
	}

	return false
}

// SetPeerCallerIdNumber gets a reference to the given string and assigns it to the PeerCallerIdNumber field.
func (o *Call) SetPeerCallerIdNumber(v string) {
	o.PeerCallerIdNumber = &v
}

// GetRecordState returns the RecordState field value if set, zero value otherwise.
func (o *Call) GetRecordState() string {
	if o == nil || IsNil(o.RecordState) {
		var ret string
		return ret
	}
	return *o.RecordState
}

// GetRecordStateOk returns a tuple with the RecordState field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetRecordStateOk() (*string, bool) {
	if o == nil || IsNil(o.RecordState) {
		return nil, false
	}
	return o.RecordState, true
}

// HasRecordState returns a boolean if a field has been set.
func (o *Call) HasRecordState() bool {
	if o != nil && !IsNil(o.RecordState) {
		return true
	}

	return false
}

// SetRecordState gets a reference to the given string and assigns it to the RecordState field.
func (o *Call) SetRecordState(v string) {
	o.RecordState = &v
}

// GetSipCallId returns the SipCallId field value if set, zero value otherwise.
func (o *Call) GetSipCallId() string {
	if o == nil || IsNil(o.SipCallId) {
		var ret string
		return ret
	}
	return *o.SipCallId
}

// GetSipCallIdOk returns a tuple with the SipCallId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetSipCallIdOk() (*string, bool) {
	if o == nil || IsNil(o.SipCallId) {
		return nil, false
	}
	return o.SipCallId, true
}

// HasSipCallId returns a boolean if a field has been set.
func (o *Call) HasSipCallId() bool {
	if o != nil && !IsNil(o.SipCallId) {
		return true
	}

	return false
}

// SetSipCallId gets a reference to the given string and assigns it to the SipCallId field.
func (o *Call) SetSipCallId(v string) {
	o.SipCallId = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Call) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Call) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Call) SetStatus(v string) {
	o.Status = &v
}

// GetTalkingTo returns the TalkingTo field value if set, zero value otherwise.
func (o *Call) GetTalkingTo() TalkingTo {
	if o == nil || IsNil(o.TalkingTo) {
		var ret TalkingTo
		return ret
	}
	return *o.TalkingTo
}

// GetTalkingToOk returns a tuple with the TalkingTo field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetTalkingToOk() (*TalkingTo, bool) {
	if o == nil || IsNil(o.TalkingTo) {
		return nil, false
	}
	return o.TalkingTo, true
}

// HasTalkingTo returns a boolean if a field has been set.
func (o *Call) HasTalkingTo() bool {
	if o != nil && !IsNil(o.TalkingTo) {
		return true
	}

	return false
}

// SetTalkingTo gets a reference to the given TalkingTo and assigns it to the TalkingTo field.
func (o *Call) SetTalkingTo(v TalkingTo) {
	o.TalkingTo = &v
}

// GetUserUuid returns the UserUuid field value if set, zero value otherwise.
func (o *Call) GetUserUuid() string {
	if o == nil || IsNil(o.UserUuid) {
		var ret string
		return ret
	}
	return *o.UserUuid
}

// GetUserUuidOk returns a tuple with the UserUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Call) GetUserUuidOk() (*string, bool) {
	if o == nil || IsNil(o.UserUuid) {
		return nil, false
	}
	return o.UserUuid, true
}

// HasUserUuid returns a boolean if a field has been set.
func (o *Call) HasUserUuid() bool {
	if o != nil && !IsNil(o.UserUuid) {
		return true
	}

	return false
}

// SetUserUuid gets a reference to the given string and assigns it to the UserUuid field.
func (o *Call) SetUserUuid(v string) {
	o.UserUuid = &v
}

func (o Call) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Call) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AnswerTime) {
		toSerialize["answer_time"] = o.AnswerTime
	}
	if !IsNil(o.Bridges) {
		toSerialize["bridges"] = o.Bridges
	}
	if !IsNil(o.CallId) {
		toSerialize["call_id"] = o.CallId
	}
	if !IsNil(o.CallerIdName) {
		toSerialize["caller_id_name"] = o.CallerIdName
	}
	if !IsNil(o.CallerIdNumber) {
		toSerialize["caller_id_number"] = o.CallerIdNumber
	}
	if !IsNil(o.ConversationId) {
		toSerialize["conversation_id"] = o.ConversationId
	}
	if !IsNil(o.CreationTime) {
		toSerialize["creation_time"] = o.CreationTime
	}
	if !IsNil(o.DialedExtension) {
		toSerialize["dialed_extension"] = o.DialedExtension
	}
	if !IsNil(o.Direction) {
		toSerialize["direction"] = o.Direction
	}
	if !IsNil(o.HangupTime) {
		toSerialize["hangup_time"] = o.HangupTime
	}
	if !IsNil(o.IsCaller) {
		toSerialize["is_caller"] = o.IsCaller
	}
	if !IsNil(o.IsVideo) {
		toSerialize["is_video"] = o.IsVideo
	}
	if !IsNil(o.LineId) {
		toSerialize["line_id"] = o.LineId
	}
	if !IsNil(o.Muted) {
		toSerialize["muted"] = o.Muted
	}
	if !IsNil(o.OnHold) {
		toSerialize["on_hold"] = o.OnHold
	}
	if !IsNil(o.PeerCallerIdName) {
		toSerialize["peer_caller_id_name"] = o.PeerCallerIdName
	}
	if !IsNil(o.PeerCallerIdNumber) {
		toSerialize["peer_caller_id_number"] = o.PeerCallerIdNumber
	}
	if !IsNil(o.RecordState) {
		toSerialize["record_state"] = o.RecordState
	}
	if !IsNil(o.SipCallId) {
		toSerialize["sip_call_id"] = o.SipCallId
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TalkingTo) {
		toSerialize["talking_to"] = o.TalkingTo
	}
	if !IsNil(o.UserUuid) {
		toSerialize["user_uuid"] = o.UserUuid
	}
	return toSerialize, nil
}

type NullableCall struct {
	value *Call
	isSet bool
}

func (v NullableCall) Get() *Call {
	return v.value
}

func (v *NullableCall) Set(val *Call) {
	v.value = val
	v.isSet = true
}

func (v NullableCall) IsSet() bool {
	return v.isSet
}

func (v *NullableCall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCall(val *Call) *NullableCall {
	return &NullableCall{value: val, isSet: true}
}

func (v NullableCall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
