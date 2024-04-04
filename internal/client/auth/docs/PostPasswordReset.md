# PostPasswordReset

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Password** | **string** | The desired password |

## Methods

### NewPostPasswordReset

`func NewPostPasswordReset(password string, ) *PostPasswordReset`

NewPostPasswordReset instantiates a new PostPasswordReset object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPostPasswordResetWithDefaults

`func NewPostPasswordResetWithDefaults() *PostPasswordReset`

NewPostPasswordResetWithDefaults instantiates a new PostPasswordReset object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPassword

`func (o *PostPasswordReset) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *PostPasswordReset) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *PostPasswordReset) SetPassword(v string)`

SetPassword sets Password field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
