# Registrar

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**BackupHost** | Pointer to **string** | Backup registrar host | [optional]
**BackupPort** | Pointer to **int32** | Backup registrar port | [optional]
**Deletable** | Pointer to **bool** | Define if the registrar can be deleted | [optional]
**Id** | Pointer to **string** | Registrar identifier | [optional]
**MainHost** | **string** | Registrar host |
**MainPort** | Pointer to **int32** | Registrar port | [optional]
**Name** | Pointer to **string** | Display name of the registrar | [optional]
**OutboundProxyHost** | Pointer to **string** | Outbound proxy host | [optional]
**OutboundProxyPort** | Pointer to **int32** | Outbound proxy port | [optional]
**ProxyBackupHost** | Pointer to **string** | Backup proxy host | [optional]
**ProxyBackupPort** | Pointer to **int32** | Backup proxy port | [optional]
**ProxyMainHost** | **string** | Proxy host. Using IPv4 is recommended to have a better integration with some provisioning plugins. (ex: Yealink DND integration) |
**ProxyMainPort** | Pointer to **int32** | Proxy port | [optional]

## Methods

### NewRegistrar

`func NewRegistrar(mainHost string, proxyMainHost string, ) *Registrar`

NewRegistrar instantiates a new Registrar object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRegistrarWithDefaults

`func NewRegistrarWithDefaults() *Registrar`

NewRegistrarWithDefaults instantiates a new Registrar object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBackupHost

`func (o *Registrar) GetBackupHost() string`

GetBackupHost returns the BackupHost field if non-nil, zero value otherwise.

### GetBackupHostOk

`func (o *Registrar) GetBackupHostOk() (*string, bool)`

GetBackupHostOk returns a tuple with the BackupHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupHost

`func (o *Registrar) SetBackupHost(v string)`

SetBackupHost sets BackupHost field to given value.

### HasBackupHost

`func (o *Registrar) HasBackupHost() bool`

HasBackupHost returns a boolean if a field has been set.

### GetBackupPort

`func (o *Registrar) GetBackupPort() int32`

GetBackupPort returns the BackupPort field if non-nil, zero value otherwise.

### GetBackupPortOk

`func (o *Registrar) GetBackupPortOk() (*int32, bool)`

GetBackupPortOk returns a tuple with the BackupPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBackupPort

`func (o *Registrar) SetBackupPort(v int32)`

SetBackupPort sets BackupPort field to given value.

### HasBackupPort

`func (o *Registrar) HasBackupPort() bool`

HasBackupPort returns a boolean if a field has been set.

### GetDeletable

`func (o *Registrar) GetDeletable() bool`

GetDeletable returns the Deletable field if non-nil, zero value otherwise.

### GetDeletableOk

`func (o *Registrar) GetDeletableOk() (*bool, bool)`

GetDeletableOk returns a tuple with the Deletable field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeletable

`func (o *Registrar) SetDeletable(v bool)`

SetDeletable sets Deletable field to given value.

### HasDeletable

`func (o *Registrar) HasDeletable() bool`

HasDeletable returns a boolean if a field has been set.

### GetId

`func (o *Registrar) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Registrar) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Registrar) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *Registrar) HasId() bool`

HasId returns a boolean if a field has been set.

### GetMainHost

`func (o *Registrar) GetMainHost() string`

GetMainHost returns the MainHost field if non-nil, zero value otherwise.

### GetMainHostOk

`func (o *Registrar) GetMainHostOk() (*string, bool)`

GetMainHostOk returns a tuple with the MainHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainHost

`func (o *Registrar) SetMainHost(v string)`

SetMainHost sets MainHost field to given value.

### GetMainPort

`func (o *Registrar) GetMainPort() int32`

GetMainPort returns the MainPort field if non-nil, zero value otherwise.

### GetMainPortOk

`func (o *Registrar) GetMainPortOk() (*int32, bool)`

GetMainPortOk returns a tuple with the MainPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMainPort

`func (o *Registrar) SetMainPort(v int32)`

SetMainPort sets MainPort field to given value.

### HasMainPort

`func (o *Registrar) HasMainPort() bool`

