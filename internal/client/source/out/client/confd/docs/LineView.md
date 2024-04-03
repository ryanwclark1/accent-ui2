# LineView

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Line ID | [optional] [readonly]
**Name** | Pointer to **string** | The name of the line | [optional] [readonly]
**Application** | Pointer to [**ApplicationRelationBase**](ApplicationRelationBase.md) |  | [optional]
**EndpointCustom** | Pointer to [**EndpointCustomRelationBase**](EndpointCustomRelationBase.md) |  | [optional]
**EndpointSccp** | Pointer to [**EndpointSccpRelationBase**](EndpointSccpRelationBase.md) |  | [optional]
**EndpointSip** | Pointer to [**EndpointSipRelationBase**](EndpointSipRelationBase.md) |  | [optional]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]
**Users** | Pointer to [**[]UserRelationBase**](UserRelationBase.md) |  | [optional] [readonly]
**CallerIdName** | Pointer to **string** | Name to display when calling | [optional]
**CallerIdNum** | Pointer to **string** | Number to display when calling | [optional]
**Context** | **string** | The name of an internal context |
**DeviceId** | Pointer to **string** | ID of the device associated to the line | [optional] [readonly]
**DeviceSlot** | Pointer to **int32** | *Deprecated* Please use &#x60;position&#x60; | [optional] [readonly]
**Position** | Pointer to **int32** | Line&#39;s position on the device | [optional]
**Protocol** | Pointer to **string** | Line&#39;s protocol | [optional] [readonly]
**ProvisioningCode** | Pointer to **string** | Code used to provision a device | [optional]
**ProvisioningExtension** | Pointer to **string** | *Deprecated* Please use &#x60;provisioning_code&#x60; | [optional] [readonly]
**Registrar** | Pointer to **string** | Name of the template line used by the device | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewLineView

`func NewLineView(context string, ) *LineView`

NewLineView instantiates a new LineView object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineViewWithDefaults

`func NewLineViewWithDefaults() *LineView`

NewLineViewWithDefaults instantiates a new LineView object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *LineView) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *LineView) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *LineView) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *LineView) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *LineView) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LineView) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LineView) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *LineView) HasName() bool`

HasName returns a boolean if a field has been set.

### GetApplication

`func (o *LineView) GetApplication() ApplicationRelationBase`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *LineView) GetApplicationOk() (*ApplicationRelationBase, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *LineView) SetApplication(v ApplicationRelationBase)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *LineView) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetEndpointCustom

`func (o *LineView) GetEndpointCustom() EndpointCustomRelationBase`

GetEndpointCustom returns the EndpointCustom field if non-nil, zero value otherwise.

### GetEndpointCustomOk

`func (o *LineView) GetEndpointCustomOk() (*EndpointCustomRelationBase, bool)`

GetEndpointCustomOk returns a tuple with the EndpointCustom field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointCustom

`func (o *LineView) SetEndpointCustom(v EndpointCustomRelationBase)`

SetEndpointCustom sets EndpointCustom field to given value.

### HasEndpointCustom

`func (o *LineView) HasEndpointCustom() bool`

HasEndpointCustom returns a boolean if a field has been set.

### GetEndpointSccp

`func (o *LineView) GetEndpointSccp() EndpointSccpRelationBase`

GetEndpointSccp returns the EndpointSccp field if non-nil, zero value otherwise.

### GetEndpointSccpOk

`func (o *LineView) GetEndpointSccpOk() (*EndpointSccpRelationBase, bool)`

GetEndpointSccpOk returns a tuple with the EndpointSccp field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSccp

`func (o *LineView) SetEndpointSccp(v EndpointSccpRelationBase)`

SetEndpointSccp sets EndpointSccp field to given value.

### HasEndpointSccp

`func (o *LineView) HasEndpointSccp() bool`

HasEndpointSccp returns a boolean if a field has been set.

### GetEndpointSip

`func (o *LineView) GetEndpointSip() EndpointSipRelationBase`

GetEndpointSip returns the EndpointSip field if non-nil, zero value otherwise.

### GetEndpointSipOk

`func (o *LineView) GetEndpointSipOk() (*EndpointSipRelationBase, bool)`

GetEndpointSipOk returns a tuple with the EndpointSip field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpointSip

`func (o *LineView) SetEndpointSip(v EndpointSipRelationBase)`

SetEndpointSip sets EndpointSip field to given value.

### HasEndpointSip

`func (o *LineView) HasEndpointSip() bool`

HasEndpointSip returns a boolean if a field has been set.

### GetExtensions

`func (o *LineView) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *LineView) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *LineView) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *LineView) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetUsers

