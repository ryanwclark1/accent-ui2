# UserRegister

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**EmailAddress** | **string** | The main email address of the new username |
**Firstname** | Pointer to **string** | The user&#39;s firstname | [optional]
**Lastname** | Pointer to **string** | The user&#39;s lastname | [optional]
**Password** | **string** | The password of the newly created username |
**Username** | **string** | The username that will identify that new username |

## Methods

### NewUserRegister

`func NewUserRegister(emailAddress string, password string, username string, ) *UserRegister`

NewUserRegister instantiates a new UserRegister object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserRegisterWithDefaults

`func NewUserRegisterWithDefaults() *UserRegister`

NewUserRegisterWithDefaults instantiates a new UserRegister object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmailAddress

`func (o *UserRegister) GetEmailAddress() string`

GetEmailAddress returns the EmailAddress field if non-nil, zero value otherwise.

### GetEmailAddressOk

`func (o *UserRegister) GetEmailAddressOk() (*string, bool)`

GetEmailAddressOk returns a tuple with the EmailAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmailAddress

`func (o *UserRegister) SetEmailAddress(v string)`

SetEmailAddress sets EmailAddress field to given value.

### GetFirstname

`func (o *UserRegister) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *UserRegister) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *UserRegister) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *UserRegister) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *UserRegister) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *UserRegister) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *UserRegister) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *UserRegister) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetPassword

`func (o *UserRegister) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *UserRegister) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *UserRegister) SetPassword(v string)`

SetPassword sets Password field to given value.

### GetUsername

`func (o *UserRegister) GetUsername() string`

GetUsername returns the Username field if non-nil, zero value otherwise.

### GetUsernameOk

`func (o *UserRegister) GetUsernameOk() (*string, bool)`

GetUsernameOk returns a tuple with the Username field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsername

`func (o *UserRegister) SetUsername(v string)`

SetUsername sets Username field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
