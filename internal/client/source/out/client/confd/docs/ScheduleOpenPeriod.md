# ScheduleOpenPeriod

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**HoursEnd** | Pointer to **string** | The end hour of the period. Format &lt;HH:MM&gt; | [optional]
**HoursStart** | Pointer to **string** | The start hour of the period. Format &lt;HH:MM&gt; | [optional]
**MonthDays** | Pointer to **[]int32** | The month days of the period. | [optional]
**Months** | Pointer to **[]int32** | The months of the period. Month start with January(1) and end with December(12) | [optional]
**WeekDays** | Pointer to **[]int32** | The week days of the period. Week day start with Monday(1) and end with Sunday(7) | [optional]

## Methods

### NewScheduleOpenPeriod

`func NewScheduleOpenPeriod() *ScheduleOpenPeriod`

NewScheduleOpenPeriod instantiates a new ScheduleOpenPeriod object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewScheduleOpenPeriodWithDefaults

`func NewScheduleOpenPeriodWithDefaults() *ScheduleOpenPeriod`

NewScheduleOpenPeriodWithDefaults instantiates a new ScheduleOpenPeriod object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHoursEnd

`func (o *ScheduleOpenPeriod) GetHoursEnd() string`

GetHoursEnd returns the HoursEnd field if non-nil, zero value otherwise.

### GetHoursEndOk

`func (o *ScheduleOpenPeriod) GetHoursEndOk() (*string, bool)`

GetHoursEndOk returns a tuple with the HoursEnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursEnd

`func (o *ScheduleOpenPeriod) SetHoursEnd(v string)`

SetHoursEnd sets HoursEnd field to given value.

### HasHoursEnd

`func (o *ScheduleOpenPeriod) HasHoursEnd() bool`

HasHoursEnd returns a boolean if a field has been set.

### GetHoursStart

`func (o *ScheduleOpenPeriod) GetHoursStart() string`

GetHoursStart returns the HoursStart field if non-nil, zero value otherwise.

### GetHoursStartOk

`func (o *ScheduleOpenPeriod) GetHoursStartOk() (*string, bool)`

GetHoursStartOk returns a tuple with the HoursStart field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHoursStart

`func (o *ScheduleOpenPeriod) SetHoursStart(v string)`

SetHoursStart sets HoursStart field to given value.

### HasHoursStart

`func (o *ScheduleOpenPeriod) HasHoursStart() bool`

HasHoursStart returns a boolean if a field has been set.

### GetMonthDays

`func (o *ScheduleOpenPeriod) GetMonthDays() []int32`

GetMonthDays returns the MonthDays field if non-nil, zero value otherwise.

### GetMonthDaysOk

`func (o *ScheduleOpenPeriod) GetMonthDaysOk() (*[]int32, bool)`

GetMonthDaysOk returns a tuple with the MonthDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonthDays

`func (o *ScheduleOpenPeriod) SetMonthDays(v []int32)`

SetMonthDays sets MonthDays field to given value.

### HasMonthDays

`func (o *ScheduleOpenPeriod) HasMonthDays() bool`

HasMonthDays returns a boolean if a field has been set.

### GetMonths

`func (o *ScheduleOpenPeriod) GetMonths() []int32`

GetMonths returns the Months field if non-nil, zero value otherwise.

### GetMonthsOk

`func (o *ScheduleOpenPeriod) GetMonthsOk() (*[]int32, bool)`

GetMonthsOk returns a tuple with the Months field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMonths

`func (o *ScheduleOpenPeriod) SetMonths(v []int32)`

SetMonths sets Months field to given value.

### HasMonths

`func (o *ScheduleOpenPeriod) HasMonths() bool`

HasMonths returns a boolean if a field has been set.

### GetWeekDays

`func (o *ScheduleOpenPeriod) GetWeekDays() []int32`

GetWeekDays returns the WeekDays field if non-nil, zero value otherwise.

### GetWeekDaysOk

`func (o *ScheduleOpenPeriod) GetWeekDaysOk() (*[]int32, bool)`

GetWeekDaysOk returns a tuple with the WeekDays field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWeekDays

`func (o *ScheduleOpenPeriod) SetWeekDays(v []int32)`

SetWeekDays sets WeekDays field to given value.

### HasWeekDays

`func (o *ScheduleOpenPeriod) HasWeekDays() bool`

HasWeekDays returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
