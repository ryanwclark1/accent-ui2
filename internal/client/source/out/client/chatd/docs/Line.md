# Line

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**State** | Pointer to **string** | The current state of the line. | [optional] [readonly]

## Methods

### NewLine

`func NewLine() *Line`

NewLine instantiates a new Line object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineWithDefaults

`func NewLineWithDefaults() *Line`

NewLineWithDefaults instantiates a new Line object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Line) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Line) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Line) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Line) HasId() bool`

HasId returns a boolean if a field has been set.

### GetState

`func (o *Line) GetState() string`

GetState returns the State field if non-nil, zero value otherwise.

### GetStateOk

`func (o *Line) GetStateOk() (*string, bool)`

GetStateOk returns a tuple with the State field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetState

`func (o *Line) SetState(v string)`

SetState sets State field to given value.

### HasState

`func (o *Line) HasState() bool`

HasState returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
