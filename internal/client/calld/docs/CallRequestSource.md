# CallRequestSource

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AllLines** | Pointer to **bool** | Use all of the lines of the user to make the call (ignored when &#x60;line_id&#x60; is specified). | [optional]
**AutoAnswer** | Pointer to **bool** | Inform the phone that it should answer automatically. Limitation: this does not work if &#x60;all_lines&#x60; is true, if &#x60;from_mobile&#x60; is true or if the phone is SCCP. | [optional]
**FromMobile** | Pointer to **bool** | Start the call from the user&#39;s mobile phone. Default is False | [optional]
**LineId** | Pointer to **int32** | ID of the line of the user used to make the call. Default is the main line of the user. | [optional]
**User** | **string** | UUID of the user making the call |

## Methods

### NewCallRequestSource

`func NewCallRequestSource(user string, ) *CallRequestSource`

NewCallRequestSource instantiates a new CallRequestSource object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestSourceWithDefaults

`func NewCallRequestSourceWithDefaults() *CallRequestSource`

NewCallRequestSourceWithDefaults instantiates a new CallRequestSource object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAllLines

`func (o *CallRequestSource) GetAllLines() bool`

GetAllLines returns the AllLines field if non-nil, zero value otherwise.

### GetAllLinesOk

`func (o *CallRequestSource) GetAllLinesOk() (*bool, bool)`

GetAllLinesOk returns a tuple with the AllLines field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAllLines

`func (o *CallRequestSource) SetAllLines(v bool)`

SetAllLines sets AllLines field to given value.

### HasAllLines

`func (o *CallRequestSource) HasAllLines() bool`

HasAllLines returns a boolean if a field has been set.

### GetAutoAnswer

`func (o *CallRequestSource) GetAutoAnswer() bool`

GetAutoAnswer returns the AutoAnswer field if non-nil, zero value otherwise.

### GetAutoAnswerOk

`func (o *CallRequestSource) GetAutoAnswerOk() (*bool, bool)`

GetAutoAnswerOk returns a tuple with the AutoAnswer field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAutoAnswer

`func (o *CallRequestSource) SetAutoAnswer(v bool)`

SetAutoAnswer sets AutoAnswer field to given value.

### HasAutoAnswer

`func (o *CallRequestSource) HasAutoAnswer() bool`

HasAutoAnswer returns a boolean if a field has been set.

### GetFromMobile

`func (o *CallRequestSource) GetFromMobile() bool`

GetFromMobile returns the FromMobile field if non-nil, zero value otherwise.

### GetFromMobileOk

`func (o *CallRequestSource) GetFromMobileOk() (*bool, bool)`

GetFromMobileOk returns a tuple with the FromMobile field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFromMobile

`func (o *CallRequestSource) SetFromMobile(v bool)`

SetFromMobile sets FromMobile field to given value.

### HasFromMobile

`func (o *CallRequestSource) HasFromMobile() bool`

HasFromMobile returns a boolean if a field has been set.

### GetLineId

`func (o *CallRequestSource) GetLineId() int32`

GetLineId returns the LineId field if non-nil, zero value otherwise.

### GetLineIdOk

`func (o *CallRequestSource) GetLineIdOk() (*int32, bool)`

GetLineIdOk returns a tuple with the LineId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLineId

`func (o *CallRequestSource) SetLineId(v int32)`

SetLineId sets LineId field to given value.

### HasLineId

`func (o *CallRequestSource) HasLineId() bool`

HasLineId returns a boolean if a field has been set.

### GetUser

`func (o *CallRequestSource) GetUser() string`

GetUser returns the User field if non-nil, zero value otherwise.

### GetUserOk

`func (o *CallRequestSource) GetUserOk() (*string, bool)`

GetUserOk returns a tuple with the User field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUser

`func (o *CallRequestSource) SetUser(v string)`

SetUser sets User field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
