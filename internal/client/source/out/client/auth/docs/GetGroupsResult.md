# GetGroupsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | **int32** | The number of groups matching the searched term. |
**Items** | [**[]GroupResult**](GroupResult.md) | A paginated list of groups |
**Total** | **int32** | The number of groups. |

## Methods

### NewGetGroupsResult

`func NewGetGroupsResult(filtered int32, items []GroupResult, total int32, ) *GetGroupsResult`

NewGetGroupsResult instantiates a new GetGroupsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetGroupsResultWithDefaults

`func NewGetGroupsResultWithDefaults() *GetGroupsResult`

NewGetGroupsResultWithDefaults instantiates a new GetGroupsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *GetGroupsResult) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *GetGroupsResult) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *GetGroupsResult) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### GetItems

`func (o *GetGroupsResult) GetItems() []GroupResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetGroupsResult) GetItemsOk() (*[]GroupResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetGroupsResult) SetItems(v []GroupResult)`

SetItems sets Items field to given value.

### GetTotal

`func (o *GetGroupsResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetGroupsResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetGroupsResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
