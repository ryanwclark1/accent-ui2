# \FunckeysAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateUserFuncKeyTemplate**](FunckeysAPI.md#AssociateUserFuncKeyTemplate) | **Put** /users/{user_id}/funckeys/templates/{template_id} | Associate a func key template to a user
[**CreateFuncKeyTemplate**](FunckeysAPI.md#CreateFuncKeyTemplate) | **Post** /funckeys/templates | Create a template of func keys
[**DeleteFuncKey**](FunckeysAPI.md#DeleteFuncKey) | **Delete** /funckeys/templates/{template_id}/{position} | Remove func key from template
[**DeleteFuncKeyTemplate**](FunckeysAPI.md#DeleteFuncKeyTemplate) | **Delete** /funckeys/templates/{template_id} | Delete func key template
[**DeleteUserFuncKey**](FunckeysAPI.md#DeleteUserFuncKey) | **Delete** /users/{user_id}/funckeys/{position} | Remove func key for user
[**DissociateUserFuncKeyTemplate**](FunckeysAPI.md#DissociateUserFuncKeyTemplate) | **Delete** /users/{user_id}/funckeys/templates/{template_id} | Dissociate a func key template to a user
[**GetFuncKey**](FunckeysAPI.md#GetFuncKey) | **Get** /funckeys/templates/{template_id}/{position} | Get a func key inside template
[**GetFuncKeyTemplate**](FunckeysAPI.md#GetFuncKeyTemplate) | **Get** /funckeys/templates/{template_id} | Get a func key template
[**GetUserFuncKey**](FunckeysAPI.md#GetUserFuncKey) | **Get** /users/{user_id}/funckeys/{position} | Get a func key for a user
[**ListFuncKeyDestinations**](FunckeysAPI.md#ListFuncKeyDestinations) | **Get** /funckeys/destinations | List of possible func key destinations and their parameters
[**ListFuncKeyTemplate**](FunckeysAPI.md#ListFuncKeyTemplate) | **Get** /funckeys/templates | List a func key template
[**ListFuncKeyTemplateUserAssociations**](FunckeysAPI.md#ListFuncKeyTemplateUserAssociations) | **Get** /funckeys/templates/{template_id}/users | List users associated to template
[**ListUserFuncKeyTemplateAssociations**](FunckeysAPI.md#ListUserFuncKeyTemplateAssociations) | **Get** /users/{user_id}/funckeys/templates | List funckey templates associated to user
[**ListUserFuncKeys**](FunckeysAPI.md#ListUserFuncKeys) | **Get** /users/{user_id}/funckeys | List func keys for a user
[**UpdateFuncKey**](FunckeysAPI.md#UpdateFuncKey) | **Put** /funckeys/templates/{template_id}/{position} | Add/Replace a func key in a template
[**UpdateFuncKeyTemplate**](FunckeysAPI.md#UpdateFuncKeyTemplate) | **Put** /funckeys/templates/{template_id} | Update a func key template
[**UpdateUserFuncKey**](FunckeysAPI.md#UpdateUserFuncKey) | **Put** /users/{user_id}/funckeys/{position} | Add/Replace a func key for a user
[**UpdateUserFuncKeys**](FunckeysAPI.md#UpdateUserFuncKeys) | **Put** /users/{user_id}/funckeys | Update func keys for a user

## AssociateUserFuncKeyTemplate

> AssociateUserFuncKeyTemplate(ctx, userId, templateId).AccentTenant(accentTenant).Execute()

Associate a func key template to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.AssociateUserFuncKeyTemplate(context.Background(), userId, templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.AssociateUserFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateFuncKeyTemplate

> FuncKeyTemplate CreateFuncKeyTemplate(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a template of func keys

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKeyTemplate() // FuncKeyTemplate | Template to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.CreateFuncKeyTemplate(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.CreateFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateFuncKeyTemplate`: FuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.CreateFuncKeyTemplate`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKeyTemplate**](FuncKeyTemplate.md) | Template to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKeyTemplate**](FuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteFuncKey

> DeleteFuncKey(ctx, templateId, position).AccentTenant(accentTenant).Execute()

Remove func key from template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.DeleteFuncKey(context.Background(), templateId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.DeleteFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteFuncKeyTemplate

> DeleteFuncKeyTemplate(ctx, templateId).AccentTenant(accentTenant).Execute()

Delete func key template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.DeleteFuncKeyTemplate(context.Background(), templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.DeleteFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteUserFuncKey

> DeleteUserFuncKey(ctx, userId, position).AccentTenant(accentTenant).Execute()

Remove func key for user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.DeleteUserFuncKey(context.Background(), userId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.DeleteUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DissociateUserFuncKeyTemplate

> DissociateUserFuncKeyTemplate(ctx, userId, templateId).AccentTenant(accentTenant).Execute()

Dissociate a func key template to a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.DissociateUserFuncKeyTemplate(context.Background(), userId, templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.DissociateUserFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFuncKey

> FuncKey GetFuncKey(ctx, templateId, position).AccentTenant(accentTenant).Execute()

Get a func key inside template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.GetFuncKey(context.Background(), templateId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.GetFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetFuncKey`: FuncKey
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.GetFuncKey`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKey**](FuncKey.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetFuncKeyTemplate

> FuncKeyTemplate GetFuncKeyTemplate(ctx, templateId).AccentTenant(accentTenant).Execute()

Get a func key template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.GetFuncKeyTemplate(context.Background(), templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.GetFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetFuncKeyTemplate`: FuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.GetFuncKeyTemplate`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKeyTemplate**](FuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserFuncKey

> FuncKey GetUserFuncKey(ctx, userId, position).AccentTenant(accentTenant).Execute()

Get a func key for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.GetUserFuncKey(context.Background(), userId, position).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.GetUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserFuncKey`: FuncKey
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.GetUserFuncKey`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKey**](FuncKey.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFuncKeyDestinations

> []FuncKeyDestination ListFuncKeyDestinations(ctx).Execute()

List of possible func key destinations and their parameters

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.ListFuncKeyDestinations(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.ListFuncKeyDestinations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListFuncKeyDestinations`: []FuncKeyDestination
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.ListFuncKeyDestinations`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListFuncKeyDestinationsRequest struct via the builder pattern

### Return type

[**[]FuncKeyDestination**](FuncKeyDestination.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFuncKeyTemplate

> FuncKeyTemplate ListFuncKeyTemplate(ctx).AccentTenant(accentTenant).Recurse(recurse).Execute()

List a func key template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.ListFuncKeyTemplate(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.ListFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListFuncKeyTemplate`: FuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.ListFuncKeyTemplate`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]

### Return type

[**FuncKeyTemplate**](FuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListFuncKeyTemplateUserAssociations

> UserFuncKeyTemplate ListFuncKeyTemplateUserAssociations(ctx, templateId).AccentTenant(accentTenant).Execute()

List users associated to template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.ListFuncKeyTemplateUserAssociations(context.Background(), templateId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.ListFuncKeyTemplateUserAssociations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListFuncKeyTemplateUserAssociations`: UserFuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.ListFuncKeyTemplateUserAssociations`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiListFuncKeyTemplateUserAssociationsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserFuncKeyTemplate**](UserFuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserFuncKeyTemplateAssociations

> UserFuncKeyTemplate ListUserFuncKeyTemplateAssociations(ctx, userId).AccentTenant(accentTenant).Execute()

List funckey templates associated to user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.ListUserFuncKeyTemplateAssociations(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.ListUserFuncKeyTemplateAssociations``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserFuncKeyTemplateAssociations`: UserFuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.ListUserFuncKeyTemplateAssociations`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserFuncKeyTemplateAssociationsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserFuncKeyTemplate**](UserFuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserFuncKeys

> FuncKeyTemplate ListUserFuncKeys(ctx, userId).AccentTenant(accentTenant).Execute()

List func keys for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FunckeysAPI.ListUserFuncKeys(context.Background(), userId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.ListUserFuncKeys``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserFuncKeys`: FuncKeyTemplate
 fmt.Fprintf(os.Stdout, "Response from `FunckeysAPI.ListUserFuncKeys`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserFuncKeysRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**FuncKeyTemplate**](FuncKeyTemplate.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateFuncKey

> UpdateFuncKey(ctx, templateId, position).Body(body).AccentTenant(accentTenant).Execute()

Add/Replace a func key in a template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKey() // FuncKey | 
 templateId := int32(56) // int32 | 
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.UpdateFuncKey(context.Background(), templateId, position).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.UpdateFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKey**](FuncKey.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateFuncKeyTemplate

> UpdateFuncKeyTemplate(ctx, templateId).Body(body).AccentTenant(accentTenant).Execute()

Update a func key template

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKeyTemplate() // FuncKeyTemplate | 
 templateId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.UpdateFuncKeyTemplate(context.Background(), templateId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.UpdateFuncKeyTemplate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**templateId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFuncKeyTemplateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKeyTemplate**](FuncKeyTemplate.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserFuncKey

> UpdateUserFuncKey(ctx, userId, position).Body(body).AccentTenant(accentTenant).Execute()

Add/Replace a func key for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKey() // FuncKey | 
 userId := "userId_example" // string | the user's ID or UUID
 position := int32(56) // int32 | position of the funckey
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.UpdateUserFuncKey(context.Background(), userId, position).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.UpdateUserFuncKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**position** | **int32** | position of the funckey |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserFuncKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKey**](FuncKey.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserFuncKeys

> UpdateUserFuncKeys(ctx, userId).Body(body).AccentTenant(accentTenant).Execute()

Update func keys for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewFuncKeyTemplate() // FuncKeyTemplate | 
 userId := "userId_example" // string | the user's ID or UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.FunckeysAPI.UpdateUserFuncKeys(context.Background(), userId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FunckeysAPI.UpdateUserFuncKeys``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserFuncKeysRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FuncKeyTemplate**](FuncKeyTemplate.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
