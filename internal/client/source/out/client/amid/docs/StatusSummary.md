# StatusSummary

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**AmiSocket** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**BusPublisher** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**RestApi** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]
**ServiceToken** | Pointer to [**ComponentWithStatus**](ComponentWithStatus.md) |  | [optional]

## Methods

### NewStatusSummary

`func NewStatusSummary() *StatusSummary`

NewStatusSummary instantiates a new StatusSummary object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewStatusSummaryWithDefaults

`func NewStatusSummaryWithDefaults() *StatusSummary`

NewStatusSummaryWithDefaults instantiates a new StatusSummary object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAmiSocket

`func (o *StatusSummary) GetAmiSocket() ComponentWithStatus`

GetAmiSocket returns the AmiSocket field if non-nil, zero value otherwise.

### GetAmiSocketOk

`func (o *StatusSummary) GetAmiSocketOk() (*ComponentWithStatus, bool)`

GetAmiSocketOk returns a tuple with the AmiSocket field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAmiSocket

`func (o *StatusSummary) SetAmiSocket(v ComponentWithStatus)`

SetAmiSocket sets AmiSocket field to given value.

### HasAmiSocket

`func (o *StatusSummary) HasAmiSocket() bool`

HasAmiSocket returns a boolean if a field has been set.

### GetBusPublisher

`func (o *StatusSummary) GetBusPublisher() ComponentWithStatus`

GetBusPublisher returns the BusPublisher field if non-nil, zero value otherwise.

### GetBusPublisherOk

`func (o *StatusSummary) GetBusPublisherOk() (*ComponentWithStatus, bool)`

GetBusPublisherOk returns a tuple with the BusPublisher field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBusPublisher

`func (o *StatusSummary) SetBusPublisher(v ComponentWithStatus)`

SetBusPublisher sets BusPublisher field to given value.

### HasBusPublisher

`func (o *StatusSummary) HasBusPublisher() bool`

HasBusPublisher returns a boolean if a field has been set.

### GetRestApi

`func (o *StatusSummary) GetRestApi() ComponentWithStatus`

GetRestApi returns the RestApi field if non-nil, zero value otherwise.

### GetRestApiOk

`func (o *StatusSummary) GetRestApiOk() (*ComponentWithStatus, bool)`

GetRestApiOk returns a tuple with the RestApi field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRestApi

`func (o *StatusSummary) SetRestApi(v ComponentWithStatus)`

SetRestApi sets RestApi field to given value.

### HasRestApi

`func (o *StatusSummary) HasRestApi() bool`

HasRestApi returns a boolean if a field has been set.

### GetServiceToken

`func (o *StatusSummary) GetServiceToken() ComponentWithStatus`

GetServiceToken returns the ServiceToken field if non-nil, zero value otherwise.

### GetServiceTokenOk

`func (o *StatusSummary) GetServiceTokenOk() (*ComponentWithStatus, bool)`

GetServiceTokenOk returns a tuple with the ServiceToken field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetServiceToken

`func (o *StatusSummary) SetServiceToken(v ComponentWithStatus)`

SetServiceToken sets ServiceToken field to given value.

### HasServiceToken

`func (o *StatusSummary) HasServiceToken() bool`

HasServiceToken returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
