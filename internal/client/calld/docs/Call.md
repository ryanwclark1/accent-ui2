# Call

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AnswerTime** | Pointer to **time.Time** |  | [optional]
**Bridges** | Pointer to **[]string** |  | [optional]
**CallId** | Pointer to **string** |  | [optional]
**CallerIdName** | Pointer to **string** |  | [optional]
**CallerIdNumber** | Pointer to **string** |  | [optional]
**ConversationId** | Pointer to **string** |  | [optional]
**CreationTime** | Pointer to **time.Time** |  | [optional]
**DialedExtension** | Pointer to **string** |  | [optional]
**Direction** | Pointer to **string** |  | [optional]
**HangupTime** | Pointer to **time.Time** |  | [optional]
**IsCaller** | Pointer to **bool** | This value is only correct when the destination of the call is a user or outgoing call. In other cases, it is always False. | [optional]
**IsVideo** | Pointer to **bool** | If this call has a video track | [optional]
**LineId** | Pointer to **int32** | Line ID of the endpoint making the call | [optional] [readonly]
**Muted** | Pointer to **bool** |  | [optional]
**OnHold** | Pointer to **bool** |  | [optional]
**PeerCallerIdName** | Pointer to **string** |  | [optional]
**PeerCallerIdNumber** | Pointer to **string** |  | [optional]
**RecordState** | Pointer to **string** |  | [optional]
**SipCallId** | Pointer to **string** | Matches the &#x60;Call-ID&#x60; SIP header of the call. This value can be &#x60;null&#x60; when not using SIP | [optional] [readonly]
**Status** | Pointer to **string** |  | [optional]
**TalkingTo** | Pointer to [**TalkingTo**](TalkingTo.md) |  | [optional]
**UserUuid** | Pointer to **string** |  | [optional]

## Methods

### NewCall

`func NewCall() *Call`

NewCall instantiates a new Call object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallWithDefaults

`func NewCallWithDefaults() *Call`

NewCallWithDefaults instantiates a new Call object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAnswerTime

`func (o *Call) GetAnswerTime() time.Time`

GetAnswerTime returns the AnswerTime field if non-nil, zero value otherwise.

### GetAnswerTimeOk

`func (o *Call) GetAnswerTimeOk() (*time.Time, bool)`

GetAnswerTimeOk returns a tuple with the AnswerTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswerTime

`func (o *Call) SetAnswerTime(v time.Time)`

SetAnswerTime sets AnswerTime field to given value.

### HasAnswerTime

`func (o *Call) HasAnswerTime() bool`

HasAnswerTime returns a boolean if a field has been set.

### GetBridges

`func (o *Call) GetBridges() []string`

GetBridges returns the Bridges field if non-nil, zero value otherwise.

### GetBridgesOk

`func (o *Call) GetBridgesOk() (*[]string, bool)`

GetBridgesOk returns a tuple with the Bridges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBridges

`func (o *Call) SetBridges(v []string)`

SetBridges sets Bridges field to given value.

### HasBridges

`func (o *Call) HasBridges() bool`

HasBridges returns a boolean if a field has been set.

### GetCallId

`func (o *Call) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Call) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Call) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *Call) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallerIdName

`func (o *Call) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Call) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Call) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Call) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNumber

`func (o *Call) GetCallerIdNumber() string`

GetCallerIdNumber returns the CallerIdNumber field if non-nil, zero value otherwise.

### GetCallerIdNumberOk

`func (o *Call) GetCallerIdNumberOk() (*string, bool)`

GetCallerIdNumberOk returns a tuple with the CallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNumber

`func (o *Call) SetCallerIdNumber(v string)`

SetCallerIdNumber sets CallerIdNumber field to given value.

### HasCallerIdNumber

`func (o *Call) HasCallerIdNumber() bool`

HasCallerIdNumber returns a boolean if a field has been set.

### GetConversationId

`func (o *Call) GetConversationId() string`

GetConversationId returns the ConversationId field if non-nil, zero value otherwise.

### GetConversationIdOk

`func (o *Call) GetConversationIdOk() (*string, bool)`

GetConversationIdOk returns a tuple with the ConversationId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationId

`func (o *Call) SetConversationId(v string)`

SetConversationId sets ConversationId field to given value.

### HasConversationId

`func (o *Call) HasConversationId() bool`

HasConversationId returns a boolean if a field has been set.

### GetCreationTime

