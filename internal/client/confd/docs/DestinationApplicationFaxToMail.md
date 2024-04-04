# DestinationApplicationFaxToMail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | Pointer to **string** | MUST be &#39;fax_to_mail&#39; | [optional]
**Email** | Pointer to **string** | The email of the application | [optional]

## Methods

### NewDestinationApplicationFaxToMail

`func NewDestinationApplicationFaxToMail(type_ string, ) *DestinationApplicationFaxToMail`

NewDestinationApplicationFaxToMail instantiates a new DestinationApplicationFaxToMail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationFaxToMailWithDefaults

`func NewDestinationApplicationFaxToMailWithDefaults() *DestinationApplicationFaxToMail`

NewDestinationApplicationFaxToMailWithDefaults instantiates a new DestinationApplicationFaxToMail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationFaxToMail) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationFaxToMail) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationFaxToMail) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationFaxToMail) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationFaxToMail) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationFaxToMail) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DestinationApplicationFaxToMail) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEmail

`func (o *DestinationApplicationFaxToMail) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *DestinationApplicationFaxToMail) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *DestinationApplicationFaxToMail) SetEmail(v string)`

SetEmail sets Email field to given value.

### HasEmail

`func (o *DestinationApplicationFaxToMail) HasEmail() bool`

HasEmail returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
