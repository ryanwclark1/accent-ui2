# CallFilterRecipientUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Timeout** | Pointer to **int32** |  | [optional]

## Methods

### NewCallFilterRecipientUser

`func NewCallFilterRecipientUser() *CallFilterRecipientUser`

NewCallFilterRecipientUser instantiates a new CallFilterRecipientUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallFilterRecipientUserWithDefaults

`func NewCallFilterRecipientUserWithDefaults() *CallFilterRecipientUser`

NewCallFilterRecipientUserWithDefaults instantiates a new CallFilterRecipientUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *CallFilterRecipientUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *CallFilterRecipientUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *CallFilterRecipientUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *CallFilterRecipientUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *CallFilterRecipientUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *CallFilterRecipientUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *CallFilterRecipientUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *CallFilterRecipientUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *CallFilterRecipientUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CallFilterRecipientUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CallFilterRecipientUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CallFilterRecipientUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *CallFilterRecipientUser) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CallFilterRecipientUser) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CallFilterRecipientUser) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CallFilterRecipientUser) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
