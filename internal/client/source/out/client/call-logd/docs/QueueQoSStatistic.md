# QueueQoSStatistic

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**From** | Pointer to **string** | Start of the statistic interval. | [optional]
**QualityOfService** | Pointer to [**[]QueueQoSStatisticQualityOfServiceInner**](QueueQoSStatisticQualityOfServiceInner.md) |  | [optional]
**QueueId** | Pointer to **int32** | ID of the corresponding queue. | [optional]
**QueueName** | Pointer to **string** | Name of the corresponding queue. | [optional]
**TenantUuid** | Pointer to **string** | Tenant UUID of the corresponding queue. | [optional]
**Until** | Pointer to **string** | End of the statistic interval. | [optional]

## Methods

### NewQueueQoSStatistic

`func NewQueueQoSStatistic() *QueueQoSStatistic`

NewQueueQoSStatistic instantiates a new QueueQoSStatistic object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueQoSStatisticWithDefaults

`func NewQueueQoSStatisticWithDefaults() *QueueQoSStatistic`

NewQueueQoSStatisticWithDefaults instantiates a new QueueQoSStatistic object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFrom

`func (o *QueueQoSStatistic) GetFrom() string`

GetFrom returns the From field if non-nil, zero value otherwise.

### GetFromOk

`func (o *QueueQoSStatistic) GetFromOk() (*string, bool)`

GetFromOk returns a tuple with the From field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFrom

`func (o *QueueQoSStatistic) SetFrom(v string)`

SetFrom sets From field to given value.

### HasFrom

`func (o *QueueQoSStatistic) HasFrom() bool`

HasFrom returns a boolean if a field has been set.

### GetQualityOfService

`func (o *QueueQoSStatistic) GetQualityOfService() []QueueQoSStatisticQualityOfServiceInner`

GetQualityOfService returns the QualityOfService field if non-nil, zero value otherwise.

### GetQualityOfServiceOk

`func (o *QueueQoSStatistic) GetQualityOfServiceOk() (*[]QueueQoSStatisticQualityOfServiceInner, bool)`

GetQualityOfServiceOk returns a tuple with the QualityOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQualityOfService

`func (o *QueueQoSStatistic) SetQualityOfService(v []QueueQoSStatisticQualityOfServiceInner)`

SetQualityOfService sets QualityOfService field to given value.

### HasQualityOfService

`func (o *QueueQoSStatistic) HasQualityOfService() bool`

HasQualityOfService returns a boolean if a field has been set.

### GetQueueId

`func (o *QueueQoSStatistic) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *QueueQoSStatistic) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *QueueQoSStatistic) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *QueueQoSStatistic) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

### GetQueueName

`func (o *QueueQoSStatistic) GetQueueName() string`

GetQueueName returns the QueueName field if non-nil, zero value otherwise.

### GetQueueNameOk

`func (o *QueueQoSStatistic) GetQueueNameOk() (*string, bool)`

GetQueueNameOk returns a tuple with the QueueName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueName

`func (o *QueueQoSStatistic) SetQueueName(v string)`

SetQueueName sets QueueName field to given value.

### HasQueueName

`func (o *QueueQoSStatistic) HasQueueName() bool`

HasQueueName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *QueueQoSStatistic) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *QueueQoSStatistic) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *QueueQoSStatistic) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *QueueQoSStatistic) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetUntil

`func (o *QueueQoSStatistic) GetUntil() string`

GetUntil returns the Until field if non-nil, zero value otherwise.

### GetUntilOk

`func (o *QueueQoSStatistic) GetUntilOk() (*string, bool)`

GetUntilOk returns a tuple with the Until field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUntil

`func (o *QueueQoSStatistic) SetUntil(v string)`

SetUntil sets Until field to given value.

### HasUntil

`func (o *QueueQoSStatistic) HasUntil() bool`

HasUntil returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
