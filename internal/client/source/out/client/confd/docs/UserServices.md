# UserServices

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Dnd** | Pointer to [**UserService**](UserService.md) |  | [optional]
**Incallfilter** | Pointer to [**UserService**](UserService.md) |  | [optional]

## Methods

### NewUserServices

`func NewUserServices() *UserServices`

NewUserServices instantiates a new UserServices object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserServicesWithDefaults

`func NewUserServicesWithDefaults() *UserServices`

NewUserServicesWithDefaults instantiates a new UserServices object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetDnd

`func (o *UserServices) GetDnd() UserService`

GetDnd returns the Dnd field if non-nil, zero value otherwise.

### GetDndOk

`func (o *UserServices) GetDndOk() (*UserService, bool)`

GetDndOk returns a tuple with the Dnd field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDnd

`func (o *UserServices) SetDnd(v UserService)`

SetDnd sets Dnd field to given value.

### HasDnd

`func (o *UserServices) HasDnd() bool`

HasDnd returns a boolean if a field has been set.

### GetIncallfilter

`func (o *UserServices) GetIncallfilter() UserService`

GetIncallfilter returns the Incallfilter field if non-nil, zero value otherwise.

### GetIncallfilterOk

`func (o *UserServices) GetIncallfilterOk() (*UserService, bool)`

GetIncallfilterOk returns a tuple with the Incallfilter field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIncallfilter

`func (o *UserServices) SetIncallfilter(v UserService)`

SetIncallfilter sets Incallfilter field to given value.

### HasIncallfilter

`func (o *UserServices) HasIncallfilter() bool`

HasIncallfilter returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
