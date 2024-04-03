# Switchboard

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Incalls** | Pointer to [**[]SwitchboardRelationIncall**](SwitchboardRelationIncall.md) |  | [optional] [readonly]
**Members** | Pointer to [**SwitchboardRelationMemberUsers**](SwitchboardRelationMemberUsers.md) |  | [optional]
**Fallbacks** | Pointer to [**SwitchboardFallbacks**](SwitchboardFallbacks.md) |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**QueueMusicOnHold** | Pointer to **string** |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Timeout** | Pointer to **int32** | Maximum time allowed for a call to be queued. When the timeout expires, the call is redirected to the no-answer fallback. | [optional]
**Uuid** | Pointer to **string** |  | [optional] [readonly]
**WaitingRoomMusicOnHold** | Pointer to **string** |  | [optional]

## Methods

### NewSwitchboard

`func NewSwitchboard() *Switchboard`

NewSwitchboard instantiates a new Switchboard object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSwitchboardWithDefaults

`func NewSwitchboardWithDefaults() *Switchboard`

NewSwitchboardWithDefaults instantiates a new Switchboard object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetIncalls

`func (o *Switchboard) GetIncalls() []SwitchboardRelationIncall`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *Switchboard) GetIncallsOk() (*[]SwitchboardRelationIncall, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *Switchboard) SetIncalls(v []SwitchboardRelationIncall)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *Switchboard) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetMembers

`func (o *Switchboard) GetMembers() SwitchboardRelationMemberUsers`

GetMembers returns the Members field if non-nil, zero value otherwise.

### GetMembersOk

`func (o *Switchboard) GetMembersOk() (*SwitchboardRelationMemberUsers, bool)`

GetMembersOk returns a tuple with the Members field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMembers

`func (o *Switchboard) SetMembers(v SwitchboardRelationMemberUsers)`

SetMembers sets Members field to given value.

### HasMembers

`func (o *Switchboard) HasMembers() bool`

HasMembers returns a boolean if a field has been set.

### GetFallbacks

`func (o *Switchboard) GetFallbacks() SwitchboardFallbacks`

GetFallbacks returns the Fallbacks field if non-nil, zero value otherwise.

### GetFallbacksOk

`func (o *Switchboard) GetFallbacksOk() (*SwitchboardFallbacks, bool)`

GetFallbacksOk returns a tuple with the Fallbacks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFallbacks

`func (o *Switchboard) SetFallbacks(v SwitchboardFallbacks)`

SetFallbacks sets Fallbacks field to given value.

### HasFallbacks

`func (o *Switchboard) HasFallbacks() bool`

HasFallbacks returns a boolean if a field has been set.

### GetName

`func (o *Switchboard) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Switchboard) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Switchboard) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Switchboard) HasName() bool`

HasName returns a boolean if a field has been set.

### GetQueueMusicOnHold

`func (o *Switchboard) GetQueueMusicOnHold() string`

GetQueueMusicOnHold returns the QueueMusicOnHold field if non-nil, zero value otherwise.

### GetQueueMusicOnHoldOk

`func (o *Switchboard) GetQueueMusicOnHoldOk() (*string, bool)`

GetQueueMusicOnHoldOk returns a tuple with the QueueMusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueMusicOnHold

`func (o *Switchboard) SetQueueMusicOnHold(v string)`

SetQueueMusicOnHold sets QueueMusicOnHold field to given value.

### HasQueueMusicOnHold

`func (o *Switchboard) HasQueueMusicOnHold() bool`

HasQueueMusicOnHold returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Switchboard) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Switchboard) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Switchboard) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Switchboard) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetTimeout

`func (o *Switchboard) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *Switchboard) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *Switchboard) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *Switchboard) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetUuid

`func (o *Switchboard) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Switchboard) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Switchboard) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Switchboard) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWaitingRoomMusicOnHold

`func (o *Switchboard) GetWaitingRoomMusicOnHold() string`

GetWaitingRoomMusicOnHold returns the WaitingRoomMusicOnHold field if non-nil, zero value otherwise.

### GetWaitingRoomMusicOnHoldOk

`func (o *Switchboard) GetWaitingRoomMusicOnHoldOk() (*string, bool)`

GetWaitingRoomMusicOnHoldOk returns a tuple with the WaitingRoomMusicOnHold field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWaitingRoomMusicOnHold

`func (o *Switchboard) SetWaitingRoomMusicOnHold(v string)`

SetWaitingRoomMusicOnHold sets WaitingRoomMusicOnHold field to given value.

### HasWaitingRoomMusicOnHold

`func (o *Switchboard) HasWaitingRoomMusicOnHold() bool`

HasWaitingRoomMusicOnHold returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
