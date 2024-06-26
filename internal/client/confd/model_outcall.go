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

// checks if the Outcall type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &Outcall{}

// Outcall struct for Outcall
type Outcall struct {
	// The id of the outgoing call
	Id *int32 `json:"id,omitempty"`
	// The name of the outcall
	Name            *string                      `json:"name,omitempty"`
	Trunks          []OutcallRelationTrunk       `json:"trunks,omitempty"`
	Extensions      []OutcallRelationExtension   `json:"extensions,omitempty"`
	Schedules       []OutcallRelationSchedule    `json:"schedules,omitempty"`
	CallPermissions []CallPermissionRelationBase `json:"call_permissions,omitempty"`
	// Additional information about the outgoing call
	Description *string `json:"description,omitempty"`
	Enabled     *bool   `json:"enabled,omitempty"`
	// Use the internal caller id
	InternalCallerId *bool `json:"internal_caller_id,omitempty"`
	// Name of the subroutine to execute in asterisk before receiving a call
	PreprocessSubroutine *string `json:"preprocess_subroutine,omitempty"`
	// Ringing time before hangup
	RingTime *int32 `json:"ring_time,omitempty"`
	// The UUID of the tenant
	TenantUuid *string `json:"tenant_uuid,omitempty"`
}

// NewOutcall instantiates a new Outcall object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewOutcall() *Outcall {
	this := Outcall{}
	var enabled bool = true
	this.Enabled = &enabled
	var internalCallerId bool = false
	this.InternalCallerId = &internalCallerId
	return &this
}

// NewOutcallWithDefaults instantiates a new Outcall object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewOutcallWithDefaults() *Outcall {
	this := Outcall{}
	var enabled bool = true
	this.Enabled = &enabled
	var internalCallerId bool = false
	this.InternalCallerId = &internalCallerId
	return &this
}

// GetId returns the Id field value if set, zero value otherwise.
func (o *Outcall) GetId() int32 {
	if o == nil || IsNil(o.Id) {
		var ret int32
		return ret
	}
	return *o.Id
}

// GetIdOk returns a tuple with the Id field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetIdOk() (*int32, bool) {
	if o == nil || IsNil(o.Id) {
		return nil, false
	}
	return o.Id, true
}

// HasId returns a boolean if a field has been set.
func (o *Outcall) HasId() bool {
	if o != nil && !IsNil(o.Id) {
		return true
	}

	return false
}

// SetId gets a reference to the given int32 and assigns it to the Id field.
func (o *Outcall) SetId(v int32) {
	o.Id = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Outcall) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Outcall) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Outcall) SetName(v string) {
	o.Name = &v
}

// GetTrunks returns the Trunks field value if set, zero value otherwise.
func (o *Outcall) GetTrunks() []OutcallRelationTrunk {
	if o == nil || IsNil(o.Trunks) {
		var ret []OutcallRelationTrunk
		return ret
	}
	return o.Trunks
}

// GetTrunksOk returns a tuple with the Trunks field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetTrunksOk() ([]OutcallRelationTrunk, bool) {
	if o == nil || IsNil(o.Trunks) {
		return nil, false
	}
	return o.Trunks, true
}

// HasTrunks returns a boolean if a field has been set.
func (o *Outcall) HasTrunks() bool {
	if o != nil && !IsNil(o.Trunks) {
		return true
	}

	return false
}

// SetTrunks gets a reference to the given []OutcallRelationTrunk and assigns it to the Trunks field.
func (o *Outcall) SetTrunks(v []OutcallRelationTrunk) {
	o.Trunks = v
}

// GetExtensions returns the Extensions field value if set, zero value otherwise.
func (o *Outcall) GetExtensions() []OutcallRelationExtension {
	if o == nil || IsNil(o.Extensions) {
		var ret []OutcallRelationExtension
		return ret
	}
	return o.Extensions
}

// GetExtensionsOk returns a tuple with the Extensions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetExtensionsOk() ([]OutcallRelationExtension, bool) {
	if o == nil || IsNil(o.Extensions) {
		return nil, false
	}
	return o.Extensions, true
}