`func (o *LineView) GetUsers() []UserRelationBase`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *LineView) GetUsersOk() (*[]UserRelationBase, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *LineView) SetUsers(v []UserRelationBase)`

SetUsers sets Users field to given value.

### HasUsers

`func (o *LineView) HasUsers() bool`

HasUsers returns a boolean if a field has been set.

### GetCallerIdName

`func (o *LineView) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *LineView) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *LineView) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *LineView) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *LineView) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *LineView) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *LineView) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *LineView) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetContext

`func (o *LineView) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *LineView) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *LineView) SetContext(v string)`

SetContext sets Context field to given value.

### GetDeviceId

`func (o *LineView) GetDeviceId() string`

GetDeviceId returns the DeviceId field if non-nil, zero value otherwise.

### GetDeviceIdOk

`func (o *LineView) GetDeviceIdOk() (*string, bool)`

GetDeviceIdOk returns a tuple with the DeviceId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceId

`func (o *LineView) SetDeviceId(v string)`

SetDeviceId sets DeviceId field to given value.

### HasDeviceId

`func (o *LineView) HasDeviceId() bool`

HasDeviceId returns a boolean if a field has been set.

### GetDeviceSlot

`func (o *LineView) GetDeviceSlot() int32`

GetDeviceSlot returns the DeviceSlot field if non-nil, zero value otherwise.

### GetDeviceSlotOk

`func (o *LineView) GetDeviceSlotOk() (*int32, bool)`

GetDeviceSlotOk returns a tuple with the DeviceSlot field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeviceSlot

`func (o *LineView) SetDeviceSlot(v int32)`

SetDeviceSlot sets DeviceSlot field to given value.

### HasDeviceSlot

`func (o *LineView) HasDeviceSlot() bool`

HasDeviceSlot returns a boolean if a field has been set.

### GetPosition

`func (o *LineView) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *LineView) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *LineView) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *LineView) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProtocol

`func (o *LineView) GetProtocol() string`

GetProtocol returns the Protocol field if non-nil, zero value otherwise.

### GetProtocolOk

`func (o *LineView) GetProtocolOk() (*string, bool)`

GetProtocolOk returns a tuple with the Protocol field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProtocol

`func (o *LineView) SetProtocol(v string)`

SetProtocol sets Protocol field to given value.

### HasProtocol

`func (o *LineView) HasProtocol() bool`

HasProtocol returns a boolean if a field has been set.

### GetProvisioningCode

`func (o *LineView) GetProvisioningCode() string`

GetProvisioningCode returns the ProvisioningCode field if non-nil, zero value otherwise.

### GetProvisioningCodeOk

`func (o *LineView) GetProvisioningCodeOk() (*string, bool)`

GetProvisioningCodeOk returns a tuple with the ProvisioningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningCode

`func (o *LineView) SetProvisioningCode(v string)`

SetProvisioningCode sets ProvisioningCode field to given value.

### HasProvisioningCode

`func (o *LineView) HasProvisioningCode() bool`

HasProvisioningCode returns a boolean if a field has been set.

### GetProvisioningExtension

`func (o *LineView) GetProvisioningExtension() string`

GetProvisioningExtension returns the ProvisioningExtension field if non-nil, zero value otherwise.

### GetProvisioningExtensionOk

`func (o *LineView) GetProvisioningExtensionOk() (*string, bool)`

GetProvisioningExtensionOk returns a tuple with the ProvisioningExtension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningExtension

`func (o *LineView) SetProvisioningExtension(v string)`

SetProvisioningExtension sets ProvisioningExtension field to given value.

### HasProvisioningExtension

`func (o *LineView) HasProvisioningExtension() bool`

HasProvisioningExtension returns a boolean if a field has been set.

### GetRegistrar

`func (o *LineView) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *LineView) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *LineView) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *LineView) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

### GetTenantUuid

`func (o *LineView) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *LineView) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *LineView) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *LineView) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
