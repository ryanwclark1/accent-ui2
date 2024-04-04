# DestinationType

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Application** | Pointer to [**DestinationApplicationApplication**](DestinationApplicationApplication.md) |  | [optional]
**Conference** | Pointer to [**DestinationConference**](DestinationConference.md) |  | [optional]
**Custom** | Pointer to [**DestinationCustom**](DestinationCustom.md) |  | [optional]
**Extension** | Pointer to [**DestinationExtension**](DestinationExtension.md) |  | [optional]
**Group** | Pointer to [**DestinationGroup**](DestinationGroup.md) |  | [optional]
**Hangup** | Pointer to [**DestinationHangupCause**](DestinationHangupCause.md) |  | [optional]
**Ivr** | Pointer to [**DestinationIVR**](DestinationIVR.md) |  | [optional]
**None** | Pointer to [**DestinationNone**](DestinationNone.md) |  | [optional]
**Outcall** | Pointer to [**DestinationOutcall**](DestinationOutcall.md) |  | [optional]
**Queue** | Pointer to [**DestinationQueue**](DestinationQueue.md) |  | [optional]
**Sound** | Pointer to [**DestinationSound**](DestinationSound.md) |  | [optional]
**Switchboard** | Pointer to [**DestinationSwitchboard**](DestinationSwitchboard.md) |  | [optional]
**User** | Pointer to [**DestinationUser**](DestinationUser.md) |  | [optional]
**Voicemail** | Pointer to [**DestinationVoicemail**](DestinationVoicemail.md) |  | [optional]

## Methods

### NewDestinationType

`func NewDestinationType() *DestinationType`

NewDestinationType instantiates a new DestinationType object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationTypeWithDefaults

`func NewDestinationTypeWithDefaults() *DestinationType`

NewDestinationTypeWithDefaults instantiates a new DestinationType object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetApplication

`func (o *DestinationType) GetApplication() DestinationApplicationApplication`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationType) GetApplicationOk() (*DestinationApplicationApplication, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationType) SetApplication(v DestinationApplicationApplication)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DestinationType) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetConference

`func (o *DestinationType) GetConference() DestinationConference`

GetConference returns the Conference field if non-nil, zero value otherwise.

### GetConferenceOk

`func (o *DestinationType) GetConferenceOk() (*DestinationConference, bool)`

GetConferenceOk returns a tuple with the Conference field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConference

`func (o *DestinationType) SetConference(v DestinationConference)`

SetConference sets Conference field to given value.

### HasConference

`func (o *DestinationType) HasConference() bool`

HasConference returns a boolean if a field has been set.

### GetCustom

`func (o *DestinationType) GetCustom() DestinationCustom`

GetCustom returns the Custom field if non-nil, zero value otherwise.

### GetCustomOk

`func (o *DestinationType) GetCustomOk() (*DestinationCustom, bool)`

GetCustomOk returns a tuple with the Custom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCustom

`func (o *DestinationType) SetCustom(v DestinationCustom)`

SetCustom sets Custom field to given value.

### HasCustom

`func (o *DestinationType) HasCustom() bool`

HasCustom returns a boolean if a field has been set.

### GetExtension

`func (o *DestinationType) GetExtension() DestinationExtension`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *DestinationType) GetExtensionOk() (*DestinationExtension, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *DestinationType) SetExtension(v DestinationExtension)`

SetExtension sets Extension field to given value.

### HasExtension

`func (o *DestinationType) HasExtension() bool`

HasExtension returns a boolean if a field has been set.

### GetGroup

`func (o *DestinationType) GetGroup() DestinationGroup`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *DestinationType) GetGroupOk() (*DestinationGroup, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *DestinationType) SetGroup(v DestinationGroup)`

SetGroup sets Group field to given value.

### HasGroup

`func (o *DestinationType) HasGroup() bool`

HasGroup returns a boolean if a field has been set.

### GetHangup

`func (o *DestinationType) GetHangup() DestinationHangupCause`

GetHangup returns the Hangup field if non-nil, zero value otherwise.

### GetHangupOk

`func (o *DestinationType) GetHangupOk() (*DestinationHangupCause, bool)`

GetHangupOk returns a tuple with the Hangup field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHangup

`func (o *DestinationType) SetHangup(v DestinationHangupCause)`

SetHangup sets Hangup field to given value.

### HasHangup

`func (o *DestinationType) HasHangup() bool`

