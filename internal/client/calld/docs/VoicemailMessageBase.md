# VoicemailMessageBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerIdName** | Pointer to **string** | The caller&#39;s name (or null if no caller ID name) | [optional] [readonly]
**CallerIdNum** | Pointer to **string** | The caller&#39;s number (or null if no caller ID number) | [optional] [readonly]
**Duration** | Pointer to **int32** | The message&#39;s duration in seconds | [optional] [readonly]
**Id** | Pointer to **string** | The message&#39;s ID | [optional] [readonly]
**Timestamp** | Pointer to **int32** | The time the message was left as a Unix time value | [optional] [readonly]

## Methods

### NewVoicemailMessageBase

`func NewVoicemailMessageBase() *VoicemailMessageBase`

NewVoicemailMessageBase instantiates a new VoicemailMessageBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailMessageBaseWithDefaults

`func NewVoicemailMessageBaseWithDefaults() *VoicemailMessageBase`

NewVoicemailMessageBaseWithDefaults instantiates a new VoicemailMessageBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdName

`func (o *VoicemailMessageBase) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *VoicemailMessageBase) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *VoicemailMessageBase) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *VoicemailMessageBase) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *VoicemailMessageBase) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *VoicemailMessageBase) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *VoicemailMessageBase) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *VoicemailMessageBase) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetDuration

`func (o *VoicemailMessageBase) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoicemailMessageBase) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoicemailMessageBase) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoicemailMessageBase) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *VoicemailMessageBase) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicemailMessageBase) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicemailMessageBase) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoicemailMessageBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *VoicemailMessageBase) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VoicemailMessageBase) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VoicemailMessageBase) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VoicemailMessageBase) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