HasMainPort returns a boolean if a field has been set.

### GetName

`func (o *Registrar) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Registrar) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Registrar) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Registrar) HasName() bool`

HasName returns a boolean if a field has been set.

### GetOutboundProxyHost

`func (o *Registrar) GetOutboundProxyHost() string`

GetOutboundProxyHost returns the OutboundProxyHost field if non-nil, zero value otherwise.

### GetOutboundProxyHostOk

`func (o *Registrar) GetOutboundProxyHostOk() (*string, bool)`

GetOutboundProxyHostOk returns a tuple with the OutboundProxyHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProxyHost

`func (o *Registrar) SetOutboundProxyHost(v string)`

SetOutboundProxyHost sets OutboundProxyHost field to given value.

### HasOutboundProxyHost

`func (o *Registrar) HasOutboundProxyHost() bool`

HasOutboundProxyHost returns a boolean if a field has been set.

### GetOutboundProxyPort

`func (o *Registrar) GetOutboundProxyPort() int32`

GetOutboundProxyPort returns the OutboundProxyPort field if non-nil, zero value otherwise.

### GetOutboundProxyPortOk

`func (o *Registrar) GetOutboundProxyPortOk() (*int32, bool)`

GetOutboundProxyPortOk returns a tuple with the OutboundProxyPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOutboundProxyPort

`func (o *Registrar) SetOutboundProxyPort(v int32)`

SetOutboundProxyPort sets OutboundProxyPort field to given value.

### HasOutboundProxyPort

`func (o *Registrar) HasOutboundProxyPort() bool`

HasOutboundProxyPort returns a boolean if a field has been set.

### GetProxyBackupHost

`func (o *Registrar) GetProxyBackupHost() string`

GetProxyBackupHost returns the ProxyBackupHost field if non-nil, zero value otherwise.

### GetProxyBackupHostOk

`func (o *Registrar) GetProxyBackupHostOk() (*string, bool)`

GetProxyBackupHostOk returns a tuple with the ProxyBackupHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyBackupHost

`func (o *Registrar) SetProxyBackupHost(v string)`

SetProxyBackupHost sets ProxyBackupHost field to given value.

### HasProxyBackupHost

`func (o *Registrar) HasProxyBackupHost() bool`

HasProxyBackupHost returns a boolean if a field has been set.

### GetProxyBackupPort

`func (o *Registrar) GetProxyBackupPort() int32`

GetProxyBackupPort returns the ProxyBackupPort field if non-nil, zero value otherwise.

### GetProxyBackupPortOk

`func (o *Registrar) GetProxyBackupPortOk() (*int32, bool)`

GetProxyBackupPortOk returns a tuple with the ProxyBackupPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyBackupPort

`func (o *Registrar) SetProxyBackupPort(v int32)`

SetProxyBackupPort sets ProxyBackupPort field to given value.

### HasProxyBackupPort

`func (o *Registrar) HasProxyBackupPort() bool`

HasProxyBackupPort returns a boolean if a field has been set.

### GetProxyMainHost

`func (o *Registrar) GetProxyMainHost() string`

GetProxyMainHost returns the ProxyMainHost field if non-nil, zero value otherwise.

### GetProxyMainHostOk

`func (o *Registrar) GetProxyMainHostOk() (*string, bool)`

GetProxyMainHostOk returns a tuple with the ProxyMainHost field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyMainHost

`func (o *Registrar) SetProxyMainHost(v string)`

SetProxyMainHost sets ProxyMainHost field to given value.

### GetProxyMainPort

`func (o *Registrar) GetProxyMainPort() int32`

GetProxyMainPort returns the ProxyMainPort field if non-nil, zero value otherwise.

### GetProxyMainPortOk

`func (o *Registrar) GetProxyMainPortOk() (*int32, bool)`

GetProxyMainPortOk returns a tuple with the ProxyMainPort field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProxyMainPort

`func (o *Registrar) SetProxyMainPort(v int32)`

SetProxyMainPort sets ProxyMainPort field to given value.

### HasProxyMainPort

`func (o *Registrar) HasProxyMainPort() bool`

HasProxyMainPort returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
