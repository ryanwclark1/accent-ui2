# Conference

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | Pointer to **string** | name to identify the conference | [optional]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]
**Incalls** | Pointer to [**[]ConferenceRelationIncall**](ConferenceRelationIncall.md) |  | [optional] [readonly]
**AdminPin** | Pointer to **string** | Administrator pin to enter in the conference | [optional]
**AnnounceJoinLeave** | Pointer to **bool** | Record name and announce join/leave | [optional] [default to false]
**AnnounceOnlyUser** | Pointer to **bool** | Announce when a participant is alone in conference | [optional] [default to true]
**AnnounceUserCount** | Pointer to **bool** | Announce the number of participants | [optional] [default to false]
**MaxUsers** | Pointer to **int32** | Maximum users allowed in the conference. This exclude admin. | [optional]
**MusicOnHold** | Pointer to **string** | Name of the MOH category to use for music on hold | [optional]
**Pin** | Pointer to **string** | Pin to enter in the conference | [optional]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before entering the conference | [optional]
**QuietJoinLeave** | Pointer to **bool** | Play &#39;beep&#39; notification when join/leave a conference | [optional] [default to false]
**Record** | Pointer to **bool** | Record the conference | [optional] [default to false]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewConference

`func NewConference() *Conference`

NewConference instantiates a new Conference object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceWithDefaults

`func NewConferenceWithDefaults() *Conference`

NewConferenceWithDefaults instantiates a new Conference object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Conference) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Conference) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Conference) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Conference) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Conference) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Conference) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Conference) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Conference) HasName() bool`

HasName returns a boolean if a field has been set.

### GetExtensions

`func (o *Conference) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Conference) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Conference) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Conference) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetIncalls

`func (o *Conference) GetIncalls() []ConferenceRelationIncall`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *Conference) GetIncallsOk() (*[]ConferenceRelationIncall, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *Conference) SetIncalls(v []ConferenceRelationIncall)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *Conference) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetAdminPin

`func (o *Conference) GetAdminPin() string`

GetAdminPin returns the AdminPin field if non-nil, zero value otherwise.

### GetAdminPinOk

`func (o *Conference) GetAdminPinOk() (*string, bool)`

GetAdminPinOk returns a tuple with the AdminPin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdminPin

`func (o *Conference) SetAdminPin(v string)`

SetAdminPin sets AdminPin field to given value.

### HasAdminPin

`func (o *Conference) HasAdminPin() bool`

HasAdminPin returns a boolean if a field has been set.

### GetAnnounceJoinLeave

`func (o *Conference) GetAnnounceJoinLeave() bool`

GetAnnounceJoinLeave returns the AnnounceJoinLeave field if non-nil, zero value otherwise.

### GetAnnounceJoinLeaveOk

`func (o *Conference) GetAnnounceJoinLeaveOk() (*bool, bool)`

GetAnnounceJoinLeaveOk returns a tuple with the AnnounceJoinLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceJoinLeave

`func (o *Conference) SetAnnounceJoinLeave(v bool)`

SetAnnounceJoinLeave sets AnnounceJoinLeave field to given value.

### HasAnnounceJoinLeave

`func (o *Conference) HasAnnounceJoinLeave() bool`

HasAnnounceJoinLeave returns a boolean if a field has been set.

### GetAnnounceOnlyUser

`func (o *Conference) GetAnnounceOnlyUser() bool`

GetAnnounceOnlyUser returns the AnnounceOnlyUser field if non-nil, zero value otherwise.

### GetAnnounceOnlyUserOk

`func (o *Conference) GetAnnounceOnlyUserOk() (*bool, bool)`

GetAnnounceOnlyUserOk returns a tuple with the AnnounceOnlyUser field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceOnlyUser

`func (o *Conference) SetAnnounceOnlyUser(v bool)`

SetAnnounceOnlyUser sets AnnounceOnlyUser field to given value.

### HasAnnounceOnlyUser

`func (o *Conference) HasAnnounceOnlyUser() bool`

HasAnnounceOnlyUser returns a boolean if a field has been set.

### GetAnnounceUserCount

`func (o *Conference) GetAnnounceUserCount() bool`

GetAnnounceUserCount returns the AnnounceUserCount field if non-nil, zero value otherwise.

### GetAnnounceUserCountOk

`func (o *Conference) GetAnnounceUserCountOk() (*bool, bool)`

GetAnnounceUserCountOk returns a tuple with the AnnounceUserCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnnounceUserCount

`func (o *Conference) SetAnnounceUserCount(v bool)`

SetAnnounceUserCount sets AnnounceUserCount field to given value.

### HasAnnounceUserCount

`func (o *Conference) HasAnnounceUserCount() bool`

HasAnnounceUserCount returns a boolean if a field has been set.

### GetMaxUsers

`func (o *Conference) GetMaxUsers() int32`

GetMaxUsers returns the MaxUsers field if non-nil, zero value otherwise.

### GetMaxUsersOk

`func (o *Conference) GetMaxUsersOk() (*int32, bool)`

GetMaxUsersOk returns a tuple with the MaxUsers field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMaxUsers

`func (o *Conference) SetMaxUsers(v int32)`

SetMaxUsers sets MaxUsers field to given value.

### HasMaxUsers

`func (o *Conference) HasMaxUsers() bool`

HasMaxUsers returns a boolean if a field has been set.

### GetMusicOnHold

`func (o *Conference) GetMusicOnHold() string`

GetMusicOnHold returns the MusicOnHold field if non-nil, zero value otherwise.

### GetMusicOnHoldOk

`func (o *Conference) GetMusicOnHoldOk() (*string, bool)`

GetMusicOnHoldOk returns a tuple with the MusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMusicOnHold

`func (o *Conference) SetMusicOnHold(v string)`

SetMusicOnHold sets MusicOnHold field to given value.

### HasMusicOnHold

`func (o *Conference) HasMusicOnHold() bool`

HasMusicOnHold returns a boolean if a field has been set.

### GetPin

`func (o *Conference) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *Conference) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *Conference) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *Conference) HasPin() bool`

HasPin returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Conference) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Conference) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Conference) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Conference) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetQuietJoinLeave

`func (o *Conference) GetQuietJoinLeave() bool`

GetQuietJoinLeave returns the QuietJoinLeave field if non-nil, zero value otherwise.

### GetQuietJoinLeaveOk

`func (o *Conference) GetQuietJoinLeaveOk() (*bool, bool)`

GetQuietJoinLeaveOk returns a tuple with the QuietJoinLeave field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQuietJoinLeave

`func (o *Conference) SetQuietJoinLeave(v bool)`

SetQuietJoinLeave sets QuietJoinLeave field to given value.

### HasQuietJoinLeave

`func (o *Conference) HasQuietJoinLeave() bool`

HasQuietJoinLeave returns a boolean if a field has been set.

### GetRecord

`func (o *Conference) GetRecord() bool`

GetRecord returns the Record field if non-nil, zero value otherwise.

### GetRecordOk

`func (o *Conference) GetRecordOk() (*bool, bool)`

GetRecordOk returns a tuple with the Record field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecord

`func (o *Conference) SetRecord(v bool)`

SetRecord sets Record field to given value.

### HasRecord

`func (o *Conference) HasRecord() bool`

HasRecord returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Conference) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Conference) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Conference) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Conference) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
