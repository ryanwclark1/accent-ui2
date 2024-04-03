# Message

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Alias** | Pointer to **string** | Alias/nickname of the sender | [optional]
**Content** | Pointer to **string** | The content of the message | [optional]
**CreatedAt** | Pointer to **string** | The date of the message&#39;s creation | [optional] [readonly]
**Room** | Pointer to [**RoomRelationBase**](RoomRelationBase.md) |  | [optional]
**TenantUuid** | Pointer to **string** | tenant uuid of the sender | [optional]
**UserUuid** | Pointer to **string** | user uuid of the sender | [optional]
**Uuid** | Pointer to **string** | The UUID of the message | [optional] [readonly]
**AccentUuid** | Pointer to **string** | accent uuid of the sender | [optional]

## Methods

### NewMessage

`func NewMessage() *Message`

NewMessage instantiates a new Message object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMessageWithDefaults

`func NewMessageWithDefaults() *Message`

NewMessageWithDefaults instantiates a new Message object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAlias

`func (o *Message) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Message) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Message) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Message) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetContent

`func (o *Message) GetContent() string`

GetContent returns the Content field if non-nil, zero value otherwise.

### GetContentOk

`func (o *Message) GetContentOk() (*string, bool)`

GetContentOk returns a tuple with the Content field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContent

`func (o *Message) SetContent(v string)`

SetContent sets Content field to given value.

### HasContent

`func (o *Message) HasContent() bool`

HasContent returns a boolean if a field has been set.

### GetCreatedAt

`func (o *Message) GetCreatedAt() string`

GetCreatedAt returns the CreatedAt field if non-nil, zero value otherwise.

### GetCreatedAtOk

`func (o *Message) GetCreatedAtOk() (*string, bool)`

GetCreatedAtOk returns a tuple with the CreatedAt field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreatedAt

`func (o *Message) SetCreatedAt(v string)`

SetCreatedAt sets CreatedAt field to given value.

### HasCreatedAt

`func (o *Message) HasCreatedAt() bool`

HasCreatedAt returns a boolean if a field has been set.

### GetRoom

`func (o *Message) GetRoom() RoomRelationBase`

GetRoom returns the Room field if non-nil, zero value otherwise.

### GetRoomOk

`func (o *Message) GetRoomOk() (*RoomRelationBase, bool)`

GetRoomOk returns a tuple with the Room field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRoom

`func (o *Message) SetRoom(v RoomRelationBase)`

SetRoom sets Room field to given value.

### HasRoom

`func (o *Message) HasRoom() bool`

HasRoom returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Message) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Message) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Message) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Message) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUserUuid

`func (o *Message) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *Message) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *Message) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### HasUserUuid

`func (o *Message) HasUserUuid() bool`

HasUserUuid returns a boolean if a field has been set.

### GetUuid

`func (o *Message) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Message) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Message) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Message) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetAccentUuid

`func (o *Message) GetAccentUuid() string`

GetAccentUuid returns the AccentUuid field if non-nil, zero value otherwise.

### GetAccentUuidOk

`func (o *Message) GetAccentUuidOk() (*string, bool)`

GetAccentUuidOk returns a tuple with the AccentUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccentUuid

`func (o *Message) SetAccentUuid(v string)`

SetAccentUuid sets AccentUuid field to given value.

### HasAccentUuid

`func (o *Message) HasAccentUuid() bool`

HasAccentUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
