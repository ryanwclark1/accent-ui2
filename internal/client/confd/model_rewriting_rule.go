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

// checks if the RewritingRule type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RewritingRule{}

// RewritingRule struct for RewritingRule
type RewritingRule struct {
	// The local address matching rule. Supports regular expressions
	Match *string `json:"match,omitempty"`
	// The replacement for the matched rule
	Replacement *string `json:"replacement,omitempty"`
}

// NewRewritingRule instantiates a new RewritingRule object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRewritingRule() *RewritingRule {
	this := RewritingRule{}
	return &this
}

// NewRewritingRuleWithDefaults instantiates a new RewritingRule object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRewritingRuleWithDefaults() *RewritingRule {
	this := RewritingRule{}
	return &this
}

// GetMatch returns the Match field value if set, zero value otherwise.
func (o *RewritingRule) GetMatch() string {
	if o == nil || IsNil(o.Match) {
		var ret string
		return ret
	}
	return *o.Match
}

// GetMatchOk returns a tuple with the Match field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewritingRule) GetMatchOk() (*string, bool) {
	if o == nil || IsNil(o.Match) {
		return nil, false
	}
	return o.Match, true
}

// HasMatch returns a boolean if a field has been set.
func (o *RewritingRule) HasMatch() bool {
	if o != nil && !IsNil(o.Match) {
		return true
	}

	return false
}

// SetMatch gets a reference to the given string and assigns it to the Match field.
func (o *RewritingRule) SetMatch(v string) {
	o.Match = &v
}

// GetReplacement returns the Replacement field value if set, zero value otherwise.
func (o *RewritingRule) GetReplacement() string {
	if o == nil || IsNil(o.Replacement) {
		var ret string
		return ret
	}
	return *o.Replacement
}

// GetReplacementOk returns a tuple with the Replacement field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RewritingRule) GetReplacementOk() (*string, bool) {
	if o == nil || IsNil(o.Replacement) {
		return nil, false
	}
	return o.Replacement, true
}

// HasReplacement returns a boolean if a field has been set.
func (o *RewritingRule) HasReplacement() bool {
	if o != nil && !IsNil(o.Replacement) {
		return true
	}

	return false
}

// SetReplacement gets a reference to the given string and assigns it to the Replacement field.
func (o *RewritingRule) SetReplacement(v string) {
	o.Replacement = &v
}

func (o RewritingRule) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RewritingRule) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Match) {
		toSerialize["match"] = o.Match
	}
	if !IsNil(o.Replacement) {
		toSerialize["replacement"] = o.Replacement
	}
	return toSerialize, nil
}

type NullableRewritingRule struct {
	value *RewritingRule
	isSet bool
}

func (v NullableRewritingRule) Get() *RewritingRule {
	return v.value
}

func (v *NullableRewritingRule) Set(val *RewritingRule) {
	v.value = val
	v.isSet = true
}

func (v NullableRewritingRule) IsSet() bool {
	return v.isSet
}

func (v *NullableRewritingRule) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRewritingRule(val *RewritingRule) *NullableRewritingRule {
	return &NullableRewritingRule{value: val, isSet: true}
}

func (v NullableRewritingRule) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRewritingRule) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
