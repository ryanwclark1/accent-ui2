# User

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**CallPermissionPassword** | Pointer to **string** | Overwrite all passwords set in call permissions associated to the user | [optional]
**CallRecordEnabled** | Pointer to **bool** | Record all calls made by this user (deprecated). If set, all &#x60;call_record_*_enabled&#x60; will be affected | [optional] [default to false]
**CallRecordIncomingExternalEnabled** | Pointer to **bool** | Record all external calls received by this user. | [optional] [default to false]
**CallRecordIncomingInternalEnabled** | Pointer to **bool** | Record all internal calls received by this user. | [optional] [default to false]
**CallRecordOutgoingExternalEnabled** | Pointer to **bool** | Record all external calls made by this user | [optional] [default to false]
**CallRecordOutgoingInternalEnabled** | Pointer to **bool** | Record all internal calls received by this user | [optional] [default to false]
**CallTransferEnabled** | Pointer to **bool** | Authorize call transfers | [optional] [default to false]
**CallerId** | Pointer to **string** | Name that appears on the phone when calling. Formatted as \&quot;Firstname Lastname\&quot; &lt; number &gt; | [optional]
**Description** | Pointer to **string** | Additional information about the user | [optional]
**DtmfHangupEnabled** | Pointer to **bool** | Authorize to hangup with DTMF | [optional] [default to false]
**Email** | Pointer to **string** | User&#39;s email. Used for directories (i.e. by accent-dird) | [optional]
**Enabled** | Pointer to **bool** | Enable/Disable the user | [optional] [default to true]
**Id** | Pointer to **int32** | User ID | [optional] [readonly]
**Language** | Pointer to **string** | User language (e.g. \&quot;en_US\&quot;) | [optional]
**MobilePhoneNumber** | Pointer to **string** | Phone number for the user’s mobile device | [optional]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**OnlineCallRecordEnabled** | Pointer to **bool** | Allow user to record a call by pressing *3 | [optional] [default to false]
**OutgoingCallerId** | Pointer to **string** | Name that appears on the phone when calling | [optional]
**Password** | Pointer to **string** | Password for connecting to the CTI (deprecated) | [optional]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]
**RingSeconds** | Pointer to **int32** | Number of seconds a user&#39;s phone will ring before falling back | [optional]
**SimultaneousCalls** | Pointer to **int32** | Number of allowed simultaneous calls on a user&#39;s phone | [optional]
**SupervisionEnabled** | Pointer to **bool** | Activate presence sharing in the accent client | [optional] [default to true]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timezone** | Pointer to **string** | User timezone | [optional]
**Userfield** | Pointer to **string** | A custom field which purpose is left to the client. If the user has no userfield, then this field is an empty string. | [optional]
**Username** | Pointer to **string** | Username for connecting to the CTI (deprecated) | [optional]
**Agent** | Pointer to [**AgentRelationBase**](AgentRelationBase.md) |  | [optional]
**Fallbacks** | Pointer to [**UserFallbacks**](UserFallbacks.md) |  | [optional]
**Forwards** | Pointer to [**UserForwards**](UserForwards.md) |  | [optional]
**Groups** | Pointer to [**[]GroupRelationBase**](GroupRelationBase.md) |  | [optional] [readonly]
**Incalls** | Pointer to [**[]UserRelationIncall**](UserRelationIncall.md) |  | [optional] [readonly]
**Lines** | Pointer to [**[]UserRelationLine**](UserRelationLine.md) |  | [optional] [readonly]
**Services** | Pointer to [**UserServices**](UserServices.md) |  | [optional]
**Switchboards** | Pointer to [**[]SwitchboardRelationBase**](SwitchboardRelationBase.md) |  | [optional] [readonly]
**Voicemail** | Pointer to [**VoicemailRelationBase**](VoicemailRelationBase.md) |  | [optional]
**Queues** | Pointer to [**[]QueueRelationBase**](QueueRelationBase.md) |  | [optional] [readonly]
**CallPickupTargetUsers** | Pointer to [**[]UserRelationBase**](UserRelationBase.md) |  | [optional]

## Methods

### NewUser

`func NewUser() *User`

NewUser instantiates a new User object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserWithDefaults

`func NewUserWithDefaults() *User`

NewUserWithDefaults instantiates a new User object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *User) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *User) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *User) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *User) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *User) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *User) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *User) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *User) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *User) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *User) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *User) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *User) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCallPermissionPassword

`func (o *User) GetCallPermissionPassword() string`

GetCallPermissionPassword returns the CallPermissionPassword field if non-nil, zero value otherwise.

### GetCallPermissionPasswordOk

`func (o *User) GetCallPermissionPasswordOk() (*string, bool)`

GetCallPermissionPasswordOk returns a tuple with the CallPermissionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissionPassword

