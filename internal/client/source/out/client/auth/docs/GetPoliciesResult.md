# GetPoliciesResult

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Items** | [**[]PolicyResult**](PolicyResult.md) | A paginated list of policies |
**Total** | **int32** | The number of policies matching the searched term |

## Methods

### NewGetPoliciesResult

`func NewGetPoliciesResult(items []PolicyResult, total int32, ) *GetPoliciesResult`

NewGetPoliciesResult instantiates a new GetPoliciesResult object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGetPoliciesResultWithDefaults

`func NewGetPoliciesResultWithDefaults() *GetPoliciesResult`

NewGetPoliciesResultWithDefaults instantiates a new GetPoliciesResult object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetItems

`func (o *GetPoliciesResult) GetItems() []PolicyResult`

GetItems returns the Items field if non-nil, zero value otherwise.

### GetItemsOk

`func (o *GetPoliciesResult) GetItemsOk() (*[]PolicyResult, bool)`

GetItemsOk returns a tuple with the Items field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetItems

`func (o *GetPoliciesResult) SetItems(v []PolicyResult)`

SetItems sets Items field to given value.

### GetTotal

`func (o *GetPoliciesResult) GetTotal() int32`

GetTotal returns the Total field if non-nil, zero value otherwise.

### GetTotalOk

`func (o *GetPoliciesResult) GetTotalOk() (*int32, bool)`

GetTotalOk returns a tuple with the Total field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTotal

`func (o *GetPoliciesResult) SetTotal(v int32)`

SetTotal sets Total field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
