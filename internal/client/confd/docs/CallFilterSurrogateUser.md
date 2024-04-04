# CallFilterSurrogateUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Firstname** | Pointer to **string** | User firstname | [optional]
**Lastname** | Pointer to **string** | User lastname | [optional]
**Uuid** | Pointer to **string** | User UUID. This ID is globally unique across multiple Accent instances | [optional] [readonly]
**Exten** | Pointer to **string** |  | [optional]
**MemberId** | Pointer to **int32** |  | [optional]

## Methods

### NewCallFilterSurrogateUser

`func NewCallFilterSurrogateUser() *CallFilterSurrogateUser`

NewCallFilterSurrogateUser instantiates a new CallFilterSurrogateUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallFilterSurrogateUserWithDefaults

`func NewCallFilterSurrogateUserWithDefaults() *CallFilterSurrogateUser`

NewCallFilterSurrogateUserWithDefaults instantiates a new CallFilterSurrogateUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstname

`func (o *CallFilterSurrogateUser) GetFirstname() string`

GetFirstname returns the Firstname field if non-nil, zero value otherwise.

### GetFirstnameOk

`func (o *CallFilterSurrogateUser) GetFirstnameOk() (*string, bool)`

GetFirstnameOk returns a tuple with the Firstname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstname

`func (o *CallFilterSurrogateUser) SetFirstname(v string)`

SetFirstname sets Firstname field to given value.

### HasFirstname

`func (o *CallFilterSurrogateUser) HasFirstname() bool`

HasFirstname returns a boolean if a field has been set.

### GetLastname

`func (o *CallFilterSurrogateUser) GetLastname() string`

GetLastname returns the Lastname field if non-nil, zero value otherwise.

### GetLastnameOk

`func (o *CallFilterSurrogateUser) GetLastnameOk() (*string, bool)`

GetLastnameOk returns a tuple with the Lastname field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastname

`func (o *CallFilterSurrogateUser) SetLastname(v string)`

SetLastname sets Lastname field to given value.

### HasLastname

`func (o *CallFilterSurrogateUser) HasLastname() bool`

HasLastname returns a boolean if a field has been set.

### GetUuid

`func (o *CallFilterSurrogateUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *CallFilterSurrogateUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *CallFilterSurrogateUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *CallFilterSurrogateUser) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetExten

`func (o *CallFilterSurrogateUser) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *CallFilterSurrogateUser) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *CallFilterSurrogateUser) SetExten(v string)`

SetExten sets Exten field to given value.

### HasExten

`func (o *CallFilterSurrogateUser) HasExten() bool`

HasExten returns a boolean if a field has been set.

### GetMemberId

`func (o *CallFilterSurrogateUser) GetMemberId() int32`

GetMemberId returns the MemberId field if non-nil, zero value otherwise.

### GetMemberIdOk

`func (o *CallFilterSurrogateUser) GetMemberIdOk() (*int32, bool)`

GetMemberIdOk returns a tuple with the MemberId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMemberId

`func (o *CallFilterSurrogateUser) SetMemberId(v int32)`

SetMemberId sets MemberId field to given value.

### HasMemberId

`func (o *CallFilterSurrogateUser) HasMemberId() bool`

HasMemberId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
