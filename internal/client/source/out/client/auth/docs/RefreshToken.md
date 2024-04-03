# RefreshToken

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ClientId** | Pointer to **string** | The &#x60;client_id&#x60; that was used to create this refresh token | [optional]
**CreatedAt** | Pointer to **string** | The time at which this token was created | [optional]
**Mobile** | Pointer to **bool** | Indicate if that refresh token was created with a mobile session type | [optional]
**TenantUuid** | Pointer to **string** | The tenant UUID of the user which created this refresh token | [optional]
**UserUuid** | Pointer to **string** | The UUID of the user which created this refresh token | [optional]

## Methods

### NewRefreshToken

`func NewRefreshToken() *RefreshToken`

NewRefreshToken instantiates a new RefreshToken object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRefreshTokenWithDefaults

`func NewRefreshTokenWithDefaults() *RefreshToken`

NewRefreshTokenWithDefaults instantiates a new RefreshToken object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetClientId

`func (o *RefreshToken) GetClientId() string`

GetClientId returns the ClientId field if non-nil, zero value otherwise.

### GetClientIdOk

`func (o *RefreshToken) GetClientIdOk() (*string, bool)`

GetClientIdOk returns a tuple with the ClientId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClientId

`func (o *RefreshToken) SetClientId(v string)`

SetClientId sets ClientId field to given value.

### HasClientId

`func (o *RefreshToken) HasClientId() bool`

HasClientId returns a boolean if a field has been set.

### GetCreatedAt

`func (o *RefreshToken) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *RefreshToken) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *RefreshToken) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *RefreshToken) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetMobile

`func (o *RefreshToken) GetMobile() bool`

GetMobile returns the Mobile field if non-nil, zero value otherwise.

### GetMobileOk

`func (o *RefreshToken) GetMobileOk() (*bool, bool)`

GetMobileOk returns a tuple with the Mobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMobile

`func (o *RefreshToken) SetMobile(v bool)`

SetMobile sets Mobile field to given value.

### HasMobile

`func (o *RefreshToken) HasMobile() bool`

HasMobile returns a boolean if a field has been set.

### GetTenantUuid

`func (o *RefreshToken) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *RefreshToken) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *RefreshToken) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *RefreshToken) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUserUuid

`func (o *RefreshToken) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *RefreshToken) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *RefreshToken) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *RefreshToken) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
