# ContextRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] [default to "<end>"]
**Start** | **string** |  | [default to "<start>"]

## Methods

### NewContextRange

`func NewContextRange(start string, ) *ContextRange`

NewContextRange instantiates a new ContextRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextRangeWithDefaults

`func NewContextRangeWithDefaults() *ContextRange`

NewContextRangeWithDefaults instantiates a new ContextRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ContextRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ContextRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ContextRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ContextRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *ContextRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ContextRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ContextRange) SetStart(v string)`

SetStart sets Start field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
