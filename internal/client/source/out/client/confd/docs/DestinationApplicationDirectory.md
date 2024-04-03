# DestinationApplicationDirectory

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Type** | **string** | MUST be &#39;application&#39; |
**Application** | **string** | MUST be &#39;directory&#39; |
**Context** | **int32** | The context of the application |

## Methods

### NewDestinationApplicationDirectory

`func NewDestinationApplicationDirectory(type_ string, application string, context int32, ) *DestinationApplicationDirectory`

NewDestinationApplicationDirectory instantiates a new DestinationApplicationDirectory object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewDestinationApplicationDirectoryWithDefaults

`func NewDestinationApplicationDirectoryWithDefaults() *DestinationApplicationDirectory`

NewDestinationApplicationDirectoryWithDefaults instantiates a new DestinationApplicationDirectory object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetType

`func (o *DestinationApplicationDirectory) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *DestinationApplicationDirectory) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *DestinationApplicationDirectory) SetType(v string)`

SetType sets Type field to given value.

### GetApplication

`func (o *DestinationApplicationDirectory) GetApplication() string`

GetApplication returns the Application field if non-nil, zero value otherwise.

### GetApplicationOk

`func (o *DestinationApplicationDirectory) GetApplicationOk() (*string, bool)`

GetApplicationOk returns a tuple with the Application field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetApplication

`func (o *DestinationApplicationDirectory) SetApplication(v string)`

SetApplication sets Application field to given value.

### GetContext

`func (o *DestinationApplicationDirectory) GetContext() int32`

GetContext returns the Context field if non-nil, zero value otherwise.

### GetContextOk

`func (o *DestinationApplicationDirectory) GetContextOk() (*int32, bool)`

GetContextOk returns a tuple with the Context field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContext

`func (o *DestinationApplicationDirectory) SetContext(v int32)`

SetContext sets Context field to given value.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
