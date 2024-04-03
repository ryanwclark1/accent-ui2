# GroupMemberExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | Pointer to **string** |  | [optional]
**Exten** | Pointer to **string** |  | [optional]
**Priority** | Pointer to **int32** | priority of extension in the group. Only used for linear ring strategy | [optional]

## Methods

### NewGroupMemberExtension

`func NewGroupMemberExtension() *GroupMemberExtension`

NewGroupMemberExtension instantiates a new GroupMemberExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGroupMemberExtensionWithDefaults

`func NewGroupMemberExtensionWithDefaults() *GroupMemberExtension`

NewGroupMemberExtensionWithDefaults instantiates a new GroupMemberExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *GroupMemberExtension) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *GroupMemberExtension) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *GroupMemberExtension) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *GroupMemberExtension) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetExten

`func (o *GroupMemberExtension) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *GroupMemberExtension) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *GroupMemberExtension) SetExten(v string)`

SetExten sets Exten field to given value.

### HasExten

`func (o *GroupMemberExtension) HasExten() bool`

HasExten returns a boolean if a field has been set.

### GetPriority

`func (o *GroupMemberExtension) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *GroupMemberExtension) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *GroupMemberExtension) SetPriority(v int32)`

SetPriority sets Priority field to given value.

### HasPriority

`func (o *GroupMemberExtension) HasPriority() bool`

HasPriority returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
