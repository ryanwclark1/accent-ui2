# ContextIncallRange

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**End** | Pointer to **string** |  | [optional] [default to "<end>"]
**Start** | **string** |  | [default to "<start>"]
**DidLength** | Pointer to **int32** | The length of the did | [optional]

## Methods

### NewContextIncallRange

`func NewContextIncallRange(start string, ) *ContextIncallRange`

NewContextIncallRange instantiates a new ContextIncallRange object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewContextIncallRangeWithDefaults

`func NewContextIncallRangeWithDefaults() *ContextIncallRange`

NewContextIncallRangeWithDefaults instantiates a new ContextIncallRange object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEnd

`func (o *ContextIncallRange) GetEnd() string`

GetEnd returns the End field if non-nil, zero value otherwise.

### GetEndOk

`func (o *ContextIncallRange) GetEndOk() (*string, bool)`

GetEndOk returns a tuple with the End field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnd

`func (o *ContextIncallRange) SetEnd(v string)`

SetEnd sets End field to given value.

### HasEnd

`func (o *ContextIncallRange) HasEnd() bool`

HasEnd returns a boolean if a field has been set.

### GetStart

`func (o *ContextIncallRange) GetStart() string`

GetStart returns the Start field if non-nil, zero value otherwise.

### GetStartOk

`func (o *ContextIncallRange) GetStartOk() (*string, bool)`

GetStartOk returns a tuple with the Start field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStart

`func (o *ContextIncallRange) SetStart(v string)`

SetStart sets Start field to given value.

### GetDidLength

`func (o *ContextIncallRange) GetDidLength() int32`

GetDidLength returns the DidLength field if non-nil, zero value otherwise.

### GetDidLengthOk

`func (o *ContextIncallRange) GetDidLengthOk() (*int32, bool)`

GetDidLengthOk returns a tuple with the DidLength field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDidLength

`func (o *ContextIncallRange) SetDidLength(v int32)`

SetDidLength sets DidLength field to given value.

### HasDidLength

`func (o *ContextIncallRange) HasDidLength() bool`

HasDidLength returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
