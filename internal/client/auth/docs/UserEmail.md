# UserEmail

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Address** | Pointer to **string** |  | [optional]
**Confirmed** | Pointer to **bool** |  | [optional]
**Main** | Pointer to **bool** |  | [optional]

## Methods

### NewUserEmail

`func NewUserEmail() *UserEmail`

NewUserEmail instantiates a new UserEmail object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserEmailWithDefaults

`func NewUserEmailWithDefaults() *UserEmail`

NewUserEmailWithDefaults instantiates a new UserEmail object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAddress

`func (o *UserEmail) GetAddress() string`

GetAddress returns the Address field if non-nil, zero value otherwise.

### GetAddressOk

`func (o *UserEmail) GetAddressOk() (*string, bool)`

GetAddressOk returns a tuple with the Address field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAddress

`func (o *UserEmail) SetAddress(v string)`

SetAddress sets Address field to given value.

### HasAddress

`func (o *UserEmail) HasAddress() bool`

HasAddress returns a boolean if a field has been set.

### GetConfirmed

`func (o *UserEmail) GetConfirmed() bool`

GetConfirmed returns the Confirmed field if non-nil, zero value otherwise.

### GetConfirmedOk

`func (o *UserEmail) GetConfirmedOk() (*bool, bool)`

GetConfirmedOk returns a tuple with the Confirmed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfirmed

`func (o *UserEmail) SetConfirmed(v bool)`

SetConfirmed sets Confirmed field to given value.

### HasConfirmed

`func (o *UserEmail) HasConfirmed() bool`

HasConfirmed returns a boolean if a field has been set.

### GetMain

`func (o *UserEmail) GetMain() bool`

GetMain returns the Main field if non-nil, zero value otherwise.

### GetMainOk

`func (o *UserEmail) GetMainOk() (*bool, bool)`

GetMainOk returns a tuple with the Main field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMain

`func (o *UserEmail) SetMain(v bool)`

SetMain sets Main field to given value.

### HasMain

`func (o *UserEmail) HasMain() bool`

HasMain returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
