# UserPostRelationIncallsIncallsInner

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**CallerIdMode** | Pointer to **string** | How the caller_id_name will be treated | [optional]
**CallerIdName** | Pointer to **string** | Name to display when calling | [optional]
**Extensions** | Pointer to [**[]UserPostRelationIncallsIncallsInnerExtensionsInner**](UserPostRelationIncallsIncallsInnerExtensionsInner.md) |  | [optional]
**GreetingSound** | Pointer to **string** | The name of the sound file to be played before redirecting the caller to the destination | [optional]
**Id** | Pointer to **int32** | The id of the incoming call | [optional] [readonly]
**PreprocessSubroutine** | Pointer to **string** | Name of the subroutine to execute in asterisk before receiving a call | [optional]

## Methods

### NewUserPostRelationIncallsIncallsInner

`func NewUserPostRelationIncallsIncallsInner() *UserPostRelationIncallsIncallsInner`

NewUserPostRelationIncallsIncallsInner instantiates a new UserPostRelationIncallsIncallsInner object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserPostRelationIncallsIncallsInnerWithDefaults

`func NewUserPostRelationIncallsIncallsInnerWithDefaults() *UserPostRelationIncallsIncallsInner`

NewUserPostRelationIncallsIncallsInnerWithDefaults instantiates a new UserPostRelationIncallsIncallsInner object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCallerIdMode

`func (o *UserPostRelationIncallsIncallsInner) GetCallerIdMode() string`

GetCallerIdMode returns the CallerIdMode field if non-nil, zero value otherwise.

### GetCallerIdModeOk

`func (o *UserPostRelationIncallsIncallsInner) GetCallerIdModeOk() (*string, bool)`

GetCallerIdModeOk returns a tuple with the CallerIdMode field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdMode

`func (o *UserPostRelationIncallsIncallsInner) SetCallerIdMode(v string)`

SetCallerIdMode sets CallerIdMode field to given value.

### HasCallerIdMode

`func (o *UserPostRelationIncallsIncallsInner) HasCallerIdMode() bool`

HasCallerIdMode returns a boolean if a field has been set.

### GetCallerIdName

`func (o *UserPostRelationIncallsIncallsInner) GetCallerIdName() string`

GetCallerIdName returns the CallerIdName field if non-nil, zero value otherwise.

### GetCallerIdNameOk

`func (o *UserPostRelationIncallsIncallsInner) GetCallerIdNameOk() (*string, bool)`

GetCallerIdNameOk returns a tuple with the CallerIdName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCallerIdName

`func (o *UserPostRelationIncallsIncallsInner) SetCallerIdName(v string)`

SetCallerIdName sets CallerIdName field to given value.

### HasCallerIdName

`func (o *UserPostRelationIncallsIncallsInner) HasCallerIdName() bool`

HasCallerIdName returns a boolean if a field has been set.

### GetExtensions

`func (o *UserPostRelationIncallsIncallsInner) GetExtensions() []UserPostRelationIncallsIncallsInnerExtensionsInner`

GetExtensions returns the Extensions field if non-nil, zero value otherwise.

### GetExtensionsOk

`func (o *UserPostRelationIncallsIncallsInner) GetExtensionsOk() (*[]UserPostRelationIncallsIncallsInnerExtensionsInner, bool)`

GetExtensionsOk returns a tuple with the Extensions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtensions

`func (o *UserPostRelationIncallsIncallsInner) SetExtensions(v []UserPostRelationIncallsIncallsInnerExtensionsInner)`

SetExtensions sets Extensions field to given value.

### HasExtensions

`func (o *UserPostRelationIncallsIncallsInner) HasExtensions() bool`

HasExtensions returns a boolean if a field has been set.

### GetGreetingSound

`func (o *UserPostRelationIncallsIncallsInner) GetGreetingSound() string`

GetGreetingSound returns the GreetingSound field if non-nil, zero value otherwise.

### GetGreetingSoundOk

`func (o *UserPostRelationIncallsIncallsInner) GetGreetingSoundOk() (*string, bool)`

GetGreetingSoundOk returns a tuple with the GreetingSound field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGreetingSound

`func (o *UserPostRelationIncallsIncallsInner) SetGreetingSound(v string)`

SetGreetingSound sets GreetingSound field to given value.

### HasGreetingSound

`func (o *UserPostRelationIncallsIncallsInner) HasGreetingSound() bool`

HasGreetingSound returns a boolean if a field has been set.

### GetId

`func (o *UserPostRelationIncallsIncallsInner) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *UserPostRelationIncallsIncallsInner) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *UserPostRelationIncallsIncallsInner) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *UserPostRelationIncallsIncallsInner) HasId() bool`

HasId returns a boolean if a field has been set.

### GetPreprocessSubroutine

`func (o *UserPostRelationIncallsIncallsInner) GetPreprocessSubroutine() string`

GetPreprocessSubroutine returns the PreprocessSubroutine field if non-nil, zero value otherwise.

### GetPreprocessSubroutineOk

`func (o *UserPostRelationIncallsIncallsInner) GetPreprocessSubroutineOk() (*string, bool)`

GetPreprocessSubroutineOk returns a tuple with the PreprocessSubroutine field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPreprocessSubroutine

`func (o *UserPostRelationIncallsIncallsInner) SetPreprocessSubroutine(v string)`

SetPreprocessSubroutine sets PreprocessSubroutine field to given value.

### HasPreprocessSubroutine

`func (o *UserPostRelationIncallsIncallsInner) HasPreprocessSubroutine() bool`

HasPreprocessSubroutine returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
