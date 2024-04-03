# DestinationUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MohUuid** | Pointer to **string** | The UUID of the music on hold to use instead of a ringback tone. | [optional]
**RingTime** | Pointer to **float32** |  | [optional]
**Type** | **string** | MUST be &#39;user&#39; |
**UserId** | **int32** | The id of the user |

## Methods

### NewDestinationUser

`func NewDestinationUser(type_ string, userId int32, ) *DestinationUser`

NewDestinationUser instantiates a new DestinationUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationUserWithDefaults

`func NewDestinationUserWithDefaults() *DestinationUser`

NewDestinationUserWithDefaults instantiates a new DestinationUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMohUuid

`func (o *DestinationUser) GetMohUuid() string`

GetMohUuid returns the MohUuid field if non-nil, zero value otherwise.

### GetMohUuidOk

`func (o *DestinationUser) GetMohUuidOk() (*string, bool)`

GetMohUuidOk returns a tuple with the MohUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMohUuid

`func (o *DestinationUser) SetMohUuid(v string)`

SetMohUuid sets MohUuid field to given value.

### HasMohUuid

`func (o *DestinationUser) HasMohUuid() bool`

HasMohUuid returns a boolean if a field has been set.

### GetRingTime

`func (o *DestinationUser) GetRingTime() float32`

GetRingTime returns the RingTime field if non-nil, zero value otherwise.

### GetRingTimeOk

`func (o *DestinationUser) GetRingTimeOk() (*float32, bool)`

GetRingTimeOk returns a tuple with the RingTime field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRingTime

`func (o *DestinationUser) SetRingTime(v float32)`

SetRingTime sets RingTime field to given value.

### HasRingTime

`func (o *DestinationUser) HasRingTime() bool`

HasRingTime returns a boolean if a field has been set.

### GetType

`func (o *DestinationUser) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationUser) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationUser) SetType(v string)`

SetType sets Type field to given value.

### GetUserId

`func (o *DestinationUser) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *DestinationUser) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *DestinationUser) SetUserId(v int32)`

SetUserId sets UserId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
