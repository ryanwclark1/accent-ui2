# UserCallRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLines** | Pointer to **bool** | Use all of the lines of the user to make the call (ignored when &#x60;line_id&#x60; is specified). | [optional]
**AutoAnswerCaller** | Pointer to **bool** | Inform the caller phone that it should answer automatically. Limitation: this does not work if &#x60;all_lines&#x60; is true, if &#x60;from_mobile&#x60; is true or if the phone is SCCP. | [optional]
**Extension** | **string** | Extension to call |
**FromMobile** | Pointer to **bool** | Start the call from the user&#39;s mobile phone. Default is False. Limitation: this feature may return a wrong call_id if the outgoing call used to dial the mobile number has more than one associated trunk. | [optional]
**LineId** | Pointer to **int32** | ID of the line of the user used to make the call. Default is the main line of the user. | [optional]
**Variables** | Pointer to **map[string]interface{}** | Channel variables to set | [optional]

## Methods

### NewUserCallRequest

`func NewUserCallRequest(extension string, ) *UserCallRequest`

NewUserCallRequest instantiates a new UserCallRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUserCallRequestWithDefaults

`func NewUserCallRequestWithDefaults() *UserCallRequest`

NewUserCallRequestWithDefaults instantiates a new UserCallRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLines

`func (o *UserCallRequest) GetAllLines() bool`

GetAllLines returns the AllLines field if non-nil, zero value otherwise.

### GetAllLinesOk

`func (o *UserCallRequest) GetAllLinesOk() (*bool, bool)`

GetAllLinesOk returns a tuple with the AllLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLines

`func (o *UserCallRequest) SetAllLines(v bool)`

SetAllLines sets AllLines field to given value.

### HasAllLines

`func (o *UserCallRequest) HasAllLines() bool`

HasAllLines returns a boolean if a field has been set.

### GetAutoAnswerCaller

`func (o *UserCallRequest) GetAutoAnswerCaller() bool`

GetAutoAnswerCaller returns the AutoAnswerCaller field if non-nil, zero value otherwise.

### GetAutoAnswerCallerOk

`func (o *UserCallRequest) GetAutoAnswerCallerOk() (*bool, bool)`

GetAutoAnswerCallerOk returns a tuple with the AutoAnswerCaller field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswerCaller

`func (o *UserCallRequest) SetAutoAnswerCaller(v bool)`

SetAutoAnswerCaller sets AutoAnswerCaller field to given value.

### HasAutoAnswerCaller

`func (o *UserCallRequest) HasAutoAnswerCaller() bool`

HasAutoAnswerCaller returns a boolean if a field has been set.

### GetExtension

`func (o *UserCallRequest) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *UserCallRequest) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *UserCallRequest) SetExtension(v string)`

SetExtension sets Extension field to given value.

### GetFromMobile

`func (o *UserCallRequest) GetFromMobile() bool`

GetFromMobile returns the FromMobile field if non-nil, zero value otherwise.

### GetFromMobileOk

`func (o *UserCallRequest) GetFromMobileOk() (*bool, bool)`

GetFromMobileOk returns a tuple with the FromMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMobile

`func (o *UserCallRequest) SetFromMobile(v bool)`

SetFromMobile sets FromMobile field to given value.

### HasFromMobile

`func (o *UserCallRequest) HasFromMobile() bool`

HasFromMobile returns a boolean if a field has been set.

### GetLineId

`func (o *UserCallRequest) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *UserCallRequest) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *UserCallRequest) SetLineId(v int32)`

SetLineId sets LineId field to given value.

### HasLineId

`func (o *UserCallRequest) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### GetVariables

`func (o *UserCallRequest) GetVariables() map[string]interface{}`

GetVariables returns the Variables field if non-nil, zero value otherwise.

### GetVariablesOk

`func (o *UserCallRequest) GetVariablesOk() (*map[string]interface{}, bool)`

GetVariablesOk returns a tuple with the Variables field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVariables

`func (o *UserCallRequest) SetVariables(v map[string]interface{})`

SetVariables sets Variables field to given value.

### HasVariables

`func (o *UserCallRequest) HasVariables() bool`

HasVariables returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
