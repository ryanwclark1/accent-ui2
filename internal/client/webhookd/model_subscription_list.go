/*
accent-webhookd

Control your webhooks from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package webhookd

import (
	"encoding/json"
)

// checks if the SubscriptionList type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &SubscriptionList{}

// SubscriptionList struct for SubscriptionList
type SubscriptionList struct {
	Items []Subscription `json:"items,omitempty"`
	Total *int32         `json:"total,omitempty"`
}

// NewSubscriptionList instantiates a new SubscriptionList object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewSubscriptionList() *SubscriptionList {
	this := SubscriptionList{}
	return &this
}

// NewSubscriptionListWithDefaults instantiates a new SubscriptionList object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewSubscriptionListWithDefaults() *SubscriptionList {
	this := SubscriptionList{}
	return &this
}

// GetItems returns the Items field value if set, zero value otherwise.
func (o *SubscriptionList) GetItems() []Subscription {
	if o == nil || IsNil(o.Items) {
		var ret []Subscription
		return ret
	}
	return o.Items
}

// GetItemsOk returns a tuple with the Items field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionList) GetItemsOk() ([]Subscription, bool) {
	if o == nil || IsNil(o.Items) {
		return nil, false
	}
	return o.Items, true
}

// HasItems returns a boolean if a field has been set.
func (o *SubscriptionList) HasItems() bool {
	if o != nil && !IsNil(o.Items) {
		return true
	}

	return false
}

// SetItems gets a reference to the given []Subscription and assigns it to the Items field.
func (o *SubscriptionList) SetItems(v []Subscription) {
	o.Items = v
}

// GetTotal returns the Total field value if set, zero value otherwise.
func (o *SubscriptionList) GetTotal() int32 {
	if o == nil || IsNil(o.Total) {
		var ret int32
		return ret
	}
	return *o.Total
}

// GetTotalOk returns a tuple with the Total field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *SubscriptionList) GetTotalOk() (*int32, bool) {
	if o == nil || IsNil(o.Total) {
		return nil, false
	}
	return o.Total, true
}

// HasTotal returns a boolean if a field has been set.
func (o *SubscriptionList) HasTotal() bool {
	if o != nil && !IsNil(o.Total) {
		return true
	}

	return false
}

// SetTotal gets a reference to the given int32 and assigns it to the Total field.
func (o *SubscriptionList) SetTotal(v int32) {
	o.Total = &v
}

func (o SubscriptionList) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o SubscriptionList) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Items) {
		toSerialize["items"] = o.Items
	}
	if !IsNil(o.Total) {
		toSerialize["total"] = o.Total
	}
	return toSerialize, nil
}

type NullableSubscriptionList struct {
	value *SubscriptionList
	isSet bool
}

func (v NullableSubscriptionList) Get() *SubscriptionList {
	return v.value
}

func (v *NullableSubscriptionList) Set(val *SubscriptionList) {
	v.value = val
	v.isSet = true
}

func (v NullableSubscriptionList) IsSet() bool {
	return v.isSet
}

func (v *NullableSubscriptionList) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableSubscriptionList(val *SubscriptionList) *NullableSubscriptionList {
	return &NullableSubscriptionList{value: val, isSet: true}
}

func (v NullableSubscriptionList) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableSubscriptionList) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
