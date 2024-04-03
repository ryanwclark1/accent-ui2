# LinkObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Href** | Pointer to **string** | Location of the resource | [optional]
**Rel** | Pointer to **string** | Relation to the resource | [optional]

## Methods

### NewLinkObject

`func NewLinkObject() *LinkObject`

NewLinkObject instantiates a new LinkObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLinkObjectWithDefaults

`func NewLinkObjectWithDefaults() *LinkObject`

NewLinkObjectWithDefaults instantiates a new LinkObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetHref

`func (o *LinkObject) GetHref() string`

GetHref returns the Href field if non-nil, zero value otherwise.

### GetHrefOk

`func (o *LinkObject) GetHrefOk() (*string, bool)`

GetHrefOk returns a tuple with the Href field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetHref

`func (o *LinkObject) SetHref(v string)`

SetHref sets Href field to given value.

### HasHref

`func (o *LinkObject) HasHref() bool`

HasHref returns a boolean if a field has been set.

### GetRel

`func (o *LinkObject) GetRel() string`

GetRel returns the Rel field if non-nil, zero value otherwise.

### GetRelOk

`func (o *LinkObject) GetRelOk() (*string, bool)`

GetRelOk returns a tuple with the Rel field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRel

`func (o *LinkObject) SetRel(v string)`

SetRel sets Rel field to given value.

### HasRel

`func (o *LinkObject) HasRel() bool`

HasRel returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
