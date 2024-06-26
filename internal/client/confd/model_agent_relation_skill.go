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

// checks if the AgentRelationSkill type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &AgentRelationSkill{}

// AgentRelationSkill struct for AgentRelationSkill
type AgentRelationSkill struct {
	Id *int32 `json:"id,omitempty"`
	// The name to identify the skill
	Name *string `json:"name,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
	// Skill's weight.
	SkillWeight *int32 `json:"skill_weight,omitempty"`
}

// NewAgentRelationSkill instantiates a new AgentRelationSkill object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAgentRelationSkill() *AgentRelationSkill {
	this := AgentRelationSkill{}
	return &this
}

// NewAgentRelationSkillWithDefaults instantiates a new AgentRelationSkill object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAgentRelationSkillWithDefaults() *AgentRelationSkill {
	this := AgentRelationSkill{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *AgentRelationSkill) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRelationSkill) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *AgentRelationSkill) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *AgentRelationSkill) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *AgentRelationSkill) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRelationSkill) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *AgentRelationSkill) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *AgentRelationSkill) SetName(v string) {
	o.Name = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *AgentRelationSkill) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRelationSkill) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *AgentRelationSkill) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *AgentRelationSkill) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

// GetSkillWeight returns the SkillWeight field value if set, zero value otherwise.
func (o *AgentRelationSkill) GetSkillWeight() int32 {
	if o == nil || IsNil(o.SkillWeight) {
		var ret int32
		return ret
	}
	return *o.SkillWeight
}

// GetSkillWeightOk returns a tuple with the SkillWeight field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *AgentRelationSkill) GetSkillWeightOk() (*int32, bool) {
	if o == nil || IsNil(o.SkillWeight) {
		return nil, false
	}
	return o.SkillWeight, true
}

// HasSkillWeight returns a boolean if a field has been set.
func (o *AgentRelationSkill) HasSkillWeight() bool {
	if o != nil && !IsNil(o.SkillWeight) {
		return true
	}

	return false
}

// SetSkillWeight gets a reference to the given int32 and assigns it to the SkillWeight field.
func (o *AgentRelationSkill) SetSkillWeight(v int32) {
	o.SkillWeight = &v
}

func (o AgentRelationSkill) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o AgentRelationSkill) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	if !IsNil(o.SkillWeight) {
		toSerialize["skill_weight"] = o.SkillWeight
	}
	return toSerialize, nil
}

type NullableAgentRelationSkill struct {
	value *AgentRelationSkill
	isSet bool
}

func (v NullableAgentRelationSkill) Get() *AgentRelationSkill {
	return v.value
}

func (v *NullableAgentRelationSkill) Set(val *AgentRelationSkill) {
	v.value = val
	v.isSet = true
}

func (v NullableAgentRelationSkill) IsSet() bool {
	return v.isSet
}

func (v *NullableAgentRelationSkill) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAgentRelationSkill(val *AgentRelationSkill) *NullableAgentRelationSkill {
	return &NullableAgentRelationSkill{value: val, isSet: true}
}

func (v NullableAgentRelationSkill) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAgentRelationSkill) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
