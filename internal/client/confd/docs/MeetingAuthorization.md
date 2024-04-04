# MeetingAuthorization

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CreationTime** | Pointer to **time.Time** |  | [optional]
**GuestName** | Pointer to **string** |  | [optional]
**SipAuthorization** | Pointer to **string** |  | [optional]
**Status** | Pointer to **string** | The status of the authorization. If the meeting does not require an authorization, the authorization will always be &#x60;accepted&#x60;. | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewMeetingAuthorization

`func NewMeetingAuthorization() *MeetingAuthorization`

NewMeetingAuthorization instantiates a new MeetingAuthorization object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMeetingAuthorizationWithDefaults

`func NewMeetingAuthorizationWithDefaults() *MeetingAuthorization`

NewMeetingAuthorizationWithDefaults instantiates a new MeetingAuthorization object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCreationTime

`func (o *MeetingAuthorization) GetCreationTime() time.Time`

GetCreationTime returns the CreationTime field if non-nil, zero value otherwise.

### GetCreationTimeOk

`func (o *MeetingAuthorization) GetCreationTimeOk() (*time.Time, bool)`

GetCreationTimeOk returns a tuple with the CreationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreationTime

`func (o *MeetingAuthorization) SetCreationTime(v time.Time)`

SetCreationTime sets CreationTime field to given value.

### HasCreationTime

`func (o *MeetingAuthorization) HasCreationTime() bool`

HasCreationTime returns a boolean if a field has been set.

### GetGuestName

`func (o *MeetingAuthorization) GetGuestName() string`

GetGuestName returns the GuestName field if non-nil, zero value otherwise.

### GetGuestNameOk

`func (o *MeetingAuthorization) GetGuestNameOk() (*string, bool)`

GetGuestNameOk returns a tuple with the GuestName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGuestName

`func (o *MeetingAuthorization) SetGuestName(v string)`

SetGuestName sets GuestName field to given value.

### HasGuestName

`func (o *MeetingAuthorization) HasGuestName() bool`

HasGuestName returns a boolean if a field has been set.

### GetSipAuthorization

`func (o *MeetingAuthorization) GetSipAuthorization() string`

GetSipAuthorization returns the SipAuthorization field if non-nil, zero value otherwise.

### GetSipAuthorizationOk

`func (o *MeetingAuthorization) GetSipAuthorizationOk() (*string, bool)`

GetSipAuthorizationOk returns a tuple with the SipAuthorization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSipAuthorization

`func (o *MeetingAuthorization) SetSipAuthorization(v string)`

SetSipAuthorization sets SipAuthorization field to given value.

### HasSipAuthorization

`func (o *MeetingAuthorization) HasSipAuthorization() bool`

HasSipAuthorization returns a boolean if a field has been set.

### GetStatus

`func (o *MeetingAuthorization) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MeetingAuthorization) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MeetingAuthorization) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MeetingAuthorization) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetUuid

`func (o *MeetingAuthorization) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *MeetingAuthorization) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *MeetingAuthorization) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *MeetingAuthorization) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
