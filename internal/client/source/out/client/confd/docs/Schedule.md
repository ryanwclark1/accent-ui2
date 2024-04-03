# Schedule

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** |  | [optional] [readonly]
**Name** | Pointer to **string** | The name to identify the schedule | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**Incalls** | Pointer to [**[]OutcallRelationBase**](OutcallRelationBase.md) |  | [optional] [readonly]
**ClosedDestination** | [**DestinationType**](DestinationType.md) |  |
**Enabled** | Pointer to **bool** |  | [optional] [default to true]
**ExceptionalPeriods** | Pointer to [**[]ScheduleExceptionalPeriod**](ScheduleExceptionalPeriod.md) |  | [optional]
**OpenPeriods** | Pointer to [**[]ScheduleOpenPeriod**](ScheduleOpenPeriod.md) |  | [optional]
**Timezone** | Pointer to **string** | The number of the schedule | [optional]

## Methods

### NewSchedule

`func NewSchedule(closedDestination DestinationType, ) *Schedule`

NewSchedule instantiates a new Schedule object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleWithDefaults

`func NewScheduleWithDefaults() *Schedule`

NewScheduleWithDefaults instantiates a new Schedule object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Schedule) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Schedule) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Schedule) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Schedule) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *Schedule) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Schedule) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Schedule) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Schedule) HasName() bool`

HasName returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Schedule) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Schedule) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Schedule) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Schedule) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetIncalls

`func (o *Schedule) GetIncalls() []OutcallRelationBase`

GetIncalls returns the Incalls field if non-nil, zero value otherwise.

### GetIncallsOk

`func (o *Schedule) GetIncallsOk() (*[]OutcallRelationBase, bool)`

GetIncallsOk returns a tuple with the Incalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncalls

`func (o *Schedule) SetIncalls(v []OutcallRelationBase)`

SetIncalls sets Incalls field to given value.

### HasIncalls

`func (o *Schedule) HasIncalls() bool`

HasIncalls returns a boolean if a field has been set.

### GetClosedDestination

`func (o *Schedule) GetClosedDestination() DestinationType`

GetClosedDestination returns the ClosedDestination field if non-nil, zero value otherwise.

### GetClosedDestinationOk

`func (o *Schedule) GetClosedDestinationOk() (*DestinationType, bool)`

GetClosedDestinationOk returns a tuple with the ClosedDestination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetClosedDestination

`func (o *Schedule) SetClosedDestination(v DestinationType)`

SetClosedDestination sets ClosedDestination field to given value.

### GetEnabled

`func (o *Schedule) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *Schedule) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *Schedule) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *Schedule) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetExceptionalPeriods

`func (o *Schedule) GetExceptionalPeriods() []ScheduleExceptionalPeriod`

GetExceptionalPeriods returns the ExceptionalPeriods field if non-nil, zero value otherwise.

### GetExceptionalPeriodsOk

`func (o *Schedule) GetExceptionalPeriodsOk() (*[]ScheduleExceptionalPeriod, bool)`

GetExceptionalPeriodsOk returns a tuple with the ExceptionalPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExceptionalPeriods

`func (o *Schedule) SetExceptionalPeriods(v []ScheduleExceptionalPeriod)`

SetExceptionalPeriods sets ExceptionalPeriods field to given value.

### HasExceptionalPeriods

`func (o *Schedule) HasExceptionalPeriods() bool`

HasExceptionalPeriods returns a boolean if a field has been set.

### GetOpenPeriods

`func (o *Schedule) GetOpenPeriods() []ScheduleOpenPeriod`

GetOpenPeriods returns the OpenPeriods field if non-nil, zero value otherwise.

### GetOpenPeriodsOk

`func (o *Schedule) GetOpenPeriodsOk() (*[]ScheduleOpenPeriod, bool)`

GetOpenPeriodsOk returns a tuple with the OpenPeriods field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOpenPeriods

`func (o *Schedule) SetOpenPeriods(v []ScheduleOpenPeriod)`

SetOpenPeriods sets OpenPeriods field to given value.

### HasOpenPeriods

`func (o *Schedule) HasOpenPeriods() bool`

HasOpenPeriods returns a boolean if a field has been set.

### GetTimezone

`func (o *Schedule) GetTimezone() string`

GetTimezone returns the Timezone field if non-nil, zero value otherwise.

### GetTimezoneOk

`func (o *Schedule) GetTimezoneOk() (*string, bool)`

GetTimezoneOk returns a tuple with the Timezone field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimezone

`func (o *Schedule) SetTimezone(v string)`

SetTimezone sets Timezone field to given value.

### HasTimezone

`func (o *Schedule) HasTimezone() bool`

HasTimezone returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