`func (o *User) SetCallPermissionPassword(v string)`

SetCallPermissionPassword sets CallPermissionPassword field to given value.

### HasCallPermissionPassword

`func (o *User) HasCallPermissionPassword() bool`

HasCallPermissionPassword returns a boolean if a field has been set.

### GetCallRecordEnabled

`func (o *User) GetCallRecordEnabled() bool`

GetCallRecordEnabled returns the CallRecordEnabled field if non-nil, zero value otherwise.

### GetCallRecordEnabledOk

`func (o *User) GetCallRecordEnabledOk() (*bool, bool)`

GetCallRecordEnabledOk returns a tuple with the CallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordEnabled

`func (o *User) SetCallRecordEnabled(v bool)`

SetCallRecordEnabled sets CallRecordEnabled field to given value.

### HasCallRecordEnabled

`func (o *User) HasCallRecordEnabled() bool`

HasCallRecordEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingExternalEnabled

`func (o *User) GetCallRecordIncomingExternalEnabled() bool`

GetCallRecordIncomingExternalEnabled returns the CallRecordIncomingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingExternalEnabledOk

`func (o *User) GetCallRecordIncomingExternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingExternalEnabledOk returns a tuple with the CallRecordIncomingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingExternalEnabled

`func (o *User) SetCallRecordIncomingExternalEnabled(v bool)`

SetCallRecordIncomingExternalEnabled sets CallRecordIncomingExternalEnabled field to given value.

### HasCallRecordIncomingExternalEnabled

`func (o *User) HasCallRecordIncomingExternalEnabled() bool`

HasCallRecordIncomingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingInternalEnabled

`func (o *User) GetCallRecordIncomingInternalEnabled() bool`

GetCallRecordIncomingInternalEnabled returns the CallRecordIncomingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingInternalEnabledOk

`func (o *User) GetCallRecordIncomingInternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingInternalEnabledOk returns a tuple with the CallRecordIncomingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingInternalEnabled

`func (o *User) SetCallRecordIncomingInternalEnabled(v bool)`

SetCallRecordIncomingInternalEnabled sets CallRecordIncomingInternalEnabled field to given value.

### HasCallRecordIncomingInternalEnabled

`func (o *User) HasCallRecordIncomingInternalEnabled() bool`

HasCallRecordIncomingInternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingExternalEnabled

`func (o *User) GetCallRecordOutgoingExternalEnabled() bool`

GetCallRecordOutgoingExternalEnabled returns the CallRecordOutgoingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingExternalEnabledOk

`func (o *User) GetCallRecordOutgoingExternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingExternalEnabledOk returns a tuple with the CallRecordOutgoingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingExternalEnabled

`func (o *User) SetCallRecordOutgoingExternalEnabled(v bool)`

SetCallRecordOutgoingExternalEnabled sets CallRecordOutgoingExternalEnabled field to given value.

### HasCallRecordOutgoingExternalEnabled

`func (o *User) HasCallRecordOutgoingExternalEnabled() bool`

HasCallRecordOutgoingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingInternalEnabled

`func (o *User) GetCallRecordOutgoingInternalEnabled() bool`

GetCallRecordOutgoingInternalEnabled returns the CallRecordOutgoingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingInternalEnabledOk

`func (o *User) GetCallRecordOutgoingInternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingInternalEnabledOk returns a tuple with the CallRecordOutgoingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingInternalEnabled

`func (o *User) SetCallRecordOutgoingInternalEnabled(v bool)`

SetCallRecordOutgoingInternalEnabled sets CallRecordOutgoingInternalEnabled field to given value.

### HasCallRecordOutgoingInternalEnabled

`func (o *User) HasCallRecordOutgoingInternalEnabled() bool`

HasCallRecordOutgoingInternalEnabled returns a boolean if a field has been set.

### GetCallTransferEnabled

`func (o *User) GetCallTransferEnabled() bool`

GetCallTransferEnabled returns the CallTransferEnabled field if non-nil, zero value otherwise.

### GetCallTransferEnabledOk

`func (o *User) GetCallTransferEnabledOk() (*bool, bool)`

GetCallTransferEnabledOk returns a tuple with the CallTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTransferEnabled

`func (o *User) SetCallTransferEnabled(v bool)`

SetCallTransferEnabled sets CallTransferEnabled field to given value.

### HasCallTransferEnabled

`func (o *User) HasCallTransferEnabled() bool`

HasCallTransferEnabled returns a boolean if a field has been set.

### GetCallerId

`func (o *User) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *User) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *User) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *User) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDescription

`func (o *User) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *User) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *User) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *User) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDtmfHangupEnabled

`func (o *User) GetDtmfHangupEnabled() bool`

GetDtmfHangupEnabled returns the DtmfHangupEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupEnabledOk

