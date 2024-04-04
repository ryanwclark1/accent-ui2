# ConfigPatchItem

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Op** | Pointer to **string** | Patch operation. Supported operations: &#x60;replace&#x60;. | [optional]
**Path** | Pointer to **string** | JSON path to operate on. Supported paths: &#x60;/debug&#x60;. | [optional]
**Value** | Pointer to **map[string]interface{}** | The new value for the operation. Type of value is dependent of &#x60;path&#x60; | [optional]

## Methods

### NewConfigPatchItem

`func NewConfigPatchItem() *ConfigPatchItem`

NewConfigPatchItem instantiates a new ConfigPatchItem object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConfigPatchItemWithDefaults

`func NewConfigPatchItemWithDefaults() *ConfigPatchItem`

NewConfigPatchItemWithDefaults instantiates a new ConfigPatchItem object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOp

`func (o *ConfigPatchItem) GetOp() string`

GetOp returns the Op field if non-nil, zero value otherwise.

### GetOpOk

`func (o *ConfigPatchItem) GetOpOk() (*string, bool)`

GetOpOk returns a tuple with the Op field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOp

`func (o *ConfigPatchItem) SetOp(v string)`

SetOp sets Op field to given value.

### HasOp

`func (o *ConfigPatchItem) HasOp() bool`

HasOp returns a boolean if a field has been set.

### GetPath

`func (o *ConfigPatchItem) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *ConfigPatchItem) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *ConfigPatchItem) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *ConfigPatchItem) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetValue

`func (o *ConfigPatchItem) GetValue() map[string]interface{}`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *ConfigPatchItem) GetValueOk() (*map[string]interface{}, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *ConfigPatchItem) SetValue(v map[string]interface{})`

SetValue sets Value field to given value.

### HasValue

`func (o *ConfigPatchItem) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
