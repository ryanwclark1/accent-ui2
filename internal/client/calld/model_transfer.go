/*
accent-calld

Control your calls from a REST API

API version: 1.0.0
Contact: help@accentvoice.io
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package calld

import (
	"encoding/json"
)

// checks if the Transfer type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Transfer{}

// Transfer struct for Transfer
type Transfer struct {
	// The behavior of the transfer
	Flow *string `json:"flow,omitempty"`
	// Unique identifier of the transfer
	Id *string `json:"id,omitempty"`
	// Call ID of the transfer initiator
	InitiatorCall *string `json:"initiator_call,omitempty"`
	// Tenant UUID of the user who initiated the transfer
	InitiatorTenantUuid *string `json:"initiator_tenant_uuid,omitempty"`
	// UUID of the user who initiated the transfer
	InitiatorUuid *string `json:"initiator_uuid,omitempty"`
	// Call ID of the recipient of the transfer. May be null when the transfer is 'starting'.
	RecipientCall *string `json:"recipient_call,omitempty"`
	// The current step of the transfer
	Status *string `json:"status,omitempty"`
	// Call ID of the call being transferred to someone else
	TransferredCall *string `json:"transferred_call,omitempty"`
}

// NewTransfer instantiates a new Transfer object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTransfer() *Transfer {
	this := Transfer{}
	var flow string = "attended"
	this.Flow = &flow
	return &this
}

// NewTransferWithDefaults instantiates a new Transfer object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTransferWithDefaults() *Transfer {
	this := Transfer{}
	var flow string = "attended"
	this.Flow = &flow
	return &this
}

// GetFlow returns the Flow field value if set, zero value otherwise.
func (o *Transfer) GetFlow() string {
	if o == nil || IsNil(o.Flow) {
		var ret string
		return ret
	}
	return *o.Flow
}

// GetFlowOk returns a tuple with the Flow field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetFlowOk() (*string, bool) {
	if o == nil || IsNil(o.Flow) {
		return nil, false
	}
	return o.Flow, true
}

// HasFlow returns a boolean if a field has been set.
func (o *Transfer) HasFlow() bool {
	if o != nil && !IsNil(o.Flow) {
		return true
	}

	return false
}

// SetFlow gets a reference to the given string and assigns it to the Flow field.
func (o *Transfer) SetFlow(v string) {
	o.Flow = &v
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Transfer) GetId() string {
	if o == nil || IsNil(o.Id) {
		var ret string
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetIdOk() (*string, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Transfer) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given string and assigns it to the Id field.
func (o *Transfer) SetId(v string) {
	o.Id = &v
}

// GetInitiatorCall returns the InitiatorCall field value if set, zero value otherwise.
func (o *Transfer) GetInitiatorCall() string {
	if o == nil || IsNil(o.InitiatorCall) {
		var ret string
		return ret
	}
	return *o.InitiatorCall
}

// GetInitiatorCallOk returns a tuple with the InitiatorCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetInitiatorCallOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorCall) {
		return nil, false
	}
	return o.InitiatorCall, true
}

// HasInitiatorCall returns a boolean if a field has been set.
func (o *Transfer) HasInitiatorCall() bool {
	if o != nil && !IsNil(o.InitiatorCall) {
		return true
	}

	return false
}

// SetInitiatorCall gets a reference to the given string and assigns it to the InitiatorCall field.
func (o *Transfer) SetInitiatorCall(v string) {
	o.InitiatorCall = &v
}

// GetInitiatorTenantUuid returns the InitiatorTenantUuid field value if set, zero value otherwise.
func (o *Transfer) GetInitiatorTenantUuid() string {
	if o == nil || IsNil(o.InitiatorTenantUuid) {
		var ret string
		return ret
	}
	return *o.InitiatorTenantUuid
}

// GetInitiatorTenantUuidOk returns a tuple with the InitiatorTenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetInitiatorTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorTenantUuid) {
		return nil, false
	}
	return o.InitiatorTenantUuid, true
}

// HasInitiatorTenantUuid returns a boolean if a field has been set.
func (o *Transfer) HasInitiatorTenantUuid() bool {
	if o != nil && !IsNil(o.InitiatorTenantUuid) {
		return true
	}

	return false
}

// SetInitiatorTenantUuid gets a reference to the given string and assigns it to the InitiatorTenantUuid field.
func (o *Transfer) SetInitiatorTenantUuid(v string) {
	o.InitiatorTenantUuid = &v
}

// GetInitiatorUuid returns the InitiatorUuid field value if set, zero value otherwise.
func (o *Transfer) GetInitiatorUuid() string {
	if o == nil || IsNil(o.InitiatorUuid) {
		var ret string
		return ret
	}
	return *o.InitiatorUuid
}

// GetInitiatorUuidOk returns a tuple with the InitiatorUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetInitiatorUuidOk() (*string, bool) {
	if o == nil || IsNil(o.InitiatorUuid) {
		return nil, false
	}
	return o.InitiatorUuid, true
}

// HasInitiatorUuid returns a boolean if a field has been set.
func (o *Transfer) HasInitiatorUuid() bool {
	if o != nil && !IsNil(o.InitiatorUuid) {
		return true
	}

	return false
}

// SetInitiatorUuid gets a reference to the given string and assigns it to the InitiatorUuid field.
func (o *Transfer) SetInitiatorUuid(v string) {
	o.InitiatorUuid = &v
}

// GetRecipientCall returns the RecipientCall field value if set, zero value otherwise.
func (o *Transfer) GetRecipientCall() string {
	if o == nil || IsNil(o.RecipientCall) {
		var ret string
		return ret
	}
	return *o.RecipientCall
}

// GetRecipientCallOk returns a tuple with the RecipientCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetRecipientCallOk() (*string, bool) {
	if o == nil || IsNil(o.RecipientCall) {
		return nil, false
	}
	return o.RecipientCall, true
}

// HasRecipientCall returns a boolean if a field has been set.
func (o *Transfer) HasRecipientCall() bool {
	if o != nil && !IsNil(o.RecipientCall) {
		return true
	}

	return false
}

// SetRecipientCall gets a reference to the given string and assigns it to the RecipientCall field.
func (o *Transfer) SetRecipientCall(v string) {
	o.RecipientCall = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *Transfer) GetStatus() string {
	if o == nil || IsNil(o.Status) {
		var ret string
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetStatusOk() (*string, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *Transfer) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given string and assigns it to the Status field.
func (o *Transfer) SetStatus(v string) {
	o.Status = &v
}

// GetTransferredCall returns the TransferredCall field value if set, zero value otherwise.
func (o *Transfer) GetTransferredCall() string {
	if o == nil || IsNil(o.TransferredCall) {
		var ret string
		return ret
	}
	return *o.TransferredCall
}

// GetTransferredCallOk returns a tuple with the TransferredCall field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Transfer) GetTransferredCallOk() (*string, bool) {
	if o == nil || IsNil(o.TransferredCall) {
		return nil, false
	}
	return o.TransferredCall, true
}

// HasTransferredCall returns a boolean if a field has been set.
func (o *Transfer) HasTransferredCall() bool {
	if o != nil && !IsNil(o.TransferredCall) {
		return true
	}

	return false
}

// SetTransferredCall gets a reference to the given string and assigns it to the TransferredCall field.
func (o *Transfer) SetTransferredCall(v string) {
	o.TransferredCall = &v
}

func (o Transfer) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Transfer) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Flow) {
		toSerialize["flow"] = o.Flow
	}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.InitiatorCall) {
		toSerialize["initiator_call"] = o.InitiatorCall
	}
	if !IsNil(o.InitiatorTenantUuid) {
		toSerialize["initiator_tenant_uuid"] = o.InitiatorTenantUuid
	}
	if !IsNil(o.InitiatorUuid) {
		toSerialize["initiator_uuid"] = o.InitiatorUuid
	}
	if !IsNil(o.RecipientCall) {
		toSerialize["recipient_call"] = o.RecipientCall
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if !IsNil(o.TransferredCall) {
		toSerialize["transferred_call"] = o.TransferredCall
	}
	return toSerialize, nil
}

type NullableTransfer struct {
	value *Transfer
	isSet bool
}

func (v NullableTransfer) Get() *Transfer {
	return v.value
}

func (v *NullableTransfer) Set(val *Transfer) {
	v.value = val
	v.isSet = true
}

func (v NullableTransfer) IsSet() bool {
	return v.isSet
}

func (v *NullableTransfer) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTransfer(val *Transfer) *NullableTransfer {
	return &NullableTransfer{value: val, isSet: true}
}

func (v NullableTransfer) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTransfer) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
