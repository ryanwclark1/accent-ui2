# SoundItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Sound**](Sound.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewSoundItems

`func NewSoundItems(total int32, ) *SoundItems`

NewSoundItems instantiates a new SoundItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSoundItemsWithDefaults

`func NewSoundItemsWithDefaults() *SoundItems`

NewSoundItemsWithDefaults instantiates a new SoundItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *SoundItems) GetItems() []Sound`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *SoundItems) GetItemsOk() (*[]Sound, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *SoundItems) SetItems(v []Sound)`

SetItems sets Items field to given value.

### HasItems

`func (o *SoundItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *SoundItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *SoundItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *SoundItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
