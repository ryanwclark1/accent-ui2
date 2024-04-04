# UserPostRelationLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**DeviceId** | Pointer to **string** | ID of the device associated to the line | [optional]
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

### NewUserPostRelationLine

`func NewUserPostRelationLine(context string, ) *UserPostRelationLine`

NewUserPostRelationLine instantiates a new UserPostRelationLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostRelationLineWithDefaults

`func NewUserPostRelationLineWithDefaults() *UserPostRelationLine`

NewUserPostRelationLineWithDefaults instantiates a new UserPostRelationLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDeviceId

`func (o *UserPostRelationLine) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *UserPostRelationLine) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *UserPostRelationLine) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *UserPostRelationLine) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetCallerIdName

`func (o *UserPostRelationLine) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *UserPostRelationLine) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *UserPostRelationLine) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *UserPostRelationLine) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *UserPostRelationLine) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *UserPostRelationLine) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *UserPostRelationLine) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *UserPostRelationLine) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetContext

`func (o *UserPostRelationLine) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *UserPostRelationLine) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *UserPostRelationLine) SetContext(v string)`

SetContext sets Context field to given value.

### GetId

`func (o *UserPostRelationLine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPostRelationLine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPostRelationLine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPostRelationLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *UserPostRelationLine) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *UserPostRelationLine) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *UserPostRelationLine) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *UserPostRelationLine) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProvisioningCode

`func (o *UserPostRelationLine) GetProvisioningCode() string`

GetProvisioningCode returns the ProvisioningCode field if non-nil, zero value otherwise.

### GetProvisioningCodeOk

`func (o *UserPostRelationLine) GetProvisioningCodeOk() (*string, bool)`

GetProvisioningCodeOk returns a tuple with the ProvisioningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningCode

`func (o *UserPostRelationLine) SetProvisioningCode(v string)`

SetProvisioningCode sets ProvisioningCode field to given value.

### HasProvisioningCode

`func (o *UserPostRelationLine) HasProvisioningCode() bool`

HasProvisioningCode returns a boolean if a field has been set.

### GetRegistrar

`func (o *UserPostRelationLine) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *UserPostRelationLine) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *UserPostRelationLine) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *UserPostRelationLine) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetExtensions

`func (o *UserPostRelationLine) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UserPostRelationLine) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UserPostRelationLine) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UserPostRelationLine) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetEndpointSip

`func (o *UserPostRelationLine) GetEndpointSip() EndpointSIP`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *UserPostRelationLine) GetEndpointSipOk() (*EndpointSIP, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *UserPostRelationLine) SetEndpointSip(v EndpointSIP)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *UserPostRelationLine) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

### GetEndpointSccp

`func (o *UserPostRelationLine) GetEndpointSccp() EndpointSccp`

GetEndpointSccp returns the EndpointSccp field if non-nil, zero value otherwise.

### GetEndpointSccpOk

`func (o *UserPostRelationLine) GetEndpointSccpOk() (*EndpointSccp, bool)`

GetEndpointSccpOk returns a tuple with the EndpointSccp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSccp

`func (o *UserPostRelationLine) SetEndpointSccp(v EndpointSccp)`

SetEndpointSccp sets EndpointSccp field to given value.

### HasEndpointSccp

`func (o *UserPostRelationLine) HasEndpointSccp() bool`

HasEndpointSccp returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *UserPostRelationLine) GetEndpointCustom() EndpointCustom`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *UserPostRelationLine) GetEndpointCustomOk() (*EndpointCustom, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *UserPostRelationLine) SetEndpointCustom(v EndpointCustom)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *UserPostRelationLine) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
