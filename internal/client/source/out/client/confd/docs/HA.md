# HA

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**NodeType** | **string** | The role of this server in a HA cluster. |
**RemoteAddress** | Pointer to **string** | The remote IP address of the other machine in the HA cluster. Must be an IP address, a hostname is invalid. | [optional]

## Methods

### NewHA

`func NewHA(nodeType string, ) *HA`

NewHA instantiates a new HA object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewHAWithDefaults

`func NewHAWithDefaults() *HA`

NewHAWithDefaults instantiates a new HA object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetNodeType

`func (o *HA) GetNodeType() string`

GetNodeType returns the NodeType field if non-nil, zero value otherwise.

### GetNodeTypeOk

`func (o *HA) GetNodeTypeOk() (*string, bool)`

GetNodeTypeOk returns a tuple with the NodeType field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNodeType

`func (o *HA) SetNodeType(v string)`

SetNodeType sets NodeType field to given value.

### GetRemoteAddress

`func (o *HA) GetRemoteAddress() string`

GetRemoteAddress returns the RemoteAddress field if non-nil, zero value otherwise.

### GetRemoteAddressOk

`func (o *HA) GetRemoteAddressOk() (*string, bool)`

GetRemoteAddressOk returns a tuple with the RemoteAddress field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetRemoteAddress

`func (o *HA) SetRemoteAddress(v string)`

SetRemoteAddress sets RemoteAddress field to given value.

### HasRemoteAddress

`func (o *HA) HasRemoteAddress() bool`

HasRemoteAddress returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
