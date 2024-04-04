# UserResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]UserEmail**](UserEmail.md) |  | [optional]
**Enabled** | Pointer to **bool** |  | [optional]
**Firstname** | Pointer to **string** |  | [optional]
**Lastname** | Pointer to **string** |  | [optional]
**Purpose** | Pointer to **string** |  | [optional]
**TenantUuid** | Pointer to **string** |  | [optional]
**Username** | Pointer to **string** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewUserResult

`func NewUserResult() *UserResult`

NewUserResult instantiates a new UserResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserResultWithDefaults

`func NewUserResultWithDefaults() *UserResult`

NewUserResultWithDefaults instantiates a new UserResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UserResult) GetEmails() []UserEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserResult) GetEmailsOk() (*[]UserEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserResult) SetEmails(v []UserEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UserResult) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetEnabled

`func (o *UserResult) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserResult) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserResult) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserResult) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstname

`func (o *UserResult) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserResult) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserResult) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserResult) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserResult) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserResult) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserResult) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserResult) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetPurpose

`func (o *UserResult) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UserResult) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UserResult) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *UserResult) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetTenantUuid

`func (o *UserResult) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *UserResult) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *UserResult) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *UserResult) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUsername

`func (o *UserResult) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserResult) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserResult) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserResult) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *UserResult) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserResult) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserResult) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserResult) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
