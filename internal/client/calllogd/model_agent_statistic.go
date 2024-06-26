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

// checks if the AgentStatistic type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentStatistic{}

// AgentStatistic struct for AgentStatistic
type AgentStatistic struct {
	// ID of the corresponding agent.
	AgentId *int32 `json:"agent_id,omitempty"`
	// The number of this agent
	AgentNumber *string `json:"agent_number,omitempty"`
	// The time spent in conversation in seconds
	ConversationTime *int32 `json:"conversation_time,omitempty"`
	// Start of the statistic interval.
	From *string `json:"from,omitempty"`
	// The time spent logged-in in seconds
	LoginTime *int32 `json:"login_time,omitempty"`
	// The time spent in pause in seconds
	PauseTime *int32 `json:"pause_time,omitempty"`
	// Tenant UUID of the corresponding queue.
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// End of the statistic interval.
	Until *string `json:"until,omitempty"`
	// The time spent in wrapup in seconds
	WrapupTime *int32 `json:"wrapup_time,omitempty"`
}

// NewAgentStatistic instantiates a new AgentStatistic object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentStatistic() *AgentStatistic {
	this := AgentStatistic{}
	return &this
}

// NewAgentStatisticWithDefaults instantiates a new AgentStatistic object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentStatisticWithDefaults() *AgentStatistic {
	this := AgentStatistic{}
	return &this
}

// GetAgentId returns the AgentId field value if set, zero value otherwise.
func (o *AgentStatistic) GetAgentId() int32 {
	if o == nil || IsNil(o.AgentId) {
		var ret int32
		return ret
	}
	return *o.AgentId
}

// GetAgentIdOk returns a tuple with the AgentId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetAgentIdOk() (*int32, bool) {
	if o == nil || IsNil(o.AgentId) {
		return nil, false
	}
	return o.AgentId, true
}

// HasAgentId returns a boolean if a field has been set.
func (o *AgentStatistic) HasAgentId() bool {
	if o != nil && !IsNil(o.AgentId) {
		return true
	}

	return false
}

// SetAgentId gets a reference to the given int32 and assigns it to the AgentId field.
func (o *AgentStatistic) SetAgentId(v int32) {
	o.AgentId = &v
}

// GetAgentNumber returns the AgentNumber field value if set, zero value otherwise.
func (o *AgentStatistic) GetAgentNumber() string {
	if o == nil || IsNil(o.AgentNumber) {
		var ret string
		return ret
	}
	return *o.AgentNumber
}

// GetAgentNumberOk returns a tuple with the AgentNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetAgentNumberOk() (*string, bool) {
	if o == nil || IsNil(o.AgentNumber) {
		return nil, false
	}
	return o.AgentNumber, true
}

// HasAgentNumber returns a boolean if a field has been set.
func (o *AgentStatistic) HasAgentNumber() bool {
	if o != nil && !IsNil(o.AgentNumber) {
		return true
	}

	return false
}

// SetAgentNumber gets a reference to the given string and assigns it to the AgentNumber field.
func (o *AgentStatistic) SetAgentNumber(v string) {
	o.AgentNumber = &v
}

// GetConversationTime returns the ConversationTime field value if set, zero value otherwise.
func (o *AgentStatistic) GetConversationTime() int32 {
	if o == nil || IsNil(o.ConversationTime) {
		var ret int32
		return ret
	}
	return *o.ConversationTime
}

// GetConversationTimeOk returns a tuple with the ConversationTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetConversationTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.ConversationTime) {
		return nil, false
	}
	return o.ConversationTime, true
}

// HasConversationTime returns a boolean if a field has been set.
func (o *AgentStatistic) HasConversationTime() bool {
	if o != nil && !IsNil(o.ConversationTime) {
		return true
	}

	return false
}

// SetConversationTime gets a reference to the given int32 and assigns it to the ConversationTime field.
func (o *AgentStatistic) SetConversationTime(v int32) {
	o.ConversationTime = &v
}

// GetFrom returns the From field value if set, zero value otherwise.
func (o *AgentStatistic) GetFrom() string {
	if o == nil || IsNil(o.From) {
		var ret string
		return ret
	}
	return *o.From
}

// GetFromOk returns a tuple with the From field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetFromOk() (*string, bool) {
	if o == nil || IsNil(o.From) {
		return nil, false
	}
	return o.From, true
}

// HasFrom returns a boolean if a field has been set.
func (o *AgentStatistic) HasFrom() bool {
	if o != nil && !IsNil(o.From) {
		return true
	}

	return false
}

// SetFrom gets a reference to the given string and assigns it to the From field.
func (o *AgentStatistic) SetFrom(v string) {
	o.From = &v
}

