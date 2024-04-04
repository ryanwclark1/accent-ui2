# ExternalAuthList

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | **int32** | The number of external auth matching the searched term. |
**Items** | [**[]ExternalAuth**](ExternalAuth.md) | A paginated list of external auth |
**Total** | **int32** | The number of external auth. |

## Methods

### NewExternalAuthList

`func NewExternalAuthList(filtered int32, items []ExternalAuth, total int32, ) *ExternalAuthList`

NewExternalAuthList instantiates a new ExternalAuthList object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewExternalAuthListWithDefaults

`func NewExternalAuthListWithDefaults() *ExternalAuthList`

NewExternalAuthListWithDefaults instantiates a new ExternalAuthList object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *ExternalAuthList) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *ExternalAuthList) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *ExternalAuthList) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### GetItems

`func (o *ExternalAuthList) GetItems() []ExternalAuth`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *ExternalAuthList) GetItemsOk() (*[]ExternalAuth, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *ExternalAuthList) SetItems(v []ExternalAuth)`

SetItems sets Items field to given value.

### GetTotal

`func (o *ExternalAuthList) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *ExternalAuthList) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *ExternalAuthList) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