// HasExtensions returns a boolean if a field has been set.
func (o *Outcall) HasExtensions() bool {
	if o != nil && !IsNil(o.Extensions) {
		return true
	}

	return false
}

// SetExtensions gets a reference to the given []OutcallRelationExtension and assigns it to the Extensions field.
func (o *Outcall) SetExtensions(v []OutcallRelationExtension) {
	o.Extensions = v
}

// GetSchedules returns the Schedules field value if set, zero value otherwise.
func (o *Outcall) GetSchedules() []OutcallRelationSchedule {
	if o == nil || IsNil(o.Schedules) {
		var ret []OutcallRelationSchedule
		return ret
	}
	return o.Schedules
}

// GetSchedulesOk returns a tuple with the Schedules field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetSchedulesOk() ([]OutcallRelationSchedule, bool) {
	if o == nil || IsNil(o.Schedules) {
		return nil, false
	}
	return o.Schedules, true
}

// HasSchedules returns a boolean if a field has been set.
func (o *Outcall) HasSchedules() bool {
	if o != nil && !IsNil(o.Schedules) {
		return true
	}

	return false
}

// SetSchedules gets a reference to the given []OutcallRelationSchedule and assigns it to the Schedules field.
func (o *Outcall) SetSchedules(v []OutcallRelationSchedule) {
	o.Schedules = v
}

// GetCallPermissions returns the CallPermissions field value if set, zero value otherwise.
func (o *Outcall) GetCallPermissions() []CallPermissionRelationBase {
	if o == nil || IsNil(o.CallPermissions) {
		var ret []CallPermissionRelationBase
		return ret
	}
	return o.CallPermissions
}

// GetCallPermissionsOk returns a tuple with the CallPermissions field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetCallPermissionsOk() ([]CallPermissionRelationBase, bool) {
	if o == nil || IsNil(o.CallPermissions) {
		return nil, false
	}
	return o.CallPermissions, true
}

// HasCallPermissions returns a boolean if a field has been set.
func (o *Outcall) HasCallPermissions() bool {
	if o != nil && !IsNil(o.CallPermissions) {
		return true
	}

	return false
}

// SetCallPermissions gets a reference to the given []CallPermissionRelationBase and assigns it to the CallPermissions field.
func (o *Outcall) SetCallPermissions(v []CallPermissionRelationBase) {
	o.CallPermissions = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *Outcall) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *Outcall) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *Outcall) SetDescription(v string) {
	o.Description = &v
}

// GetEnabled returns the Enabled field value if set, zero value otherwise.
func (o *Outcall) GetEnabled() bool {
	if o == nil || IsNil(o.Enabled) {
		var ret bool
		return ret
	}
	return *o.Enabled
}

// GetEnabledOk returns a tuple with the Enabled field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetEnabledOk() (*bool, bool) {
	if o == nil || IsNil(o.Enabled) {
		return nil, false
	}
	return o.Enabled, true
}

// HasEnabled returns a boolean if a field has been set.
func (o *Outcall) HasEnabled() bool {
	if o != nil && !IsNil(o.Enabled) {
		return true
	}

	return false
}

// SetEnabled gets a reference to the given bool and assigns it to the Enabled field.
func (o *Outcall) SetEnabled(v bool) {
	o.Enabled = &v
}

// GetInternalCallerId returns the InternalCallerId field value if set, zero value otherwise.
func (o *Outcall) GetInternalCallerId() bool {
	if o == nil || IsNil(o.InternalCallerId) {
		var ret bool
		return ret
	}
	return *o.InternalCallerId
}

// GetInternalCallerIdOk returns a tuple with the InternalCallerId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetInternalCallerIdOk() (*bool, bool) {
	if o == nil || IsNil(o.InternalCallerId) {
		return nil, false
	}
	return o.InternalCallerId, true
}

// HasInternalCallerId returns a boolean if a field has been set.
func (o *Outcall) HasInternalCallerId() bool {
	if o != nil && !IsNil(o.InternalCallerId) {
		return true
	}

	return false
}

// SetInternalCallerId gets a reference to the given bool and assigns it to the InternalCallerId field.
func (o *Outcall) SetInternalCallerId(v bool) {
	o.InternalCallerId = &v
}

