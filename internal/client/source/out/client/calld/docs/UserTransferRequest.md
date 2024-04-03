# UserTransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Exten** | **string** | Extension of the recipient of the transfer |
**Flow** | Pointer to [**TransferFlow**](TransferFlow.md) |  | [optional] [default to TRANSFERFLOW_ATTENDED]
**InitiatorCall** | **string** | Call ID of the transfer initiator. This call must be owned by the authenticated user. |
**Timeout** | Pointer to **int32** | Maximum ringing time before cancelling the transfer (in seconds). Default (or null) is an unlimited ring time. | [optional]

## Methods

### NewUserTransferRequest

`func NewUserTransferRequest(exten string, initiatorCall string, ) *UserTransferRequest`

NewUserTransferRequest instantiates a new UserTransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserTransferRequestWithDefaults

`func NewUserTransferRequestWithDefaults() *UserTransferRequest`

NewUserTransferRequestWithDefaults instantiates a new UserTransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetExten

`func (o *UserTransferRequest) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *UserTransferRequest) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *UserTransferRequest) SetExten(v string)`

SetExten sets Exten field to given value.

### GetFlow

`func (o *UserTransferRequest) GetFlow() TransferFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *UserTransferRequest) GetFlowOk() (*TransferFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *UserTransferRequest) SetFlow(v TransferFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *UserTransferRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetInitiatorCall

`func (o *UserTransferRequest) GetInitiatorCall() string`

GetInitiatorCall returns the InitiatorCall field if non-nil, zero value otherwise.

### GetInitiatorCallOk

`func (o *UserTransferRequest) GetInitiatorCallOk() (*string, bool)`

GetInitiatorCallOk returns a tuple with the InitiatorCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCall

`func (o *UserTransferRequest) SetInitiatorCall(v string)`

SetInitiatorCall sets InitiatorCall field to given value.

### GetTimeout

`func (o *UserTransferRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *UserTransferRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *UserTransferRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *UserTransferRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
