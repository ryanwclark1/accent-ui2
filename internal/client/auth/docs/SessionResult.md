# SessionResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Mobile** | Pointer to **bool** |  | [optional]
**TenantUuid** | Pointer to **string** |  | [optional]
**UserUuid** | Pointer to **string** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewSessionResult

`func NewSessionResult() *SessionResult`

NewSessionResult instantiates a new SessionResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSessionResultWithDefaults

`func NewSessionResultWithDefaults() *SessionResult`

NewSessionResultWithDefaults instantiates a new SessionResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMobile

`func (o *SessionResult) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *SessionResult) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *SessionResult) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *SessionResult) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTenantUuid

`func (o *SessionResult) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *SessionResult) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *SessionResult) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *SessionResult) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUserUuid

`func (o *SessionResult) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *SessionResult) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *SessionResult) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *SessionResult) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetUuid

`func (o *SessionResult) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *SessionResult) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *SessionResult) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *SessionResult) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