HasHangup returns a boolean if a field has been set.

### GetIvr

`func (o *DestinationType) GetIvr() DestinationIVR`

GetIvr returns the Ivr field if non-nil, zero value otherwise.

### GetIvrOk

`func (o *DestinationType) GetIvrOk() (*DestinationIVR, bool)`

GetIvrOk returns a tuple with the Ivr field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIvr

`func (o *DestinationType) SetIvr(v DestinationIVR)`

SetIvr sets Ivr field to given value.

### HasIvr

`func (o *DestinationType) HasIvr() bool`

HasIvr returns a boolean if a field has been set.

### GetNone

`func (o *DestinationType) GetNone() DestinationNone`

GetNone returns the None field if non-nil, zero value otherwise.

### GetNoneOk

`func (o *DestinationType) GetNoneOk() (*DestinationNone, bool)`

GetNoneOk returns a tuple with the None field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNone

`func (o *DestinationType) SetNone(v DestinationNone)`

SetNone sets None field to given value.

### HasNone

`func (o *DestinationType) HasNone() bool`

HasNone returns a boolean if a field has been set.

### GetOutcall

`func (o *DestinationType) GetOutcall() DestinationOutcall`

GetOutcall returns the Outcall field if non-nil, zero value otherwise.

### GetOutcallOk

`func (o *DestinationType) GetOutcallOk() (*DestinationOutcall, bool)`

GetOutcallOk returns a tuple with the Outcall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcall

`func (o *DestinationType) SetOutcall(v DestinationOutcall)`

SetOutcall sets Outcall field to given value.

### HasOutcall

`func (o *DestinationType) HasOutcall() bool`

HasOutcall returns a boolean if a field has been set.

### GetQueue

`func (o *DestinationType) GetQueue() DestinationQueue`

GetQueue returns the Queue field if non-nil, zero value otherwise.

### GetQueueOk

`func (o *DestinationType) GetQueueOk() (*DestinationQueue, bool)`

GetQueueOk returns a tuple with the Queue field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueue

`func (o *DestinationType) SetQueue(v DestinationQueue)`

SetQueue sets Queue field to given value.

### HasQueue

`func (o *DestinationType) HasQueue() bool`

HasQueue returns a boolean if a field has been set.

### GetSound

`func (o *DestinationType) GetSound() DestinationSound`

GetSound returns the Sound field if non-nil, zero value otherwise.

### GetSoundOk

`func (o *DestinationType) GetSoundOk() (*DestinationSound, bool)`

GetSoundOk returns a tuple with the Sound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSound

`func (o *DestinationType) SetSound(v DestinationSound)`

SetSound sets Sound field to given value.

### HasSound

`func (o *DestinationType) HasSound() bool`

HasSound returns a boolean if a field has been set.

### GetSwitchboard

`func (o *DestinationType) GetSwitchboard() DestinationSwitchboard`

GetSwitchboard returns the Switchboard field if non-nil, zero value otherwise.

### GetSwitchboardOk

`func (o *DestinationType) GetSwitchboardOk() (*DestinationSwitchboard, bool)`

GetSwitchboardOk returns a tuple with the Switchboard field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboard

`func (o *DestinationType) SetSwitchboard(v DestinationSwitchboard)`

SetSwitchboard sets Switchboard field to given value.

### HasSwitchboard

`func (o *DestinationType) HasSwitchboard() bool`

HasSwitchboard returns a boolean if a field has been set.

### GetUser

`func (o *DestinationType) GetUser() DestinationUser`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *DestinationType) GetUserOk() (*DestinationUser, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *DestinationType) SetUser(v DestinationUser)`

SetUser sets User field to given value.

### HasUser

`func (o *DestinationType) HasUser() bool`

HasUser returns a boolean if a field has been set.

### GetVoicemail

`func (o *DestinationType) GetVoicemail() DestinationVoicemail`

GetVoicemail returns the Voicemail field if non-nil, zero value otherwise.

### GetVoicemailOk

`func (o *DestinationType) GetVoicemailOk() (*DestinationVoicemail, bool)`

GetVoicemailOk returns a tuple with the Voicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemail

`func (o *DestinationType) SetVoicemail(v DestinationVoicemail)`

SetVoicemail sets Voicemail field to given value.

### HasVoicemail

`func (o *DestinationType) HasVoicemail() bool`

HasVoicemail returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
