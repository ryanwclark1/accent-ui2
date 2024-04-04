# Participant

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Admin** | Pointer to **bool** | Is the participant an admin of the conference? | [optional]
**CallId** | Pointer to **string** | The ID of the participant&#39;s call | [optional]
**CallerIdName** | Pointer to **string** | The participant&#39;s name | [optional]
**CallerIdNum** | Pointer to **string** | The participant&#39;s number | [optional]
**Id** | Pointer to **string** | The participant&#39;s ID | [optional]
**JoinTime** | Pointer to **int32** | Elapsed seconds since the participant joined the conference | [optional]
**Language** | Pointer to **string** | The participant&#39;s language | [optional]
**Muted** | Pointer to **bool** | Is the participant muted? | [optional]
**UserUuid** | Pointer to **string** | The UUID of the participant&#39;s user. &#x60;null&#x60; when there is no user. | [optional]

## Methods

### NewParticipant

`func NewParticipant() *Participant`

NewParticipant instantiates a new Participant object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParticipantWithDefaults

`func NewParticipantWithDefaults() *Participant`

NewParticipantWithDefaults instantiates a new Participant object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAdmin

`func (o *Participant) GetAdmin() bool`

GetAdmin returns the Admin field if non-nil, zero value otherwise.

### GetAdminOk

`func (o *Participant) GetAdminOk() (*bool, bool)`

GetAdminOk returns a tuple with the Admin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAdmin

`func (o *Participant) SetAdmin(v bool)`

SetAdmin sets Admin field to given value.

### HasAdmin

`func (o *Participant) HasAdmin() bool`

HasAdmin returns a boolean if a field has been set.

### GetCallId

`func (o *Participant) GetCallId() string`

GetCallId returns the CallId field if non-nil, zero value otherwise.

### GetCallIdOk

`func (o *Participant) GetCallIdOk() (*string, bool)`

GetCallIdOk returns a tuple with the CallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallId

`func (o *Participant) SetCallId(v string)`

SetCallId sets CallId field to given value.

### HasCallId

`func (o *Participant) HasCallId() bool`

HasCallId returns a boolean if a field has been set.

### GetCallerIdName

`func (o *Participant) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Participant) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Participant) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Participant) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *Participant) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *Participant) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *Participant) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *Participant) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetId

`func (o *Participant) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Participant) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Participant) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Participant) HasId() bool`

HasId returns a boolean if a field has been set.

### GetJoinTime

`func (o *Participant) GetJoinTime() int32`

GetJoinTime returns the JoinTime field if non-nil, zero value otherwise.

### GetJoinTimeOk

`func (o *Participant) GetJoinTimeOk() (*int32, bool)`

GetJoinTimeOk returns a tuple with the JoinTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetJoinTime

`func (o *Participant) SetJoinTime(v int32)`

SetJoinTime sets JoinTime field to given value.

### HasJoinTime

`func (o *Participant) HasJoinTime() bool`

HasJoinTime returns a boolean if a field has been set.

### GetLanguage

`func (o *Participant) GetLanguage() string`

GetLanguage returns the Language field if non-nil, zero value otherwise.

### GetLanguageOk

`func (o *Participant) GetLanguageOk() (*string, bool)`

GetLanguageOk returns a tuple with the Language field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLanguage

`func (o *Participant) SetLanguage(v string)`

SetLanguage sets Language field to given value.

### HasLanguage

`func (o *Participant) HasLanguage() bool`

HasLanguage returns a boolean if a field has been set.

### GetMuted

`func (o *Participant) GetMuted() bool`

GetMuted returns the Muted field if non-nil, zero value otherwise.

### GetMutedOk

`func (o *Participant) GetMutedOk() (*bool, bool)`

GetMutedOk returns a tuple with the Muted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMuted

`func (o *Participant) SetMuted(v bool)`

SetMuted sets Muted field to given value.

### HasMuted

`func (o *Participant) HasMuted() bool`

HasMuted returns a boolean if a field has been set.

### GetUserUuid

`func (o *Participant) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *Participant) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *Participant) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *Participant) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
