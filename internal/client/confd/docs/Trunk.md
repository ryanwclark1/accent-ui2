# Trunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the trunk | [optional] [readonly]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]
**EndpointCustom** | Pointer to [**EndpointCustomRelationBase**](EndpointCustomRelationBase.md) |  | [optional]
**EndpointIax** | Pointer to [**EndpointIAXRelationBase**](EndpointIAXRelationBase.md) |  | [optional]
**EndpointSip** | Pointer to [**EndpointSipRelationBase**](EndpointSipRelationBase.md) |  | [optional]
**Outcalls** | Pointer to [**[]OutcallRelationBase**](OutcallRelationBase.md) |  | [optional] [readonly]
**Context** | Pointer to **string** | The context of the trunk | [optional]
**TwilioIncoming** | Pointer to **bool** | Use this trunk&#39;s settings to handle incoming calls from Twilio | [optional]

## Methods

### NewTrunk

`func NewTrunk() *Trunk`

NewTrunk instantiates a new Trunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewTrunkWithDefaults

`func NewTrunkWithDefaults() *Trunk`

NewTrunkWithDefaults instantiates a new Trunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Trunk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Trunk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Trunk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Trunk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Trunk) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Trunk) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Trunk) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Trunk) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *Trunk) GetEndpointCustom() EndpointCustomRelationBase`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *Trunk) GetEndpointCustomOk() (*EndpointCustomRelationBase, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *Trunk) SetEndpointCustom(v EndpointCustomRelationBase)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *Trunk) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

### GetEndpointIax

`func (o *Trunk) GetEndpointIax() EndpointIAXRelationBase`

GetEndpointIax returns the EndpointIax field if non-nil, zero value otherwise.

### GetEndpointIaxOk

`func (o *Trunk) GetEndpointIaxOk() (*EndpointIAXRelationBase, bool)`

GetEndpointIaxOk returns a tuple with the EndpointIax field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointIax

`func (o *Trunk) SetEndpointIax(v EndpointIAXRelationBase)`

SetEndpointIax sets EndpointIax field to given value.

### HasEndpointIax

`func (o *Trunk) HasEndpointIax() bool`

HasEndpointIax returns a boolean if a field has been set.

### GetEndpointSip

`func (o *Trunk) GetEndpointSip() EndpointSipRelationBase`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *Trunk) GetEndpointSipOk() (*EndpointSipRelationBase, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *Trunk) SetEndpointSip(v EndpointSipRelationBase)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *Trunk) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

### GetOutcalls

`func (o *Trunk) GetOutcalls() []OutcallRelationBase`

GetOutcalls returns the Outcalls field if non-nil, zero value otherwise.

### GetOutcallsOk

`func (o *Trunk) GetOutcallsOk() (*[]OutcallRelationBase, bool)`

GetOutcallsOk returns a tuple with the Outcalls field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutcalls

`func (o *Trunk) SetOutcalls(v []OutcallRelationBase)`

SetOutcalls sets Outcalls field to given value.

### HasOutcalls

`func (o *Trunk) HasOutcalls() bool`

HasOutcalls returns a boolean if a field has been set.

### GetContext

`func (o *Trunk) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Trunk) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Trunk) SetContext(v string)`

SetContext sets Context field to given value.

### HasContext

`func (o *Trunk) HasContext() bool`

HasContext returns a boolean if a field has been set.

### GetTwilioIncoming

`func (o *Trunk) GetTwilioIncoming() bool`

GetTwilioIncoming returns the TwilioIncoming field if non-nil, zero value otherwise.

### GetTwilioIncomingOk

`func (o *Trunk) GetTwilioIncomingOk() (*bool, bool)`

GetTwilioIncomingOk returns a tuple with the TwilioIncoming field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTwilioIncoming

`func (o *Trunk) SetTwilioIncoming(v bool)`

SetTwilioIncoming sets TwilioIncoming field to given value.

### HasTwilioIncoming

`func (o *Trunk) HasTwilioIncoming() bool`

HasTwilioIncoming returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
