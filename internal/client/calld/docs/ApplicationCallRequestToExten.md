# ApplicationCallRequestToExten

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Autoanswer** | Pointer to **bool** |  | [optional] [default to false]
**Context** | **string** |  |
**DisplayedCallerIdName** | Pointer to **string** |  | [optional]
**DisplayedCallerIdNumber** | Pointer to **string** |  | [optional]
**Exten** | **string** |  |
**Variables** | Pointer to **map[string]interface{}** | channel variables that should be assigned on this new channel | [optional]

## Methods

### NewApplicationCallRequestToExten

`func NewApplicationCallRequestToExten(context string, exten string, ) *ApplicationCallRequestToExten`

NewApplicationCallRequestToExten instantiates a new ApplicationCallRequestToExten object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewApplicationCallRequestToExtenWithDefaults

`func NewApplicationCallRequestToExtenWithDefaults() *ApplicationCallRequestToExten`

NewApplicationCallRequestToExtenWithDefaults instantiates a new ApplicationCallRequestToExten object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAutoanswer

`func (o *ApplicationCallRequestToExten) GetAutoanswer() bool`

GetAutoanswer returns the Autoanswer field if non-nil, zero value otherwise.

### GetAutoanswerOk

`func (o *ApplicationCallRequestToExten) GetAutoanswerOk() (*bool, bool)`

GetAutoanswerOk returns a tuple with the Autoanswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoanswer

`func (o *ApplicationCallRequestToExten) SetAutoanswer(v bool)`

SetAutoanswer sets Autoanswer field to given value.

### HasAutoanswer

`func (o *ApplicationCallRequestToExten) HasAutoanswer() bool`

HasAutoanswer returns a boolean if a field has been set.

### GetContext

`func (o *ApplicationCallRequestToExten) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *ApplicationCallRequestToExten) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *ApplicationCallRequestToExten) SetContext(v string)`

SetContext sets Context field to given value.

### GetDisplayedCallerIdName

`func (o *ApplicationCallRequestToExten) GetDisplayedCallerIdName() string`

GetDisplayedCallerIdName returns the DisplayedCallerIdName field if non-nil, zero value otherwise.

### GetDisplayedCallerIdNameOk

`func (o *ApplicationCallRequestToExten) GetDisplayedCallerIdNameOk() (*string, bool)`

GetDisplayedCallerIdNameOk returns a tuple with the DisplayedCallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedCallerIdName

`func (o *ApplicationCallRequestToExten) SetDisplayedCallerIdName(v string)`

SetDisplayedCallerIdName sets DisplayedCallerIdName field to given value.

### HasDisplayedCallerIdName

`func (o *ApplicationCallRequestToExten) HasDisplayedCallerIdName() bool`

HasDisplayedCallerIdName returns a boolean if a field has been set.

### GetDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToExten) GetDisplayedCallerIdNumber() string`

GetDisplayedCallerIdNumber returns the DisplayedCallerIdNumber field if non-nil, zero value otherwise.

### GetDisplayedCallerIdNumberOk

`func (o *ApplicationCallRequestToExten) GetDisplayedCallerIdNumberOk() (*string, bool)`

GetDisplayedCallerIdNumberOk returns a tuple with the DisplayedCallerIdNumber field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToExten) SetDisplayedCallerIdNumber(v string)`

SetDisplayedCallerIdNumber sets DisplayedCallerIdNumber field to given value.

### HasDisplayedCallerIdNumber

`func (o *ApplicationCallRequestToExten) HasDisplayedCallerIdNumber() bool`

HasDisplayedCallerIdNumber returns a boolean if a field has been set.

### GetExten

`func (o *ApplicationCallRequestToExten) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *ApplicationCallRequestToExten) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *ApplicationCallRequestToExten) SetExten(v string)`

SetExten sets Exten field to given value.

### GetVariables

`func (o *ApplicationCallRequestToExten) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *ApplicationCallRequestToExten) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *ApplicationCallRequestToExten) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *ApplicationCallRequestToExten) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
