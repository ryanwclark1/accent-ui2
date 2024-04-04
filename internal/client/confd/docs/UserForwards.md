# UserForwards

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Busy** | Pointer to [**UserForward**](UserForward.md) |  | [optional]
**Noanswer** | Pointer to [**UserForward**](UserForward.md) |  | [optional]
**Unconditional** | Pointer to [**UserForward**](UserForward.md) |  | [optional]

## Methods

### NewUserForwards

`func NewUserForwards() *UserForwards`

NewUserForwards instantiates a new UserForwards object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserForwardsWithDefaults

`func NewUserForwardsWithDefaults() *UserForwards`

NewUserForwardsWithDefaults instantiates a new UserForwards object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetBusy

`func (o *UserForwards) GetBusy() UserForward`

GetBusy returns the Busy field if non-nil, zero value otherwise.

### GetBusyOk

`func (o *UserForwards) GetBusyOk() (*UserForward, bool)`

GetBusyOk returns a tuple with the Busy field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusy

`func (o *UserForwards) SetBusy(v UserForward)`

SetBusy sets Busy field to given value.

### HasBusy

`func (o *UserForwards) HasBusy() bool`

HasBusy returns a boolean if a field has been set.

### GetNoanswer

`func (o *UserForwards) GetNoanswer() UserForward`

GetNoanswer returns the Noanswer field if non-nil, zero value otherwise.

### GetNoanswerOk

`func (o *UserForwards) GetNoanswerOk() (*UserForward, bool)`

GetNoanswerOk returns a tuple with the Noanswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNoanswer

`func (o *UserForwards) SetNoanswer(v UserForward)`

SetNoanswer sets Noanswer field to given value.

### HasNoanswer

`func (o *UserForwards) HasNoanswer() bool`

HasNoanswer returns a boolean if a field has been set.

### GetUnconditional

`func (o *UserForwards) GetUnconditional() UserForward`

GetUnconditional returns the Unconditional field if non-nil, zero value otherwise.

### GetUnconditionalOk

`func (o *UserForwards) GetUnconditionalOk() (*UserForward, bool)`

GetUnconditionalOk returns a tuple with the Unconditional field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnconditional

`func (o *UserForwards) SetUnconditional(v UserForward)`

SetUnconditional sets Unconditional field to given value.

### HasUnconditional

`func (o *UserForwards) HasUnconditional() bool`

HasUnconditional returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
