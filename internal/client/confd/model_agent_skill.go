/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"encoding/json"
)

// checks if the AgentSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentSkill{}

// AgentSkill struct for AgentSkill
type AgentSkill struct {
	// Skill's weight.
	SkillWeight *int32 `json:"skill_weight,omitempty"`
}

// NewAgentSkill instantiates a new AgentSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentSkill() *AgentSkill {
	this := AgentSkill{}
	return &this
}

// NewAgentSkillWithDefaults instantiates a new AgentSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentSkillWithDefaults() *AgentSkill {
	this := AgentSkill{}
	return &this
}

// GetSkillWeight returns the SkillWeight field value if set, zero value otherwise.
func (o *AgentSkill) GetSkillWeight() int32 {
	if o == nil || IsNil(o.SkillWeight) {
		var ret int32
		return ret
	}
	return *o.SkillWeight
}

// GetSkillWeightOk returns a tuple with the SkillWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentSkill) GetSkillWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SkillWeight) {
		return nil, false
	}
	return o.SkillWeight, true
}

// HasSkillWeight returns a boolean if a field has been set.
func (o *AgentSkill) HasSkillWeight() bool {
	if o != nil && !IsNil(o.SkillWeight) {
		return true
	}

	return false
}

// SetSkillWeight gets a reference to the given int32 and assigns it to the SkillWeight field.
func (o *AgentSkill) SetSkillWeight(v int32) {
	o.SkillWeight = &v
}

func (o AgentSkill) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.SkillWeight) {
		toSerialize["skill_weight"] = o.SkillWeight
	}
	return toSerialize, nil
}

type NullableAgentSkill struct {
	value *AgentSkill
	isSet bool
}

func (v NullableAgentSkill) Get() *AgentSkill {
	return v.value
}

func (v *NullableAgentSkill) Set(val *AgentSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentSkill(val *AgentSkill) *NullableAgentSkill {
	return &NullableAgentSkill{value: val, isSet: true}
}

func (v NullableAgentSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
