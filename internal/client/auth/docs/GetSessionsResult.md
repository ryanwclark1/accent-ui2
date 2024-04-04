# GetSessionsResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Filtered** | **int32** | The number of sessions matching the searched term. |
**Items** | [**[]SessionResult**](SessionResult.md) | A paginated list of sessions |
**Total** | **int32** | The number of sessions. |

## Methods

### NewGetSessionsResult

`func NewGetSessionsResult(filtered int32, items []SessionResult, total int32, ) *GetSessionsResult`

NewGetSessionsResult instantiates a new GetSessionsResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetSessionsResultWithDefaults

`func NewGetSessionsResultWithDefaults() *GetSessionsResult`

NewGetSessionsResultWithDefaults instantiates a new GetSessionsResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFiltered

`func (o *GetSessionsResult) GetFiltered() int32`

GetFiltered returns the Filtered field if non-nil, zero value otherwise.

### GetFilteredOk

`func (o *GetSessionsResult) GetFilteredOk() (*int32, bool)`

GetFilteredOk returns a tuple with the Filtered field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFiltered

`func (o *GetSessionsResult) SetFiltered(v int32)`

SetFiltered sets Filtered field to given value.

### GetItems

`func (o *GetSessionsResult) GetItems() []SessionResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetSessionsResult) GetItemsOk() (*[]SessionResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetSessionsResult) SetItems(v []SessionResult)`

SetItems sets Items field to given value.

### GetTotal

`func (o *GetSessionsResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetSessionsResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetSessionsResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
