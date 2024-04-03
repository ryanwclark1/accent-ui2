# BaseUser

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

## Methods

### NewBaseUser

`func NewBaseUser() *BaseUser`

NewBaseUser instantiates a new BaseUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewBaseUserWithDefaults

`func NewBaseUserWithDefaults() *BaseUser`

NewBaseUserWithDefaults instantiates a new BaseUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *BaseUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *BaseUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *BaseUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *BaseUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *BaseUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *BaseUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *BaseUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *BaseUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *BaseUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *BaseUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *BaseUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *BaseUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetCallPermissionPassword

`func (o *BaseUser) GetCallPermissionPassword() string`

GetCallPermissionPassword returns the CallPermissionPassword field if non-nil, zero value otherwise.

### GetCallPermissionPasswordOk

`func (o *BaseUser) GetCallPermissionPasswordOk() (*string, bool)`

GetCallPermissionPasswordOk returns a tuple with the CallPermissionPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissionPassword

`func (o *BaseUser) SetCallPermissionPassword(v string)`

SetCallPermissionPassword sets CallPermissionPassword field to given value.

### HasCallPermissionPassword

`func (o *BaseUser) HasCallPermissionPassword() bool`

HasCallPermissionPassword returns a boolean if a field has been set.

### GetCallRecordEnabled

`func (o *BaseUser) GetCallRecordEnabled() bool`

GetCallRecordEnabled returns the CallRecordEnabled field if non-nil, zero value otherwise.

### GetCallRecordEnabledOk

`func (o *BaseUser) GetCallRecordEnabledOk() (*bool, bool)`

GetCallRecordEnabledOk returns a tuple with the CallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordEnabled

`func (o *BaseUser) SetCallRecordEnabled(v bool)`

SetCallRecordEnabled sets CallRecordEnabled field to given value.

### HasCallRecordEnabled

`func (o *BaseUser) HasCallRecordEnabled() bool`

HasCallRecordEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingExternalEnabled

`func (o *BaseUser) GetCallRecordIncomingExternalEnabled() bool`

GetCallRecordIncomingExternalEnabled returns the CallRecordIncomingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingExternalEnabledOk

`func (o *BaseUser) GetCallRecordIncomingExternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingExternalEnabledOk returns a tuple with the CallRecordIncomingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingExternalEnabled

`func (o *BaseUser) SetCallRecordIncomingExternalEnabled(v bool)`

SetCallRecordIncomingExternalEnabled sets CallRecordIncomingExternalEnabled field to given value.

### HasCallRecordIncomingExternalEnabled

`func (o *BaseUser) HasCallRecordIncomingExternalEnabled() bool`

HasCallRecordIncomingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordIncomingInternalEnabled

`func (o *BaseUser) GetCallRecordIncomingInternalEnabled() bool`

GetCallRecordIncomingInternalEnabled returns the CallRecordIncomingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordIncomingInternalEnabledOk

`func (o *BaseUser) GetCallRecordIncomingInternalEnabledOk() (*bool, bool)`

GetCallRecordIncomingInternalEnabledOk returns a tuple with the CallRecordIncomingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordIncomingInternalEnabled

`func (o *BaseUser) SetCallRecordIncomingInternalEnabled(v bool)`

SetCallRecordIncomingInternalEnabled sets CallRecordIncomingInternalEnabled field to given value.

### HasCallRecordIncomingInternalEnabled

`func (o *BaseUser) HasCallRecordIncomingInternalEnabled() bool`

HasCallRecordIncomingInternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingExternalEnabled

`func (o *BaseUser) GetCallRecordOutgoingExternalEnabled() bool`

GetCallRecordOutgoingExternalEnabled returns the CallRecordOutgoingExternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingExternalEnabledOk

`func (o *BaseUser) GetCallRecordOutgoingExternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingExternalEnabledOk returns a tuple with the CallRecordOutgoingExternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingExternalEnabled

`func (o *BaseUser) SetCallRecordOutgoingExternalEnabled(v bool)`

SetCallRecordOutgoingExternalEnabled sets CallRecordOutgoingExternalEnabled field to given value.

### HasCallRecordOutgoingExternalEnabled

`func (o *BaseUser) HasCallRecordOutgoingExternalEnabled() bool`

HasCallRecordOutgoingExternalEnabled returns a boolean if a field has been set.

### GetCallRecordOutgoingInternalEnabled

`func (o *BaseUser) GetCallRecordOutgoingInternalEnabled() bool`

