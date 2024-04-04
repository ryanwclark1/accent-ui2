# ConferenceItems

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | Pointer to [**[]Conference**](Conference.md) |  | [optional]
**Total** | **int32** |  |

## Methods

### NewConferenceItems

`func NewConferenceItems(total int32, ) *ConferenceItems`

NewConferenceItems instantiates a new ConferenceItems object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConferenceItemsWithDefaults

`func NewConferenceItemsWithDefaults() *ConferenceItems`

NewConferenceItemsWithDefaults instantiates a new ConferenceItems object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *ConferenceItems) GetItems() []Conference`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ConferenceItems) GetItemsOk() (*[]Conference, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ConferenceItems) SetItems(v []Conference)`

SetItems sets Items field to given value.

### HasItems

`func (o *ConferenceItems) HasItems() bool`

HasItems returns a boolean if a field has been set.

### GetTotal

`func (o *ConferenceItems) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ConferenceItems) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ConferenceItems) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
