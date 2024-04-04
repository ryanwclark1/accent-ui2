# AgentsStatistics

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AgentStatistic**](AgentStatistic.md) |  | [optional]
**Total** | Pointer to **int32** |  | [optional]

## Methods

### NewAgentsStatistics

`func NewAgentsStatistics() *AgentsStatistics`

NewAgentsStatistics instantiates a new AgentsStatistics object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentsStatisticsWithDefaults

`func NewAgentsStatisticsWithDefaults() *AgentsStatistics`

NewAgentsStatisticsWithDefaults instantiates a new AgentsStatistics object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AgentsStatistics) GetItems() []AgentStatistic`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentsStatistics) GetItemsOk() (*[]AgentStatistic, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentsStatistics) SetItems(v []AgentStatistic)`

SetItems sets Items field to given value.

### HasItems

`func (o *AgentsStatistics) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *AgentsStatistics) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AgentsStatistics) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AgentsStatistics) SetTotal(v int32)`

SetTotal sets Total field to given value.

### HasTotal

`func (o *AgentsStatistics) HasTotal() bool`

HasTotal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
