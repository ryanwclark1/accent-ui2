# Room

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Uuid** | Pointer to **string** | The UUID of the room | [optional] [readonly]
**Name** | Pointer to **string** | The name of the room | [optional]
**Users** | [**[]RoomUser**](RoomUser.md) |  |

## Methods

### NewRoom

`func NewRoom(users []RoomUser, ) *Room`

NewRoom instantiates a new Room object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewRoomWithDefaults

`func NewRoomWithDefaults() *Room`

NewRoomWithDefaults instantiates a new Room object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetUuid

`func (o *Room) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *Room) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *Room) SetUuid(v string)`

SetUuid sets Uuid field to given value.

### HasUuid

`func (o *Room) HasUuid() bool`

HasUuid returns a boolean if a field has been set.

### GetName

`func (o *Room) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Room) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Room) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *Room) HasName() bool`

HasName returns a boolean if a field has been set.

### GetUsers

`func (o *Room) GetUsers() []RoomUser`

GetUsers returns the Users field if non-nil, zero value otherwise.

### GetUsersOk

`func (o *Room) GetUsersOk() (*[]RoomUser, bool)`

GetUsersOk returns a tuple with the Users field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUsers

`func (o *Room) SetUsers(v []RoomUser)`

SetUsers sets Users field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
