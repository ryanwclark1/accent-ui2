# DestinationHangupCause

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Busy** | Pointer to [**DestinationHangupBusy**](DestinationHangupBusy.md) |  | [optional]
**Congestion** | Pointer to [**DestinationHangupCongestion**](DestinationHangupCongestion.md) |  | [optional]
**Normal** | Pointer to [**DestinationHangupNormal**](DestinationHangupNormal.md) |  | [optional]

## Methods

### NewDestinationHangupCause

`func NewDestinationHangupCause() *DestinationHangupCause`

NewDestinationHangupCause instantiates a new DestinationHangupCause object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationHangupCauseWithDefaults

`func NewDestinationHangupCauseWithDefaults() *DestinationHangupCause`

NewDestinationHangupCauseWithDefaults instantiates a new DestinationHangupCause object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusy

`func (o *DestinationHangupCause) GetBusy() DestinationHangupBusy`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *DestinationHangupCause) GetBusyOk() (*DestinationHangupBusy, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *DestinationHangupCause) SetBusy(v DestinationHangupBusy)`

SetBusy sets Busy field to given value.

### HasBusy

`func (o *DestinationHangupCause) HasBusy() bool`

HasBusy returns a boolean if a field has been set.

### GetCongestion

`func (o *DestinationHangupCause) GetCongestion() DestinationHangupCongestion`

GetCongestion returns the Congestion field if non-nil, zero value otherwise.

### GetCongestionOk

`func (o *DestinationHangupCause) GetCongestionOk() (*DestinationHangupCongestion, bool)`

GetCongestionOk returns a tuple with the Congestion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCongestion

`func (o *DestinationHangupCause) SetCongestion(v DestinationHangupCongestion)`

SetCongestion sets Congestion field to given value.

### HasCongestion

`func (o *DestinationHangupCause) HasCongestion() bool`

HasCongestion returns a boolean if a field has been set.

### GetNormal

`func (o *DestinationHangupCause) GetNormal() DestinationHangupNormal`

GetNormal returns the Normal field if non-nil, zero value otherwise.

### GetNormalOk

`func (o *DestinationHangupCause) GetNormalOk() (*DestinationHangupNormal, bool)`

GetNormalOk returns a tuple with the Normal field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNormal

`func (o *DestinationHangupCause) SetNormal(v DestinationHangupNormal)`

SetNormal sets Normal field to given value.

### HasNormal

`func (o *DestinationHangupCause) HasNormal() bool`

HasNormal returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
