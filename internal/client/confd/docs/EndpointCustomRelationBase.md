# EndpointCustomRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | Custom Endpoint ID | [optional] [readonly]
**Interface** | Pointer to **string** | Asterisk interface to use. (e.g. &#39;dahdi/i1&#39;) | [optional]

## Methods

### NewEndpointCustomRelationBase

`func NewEndpointCustomRelationBase() *EndpointCustomRelationBase`

NewEndpointCustomRelationBase instantiates a new EndpointCustomRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointCustomRelationBaseWithDefaults

`func NewEndpointCustomRelationBaseWithDefaults() *EndpointCustomRelationBase`

NewEndpointCustomRelationBaseWithDefaults instantiates a new EndpointCustomRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *EndpointCustomRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *EndpointCustomRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *EndpointCustomRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *EndpointCustomRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetInterface

`func (o *EndpointCustomRelationBase) GetInterface() string`

GetInterface returns the Interface field if non-nil, zero value otherwise.

### GetInterfaceOk

`func (o *EndpointCustomRelationBase) GetInterfaceOk() (*string, bool)`

GetInterfaceOk returns a tuple with the Interface field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetInterface

`func (o *EndpointCustomRelationBase) SetInterface(v string)`

SetInterface sets Interface field to given value.

### HasInterface

`func (o *EndpointCustomRelationBase) HasInterface() bool`

HasInterface returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
