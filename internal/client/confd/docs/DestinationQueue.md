# DestinationQueue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | **int32** | The id of the queue |
**RingTime** | Pointer to **float32** |  | [optional]
**SkillRuleId** | Pointer to **int32** | The id of the skill rule | [optional]
**SkillRuleVariables** | Pointer to **map[string]interface{}** | key-value where key represents the variable of the skill rule and value represents a value | [optional]
**Type** | **string** | MUST be &#39;queue&#39; |

## Methods

### NewDestinationQueue

`func NewDestinationQueue(queueId int32, type_ string, ) *DestinationQueue`

NewDestinationQueue instantiates a new DestinationQueue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationQueueWithDefaults

`func NewDestinationQueueWithDefaults() *DestinationQueue`

NewDestinationQueueWithDefaults instantiates a new DestinationQueue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *DestinationQueue) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *DestinationQueue) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *DestinationQueue) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### GetRingTime

`func (o *DestinationQueue) GetRingTime() float32`

GetRingTime returns the RingTime field if non-nil, zero value otherwise.

### GetRingTimeOk

`func (o *DestinationQueue) GetRingTimeOk() (*float32, bool)`

GetRingTimeOk returns a tuple with the RingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTime

`func (o *DestinationQueue) SetRingTime(v float32)`

SetRingTime sets RingTime field to given value.

### HasRingTime

`func (o *DestinationQueue) HasRingTime() bool`

HasRingTime returns a boolean if a field has been set.

### GetSkillRuleId

`func (o *DestinationQueue) GetSkillRuleId() int32`

GetSkillRuleId returns the SkillRuleId field if non-nil, zero value otherwise.

### GetSkillRuleIdOk

`func (o *DestinationQueue) GetSkillRuleIdOk() (*int32, bool)`

GetSkillRuleIdOk returns a tuple with the SkillRuleId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillRuleId

`func (o *DestinationQueue) SetSkillRuleId(v int32)`

SetSkillRuleId sets SkillRuleId field to given value.

### HasSkillRuleId

`func (o *DestinationQueue) HasSkillRuleId() bool`

HasSkillRuleId returns a boolean if a field has been set.

### GetSkillRuleVariables

`func (o *DestinationQueue) GetSkillRuleVariables() map[string]interface{}`

GetSkillRuleVariables returns the SkillRuleVariables field if non-nil, zero value otherwise.

### GetSkillRuleVariablesOk

`func (o *DestinationQueue) GetSkillRuleVariablesOk() (*map[string]interface{}, bool)`

GetSkillRuleVariablesOk returns a tuple with the SkillRuleVariables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSkillRuleVariables

`func (o *DestinationQueue) SetSkillRuleVariables(v map[string]interface{})`

SetSkillRuleVariables sets SkillRuleVariables field to given value.

### HasSkillRuleVariables

`func (o *DestinationQueue) HasSkillRuleVariables() bool`

HasSkillRuleVariables returns a boolean if a field has been set.

### GetType

`func (o *DestinationQueue) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationQueue) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationQueue) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
