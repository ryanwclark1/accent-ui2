# CallPickupTargets

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to [**CallPickupTargetUsers**](CallPickupTargetUsers.md) |  | [optional]

## Methods

### NewCallPickupTargets

`func NewCallPickupTargets() *CallPickupTargets`

NewCallPickupTargets instantiates a new CallPickupTargets object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPickupTargetsWithDefaults

`func NewCallPickupTargetsWithDefaults() *CallPickupTargets`

NewCallPickupTargetsWithDefaults instantiates a new CallPickupTargets object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *CallPickupTargets) GetRecipients() CallPickupTargetUsers`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CallPickupTargets) GetRecipientsOk() (*CallPickupTargetUsers, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CallPickupTargets) SetRecipients(v CallPickupTargetUsers)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CallPickupTargets) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
