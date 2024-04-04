# UserPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Agent** | Pointer to [**UserRelationAgent**](UserRelationAgent.md) |  | [optional]
**Auth** | Pointer to [**AuthUserPostAuth**](AuthUserPostAuth.md) |  | [optional]
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

### NewUserPostResponse

`func NewUserPostResponse() *UserPostResponse`

NewUserPostResponse instantiates a new UserPostResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostResponseWithDefaults

`func NewUserPostResponseWithDefaults() *UserPostResponse`

NewUserPostResponseWithDefaults instantiates a new UserPostResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgent

`func (o *UserPostResponse) GetAgent() UserRelationAgent`

GetAgent returns the Agent field if non-nil, zero value otherwise.

### GetAgentOk

`func (o *UserPostResponse) GetAgentOk() (*UserRelationAgent, bool)`

GetAgentOk returns a tuple with the Agent field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgent

`func (o *UserPostResponse) SetAgent(v UserRelationAgent)`

SetAgent sets Agent field to given value.

### HasAgent

`func (o *UserPostResponse) HasAgent() bool`

HasAgent returns a boolean if a field has been set.

### GetAuth

`func (o *UserPostResponse) GetAuth() AuthUserPostAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *UserPostResponse) GetAuthOk() (*AuthUserPostAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *UserPostResponse) SetAuth(v AuthUserPostAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *UserPostResponse) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

### GetCallPermissionPassword

`func (o *UserPostResponse) GetCallPermissionPassword() string`

GetCallPermissionPassword returns the CallPermissionPassword field if non-nil, zero value otherwise.

### GetCallPermissionPasswordOk

`func (o *UserPostResponse) GetCallPermissionPasswordOk() (*string, bool)`

GetCallPermissionPasswordOk returns a tuple with the CallPermissionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissionPassword

`func (o *UserPostResponse) SetCallPermissionPassword(v string)`

SetCallPermissionPassword sets CallPermissionPassword field to given value.

### HasCallPermissionPassword

`func (o *UserPostResponse) HasCallPermissionPassword() bool`

HasCallPermissionPassword returns a boolean if a field has been set.

### GetCallRecordEnabled

`func (o *UserPostResponse) GetCallRecordEnabled() bool`

GetCallRecordEnabled returns the CallRecordEnabled field if non-nil, zero value otherwise.

### GetCallRecordEnabledOk

`func (o *UserPostResponse) GetCallRecordEnabledOk() (*bool, bool)`

GetCallRecordEnabledOk returns a tuple with the CallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordEnabled

`func (o *UserPostResponse) SetCallRecordEnabled(v bool)`

SetCallRecordEnabled sets CallRecordEnabled field to given value.

### HasCallRecordEnabled

`func (o *UserPostResponse) HasCallRecordEnabled() bool`

HasCallRecordEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingExternalEnabled

`func (o *UserPostResponse) GetCallRecordIncomingExternalEnabled() bool`

GetCallRecordIncomingExternalEnabled returns the CallRecordIncomingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingExternalEnabledOk

`func (o *UserPostResponse) GetCallRecordIncomingExternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingExternalEnabledOk returns a tuple with the CallRecordIncomingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingExternalEnabled

`func (o *UserPostResponse) SetCallRecordIncomingExternalEnabled(v bool)`

SetCallRecordIncomingExternalEnabled sets CallRecordIncomingExternalEnabled field to given value.

### HasCallRecordIncomingExternalEnabled

`func (o *UserPostResponse) HasCallRecordIncomingExternalEnabled() bool`

HasCallRecordIncomingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingInternalEnabled

`func (o *UserPostResponse) GetCallRecordIncomingInternalEnabled() bool`

GetCallRecordIncomingInternalEnabled returns the CallRecordIncomingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingInternalEnabledOk

`func (o *UserPostResponse) GetCallRecordIncomingInternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingInternalEnabledOk returns a tuple with the CallRecordIncomingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingInternalEnabled

`func (o *UserPostResponse) SetCallRecordIncomingInternalEnabled(v bool)`

SetCallRecordIncomingInternalEnabled sets CallRecordIncomingInternalEnabled field to given value.

### HasCallRecordIncomingInternalEnabled

`func (o *UserPostResponse) HasCallRecordIncomingInternalEnabled() bool`

HasCallRecordIncomingInternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingExternalEnabled

`func (o *UserPostResponse) GetCallRecordOutgoingExternalEnabled() bool`

GetCallRecordOutgoingExternalEnabled returns the CallRecordOutgoingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingExternalEnabledOk

`func (o *UserPostResponse) GetCallRecordOutgoingExternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingExternalEnabledOk returns a tuple with the CallRecordOutgoingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingExternalEnabled

