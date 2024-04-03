# \TokenAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CheckTokenContext**](TokenAPI.md#CheckTokenContext) | **Head** /token/{token} | Checks if a token is valid
[**CheckTokenScopes**](TokenAPI.md#CheckTokenScopes) | **Post** /token/{token}/scopes/check | Check a token against scopes
[**CreateToken**](TokenAPI.md#CreateToken) | **Post** /token | Creates a token
[**DeleteRefreshTokens**](TokenAPI.md#DeleteRefreshTokens) | **Delete** /users/{user_uuid_or_me}/tokens/{client_id} | Delete a user&#39;s refresh token
[**GetToken**](TokenAPI.md#GetToken) | **Get** /token/{token} | Retrieves token data
[**GetTokens**](TokenAPI.md#GetTokens) | **Get** /tokens | Retrieve a list of refresh tokens that have been created on the system
[**GetUserTokens**](TokenAPI.md#GetUserTokens) | **Get** /users/{user_uuid_or_me}/tokens | Retrieve a user&#39;s refresh token list
[**RevokeToken**](TokenAPI.md#RevokeToken) | **Delete** /token/{token} | Revoke a token

## CheckTokenContext

> CheckTokenContext(ctx, token).Scope(scope).Tenant(tenant).Execute()

Checks if a token is valid

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
 token := "token_example" // string | The token to query
 scope := "scope_example" // string | The required ACL (optional)
 tenant := "tenant_example" // string | A tenant UUID to check against (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TokenAPI.CheckTokenContext(context.Background(), token).Scope(scope).Tenant(tenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.CheckTokenContext``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token to query |

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTokenContextRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | The required ACL |
 **tenant** | **string** | A tenant UUID to check against |

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

## CheckTokenScopes

> ScopeList CheckTokenScopes(ctx, token).Body(body).Execute()

Check a token against scopes

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
 body := *openapiclient.NewScopeCheckRequest([]string{"Scopes_example"}) // ScopeCheckRequest | The token scopes check parameters
 token := "token_example" // string | The token to query

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TokenAPI.CheckTokenScopes(context.Background(), token).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.CheckTokenScopes``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CheckTokenScopes`: ScopeList
 fmt.Fprintf(os.Stdout, "Response from `TokenAPI.CheckTokenScopes`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token to query |

### Other Parameters

Other parameters are passed through a pointer to a apiCheckTokenScopesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ScopeCheckRequest**](ScopeCheckRequest.md) | The token scopes check parameters |

### Return type

[**ScopeList**](ScopeList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateToken

> Token CreateToken(ctx).Body(body).AccentSessionType(accentSessionType).Execute()

Creates a token

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
 body := *openapiclient.NewCreateTokenRequest() // CreateTokenRequest | The token creation parameters (optional)
 accentSessionType := "accentSessionType_example" // string | The session type (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TokenAPI.CreateToken(context.Background()).Body(body).AccentSessionType(accentSessionType).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.CreateToken``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateToken`: Token
 fmt.Fprintf(os.Stdout, "Response from `TokenAPI.CreateToken`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTokenRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateTokenRequest**](CreateTokenRequest.md) | The token creation parameters |
 **accentSessionType** | **string** | The session type |

### Return type

[**Token**](Token.md)

### Authorization

[accent_auth_basic](../README.md#accent_auth_basic)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteRefreshTokens

> DeleteRefreshTokens(ctx, userUuidOrMe, clientId).Execute()

Delete a user's refresh token

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
 userUuidOrMe := "userUuidOrMe_example" // string | The UUID of the user or `me` to refer to the user doing the query
 clientId := "clientId_example" // string | The client_id of the refresh token to revoke

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TokenAPI.DeleteRefreshTokens(context.Background(), userUuidOrMe, clientId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.DeleteRefreshTokens``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuidOrMe** | **string** | The UUID of the user or &#x60;me&#x60; to refer to the user doing the query |
**clientId** | **string** | The client_id of the refresh token to revoke |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRefreshTokensRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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

## GetToken

> Token GetToken(ctx, token).Scope(scope).Tenant(tenant).Execute()

Retrieves token data

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
 token := "token_example" // string | The token to query
 scope := "scope_example" // string | The required ACL (optional)
 tenant := "tenant_example" // string | A tenant UUID to check against (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TokenAPI.GetToken(context.Background(), token).Scope(scope).Tenant(tenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetToken``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetToken`: Token
 fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetToken`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token to query |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokenRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **scope** | **string** | The required ACL |
 **tenant** | **string** | A tenant UUID to check against |

### Return type

[**Token**](Token.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetTokens

> RefreshTokenList GetTokens(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve a list of refresh tokens that have been created on the system

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TokenAPI.GetTokens(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetTokens``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetTokens`: RefreshTokenList
 fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetTokens`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetTokensRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**RefreshTokenList**](RefreshTokenList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserTokens

> RefreshTokenList GetUserTokens(ctx, userUuidOrMe).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Retrieve a user's refresh token list

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
 userUuidOrMe := "userUuidOrMe_example" // string | The UUID of the user or `me` to refer to the user doing the query
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TokenAPI.GetUserTokens(context.Background(), userUuidOrMe).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.GetUserTokens``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserTokens`: RefreshTokenList
 fmt.Fprintf(os.Stdout, "Response from `TokenAPI.GetUserTokens`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuidOrMe** | **string** | The UUID of the user or &#x60;me&#x60; to refer to the user doing the query |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserTokensRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**RefreshTokenList**](RefreshTokenList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RevokeToken

> RevokeToken(ctx, token).Execute()

Revoke a token

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
 token := "token_example" // string | The token to query

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TokenAPI.RevokeToken(context.Background(), token).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TokenAPI.RevokeToken``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**token** | **string** | The token to query |

### Other Parameters

Other parameters are passed through a pointer to a apiRevokeTokenRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

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
