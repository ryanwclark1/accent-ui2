# FuncKeyTemplate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The ID of the funckey template | [optional] [readonly]
**Keys** | Pointer to **map[string]interface{}** |  | [optional]
**Name** | Pointer to **string** |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewFuncKeyTemplate

`func NewFuncKeyTemplate() *FuncKeyTemplate`

NewFuncKeyTemplate instantiates a new FuncKeyTemplate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuncKeyTemplateWithDefaults

`func NewFuncKeyTemplateWithDefaults() *FuncKeyTemplate`

NewFuncKeyTemplateWithDefaults instantiates a new FuncKeyTemplate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *FuncKeyTemplate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *FuncKeyTemplate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *FuncKeyTemplate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *FuncKeyTemplate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetKeys

`func (o *FuncKeyTemplate) GetKeys() map[string]interface{}`

GetKeys returns the Keys field if non-nil, zero value otherwise.

### GetKeysOk

`func (o *FuncKeyTemplate) GetKeysOk() (*map[string]interface{}, bool)`

GetKeysOk returns a tuple with the Keys field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKeys

`func (o *FuncKeyTemplate) SetKeys(v map[string]interface{})`

SetKeys sets Keys field to given value.

### HasKeys

`func (o *FuncKeyTemplate) HasKeys() bool`

HasKeys returns a boolean if a field has been set.

### GetName

`func (o *FuncKeyTemplate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FuncKeyTemplate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FuncKeyTemplate) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FuncKeyTemplate) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *FuncKeyTemplate) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *FuncKeyTemplate) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *FuncKeyTemplate) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *FuncKeyTemplate) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
