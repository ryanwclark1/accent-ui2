# ConnectCallToUserRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Timeout** | Pointer to **int32** | timeout in seconds for the dial attempt to the targeted user, or null for no timeout(infinite ring time). Omission leads to a default timeout of 30s.  | [optional]

## Methods

### NewConnectCallToUserRequest

`func NewConnectCallToUserRequest() *ConnectCallToUserRequest`

NewConnectCallToUserRequest instantiates a new ConnectCallToUserRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewConnectCallToUserRequestWithDefaults

`func NewConnectCallToUserRequestWithDefaults() *ConnectCallToUserRequest`

NewConnectCallToUserRequestWithDefaults instantiates a new ConnectCallToUserRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetTimeout

`func (o *ConnectCallToUserRequest) GetTimeout() int32`

GetTimeout returns the Timeout field if non-nil, zero value otherwise.

### GetTimeoutOk

`func (o *ConnectCallToUserRequest) GetTimeoutOk() (*int32, bool)`

GetTimeoutOk returns a tuple with the Timeout field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTimeout

`func (o *ConnectCallToUserRequest) SetTimeout(v int32)`

SetTimeout sets Timeout field to given value.

### HasTimeout

`func (o *ConnectCallToUserRequest) HasTimeout() bool`

HasTimeout returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
