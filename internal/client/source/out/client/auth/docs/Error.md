# Error

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Reason** | Pointer to **[]string** |  | [optional]
**StatusCode** | Pointer to **int32** |  | [optional]
**Timestamp** | Pointer to **[]string** |  | [optional]

## Methods

### NewError

`func NewError() *Error`

NewError instantiates a new Error object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewErrorWithDefaults

`func NewErrorWithDefaults() *Error`

NewErrorWithDefaults instantiates a new Error object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetReason

`func (o *Error) GetReason() []string`

GetReason returns the Reason field if non-nil, zero value otherwise.

### GetReasonOk

`func (o *Error) GetReasonOk() (*[]string, bool)`

GetReasonOk returns a tuple with the Reason field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetReason

`func (o *Error) SetReason(v []string)`

SetReason sets Reason field to given value.

### HasReason

`func (o *Error) HasReason() bool`

HasReason returns a boolean if a field has been set.

### GetStatusCode

`func (o *Error) GetStatusCode() int32`

GetStatusCode returns the StatusCode field if non-nil, zero value otherwise.

### GetStatusCodeOk

`func (o *Error) GetStatusCodeOk() (*int32, bool)`

GetStatusCodeOk returns a tuple with the StatusCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatusCode

`func (o *Error) SetStatusCode(v int32)`

SetStatusCode sets StatusCode field to given value.

### HasStatusCode

`func (o *Error) HasStatusCode() bool`

HasStatusCode returns a boolean if a field has been set.

### GetTimestamp

`func (o *Error) GetTimestamp() []string`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *Error) GetTimestampOk() (*[]string, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *Error) SetTimestamp(v []string)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *Error) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
