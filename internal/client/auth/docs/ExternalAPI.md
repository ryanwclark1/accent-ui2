# \ExternalAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteExternalAuthConfig**](ExternalAPI.md#DeleteExternalAuthConfig) | **Delete** /external/{auth_type}/config | Delete the client id and client secret
[**GetExternalAuthConfig**](ExternalAPI.md#GetExternalAuthConfig) | **Get** /external/{auth_type}/config | Retrieve the client id and client secret
[**GetExternalAuthUsers**](ExternalAPI.md#GetExternalAuthUsers) | **Get** /external/{auth_type}/users | Retrieves the list of connected users to this external source
[**GetUserExternalAuth**](ExternalAPI.md#GetUserExternalAuth) | **Get** /users/{user_uuid}/external | Retrieves the list of the users external auth data
[**PostExternalAuthConfig**](ExternalAPI.md#PostExternalAuthConfig) | **Post** /external/{auth_type}/config | Add configuration for the given auth_type
[**UpdateExternalAuthConfig**](ExternalAPI.md#UpdateExternalAuthConfig) | **Put** /external/{auth_type}/config | Update configuration for the given auth_type

## DeleteExternalAuthConfig

> DeleteExternalAuthConfig(ctx, authType).AccentTenant(accentTenant).Execute()

Delete the client id and client secret

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 authType := "authType_example" // string | External auth type name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAPI.DeleteExternalAuthConfig(context.Background(), authType).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.DeleteExternalAuthConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authType** | **string** | External auth type name |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalAuthConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetExternalAuthConfig

> ExternalConfig GetExternalAuthConfig(ctx, authType).AccentTenant(accentTenant).Execute()

Retrieve the client id and client secret

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 authType := "authType_example" // string | External auth type name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAPI.GetExternalAuthConfig(context.Background(), authType).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetExternalAuthConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExternalAuthConfig`: ExternalConfig
 fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetExternalAuthConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authType** | **string** | External auth type name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAuthConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**ExternalConfig**](ExternalConfig.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetExternalAuthUsers

> ExternalAuthUserList GetExternalAuthUsers(ctx, authType).AccentTenant(accentTenant).Recurse(recurse).Limit(limit).Offset(offset).Execute()

Retrieves the list of connected users to this external source

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 authType := "authType_example" // string | External auth type name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAPI.GetExternalAuthUsers(context.Background(), authType).AccentTenant(accentTenant).Recurse(recurse).Limit(limit).Offset(offset).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetExternalAuthUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExternalAuthUsers`: ExternalAuthUserList
 fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetExternalAuthUsers`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authType** | **string** | External auth type name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAuthUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]

### Return type

[**ExternalAuthUserList**](ExternalAuthUserList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserExternalAuth

> ExternalAuthList GetUserExternalAuth(ctx, userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieves the list of the users external auth data

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 userUuid := "userUuid_example" // string | The UUID of the user
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAPI.GetUserExternalAuth(context.Background(), userUuid).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.GetUserExternalAuth``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserExternalAuth`: ExternalAuthList
 fmt.Fprintf(os.Stdout, "Response from `ExternalAPI.GetUserExternalAuth`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserExternalAuthRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**ExternalAuthList**](ExternalAuthList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostExternalAuthConfig

> PostExternalAuthConfig(ctx, authType).Config(config).AccentTenant(accentTenant).Execute()

Add configuration for the given auth_type

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 config := *openapiclient.NewExternalConfig() // ExternalConfig | JSON object holding configuration for the given authentication type
 authType := "authType_example" // string | External auth type name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAPI.PostExternalAuthConfig(context.Background(), authType).Config(config).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.PostExternalAuthConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authType** | **string** | External auth type name |

### Other Parameters

Other parameters are passed through a pointer to a apiPostExternalAuthConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | [**ExternalConfig**](ExternalConfig.md) | JSON object holding configuration for the given authentication type |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateExternalAuthConfig

> UpdateExternalAuthConfig(ctx, authType).Config(config).AccentTenant(accentTenant).Execute()

Update configuration for the given auth_type

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/auth"
)

func main() {
 config := *openapiclient.NewExternalConfig() // ExternalConfig | JSON object holding configuration for the given authentication type
 authType := "authType_example" // string | External auth type name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAPI.UpdateExternalAuthConfig(context.Background(), authType).Config(config).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAPI.UpdateExternalAuthConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**authType** | **string** | External auth type name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalAuthConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **config** | [**ExternalConfig**](ExternalConfig.md) | JSON object holding configuration for the given authentication type |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
