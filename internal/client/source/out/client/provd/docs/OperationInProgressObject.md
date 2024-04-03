# OperationInProgressObject

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Status** | Pointer to **string** | The &#x60;&#x60;status&#x60;&#x60; field describes the current status of the operation. The format is &#x60;&#x60;[label|]state[;current[/end]](\\(sub_oips\\))*&#x60;&#x60;. Here are some examples:* progress * download|progress * download|progress;10 * download|progress;10/100 * download|progress(file_1|progress;20/100)(file_2|waiting;0/50) * download|progress;20/150(file_1|progress)(file_2|waiting) * op|progress(op1|progress(op11|progress)(op12|waiting))(op2|progress)  The state of an operation is either &#x60;&#x60;waiting&#x60;&#x60;, &#x60;&#x60;progress&#x60;&#x60;, &#x60;&#x60;success&#x60;&#x60; or &#x60;&#x60;fail&#x60;&#x60;.  | [optional]

## Methods

### NewOperationInProgressObject

`func NewOperationInProgressObject() *OperationInProgressObject`

NewOperationInProgressObject instantiates a new OperationInProgressObject object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewOperationInProgressObjectWithDefaults

`func NewOperationInProgressObjectWithDefaults() *OperationInProgressObject`

NewOperationInProgressObjectWithDefaults instantiates a new OperationInProgressObject object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStatus

`func (o *OperationInProgressObject) GetStatus() string`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *OperationInProgressObject) GetStatusOk() (*string, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *OperationInProgressObject) SetStatus(v string)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *OperationInProgressObject) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
