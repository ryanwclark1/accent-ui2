# Context

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the context | [optional] [readonly]
**Name** | Pointer to **string** | The name used by Asterisk | [optional] [readonly]
**Uuid** | Pointer to **string** | The UUID of the context | [optional] [readonly]
**Contexts** | Pointer to [**[]ContextRelationBase**](ContextRelationBase.md) |  | [optional] [readonly]
**ConferenceRoomRanges** | Pointer to [**[]ContextRange**](ContextRange.md) |  | [optional]
**Description** | Pointer to **string** | Additional information about the context | [optional]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**GroupRanges** | Pointer to [**[]ContextRange**](ContextRange.md) |  | [optional]
**IncallRanges** | Pointer to [**[]ContextIncallRange**](ContextIncallRange.md) |  | [optional]
**Label** | Pointer to **string** | The label of the context | [optional]
**QueueRanges** | Pointer to [**[]ContextRange**](ContextRange.md) |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Type** | Pointer to **string** |  | [optional] [default to "internal"]
**UserRanges** | Pointer to [**[]ContextRange**](ContextRange.md) |  | [optional]

## Methods

### NewContext

`func NewContext() *Context`

NewContext instantiates a new Context object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextWithDefaults

`func NewContextWithDefaults() *Context`

NewContextWithDefaults instantiates a new Context object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Context) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Context) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Context) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Context) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Context) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Context) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Context) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Context) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *Context) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Context) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Context) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Context) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetContexts

`func (o *Context) GetContexts() []ContextRelationBase`

GetContexts returns the Contexts field if non-nil, zero value otherwise.

### GetContextsOk

`func (o *Context) GetContextsOk() (*[]ContextRelationBase, bool)`

GetContextsOk returns a tuple with the Contexts field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContexts

`func (o *Context) SetContexts(v []ContextRelationBase)`

SetContexts sets Contexts field to given value.

### HasContexts

`func (o *Context) HasContexts() bool`

HasContexts returns a boolean if a field has been set.

### GetConferenceRoomRanges

`func (o *Context) GetConferenceRoomRanges() []ContextRange`

GetConferenceRoomRanges returns the ConferenceRoomRanges field if non-nil, zero value otherwise.

### GetConferenceRoomRangesOk

`func (o *Context) GetConferenceRoomRangesOk() (*[]ContextRange, bool)`

GetConferenceRoomRangesOk returns a tuple with the ConferenceRoomRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConferenceRoomRanges

`func (o *Context) SetConferenceRoomRanges(v []ContextRange)`

SetConferenceRoomRanges sets ConferenceRoomRanges field to given value.

### HasConferenceRoomRanges

`func (o *Context) HasConferenceRoomRanges() bool`

HasConferenceRoomRanges returns a boolean if a field has been set.

### GetDescription

`func (o *Context) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Context) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Context) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Context) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Context) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Context) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Context) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Context) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetGroupRanges

`func (o *Context) GetGroupRanges() []ContextRange`

GetGroupRanges returns the GroupRanges field if non-nil, zero value otherwise.

### GetGroupRangesOk

`func (o *Context) GetGroupRangesOk() (*[]ContextRange, bool)`

GetGroupRangesOk returns a tuple with the GroupRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupRanges

`func (o *Context) SetGroupRanges(v []ContextRange)`

SetGroupRanges sets GroupRanges field to given value.

### HasGroupRanges

`func (o *Context) HasGroupRanges() bool`

HasGroupRanges returns a boolean if a field has been set.

### GetIncallRanges

`func (o *Context) GetIncallRanges() []ContextIncallRange`

GetIncallRanges returns the IncallRanges field if non-nil, zero value otherwise.

### GetIncallRangesOk

`func (o *Context) GetIncallRangesOk() (*[]ContextIncallRange, bool)`

GetIncallRangesOk returns a tuple with the IncallRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncallRanges

`func (o *Context) SetIncallRanges(v []ContextIncallRange)`

SetIncallRanges sets IncallRanges field to given value.

### HasIncallRanges

`func (o *Context) HasIncallRanges() bool`

HasIncallRanges returns a boolean if a field has been set.

### GetLabel

`func (o *Context) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Context) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Context) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *Context) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetQueueRanges

`func (o *Context) GetQueueRanges() []ContextRange`

GetQueueRanges returns the QueueRanges field if non-nil, zero value otherwise.

### GetQueueRangesOk

`func (o *Context) GetQueueRangesOk() (*[]ContextRange, bool)`

GetQueueRangesOk returns a tuple with the QueueRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueRanges

`func (o *Context) SetQueueRanges(v []ContextRange)`

SetQueueRanges sets QueueRanges field to given value.

### HasQueueRanges

`func (o *Context) HasQueueRanges() bool`

HasQueueRanges returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Context) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Context) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Context) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Context) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetType

`func (o *Context) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Context) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Context) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *Context) HasType() bool`

HasType returns a boolean if a field has been set.

### GetUserRanges

`func (o *Context) GetUserRanges() []ContextRange`

GetUserRanges returns the UserRanges field if non-nil, zero value otherwise.

### GetUserRangesOk

`func (o *Context) GetUserRangesOk() (*[]ContextRange, bool)`

GetUserRangesOk returns a tuple with the UserRanges field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserRanges

`func (o *Context) SetUserRanges(v []ContextRange)`

SetUserRanges sets UserRanges field to given value.

### HasUserRanges

`func (o *Context) HasUserRanges() bool`

HasUserRanges returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
