# UserEmailList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | Pointer to [**[]UserEmailListEmailsInner**](UserEmailListEmailsInner.md) |  | [optional]

## Methods

### NewUserEmailList

`func NewUserEmailList() *UserEmailList`

NewUserEmailList instantiates a new UserEmailList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEmailListWithDefaults

`func NewUserEmailListWithDefaults() *UserEmailList`

NewUserEmailListWithDefaults instantiates a new UserEmailList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *UserEmailList) GetEmails() []UserEmailListEmailsInner`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *UserEmailList) GetEmailsOk() (*[]UserEmailListEmailsInner, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *UserEmailList) SetEmails(v []UserEmailListEmailsInner)`

SetEmails sets Emails field to given value.

### HasEmails

`func (o *UserEmailList) HasEmails() bool`

HasEmails returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
