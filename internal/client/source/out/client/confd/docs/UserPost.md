# UserPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**UserRelationAgent**](UserRelationAgent.md) |  | [optional]
**Auth** | Pointer to **map[string]interface{}** |  | [optional]
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
**Fallbacks** | Pointer to [**UserFallbacks**](UserFallbacks.md) |  | [optional]
**Firstname** | Pointer to **string** | User firstname | [optional]
**Forwards** | Pointer to [**UserForwards**](UserForwards.md) |  | [optional]
**Groups** | Pointer to **[]map[string]interface{}** |  | [optional]
**Id** | Pointer to **int32** | User ID | [optional] [readonly]
**Language** | Pointer to **string** | User language (e.g. \&quot;en_US\&quot;) | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Lines** | Pointer to [**[]UserPostRelationLine**](UserPostRelationLine.md) |  | [optional]
**MobilePhoneNumber** | Pointer to **string** | Phone number for the user’s mobile device | [optional]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**OnlineCallRecordEnabled** | Pointer to **bool** | Allow user to record a call by pressing *3 | [optional] [default to false]
**OutgoingCallerId** | Pointer to **string** | Name that appears on the phone when calling | [optional]
**Password** | Pointer to **string** | Password for connecting to the CTI (deprecated) | [optional]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]
**RingSeconds** | Pointer to **int32** | Number of seconds a user&#39;s phone will ring before falling back | [optional]
**SimultaneousCalls** | Pointer to **int32** | Number of allowed simultaneous calls on a user&#39;s phone | [optional]
**SupervisionEnabled** | Pointer to **bool** | Activate presence sharing in the accent client | [optional] [default to true]
**Switchboards** | Pointer to **[]map[string]interface{}** |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timezone** | Pointer to **string** | User timezone | [optional]
**Userfield** | Pointer to **string** | A custom field which purpose is left to the client. If the user has no userfield, then this field is an empty string. | [optional]
**Username** | Pointer to **string** | Username for connecting to the CTI (deprecated) | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Voicemail** | Pointer to [**Voicemail**](Voicemail.md) |  | [optional]
**Incalls** | Pointer to [**[]UserPostRelationIncallsIncallsInner**](UserPostRelationIncallsIncallsInner.md) |  | [optional]

## Methods

### NewUserPost

`func NewUserPost() *UserPost`

NewUserPost instantiates a new UserPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostWithDefaults

`func NewUserPostWithDefaults() *UserPost`

NewUserPostWithDefaults instantiates a new UserPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *UserPost) GetAgent() UserRelationAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *UserPost) GetAgentOk() (*UserRelationAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *UserPost) SetAgent(v UserRelationAgent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *UserPost) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAuth

`func (o *UserPost) GetAuth() map[string]interface{}`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UserPost) GetAuthOk() (*map[string]interface{}, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UserPost) SetAuth(v map[string]interface{})`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UserPost) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCallPermissionPassword

`func (o *UserPost) GetCallPermissionPassword() string`

GetCallPermissionPassword returns the CallPermissionPassword field if non-nil, zero value otherwise.

### GetCallPermissionPasswordOk

`func (o *UserPost) GetCallPermissionPasswordOk() (*string, bool)`

GetCallPermissionPasswordOk returns a tuple with the CallPermissionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissionPassword

`func (o *UserPost) SetCallPermissionPassword(v string)`

SetCallPermissionPassword sets CallPermissionPassword field to given value.

### HasCallPermissionPassword

`func (o *UserPost) HasCallPermissionPassword() bool`

HasCallPermissionPassword returns a boolean if a field has been set.

### GetCallRecordEnabled

`func (o *UserPost) GetCallRecordEnabled() bool`

GetCallRecordEnabled returns the CallRecordEnabled field if non-nil, zero value otherwise.

### GetCallRecordEnabledOk

`func (o *UserPost) GetCallRecordEnabledOk() (*bool, bool)`

GetCallRecordEnabledOk returns a tuple with the CallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordEnabled

`func (o *UserPost) SetCallRecordEnabled(v bool)`

SetCallRecordEnabled sets CallRecordEnabled field to given value.

### HasCallRecordEnabled

`func (o *UserPost) HasCallRecordEnabled() bool`

HasCallRecordEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingExternalEnabled

`func (o *UserPost) GetCallRecordIncomingExternalEnabled() bool`

GetCallRecordIncomingExternalEnabled returns the CallRecordIncomingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingExternalEnabledOk

`func (o *UserPost) GetCallRecordIncomingExternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingExternalEnabledOk returns a tuple with the CallRecordIncomingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingExternalEnabled

