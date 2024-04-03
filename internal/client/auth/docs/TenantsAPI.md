# \TenantsAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateTenant**](TenantsAPI.md#CreateTenant) | **Post** /tenants | Creates a new tenant
[**DeleteTenant**](TenantsAPI.md#DeleteTenant) | **Delete** /tenants/{tenant_uuid} | Delete a tenant
[**GetTenant**](TenantsAPI.md#GetTenant) | **Get** /tenants/{tenant_uuid} | Retrieves the details of a tenant
[**GetTenants**](TenantsAPI.md#GetTenants) | **Get** /tenants | Retrieves the list of tenants
[**UpdateTenant**](TenantsAPI.md#UpdateTenant) | **Put** /tenants/{tenant_uuid} | Modify a tenant

## CreateTenant

> TenantPostResponse CreateTenant(ctx).Body(body).AccentTenant(accentTenant).Execute()

Creates a new tenant

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
 body := *openapiclient.NewTenantCreate() // TenantCreate | The tenant creation parameters (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TenantsAPI.CreateTenant(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.CreateTenant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateTenant`: TenantPostResponse
 fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.CreateTenant`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateTenantRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantCreate**](TenantCreate.md) | The tenant creation parameters |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**TenantPostResponse**](TenantPostResponse.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteTenant

> DeleteTenant(ctx, tenantUuid).Execute()

Delete a tenant

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
 tenantUuid := "tenantUuid_example" // string | The UUID of the tenant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.TenantsAPI.DeleteTenant(context.Background(), tenantUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.DeleteTenant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | The UUID of the tenant |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteTenantRequest struct via the builder pattern

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

## GetTenant

> TenantResult GetTenant(ctx, tenantUuid).Execute()

Retrieves the details of a tenant

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
 tenantUuid := "tenantUuid_example" // string | The UUID of the tenant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TenantsAPI.GetTenant(context.Background(), tenantUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetTenant`: TenantResult
 fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenant`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | The UUID of the tenant |

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**TenantResult**](TenantResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetTenants

> TenantList GetTenants(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Execute()

Retrieves the list of tenants

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
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | The limit defines the number of individual objects that are returned (optional)
 offset := int32(56) // int32 | The offset defines the offsets the start by the number specified (optional) (default to 0)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TenantsAPI.GetTenants(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.GetTenants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetTenants`: TenantList
 fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.GetTenants`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetTenantsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | The limit defines the number of individual objects that are returned |
 **offset** | **int32** | The offset defines the offsets the start by the number specified | [default to 0]
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**TenantList**](TenantList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateTenant

> TenantPostResponse UpdateTenant(ctx, tenantUuid).Body(body).Execute()

Modify a tenant

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
 body := *openapiclient.NewTenantEdit() // TenantEdit | The tenant parameters
 tenantUuid := "tenantUuid_example" // string | The UUID of the tenant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.TenantsAPI.UpdateTenant(context.Background(), tenantUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `TenantsAPI.UpdateTenant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateTenant`: TenantPostResponse
 fmt.Fprintf(os.Stdout, "Response from `TenantsAPI.UpdateTenant`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**tenantUuid** | **string** | The UUID of the tenant |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateTenantRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**TenantEdit**](TenantEdit.md) | The tenant parameters |

### Return type

[**TenantPostResponse**](TenantPostResponse.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
