# LineCreate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerIdName** | Pointer to **string** | Name to display when calling | [optional]
**CallerIdNum** | Pointer to **string** | Number to display when calling | [optional]
**Context** | **string** | The name of an internal context |
**Id** | Pointer to **int32** | Line ID | [optional] [readonly]
**Position** | Pointer to **int32** | Line&#39;s position on the device | [optional]
**ProvisioningCode** | Pointer to **string** | Code used to provision a device | [optional]
**Registrar** | Pointer to **string** | Name of the template line used by the device | [optional]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional]
**EndpointSip** | Pointer to [**EndpointSIP**](EndpointSIP.md) |  | [optional]
**EndpointSccp** | Pointer to [**EndpointSccp**](EndpointSccp.md) |  | [optional]
**EndpointCustom** | Pointer to [**EndpointCustom**](EndpointCustom.md) |  | [optional]

## Methods

### NewLineCreate

`func NewLineCreate(context string, ) *LineCreate`

NewLineCreate instantiates a new LineCreate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineCreateWithDefaults

`func NewLineCreateWithDefaults() *LineCreate`

NewLineCreateWithDefaults instantiates a new LineCreate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdName

`func (o *LineCreate) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *LineCreate) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *LineCreate) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *LineCreate) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *LineCreate) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *LineCreate) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *LineCreate) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *LineCreate) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetContext

`func (o *LineCreate) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LineCreate) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LineCreate) SetContext(v string)`

SetContext sets Context field to given value.

### GetId

`func (o *LineCreate) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineCreate) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineCreate) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LineCreate) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *LineCreate) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LineCreate) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LineCreate) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LineCreate) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProvisioningCode

`func (o *LineCreate) GetProvisioningCode() string`

GetProvisioningCode returns the ProvisioningCode field if non-nil, zero value otherwise.

### GetProvisioningCodeOk

`func (o *LineCreate) GetProvisioningCodeOk() (*string, bool)`

GetProvisioningCodeOk returns a tuple with the ProvisioningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningCode

`func (o *LineCreate) SetProvisioningCode(v string)`

SetProvisioningCode sets ProvisioningCode field to given value.

### HasProvisioningCode

`func (o *LineCreate) HasProvisioningCode() bool`

HasProvisioningCode returns a boolean if a field has been set.

### GetRegistrar

`func (o *LineCreate) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *LineCreate) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *LineCreate) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *LineCreate) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetExtensions

`func (o *LineCreate) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *LineCreate) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *LineCreate) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *LineCreate) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetEndpointSip

`func (o *LineCreate) GetEndpointSip() EndpointSIP`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *LineCreate) GetEndpointSipOk() (*EndpointSIP, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *LineCreate) SetEndpointSip(v EndpointSIP)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *LineCreate) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

### GetEndpointSccp

`func (o *LineCreate) GetEndpointSccp() EndpointSccp`

GetEndpointSccp returns the EndpointSccp field if non-nil, zero value otherwise.

### GetEndpointSccpOk

`func (o *LineCreate) GetEndpointSccpOk() (*EndpointSccp, bool)`

GetEndpointSccpOk returns a tuple with the EndpointSccp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSccp

`func (o *LineCreate) SetEndpointSccp(v EndpointSccp)`

SetEndpointSccp sets EndpointSccp field to given value.

### HasEndpointSccp

`func (o *LineCreate) HasEndpointSccp() bool`

HasEndpointSccp returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *LineCreate) GetEndpointCustom() EndpointCustom`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *LineCreate) GetEndpointCustomOk() (*EndpointCustom, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *LineCreate) SetEndpointCustom(v EndpointCustom)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *LineCreate) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
