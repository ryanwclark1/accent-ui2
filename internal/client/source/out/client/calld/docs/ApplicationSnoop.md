# ApplicationSnoop

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**SnoopedCallId** | Pointer to **string** |  | [optional] [readonly]
**SnoopingCallId** | **string** |  |
**Uuid** | Pointer to **string** |  | [optional] [readonly]
**WhisperMode** | **string** |  |

## Methods

### NewApplicationSnoop

`func NewApplicationSnoop(snoopingCallId string, whisperMode string, ) *ApplicationSnoop`

NewApplicationSnoop instantiates a new ApplicationSnoop object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationSnoopWithDefaults

`func NewApplicationSnoopWithDefaults() *ApplicationSnoop`

NewApplicationSnoopWithDefaults instantiates a new ApplicationSnoop object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSnoopedCallId

`func (o *ApplicationSnoop) GetSnoopedCallId() string`

GetSnoopedCallId returns the SnoopedCallId field if non-nil, zero value otherwise.

### GetSnoopedCallIdOk

`func (o *ApplicationSnoop) GetSnoopedCallIdOk() (*string, bool)`

GetSnoopedCallIdOk returns a tuple with the SnoopedCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopedCallId

`func (o *ApplicationSnoop) SetSnoopedCallId(v string)`

SetSnoopedCallId sets SnoopedCallId field to given value.

### HasSnoopedCallId

`func (o *ApplicationSnoop) HasSnoopedCallId() bool`

HasSnoopedCallId returns a boolean if a field has been set.

### GetSnoopingCallId

`func (o *ApplicationSnoop) GetSnoopingCallId() string`

GetSnoopingCallId returns the SnoopingCallId field if non-nil, zero value otherwise.

### GetSnoopingCallIdOk

`func (o *ApplicationSnoop) GetSnoopingCallIdOk() (*string, bool)`

GetSnoopingCallIdOk returns a tuple with the SnoopingCallId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSnoopingCallId

`func (o *ApplicationSnoop) SetSnoopingCallId(v string)`

SetSnoopingCallId sets SnoopingCallId field to given value.

### GetUuid

`func (o *ApplicationSnoop) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *ApplicationSnoop) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *ApplicationSnoop) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *ApplicationSnoop) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetWhisperMode

`func (o *ApplicationSnoop) GetWhisperMode() string`

GetWhisperMode returns the WhisperMode field if non-nil, zero value otherwise.

### GetWhisperModeOk

`func (o *ApplicationSnoop) GetWhisperModeOk() (*string, bool)`

GetWhisperModeOk returns a tuple with the WhisperMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWhisperMode

`func (o *ApplicationSnoop) SetWhisperMode(v string)`

SetWhisperMode sets WhisperMode field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
