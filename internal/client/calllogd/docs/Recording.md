# Recording

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Deleted** | Pointer to **bool** |  | [optional]
**EndTime** | Pointer to **time.Time** |  | [optional]
**Filename** | Pointer to **string** |  | [optional]
**StartTime** | Pointer to **time.Time** |  | [optional]
**Uuid** | Pointer to **string** |  | [optional]

## Methods

### NewRecording

`func NewRecording() *Recording`

NewRecording instantiates a new Recording object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRecordingWithDefaults

`func NewRecordingWithDefaults() *Recording`

NewRecordingWithDefaults instantiates a new Recording object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeleted

`func (o *Recording) GetDeleted() bool`

GetDeleted returns the Deleted field if non-nil, zero value otherwise.

### GetDeletedOk

`func (o *Recording) GetDeletedOk() (*bool, bool)`

GetDeletedOk returns a tuple with the Deleted field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeleted

`func (o *Recording) SetDeleted(v bool)`

SetDeleted sets Deleted field to given value.

### HasDeleted

`func (o *Recording) HasDeleted() bool`

HasDeleted returns a boolean if a field has been set.

### GetEndTime

`func (o *Recording) GetEndTime() time.Time`

GetEndTime returns the EndTime field if non-nil, zero value otherwise.

### GetEndTimeOk

`func (o *Recording) GetEndTimeOk() (*time.Time, bool)`

GetEndTimeOk returns a tuple with the EndTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndTime

`func (o *Recording) SetEndTime(v time.Time)`

SetEndTime sets EndTime field to given value.

### HasEndTime

`func (o *Recording) HasEndTime() bool`

HasEndTime returns a boolean if a field has been set.

### GetFilename

`func (o *Recording) GetFilename() string`

GetFilename returns the Filename field if non-nil, zero value otherwise.

### GetFilenameOk

`func (o *Recording) GetFilenameOk() (*string, bool)`

GetFilenameOk returns a tuple with the Filename field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFilename

`func (o *Recording) SetFilename(v string)`

SetFilename sets Filename field to given value.

### HasFilename

`func (o *Recording) HasFilename() bool`

HasFilename returns a boolean if a field has been set.

### GetStartTime

`func (o *Recording) GetStartTime() time.Time`

GetStartTime returns the StartTime field if non-nil, zero value otherwise.

### GetStartTimeOk

`func (o *Recording) GetStartTimeOk() (*time.Time, bool)`

GetStartTimeOk returns a tuple with the StartTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStartTime

`func (o *Recording) SetStartTime(v time.Time)`

SetStartTime sets StartTime field to given value.

### HasStartTime

`func (o *Recording) HasStartTime() bool`

HasStartTime returns a boolean if a field has been set.

### GetUuid

`func (o *Recording) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Recording) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Recording) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Recording) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