`func (o *UserPost) SetCallRecordIncomingExternalEnabled(v bool)`

SetCallRecordIncomingExternalEnabled sets CallRecordIncomingExternalEnabled field to given value.

### HasCallRecordIncomingExternalEnabled

`func (o *UserPost) HasCallRecordIncomingExternalEnabled() bool`

HasCallRecordIncomingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingInternalEnabled

`func (o *UserPost) GetCallRecordIncomingInternalEnabled() bool`

GetCallRecordIncomingInternalEnabled returns the CallRecordIncomingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingInternalEnabledOk

`func (o *UserPost) GetCallRecordIncomingInternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingInternalEnabledOk returns a tuple with the CallRecordIncomingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingInternalEnabled

`func (o *UserPost) SetCallRecordIncomingInternalEnabled(v bool)`

SetCallRecordIncomingInternalEnabled sets CallRecordIncomingInternalEnabled field to given value.

### HasCallRecordIncomingInternalEnabled

`func (o *UserPost) HasCallRecordIncomingInternalEnabled() bool`

HasCallRecordIncomingInternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingExternalEnabled

`func (o *UserPost) GetCallRecordOutgoingExternalEnabled() bool`

GetCallRecordOutgoingExternalEnabled returns the CallRecordOutgoingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingExternalEnabledOk

`func (o *UserPost) GetCallRecordOutgoingExternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingExternalEnabledOk returns a tuple with the CallRecordOutgoingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingExternalEnabled

`func (o *UserPost) SetCallRecordOutgoingExternalEnabled(v bool)`

SetCallRecordOutgoingExternalEnabled sets CallRecordOutgoingExternalEnabled field to given value.

### HasCallRecordOutgoingExternalEnabled

`func (o *UserPost) HasCallRecordOutgoingExternalEnabled() bool`

HasCallRecordOutgoingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingInternalEnabled

`func (o *UserPost) GetCallRecordOutgoingInternalEnabled() bool`

GetCallRecordOutgoingInternalEnabled returns the CallRecordOutgoingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingInternalEnabledOk

`func (o *UserPost) GetCallRecordOutgoingInternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingInternalEnabledOk returns a tuple with the CallRecordOutgoingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingInternalEnabled

`func (o *UserPost) SetCallRecordOutgoingInternalEnabled(v bool)`

SetCallRecordOutgoingInternalEnabled sets CallRecordOutgoingInternalEnabled field to given value.

### HasCallRecordOutgoingInternalEnabled

`func (o *UserPost) HasCallRecordOutgoingInternalEnabled() bool`

HasCallRecordOutgoingInternalEnabled returns a boolean if a field has been set.

### GetCallTransferEnabled

`func (o *UserPost) GetCallTransferEnabled() bool`

GetCallTransferEnabled returns the CallTransferEnabled field if non-nil, zero value otherwise.

### GetCallTransferEnabledOk

`func (o *UserPost) GetCallTransferEnabledOk() (*bool, bool)`

GetCallTransferEnabledOk returns a tuple with the CallTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTransferEnabled

`func (o *UserPost) SetCallTransferEnabled(v bool)`

SetCallTransferEnabled sets CallTransferEnabled field to given value.

### HasCallTransferEnabled

`func (o *UserPost) HasCallTransferEnabled() bool`

HasCallTransferEnabled returns a boolean if a field has been set.

### GetCallerId

`func (o *UserPost) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *UserPost) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *UserPost) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *UserPost) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDescription

