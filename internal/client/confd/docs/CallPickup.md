# CallPickup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Recipients** | Pointer to [**CallPickupTargetUsers**](CallPickupTargetUsers.md) |  | [optional]
**Surrogates** | Pointer to [**CallPickupInterceptorUsers**](CallPickupInterceptorUsers.md) |  | [optional]
**Description** | Pointer to **string** | Additional information about the call pickup | [optional]
**Enabled** | Pointer to **bool** | Disable or enable the call pickup | [optional] [default to true]
**Id** | Pointer to **int32** | The id of the call pickup | [optional] [readonly]
**Name** | **string** | The name of the call pickup |
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewCallPickup

`func NewCallPickup(name string, ) *CallPickup`

NewCallPickup instantiates a new CallPickup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallPickupWithDefaults

`func NewCallPickupWithDefaults() *CallPickup`

NewCallPickupWithDefaults instantiates a new CallPickup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetRecipients

`func (o *CallPickup) GetRecipients() CallPickupTargetUsers`

GetRecipients returns the Recipients field if non-nil, zero value otherwise.

### GetRecipientsOk

`func (o *CallPickup) GetRecipientsOk() (*CallPickupTargetUsers, bool)`

GetRecipientsOk returns a tuple with the Recipients field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRecipients

`func (o *CallPickup) SetRecipients(v CallPickupTargetUsers)`

SetRecipients sets Recipients field to given value.

### HasRecipients

`func (o *CallPickup) HasRecipients() bool`

HasRecipients returns a boolean if a field has been set.

### GetSurrogates

`func (o *CallPickup) GetSurrogates() CallPickupInterceptorUsers`

GetSurrogates returns the Surrogates field if non-nil, zero value otherwise.

### GetSurrogatesOk

`func (o *CallPickup) GetSurrogatesOk() (*CallPickupInterceptorUsers, bool)`

GetSurrogatesOk returns a tuple with the Surrogates field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSurrogates

`func (o *CallPickup) SetSurrogates(v CallPickupInterceptorUsers)`

SetSurrogates sets Surrogates field to given value.

### HasSurrogates

`func (o *CallPickup) HasSurrogates() bool`

HasSurrogates returns a boolean if a field has been set.

### GetDescription

`func (o *CallPickup) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *CallPickup) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *CallPickup) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *CallPickup) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetEnabled

`func (o *CallPickup) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *CallPickup) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *CallPickup) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *CallPickup) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetId

`func (o *CallPickup) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *CallPickup) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *CallPickup) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *CallPickup) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *CallPickup) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *CallPickup) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *CallPickup) SetName(v string)`

SetName sets Name field to given value.

### GetTenantUuid

`func (o *CallPickup) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *CallPickup) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *CallPickup) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *CallPickup) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
