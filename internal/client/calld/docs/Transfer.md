# Transfer

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Flow** | Pointer to **string** | The behavior of the transfer | [optional] [default to "attended"]
**Id** | Pointer to **string** | Unique identifier of the transfer | [optional]
**InitiatorCall** | Pointer to **string** | Call ID of the transfer initiator | [optional]
**InitiatorTenantUuid** | Pointer to **string** | Tenant UUID of the user who initiated the transfer | [optional]
**InitiatorUuid** | Pointer to **string** | UUID of the user who initiated the transfer | [optional]
**RecipientCall** | Pointer to **string** | Call ID of the recipient of the transfer. May be null when the transfer is &#39;starting&#39;. | [optional]
**Status** | Pointer to **string** | The current step of the transfer | [optional]
**TransferredCall** | Pointer to **string** | Call ID of the call being transferred to someone else | [optional]

## Methods

### NewTransfer

`func NewTransfer() *Transfer`

NewTransfer instantiates a new Transfer object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTransferWithDefaults

`func NewTransferWithDefaults() *Transfer`

NewTransferWithDefaults instantiates a new Transfer object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFlow

`func (o *Transfer) GetFlow() string`

GetFlow returns the Flow field if non-nil, zero value otherwise.

### GetFlowOk

`func (o *Transfer) GetFlowOk() (*string, bool)`

GetFlowOk returns a tuple with the Flow field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFlow

`func (o *Transfer) SetFlow(v string)`

SetFlow sets Flow field to given value.

### HasFlow

`func (o *Transfer) HasFlow() bool`

HasFlow returns a boolean if a field has been set.

### GetId

`func (o *Transfer) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Transfer) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Transfer) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Transfer) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInitiatorCall

`func (o *Transfer) GetInitiatorCall() string`

GetInitiatorCall returns the InitiatorCall field if non-nil, zero value otherwise.

### GetInitiatorCallOk

`func (o *Transfer) GetInitiatorCallOk() (*string, bool)`

GetInitiatorCallOk returns a tuple with the InitiatorCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCall

`func (o *Transfer) SetInitiatorCall(v string)`

SetInitiatorCall sets InitiatorCall field to given value.

### HasInitiatorCall

`func (o *Transfer) HasInitiatorCall() bool`

HasInitiatorCall returns a boolean if a field has been set.

### GetInitiatorTenantUuid

`func (o *Transfer) GetInitiatorTenantUuid() string`

GetInitiatorTenantUuid returns the InitiatorTenantUuid field if non-nil, zero value otherwise.

### GetInitiatorTenantUuidOk

`func (o *Transfer) GetInitiatorTenantUuidOk() (*string, bool)`

GetInitiatorTenantUuidOk returns a tuple with the InitiatorTenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorTenantUuid

`func (o *Transfer) SetInitiatorTenantUuid(v string)`

SetInitiatorTenantUuid sets InitiatorTenantUuid field to given value.

### HasInitiatorTenantUuid

`func (o *Transfer) HasInitiatorTenantUuid() bool`

HasInitiatorTenantUuid returns a boolean if a field has been set.

### GetInitiatorUuid

`func (o *Transfer) GetInitiatorUuid() string`

GetInitiatorUuid returns the InitiatorUuid field if non-nil, zero value otherwise.

### GetInitiatorUuidOk

`func (o *Transfer) GetInitiatorUuidOk() (*string, bool)`

GetInitiatorUuidOk returns a tuple with the InitiatorUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorUuid

`func (o *Transfer) SetInitiatorUuid(v string)`

SetInitiatorUuid sets InitiatorUuid field to given value.

### HasInitiatorUuid

`func (o *Transfer) HasInitiatorUuid() bool`

HasInitiatorUuid returns a boolean if a field has been set.

### GetRecipientCall

`func (o *Transfer) GetRecipientCall() string`

GetRecipientCall returns the RecipientCall field if non-nil, zero value otherwise.

### GetRecipientCallOk

`func (o *Transfer) GetRecipientCallOk() (*string, bool)`

GetRecipientCallOk returns a tuple with the RecipientCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCall

`func (o *Transfer) SetRecipientCall(v string)`

SetRecipientCall sets RecipientCall field to given value.

### HasRecipientCall

`func (o *Transfer) HasRecipientCall() bool`

HasRecipientCall returns a boolean if a field has been set.

### GetStatus

`func (o *Transfer) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Transfer) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Transfer) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Transfer) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetTransferredCall

`func (o *Transfer) GetTransferredCall() string`

GetTransferredCall returns the TransferredCall field if non-nil, zero value otherwise.

### GetTransferredCallOk

`func (o *Transfer) GetTransferredCallOk() (*string, bool)`

GetTransferredCallOk returns a tuple with the TransferredCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTransferredCall

`func (o *Transfer) SetTransferredCall(v string)`

SetTransferredCall sets TransferredCall field to given value.

### HasTransferredCall

`func (o *Transfer) HasTransferredCall() bool`

HasTransferredCall returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
