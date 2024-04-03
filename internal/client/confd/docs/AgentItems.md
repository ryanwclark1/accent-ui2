# AgentItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Agent**](Agent.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewAgentItems

`func NewAgentItems(total int32, ) *AgentItems`

NewAgentItems instantiates a new AgentItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentItemsWithDefaults

`func NewAgentItemsWithDefaults() *AgentItems`

NewAgentItemsWithDefaults instantiates a new AgentItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AgentItems) GetItems() []Agent`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AgentItems) GetItemsOk() (*[]Agent, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AgentItems) SetItems(v []Agent)`

SetItems sets Items field to given value.

### HasItems

`func (o *AgentItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *AgentItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AgentItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AgentItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
