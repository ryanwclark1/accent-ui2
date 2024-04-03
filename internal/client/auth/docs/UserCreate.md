# UserCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | Pointer to **string** | The main email address of the new username | [optional]
**Enabled** | Pointer to **bool** |  | [optional]
**Firstname** | Pointer to **string** | The user&#39;s firstname | [optional]
**Lastname** | Pointer to **string** | The user&#39;s lastname | [optional]
**Password** | Pointer to **string** | The password of the newly created username | [optional]
**Purpose** | Pointer to **string** |  | [optional] [default to "user"]
**Username** | Pointer to **string** | The username that will identify that new username | [optional]
**Uuid** | Pointer to **string** | The user&#39;s UUID | [optional]

## Methods

### NewUserCreate

`func NewUserCreate() *UserCreate`

NewUserCreate instantiates a new UserCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCreateWithDefaults

`func NewUserCreateWithDefaults() *UserCreate`

NewUserCreateWithDefaults instantiates a new UserCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UserCreate) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserCreate) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserCreate) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### HasEmailAddress

`func (o *UserCreate) HasEmailAddress() bool`

HasEmailAddress returns a boolean if a field has been set.

### GetEnabled

`func (o *UserCreate) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserCreate) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserCreate) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserCreate) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstname

`func (o *UserCreate) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserCreate) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserCreate) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserCreate) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserCreate) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserCreate) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserCreate) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserCreate) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetPassword

`func (o *UserCreate) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserCreate) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserCreate) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *UserCreate) HasPassword() bool`

HasPassword returns a boolean if a field has been set.

### GetPurpose

`func (o *UserCreate) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UserCreate) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UserCreate) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *UserCreate) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetUsername

`func (o *UserCreate) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserCreate) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserCreate) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserCreate) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

### GetUuid

`func (o *UserCreate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *UserCreate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *UserCreate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *UserCreate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
