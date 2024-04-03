# DestinationHangupBusy

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;hangup&#39; |
**Cause** | Pointer to **string** | MUST be &#39;busy&#39; | [optional]
**Timeout** | Pointer to **int32** | The timeout of the hangup | [optional]

## Methods

### NewDestinationHangupBusy

`func NewDestinationHangupBusy(type_ string, ) *DestinationHangupBusy`

NewDestinationHangupBusy instantiates a new DestinationHangupBusy object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationHangupBusyWithDefaults

`func NewDestinationHangupBusyWithDefaults() *DestinationHangupBusy`

NewDestinationHangupBusyWithDefaults instantiates a new DestinationHangupBusy object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationHangupBusy) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationHangupBusy) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationHangupBusy) SetType(v string)`

SetType sets Type field to given value.

### GetCause

`func (o *DestinationHangupBusy) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *DestinationHangupBusy) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *DestinationHangupBusy) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *DestinationHangupBusy) HasCause() bool`

HasCause returns a boolean if a field has been set.

### GetTimeout

`func (o *DestinationHangupBusy) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *DestinationHangupBusy) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *DestinationHangupBusy) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *DestinationHangupBusy) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
