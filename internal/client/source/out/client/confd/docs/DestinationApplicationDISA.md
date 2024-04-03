# DestinationApplicationDISA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | **string** | MUST be &#39;disa&#39; |
**Context** | **int32** | The context of the application |
**Pin** | Pointer to **string** | the pin of the application | [optional]

## Methods

### NewDestinationApplicationDISA

`func NewDestinationApplicationDISA(type_ string, application string, context int32, ) *DestinationApplicationDISA`

NewDestinationApplicationDISA instantiates a new DestinationApplicationDISA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationDISAWithDefaults

`func NewDestinationApplicationDISAWithDefaults() *DestinationApplicationDISA`

NewDestinationApplicationDISAWithDefaults instantiates a new DestinationApplicationDISA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationDISA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationDISA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationDISA) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationDISA) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationDISA) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationDISA) SetApplication(v string)`

SetApplication sets Application field to given value.

### GetContext

`func (o *DestinationApplicationDISA) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationApplicationDISA) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationApplicationDISA) SetContext(v int32)`

SetContext sets Context field to given value.

### GetPin

`func (o *DestinationApplicationDISA) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *DestinationApplicationDISA) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *DestinationApplicationDISA) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *DestinationApplicationDISA) HasPin() bool`

HasPin returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
