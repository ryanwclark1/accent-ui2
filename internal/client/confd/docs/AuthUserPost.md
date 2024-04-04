# AuthUserPost

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Auth** | Pointer to [**AuthUserPostAuth**](AuthUserPostAuth.md) |  | [optional]

## Methods

### NewAuthUserPost

`func NewAuthUserPost() *AuthUserPost`

NewAuthUserPost instantiates a new AuthUserPost object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewAuthUserPostWithDefaults

`func NewAuthUserPostWithDefaults() *AuthUserPost`

NewAuthUserPostWithDefaults instantiates a new AuthUserPost object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetAuth

`func (o *AuthUserPost) GetAuth() AuthUserPostAuth`

GetAuth returns the Auth field if non-nil, zero value otherwise.

### GetAuthOk

`func (o *AuthUserPost) GetAuthOk() (*AuthUserPostAuth, bool)`

GetAuthOk returns a tuple with the Auth field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAuth

`func (o *AuthUserPost) SetAuth(v AuthUserPostAuth)`

SetAuth sets Auth field to given value.

### HasAuth

`func (o *AuthUserPost) HasAuth() bool`

HasAuth returns a boolean if a field has been set.

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)
