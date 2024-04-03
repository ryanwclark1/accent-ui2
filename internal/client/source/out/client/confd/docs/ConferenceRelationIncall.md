# ConferenceRelationIncall

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the incoming call | [optional] [readonly]
**Extensions** | Pointer to [**[]ExtensionRelationBase**](ExtensionRelationBase.md) |  | [optional] [readonly]

## Methods

### NewConferenceRelationIncall

`func NewConferenceRelationIncall() *ConferenceRelationIncall`

NewConferenceRelationIncall instantiates a new ConferenceRelationIncall object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceRelationIncallWithDefaults

`func NewConferenceRelationIncallWithDefaults() *ConferenceRelationIncall`

NewConferenceRelationIncallWithDefaults instantiates a new ConferenceRelationIncall object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *ConferenceRelationIncall) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *ConferenceRelationIncall) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *ConferenceRelationIncall) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *ConferenceRelationIncall) HasId() bool`

HasId returns a boolean if a field has been set.

### GetExtensions

`func (o *ConferenceRelationIncall) GetExtensions() []ExtensionRelationBase`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *ConferenceRelationIncall) GetExtensionsOk() (*[]ExtensionRelationBase, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *ConferenceRelationIncall) SetExtensions(v []ExtensionRelationBase)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *ConferenceRelationIncall) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
