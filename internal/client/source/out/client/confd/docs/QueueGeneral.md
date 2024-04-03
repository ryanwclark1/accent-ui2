# QueueGeneral

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Options** | Pointer to **map[string]interface{}** | General Queue options. These options must be unique and unordered. Option must have the following form:  &#x60;&#x60;&#x60; {   \&quot;options\&quot;: {     \&quot;name1\&quot;: \&quot;value1\&quot;,     ...   } } &#x60;&#x60;&#x60;   | [optional]

## Methods

### NewQueueGeneral

`func NewQueueGeneral() *QueueGeneral`

NewQueueGeneral instantiates a new QueueGeneral object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueGeneralWithDefaults

`func NewQueueGeneralWithDefaults() *QueueGeneral`

NewQueueGeneralWithDefaults instantiates a new QueueGeneral object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetOptions

`func (o *QueueGeneral) GetOptions() map[string]interface{}`

GetOptions returns the Options field if non-nil, zero value otherwise.

### GetOptionsOk

`func (o *QueueGeneral) GetOptionsOk() (*map[string]interface{}, bool)`

GetOptionsOk returns a tuple with the Options field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOptions

`func (o *QueueGeneral) SetOptions(v map[string]interface{})`

SetOptions sets Options field to given value.

### HasOptions

`func (o *QueueGeneral) HasOptions() bool`

HasOptions returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