`func (o *User) GetDtmfHangupEnabledOk() (*bool, bool)`

GetDtmfHangupEnabledOk returns a tuple with the DtmfHangupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupEnabled

`func (o *User) SetDtmfHangupEnabled(v bool)`

SetDtmfHangupEnabled sets DtmfHangupEnabled field to given value.

### HasDtmfHangupEnabled

`func (o *User) HasDtmfHangupEnabled() bool`

HasDtmfHangupEnabled returns a boolean if a field has been set.

### GetEmail

`func (o *User) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *User) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *User) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *User) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *User) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *User) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *User) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *User) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *User) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *User) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *User) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *User) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *User) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *User) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *User) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *User) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *User) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *User) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *User) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *User) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *User) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *User) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *User) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *User) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetOnlineCallRecordEnabled

`func (o *User) GetOnlineCallRecordEnabled() bool`

GetOnlineCallRecordEnabled returns the OnlineCallRecordEnabled field if non-nil, zero value otherwise.

### GetOnlineCallRecordEnabledOk

`func (o *User) GetOnlineCallRecordEnabledOk() (*bool, bool)`

GetOnlineCallRecordEnabledOk returns a tuple with the OnlineCallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineCallRecordEnabled

`func (o *User) SetOnlineCallRecordEnabled(v bool)`

SetOnlineCallRecordEnabled sets OnlineCallRecordEnabled field to given value.

### HasOnlineCallRecordEnabled

`func (o *User) HasOnlineCallRecordEnabled() bool`

HasOnlineCallRecordEnabled returns a boolean if a field has been set.

### GetOutgoingCallerId

`func (o *User) GetOutgoingCallerId() string`

GetOutgoingCallerId returns the OutgoingCallerId field if non-nil, zero value otherwise.

### GetOutgoingCallerIdOk

`func (o *User) GetOutgoingCallerIdOk() (*string, bool)`

GetOutgoingCallerIdOk returns a tuple with the OutgoingCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallerId

`func (o *User) SetOutgoingCallerId(v string)`

SetOutgoingCallerId sets OutgoingCallerId field to given value.

### HasOutgoingCallerId

`func (o *User) HasOutgoingCallerId() bool`

HasOutgoingCallerId returns a boolean if a field has been set.

### GetPassword

`func (o *User) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *User) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *User) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *User) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *User) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *User) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *User) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *User) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRingSeconds

`func (o *User) GetRingSeconds() int32`

GetRingSeconds returns the RingSeconds field if non-nil, zero value otherwise.

### GetRingSecondsOk

`func (o *User) GetRingSecondsOk() (*int32, bool)`

GetRingSecondsOk returns a tuple with the RingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSeconds

`func (o *User) SetRingSeconds(v int32)`

SetRingSeconds sets RingSeconds field to given value.

### HasRingSeconds

`func (o *User) HasRingSeconds() bool`

HasRingSeconds returns a boolean if a field has been set.

### GetSimultaneousCalls

`func (o *User) GetSimultaneousCalls() int32`

GetSimultaneousCalls returns the SimultaneousCalls field if non-nil, zero value otherwise.

### GetSimultaneousCallsOk

`func (o *User) GetSimultaneousCallsOk() (*int32, bool)`

GetSimultaneousCallsOk returns a tuple with the SimultaneousCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousCalls

`func (o *User) SetSimultaneousCalls(v int32)`

SetSimultaneousCalls sets SimultaneousCalls field to given value.

### HasSimultaneousCalls

`func (o *User) HasSimultaneousCalls() bool`

HasSimultaneousCalls returns a boolean if a field has been set.

### GetSupervisionEnabled

`func (o *User) GetSupervisionEnabled() bool`

GetSupervisionEnabled returns the SupervisionEnabled field if non-nil, zero value otherwise.

### GetSupervisionEnabledOk

`func (o *User) GetSupervisionEnabledOk() (*bool, bool)`

GetSupervisionEnabledOk returns a tuple with the SupervisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionEnabled

`func (o *User) SetSupervisionEnabled(v bool)`

SetSupervisionEnabled sets SupervisionEnabled field to given value.

### HasSupervisionEnabled

`func (o *User) HasSupervisionEnabled() bool`

HasSupervisionEnabled returns a boolean if a field has been set.

### GetTenantUuid

`func (o *User) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *User) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *User) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *User) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *User) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *User) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *User) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *User) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserfield

`func (o *User) GetUserfield() string`

GetUserfield returns the Userfield field if non-nil, zero value otherwise.

### GetUserfieldOk

`func (o *User) GetUserfieldOk() (*string, bool)`

GetUserfieldOk returns a tuple with the Userfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfield

`func (o *User) SetUserfield(v string)`

SetUserfield sets Userfield field to given value.

### HasUserfield

`func (o *User) HasUserfield() bool`

HasUserfield returns a boolean if a field has been set.

### GetUsername

`func (o *User) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *User) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *User) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *User) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetAgent

