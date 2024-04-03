# DestinationApplicationCustom

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | Pointer to **string** | MUST be &#39;custom&#39; | [optional]
**ApplicationUuid** | Pointer to **string** | The UUID of the application. | [optional]

## Methods

### NewDestinationApplicationCustom

`func NewDestinationApplicationCustom(type_ string, ) *DestinationApplicationCustom`

NewDestinationApplicationCustom instantiates a new DestinationApplicationCustom object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationCustomWithDefaults

`func NewDestinationApplicationCustomWithDefaults() *DestinationApplicationCustom`

NewDestinationApplicationCustomWithDefaults instantiates a new DestinationApplicationCustom object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationCustom) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationCustom) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationCustom) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationCustom) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationCustom) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationCustom) SetApplication(v string)`

SetApplication sets Application field to given value.

### HasApplication

`func (o *DestinationApplicationCustom) HasApplication() bool`

HasApplication returns a boolean if a field has been set.

### GetApplicationUuid

`func (o *DestinationApplicationCustom) GetApplicationUuid() string`

GetApplicationUuid returns the ApplicationUuid field if non-nil, zero value otherwise.

### GetApplicationUuidOk

`func (o *DestinationApplicationCustom) GetApplicationUuidOk() (*string, bool)`

GetApplicationUuidOk returns a tuple with the ApplicationUuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplicationUuid

`func (o *DestinationApplicationCustom) SetApplicationUuid(v string)`

SetApplicationUuid sets ApplicationUuid field to given value.

### HasApplicationUuid

`func (o *DestinationApplicationCustom) HasApplicationUuid() bool`

HasApplicationUuid returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
