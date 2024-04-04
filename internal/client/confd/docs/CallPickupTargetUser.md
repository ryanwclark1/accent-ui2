# CallPickupTargetUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Timeout** | Pointer to **int32** |  | [optional]

## Methods

### NewCallPickupTargetUser

`func NewCallPickupTargetUser() *CallPickupTargetUser`

NewCallPickupTargetUser instantiates a new CallPickupTargetUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPickupTargetUserWithDefaults

`func NewCallPickupTargetUserWithDefaults() *CallPickupTargetUser`

NewCallPickupTargetUserWithDefaults instantiates a new CallPickupTargetUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *CallPickupTargetUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *CallPickupTargetUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *CallPickupTargetUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *CallPickupTargetUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *CallPickupTargetUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *CallPickupTargetUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *CallPickupTargetUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *CallPickupTargetUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *CallPickupTargetUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CallPickupTargetUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CallPickupTargetUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CallPickupTargetUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *CallPickupTargetUser) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *CallPickupTargetUser) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *CallPickupTargetUser) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *CallPickupTargetUser) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
