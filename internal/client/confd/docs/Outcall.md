# Outcall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the outgoing call | [optional] [readonly]
**Name** | Pointer to **string** | The name of the outcall | [optional]
**Trunks** | Pointer to [**[]OutcallRelationTrunk**](OutcallRelationTrunk.md) |  | [optional] [readonly]
**Extensions** | Pointer to [**[]OutcallRelationExtension**](OutcallRelationExtension.md) |  | [optional] [readonly]
**Schedules** | Pointer to [**[]OutcallRelationSchedule**](OutcallRelationSchedule.md) |  | [optional] [readonly]
**CallPermissions** | Pointer to [**[]CallPermissionRelationBase**](CallPermissionRelationBase.md) |  | [optional] [readonly]
**Description** | Pointer to **string** | Additional information about the outgoing call | [optional]
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**InternalCallerId** | Pointer to **bool** | Use the internal caller id | [optional] [default to false]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]
**RingTime** | Pointer to **int32** | Ringing time before hangup | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewOutcall

`func NewOutcall() *Outcall`

NewOutcall instantiates a new Outcall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcallWithDefaults

`func NewOutcallWithDefaults() *Outcall`

NewOutcallWithDefaults instantiates a new Outcall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Outcall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Outcall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Outcall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Outcall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Outcall) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Outcall) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Outcall) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Outcall) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTrunks

`func (o *Outcall) GetTrunks() []OutcallRelationTrunk`

GetTrunks returns the Trunks field if non-nil, zero value otherwise.

### GetTrunksOk

`func (o *Outcall) GetTrunksOk() (*[]OutcallRelationTrunk, bool)`

GetTrunksOk returns a tuple with the Trunks field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunks

`func (o *Outcall) SetTrunks(v []OutcallRelationTrunk)`

SetTrunks sets Trunks field to given value.

### HasTrunks

`func (o *Outcall) HasTrunks() bool`

HasTrunks returns a boolean if a field has been set.

### GetExtensions

`func (o *Outcall) GetExtensions() []OutcallRelationExtension`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *Outcall) GetExtensionsOk() (*[]OutcallRelationExtension, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *Outcall) SetExtensions(v []OutcallRelationExtension)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *Outcall) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetSchedules

`func (o *Outcall) GetSchedules() []OutcallRelationSchedule`

GetSchedules returns the Schedules field if non-nil, zero value otherwise.

### GetSchedulesOk

`func (o *Outcall) GetSchedulesOk() (*[]OutcallRelationSchedule, bool)`

GetSchedulesOk returns a tuple with the Schedules field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSchedules

`func (o *Outcall) SetSchedules(v []OutcallRelationSchedule)`

SetSchedules sets Schedules field to given value.

### HasSchedules

`func (o *Outcall) HasSchedules() bool`

HasSchedules returns a boolean if a field has been set.

### GetCallPermissions

`func (o *Outcall) GetCallPermissions() []CallPermissionRelationBase`

GetCallPermissions returns the CallPermissions field if non-nil, zero value otherwise.

### GetCallPermissionsOk

`func (o *Outcall) GetCallPermissionsOk() (*[]CallPermissionRelationBase, bool)`

GetCallPermissionsOk returns a tuple with the CallPermissions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissions

`func (o *Outcall) SetCallPermissions(v []CallPermissionRelationBase)`

SetCallPermissions sets CallPermissions field to given value.

### HasCallPermissions

`func (o *Outcall) HasCallPermissions() bool`

HasCallPermissions returns a boolean if a field has been set.

### GetDescription

`func (o *Outcall) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Outcall) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Outcall) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Outcall) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *Outcall) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Outcall) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Outcall) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Outcall) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetInternalCallerId

`func (o *Outcall) GetInternalCallerId() bool`

GetInternalCallerId returns the InternalCallerId field if non-nil, zero value otherwise.

### GetInternalCallerIdOk

`func (o *Outcall) GetInternalCallerIdOk() (*bool, bool)`

GetInternalCallerIdOk returns a tuple with the InternalCallerId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInternalCallerId

`func (o *Outcall) SetInternalCallerId(v bool)`

SetInternalCallerId sets InternalCallerId field to given value.

### HasInternalCallerId

`func (o *Outcall) HasInternalCallerId() bool`

HasInternalCallerId returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *Outcall) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *Outcall) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *Outcall) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *Outcall) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

### GetRingTime

`func (o *Outcall) GetRingTime() int32`

GetRingTime returns the RingTime field if non-nil, zero value otherwise.

### GetRingTimeOk

`func (o *Outcall) GetRingTimeOk() (*int32, bool)`

GetRingTimeOk returns a tuple with the RingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTime

`func (o *Outcall) SetRingTime(v int32)`

SetRingTime sets RingTime field to given value.

### HasRingTime

`func (o *Outcall) HasRingTime() bool`

HasRingTime returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Outcall) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Outcall) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Outcall) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Outcall) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
