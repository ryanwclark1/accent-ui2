# EndpointSipRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AuthSectionOptions** | Pointer to **[][]string** | A list of PJSIP auth section options for this endpoint. Only &#x60;username&#x60; is supported | [optional]
**Label** | Pointer to **string** | The human readable name for this configuration | [optional]
**Name** | Pointer to **string** | The name of the PJSIP entity, auto-generated if missing | [optional]
**Uuid** | Pointer to **string** | The UUID of this resource | [optional] [readonly]

## Methods

### NewEndpointSipRelationBase

`func NewEndpointSipRelationBase() *EndpointSipRelationBase`

NewEndpointSipRelationBase instantiates a new EndpointSipRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEndpointSipRelationBaseWithDefaults

`func NewEndpointSipRelationBaseWithDefaults() *EndpointSipRelationBase`

NewEndpointSipRelationBaseWithDefaults instantiates a new EndpointSipRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuthSectionOptions

`func (o *EndpointSipRelationBase) GetAuthSectionOptions() [][]string`

GetAuthSectionOptions returns the AuthSectionOptions field if non-nil, zero value otherwise.

### GetAuthSectionOptionsOk

`func (o *EndpointSipRelationBase) GetAuthSectionOptionsOk() (*[][]string, bool)`

GetAuthSectionOptionsOk returns a tuple with the AuthSectionOptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuthSectionOptions

`func (o *EndpointSipRelationBase) SetAuthSectionOptions(v [][]string)`

SetAuthSectionOptions sets AuthSectionOptions field to given value.

### HasAuthSectionOptions

`func (o *EndpointSipRelationBase) HasAuthSectionOptions() bool`

HasAuthSectionOptions returns a boolean if a field has been set.

### GetLabel

`func (o *EndpointSipRelationBase) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *EndpointSipRelationBase) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *EndpointSipRelationBase) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *EndpointSipRelationBase) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *EndpointSipRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EndpointSipRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EndpointSipRelationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *EndpointSipRelationBase) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUuid

`func (o *EndpointSipRelationBase) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EndpointSipRelationBase) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EndpointSipRelationBase) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *EndpointSipRelationBase) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
