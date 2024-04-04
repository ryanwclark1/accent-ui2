# \RegistrarsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRegistrar**](RegistrarsAPI.md#CreateRegistrar) | **Post** /registrars | Create registrar
[**DeleteRegistrar**](RegistrarsAPI.md#DeleteRegistrar) | **Delete** /registrars/{registrar_id} | Delete registrar
[**GetRegistrar**](RegistrarsAPI.md#GetRegistrar) | **Get** /registrars/{registrar_id} | Get registrar
[**GetRegistrars**](RegistrarsAPI.md#GetRegistrars) | **Get** /registrars | Get registrars
[**UpdateRegistrar**](RegistrarsAPI.md#UpdateRegistrar) | **Put** /registrars/{registrar_id} | Update registrar

## CreateRegistrar

> Registrar CreateRegistrar(ctx).Body(body).Execute()

Create registrar

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewRegistrar("MainHost_example", "ProxyMainHost_example") // Registrar | Registrar to create

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.RegistrarsAPI.CreateRegistrar(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistrarsAPI.CreateRegistrar``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateRegistrar`: Registrar
 fmt.Fprintf(os.Stdout, "Response from `RegistrarsAPI.CreateRegistrar`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegistrarRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Registrar**](Registrar.md) | Registrar to create |

### Return type

[**Registrar**](Registrar.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteRegistrar

> DeleteRegistrar(ctx, registrarId).Execute()

Delete registrar

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 registrarId := "registrarId_example" // string | Registrar ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.RegistrarsAPI.DeleteRegistrar(context.Background(), registrarId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistrarsAPI.DeleteRegistrar``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrarId** | **string** | Registrar ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegistrarRequest struct via the builder pattern

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

## GetRegistrar

> Registrar GetRegistrar(ctx, registrarId).Execute()

Get registrar

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 registrarId := "registrarId_example" // string | Registrar ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.RegistrarsAPI.GetRegistrar(context.Background(), registrarId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistrarsAPI.GetRegistrar``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetRegistrar`: Registrar
 fmt.Fprintf(os.Stdout, "Response from `RegistrarsAPI.GetRegistrar`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrarId** | **string** | Registrar ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrarRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Registrar**](Registrar.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetRegistrars

> RegistrarItems GetRegistrars(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

Get registrars

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.RegistrarsAPI.GetRegistrars(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistrarsAPI.GetRegistrars``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetRegistrars`: RegistrarItems
 fmt.Fprintf(os.Stdout, "Response from `RegistrarsAPI.GetRegistrars`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegistrarsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**RegistrarItems**](RegistrarItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateRegistrar

> UpdateRegistrar(ctx, registrarId).Body(body).Execute()

Update registrar

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewRegistrar("MainHost_example", "ProxyMainHost_example") // Registrar | Registrar body
 registrarId := "registrarId_example" // string | Registrar ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.RegistrarsAPI.UpdateRegistrar(context.Background(), registrarId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistrarsAPI.UpdateRegistrar``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registrarId** | **string** | Registrar ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegistrarRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Registrar**](Registrar.md) | Registrar body |

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
