# AccessFeatureItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]AccessFeature**](AccessFeature.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewAccessFeatureItems

`func NewAccessFeatureItems(total int32, ) *AccessFeatureItems`

NewAccessFeatureItems instantiates a new AccessFeatureItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAccessFeatureItemsWithDefaults

`func NewAccessFeatureItemsWithDefaults() *AccessFeatureItems`

NewAccessFeatureItemsWithDefaults instantiates a new AccessFeatureItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *AccessFeatureItems) GetItems() []AccessFeature`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *AccessFeatureItems) GetItemsOk() (*[]AccessFeature, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *AccessFeatureItems) SetItems(v []AccessFeature)`

SetItems sets Items field to given value.

### HasItems

`func (o *AccessFeatureItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *AccessFeatureItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *AccessFeatureItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *AccessFeatureItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
