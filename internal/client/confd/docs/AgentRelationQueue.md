# AgentRelationQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the queue | [optional] [readonly]
**Label** | Pointer to **string** | The label of the queue | [optional]
**Name** | Pointer to **string** | The name of the queue. Cannot be modified | [optional]
**Penalty** | Pointer to **int32** | Agent&#39;s penalty. A priority used for distributing calls. | [optional]

## Methods

### NewAgentRelationQueue

`func NewAgentRelationQueue() *AgentRelationQueue`

NewAgentRelationQueue instantiates a new AgentRelationQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentRelationQueueWithDefaults

`func NewAgentRelationQueueWithDefaults() *AgentRelationQueue`

NewAgentRelationQueueWithDefaults instantiates a new AgentRelationQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *AgentRelationQueue) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *AgentRelationQueue) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *AgentRelationQueue) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *AgentRelationQueue) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *AgentRelationQueue) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *AgentRelationQueue) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *AgentRelationQueue) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *AgentRelationQueue) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *AgentRelationQueue) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *AgentRelationQueue) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *AgentRelationQueue) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *AgentRelationQueue) HasName() bool`

HasName returns a boolean if a field has been set.

### GetPenalty

`func (o *AgentRelationQueue) GetPenalty() int32`

GetPenalty returns the Penalty field if non-nil, zero value otherwise.

### GetPenaltyOk

`func (o *AgentRelationQueue) GetPenaltyOk() (*int32, bool)`

GetPenaltyOk returns a tuple with the Penalty field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPenalty

`func (o *AgentRelationQueue) SetPenalty(v int32)`

SetPenalty sets Penalty field to given value.

### HasPenalty

`func (o *AgentRelationQueue) HasPenalty() bool`

HasPenalty returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
