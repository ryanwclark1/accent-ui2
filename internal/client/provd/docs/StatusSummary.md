# StatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**RestApi** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]

## Methods

### NewStatusSummary

`func NewStatusSummary() *StatusSummary`

NewStatusSummary instantiates a new StatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSummaryWithDefaults

`func NewStatusSummaryWithDefaults() *StatusSummary`

NewStatusSummaryWithDefaults instantiates a new StatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRestApi

`func (o *StatusSummary) GetRestApi() ComponentWithStatus`

GetRestApi returns the RestApi field if non-nil, zero value otherwise.

### GetRestApiOk

`func (o *StatusSummary) GetRestApiOk() (*ComponentWithStatus, bool)`

GetRestApiOk returns a tuple with the RestApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApi

`func (o *StatusSummary) SetRestApi(v ComponentWithStatus)`

SetRestApi sets RestApi field to given value.

### HasRestApi

`func (o *StatusSummary) HasRestApi() bool`

HasRestApi returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
