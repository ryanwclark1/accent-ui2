# TrunkRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the trunk | [optional] [readonly]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewTrunkRelationBase

`func NewTrunkRelationBase() *TrunkRelationBase`

NewTrunkRelationBase instantiates a new TrunkRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkRelationBaseWithDefaults

`func NewTrunkRelationBaseWithDefaults() *TrunkRelationBase`

NewTrunkRelationBaseWithDefaults instantiates a new TrunkRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *TrunkRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *TrunkRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *TrunkRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *TrunkRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantUuid

`func (o *TrunkRelationBase) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *TrunkRelationBase) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *TrunkRelationBase) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *TrunkRelationBase) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
