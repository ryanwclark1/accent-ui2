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

// checks if the QueueFallbacks type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &QueueFallbacks{}

// QueueFallbacks struct for QueueFallbacks
type QueueFallbacks struct {
	BusyDestination       *DestinationType `json:"busy_destination,omitempty"`
	CongestionDestination *DestinationType `json:"congestion_destination,omitempty"`
	FailDestination       *DestinationType `json:"fail_destination,omitempty"`
	NoanswerDestination   *DestinationType `json:"noanswer_destination,omitempty"`
}

// NewQueueFallbacks instantiates a new QueueFallbacks object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewQueueFallbacks() *QueueFallbacks {
	this := QueueFallbacks{}
	return &this
}

// NewQueueFallbacksWithDefaults instantiates a new QueueFallbacks object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewQueueFallbacksWithDefaults() *QueueFallbacks {
	this := QueueFallbacks{}
	return &this
}

// GetBusyDestination returns the BusyDestination field value if set, zero value otherwise.
func (o *QueueFallbacks) GetBusyDestination() DestinationType {
	if o == nil || IsNil(o.BusyDestination) {
		var ret DestinationType
		return ret
	}
	return *o.BusyDestination
}

// GetBusyDestinationOk returns a tuple with the BusyDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueFallbacks) GetBusyDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.BusyDestination) {
		return nil, false
	}
	return o.BusyDestination, true
}

// HasBusyDestination returns a boolean if a field has been set.
func (o *QueueFallbacks) HasBusyDestination() bool {
	if o != nil && !IsNil(o.BusyDestination) {
		return true
	}

	return false
}

// SetBusyDestination gets a reference to the given DestinationType and assigns it to the BusyDestination field.
func (o *QueueFallbacks) SetBusyDestination(v DestinationType) {
	o.BusyDestination = &v
}

// GetCongestionDestination returns the CongestionDestination field value if set, zero value otherwise.
func (o *QueueFallbacks) GetCongestionDestination() DestinationType {
	if o == nil || IsNil(o.CongestionDestination) {
		var ret DestinationType
		return ret
	}
	return *o.CongestionDestination
}

// GetCongestionDestinationOk returns a tuple with the CongestionDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueFallbacks) GetCongestionDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.CongestionDestination) {
		return nil, false
	}
	return o.CongestionDestination, true
}

// HasCongestionDestination returns a boolean if a field has been set.
func (o *QueueFallbacks) HasCongestionDestination() bool {
	if o != nil && !IsNil(o.CongestionDestination) {
		return true
	}

	return false
}

// SetCongestionDestination gets a reference to the given DestinationType and assigns it to the CongestionDestination field.
func (o *QueueFallbacks) SetCongestionDestination(v DestinationType) {
	o.CongestionDestination = &v
}

// GetFailDestination returns the FailDestination field value if set, zero value otherwise.
func (o *QueueFallbacks) GetFailDestination() DestinationType {
	if o == nil || IsNil(o.FailDestination) {
		var ret DestinationType
		return ret
	}
	return *o.FailDestination
}

// GetFailDestinationOk returns a tuple with the FailDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueFallbacks) GetFailDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.FailDestination) {
		return nil, false
	}
	return o.FailDestination, true
}

// HasFailDestination returns a boolean if a field has been set.
func (o *QueueFallbacks) HasFailDestination() bool {
	if o != nil && !IsNil(o.FailDestination) {
		return true
	}

	return false
}

// SetFailDestination gets a reference to the given DestinationType and assigns it to the FailDestination field.
func (o *QueueFallbacks) SetFailDestination(v DestinationType) {
	o.FailDestination = &v
}

// GetNoanswerDestination returns the NoanswerDestination field value if set, zero value otherwise.
func (o *QueueFallbacks) GetNoanswerDestination() DestinationType {
	if o == nil || IsNil(o.NoanswerDestination) {
		var ret DestinationType
		return ret
	}
	return *o.NoanswerDestination
}

// GetNoanswerDestinationOk returns a tuple with the NoanswerDestination field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *QueueFallbacks) GetNoanswerDestinationOk() (*DestinationType, bool) {
	if o == nil || IsNil(o.NoanswerDestination) {
		return nil, false
	}
	return o.NoanswerDestination, true
}

// HasNoanswerDestination returns a boolean if a field has been set.
func (o *QueueFallbacks) HasNoanswerDestination() bool {
	if o != nil && !IsNil(o.NoanswerDestination) {
		return true
	}

	return false
}

// SetNoanswerDestination gets a reference to the given DestinationType and assigns it to the NoanswerDestination field.
func (o *QueueFallbacks) SetNoanswerDestination(v DestinationType) {
	o.NoanswerDestination = &v
}

func (o QueueFallbacks) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o QueueFallbacks) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.BusyDestination) {
		toSerialize["busy_destination"] = o.BusyDestination
	}
	if !IsNil(o.CongestionDestination) {
		toSerialize["congestion_destination"] = o.CongestionDestination
	}
	if !IsNil(o.FailDestination) {
		toSerialize["fail_destination"] = o.FailDestination
	}
	if !IsNil(o.NoanswerDestination) {
		toSerialize["noanswer_destination"] = o.NoanswerDestination
	}
	return toSerialize, nil
}

type NullableQueueFallbacks struct {
	value *QueueFallbacks
	isSet bool
}

func (v NullableQueueFallbacks) Get() *QueueFallbacks {
	return v.value
}

func (v *NullableQueueFallbacks) Set(val *QueueFallbacks) {
	v.value = val
	v.isSet = true
}

func (v NullableQueueFallbacks) IsSet() bool {
	return v.isSet
}

func (v *NullableQueueFallbacks) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableQueueFallbacks(val *QueueFallbacks) *NullableQueueFallbacks {
	return &NullableQueueFallbacks{value: val, isSet: true}
}

func (v NullableQueueFallbacks) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableQueueFallbacks) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
