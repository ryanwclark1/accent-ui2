# Relocate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Completions** | Pointer to [**[]RelocateCompletion**](RelocateCompletion.md) |  | [optional]
**Initiator** | Pointer to **string** | The user UUID of the relocate initiator | [optional]
**InitiatorCall** | Pointer to **string** | Call ID of the relocate initiator | [optional]
**RecipientCall** | Pointer to **string** | Call ID of the recipient of the relocate. | [optional]
**RelocatedCall** | Pointer to **string** | Call ID of the call being relocated to someone else | [optional]
**Uuid** | Pointer to **string** | Unique identifier of the relocate | [optional]

## Methods

### NewRelocate

`func NewRelocate() *Relocate`

NewRelocate instantiates a new Relocate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRelocateWithDefaults

`func NewRelocateWithDefaults() *Relocate`

NewRelocateWithDefaults instantiates a new Relocate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCompletions

`func (o *Relocate) GetCompletions() []RelocateCompletion`

GetCompletions returns the Completions field if non-nil, zero value otherwise.

### GetCompletionsOk

`func (o *Relocate) GetCompletionsOk() (*[]RelocateCompletion, bool)`

GetCompletionsOk returns a tuple with the Completions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompletions

`func (o *Relocate) SetCompletions(v []RelocateCompletion)`

SetCompletions sets Completions field to given value.

### HasCompletions

`func (o *Relocate) HasCompletions() bool`

HasCompletions returns a boolean if a field has been set.

### GetInitiator

`func (o *Relocate) GetInitiator() string`

GetInitiator returns the Initiator field if non-nil, zero value otherwise.

### GetInitiatorOk

`func (o *Relocate) GetInitiatorOk() (*string, bool)`

GetInitiatorOk returns a tuple with the Initiator field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiator

`func (o *Relocate) SetInitiator(v string)`

SetInitiator sets Initiator field to given value.

### HasInitiator

`func (o *Relocate) HasInitiator() bool`

HasInitiator returns a boolean if a field has been set.

### GetInitiatorCall

`func (o *Relocate) GetInitiatorCall() string`

GetInitiatorCall returns the InitiatorCall field if non-nil, zero value otherwise.

### GetInitiatorCallOk

`func (o *Relocate) GetInitiatorCallOk() (*string, bool)`

GetInitiatorCallOk returns a tuple with the InitiatorCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInitiatorCall

`func (o *Relocate) SetInitiatorCall(v string)`

SetInitiatorCall sets InitiatorCall field to given value.

### HasInitiatorCall

`func (o *Relocate) HasInitiatorCall() bool`

HasInitiatorCall returns a boolean if a field has been set.

### GetRecipientCall

`func (o *Relocate) GetRecipientCall() string`

GetRecipientCall returns the RecipientCall field if non-nil, zero value otherwise.

### GetRecipientCallOk

`func (o *Relocate) GetRecipientCallOk() (*string, bool)`

GetRecipientCallOk returns a tuple with the RecipientCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipientCall

`func (o *Relocate) SetRecipientCall(v string)`

SetRecipientCall sets RecipientCall field to given value.

### HasRecipientCall

`func (o *Relocate) HasRecipientCall() bool`

HasRecipientCall returns a boolean if a field has been set.

### GetRelocatedCall

`func (o *Relocate) GetRelocatedCall() string`

GetRelocatedCall returns the RelocatedCall field if non-nil, zero value otherwise.

### GetRelocatedCallOk

`func (o *Relocate) GetRelocatedCallOk() (*string, bool)`

GetRelocatedCallOk returns a tuple with the RelocatedCall field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRelocatedCall

`func (o *Relocate) SetRelocatedCall(v string)`

SetRelocatedCall sets RelocatedCall field to given value.

### HasRelocatedCall

`func (o *Relocate) HasRelocatedCall() bool`

HasRelocatedCall returns a boolean if a field has been set.

### GetUuid

`func (o *Relocate) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Relocate) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Relocate) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Relocate) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
