# FuncKeyDestinationParameter

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Collection** | Pointer to **string** | URL towards a collection of entities representing possible values as a destination | [optional]
**Name** | Pointer to **string** | Parameter name | [optional]
**Values** | Pointer to **[]string** | Array of values to choose from for this parameter | [optional]

## Methods

### NewFuncKeyDestinationParameter

`func NewFuncKeyDestinationParameter() *FuncKeyDestinationParameter`

NewFuncKeyDestinationParameter instantiates a new FuncKeyDestinationParameter object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewFuncKeyDestinationParameterWithDefaults

`func NewFuncKeyDestinationParameterWithDefaults() *FuncKeyDestinationParameter`

NewFuncKeyDestinationParameterWithDefaults instantiates a new FuncKeyDestinationParameter object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCollection

`func (o *FuncKeyDestinationParameter) GetCollection() string`

GetCollection returns the Collection field if non-nil, zero value otherwise.

### GetCollectionOk

`func (o *FuncKeyDestinationParameter) GetCollectionOk() (*string, bool)`

GetCollectionOk returns a tuple with the Collection field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCollection

`func (o *FuncKeyDestinationParameter) SetCollection(v string)`

SetCollection sets Collection field to given value.

### HasCollection

`func (o *FuncKeyDestinationParameter) HasCollection() bool`

HasCollection returns a boolean if a field has been set.

### GetName

`func (o *FuncKeyDestinationParameter) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *FuncKeyDestinationParameter) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *FuncKeyDestinationParameter) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *FuncKeyDestinationParameter) HasName() bool`

HasName returns a boolean if a field has been set.

### GetValues

`func (o *FuncKeyDestinationParameter) GetValues() []string`

GetValues returns the Values field if non-nil, zero value otherwise.

### GetValuesOk

`func (o *FuncKeyDestinationParameter) GetValuesOk() (*[]string, bool)`

GetValuesOk returns a tuple with the Values field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValues

`func (o *FuncKeyDestinationParameter) SetValues(v []string)`

SetValues sets Values field to given value.

### HasValues

`func (o *FuncKeyDestinationParameter) HasValues() bool`

HasValues returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
