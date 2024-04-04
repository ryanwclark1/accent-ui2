# StatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BusConsumer** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**MasterTenant** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**PresenceInitialization** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**RestApi** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]

## Methods

### NewStatusSummary

`func NewStatusSummary() *StatusSummary`

NewStatusSummary instantiates a new StatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSummaryWithDefaults

`func NewStatusSummaryWithDefaults() *StatusSummary`

NewStatusSummaryWithDefaults instantiates a new StatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusConsumer

`func (o *StatusSummary) GetBusConsumer() ComponentWithStatus`

GetBusConsumer returns the BusConsumer field if non-nil, zero value otherwise.

### GetBusConsumerOk

`func (o *StatusSummary) GetBusConsumerOk() (*ComponentWithStatus, bool)`

GetBusConsumerOk returns a tuple with the BusConsumer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusConsumer

`func (o *StatusSummary) SetBusConsumer(v ComponentWithStatus)`

SetBusConsumer sets BusConsumer field to given value.

### HasBusConsumer

`func (o *StatusSummary) HasBusConsumer() bool`

HasBusConsumer returns a boolean if a field has been set.

### GetMasterTenant

`func (o *StatusSummary) GetMasterTenant() ComponentWithStatus`

GetMasterTenant returns the MasterTenant field if non-nil, zero value otherwise.

### GetMasterTenantOk

`func (o *StatusSummary) GetMasterTenantOk() (*ComponentWithStatus, bool)`

GetMasterTenantOk returns a tuple with the MasterTenant field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMasterTenant

`func (o *StatusSummary) SetMasterTenant(v ComponentWithStatus)`

SetMasterTenant sets MasterTenant field to given value.

### HasMasterTenant

`func (o *StatusSummary) HasMasterTenant() bool`

HasMasterTenant returns a boolean if a field has been set.

### GetPresenceInitialization

`func (o *StatusSummary) GetPresenceInitialization() ComponentWithStatus`

GetPresenceInitialization returns the PresenceInitialization field if non-nil, zero value otherwise.

### GetPresenceInitializationOk

`func (o *StatusSummary) GetPresenceInitializationOk() (*ComponentWithStatus, bool)`

GetPresenceInitializationOk returns a tuple with the PresenceInitialization field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPresenceInitialization

`func (o *StatusSummary) SetPresenceInitialization(v ComponentWithStatus)`

SetPresenceInitialization sets PresenceInitialization field to given value.

### HasPresenceInitialization

`func (o *StatusSummary) HasPresenceInitialization() bool`

HasPresenceInitialization returns a boolean if a field has been set.

### GetRestApi

`func (o *StatusSummary) GetRestApi() ComponentWithStatus`

GetRestApi returns the RestApi field if non-nil, zero value otherwise.

### GetRestApiOk

`func (o *StatusSummary) GetRestApiOk() (*ComponentWithStatus, bool)`

GetRestApiOk returns a tuple with the RestApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApi

`func (o *StatusSummary) SetRestApi(v ComponentWithStatus)`

SetRestApi sets RestApi field to given value.

### HasRestApi

`func (o *StatusSummary) HasRestApi() bool`

HasRestApi returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
