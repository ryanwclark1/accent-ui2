# AdminUserEmailList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]AdminUserEmailListEmailsInner**](AdminUserEmailListEmailsInner.md) |  | [optional]

## Methods

### NewAdminUserEmailList

`func NewAdminUserEmailList() *AdminUserEmailList`

NewAdminUserEmailList instantiates a new AdminUserEmailList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAdminUserEmailListWithDefaults

`func NewAdminUserEmailListWithDefaults() *AdminUserEmailList`

NewAdminUserEmailListWithDefaults instantiates a new AdminUserEmailList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *AdminUserEmailList) GetEmails() []AdminUserEmailListEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *AdminUserEmailList) GetEmailsOk() (*[]AdminUserEmailListEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *AdminUserEmailList) SetEmails(v []AdminUserEmailListEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *AdminUserEmailList) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
