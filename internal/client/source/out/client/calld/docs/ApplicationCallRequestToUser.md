# ApplicationCallRequestToUser

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoanswer** | Pointer to **bool** |  | [optional] [default to false]
**DisplayedCallerIdName** | Pointer to **string** |  | [optional]
**DisplayedCallerIdNumber** | Pointer to **string** |  | [optional]
**UserUuid** | **string** |  |
**Variables** | Pointer to **map[string]interface{}** | channel variables that should be assigned on this new channel | [optional]

## Methods

### NewApplicationCallRequestToUser

`func NewApplicationCallRequestToUser(userUuid string, ) *ApplicationCallRequestToUser`

NewApplicationCallRequestToUser instantiates a new ApplicationCallRequestToUser object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCallRequestToUserWithDefaults

`func NewApplicationCallRequestToUserWithDefaults() *ApplicationCallRequestToUser`

NewApplicationCallRequestToUserWithDefaults instantiates a new ApplicationCallRequestToUser object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoanswer

`func (o *ApplicationCallRequestToUser) GetAutoanswer() bool`

GetAutoanswer returns the Autoanswer field if non-nil, zero value otherwise.

### GetAutoanswerOk

`func (o *ApplicationCallRequestToUser) GetAutoanswerOk() (*bool, bool)`

GetAutoanswerOk returns a tuple with the Autoanswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoanswer

`func (o *ApplicationCallRequestToUser) SetAutoanswer(v bool)`

SetAutoanswer sets Autoanswer field to given value.

### HasAutoanswer

`func (o *ApplicationCallRequestToUser) HasAutoanswer() bool`

HasAutoanswer returns a boolean if a field has been set.

### GetDisplayedCallerIdName

`func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdName() string`

GetDisplayedCallerIdName returns the DisplayedCallerIdName field if non-nil, zero value otherwise.

### GetDisplayedCallerIdNameOk

`func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNameOk() (*string, bool)`

GetDisplayedCallerIdNameOk returns a tuple with the DisplayedCallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedCallerIdName

`func (o *ApplicationCallRequestToUser) SetDisplayedCallerIdName(v string)`

SetDisplayedCallerIdName sets DisplayedCallerIdName field to given value.

### HasDisplayedCallerIdName

`func (o *ApplicationCallRequestToUser) HasDisplayedCallerIdName() bool`

HasDisplayedCallerIdName returns a boolean if a field has been set.

### GetDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNumber() string`

GetDisplayedCallerIdNumber returns the DisplayedCallerIdNumber field if non-nil, zero value otherwise.

### GetDisplayedCallerIdNumberOk

`func (o *ApplicationCallRequestToUser) GetDisplayedCallerIdNumberOk() (*string, bool)`

GetDisplayedCallerIdNumberOk returns a tuple with the DisplayedCallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToUser) SetDisplayedCallerIdNumber(v string)`

SetDisplayedCallerIdNumber sets DisplayedCallerIdNumber field to given value.

### HasDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToUser) HasDisplayedCallerIdNumber() bool`

HasDisplayedCallerIdNumber returns a boolean if a field has been set.

### GetUserUuid

`func (o *ApplicationCallRequestToUser) GetUserUuid() string`

GetUserUuid returns the UserUuid field if non-nil, zero value otherwise.

### GetUserUuidOk

`func (o *ApplicationCallRequestToUser) GetUserUuidOk() (*string, bool)`

GetUserUuidOk returns a tuple with the UserUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUserUuid

`func (o *ApplicationCallRequestToUser) SetUserUuid(v string)`

SetUserUuid sets UserUuid field to given value.

### GetVariables

`func (o *ApplicationCallRequestToUser) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ApplicationCallRequestToUser) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ApplicationCallRequestToUser) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ApplicationCallRequestToUser) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
