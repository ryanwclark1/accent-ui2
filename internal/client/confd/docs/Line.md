# Line

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

## Methods

### NewLine

`func NewLine(context string, ) *Line`

NewLine instantiates a new Line object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLineWithDefaults

`func NewLineWithDefaults() *Line`

NewLineWithDefaults instantiates a new Line object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdName

`func (o *Line) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *Line) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *Line) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *Line) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetCallerIdNum

`func (o *Line) GetCallerIdNum() string`

GetCallerIdNum returns the CallerIdNum field if non-nil, zero value otherwise.

### GetCallerIdNumOk

`func (o *Line) GetCallerIdNumOk() (*string, bool)`

GetCallerIdNumOk returns a tuple with the CallerIdNum field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdNum

`func (o *Line) SetCallerIdNum(v string)`

SetCallerIdNum sets CallerIdNum field to given value.

### HasCallerIdNum

`func (o *Line) HasCallerIdNum() bool`

HasCallerIdNum returns a boolean if a field has been set.

### GetContext

`func (o *Line) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *Line) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *Line) SetContext(v string)`

SetContext sets Context field to given value.

### GetId

`func (o *Line) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Line) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Line) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *Line) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPosition

`func (o *Line) GetPosition() int32`

GetPosition returns the Position field if non-nil, zero value otherwise.

### GetPositionOk

`func (o *Line) GetPositionOk() (*int32, bool)`

GetPositionOk returns a tuple with the Position field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPosition

`func (o *Line) SetPosition(v int32)`

SetPosition sets Position field to given value.

### HasPosition

`func (o *Line) HasPosition() bool`

HasPosition returns a boolean if a field has been set.

### GetProvisioningCode

`func (o *Line) GetProvisioningCode() string`

GetProvisioningCode returns the ProvisioningCode field if non-nil, zero value otherwise.

### GetProvisioningCodeOk

`func (o *Line) GetProvisioningCodeOk() (*string, bool)`

GetProvisioningCodeOk returns a tuple with the ProvisioningCode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioningCode

`func (o *Line) SetProvisioningCode(v string)`

SetProvisioningCode sets ProvisioningCode field to given value.

### HasProvisioningCode

`func (o *Line) HasProvisioningCode() bool`

HasProvisioningCode returns a boolean if a field has been set.

### GetRegistrar

`func (o *Line) GetRegistrar() string`

GetRegistrar returns the Registrar field if non-nil, zero value otherwise.

### GetRegistrarOk

`func (o *Line) GetRegistrarOk() (*string, bool)`

GetRegistrarOk returns a tuple with the Registrar field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistrar

`func (o *Line) SetRegistrar(v string)`

SetRegistrar sets Registrar field to given value.

### HasRegistrar

`func (o *Line) HasRegistrar() bool`

HasRegistrar returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