// GetLoginTime returns the LoginTime field value if set, zero value otherwise.
func (o *AgentStatistic) GetLoginTime() int32 {
	if o == nil || IsNil(o.LoginTime) {
		var ret int32
		return ret
	}
	return *o.LoginTime
}

// GetLoginTimeOk returns a tuple with the LoginTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetLoginTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.LoginTime) {
		return nil, false
	}
	return o.LoginTime, true
}

// HasLoginTime returns a boolean if a field has been set.
func (o *AgentStatistic) HasLoginTime() bool {
	if o != nil && !IsNil(o.LoginTime) {
		return true
	}

	return false
}

// SetLoginTime gets a reference to the given int32 and assigns it to the LoginTime field.
func (o *AgentStatistic) SetLoginTime(v int32) {
	o.LoginTime = &v
}

// GetPauseTime returns the PauseTime field value if set, zero value otherwise.
func (o *AgentStatistic) GetPauseTime() int32 {
	if o == nil || IsNil(o.PauseTime) {
		var ret int32
		return ret
	}
	return *o.PauseTime
}

// GetPauseTimeOk returns a tuple with the PauseTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetPauseTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.PauseTime) {
		return nil, false
	}
	return o.PauseTime, true
}

// HasPauseTime returns a boolean if a field has been set.
func (o *AgentStatistic) HasPauseTime() bool {
	if o != nil && !IsNil(o.PauseTime) {
		return true
	}

	return false
}

// SetPauseTime gets a reference to the given int32 and assigns it to the PauseTime field.
func (o *AgentStatistic) SetPauseTime(v int32) {
	o.PauseTime = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *AgentStatistic) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *AgentStatistic) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *AgentStatistic) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetUntil returns the Until field value if set, zero value otherwise.
func (o *AgentStatistic) GetUntil() string {
	if o == nil || IsNil(o.Until) {
		var ret string
		return ret
	}
	return *o.Until
}

// GetUntilOk returns a tuple with the Until field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetUntilOk() (*string, bool) {
	if o == nil || IsNil(o.Until) {
		return nil, false
	}
	return o.Until, true
}

// HasUntil returns a boolean if a field has been set.
func (o *AgentStatistic) HasUntil() bool {
	if o != nil && !IsNil(o.Until) {
		return true
	}

	return false
}

// SetUntil gets a reference to the given string and assigns it to the Until field.
func (o *AgentStatistic) SetUntil(v string) {
	o.Until = &v
}

// GetWrapupTime returns the WrapupTime field value if set, zero value otherwise.
func (o *AgentStatistic) GetWrapupTime() int32 {
	if o == nil || IsNil(o.WrapupTime) {
		var ret int32
		return ret
	}
	return *o.WrapupTime
}

// GetWrapupTimeOk returns a tuple with the WrapupTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentStatistic) GetWrapupTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.WrapupTime) {
		return nil, false
	}
	return o.WrapupTime, true
}

// HasWrapupTime returns a boolean if a field has been set.
func (o *AgentStatistic) HasWrapupTime() bool {
	if o != nil && !IsNil(o.WrapupTime) {
		return true
	}

	return false
}

// SetWrapupTime gets a reference to the given int32 and assigns it to the WrapupTime field.
func (o *AgentStatistic) SetWrapupTime(v int32) {
	o.WrapupTime = &v
}

func (o AgentStatistic) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentStatistic) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.AgentId) {
		toSerialize["agent_id"] = o.AgentId
	}
	if !IsNil(o.AgentNumber) {
		toSerialize["agent_number"] = o.AgentNumber
	}
	if !IsNil(o.ConversationTime) {
		toSerialize["conversation_time"] = o.ConversationTime
	}
	if !IsNil(o.From) {
		toSerialize["from"] = o.From
	}
	if !IsNil(o.LoginTime) {
		toSerialize["login_time"] = o.LoginTime
	}
	if !IsNil(o.PauseTime) {
		toSerialize["pause_time"] = o.PauseTime
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.Until) {
		toSerialize["until"] = o.Until
	}
	if !IsNil(o.WrapupTime) {
		toSerialize["wrapup_time"] = o.WrapupTime
	}
	return toSerialize, nil
}

type NullableAgentStatistic struct {
	value *AgentStatistic
	isSet bool
}

func (v NullableAgentStatistic) Get() *AgentStatistic {
	return v.value
}

func (v *NullableAgentStatistic) Set(val *AgentStatistic) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentStatistic) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentStatistic) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentStatistic(val *AgentStatistic) *NullableAgentStatistic {
	return &NullableAgentStatistic{value: val, isSet: true}
}

func (v NullableAgentStatistic) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentStatistic) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
