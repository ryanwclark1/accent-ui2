# PluginServer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Param** | Pointer to **string** | The plugins repository URL | [optional]

## Methods

### NewPluginServer

`func NewPluginServer() *PluginServer`

NewPluginServer instantiates a new PluginServer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPluginServerWithDefaults

`func NewPluginServerWithDefaults() *PluginServer`

NewPluginServerWithDefaults instantiates a new PluginServer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetParam

`func (o *PluginServer) GetParam() string`

GetParam returns the Param field if non-nil, zero value otherwise.

### GetParamOk

`func (o *PluginServer) GetParamOk() (*string, bool)`

GetParamOk returns a tuple with the Param field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetParam

`func (o *PluginServer) SetParam(v string)`

SetParam sets Param field to given value.

### HasParam

`func (o *PluginServer) HasParam() bool`

HasParam returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
