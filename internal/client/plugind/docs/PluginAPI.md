# \PluginAPI

All URIs are relative to *<http://api.accentvoice.io/0.2>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetMarket**](PluginAPI.md#GetMarket) | **Get** /market | List plugins available on the configured market
[**GetMarketPlugin**](PluginAPI.md#GetMarketPlugin) | **Get** /market/{namespace}/{name} | Fetch the information about a plugin from the market
[**GetPlugin**](PluginAPI.md#GetPlugin) | **Get** /plugins/{namespace}/{name} | Fetch the information about a plugin that has been installed
[**GetPlugins**](PluginAPI.md#GetPlugins) | **Get** /plugins | List installed plugins
[**InstallPlugin**](PluginAPI.md#InstallPlugin) | **Post** /plugins | Install a plugin
[**UninstallPlugin**](PluginAPI.md#UninstallPlugin) | **Delete** /plugins/{namespace}/{name} | Uninstall a plugin

## GetMarket

> GetMarketResult GetMarket(ctx).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).Namespace(namespace).Name(name).Installed(installed).Execute()

List plugins available on the configured market

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 search := "search_example" // string | Filter list of items, this search is applied to all fields and is not strict (eba matches Sébastien) (optional)
 namespace := "namespace_example" // string | Search by namespace (optional)
 name := "name_example" // string | Search by name (optional)
 installed := true // bool | Filter installed plugins (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginAPI.GetMarket(context.Background()).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).Namespace(namespace).Name(name).Installed(installed).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetMarket``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetMarket`: GetMarketResult
 fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetMarket`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **search** | **string** | Filter list of items, this search is applied to all fields and is not strict (eba matches Sébastien) |
 **namespace** | **string** | Search by namespace |
 **name** | **string** | Search by name |
 **installed** | **bool** | Filter installed plugins |

### Return type

[**GetMarketResult**](GetMarketResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetMarketPlugin

> MarketPluginList GetMarketPlugin(ctx, namespace, name).Execute()

Fetch the information about a plugin from the market

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {
 namespace := "namespace_example" // string | The plugin's namespace
 name := "name_example" // string | The plugin's name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginAPI.GetMarketPlugin(context.Background(), namespace, name).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetMarketPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetMarketPlugin`: MarketPluginList
 fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetMarketPlugin`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The plugin&#39;s namespace |
**name** | **string** | The plugin&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetMarketPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MarketPluginList**](MarketPluginList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPlugin

> PluginMetadata GetPlugin(ctx, namespace, name).Execute()

Fetch the information about a plugin that has been installed

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {
 namespace := "namespace_example" // string | The plugin's namespace
 name := "name_example" // string | The plugin's name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginAPI.GetPlugin(context.Background(), namespace, name).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPlugin`: PluginMetadata
 fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPlugin`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The plugin&#39;s namespace |
**name** | **string** | The plugin&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PluginMetadata**](PluginMetadata.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPlugins

> GetPluginsResult GetPlugins(ctx).Execute()

List installed plugins

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginAPI.GetPlugins(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.GetPlugins``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPlugins`: GetPluginsResult
 fmt.Fprintf(os.Stdout, "Response from `PluginAPI.GetPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPluginsRequest struct via the builder pattern

### Return type

[**GetPluginsResult**](GetPluginsResult.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## InstallPlugin

> InstallResponse InstallPlugin(ctx).Body(body).Reinstall(reinstall).Execute()

Install a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {
 body := *openapiclient.NewPluginInstallParameters("Method_example") // PluginInstallParameters | The plugins' installation parameters
 reinstall := true // bool | With this option the plugin will be reinstalled if it is already installed (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginAPI.InstallPlugin(context.Background()).Body(body).Reinstall(reinstall).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.InstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `InstallPlugin`: InstallResponse
 fmt.Fprintf(os.Stdout, "Response from `PluginAPI.InstallPlugin`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiInstallPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PluginInstallParameters**](PluginInstallParameters.md) | The plugins&#39; installation parameters |
 **reinstall** | **bool** | With this option the plugin will be reinstalled if it is already installed |

### Return type

[**InstallResponse**](InstallResponse.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UninstallPlugin

> UninstallPlugin(ctx, namespace, name).Execute()

Uninstall a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/plugind"
)

func main() {
 namespace := "namespace_example" // string | The plugin's namespace
 name := "name_example" // string | The plugin's name

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginAPI.UninstallPlugin(context.Background(), namespace, name).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginAPI.UninstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**namespace** | **string** | The plugin&#39;s namespace |
**name** | **string** | The plugin&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiUninstallPluginRequest struct via the builder pattern

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
