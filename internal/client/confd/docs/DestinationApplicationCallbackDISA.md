# DestinationApplicationCallbackDISA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | **string** | MUST be &#39;callback_disa&#39; |
**Context** | **int32** | The context of the application |
**Pin** | Pointer to **string** | the pin of the application | [optional]

## Methods

### NewDestinationApplicationCallbackDISA

`func NewDestinationApplicationCallbackDISA(type_ string, application string, context int32, ) *DestinationApplicationCallbackDISA`

NewDestinationApplicationCallbackDISA instantiates a new DestinationApplicationCallbackDISA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationCallbackDISAWithDefaults

`func NewDestinationApplicationCallbackDISAWithDefaults() *DestinationApplicationCallbackDISA`

NewDestinationApplicationCallbackDISAWithDefaults instantiates a new DestinationApplicationCallbackDISA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationCallbackDISA) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationCallbackDISA) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationCallbackDISA) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationCallbackDISA) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationCallbackDISA) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationCallbackDISA) SetApplication(v string)`

SetApplication sets Application field to given value.

### GetContext

`func (o *DestinationApplicationCallbackDISA) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationApplicationCallbackDISA) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationApplicationCallbackDISA) SetContext(v int32)`

SetContext sets Context field to given value.

### GetPin

`func (o *DestinationApplicationCallbackDISA) GetPin() string`

GetPin returns the Pin field if non-nil, zero value otherwise.

### GetPinOk

`func (o *DestinationApplicationCallbackDISA) GetPinOk() (*string, bool)`

GetPinOk returns a tuple with the Pin field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPin

`func (o *DestinationApplicationCallbackDISA) SetPin(v string)`

SetPin sets Pin field to given value.

### HasPin

`func (o *DestinationApplicationCallbackDISA) HasPin() bool`

HasPin returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
