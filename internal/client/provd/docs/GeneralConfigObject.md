# GeneralConfigObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Description** | Pointer to **string** |  | [optional]
**Id** | Pointer to **string** |  | [optional]
**Links** | Pointer to [**LinksObject**](LinksObject.md) |  | [optional]
**Value** | Pointer to **string** |  | [optional]

## Methods

### NewGeneralConfigObject

`func NewGeneralConfigObject() *GeneralConfigObject`

NewGeneralConfigObject instantiates a new GeneralConfigObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGeneralConfigObjectWithDefaults

`func NewGeneralConfigObjectWithDefaults() *GeneralConfigObject`

NewGeneralConfigObjectWithDefaults instantiates a new GeneralConfigObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDescription

`func (o *GeneralConfigObject) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *GeneralConfigObject) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *GeneralConfigObject) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *GeneralConfigObject) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetId

`func (o *GeneralConfigObject) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *GeneralConfigObject) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *GeneralConfigObject) SetId(v string)`

SetId sets Id field to given value.

### HasId

`func (o *GeneralConfigObject) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLinks

`func (o *GeneralConfigObject) GetLinks() LinksObject`

GetLinks returns the Links field if non-nil, zero value otherwise.

### GetLinksOk

`func (o *GeneralConfigObject) GetLinksOk() (*LinksObject, bool)`

GetLinksOk returns a tuple with the Links field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLinks

`func (o *GeneralConfigObject) SetLinks(v LinksObject)`

SetLinks sets Links field to given value.

### HasLinks

`func (o *GeneralConfigObject) HasLinks() bool`

HasLinks returns a boolean if a field has been set.

### GetValue

`func (o *GeneralConfigObject) GetValue() string`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *GeneralConfigObject) GetValueOk() (*string, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *GeneralConfigObject) SetValue(v string)`

SetValue sets Value field to given value.

### HasValue

`func (o *GeneralConfigObject) HasValue() bool`

HasValue returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
