# UserPostResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]UserEmail**](UserEmail.md) |  | [optional]
**Purpose** | Pointer to **string** |  | [optional]
**Username** | Pointer to **string** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

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

### GetEmails

`func (o *UserPostResponse) GetEmails() []UserEmail`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserPostResponse) GetEmailsOk() (*[]UserEmail, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserPostResponse) SetEmails(v []UserEmail)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UserPostResponse) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

### GetPurpose

`func (o *UserPostResponse) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UserPostResponse) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UserPostResponse) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *UserPostResponse) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

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

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
