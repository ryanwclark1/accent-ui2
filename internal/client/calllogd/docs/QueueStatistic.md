# QueueStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abandoned** | Pointer to **int32** | Number of calls that were abandoned while they were waiting for an answer. | [optional]
**Answered** | Pointer to **int32** | Number of calls answered by an agent. | [optional]
**AnsweredRate** | Pointer to **float32** | The number of answered called over (received calls - closed calls) | [optional]
**AverageWaitingTime** | Pointer to **int32** | The average waiting time of calls | [optional]
**Blocked** | Pointer to **int32** | Number of calls received when no agent was available, when there was no agent to take the call, when the join an empty queue condition is reached, or when the drop callers if no agent condition is reached. | [optional]
**Closed** | Pointer to **int32** | Number of calls received when the queue was closed. | [optional]
**From** | Pointer to **string** | Start of the statistic interval. | [optional]
**NotAnswered** | Pointer to **int32** | Number of calls that reached the ring timeout delay. | [optional]
**QualityOfService** | Pointer to **float32** | Percentage based on the number of calls answered in less than the defined quality of service threshold over the number of answered calls. | [optional]
**QueueId** | Pointer to **int32** | ID of the corresponding queue. | [optional]
**QueueName** | Pointer to **string** | Name of the corresponding queue. | [optional]
**Received** | Pointer to **int32** | Total number of calls received in the interval. | [optional]
**Saturated** | Pointer to **int32** | Number of calls received when the queue was full or when one of the diversion parameter was reached. | [optional]
**TenantUuid** | Pointer to **string** | Tenant UUID of the corresponding queue. | [optional]
**Until** | Pointer to **string** | End of the statistic interval. | [optional]

## Methods

### NewQueueStatistic

`func NewQueueStatistic() *QueueStatistic`

NewQueueStatistic instantiates a new QueueStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueStatisticWithDefaults

`func NewQueueStatisticWithDefaults() *QueueStatistic`

NewQueueStatisticWithDefaults instantiates a new QueueStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandoned

`func (o *QueueStatistic) GetAbandoned() int32`

GetAbandoned returns the Abandoned field if non-nil, zero value otherwise.

### GetAbandonedOk

`func (o *QueueStatistic) GetAbandonedOk() (*int32, bool)`

GetAbandonedOk returns a tuple with the Abandoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandoned

`func (o *QueueStatistic) SetAbandoned(v int32)`

SetAbandoned sets Abandoned field to given value.

### HasAbandoned

`func (o *QueueStatistic) HasAbandoned() bool`

HasAbandoned returns a boolean if a field has been set.

### GetAnswered

`func (o *QueueStatistic) GetAnswered() int32`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *QueueStatistic) GetAnsweredOk() (*int32, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *QueueStatistic) SetAnswered(v int32)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *QueueStatistic) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetAnsweredRate

`func (o *QueueStatistic) GetAnsweredRate() float32`

GetAnsweredRate returns the AnsweredRate field if non-nil, zero value otherwise.

### GetAnsweredRateOk

`func (o *QueueStatistic) GetAnsweredRateOk() (*float32, bool)`

GetAnsweredRateOk returns a tuple with the AnsweredRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnsweredRate

`func (o *QueueStatistic) SetAnsweredRate(v float32)`

SetAnsweredRate sets AnsweredRate field to given value.

### HasAnsweredRate

`func (o *QueueStatistic) HasAnsweredRate() bool`

HasAnsweredRate returns a boolean if a field has been set.

### GetAverageWaitingTime

`func (o *QueueStatistic) GetAverageWaitingTime() int32`

GetAverageWaitingTime returns the AverageWaitingTime field if non-nil, zero value otherwise.

### GetAverageWaitingTimeOk

`func (o *QueueStatistic) GetAverageWaitingTimeOk() (*int32, bool)`

GetAverageWaitingTimeOk returns a tuple with the AverageWaitingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAverageWaitingTime

`func (o *QueueStatistic) SetAverageWaitingTime(v int32)`

SetAverageWaitingTime sets AverageWaitingTime field to given value.

### HasAverageWaitingTime

`func (o *QueueStatistic) HasAverageWaitingTime() bool`

