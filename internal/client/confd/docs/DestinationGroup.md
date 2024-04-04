# DestinationGroup

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GroupId** | **int32** | The id of the group |
**RingTime** | Pointer to **float32** |  | [optional]
**Type** | **string** | MUST be &#39;group&#39; |

## Methods

### NewDestinationGroup

`func NewDestinationGroup(groupId int32, type_ string, ) *DestinationGroup`

NewDestinationGroup instantiates a new DestinationGroup object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationGroupWithDefaults

`func NewDestinationGroupWithDefaults() *DestinationGroup`

NewDestinationGroupWithDefaults instantiates a new DestinationGroup object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGroupId

`func (o *DestinationGroup) GetGroupId() int32`

GetGroupId returns the GroupId field if non-nil, zero value otherwise.

### GetGroupIdOk

`func (o *DestinationGroup) GetGroupIdOk() (*int32, bool)`

GetGroupIdOk returns a tuple with the GroupId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroupId

`func (o *DestinationGroup) SetGroupId(v int32)`

SetGroupId sets GroupId field to given value.

### GetRingTime

`func (o *DestinationGroup) GetRingTime() float32`

GetRingTime returns the RingTime field if non-nil, zero value otherwise.

### GetRingTimeOk

`func (o *DestinationGroup) GetRingTimeOk() (*float32, bool)`

GetRingTimeOk returns a tuple with the RingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTime

`func (o *DestinationGroup) SetRingTime(v float32)`

SetRingTime sets RingTime field to given value.

### HasRingTime

`func (o *DestinationGroup) HasRingTime() bool`

HasRingTime returns a boolean if a field has been set.

### GetType

`func (o *DestinationGroup) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationGroup) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationGroup) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
