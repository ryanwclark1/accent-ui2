# ComponentWithStatus

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to [**StatusValue**](StatusValue.md) |  | [optional]

## Methods

### NewComponentWithStatus

`func NewComponentWithStatus() *ComponentWithStatus`

NewComponentWithStatus instantiates a new ComponentWithStatus object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewComponentWithStatusWithDefaults

`func NewComponentWithStatusWithDefaults() *ComponentWithStatus`

NewComponentWithStatusWithDefaults instantiates a new ComponentWithStatus object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *ComponentWithStatus) GetStatus() StatusValue`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *ComponentWithStatus) GetStatusOk() (*StatusValue, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *ComponentWithStatus) SetStatus(v StatusValue)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *ComponentWithStatus) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
