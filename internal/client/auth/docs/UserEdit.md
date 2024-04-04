# UserEdit

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Enabled** | Pointer to **bool** |  | [optional]
**Firstname** | Pointer to **string** | The user&#39;s firstname | [optional]
**Lastname** | Pointer to **string** | The user&#39;s lastname | [optional]
**Purpose** | Pointer to **string** |  | [optional] [default to "user"]
**Username** | Pointer to **string** | The username that will identify that new username | [optional]

## Methods

### NewUserEdit

`func NewUserEdit() *UserEdit`

NewUserEdit instantiates a new UserEdit object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEditWithDefaults

`func NewUserEditWithDefaults() *UserEdit`

NewUserEditWithDefaults instantiates a new UserEdit object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnabled

`func (o *UserEdit) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *UserEdit) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *UserEdit) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *UserEdit) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetFirstname

`func (o *UserEdit) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserEdit) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserEdit) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserEdit) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserEdit) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserEdit) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserEdit) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserEdit) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetPurpose

`func (o *UserEdit) GetPurpose() string`

GetPurpose returns the Purpose field if non-nil, zero value otherwise.

### GetPurposeOk

`func (o *UserEdit) GetPurposeOk() (*string, bool)`

GetPurposeOk returns a tuple with the Purpose field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPurpose

`func (o *UserEdit) SetPurpose(v string)`

SetPurpose sets Purpose field to given value.

### HasPurpose

`func (o *UserEdit) HasPurpose() bool`

HasPurpose returns a boolean if a field has been set.

### GetUsername

`func (o *UserEdit) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserEdit) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserEdit) SetUsername(v string)`

SetUsername sets Username field to given value.

### HasUsername

`func (o *UserEdit) HasUsername() bool`

HasUsername returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