GetCallRecordOutgoingInternalEnabled returns the CallRecordOutgoingInternalEnabled field if non-nil, zero value otherwise.

### GetCallRecordOutgoingInternalEnabledOk

`func (o *BaseUser) GetCallRecordOutgoingInternalEnabledOk() (*bool, bool)`

GetCallRecordOutgoingInternalEnabledOk returns a tuple with the CallRecordOutgoingInternalEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallRecordOutgoingInternalEnabled

`func (o *BaseUser) SetCallRecordOutgoingInternalEnabled(v bool)`

SetCallRecordOutgoingInternalEnabled sets CallRecordOutgoingInternalEnabled field to given value.

### HasCallRecordOutgoingInternalEnabled

`func (o *BaseUser) HasCallRecordOutgoingInternalEnabled() bool`

HasCallRecordOutgoingInternalEnabled returns a boolean if a field has been set.

### GetCallTransferEnabled

`func (o *BaseUser) GetCallTransferEnabled() bool`

GetCallTransferEnabled returns the CallTransferEnabled field if non-nil, zero value otherwise.

### GetCallTransferEnabledOk

`func (o *BaseUser) GetCallTransferEnabledOk() (*bool, bool)`

GetCallTransferEnabledOk returns a tuple with the CallTransferEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallTransferEnabled

`func (o *BaseUser) SetCallTransferEnabled(v bool)`

SetCallTransferEnabled sets CallTransferEnabled field to given value.

### HasCallTransferEnabled

`func (o *BaseUser) HasCallTransferEnabled() bool`

HasCallTransferEnabled returns a boolean if a field has been set.

### GetCallerId

`func (o *BaseUser) GetCallerId() string`

GetCallerId returns the CallerId field if non-nil, zero value otherwise.

### GetCallerIdOk

`func (o *BaseUser) GetCallerIdOk() (*string, bool)`

GetCallerIdOk returns a tuple with the CallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerId

`func (o *BaseUser) SetCallerId(v string)`

SetCallerId sets CallerId field to given value.

### HasCallerId

`func (o *BaseUser) HasCallerId() bool`

HasCallerId returns a boolean if a field has been set.

### GetDescription

`func (o *BaseUser) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *BaseUser) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *BaseUser) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *BaseUser) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetDtmfHangupEnabled

`func (o *BaseUser) GetDtmfHangupEnabled() bool`

GetDtmfHangupEnabled returns the DtmfHangupEnabled field if non-nil, zero value otherwise.

### GetDtmfHangupEnabledOk

`func (o *BaseUser) GetDtmfHangupEnabledOk() (*bool, bool)`

GetDtmfHangupEnabledOk returns a tuple with the DtmfHangupEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDtmfHangupEnabled

`func (o *BaseUser) SetDtmfHangupEnabled(v bool)`

SetDtmfHangupEnabled sets DtmfHangupEnabled field to given value.

### HasDtmfHangupEnabled

`func (o *BaseUser) HasDtmfHangupEnabled() bool`

HasDtmfHangupEnabled returns a boolean if a field has been set.

### GetEmail

`func (o *BaseUser) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *BaseUser) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *BaseUser) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *BaseUser) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

### GetEnabled

`func (o *BaseUser) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *BaseUser) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *BaseUser) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *BaseUser) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *BaseUser) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *BaseUser) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *BaseUser) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *BaseUser) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLanguage

`func (o *BaseUser) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *BaseUser) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *BaseUser) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *BaseUser) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMobilePhoneNumber

`func (o *BaseUser) GetMobilePhoneNumber() string`

GetMobilePhoneNumber returns the MobilePhoneNumber field if non-nil, zero value otherwise.

### GetMobilePhoneNumberOk

`func (o *BaseUser) GetMobilePhoneNumberOk() (*string, bool)`

GetMobilePhoneNumberOk returns a tuple with the MobilePhoneNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobilePhoneNumber

`func (o *BaseUser) SetMobilePhoneNumber(v string)`

SetMobilePhoneNumber sets MobilePhoneNumber field to given value.

### HasMobilePhoneNumber

`func (o *BaseUser) HasMobilePhoneNumber() bool`

HasMobilePhoneNumber returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *BaseUser) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *BaseUser) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *BaseUser) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *BaseUser) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetOnlineCallRecordEnabled

`func (o *BaseUser) GetOnlineCallRecordEnabled() bool`

GetOnlineCallRecordEnabled returns the OnlineCallRecordEnabled field if non-nil, zero value otherwise.

