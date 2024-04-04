# AgentStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AgentId** | Pointer to **int32** | ID of the corresponding agent. | [optional]
**AgentNumber** | Pointer to **string** | The number of this agent | [optional]
**ConversationTime** | Pointer to **int32** | The time spent in conversation in seconds | [optional]
**From** | Pointer to **string** | Start of the statistic interval. | [optional]
**LoginTime** | Pointer to **int32** | The time spent logged-in in seconds | [optional]
**PauseTime** | Pointer to **int32** | The time spent in pause in seconds | [optional]
**TenantUuid** | Pointer to **string** | Tenant UUID of the corresponding queue. | [optional]
**Until** | Pointer to **string** | End of the statistic interval. | [optional]
**WrapupTime** | Pointer to **int32** | The time spent in wrapup in seconds | [optional]

## Methods

### NewAgentStatistic

`func NewAgentStatistic() *AgentStatistic`

NewAgentStatistic instantiates a new AgentStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAgentStatisticWithDefaults

`func NewAgentStatisticWithDefaults() *AgentStatistic`

NewAgentStatisticWithDefaults instantiates a new AgentStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAgentId

`func (o *AgentStatistic) GetAgentId() int32`

GetAgentId returns the AgentId field if non-nil, zero value otherwise.

### GetAgentIdOk

`func (o *AgentStatistic) GetAgentIdOk() (*int32, bool)`

GetAgentIdOk returns a tuple with the AgentId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentId

`func (o *AgentStatistic) SetAgentId(v int32)`

SetAgentId sets AgentId field to given value.

### HasAgentId

`func (o *AgentStatistic) HasAgentId() bool`

HasAgentId returns a boolean if a field has been set.

### GetAgentNumber

`func (o *AgentStatistic) GetAgentNumber() string`

GetAgentNumber returns the AgentNumber field if non-nil, zero value otherwise.

### GetAgentNumberOk

`func (o *AgentStatistic) GetAgentNumberOk() (*string, bool)`

GetAgentNumberOk returns a tuple with the AgentNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAgentNumber

`func (o *AgentStatistic) SetAgentNumber(v string)`

SetAgentNumber sets AgentNumber field to given value.

### HasAgentNumber

`func (o *AgentStatistic) HasAgentNumber() bool`

HasAgentNumber returns a boolean if a field has been set.

### GetConversationTime

`func (o *AgentStatistic) GetConversationTime() int32`

GetConversationTime returns the ConversationTime field if non-nil, zero value otherwise.

### GetConversationTimeOk

`func (o *AgentStatistic) GetConversationTimeOk() (*int32, bool)`

GetConversationTimeOk returns a tuple with the ConversationTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConversationTime

`func (o *AgentStatistic) SetConversationTime(v int32)`

SetConversationTime sets ConversationTime field to given value.

### HasConversationTime

`func (o *AgentStatistic) HasConversationTime() bool`

HasConversationTime returns a boolean if a field has been set.

### GetFrom

`func (o *AgentStatistic) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *AgentStatistic) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *AgentStatistic) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *AgentStatistic) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetLoginTime

`func (o *AgentStatistic) GetLoginTime() int32`

GetLoginTime returns the LoginTime field if non-nil, zero value otherwise.

### GetLoginTimeOk

`func (o *AgentStatistic) GetLoginTimeOk() (*int32, bool)`

GetLoginTimeOk returns a tuple with the LoginTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLoginTime

`func (o *AgentStatistic) SetLoginTime(v int32)`

SetLoginTime sets LoginTime field to given value.

### HasLoginTime

`func (o *AgentStatistic) HasLoginTime() bool`

HasLoginTime returns a boolean if a field has been set.

### GetPauseTime

`func (o *AgentStatistic) GetPauseTime() int32`

GetPauseTime returns the PauseTime field if non-nil, zero value otherwise.

### GetPauseTimeOk

`func (o *AgentStatistic) GetPauseTimeOk() (*int32, bool)`

GetPauseTimeOk returns a tuple with the PauseTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPauseTime

`func (o *AgentStatistic) SetPauseTime(v int32)`

SetPauseTime sets PauseTime field to given value.

### HasPauseTime

`func (o *AgentStatistic) HasPauseTime() bool`

HasPauseTime returns a boolean if a field has been set.

### GetTenantUuid

`func (o *AgentStatistic) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *AgentStatistic) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *AgentStatistic) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *AgentStatistic) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUntil

`func (o *AgentStatistic) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *AgentStatistic) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *AgentStatistic) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *AgentStatistic) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

### GetWrapupTime

`func (o *AgentStatistic) GetWrapupTime() int32`

GetWrapupTime returns the WrapupTime field if non-nil, zero value otherwise.

### GetWrapupTimeOk

`func (o *AgentStatistic) GetWrapupTimeOk() (*int32, bool)`

GetWrapupTimeOk returns a tuple with the WrapupTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWrapupTime

`func (o *AgentStatistic) SetWrapupTime(v int32)`

SetWrapupTime sets WrapupTime field to given value.

### HasWrapupTime

`func (o *AgentStatistic) HasWrapupTime() bool`

HasWrapupTime returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