// GetPreprocessSubroutine returns the PreprocessSubroutine field value if set, zero value otherwise.
func (o *Outcall) GetPreprocessSubroutine() string {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		var ret string
		return ret
	}
	return *o.PreprocessSubroutine
}

// GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetPreprocessSubroutineOk() (*string, bool) {
	if o == nil || IsNil(o.PreprocessSubroutine) {
		return nil, false
	}
	return o.PreprocessSubroutine, true
}

// HasPreprocessSubroutine returns a boolean if a field has been set.
func (o *Outcall) HasPreprocessSubroutine() bool {
	if o != nil && !IsNil(o.PreprocessSubroutine) {
		return true
	}

	return false
}

// SetPreprocessSubroutine gets a reference to the given string and assigns it to the PreprocessSubroutine field.
func (o *Outcall) SetPreprocessSubroutine(v string) {
	o.PreprocessSubroutine = &v
}

// GetRingTime returns the RingTime field value if set, zero value otherwise.
func (o *Outcall) GetRingTime() int32 {
	if o == nil || IsNil(o.RingTime) {
		var ret int32
		return ret
	}
	return *o.RingTime
}

// GetRingTimeOk returns a tuple with the RingTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetRingTimeOk() (*int32, bool) {
	if o == nil || IsNil(o.RingTime) {
		return nil, false
	}
	return o.RingTime, true
}

// HasRingTime returns a boolean if a field has been set.
func (o *Outcall) HasRingTime() bool {
	if o != nil && !IsNil(o.RingTime) {
		return true
	}

	return false
}

// SetRingTime gets a reference to the given int32 and assigns it to the RingTime field.
func (o *Outcall) SetRingTime(v int32) {
	o.RingTime = &v
}

// GetTenantUuid returns the TenantUuid field value if set, zero value otherwise.
func (o *Outcall) GetTenantUuid() string {
	if o == nil || IsNil(o.TenantUuid) {
		var ret string
		return ret
	}
	return *o.TenantUuid
}

// GetTenantUuidOk returns a tuple with the TenantUuid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Outcall) GetTenantUuidOk() (*string, bool) {
	if o == nil || IsNil(o.TenantUuid) {
		return nil, false
	}
	return o.TenantUuid, true
}

// HasTenantUuid returns a boolean if a field has been set.
func (o *Outcall) HasTenantUuid() bool {
	if o != nil && !IsNil(o.TenantUuid) {
		return true
	}

	return false
}

// SetTenantUuid gets a reference to the given string and assigns it to the TenantUuid field.
func (o *Outcall) SetTenantUuid(v string) {
	o.TenantUuid = &v
}

func (o Outcall) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o Outcall) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Id) {
		toSerialize["id"] = o.Id
	}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Trunks) {
		toSerialize["trunks"] = o.Trunks
	}
	if !IsNil(o.Extensions) {
		toSerialize["extensions"] = o.Extensions
	}
	if !IsNil(o.Schedules) {
		toSerialize["schedules"] = o.Schedules
	}
	if !IsNil(o.CallPermissions) {
		toSerialize["call_permissions"] = o.CallPermissions
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Enabled) {
		toSerialize["enabled"] = o.Enabled
	}
	if !IsNil(o.InternalCallerId) {
		toSerialize["internal_caller_id"] = o.InternalCallerId
	}
	if !IsNil(o.PreprocessSubroutine) {
		toSerialize["preprocess_subroutine"] = o.PreprocessSubroutine
	}
	if !IsNil(o.RingTime) {
		toSerialize["ring_time"] = o.RingTime
	}
	if !IsNil(o.TenantUuid) {
		toSerialize["tenant_uuid"] = o.TenantUuid
	}
	return toSerialize, nil
}

type NullableOutcall struct {
	value *Outcall
	isSet bool
}

func (v NullableOutcall) Get() *Outcall {
	return v.value
}

func (v *NullableOutcall) Set(val *Outcall) {
	v.value = val
	v.isSet = true
}

func (v NullableOutcall) IsSet() bool {
	return v.isSet
}

func (v *NullableOutcall) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutcall(val *Outcall) *NullableOutcall {
	return &NullableOutcall{value: val, isSet: true}
}

func (v NullableOutcall) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutcall) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
