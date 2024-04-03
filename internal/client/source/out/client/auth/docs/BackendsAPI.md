# \BackendsAPI

All URIs are relative to *<http://api.accentvoice.io/0.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteLDAPBackendConfig**](BackendsAPI.md#DeleteLDAPBackendConfig) | **Delete** /backends/ldap | Delete current tenant&#39;s LDAP backend configuration
[**GetBackends**](BackendsAPI.md#GetBackends) | **Get** /backends | Get list of activated backends
[**GetLDAPBackendConfig**](BackendsAPI.md#GetLDAPBackendConfig) | **Get** /backends/ldap | Get current tenant&#39;s LDAP backend configuration. If there is no configuration, all the fields will be &#x60;null&#x60;.
[**UpdateLDAPBackendConfig**](BackendsAPI.md#UpdateLDAPBackendConfig) | **Put** /backends/ldap | Update current tenant&#39;s LDAP backend configuration

## DeleteLDAPBackendConfig

> DeleteLDAPBackendConfig(ctx).AccentTenant(accentTenant).Execute()

Delete current tenant's LDAP backend configuration

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.BackendsAPI.DeleteLDAPBackendConfig(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `BackendsAPI.DeleteLDAPBackendConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLDAPBackendConfigRequest struct via the builder pattern

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

## GetBackends

> BackendList GetBackends(ctx).Execute()

Get list of activated backends

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.BackendsAPI.GetBackends(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `BackendsAPI.GetBackends``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetBackends`: BackendList
 fmt.Fprintf(os.Stdout, "Response from `BackendsAPI.GetBackends`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetBackendsRequest struct via the builder pattern

### Return type

[**BackendList**](BackendList.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLDAPBackendConfig

> LDAPBackendConfig GetLDAPBackendConfig(ctx).AccentTenant(accentTenant).Execute()

Get current tenant's LDAP backend configuration. If there is no configuration, all the fields will be `null`.

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.BackendsAPI.GetLDAPBackendConfig(context.Background()).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `BackendsAPI.GetLDAPBackendConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetLDAPBackendConfig`: LDAPBackendConfig
 fmt.Fprintf(os.Stdout, "Response from `BackendsAPI.GetLDAPBackendConfig`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetLDAPBackendConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LDAPBackendConfig**](LDAPBackendConfig.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLDAPBackendConfig

> LDAPBackendConfig UpdateLDAPBackendConfig(ctx).Body(body).AccentTenant(accentTenant).Execute()

Update current tenant's LDAP backend configuration

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
 body := *openapiclient.NewLDAPBackendConfigEdit("Host_example", int32(389), "OU=people,DC=accent-voice,DC=org", "mail", "uid") // LDAPBackendConfigEdit | The LDAP backend configuration
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.BackendsAPI.UpdateLDAPBackendConfig(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `BackendsAPI.UpdateLDAPBackendConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UpdateLDAPBackendConfig`: LDAPBackendConfig
 fmt.Fprintf(os.Stdout, "Response from `BackendsAPI.UpdateLDAPBackendConfig`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLDAPBackendConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LDAPBackendConfigEdit**](LDAPBackendConfigEdit.md) | The LDAP backend configuration |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LDAPBackendConfig**](LDAPBackendConfig.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
