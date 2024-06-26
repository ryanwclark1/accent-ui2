/*
accent-chatd

Control your message and presence from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package chatd

import (
	"encoding/json"
)

// checks if the APIError type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &APIError{}

// APIError struct for APIError
type APIError struct {
	// Additional information about the error. The keys are specific to each error.
	Details map[string]interface{} `json:"details,omitempty"`
	// Identifier of the type of error. It is more precise than the HTTP status code.
	ErrorId *string `json:"error_id,omitempty"`
	// Human readable explanation of the error
	Message *string `json:"message,omitempty"`
	// Time when the error occured
	Timestamp *float32 `json:"timestamp,omitempty"`
	// Resource name of the error
	Resource *string `json:"resource,omitempty"`
}

// NewAPIError instantiates a new APIError object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewAPIError() *APIError {
	this := APIError{}
	return &this
}

// NewAPIErrorWithDefaults instantiates a new APIError object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewAPIErrorWithDefaults() *APIError {
	this := APIError{}
	return &this
}

// GetDetails returns the Details field value if set, zero value otherwise.
func (o *APIError) GetDetails() map[string]interface{} {
	if o == nil || IsNil(o.Details) {
		var ret map[string]interface{}
		return ret
	}
	return o.Details
}

// GetDetailsOk returns a tuple with the Details field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetDetailsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.Details) {
		return map[string]interface{}{}, false
	}
	return o.Details, true
}

// HasDetails returns a boolean if a field has been set.
func (o *APIError) HasDetails() bool {
	if o != nil && !IsNil(o.Details) {
		return true
	}

	return false
}

// SetDetails gets a reference to the given map[string]interface{} and assigns it to the Details field.
func (o *APIError) SetDetails(v map[string]interface{}) {
	o.Details = v
}

// GetErrorId returns the ErrorId field value if set, zero value otherwise.
func (o *APIError) GetErrorId() string {
	if o == nil || IsNil(o.ErrorId) {
		var ret string
		return ret
	}
	return *o.ErrorId
}

// GetErrorIdOk returns a tuple with the ErrorId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetErrorIdOk() (*string, bool) {
	if o == nil || IsNil(o.ErrorId) {
		return nil, false
	}
	return o.ErrorId, true
}

// HasErrorId returns a boolean if a field has been set.
func (o *APIError) HasErrorId() bool {
	if o != nil && !IsNil(o.ErrorId) {
		return true
	}

	return false
}

// SetErrorId gets a reference to the given string and assigns it to the ErrorId field.
func (o *APIError) SetErrorId(v string) {
	o.ErrorId = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *APIError) GetMessage() string {
	if o == nil || IsNil(o.Message) {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetMessageOk() (*string, bool) {
	if o == nil || IsNil(o.Message) {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *APIError) HasMessage() bool {
	if o != nil && !IsNil(o.Message) {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *APIError) SetMessage(v string) {
	o.Message = &v
}

// GetTimestamp returns the Timestamp field value if set, zero value otherwise.
func (o *APIError) GetTimestamp() float32 {
	if o == nil || IsNil(o.Timestamp) {
		var ret float32
		return ret
	}
	return *o.Timestamp
}

// GetTimestampOk returns a tuple with the Timestamp field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetTimestampOk() (*float32, bool) {
	if o == nil || IsNil(o.Timestamp) {
		return nil, false
	}
	return o.Timestamp, true
}

// HasTimestamp returns a boolean if a field has been set.
func (o *APIError) HasTimestamp() bool {
	if o != nil && !IsNil(o.Timestamp) {
		return true
	}

	return false
}

// SetTimestamp gets a reference to the given float32 and assigns it to the Timestamp field.
func (o *APIError) SetTimestamp(v float32) {
	o.Timestamp = &v
}

// GetResource returns the Resource field value if set, zero value otherwise.
func (o *APIError) GetResource() string {
	if o == nil || IsNil(o.Resource) {
		var ret string
		return ret
	}
	return *o.Resource
}

// GetResourceOk returns a tuple with the Resource field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *APIError) GetResourceOk() (*string, bool) {
	if o == nil || IsNil(o.Resource) {
		return nil, false
	}
	return o.Resource, true
}

// HasResource returns a boolean if a field has been set.
func (o *APIError) HasResource() bool {
	if o != nil && !IsNil(o.Resource) {
		return true
	}

	return false
}

// SetResource gets a reference to the given string and assigns it to the Resource field.
func (o *APIError) SetResource(v string) {
	o.Resource = &v
}

func (o APIError) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o APIError) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Details) {
		toSerialize["details"] = o.Details
	}
	if !IsNil(o.ErrorId) {
		toSerialize["error_id"] = o.ErrorId
	}
	if !IsNil(o.Message) {
		toSerialize["message"] = o.Message
	}
	if !IsNil(o.Timestamp) {
		toSerialize["timestamp"] = o.Timestamp
	}
	if !IsNil(o.Resource) {
		toSerialize["resource"] = o.Resource
	}
	return toSerialize, nil
}

type NullableAPIError struct {
	value *APIError
	isSet bool
}

func (v NullableAPIError) Get() *APIError {
	return v.value
}

func (v *NullableAPIError) Set(val *APIError) {
	v.value = val
	v.isSet = true
}

func (v NullableAPIError) IsSet() bool {
	return v.isSet
}

func (v *NullableAPIError) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableAPIError(val *APIError) *NullableAPIError {
	return &NullableAPIError{value: val, isSet: true}
}

func (v NullableAPIError) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableAPIError) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
