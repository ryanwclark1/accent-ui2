# EndpointCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Custom Endpoint ID | [optional] [readonly]
**Interface** | Pointer to **string** | Asterisk interface to use. (e.g. &#39;dahdi/i1&#39;) | [optional]
**Trunk** | Pointer to [**TrunkRelationBase**](TrunkRelationBase.md) |  | [optional]
**Line** | Pointer to [**[]LineRelationBase**](LineRelationBase.md) |  | [optional] [readonly]
**Enabled** | Pointer to **bool** | Endpoint can be used when it is enabled | [optional] [default to true]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewEndpointCustom

`func NewEndpointCustom() *EndpointCustom`

NewEndpointCustom instantiates a new EndpointCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointCustomWithDefaults

`func NewEndpointCustomWithDefaults() *EndpointCustom`

NewEndpointCustomWithDefaults instantiates a new EndpointCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointCustom) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointCustom) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointCustom) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointCustom) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterface

`func (o *EndpointCustom) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *EndpointCustom) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *EndpointCustom) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *EndpointCustom) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

### GetTrunk

`func (o *EndpointCustom) GetTrunk() TrunkRelationBase`

GetTrunk returns the Trunk field if non-nil, zero value otherwise.

### GetTrunkOk

`func (o *EndpointCustom) GetTrunkOk() (*TrunkRelationBase, bool)`

GetTrunkOk returns a tuple with the Trunk field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTrunk

`func (o *EndpointCustom) SetTrunk(v TrunkRelationBase)`

SetTrunk sets Trunk field to given value.

### HasTrunk

`func (o *EndpointCustom) HasTrunk() bool`

HasTrunk returns a boolean if a field has been set.

### GetLine

`func (o *EndpointCustom) GetLine() []LineRelationBase`

GetLine returns the Line field if non-nil, zero value otherwise.

### GetLineOk

`func (o *EndpointCustom) GetLineOk() (*[]LineRelationBase, bool)`

GetLineOk returns a tuple with the Line field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLine

`func (o *EndpointCustom) SetLine(v []LineRelationBase)`

SetLine sets Line field to given value.

### HasLine

`func (o *EndpointCustom) HasLine() bool`

HasLine returns a boolean if a field has been set.

### GetEnabled

`func (o *EndpointCustom) GetEnabled() bool`

GetEnabled returns the Enabled field if non-nil, zero value otherwise.

### GetEnabledOk

`func (o *EndpointCustom) GetEnabledOk() (*bool, bool)`

GetEnabledOk returns a tuple with the Enabled field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEnabled

`func (o *EndpointCustom) SetEnabled(v bool)`

SetEnabled sets Enabled field to given value.

### HasEnabled

`func (o *EndpointCustom) HasEnabled() bool`

HasEnabled returns a boolean if a field has been set.

### GetTenantUuid

`func (o *EndpointCustom) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *EndpointCustom) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *EndpointCustom) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *EndpointCustom) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
