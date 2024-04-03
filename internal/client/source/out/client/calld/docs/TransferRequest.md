# TransferRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Context of the recipient of the transfer |
**Exten** | **string** | Extension of the recipient of the transfer |
**Flow** | Pointer to [**TransferFlow**](TransferFlow.md) |  | [optional] [default to TRANSFERFLOW_ATTENDED]
**InitiatorCall** | **string** | Call ID of the transfer initiator |
**Timeout** | Pointer to **int32** | Maximum ringing time before cancelling the transfer (in seconds). Default (or null) is an unlimited ring time. | [optional]
**TransferredCall** | **string** | Call ID of the call being transferred to someone else |
**Variables** | Pointer to **map[string]interface{}** | Channel variables to set on the recipient call | [optional]

## Methods

### NewTransferRequest

`func NewTransferRequest(context string, exten string, initiatorCall string, transferredCall string, ) *TransferRequest`

NewTransferRequest instantiates a new TransferRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferRequestWithDefaults

`func NewTransferRequestWithDefaults() *TransferRequest`

NewTransferRequestWithDefaults instantiates a new TransferRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *TransferRequest) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *TransferRequest) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *TransferRequest) SetContext(v string)`

SetContext sets Context field to given value.

### GetExten

`func (o *TransferRequest) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *TransferRequest) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *TransferRequest) SetExten(v string)`

SetExten sets Exten field to given value.

### GetFlow

`func (o *TransferRequest) GetFlow() TransferFlow`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *TransferRequest) GetFlowOk() (*TransferFlow, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *TransferRequest) SetFlow(v TransferFlow)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *TransferRequest) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetInitiatorCall

`func (o *TransferRequest) GetInitiatorCall() string`

GetInitiatorCall returns the InitiatorCall field if non-nil, zero value otherwise.

### GetInitiatorCallOk

`func (o *TransferRequest) GetInitiatorCallOk() (*string, bool)`

GetInitiatorCallOk returns a tuple with the InitiatorCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCall

`func (o *TransferRequest) SetInitiatorCall(v string)`

SetInitiatorCall sets InitiatorCall field to given value.

### GetTimeout

`func (o *TransferRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *TransferRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *TransferRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *TransferRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

### GetTransferredCall

`func (o *TransferRequest) GetTransferredCall() string`

GetTransferredCall returns the TransferredCall field if non-nil, zero value otherwise.

### GetTransferredCallOk

`func (o *TransferRequest) GetTransferredCallOk() (*string, bool)`

GetTransferredCallOk returns a tuple with the TransferredCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferredCall

`func (o *TransferRequest) SetTransferredCall(v string)`

SetTransferredCall sets TransferredCall field to given value.

### GetVariables

`func (o *TransferRequest) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *TransferRequest) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *TransferRequest) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *TransferRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
