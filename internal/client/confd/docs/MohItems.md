# MohItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Moh**](Moh.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewMohItems

`func NewMohItems(total int32, ) *MohItems`

NewMohItems instantiates a new MohItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMohItemsWithDefaults

`func NewMohItemsWithDefaults() *MohItems`

NewMohItemsWithDefaults instantiates a new MohItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *MohItems) GetItems() []Moh`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *MohItems) GetItemsOk() (*[]Moh, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *MohItems) SetItems(v []Moh)`

SetItems sets Items field to given value.

### HasItems

`func (o *MohItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *MohItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *MohItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *MohItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