HasAverageWaitingTime returns a boolean if a field has been set.

### GetBlocked

`func (o *QueueStatistic) GetBlocked() int32`

GetBlocked returns the Blocked field if non-nil, zero value otherwise.

### GetBlockedOk

`func (o *QueueStatistic) GetBlockedOk() (*int32, bool)`

GetBlockedOk returns a tuple with the Blocked field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBlocked

`func (o *QueueStatistic) SetBlocked(v int32)`

SetBlocked sets Blocked field to given value.

### HasBlocked

`func (o *QueueStatistic) HasBlocked() bool`

HasBlocked returns a boolean if a field has been set.

### GetClosed

`func (o *QueueStatistic) GetClosed() int32`

GetClosed returns the Closed field if non-nil, zero value otherwise.

### GetClosedOk

`func (o *QueueStatistic) GetClosedOk() (*int32, bool)`

GetClosedOk returns a tuple with the Closed field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosed

`func (o *QueueStatistic) SetClosed(v int32)`

SetClosed sets Closed field to given value.

### HasClosed

`func (o *QueueStatistic) HasClosed() bool`

HasClosed returns a boolean if a field has been set.

### GetFrom

`func (o *QueueStatistic) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *QueueStatistic) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *QueueStatistic) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *QueueStatistic) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetNotAnswered

`func (o *QueueStatistic) GetNotAnswered() int32`

GetNotAnswered returns the NotAnswered field if non-nil, zero value otherwise.

### GetNotAnsweredOk

`func (o *QueueStatistic) GetNotAnsweredOk() (*int32, bool)`

GetNotAnsweredOk returns a tuple with the NotAnswered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNotAnswered

`func (o *QueueStatistic) SetNotAnswered(v int32)`

SetNotAnswered sets NotAnswered field to given value.

### HasNotAnswered

`func (o *QueueStatistic) HasNotAnswered() bool`

HasNotAnswered returns a boolean if a field has been set.

### GetQualityOfService

`func (o *QueueStatistic) GetQualityOfService() float32`

GetQualityOfService returns the QualityOfService field if non-nil, zero value otherwise.

### GetQualityOfServiceOk

`func (o *QueueStatistic) GetQualityOfServiceOk() (*float32, bool)`

GetQualityOfServiceOk returns a tuple with the QualityOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOfService

`func (o *QueueStatistic) SetQualityOfService(v float32)`

SetQualityOfService sets QualityOfService field to given value.

### HasQualityOfService

`func (o *QueueStatistic) HasQualityOfService() bool`

HasQualityOfService returns a boolean if a field has been set.

### GetQueueId

`func (o *QueueStatistic) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *QueueStatistic) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *QueueStatistic) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *QueueStatistic) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetQueueName

`func (o *QueueStatistic) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *QueueStatistic) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *QueueStatistic) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *QueueStatistic) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetReceived

`func (o *QueueStatistic) GetReceived() int32`

GetReceived returns the Received field if non-nil, zero value otherwise.

### GetReceivedOk

`func (o *QueueStatistic) GetReceivedOk() (*int32, bool)`

GetReceivedOk returns a tuple with the Received field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReceived

`func (o *QueueStatistic) SetReceived(v int32)`

SetReceived sets Received field to given value.

### HasReceived

`func (o *QueueStatistic) HasReceived() bool`

HasReceived returns a boolean if a field has been set.

### GetSaturated

`func (o *QueueStatistic) GetSaturated() int32`

GetSaturated returns the Saturated field if non-nil, zero value otherwise.

### GetSaturatedOk

`func (o *QueueStatistic) GetSaturatedOk() (*int32, bool)`

GetSaturatedOk returns a tuple with the Saturated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSaturated

`func (o *QueueStatistic) SetSaturated(v int32)`

SetSaturated sets Saturated field to given value.

### HasSaturated

`func (o *QueueStatistic) HasSaturated() bool`

HasSaturated returns a boolean if a field has been set.

### GetTenantUuid

`func (o *QueueStatistic) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *QueueStatistic) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *QueueStatistic) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *QueueStatistic) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUntil

`func (o *QueueStatistic) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *QueueStatistic) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *QueueStatistic) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *QueueStatistic) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