`func (o *UserPostResponse) SetCallRecordOutgoingExternalEnabled(v bool)`

SetCallRecordOutgoingExternalEnabled sets CallRecordOutgoingExternalEnabled field to given value.

### HasCallRecordOutgoingExternalEnabled

`func (o *UserPostResponse) HasCallRecordOutgoingExternalEnabled() bool`

HasCallRecordOutgoingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingInternalEnabled

`func (o *UserPostResponse) GetCallRecordOutgoingInternalEnabled() bool`

GetCallRecordOutgoingInternalEnabled returns the CallRecordOutgoingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingInternalEnabledOk

`func (o *UserPostResponse) GetCallRecordOutgoingInternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingInternalEnabledOk returns a tuple with the CallRecordOutgoingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingInternalEnabled

`func (o *UserPostResponse) SetCallRecordOutgoingInternalEnabled(v bool)`

SetCallRecordOutgoingInternalEnabled sets CallRecordOutgoingInternalEnabled field to given value.

### HasCallRecordOutgoingInternalEnabled

`func (o *UserPostResponse) HasCallRecordOutgoingInternalEnabled() bool`

HasCallRecordOutgoingInternalEnabled returns a boolean if a field has been set.

### GetCallTransferEnabled

`func (o *UserPostResponse) GetCallTransferEnabled() bool`

GetCallTransferEnabled returns the CallTransferEnabled field if non-nil, zero value otherwise.

### GetCallTransferEnabledOk

`func (o *UserPostResponse) GetCallTransferEnabledOk() (*bool, bool)`

GetCallTransferEnabledOk returns a tuple with the CallTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTransferEnabled

`func (o *UserPostResponse) SetCallTransferEnabled(v bool)`

SetCallTransferEnabled sets CallTransferEnabled field to given value.

### HasCallTransferEnabled

`func (o *UserPostResponse) HasCallTransferEnabled() bool`

HasCallTransferEnabled returns a boolean if a field has been set.

### GetCallerId

`func (o *UserPostResponse) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *UserPostResponse) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *UserPostResponse) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *UserPostResponse) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDescription

`func (o *UserPostResponse) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *UserPostResponse) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *UserPostResponse) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *UserPostResponse) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDtmfHangupEnabled

`func (o *UserPostResponse) GetDtmfHangupEnabled() bool`

GetDtmfHangupEnabled returns the DtmfHangupEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupEnabledOk

`func (o *UserPostResponse) GetDtmfHangupEnabledOk() (*bool, bool)`

GetDtmfHangupEnabledOk returns a tuple with the DtmfHangupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupEnabled

`func (o *UserPostResponse) SetDtmfHangupEnabled(v bool)`

SetDtmfHangupEnabled sets DtmfHangupEnabled field to given value.

### HasDtmfHangupEnabled

`func (o *UserPostResponse) HasDtmfHangupEnabled() bool`

HasDtmfHangupEnabled returns a boolean if a field has been set.

### GetEmail

`func (o *UserPostResponse) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *UserPostResponse) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *UserPostResponse) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *UserPostResponse) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *UserPostResponse) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserPostResponse) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserPostResponse) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserPostResponse) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFallbacks

`func (o *UserPostResponse) GetFallbacks() UserFallbacks`

GetFallbacks returns the Fallbacks field if non-nil, zero value otherwise.

### GetFallbacksOk

`func (o *UserPostResponse) GetFallbacksOk() (*UserFallbacks, bool)`

GetFallbacksOk returns a tuple with the Fallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbacks

`func (o *UserPostResponse) SetFallbacks(v UserFallbacks)`

SetFallbacks sets Fallbacks field to given value.

### HasFallbacks

`func (o *UserPostResponse) HasFallbacks() bool`

HasFallbacks returns a boolean if a field has been set.

### GetFirstname