### GetOnlineCallRecordEnabledOk

`func (o *BaseUser) GetOnlineCallRecordEnabledOk() (*bool, bool)`

GetOnlineCallRecordEnabledOk returns a tuple with the OnlineCallRecordEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOnlineCallRecordEnabled

`func (o *BaseUser) SetOnlineCallRecordEnabled(v bool)`

SetOnlineCallRecordEnabled sets OnlineCallRecordEnabled field to given value.

### HasOnlineCallRecordEnabled

`func (o *BaseUser) HasOnlineCallRecordEnabled() bool`

HasOnlineCallRecordEnabled returns a boolean if a field has been set.

### GetOutgoingCallerId

`func (o *BaseUser) GetOutgoingCallerId() string`

GetOutgoingCallerId returns the OutgoingCallerId field if non-nil, zero value otherwise.

### GetOutgoingCallerIdOk

`func (o *BaseUser) GetOutgoingCallerIdOk() (*string, bool)`

GetOutgoingCallerIdOk returns a tuple with the OutgoingCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutgoingCallerId

`func (o *BaseUser) SetOutgoingCallerId(v string)`

SetOutgoingCallerId sets OutgoingCallerId field to given value.

### HasOutgoingCallerId

`func (o *BaseUser) HasOutgoingCallerId() bool`

HasOutgoingCallerId returns a boolean if a field has been set.

### GetPassword

`func (o *BaseUser) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *BaseUser) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *BaseUser) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *BaseUser) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *BaseUser) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *BaseUser) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *BaseUser) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *BaseUser) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRingSeconds

`func (o *BaseUser) GetRingSeconds() int32`

GetRingSeconds returns the RingSeconds field if non-nil, zero value otherwise.

### GetRingSecondsOk

`func (o *BaseUser) GetRingSecondsOk() (*int32, bool)`

GetRingSecondsOk returns a tuple with the RingSeconds field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingSeconds

`func (o *BaseUser) SetRingSeconds(v int32)`

SetRingSeconds sets RingSeconds field to given value.

### HasRingSeconds

`func (o *BaseUser) HasRingSeconds() bool`

HasRingSeconds returns a boolean if a field has been set.

### GetSimultaneousCalls

`func (o *BaseUser) GetSimultaneousCalls() int32`

GetSimultaneousCalls returns the SimultaneousCalls field if non-nil, zero value otherwise.

### GetSimultaneousCallsOk

`func (o *BaseUser) GetSimultaneousCallsOk() (*int32, bool)`

GetSimultaneousCallsOk returns a tuple with the SimultaneousCalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSimultaneousCalls

`func (o *BaseUser) SetSimultaneousCalls(v int32)`

SetSimultaneousCalls sets SimultaneousCalls field to given value.

### HasSimultaneousCalls

`func (o *BaseUser) HasSimultaneousCalls() bool`

HasSimultaneousCalls returns a boolean if a field has been set.

### GetSupervisionEnabled

`func (o *BaseUser) GetSupervisionEnabled() bool`

GetSupervisionEnabled returns the SupervisionEnabled field if non-nil, zero value otherwise.

### GetSupervisionEnabledOk

`func (o *BaseUser) GetSupervisionEnabledOk() (*bool, bool)`

GetSupervisionEnabledOk returns a tuple with the SupervisionEnabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupervisionEnabled

`func (o *BaseUser) SetSupervisionEnabled(v bool)`

SetSupervisionEnabled sets SupervisionEnabled field to given value.

### HasSupervisionEnabled

`func (o *BaseUser) HasSupervisionEnabled() bool`

HasSupervisionEnabled returns a boolean if a field has been set.

### GetTenantUuid

`func (o *BaseUser) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *BaseUser) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *BaseUser) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *BaseUser) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimezone

`func (o *BaseUser) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *BaseUser) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *BaseUser) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *BaseUser) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

### GetUserfield

`func (o *BaseUser) GetUserfield() string`

GetUserfield returns the Userfield field if non-nil, zero value otherwise.

### GetUserfieldOk

`func (o *BaseUser) GetUserfieldOk() (*string, bool)`

GetUserfieldOk returns a tuple with the Userfield field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserfield

`func (o *BaseUser) SetUserfield(v string)`

SetUserfield sets Userfield field to given value.

### HasUserfield

`func (o *BaseUser) HasUserfield() bool`

HasUserfield returns a boolean if a field has been set.

### GetUsername

`func (o *BaseUser) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *BaseUser) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *BaseUser) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *BaseUser) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