`func (o *User) GetAgent() AgentRelationBase`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *User) GetAgentOk() (*AgentRelationBase, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *User) SetAgent(v AgentRelationBase)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *User) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetFallbacks

`func (o *User) GetFallbacks() UserFallbacks`

GetFallbacks returns the Fallbacks field if non-nil, zero value otherwise.

### GetFallbacksOk

`func (o *User) GetFallbacksOk() (*UserFallbacks, bool)`

GetFallbacksOk returns a tuple with the Fallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbacks

`func (o *User) SetFallbacks(v UserFallbacks)`

SetFallbacks sets Fallbacks field to given value.

### HasFallbacks

`func (o *User) HasFallbacks() bool`

HasFallbacks returns a boolean if a field has been set.

### GetForwards

`func (o *User) GetForwards() UserForwards`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *User) GetForwardsOk() (*UserForwards, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *User) SetForwards(v UserForwards)`

SetForwards sets Forwards field to given value.

### HasForwards

`func (o *User) HasForwards() bool`

HasForwards returns a boolean if a field has been set.

### GetGroups

`func (o *User) GetGroups() []GroupRelationBase`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *User) GetGroupsOk() (*[]GroupRelationBase, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *User) SetGroups(v []GroupRelationBase)`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *User) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetIncalls

`func (o *User) GetIncalls() []UserRelationIncall`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *User) GetIncallsOk() (*[]UserRelationIncall, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *User) SetIncalls(v []UserRelationIncall)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *User) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetLines

`func (o *User) GetLines() []UserRelationLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *User) GetLinesOk() (*[]UserRelationLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *User) SetLines(v []UserRelationLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *User) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetServices

`func (o *User) GetServices() UserServices`

GetServices returns the Services field if non-nil, zero value otherwise.

### GetServicesOk

`func (o *User) GetServicesOk() (*UserServices, bool)`

GetServicesOk returns a tuple with the Services field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServices

`func (o *User) SetServices(v UserServices)`

SetServices sets Services field to given value.

### HasServices

`func (o *User) HasServices() bool`

HasServices returns a boolean if a field has been set.

### GetSwitchboards

`func (o *User) GetSwitchboards() []SwitchboardRelationBase`

GetSwitchboards returns the Switchboards field if non-nil, zero value otherwise.

### GetSwitchboardsOk

`func (o *User) GetSwitchboardsOk() (*[]SwitchboardRelationBase, bool)`

GetSwitchboardsOk returns a tuple with the Switchboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboards

`func (o *User) SetSwitchboards(v []SwitchboardRelationBase)`

SetSwitchboards sets Switchboards field to given value.

### HasSwitchboards

`func (o *User) HasSwitchboards() bool`

HasSwitchboards returns a boolean if a field has been set.

### GetVoicemail

`func (o *User) GetVoicemail() VoicemailRelationBase`

GetVoicemail returns the Voicemail field if non-nil, zero value otherwise.

### GetVoicemailOk

`func (o *User) GetVoicemailOk() (*VoicemailRelationBase, bool)`

GetVoicemailOk returns a tuple with the Voicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemail

`func (o *User) SetVoicemail(v VoicemailRelationBase)`

SetVoicemail sets Voicemail field to given value.

### HasVoicemail

`func (o *User) HasVoicemail() bool`

HasVoicemail returns a boolean if a field has been set.

### GetQueues

`func (o *User) GetQueues() []QueueRelationBase`

GetQueues returns the Queues field if non-nil, zero value otherwise.

### GetQueuesOk

`func (o *User) GetQueuesOk() (*[]QueueRelationBase, bool)`

GetQueuesOk returns a tuple with the Queues field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueues

`func (o *User) SetQueues(v []QueueRelationBase)`

SetQueues sets Queues field to given value.

### HasQueues

`func (o *User) HasQueues() bool`

HasQueues returns a boolean if a field has been set.

### GetCallPickupTargetUsers

`func (o *User) GetCallPickupTargetUsers() []UserRelationBase`

GetCallPickupTargetUsers returns the CallPickupTargetUsers field if non-nil, zero value otherwise.

### GetCallPickupTargetUsersOk

`func (o *User) GetCallPickupTargetUsersOk() (*[]UserRelationBase, bool)`

GetCallPickupTargetUsersOk returns a tuple with the CallPickupTargetUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPickupTargetUsers

`func (o *User) SetCallPickupTargetUsers(v []UserRelationBase)`

SetCallPickupTargetUsers sets CallPickupTargetUsers field to given value.

### HasCallPickupTargetUsers

`func (o *User) HasCallPickupTargetUsers() bool`

HasCallPickupTargetUsers returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