`func (o *UserPost) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserPost) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserPost) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserPost) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDtmfHangupEnabled

`func (o *UserPost) GetDtmfHangupEnabled() bool`

GetDtmfHangupEnabled returns the DtmfHangupEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupEnabledOk

`func (o *UserPost) GetDtmfHangupEnabledOk() (*bool, bool)`

GetDtmfHangupEnabledOk returns a tuple with the DtmfHangupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupEnabled

`func (o *UserPost) SetDtmfHangupEnabled(v bool)`

SetDtmfHangupEnabled sets DtmfHangupEnabled field to given value.

### HasDtmfHangupEnabled

`func (o *UserPost) HasDtmfHangupEnabled() bool`

HasDtmfHangupEnabled returns a boolean if a field has been set.

### GetEmail

`func (o *UserPost) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPost) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPost) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPost) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *UserPost) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserPost) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserPost) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserPost) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbacks

`func (o *UserPost) GetFallbacks() UserFallbacks`

GetFallbacks returns the Fallbacks field if non-nil, zero value otherwise.

### GetFallbacksOk

`func (o *UserPost) GetFallbacksOk() (*UserFallbacks, bool)`

GetFallbacksOk returns a tuple with the Fallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbacks

`func (o *UserPost) SetFallbacks(v UserFallbacks)`

SetFallbacks sets Fallbacks field to given value.

### HasFallbacks

`func (o *UserPost) HasFallbacks() bool`

HasFallbacks returns a boolean if a field has been set.

### GetFirstname

`func (o *UserPost) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserPost) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserPost) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserPost) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetForwards

`func (o *UserPost) GetForwards() UserForwards`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *UserPost) GetForwardsOk() (*UserForwards, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *UserPost) SetForwards(v UserForwards)`

SetForwards sets Forwards field to given value.

### HasForwards

`func (o *UserPost) HasForwards() bool`

HasForwards returns a boolean if a field has been set.

### GetGroups

`func (o *UserPost) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserPost) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserPost) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserPost) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *UserPost) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPost) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPost) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPost) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *UserPost) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserPost) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserPost) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserPost) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastname

`func (o *UserPost) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserPost) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserPost) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserPost) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLines

`func (o *UserPost) GetLines() []UserPostRelationLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UserPost) GetLinesOk() (*[]UserPostRelationLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UserPost) SetLines(v []UserPostRelationLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UserPost) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *UserPost) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *UserPost) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *UserPost) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *UserPost) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *UserPost) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *UserPost) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *UserPost) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *UserPost) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetOnlineCallRecordEnabled

`func (o *UserPost) GetOnlineCallRecordEnabled() bool`

GetOnlineCallRecordEnabled returns the OnlineCallRecordEnabled field if non-nil, zero value otherwise.

### GetOnlineCallRecordEnabledOk

`func (o *UserPost) GetOnlineCallRecordEnabledOk() (*bool, bool)`

GetOnlineCallRecordEnabledOk returns a tuple with the OnlineCallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineCallRecordEnabled

`func (o *UserPost) SetOnlineCallRecordEnabled(v bool)`

SetOnlineCallRecordEnabled sets OnlineCallRecordEnabled field to given value.

### HasOnlineCallRecordEnabled

`func (o *UserPost) HasOnlineCallRecordEnabled() bool`

HasOnlineCallRecordEnabled returns a boolean if a field has been set.

### GetOutgoingCallerId

`func (o *UserPost) GetOutgoingCallerId() string`

GetOutgoingCallerId returns the OutgoingCallerId field if non-nil, zero value otherwise.

### GetOutgoingCallerIdOk

`func (o *UserPost) GetOutgoingCallerIdOk() (*string, bool)`

GetOutgoingCallerIdOk returns a tuple with the OutgoingCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallerId

`func (o *UserPost) SetOutgoingCallerId(v string)`

SetOutgoingCallerId sets OutgoingCallerId field to given value.

### HasOutgoingCallerId

`func (o *UserPost) HasOutgoingCallerId() bool`

HasOutgoingCallerId returns a boolean if a field has been set.

### GetPassword

`func (o *UserPost) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPost) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPost) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPost) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *UserPost) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *UserPost) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *UserPost) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *UserPost) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRingSeconds

`func (o *UserPost) GetRingSeconds() int32`

GetRingSeconds returns the RingSeconds field if non-nil, zero value otherwise.

### GetRingSecondsOk

