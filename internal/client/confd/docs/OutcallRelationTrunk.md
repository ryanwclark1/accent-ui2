# OutcallRelationTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the trunk | [optional] [readonly]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**EndpointCustom** | Pointer to [**EndpointCustomRelationBase**](EndpointCustomRelationBase.md) |  | [optional]
**EndpointIax** | Pointer to [**EndpointIAXRelationBase**](EndpointIAXRelationBase.md) |  | [optional]
**EndpointSip** | Pointer to [**EndpointSipRelationBase**](EndpointSipRelationBase.md) |  | [optional]

## Methods

### NewOutcallRelationTrunk

`func NewOutcallRelationTrunk() *OutcallRelationTrunk`

NewOutcallRelationTrunk instantiates a new OutcallRelationTrunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOutcallRelationTrunkWithDefaults

`func NewOutcallRelationTrunkWithDefaults() *OutcallRelationTrunk`

NewOutcallRelationTrunkWithDefaults instantiates a new OutcallRelationTrunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *OutcallRelationTrunk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *OutcallRelationTrunk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *OutcallRelationTrunk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *OutcallRelationTrunk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantUuid

`func (o *OutcallRelationTrunk) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *OutcallRelationTrunk) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *OutcallRelationTrunk) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *OutcallRelationTrunk) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *OutcallRelationTrunk) GetEndpointCustom() EndpointCustomRelationBase`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *OutcallRelationTrunk) GetEndpointCustomOk() (*EndpointCustomRelationBase, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *OutcallRelationTrunk) SetEndpointCustom(v EndpointCustomRelationBase)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *OutcallRelationTrunk) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

### GetEndpointIax

`func (o *OutcallRelationTrunk) GetEndpointIax() EndpointIAXRelationBase`

GetEndpointIax returns the EndpointIax field if non-nil, zero value otherwise.

### GetEndpointIaxOk

`func (o *OutcallRelationTrunk) GetEndpointIaxOk() (*EndpointIAXRelationBase, bool)`

GetEndpointIaxOk returns a tuple with the EndpointIax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIax

`func (o *OutcallRelationTrunk) SetEndpointIax(v EndpointIAXRelationBase)`

SetEndpointIax sets EndpointIax field to given value.

### HasEndpointIax

`func (o *OutcallRelationTrunk) HasEndpointIax() bool`

HasEndpointIax returns a boolean if a field has been set.

### GetEndpointSip

`func (o *OutcallRelationTrunk) GetEndpointSip() EndpointSipRelationBase`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *OutcallRelationTrunk) GetEndpointSipOk() (*EndpointSipRelationBase, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *OutcallRelationTrunk) SetEndpointSip(v EndpointSipRelationBase)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *OutcallRelationTrunk) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
