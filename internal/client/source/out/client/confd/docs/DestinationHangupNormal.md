# DestinationHangupNormal

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;hangup&#39; |
**Cause** | Pointer to **string** | MUST be &#39;normal&#39; | [optional]

## Methods

### NewDestinationHangupNormal

`func NewDestinationHangupNormal(type_ string, ) *DestinationHangupNormal`

NewDestinationHangupNormal instantiates a new DestinationHangupNormal object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationHangupNormalWithDefaults

`func NewDestinationHangupNormalWithDefaults() *DestinationHangupNormal`

NewDestinationHangupNormalWithDefaults instantiates a new DestinationHangupNormal object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationHangupNormal) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationHangupNormal) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationHangupNormal) SetType(v string)`

SetType sets Type field to given value.

### GetCause

`func (o *DestinationHangupNormal) GetCause() string`

GetCause returns the Cause field if non-nil, zero value otherwise.

### GetCauseOk

`func (o *DestinationHangupNormal) GetCauseOk() (*string, bool)`

GetCauseOk returns a tuple with the Cause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCause

`func (o *DestinationHangupNormal) SetCause(v string)`

SetCause sets Cause field to given value.

### HasCause

`func (o *DestinationHangupNormal) HasCause() bool`

HasCause returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
