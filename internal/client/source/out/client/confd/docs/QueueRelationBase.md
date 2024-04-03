# QueueRelationBase

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | Pointer to **int32** | The id of the queue | [optional] [readonly]
**Label** | Pointer to **string** | The label of the queue | [optional]
**Name** | Pointer to **string** | The name of the queue. Cannot be modified | [optional]

## Methods

### NewQueueRelationBase

`func NewQueueRelationBase() *QueueRelationBase`

NewQueueRelationBase instantiates a new QueueRelationBase object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueRelationBaseWithDefaults

`func NewQueueRelationBaseWithDefaults() *QueueRelationBase`

NewQueueRelationBaseWithDefaults instantiates a new QueueRelationBase object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *QueueRelationBase) GetId() int32`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *QueueRelationBase) GetIdOk() (*int32, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *QueueRelationBase) SetId(v int32)`

SetId sets Id field to given value.

### HasId

`func (o *QueueRelationBase) HasId() bool`

HasId returns a boolean if a field has been set.

### GetLabel

`func (o *QueueRelationBase) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *QueueRelationBase) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *QueueRelationBase) SetLabel(v string)`

SetLabel sets Label field to given value.

### HasLabel

`func (o *QueueRelationBase) HasLabel() bool`

HasLabel returns a boolean if a field has been set.

### GetName

`func (o *QueueRelationBase) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *QueueRelationBase) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *QueueRelationBase) SetName(v string)`

SetName sets Name field to given value.

### HasName

`func (o *QueueRelationBase) HasName() bool`

HasName returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
