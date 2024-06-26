/*
accent-confd

Confd exposes an API for managing core resources on a Accent server such as users, extensions, devices, voicemails, queues, etc. Resources can be associated together to provide additional functionality. For example: By associating a voicemail with a user, calls will automatically fallback on to the voicemail when the user cannot answer.  Implementation notes ====================  Errors ------  Responses containing errors will have a status code in the 400 or 500 class. A list of error messages will be returned in the body of the response as a JSON-encoded array:  ~~~ [     \"Input error - User not found\",     \"Resource error - User not associated to a line\" ] ~~~   Updating resources via PUT --------------------------  When updating a resource, all fields become optional. In other words, only values that have been changed need to be sent to the server. Please note that this behavior may change in future versions of the API.

API version: 1.1
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package confd

import (
	"bytes"
	"encoding/json"
	"fmt"
)

// checks if the SkillRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SkillRule{}

// SkillRule struct for SkillRule
type SkillRule struct {
	Id *int32 `json:"id,omitempty"`
	// The name to identify the skill rule
	Name  string          `json:"name"`
	Rules []SkillRuleRule `json:"rules,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
}

type _SkillRule SkillRule

// NewSkillRule instantiates a new SkillRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSkillRule(name string) *SkillRule {
	this := SkillRule{}
	this.Name = name
	return &this
}

// NewSkillRuleWithDefaults instantiates a new SkillRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSkillRuleWithDefaults() *SkillRule {
	this := SkillRule{}
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *SkillRule) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillRule) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *SkillRule) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *SkillRule) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value
func (o *SkillRule) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *SkillRule) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *SkillRule) SetName(v string) {
	o.Name = v
}

// GetRules returns the Rules field value if set, zero value otherwise.
func (o *SkillRule) GetRules() []SkillRuleRule {
	if o == nil || IsNil(o.Rules) {
		var ret []SkillRuleRule
		return ret
	}
	return o.Rules
}

// GetRulesOk returns a tuple with the Rules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillRule) GetRulesOk() ([]SkillRuleRule, bool) {
	if o == nil || IsNil(o.Rules) {
		return nil, false
	}
	return o.Rules, true
}

// HasRules returns a boolean if a field has been set.
func (o *SkillRule) HasRules() bool {
	if o != nil && !IsNil(o.Rules) {
		return true
	}

	return false
}

// SetRules gets a reference to the given []SkillRuleRule and assigns it to the Rules field.
func (o *SkillRule) SetRules(v []SkillRuleRule) {
	o.Rules = v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *SkillRule) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SkillRule) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *SkillRule) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *SkillRule) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

func (o SkillRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SkillRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	toSerialize["name"] = o.Name
	if !IsNil(o.Rules) {
		toSerialize["rules"] = o.Rules
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	return toSerialize, nil
}

func (o *SkillRule) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varSkillRule := _SkillRule{}

	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.DisallowUnknownFields()
	err = decoder.Decode(&varSkillRule)

	if err != nil {
		return err
	}

	*o = SkillRule(varSkillRule)

	return err
}

type NullableSkillRule struct {
	value *SkillRule
	isSet bool
}

func (v NullableSkillRule) Get() *SkillRule {
	return v.value
}

func (v *NullableSkillRule) Set(val *SkillRule) {
	v.value = val
	v.isSet = true
}

func (v NullableSkillRule) IsSet() bool {
	return v.isSet
}

func (v *NullableSkillRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSkillRule(val *SkillRule) *NullableSkillRule {
	return &NullableSkillRule{value: val, isSet: true}
}

func (v NullableSkillRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSkillRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
