# ScheduleExceptionalPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursEnd** | Pointer to **string** | The end hour of the period. Format &lt;HH:MM&gt; | [optional]
**HoursStart** | Pointer to **string** | The start hour of the period. Format &lt;HH:MM&gt; | [optional]
**MonthDays** | Pointer to **[]int32** | The month days of the period. | [optional]
**Months** | Pointer to **[]int32** | The months of the period. Month start with January(1) and end with December(12) | [optional]
**WeekDays** | Pointer to **[]int32** | The week days of the period. Week day start with Monday(1) and end with Sunday(7) | [optional]
**Destination** | Pointer to [**DestinationType**](DestinationType.md) |  | [optional]

## Methods

### NewScheduleExceptionalPeriod

`func NewScheduleExceptionalPeriod() *ScheduleExceptionalPeriod`

NewScheduleExceptionalPeriod instantiates a new ScheduleExceptionalPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleExceptionalPeriodWithDefaults

`func NewScheduleExceptionalPeriodWithDefaults() *ScheduleExceptionalPeriod`

NewScheduleExceptionalPeriodWithDefaults instantiates a new ScheduleExceptionalPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursEnd

`func (o *ScheduleExceptionalPeriod) GetHoursEnd() string`

GetHoursEnd returns the HoursEnd field if non-nil, zero value otherwise.

### GetHoursEndOk

`func (o *ScheduleExceptionalPeriod) GetHoursEndOk() (*string, bool)`

GetHoursEndOk returns a tuple with the HoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursEnd

`func (o *ScheduleExceptionalPeriod) SetHoursEnd(v string)`

SetHoursEnd sets HoursEnd field to given value.

### HasHoursEnd

`func (o *ScheduleExceptionalPeriod) HasHoursEnd() bool`

HasHoursEnd returns a boolean if a field has been set.

### GetHoursStart

`func (o *ScheduleExceptionalPeriod) GetHoursStart() string`

GetHoursStart returns the HoursStart field if non-nil, zero value otherwise.

### GetHoursStartOk

`func (o *ScheduleExceptionalPeriod) GetHoursStartOk() (*string, bool)`

GetHoursStartOk returns a tuple with the HoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursStart

`func (o *ScheduleExceptionalPeriod) SetHoursStart(v string)`

SetHoursStart sets HoursStart field to given value.

### HasHoursStart

`func (o *ScheduleExceptionalPeriod) HasHoursStart() bool`

HasHoursStart returns a boolean if a field has been set.

### GetMonthDays

`func (o *ScheduleExceptionalPeriod) GetMonthDays() []int32`

GetMonthDays returns the MonthDays field if non-nil, zero value otherwise.

### GetMonthDaysOk

`func (o *ScheduleExceptionalPeriod) GetMonthDaysOk() (*[]int32, bool)`

GetMonthDaysOk returns a tuple with the MonthDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDays

`func (o *ScheduleExceptionalPeriod) SetMonthDays(v []int32)`

SetMonthDays sets MonthDays field to given value.

### HasMonthDays

`func (o *ScheduleExceptionalPeriod) HasMonthDays() bool`

HasMonthDays returns a boolean if a field has been set.

### GetMonths

`func (o *ScheduleExceptionalPeriod) GetMonths() []int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ScheduleExceptionalPeriod) GetMonthsOk() (*[]int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ScheduleExceptionalPeriod) SetMonths(v []int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ScheduleExceptionalPeriod) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetWeekDays

`func (o *ScheduleExceptionalPeriod) GetWeekDays() []int32`

GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.

### GetWeekDaysOk

`func (o *ScheduleExceptionalPeriod) GetWeekDaysOk() (*[]int32, bool)`

GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDays

`func (o *ScheduleExceptionalPeriod) SetWeekDays(v []int32)`

SetWeekDays sets WeekDays field to given value.

### HasWeekDays

`func (o *ScheduleExceptionalPeriod) HasWeekDays() bool`

HasWeekDays returns a boolean if a field has been set.

### GetDestination

`func (o *ScheduleExceptionalPeriod) GetDestination() DestinationType`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *ScheduleExceptionalPeriod) GetDestinationOk() (*DestinationType, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *ScheduleExceptionalPeriod) SetDestination(v DestinationType)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *ScheduleExceptionalPeriod) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
