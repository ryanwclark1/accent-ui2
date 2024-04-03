# UserCallPermission

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallPermissionId** | **int32** | Call Permission&#39;s ID |
**UserId** | **int32** | User&#39;s ID |

## Methods

### NewUserCallPermission

`func NewUserCallPermission(callPermissionId int32, userId int32, ) *UserCallPermission`

NewUserCallPermission instantiates a new UserCallPermission object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCallPermissionWithDefaults

`func NewUserCallPermissionWithDefaults() *UserCallPermission`

NewUserCallPermissionWithDefaults instantiates a new UserCallPermission object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallPermissionId

`func (o *UserCallPermission) GetCallPermissionId() int32`

GetCallPermissionId returns the CallPermissionId field if non-nil, zero value otherwise.

### GetCallPermissionIdOk

`func (o *UserCallPermission) GetCallPermissionIdOk() (*int32, bool)`

GetCallPermissionIdOk returns a tuple with the CallPermissionId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallPermissionId

`func (o *UserCallPermission) SetCallPermissionId(v int32)`

SetCallPermissionId sets CallPermissionId field to given value.

### GetUserId

`func (o *UserCallPermission) GetUserId() int32`

GetUserId returns the UserId field if non-nil, zero value otherwise.

### GetUserIdOk

`func (o *UserCallPermission) GetUserIdOk() (*int32, bool)`

GetUserIdOk returns a tuple with the UserId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserId

`func (o *UserCallPermission) SetUserId(v int32)`

SetUserId sets UserId field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
