# WizardDiscoverGateway

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Gateway** | Pointer to **string** | Gateway IPv4 address | [optional]
**Interface** | Pointer to **string** | Interface name (e.g. eth0) | [optional]

## Methods

### NewWizardDiscoverGateway

`func NewWizardDiscoverGateway() *WizardDiscoverGateway`

NewWizardDiscoverGateway instantiates a new WizardDiscoverGateway object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewWizardDiscoverGatewayWithDefaults

`func NewWizardDiscoverGatewayWithDefaults() *WizardDiscoverGateway`

NewWizardDiscoverGatewayWithDefaults instantiates a new WizardDiscoverGateway object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGateway

`func (o *WizardDiscoverGateway) GetGateway() string`

GetGateway returns the Gateway field if non-nil, zero value otherwise.

### GetGatewayOk

`func (o *WizardDiscoverGateway) GetGatewayOk() (*string, bool)`

GetGatewayOk returns a tuple with the Gateway field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGateway

`func (o *WizardDiscoverGateway) SetGateway(v string)`

SetGateway sets Gateway field to given value.

### HasGateway

`func (o *WizardDiscoverGateway) HasGateway() bool`

HasGateway returns a boolean if a field has been set.

### GetInterface

`func (o *WizardDiscoverGateway) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *WizardDiscoverGateway) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *WizardDiscoverGateway) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *WizardDiscoverGateway) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
