# RoomUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**TenantUuid** | Pointer to **string** | The tenant of the user_uuid. Default to the same tenant as the token owner | [optional]
**Uuid** | **string** |  |
**AccentUuid** | Pointer to **string** | The accent of the tenant_uuid. Default to the same accent as the token owner | [optional]

## Methods

### NewRoomUser

`func NewRoomUser(uuid string, ) *RoomUser`

NewRoomUser instantiates a new RoomUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomUserWithDefaults

`func NewRoomUserWithDefaults() *RoomUser`

NewRoomUserWithDefaults instantiates a new RoomUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTenantUuid

`func (o *RoomUser) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *RoomUser) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *RoomUser) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *RoomUser) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUuid

`func (o *RoomUser) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *RoomUser) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *RoomUser) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### GetAccentUuid

`func (o *RoomUser) GetAccentUuid() string`

GetAccentUuid returns the AccentUuid field if non-nil, zero value otherwise.

### GetAccentUuidOk

`func (o *RoomUser) GetAccentUuidOk() (*string, bool)`

GetAccentUuidOk returns a tuple with the AccentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentUuid

`func (o *RoomUser) SetAccentUuid(v string)`

SetAccentUuid sets AccentUuid field to given value.

### HasAccentUuid

`func (o *RoomUser) HasAccentUuid() bool`

HasAccentUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
