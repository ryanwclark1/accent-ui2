# CallRequestDestination

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** |  |
**Extension** | **string** |  |
**Priority** | **int32** |  |

## Methods

### NewCallRequestDestination

`func NewCallRequestDestination(context string, extension string, priority int32, ) *CallRequestDestination`

NewCallRequestDestination instantiates a new CallRequestDestination object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCallRequestDestinationWithDefaults

`func NewCallRequestDestinationWithDefaults() *CallRequestDestination`

NewCallRequestDestinationWithDefaults instantiates a new CallRequestDestination object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *CallRequestDestination) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *CallRequestDestination) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *CallRequestDestination) SetContext(v string)`

SetContext sets Context field to given value.

### GetExtension

`func (o *CallRequestDestination) GetExtension() string`

GetExtension returns the Extension field if non-nil, zero value otherwise.

### GetExtensionOk

`func (o *CallRequestDestination) GetExtensionOk() (*string, bool)`

GetExtensionOk returns a tuple with the Extension field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExtension

`func (o *CallRequestDestination) SetExtension(v string)`

SetExtension sets Extension field to given value.

### GetPriority

`func (o *CallRequestDestination) GetPriority() int32`

GetPriority returns the Priority field if non-nil, zero value otherwise.

### GetPriorityOk

`func (o *CallRequestDestination) GetPriorityOk() (*int32, bool)`

GetPriorityOk returns a tuple with the Priority field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPriority

`func (o *CallRequestDestination) SetPriority(v int32)`

SetPriority sets Priority field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
