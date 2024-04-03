# QueuesStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]QueueStatistic**](QueueStatistic.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewQueuesStatistics

`func NewQueuesStatistics() *QueuesStatistics`

NewQueuesStatistics instantiates a new QueuesStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueuesStatisticsWithDefaults

`func NewQueuesStatisticsWithDefaults() *QueuesStatistics`

NewQueuesStatisticsWithDefaults instantiates a new QueuesStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *QueuesStatistics) GetItems() []QueueStatistic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *QueuesStatistics) GetItemsOk() (*[]QueueStatistic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *QueuesStatistics) SetItems(v []QueueStatistic)`

SetItems sets Items field to given value.

### HasItems

`func (o *QueuesStatistics) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *QueuesStatistics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *QueuesStatistics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *QueuesStatistics) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *QueuesStatistics) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
