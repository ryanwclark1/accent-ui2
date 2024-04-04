# CallPermissionRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the call permission | [optional] [readonly]
**Name** | **string** | The name of the call permission |
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewCallPermissionRelationBase

`func NewCallPermissionRelationBase(name string, ) *CallPermissionRelationBase`

NewCallPermissionRelationBase instantiates a new CallPermissionRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPermissionRelationBaseWithDefaults

`func NewCallPermissionRelationBaseWithDefaults() *CallPermissionRelationBase`

NewCallPermissionRelationBaseWithDefaults instantiates a new CallPermissionRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *CallPermissionRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallPermissionRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallPermissionRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CallPermissionRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CallPermissionRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallPermissionRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallPermissionRelationBase) SetName(v string)`

SetName sets Name field to given value.

### GetTenantUuid

`func (o *CallPermissionRelationBase) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *CallPermissionRelationBase) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *CallPermissionRelationBase) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *CallPermissionRelationBase) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
