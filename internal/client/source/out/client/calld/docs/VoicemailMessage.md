# VoicemailMessage

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerIdName** | Pointer to **string** | The caller&#39;s name (or null if no caller ID name) | [optional] [readonly]
**CallerIdNum** | Pointer to **string** | The caller&#39;s number (or null if no caller ID number) | [optional] [readonly]
**Duration** | Pointer to **int32** | The message&#39;s duration in seconds | [optional] [readonly]
**Id** | Pointer to **string** | The message&#39;s ID | [optional] [readonly]
**Timestamp** | Pointer to **int32** | The time the message was left as a Unix time value | [optional] [readonly]
**Folder** | Pointer to [**VoicemailFolderBase**](VoicemailFolderBase.md) |  | [optional]

## Methods

### NewVoicemailMessage

`func NewVoicemailMessage() *VoicemailMessage`

NewVoicemailMessage instantiates a new VoicemailMessage object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewVoicemailMessageWithDefaults

`func NewVoicemailMessageWithDefaults() *VoicemailMessage`

NewVoicemailMessageWithDefaults instantiates a new VoicemailMessage object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdName

`func (o *VoicemailMessage) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *VoicemailMessage) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *VoicemailMessage) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *VoicemailMessage) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *VoicemailMessage) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *VoicemailMessage) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *VoicemailMessage) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *VoicemailMessage) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetDuration

`func (o *VoicemailMessage) GetDuration() int32`

GetDuration returns the Duration field if non-nil, zero value otherwise.

### GetDurationOk

`func (o *VoicemailMessage) GetDurationOk() (*int32, bool)`

GetDurationOk returns a tuple with the Duration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDuration

`func (o *VoicemailMessage) SetDuration(v int32)`

SetDuration sets Duration field to given value.

### HasDuration

`func (o *VoicemailMessage) HasDuration() bool`

HasDuration returns a boolean if a field has been set.

### GetId

`func (o *VoicemailMessage) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *VoicemailMessage) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *VoicemailMessage) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *VoicemailMessage) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTimestamp

`func (o *VoicemailMessage) GetTimestamp() int32`

GetTimestamp returns the Timestamp field if non-nil, zero value otherwise.

### GetTimestampOk

`func (o *VoicemailMessage) GetTimestampOk() (*int32, bool)`

GetTimestampOk returns a tuple with the Timestamp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimestamp

`func (o *VoicemailMessage) SetTimestamp(v int32)`

SetTimestamp sets Timestamp field to given value.

### HasTimestamp

`func (o *VoicemailMessage) HasTimestamp() bool`

HasTimestamp returns a boolean if a field has been set.

### GetFolder

`func (o *VoicemailMessage) GetFolder() VoicemailFolderBase`

GetFolder returns the Folder field if non-nil, zero value otherwise.

### GetFolderOk

`func (o *VoicemailMessage) GetFolderOk() (*VoicemailFolderBase, bool)`

GetFolderOk returns a tuple with the Folder field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFolder

`func (o *VoicemailMessage) SetFolder(v VoicemailFolderBase)`

SetFolder sets Folder field to given value.

### HasFolder

`func (o *VoicemailMessage) HasFolder() bool`

HasFolder returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
