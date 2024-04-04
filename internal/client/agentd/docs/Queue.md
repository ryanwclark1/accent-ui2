# Queue

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**QueueId** | Pointer to **int32** | Queue&#39;s ID | [optional]

## Methods

### NewQueue

`func NewQueue() *Queue`

NewQueue instantiates a new Queue object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewQueueWithDefaults

`func NewQueueWithDefaults() *Queue`

NewQueueWithDefaults instantiates a new Queue object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetQueueId

`func (o *Queue) GetQueueId() int32`

GetQueueId returns the QueueId field if non-nil, zero value otherwise.

### GetQueueIdOk

`func (o *Queue) GetQueueIdOk() (*int32, bool)`

GetQueueIdOk returns a tuple with the QueueId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetQueueId

`func (o *Queue) SetQueueId(v int32)`

SetQueueId sets QueueId field to given value.

### HasQueueId

`func (o *Queue) HasQueueId() bool`

HasQueueId returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