`func (o *UserPostResponse) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserPostResponse) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserPostResponse) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserPostResponse) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetForwards

`func (o *UserPostResponse) GetForwards() UserForwards`

GetForwards returns the Forwards field if non-nil, zero value otherwise.

### GetForwardsOk

`func (o *UserPostResponse) GetForwardsOk() (*UserForwards, bool)`

GetForwardsOk returns a tuple with the Forwards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetForwards

`func (o *UserPostResponse) SetForwards(v UserForwards)`

SetForwards sets Forwards field to given value.

### HasForwards

`func (o *UserPostResponse) HasForwards() bool`

HasForwards returns a boolean if a field has been set.

### GetGroups

`func (o *UserPostResponse) GetGroups() []map[string]interface{}`

GetGroups returns the Groups field if non-nil, zero value otherwise.

### GetGroupsOk

`func (o *UserPostResponse) GetGroupsOk() (*[]map[string]interface{}, bool)`

GetGroupsOk returns a tuple with the Groups field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroups

`func (o *UserPostResponse) SetGroups(v []map[string]interface{})`

SetGroups sets Groups field to given value.

### HasGroups

`func (o *UserPostResponse) HasGroups() bool`

HasGroups returns a boolean if a field has been set.

### GetId

`func (o *UserPostResponse) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPostResponse) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPostResponse) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPostResponse) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *UserPostResponse) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *UserPostResponse) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *UserPostResponse) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *UserPostResponse) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetLastname

`func (o *UserPostResponse) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserPostResponse) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserPostResponse) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserPostResponse) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetLines

`func (o *UserPostResponse) GetLines() []UserPostRelationLine`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *UserPostResponse) GetLinesOk() (*[]UserPostRelationLine, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *UserPostResponse) SetLines(v []UserPostRelationLine)`

SetLines sets Lines field to given value.

### HasLines

`func (o *UserPostResponse) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *UserPostResponse) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *UserPostResponse) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *UserPostResponse) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *UserPostResponse) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *UserPostResponse) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *UserPostResponse) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *UserPostResponse) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *UserPostResponse) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetOnlineCallRecordEnabled

`func (o *UserPostResponse) GetOnlineCallRecordEnabled() bool`

GetOnlineCallRecordEnabled returns the OnlineCallRecordEnabled field if non-nil, zero value otherwise.

### GetOnlineCallRecordEnabledOk

`func (o *UserPostResponse) GetOnlineCallRecordEnabledOk() (*bool, bool)`

GetOnlineCallRecordEnabledOk returns a tuple with the OnlineCallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineCallRecordEnabled

`func (o *UserPostResponse) SetOnlineCallRecordEnabled(v bool)`

SetOnlineCallRecordEnabled sets OnlineCallRecordEnabled field to given value.

### HasOnlineCallRecordEnabled

`func (o *UserPostResponse) HasOnlineCallRecordEnabled() bool`

HasOnlineCallRecordEnabled returns a boolean if a field has been set.

### GetOutgoingCallerId

`func (o *UserPostResponse) GetOutgoingCallerId() string`

GetOutgoingCallerId returns the OutgoingCallerId field if non-nil, zero value otherwise.

### GetOutgoingCallerIdOk

`func (o *UserPostResponse) GetOutgoingCallerIdOk() (*string, bool)`

GetOutgoingCallerIdOk returns a tuple with the OutgoingCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallerId

`func (o *UserPostResponse) SetOutgoingCallerId(v string)`

SetOutgoingCallerId sets OutgoingCallerId field to given value.

### HasOutgoingCallerId

`func (o *UserPostResponse) HasOutgoingCallerId() bool`

HasOutgoingCallerId returns a boolean if a field has been set.

### GetPassword

`func (o *UserPostResponse) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserPostResponse) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserPostResponse) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserPostResponse) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *UserPostResponse) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *UserPostResponse) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *UserPostResponse) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *UserPostResponse) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRingSeconds

`func (o *UserPostResponse) GetRingSeconds() int32`

GetRingSeconds returns the RingSeconds field if non-nil, zero value otherwise.

### GetRingSecondsOk

`func (o *UserPostResponse) GetRingSecondsOk() (*int32, bool)`

GetRingSecondsOk returns a tuple with the RingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSeconds

`func (o *UserPostResponse) SetRingSeconds(v int32)`