`func (o *Call) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *Call) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *Call) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *Call) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetDialedExtension

`func (o *Call) GetDialedExtension() string`

GetDialedExtension returns the DialedExtension field if non-nil, zero value otherwise.

### GetDialedExtensionOk

`func (o *Call) GetDialedExtensionOk() (*string, bool)`

GetDialedExtensionOk returns a tuple with the DialedExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDialedExtension

`func (o *Call) SetDialedExtension(v string)`

SetDialedExtension sets DialedExtension field to given value.

### HasDialedExtension

`func (o *Call) HasDialedExtension() bool`

HasDialedExtension returns a boolean if a field has been set.

### GetDirection

`func (o *Call) GetDirection() string`

GetDirection returns the Direction field if non-nil, zero value otherwise.

### GetDirectionOk

`func (o *Call) GetDirectionOk() (*string, bool)`

GetDirectionOk returns a tuple with the Direction field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDirection

`func (o *Call) SetDirection(v string)`

SetDirection sets Direction field to given value.

### HasDirection

`func (o *Call) HasDirection() bool`

HasDirection returns a boolean if a field has been set.

### GetHangupTime

`func (o *Call) GetHangupTime() time.Time`

GetHangupTime returns the HangupTime field if non-nil, zero value otherwise.

### GetHangupTimeOk

`func (o *Call) GetHangupTimeOk() (*time.Time, bool)`

GetHangupTimeOk returns a tuple with the HangupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangupTime

`func (o *Call) SetHangupTime(v time.Time)`

SetHangupTime sets HangupTime field to given value.

### HasHangupTime

`func (o *Call) HasHangupTime() bool`

HasHangupTime returns a boolean if a field has been set.

### GetIsCaller

`func (o *Call) GetIsCaller() bool`

GetIsCaller returns the IsCaller field if non-nil, zero value otherwise.

### GetIsCallerOk

`func (o *Call) GetIsCallerOk() (*bool, bool)`

GetIsCallerOk returns a tuple with the IsCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsCaller

`func (o *Call) SetIsCaller(v bool)`

SetIsCaller sets IsCaller field to given value.

### HasIsCaller

`func (o *Call) HasIsCaller() bool`

HasIsCaller returns a boolean if a field has been set.

### GetIsVideo

`func (o *Call) GetIsVideo() bool`

GetIsVideo returns the IsVideo field if non-nil, zero value otherwise.

### GetIsVideoOk

`func (o *Call) GetIsVideoOk() (*bool, bool)`

GetIsVideoOk returns a tuple with the IsVideo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIsVideo

`func (o *Call) SetIsVideo(v bool)`

SetIsVideo sets IsVideo field to given value.

### HasIsVideo

`func (o *Call) HasIsVideo() bool`

HasIsVideo returns a boolean if a field has been set.

### GetLineId

`func (o *Call) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *Call) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *Call) SetLineId(v int32)`

SetLineId sets LineId field to given value.

### HasLineId

`func (o *Call) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### GetMuted

`func (o *Call) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Call) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Call) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *Call) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetOnHold

`func (o *Call) GetOnHold() bool`

GetOnHold returns the OnHold field if non-nil, zero value otherwise.

### GetOnHoldOk

`func (o *Call) GetOnHoldOk() (*bool, bool)`

GetOnHoldOk returns a tuple with the OnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnHold

`func (o *Call) SetOnHold(v bool)`

SetOnHold sets OnHold field to given value.

### HasOnHold

`func (o *Call) HasOnHold() bool`

HasOnHold returns a boolean if a field has been set.

### GetPeerCallerIdName

`func (o *Call) GetPeerCallerIdName() string`

GetPeerCallerIdName returns the PeerCallerIdName field if non-nil, zero value otherwise.

### GetPeerCallerIdNameOk

`func (o *Call) GetPeerCallerIdNameOk() (*string, bool)`

GetPeerCallerIdNameOk returns a tuple with the PeerCallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCallerIdName

`func (o *Call) SetPeerCallerIdName(v string)`

SetPeerCallerIdName sets PeerCallerIdName field to given value.

### HasPeerCallerIdName

`func (o *Call) HasPeerCallerIdName() bool`

HasPeerCallerIdName returns a boolean if a field has been set.

### GetPeerCallerIdNumber

`func (o *Call) GetPeerCallerIdNumber() string`

GetPeerCallerIdNumber returns the PeerCallerIdNumber field if non-nil, zero value otherwise.

### GetPeerCallerIdNumberOk

`func (o *Call) GetPeerCallerIdNumberOk() (*string, bool)`

GetPeerCallerIdNumberOk returns a tuple with the PeerCallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPeerCallerIdNumber

`func (o *Call) SetPeerCallerIdNumber(v string)`

SetPeerCallerIdNumber sets PeerCallerIdNumber field to given value.

### HasPeerCallerIdNumber

`func (o *Call) HasPeerCallerIdNumber() bool`

HasPeerCallerIdNumber returns a boolean if a field has been set.

### GetRecordState

`func (o *Call) GetRecordState() string`

GetRecordState returns the RecordState field if non-nil, zero value otherwise.

### GetRecordStateOk

`func (o *Call) GetRecordStateOk() (*string, bool)`

GetRecordStateOk returns a tuple with the RecordState field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecordState

`func (o *Call) SetRecordState(v string)`

SetRecordState sets RecordState field to given value.

### HasRecordState

`func (o *Call) HasRecordState() bool`

HasRecordState returns a boolean if a field has been set.

### GetSipCallId

`func (o *Call) GetSipCallId() string`

GetSipCallId returns the SipCallId field if non-nil, zero value otherwise.

### GetSipCallIdOk

`func (o *Call) GetSipCallIdOk() (*string, bool)`

GetSipCallIdOk returns a tuple with the SipCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipCallId

`func (o *Call) SetSipCallId(v string)`

SetSipCallId sets SipCallId field to given value.

### HasSipCallId

`func (o *Call) HasSipCallId() bool`

HasSipCallId returns a boolean if a field has been set.

### GetStatus

`func (o *Call) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Call) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Call) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Call) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTalkingTo

`func (o *Call) GetTalkingTo() TalkingTo`

GetTalkingTo returns the TalkingTo field if non-nil, zero value otherwise.

### GetTalkingToOk

`func (o *Call) GetTalkingToOk() (*TalkingTo, bool)`

GetTalkingToOk returns a tuple with the TalkingTo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTalkingTo

`func (o *Call) SetTalkingTo(v TalkingTo)`

SetTalkingTo sets TalkingTo field to given value.

### HasTalkingTo

`func (o *Call) HasTalkingTo() bool`

HasTalkingTo returns a boolean if a field has been set.

### GetUserUuid

`func (o *Call) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *Call) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *Call) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *Call) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
