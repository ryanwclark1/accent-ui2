# UserImportErrorErrorsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Details** | Pointer to [**UserImportErrorErrorsInnerDetails**](UserImportErrorErrorsInnerDetails.md) |  | [optional]
**Message** | Pointer to **string** | Error message | [optional]
**Timestamp** | Pointer to **int32** | Time at which the error occurred | [optional]

## Methods

### NewUserImportErrorErrorsInner

`func NewUserImportErrorErrorsInner() *UserImportErrorErrorsInner`

NewUserImportErrorErrorsInner instantiates a new UserImportErrorErrorsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserImportErrorErrorsInnerWithDefaults

`func NewUserImportErrorErrorsInnerWithDefaults() *UserImportErrorErrorsInner`

NewUserImportErrorErrorsInnerWithDefaults instantiates a new UserImportErrorErrorsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDetails

`func (o *UserImportErrorErrorsInner) GetDetails() UserImportErrorErrorsInnerDetails`

GetDetails returns the Details field if non-nil, zero value otherwise.

### GetDetailsOk

`func (o *UserImportErrorErrorsInner) GetDetailsOk() (*UserImportErrorErrorsInnerDetails, bool)`

GetDetailsOk returns a tuple with the Details field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDetails

`func (o *UserImportErrorErrorsInner) SetDetails(v UserImportErrorErrorsInnerDetails)`

SetDetails sets Details field to given value.

### HasDetails

`func (o *UserImportErrorErrorsInner) HasDetails() bool`

HasDetails returns a boolean if a field has been set.

### GetMessage

`func (o *UserImportErrorErrorsInner) GetMessage() string`

GetMessage returns the Message field if non-nil, zero value otherwise.

### GetMessageOk

`func (o *UserImportErrorErrorsInner) GetMessageOk() (*string, bool)`

GetMessageOk returns a tuple with the Message field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMessage

`func (o *UserImportErrorErrorsInner) SetMessage(v string)`

SetMessage sets Message field to given value.

### HasMessage

`func (o *UserImportErrorErrorsInner) HasMessage() bool`

HasMessage returns a boolean if a field has been set.

### GetTimestamp

`func (o *UserImportErrorErrorsInner) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *UserImportErrorErrorsInner) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *UserImportErrorErrorsInner) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *UserImportErrorErrorsInner) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
