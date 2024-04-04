# EndpointTrunk

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CurrentCallCount** | Pointer to **int32** | The number of ongoing calls on that trunk | [optional]
**Id** | Pointer to **int32** | The ID of the matching confd trunk | [optional]
**Name** | Pointer to **string** | The name of that given endpoint in Asterisk | [optional]
**Registered** | Pointer to **bool** | Wether or not this trunk is registered. | [optional]
**Technology** | Pointer to **string** | The technology of that endpoint only (SIP, IAX or custom) | [optional]

## Methods

### NewEndpointTrunk

`func NewEndpointTrunk() *EndpointTrunk`

NewEndpointTrunk instantiates a new EndpointTrunk object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointTrunkWithDefaults

`func NewEndpointTrunkWithDefaults() *EndpointTrunk`

NewEndpointTrunkWithDefaults instantiates a new EndpointTrunk object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCurrentCallCount

`func (o *EndpointTrunk) GetCurrentCallCount() int32`

GetCurrentCallCount returns the CurrentCallCount field if non-nil, zero value otherwise.

### GetCurrentCallCountOk

`func (o *EndpointTrunk) GetCurrentCallCountOk() (*int32, bool)`

GetCurrentCallCountOk returns a tuple with the CurrentCallCount field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCurrentCallCount

`func (o *EndpointTrunk) SetCurrentCallCount(v int32)`

SetCurrentCallCount sets CurrentCallCount field to given value.

### HasCurrentCallCount

`func (o *EndpointTrunk) HasCurrentCallCount() bool`

HasCurrentCallCount returns a boolean if a field has been set.

### GetId

`func (o *EndpointTrunk) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointTrunk) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointTrunk) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointTrunk) HasId() bool`

HasId returns a boolean if a field has been set.

### GetName

`func (o *EndpointTrunk) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointTrunk) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointTrunk) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointTrunk) HasName() bool`

HasName returns a boolean if a field has been set.

### GetRegistered

`func (o *EndpointTrunk) GetRegistered() bool`

GetRegistered returns the Registered field if non-nil, zero value otherwise.

### GetRegisteredOk

`func (o *EndpointTrunk) GetRegisteredOk() (*bool, bool)`

GetRegisteredOk returns a tuple with the Registered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRegistered

`func (o *EndpointTrunk) SetRegistered(v bool)`

SetRegistered sets Registered field to given value.

### HasRegistered

`func (o *EndpointTrunk) HasRegistered() bool`

HasRegistered returns a boolean if a field has been set.

### GetTechnology

`func (o *EndpointTrunk) GetTechnology() string`

GetTechnology returns the Technology field if non-nil, zero value otherwise.

### GetTechnologyOk

`func (o *EndpointTrunk) GetTechnologyOk() (*string, bool)`

GetTechnologyOk returns a tuple with the Technology field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTechnology

`func (o *EndpointTrunk) SetTechnology(v string)`

SetTechnology sets Technology field to given value.

### HasTechnology

`func (o *EndpointTrunk) HasTechnology() bool`

HasTechnology returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
