# PasswordChange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NewPassword** | **string** | The desired password |
**OldPassword** | **string** | The old password |

## Methods

### NewPasswordChange

`func NewPasswordChange(newPassword string, oldPassword string, ) *PasswordChange`

NewPasswordChange instantiates a new PasswordChange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordChangeWithDefaults

`func NewPasswordChangeWithDefaults() *PasswordChange`

NewPasswordChangeWithDefaults instantiates a new PasswordChange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNewPassword

`func (o *PasswordChange) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordChange) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordChange) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### GetOldPassword

`func (o *PasswordChange) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PasswordChange) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PasswordChange) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
