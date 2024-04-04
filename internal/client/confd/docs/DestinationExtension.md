# DestinationExtension

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Context** | **string** | Context of the extension |
**Exten** | **string** |  |
**Type** | **string** | MUST be &#39;extension&#39; |

## Methods

### NewDestinationExtension

`func NewDestinationExtension(context string, exten string, type_ string, ) *DestinationExtension`

NewDestinationExtension instantiates a new DestinationExtension object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationExtensionWithDefaults

`func NewDestinationExtensionWithDefaults() *DestinationExtension`

NewDestinationExtensionWithDefaults instantiates a new DestinationExtension object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetContext

`func (o *DestinationExtension) GetContext() string`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationExtension) GetContextOk() (*string, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationExtension) SetContext(v string)`

SetContext sets Context field to given value.

### GetExten

`func (o *DestinationExtension) GetExten() string`

GetExten returns the Exten field if non-nil, zero value otherwise.

### GetExtenOk

`func (o *DestinationExtension) GetExtenOk() (*string, bool)`

GetExtenOk returns a tuple with the Exten field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetExten

`func (o *DestinationExtension) SetExten(v string)`

SetExten sets Exten field to given value.

### GetType

`func (o *DestinationExtension) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationExtension) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationExtension) SetType(v string)`

SetType sets Type field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