`func (o *UserPost) GetRingSecondsOk() (*int32, bool)`

GetRingSecondsOk returns a tuple with the RingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSeconds

`func (o *UserPost) SetRingSeconds(v int32)`

SetRingSeconds sets RingSeconds field to given value.

### HasRingSeconds

`func (o *UserPost) HasRingSeconds() bool`

HasRingSeconds returns a boolean if a field has been set.

### GetSimultaneousCalls

`func (o *UserPost) GetSimultaneousCalls() int32`

GetSimultaneousCalls returns the SimultaneousCalls field if non-nil, zero value otherwise.

### GetSimultaneousCallsOk

`func (o *UserPost) GetSimultaneousCallsOk() (*int32, bool)`

GetSimultaneousCallsOk returns a tuple with the SimultaneousCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousCalls

`func (o *UserPost) SetSimultaneousCalls(v int32)`

SetSimultaneousCalls sets SimultaneousCalls field to given value.

### HasSimultaneousCalls

`func (o *UserPost) HasSimultaneousCalls() bool`

HasSimultaneousCalls returns a boolean if a field has been set.

### GetSupervisionEnabled

`func (o *UserPost) GetSupervisionEnabled() bool`

GetSupervisionEnabled returns the SupervisionEnabled field if non-nil, zero value otherwise.

### GetSupervisionEnabledOk

`func (o *UserPost) GetSupervisionEnabledOk() (*bool, bool)`

GetSupervisionEnabledOk returns a tuple with the SupervisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionEnabled

`func (o *UserPost) SetSupervisionEnabled(v bool)`

SetSupervisionEnabled sets SupervisionEnabled field to given value.

### HasSupervisionEnabled

`func (o *UserPost) HasSupervisionEnabled() bool`

HasSupervisionEnabled returns a boolean if a field has been set.

### GetSwitchboards

`func (o *UserPost) GetSwitchboards() []map[string]interface{}`

GetSwitchboards returns the Switchboards field if non-nil, zero value otherwise.

### GetSwitchboardsOk

`func (o *UserPost) GetSwitchboardsOk() (*[]map[string]interface{}, bool)`

GetSwitchboardsOk returns a tuple with the Switchboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboards

`func (o *UserPost) SetSwitchboards(v []map[string]interface{})`

SetSwitchboards sets Switchboards field to given value.

### HasSwitchboards

`func (o *UserPost) HasSwitchboards() bool`

HasSwitchboards returns a boolean if a field has been set.

### GetTenantUuid

`func (o *UserPost) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *UserPost) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *UserPost) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *UserPost) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *UserPost) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserPost) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserPost) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserPost) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserfield

`func (o *UserPost) GetUserfield() string`

GetUserfield returns the Userfield field if non-nil, zero value otherwise.

### GetUserfieldOk

`func (o *UserPost) GetUserfieldOk() (*string, bool)`

GetUserfieldOk returns a tuple with the Userfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfield

`func (o *UserPost) SetUserfield(v string)`

SetUserfield sets Userfield field to given value.

### HasUserfield

`func (o *UserPost) HasUserfield() bool`

HasUserfield returns a boolean if a field has been set.

### GetUsername

`func (o *UserPost) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPost) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPost) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserPost) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *UserPost) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserPost) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserPost) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserPost) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVoicemail

`func (o *UserPost) GetVoicemail() Voicemail`

GetVoicemail returns the Voicemail field if non-nil, zero value otherwise.

### GetVoicemailOk

`func (o *UserPost) GetVoicemailOk() (*Voicemail, bool)`

GetVoicemailOk returns a tuple with the Voicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemail

`func (o *UserPost) SetVoicemail(v Voicemail)`

SetVoicemail sets Voicemail field to given value.

### HasVoicemail

`func (o *UserPost) HasVoicemail() bool`

HasVoicemail returns a boolean if a field has been set.

### GetIncalls

`func (o *UserPost) GetIncalls() []UserPostRelationIncallsIncallsInner`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *UserPost) GetIncallsOk() (*[]UserPostRelationIncallsIncallsInner, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *UserPost) SetIncalls(v []UserPostRelationIncallsIncallsInner)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *UserPost) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
