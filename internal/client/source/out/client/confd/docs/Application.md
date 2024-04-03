# Application

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | Pointer to **string** | Application name | [optional]
**Uuid** | Pointer to **string** | Application UUID | [optional] [readonly]
**Lines** | Pointer to [**[]LineRelationBase**](LineRelationBase.md) |  | [optional] [readonly]
**Destination** | Pointer to **string** | Destination where the call entering in the application will be sent | [optional]
**DestinationOptions** | Pointer to [**ApplicationDestinationNode**](ApplicationDestinationNode.md) |  | [optional]
**TenantUuid** | Pointer to **string** | The UUID of the tenant | [optional] [readonly]

## Methods

### NewApplication

`func NewApplication() *Application`

NewApplication instantiates a new Application object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationWithDefaults

`func NewApplicationWithDefaults() *Application`

NewApplicationWithDefaults instantiates a new Application object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Application) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Application) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Application) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Application) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *Application) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Application) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Application) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Application) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetLines

`func (o *Application) GetLines() []LineRelationBase`

GetLines returns the Lines field if non-nil, zero value otherwise.

### GetLinesOk

`func (o *Application) GetLinesOk() (*[]LineRelationBase, bool)`

GetLinesOk returns a tuple with the Lines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLines

`func (o *Application) SetLines(v []LineRelationBase)`

SetLines sets Lines field to given value.

### HasLines

`func (o *Application) HasLines() bool`

HasLines returns a boolean if a field has been set.

### GetDestination

`func (o *Application) GetDestination() string`

GetDestination returns the Destination field if non-nil, zero value otherwise.

### GetDestinationOk

`func (o *Application) GetDestinationOk() (*string, bool)`

GetDestinationOk returns a tuple with the Destination field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestination

`func (o *Application) SetDestination(v string)`

SetDestination sets Destination field to given value.

### HasDestination

`func (o *Application) HasDestination() bool`

HasDestination returns a boolean if a field has been set.

### GetDestinationOptions

`func (o *Application) GetDestinationOptions() ApplicationDestinationNode`

GetDestinationOptions returns the DestinationOptions field if non-nil, zero value otherwise.

### GetDestinationOptionsOk

`func (o *Application) GetDestinationOptionsOk() (*ApplicationDestinationNode, bool)`

GetDestinationOptionsOk returns a tuple with the DestinationOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDestinationOptions

`func (o *Application) SetDestinationOptions(v ApplicationDestinationNode)`

SetDestinationOptions sets DestinationOptions field to given value.

### HasDestinationOptions

`func (o *Application) HasDestinationOptions() bool`

HasDestinationOptions returns a boolean if a field has been set.

### GetTenantUuid

`func (o *Application) GetTenantUuid() string`

GetTenantUuid returns the TenantUuid field if non-nil, zero value otherwise.

### GetTenantUuidOk

`func (o *Application) GetTenantUuidOk() (*string, bool)`

GetTenantUuidOk returns a tuple with the TenantUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTenantUuid

`func (o *Application) SetTenantUuid(v string)`

SetTenantUuid sets TenantUuid field to given value.

### HasTenantUuid

`func (o *Application) HasTenantUuid() bool`

HasTenantUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