SetRingSeconds sets RingSeconds field to given value.

### HasRingSeconds

`func (o *UserPostResponse) HasRingSeconds() bool`

HasRingSeconds returns a boolean if a field has been set.

### GetSimultaneousCalls

`func (o *UserPostResponse) GetSimultaneousCalls() int32`

GetSimultaneousCalls returns the SimultaneousCalls field if non-nil, zero value otherwise.

### GetSimultaneousCallsOk

`func (o *UserPostResponse) GetSimultaneousCallsOk() (*int32, bool)`

GetSimultaneousCallsOk returns a tuple with the SimultaneousCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousCalls

`func (o *UserPostResponse) SetSimultaneousCalls(v int32)`

SetSimultaneousCalls sets SimultaneousCalls field to given value.

### HasSimultaneousCalls

`func (o *UserPostResponse) HasSimultaneousCalls() bool`

HasSimultaneousCalls returns a boolean if a field has been set.

### GetSupervisionEnabled

`func (o *UserPostResponse) GetSupervisionEnabled() bool`

GetSupervisionEnabled returns the SupervisionEnabled field if non-nil, zero value otherwise.

### GetSupervisionEnabledOk

`func (o *UserPostResponse) GetSupervisionEnabledOk() (*bool, bool)`

GetSupervisionEnabledOk returns a tuple with the SupervisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionEnabled

`func (o *UserPostResponse) SetSupervisionEnabled(v bool)`

SetSupervisionEnabled sets SupervisionEnabled field to given value.

### HasSupervisionEnabled

`func (o *UserPostResponse) HasSupervisionEnabled() bool`

HasSupervisionEnabled returns a boolean if a field has been set.

### GetSwitchboards

`func (o *UserPostResponse) GetSwitchboards() []map[string]interface{}`

GetSwitchboards returns the Switchboards field if non-nil, zero value otherwise.

### GetSwitchboardsOk

`func (o *UserPostResponse) GetSwitchboardsOk() (*[]map[string]interface{}, bool)`

GetSwitchboardsOk returns a tuple with the Switchboards field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSwitchboards

`func (o *UserPostResponse) SetSwitchboards(v []map[string]interface{})`

SetSwitchboards sets Switchboards field to given value.

### HasSwitchboards

`func (o *UserPostResponse) HasSwitchboards() bool`

HasSwitchboards returns a boolean if a field has been set.

### GetTenantUuid

`func (o *UserPostResponse) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *UserPostResponse) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *UserPostResponse) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *UserPostResponse) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *UserPostResponse) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *UserPostResponse) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *UserPostResponse) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *UserPostResponse) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserfield

`func (o *UserPostResponse) GetUserfield() string`

GetUserfield returns the Userfield field if non-nil, zero value otherwise.

### GetUserfieldOk

`func (o *UserPostResponse) GetUserfieldOk() (*string, bool)`

GetUserfieldOk returns a tuple with the Userfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfield

`func (o *UserPostResponse) SetUserfield(v string)`

SetUserfield sets Userfield field to given value.

### HasUserfield

`func (o *UserPostResponse) HasUserfield() bool`

HasUserfield returns a boolean if a field has been set.

### GetUsername

`func (o *UserPostResponse) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserPostResponse) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserPostResponse) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserPostResponse) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *UserPostResponse) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserPostResponse) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserPostResponse) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserPostResponse) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetVoicemail

`func (o *UserPostResponse) GetVoicemail() Voicemail`

GetVoicemail returns the Voicemail field if non-nil, zero value otherwise.

### GetVoicemailOk

`func (o *UserPostResponse) GetVoicemailOk() (*Voicemail, bool)`

GetVoicemailOk returns a tuple with the Voicemail field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVoicemail

`func (o *UserPostResponse) SetVoicemail(v Voicemail)`

SetVoicemail sets Voicemail field to given value.

### HasVoicemail

`func (o *UserPostResponse) HasVoicemail() bool`

HasVoicemail returns a boolean if a field has been set.

### GetIncalls

`func (o *UserPostResponse) GetIncalls() []UserPostRelationIncallsIncallsInner`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *UserPostResponse) GetIncallsOk() (*[]UserPostRelationIncallsIncallsInner, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *UserPostResponse) SetIncalls(v []UserPostRelationIncallsIncallsInner)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *UserPostResponse) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
