/*
accent-call-logd

Consult call logs from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calllogd

import (
	"encoding/json"
)

// checks if the QueueStatistic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueStatistic{}

// QueueStatistic struct for QueueStatistic
type QueueStatistic struct {
	// Number of calls that were abandoned while they were waiting for an answer.
	Abandoned *int32 `json:"abandoned,omitempty"`
	// Number of calls answered by an agent.
	Answered *int32 `json:"answered,omitempty"`
	// The number of answered called over (received calls - closed calls)
	AnsweredRate *float32 `json:"answered_rate,omitempty"`
	// The average waiting time of calls
	AverageWaitingTime *int32 `json:"average_waiting_time,omitempty"`
	// Number of calls received when no agent was available, when there was no agent to take the call, when the join an empty queue condition is reached, or when the drop callers if no agent condition is reached.
	Blocked *int32 `json:"blocked,omitempty"`
	// Number of calls received when the queue was closed.
	Closed *int32 `json:"closed,omitempty"`
	// Start of the statistic interval.
	From *string `json:"from,omitempty"`
	// Number of calls that reached the ring timeout delay.
	NotAnswered *int32 `json:"not_answered,omitempty"`
	// Percentage based on the number of calls answered in less than the defined quality of service threshold over the number of answered calls.
	QualityOfService *float32 `json:"quality_of_service,omitempty"`
	// ID of the corresponding queue.
	QueueId *int32 `json:"queue_id,omitempty"`
	// Name of the corresponding queue.
	QueueName *string `json:"queue_name,omitempty"`
	// Total number of calls received in the interval.
	Received *int32 `json:"received,omitempty"`
	// Number of calls received when the queue was full or when one of the diversion parameter was reached.
	Saturated *int32 `json:"saturated,omitempty"`
	// Tenant UUID of the corresponding queue.
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// End of the statistic interval.
	Until *string `json:"until,omitempty"`
}

// NewQueueStatistic instantiates a new QueueStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueStatistic() *QueueStatistic {
	this := QueueStatistic{}
	return &this
}

// NewQueueStatisticWithDefaults instantiates a new QueueStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueStatisticWithDefaults() *QueueStatistic {
	this := QueueStatistic{}
	return &this
}

// GetAbandoned returns the Abandoned field value if set, zero value otherwise.
func (o *QueueStatistic) GetAbandoned() int32 {
	if o == nil || IsNil(o.Abandoned) {
		var ret int32
		return ret
	}
	return *o.Abandoned
}

// GetAbandonedOk returns a tuple with the Abandoned field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetAbandonedOk() (*int32, bool) {
	if o == nil || IsNil(o.Abandoned) {
		return nil, false
	}
	return o.Abandoned, true
}

// HasAbandoned returns a boolean if a field has been set.
func (o *QueueStatistic) HasAbandoned() bool {
	if o != nil && !IsNil(o.Abandoned) {
		return true
	}

	return false
}

// SetAbandoned gets a reference to the given int32 and assigns it to the Abandoned field.
func (o *QueueStatistic) SetAbandoned(v int32) {
	o.Abandoned = &v
}

// GetAnswered returns the Answered field value if set, zero value otherwise.
func (o *QueueStatistic) GetAnswered() int32 {
	if o == nil || IsNil(o.Answered) {
		var ret int32
		return ret
	}
	return *o.Answered
}

// GetAnsweredOk returns a tuple with the Answered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetAnsweredOk() (*int32, bool) {
	if o == nil || IsNil(o.Answered) {
		return nil, false
	}
	return o.Answered, true
}

// HasAnswered returns a boolean if a field has been set.
func (o *QueueStatistic) HasAnswered() bool {
	if o != nil && !IsNil(o.Answered) {
		return true
	}

	return false
}

// SetAnswered gets a reference to the given int32 and assigns it to the Answered field.
func (o *QueueStatistic) SetAnswered(v int32) {
	o.Answered = &v
}

// GetAnsweredRate returns the AnsweredRate field value if set, zero value otherwise.
func (o *QueueStatistic) GetAnsweredRate() float32 {
	if o == nil || IsNil(o.AnsweredRate) {
		var ret float32
		return ret
	}
	return *o.AnsweredRate
}

// GetAnsweredRateOk returns a tuple with the AnsweredRate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetAnsweredRateOk() (*float32, bool) {
	if o == nil || IsNil(o.AnsweredRate) {
		return nil, false
	}
	return o.AnsweredRate, true
}

// HasAnsweredRate returns a boolean if a field has been set.
func (o *QueueStatistic) HasAnsweredRate() bool {
	if o != nil && !IsNil(o.AnsweredRate) {
		return true
	}

	return false
}

// SetAnsweredRate gets a reference to the given float32 and assigns it to the AnsweredRate field.
func (o *QueueStatistic) SetAnsweredRate(v float32) {
	o.AnsweredRate = &v
}

// GetAverageWaitingTime returns the AverageWaitingTime field value if set, zero value otherwise.
func (o *QueueStatistic) GetAverageWaitingTime() int32 {
	if o == nil || IsNil(o.AverageWaitingTime) {
		var ret int32
		return ret
	}
	return *o.AverageWaitingTime
}

// GetAverageWaitingTimeOk returns a tuple with the AverageWaitingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetAverageWaitingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.AverageWaitingTime) {
		return nil, false
	}
	return o.AverageWaitingTime, true
}

// HasAverageWaitingTime returns a boolean if a field has been set.
func (o *QueueStatistic) HasAverageWaitingTime() bool {
	if o != nil && !IsNil(o.AverageWaitingTime) {
		return true
	}

	return false
}

// SetAverageWaitingTime gets a reference to the given int32 and assigns it to the AverageWaitingTime field.
func (o *QueueStatistic) SetAverageWaitingTime(v int32) {
	o.AverageWaitingTime = &v
}

// GetBlocked returns the Blocked field value if set, zero value otherwise.
func (o *QueueStatistic) GetBlocked() int32 {
	if o == nil || IsNil(o.Blocked) {
		var ret int32
		return ret
	}
	return *o.Blocked
}

// GetBlockedOk returns a tuple with the Blocked field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetBlockedOk() (*int32, bool) {
	if o == nil || IsNil(o.Blocked) {
		return nil, false
	}
	return o.Blocked, true
}

// HasBlocked returns a boolean if a field has been set.
func (o *QueueStatistic) HasBlocked() bool {
	if o != nil && !IsNil(o.Blocked) {
		return true
	}

	return false
}

// SetBlocked gets a reference to the given int32 and assigns it to the Blocked field.
func (o *QueueStatistic) SetBlocked(v int32) {
	o.Blocked = &v
}

// GetClosed returns the Closed field value if set, zero value otherwise.
func (o *QueueStatistic) GetClosed() int32 {
	if o == nil || IsNil(o.Closed) {
		var ret int32
		return ret
	}
	return *o.Closed
}

// GetClosedOk returns a tuple with the Closed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetClosedOk() (*int32, bool) {
	if o == nil || IsNil(o.Closed) {
		return nil, false
	}
	return o.Closed, true
}

// HasClosed returns a boolean if a field has been set.
func (o *QueueStatistic) HasClosed() bool {
	if o != nil && !IsNil(o.Closed) {
		return true
	}

	return false
}

// SetClosed gets a reference to the given int32 and assigns it to the Closed field.
func (o *QueueStatistic) SetClosed(v int32) {
	o.Closed = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *QueueStatistic) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *QueueStatistic) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *QueueStatistic) SetFrom(v string) {
	o.From = &v
}

// GetNotAnswered returns the NotAnswered field value if set, zero value otherwise.
func (o *QueueStatistic) GetNotAnswered() int32 {
	if o == nil || IsNil(o.NotAnswered) {
		var ret int32
		return ret
	}
	return *o.NotAnswered
}

// GetNotAnsweredOk returns a tuple with the NotAnswered field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetNotAnsweredOk() (*int32, bool) {
	if o == nil || IsNil(o.NotAnswered) {
		return nil, false
	}
	return o.NotAnswered, true
}

// HasNotAnswered returns a boolean if a field has been set.
func (o *QueueStatistic) HasNotAnswered() bool {
	if o != nil && !IsNil(o.NotAnswered) {
		return true
	}

	return false
}

// SetNotAnswered gets a reference to the given int32 and assigns it to the NotAnswered field.
func (o *QueueStatistic) SetNotAnswered(v int32) {
	o.NotAnswered = &v
}

// GetQualityOfService returns the QualityOfService field value if set, zero value otherwise.
func (o *QueueStatistic) GetQualityOfService() float32 {
	if o == nil || IsNil(o.QualityOfService) {
		var ret float32
		return ret
	}
	return *o.QualityOfService
}

// GetQualityOfServiceOk returns a tuple with the QualityOfService field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetQualityOfServiceOk() (*float32, bool) {
	if o == nil || IsNil(o.QualityOfService) {
		return nil, false
	}
	return o.QualityOfService, true
}

// HasQualityOfService returns a boolean if a field has been set.
func (o *QueueStatistic) HasQualityOfService() bool {
	if o != nil && !IsNil(o.QualityOfService) {
		return true
	}

	return false
}

// SetQualityOfService gets a reference to the given float32 and assigns it to the QualityOfService field.
func (o *QueueStatistic) SetQualityOfService(v float32) {
	o.QualityOfService = &v
}

// GetQueueId returns the QueueId field value if set, zero value otherwise.
func (o *QueueStatistic) GetQueueId() int32 {
	if o == nil || IsNil(o.QueueId) {
		var ret int32
		return ret
	}
	return *o.QueueId
}

// GetQueueIdOk returns a tuple with the QueueId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetQueueIdOk() (*int32, bool) {
	if o == nil || IsNil(o.QueueId) {
		return nil, false
	}
	return o.QueueId, true
}

// HasQueueId returns a boolean if a field has been set.
func (o *QueueStatistic) HasQueueId() bool {
	if o != nil && !IsNil(o.QueueId) {
		return true
	}

	return false
}

// SetQueueId gets a reference to the given int32 and assigns it to the QueueId field.
func (o *QueueStatistic) SetQueueId(v int32) {
	o.QueueId = &v
}

// GetQueueName returns the QueueName field value if set, zero value otherwise.
func (o *QueueStatistic) GetQueueName() string {
	if o == nil || IsNil(o.QueueName) {
		var ret string
		return ret
	}
	return *o.QueueName
}

// GetQueueNameOk returns a tuple with the QueueName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetQueueNameOk() (*string, bool) {
	if o == nil || IsNil(o.QueueName) {
		return nil, false
	}
	return o.QueueName, true
}

// HasQueueName returns a boolean if a field has been set.
func (o *QueueStatistic) HasQueueName() bool {
	if o != nil && !IsNil(o.QueueName) {
		return true
	}

	return false
}

// SetQueueName gets a reference to the given string and assigns it to the QueueName field.
func (o *QueueStatistic) SetQueueName(v string) {
	o.QueueName = &v
}

// GetReceived returns the Received field value if set, zero value otherwise.
func (o *QueueStatistic) GetReceived() int32 {
	if o == nil || IsNil(o.Received) {
		var ret int32
		return ret
	}
	return *o.Received
}

// GetReceivedOk returns a tuple with the Received field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetReceivedOk() (*int32, bool) {
	if o == nil || IsNil(o.Received) {
		return nil, false
	}
	return o.Received, true
}

// HasReceived returns a boolean if a field has been set.
func (o *QueueStatistic) HasReceived() bool {
	if o != nil && !IsNil(o.Received) {
		return true
	}

	return false
}

// SetReceived gets a reference to the given int32 and assigns it to the Received field.
func (o *QueueStatistic) SetReceived(v int32) {
	o.Received = &v
}

// GetSaturated returns the Saturated field value if set, zero value otherwise.
func (o *QueueStatistic) GetSaturated() int32 {
	if o == nil || IsNil(o.Saturated) {
		var ret int32
		return ret
	}
	return *o.Saturated
}

// GetSaturatedOk returns a tuple with the Saturated field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetSaturatedOk() (*int32, bool) {
	if o == nil || IsNil(o.Saturated) {
		return nil, false
	}
	return o.Saturated, true
}

// HasSaturated returns a boolean if a field has been set.
func (o *QueueStatistic) HasSaturated() bool {
	if o != nil && !IsNil(o.Saturated) {
		return true
	}

	return false
}

// SetSaturated gets a reference to the given int32 and assigns it to the Saturated field.
func (o *QueueStatistic) SetSaturated(v int32) {
	o.Saturated = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *QueueStatistic) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *QueueStatistic) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *QueueStatistic) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *QueueStatistic) GetUntil() string {
	if o == nil || IsNil(o.Until) {
		var ret string
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueStatistic) GetUntilOk() (*string, bool) {
	if o == nil || IsNil(o.Until) {
		return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *QueueStatistic) HasUntil() bool {
	if o != nil && !IsNil(o.Until) {
		return true
	}

	return false
}

// SetUntil gets a reference to the given string and assigns it to the Until field.
func (o *QueueStatistic) SetUntil(v string) {
	o.Until = &v
}

func (o QueueStatistic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueStatistic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Abandoned) {
		toSerialize["abandoned"] = o.Abandoned
	}
	if !IsNil(o.Answered) {
		toSerialize["answered"] = o.Answered
	}
	if !IsNil(o.AnsweredRate) {
		toSerialize["answered_rate"] = o.AnsweredRate
	}
	if !IsNil(o.AverageWaitingTime) {
		toSerialize["average_waiting_time"] = o.AverageWaitingTime
	}
	if !IsNil(o.Blocked) {
		toSerialize["blocked"] = o.Blocked
	}
	if !IsNil(o.Closed) {
		toSerialize["closed"] = o.Closed
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.NotAnswered) {
		toSerialize["not_answered"] = o.NotAnswered
	}
	if !IsNil(o.QualityOfService) {
		toSerialize["quality_of_service"] = o.QualityOfService
	}
	if !IsNil(o.QueueId) {
		toSerialize["queue_id"] = o.QueueId
	}
	if !IsNil(o.QueueName) {
		toSerialize["queue_name"] = o.QueueName
	}
	if !IsNil(o.Received) {
		toSerialize["received"] = o.Received
	}
	if !IsNil(o.Saturated) {
		toSerialize["saturated"] = o.Saturated
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Until) {
		toSerialize["until"] = o.Until
	}
	return toSerialize, nil
}

type NullableQueueStatistic struct {
	value *QueueStatistic
	isSet bool
}

func (v NullableQueueStatistic) Get() *QueueStatistic {
	return v.value
}

func (v *NullableQueueStatistic) Set(val *QueueStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueStatistic(val *QueueStatistic) *NullableQueueStatistic {
	return &NullableQueueStatistic{value: val, isSet: true}
}

func (v NullableQueueStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
