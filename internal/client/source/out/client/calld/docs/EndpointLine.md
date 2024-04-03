# EndpointLine

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCallCount** | Pointer to **int32** | The number of ongoing calls on that line | [optional]
**Id** | Pointer to **int32** | The ID of the matching confd line | [optional]
**Name** | Pointer to **string** | The name of that given endpoint in Asterisk | [optional]
**Registered** | Pointer to **bool** | Wether or not this trunk is registered. | [optional]
**Technology** | Pointer to **string** | The technology of that endpoint only (SIP, SCCP or custom) | [optional]

## Methods

### NewEndpointLine

`func NewEndpointLine() *EndpointLine`

NewEndpointLine instantiates a new EndpointLine object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointLineWithDefaults

`func NewEndpointLineWithDefaults() *EndpointLine`

NewEndpointLineWithDefaults instantiates a new EndpointLine object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCallCount

`func (o *EndpointLine) GetCurrentCallCount() int32`

GetCurrentCallCount returns the CurrentCallCount field if non-nil, zero value otherwise.

### GetCurrentCallCountOk

`func (o *EndpointLine) GetCurrentCallCountOk() (*int32, bool)`

GetCurrentCallCountOk returns a tuple with the CurrentCallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCallCount

`func (o *EndpointLine) SetCurrentCallCount(v int32)`

SetCurrentCallCount sets CurrentCallCount field to given value.

### HasCurrentCallCount

`func (o *EndpointLine) HasCurrentCallCount() bool`

HasCurrentCallCount returns a boolean if a field has been set.

### GetId

`func (o *EndpointLine) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointLine) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointLine) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointLine) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EndpointLine) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointLine) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointLine) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointLine) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistered

`func (o *EndpointLine) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *EndpointLine) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *EndpointLine) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *EndpointLine) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetTechnology

`func (o *EndpointLine) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *EndpointLine) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *EndpointLine) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *EndpointLine) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
