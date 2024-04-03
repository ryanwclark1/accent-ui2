# QueueQoSStatisticQualityOfServiceInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Abandoned** | Pointer to **int32** | Number of calls that were abandoned while they were waiting for an answer. | [optional]
**Answered** | Pointer to **int32** | Number of calls answered by an agent. | [optional]
**Max** | Pointer to **int32** | Maximum of the QoS interval | [optional]
**Min** | Pointer to **int32** | Minimum of the QoS interval | [optional]

## Methods

### NewQueueQoSStatisticQualityOfServiceInner

`func NewQueueQoSStatisticQualityOfServiceInner() *QueueQoSStatisticQualityOfServiceInner`

NewQueueQoSStatisticQualityOfServiceInner instantiates a new QueueQoSStatisticQualityOfServiceInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueQoSStatisticQualityOfServiceInnerWithDefaults

`func NewQueueQoSStatisticQualityOfServiceInnerWithDefaults() *QueueQoSStatisticQualityOfServiceInner`

NewQueueQoSStatisticQualityOfServiceInnerWithDefaults instantiates a new QueueQoSStatisticQualityOfServiceInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAbandoned

`func (o *QueueQoSStatisticQualityOfServiceInner) GetAbandoned() int32`

GetAbandoned returns the Abandoned field if non-nil, zero value otherwise.

### GetAbandonedOk

`func (o *QueueQoSStatisticQualityOfServiceInner) GetAbandonedOk() (*int32, bool)`

GetAbandonedOk returns a tuple with the Abandoned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAbandoned

`func (o *QueueQoSStatisticQualityOfServiceInner) SetAbandoned(v int32)`

SetAbandoned sets Abandoned field to given value.

### HasAbandoned

`func (o *QueueQoSStatisticQualityOfServiceInner) HasAbandoned() bool`

HasAbandoned returns a boolean if a field has been set.

### GetAnswered

`func (o *QueueQoSStatisticQualityOfServiceInner) GetAnswered() int32`

GetAnswered returns the Answered field if non-nil, zero value otherwise.

### GetAnsweredOk

`func (o *QueueQoSStatisticQualityOfServiceInner) GetAnsweredOk() (*int32, bool)`

GetAnsweredOk returns a tuple with the Answered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAnswered

`func (o *QueueQoSStatisticQualityOfServiceInner) SetAnswered(v int32)`

SetAnswered sets Answered field to given value.

### HasAnswered

`func (o *QueueQoSStatisticQualityOfServiceInner) HasAnswered() bool`

HasAnswered returns a boolean if a field has been set.

### GetMax

`func (o *QueueQoSStatisticQualityOfServiceInner) GetMax() int32`

GetMax returns the Max field if non-nil, zero value otherwise.

### GetMaxOk

`func (o *QueueQoSStatisticQualityOfServiceInner) GetMaxOk() (*int32, bool)`

GetMaxOk returns a tuple with the Max field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMax

`func (o *QueueQoSStatisticQualityOfServiceInner) SetMax(v int32)`

SetMax sets Max field to given value.

### HasMax

`func (o *QueueQoSStatisticQualityOfServiceInner) HasMax() bool`

HasMax returns a boolean if a field has been set.

### GetMin

`func (o *QueueQoSStatisticQualityOfServiceInner) GetMin() int32`

GetMin returns the Min field if non-nil, zero value otherwise.

### GetMinOk

`func (o *QueueQoSStatisticQualityOfServiceInner) GetMinOk() (*int32, bool)`

GetMinOk returns a tuple with the Min field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMin

`func (o *QueueQoSStatisticQualityOfServiceInner) SetMin(v int32)`

SetMin sets Min field to given value.

### HasMin

`func (o *QueueQoSStatisticQualityOfServiceInner) HasMin() bool`

HasMin returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
