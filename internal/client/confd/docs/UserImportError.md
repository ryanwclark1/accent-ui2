# UserImportError

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Errors** | Pointer to [**[]UserImportErrorErrorsInner**](UserImportErrorErrorsInner.md) |  | [optional]

## Methods

### NewUserImportError

`func NewUserImportError() *UserImportError`

NewUserImportError instantiates a new UserImportError object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportErrorWithDefaults

`func NewUserImportErrorWithDefaults() *UserImportError`

NewUserImportErrorWithDefaults instantiates a new UserImportError object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetErrors

`func (o *UserImportError) GetErrors() []UserImportErrorErrorsInner`

GetErrors returns the Errors field if non-nil, zero value otherwise.

### GetErrorsOk

`func (o *UserImportError) GetErrorsOk() (*[]UserImportErrorErrorsInner, bool)`

GetErrorsOk returns a tuple with the Errors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrors

`func (o *UserImportError) SetErrors(v []UserImportErrorErrorsInner)`

SetErrors sets Errors field to given value.

### HasErrors

`func (o *UserImportError) HasErrors() bool`

HasErrors returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
