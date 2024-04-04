# \PluginsAPI

All URIs are relative to *<http://api.accentvoice.io/0.2>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeletePgMgrInstall**](PluginsAPI.md#DeletePgMgrInstall) | **Delete** /pg_mgr/install/install/{operation_id} | Delete the Operation In Progress
[**DeletePgMgrInstallMonitor**](PluginsAPI.md#DeletePgMgrInstallMonitor) | **Delete** /pg_mgr/plugins/{plugin_id}/install/install/{operation_id} | Delete the Operation In Progress
[**DeletePgMgrMonitor**](PluginsAPI.md#DeletePgMgrMonitor) | **Delete** /pg_mgr/install/update/{operation_id} | Delete the Operation In Progress
[**DeletePgMgrPluginUpgradeMonitor**](PluginsAPI.md#DeletePgMgrPluginUpgradeMonitor) | **Delete** /pg_mgr/plugins/{plugin_id}/install/upgrade/{operation_id} | Delete the Operation In Progress
[**DeletePgMgrUpgradeMonitor**](PluginsAPI.md#DeletePgMgrUpgradeMonitor) | **Delete** /pg_mgr/install/upgrade/{operation_id} | Delete the Operation In Progress
[**GetPgMgr**](PluginsAPI.md#GetPgMgr) | **Get** /pg_mgr | Get the Plugin Manager resource
[**GetPgMgrInstall**](PluginsAPI.md#GetPgMgrInstall) | **Get** /pg_mgr/install | Get the installation service resources
[**GetPgMgrInstallStatus**](PluginsAPI.md#GetPgMgrInstallStatus) | **Get** /pg_mgr/install/install/{operation_id} | Get the status of a plugin installation Operation In Progress
[**GetPgMgrInstallable**](PluginsAPI.md#GetPgMgrInstallable) | **Get** /pg_mgr/install/installable | Get the installable plugins list
[**GetPgMgrInstallableList**](PluginsAPI.md#GetPgMgrInstallableList) | **Get** /pg_mgr/plugins/{plugin_id}/install/installable | Get the installable packages list
[**GetPgMgrInstalled**](PluginsAPI.md#GetPgMgrInstalled) | **Get** /pg_mgr/install/installed | Get the installed plugins list
[**GetPgMgrInstalledList**](PluginsAPI.md#GetPgMgrInstalledList) | **Get** /pg_mgr/plugins/{plugin_id}/install/installed | Get the installed packages list
[**GetPgMgrPlugin**](PluginsAPI.md#GetPgMgrPlugin) | **Get** /pg_mgr/plugins/{plugin_id} | Get the resources of a specific plugin
[**GetPgMgrPluginInfo**](PluginsAPI.md#GetPgMgrPluginInfo) | **Get** /pg_mgr/plugins/{plugin_id}/info | Get the information of a plugin
[**GetPgMgrPluginInstall**](PluginsAPI.md#GetPgMgrPluginInstall) | **Get** /pg_mgr/plugins/{plugin_id}/install | Get the package installation service resources
[**GetPgMgrPluginInstallStatus**](PluginsAPI.md#GetPgMgrPluginInstallStatus) | **Get** /pg_mgr/plugins/{plugin_id}/install/install/{operation_id} | Get the status of a package installation Operation In Progress
[**GetPgMgrPluginUpgradeStatus**](PluginsAPI.md#GetPgMgrPluginUpgradeStatus) | **Get** /pg_mgr/plugins/{plugin_id}/install/upgrade/{operation_id} | Get the status of a package upgrade Operation In Progress
[**GetPgMgrPlugins**](PluginsAPI.md#GetPgMgrPlugins) | **Get** /pg_mgr/plugins | List the installed plugins
[**GetPgMgrUpdateStatus**](PluginsAPI.md#GetPgMgrUpdateStatus) | **Get** /pg_mgr/install/update/{operation_id} | Get the status of a plugin database update Operation In Progress
[**GetPgMgrUpgradeStatus**](PluginsAPI.md#GetPgMgrUpgradeStatus) | **Get** /pg_mgr/install/upgrade/{operation_id} | Get the status of a plugin upgrade Operation In Progress
[**PostPgMgrInstallPlugin**](PluginsAPI.md#PostPgMgrInstallPlugin) | **Post** /pg_mgr/install/install | Install a plugin
[**PostPgMgrPluginInstallPlugin**](PluginsAPI.md#PostPgMgrPluginInstallPlugin) | **Post** /pg_mgr/plugins/{plugin_id}/install/install | Install a package
[**PostPgMgrPluginUninstallPlugin**](PluginsAPI.md#PostPgMgrPluginUninstallPlugin) | **Post** /pg_mgr/plugins/{plugin_id}/install/uninstall | Uninstall a package
[**PostPgMgrReload**](PluginsAPI.md#PostPgMgrReload) | **Post** /pg_mgr/reload | Reload a plugin
[**PostPgMgrUninstallPlugin**](PluginsAPI.md#PostPgMgrUninstallPlugin) | **Post** /pg_mgr/install/uninstall | Uninstall a plugin
[**PostPgMgrUpdateList**](PluginsAPI.md#PostPgMgrUpdateList) | **Post** /pg_mgr/install/update | Update the List of installable plugins
[**PostPgMgrUpgradePlugin**](PluginsAPI.md#PostPgMgrUpgradePlugin) | **Post** /pg_mgr/install/upgrade | Upgrade a plugin

## DeletePgMgrInstall

> DeletePgMgrInstall(ctx, operationId).Execute()

Delete the Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.DeletePgMgrInstall(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePgMgrInstall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgMgrInstallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePgMgrInstallMonitor

> DeletePgMgrInstallMonitor(ctx, pluginId, operationId).Execute()

Delete the Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.DeletePgMgrInstallMonitor(context.Background(), pluginId, operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePgMgrInstallMonitor``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgMgrInstallMonitorRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePgMgrMonitor

> DeletePgMgrMonitor(ctx, operationId).Execute()

Delete the Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.DeletePgMgrMonitor(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePgMgrMonitor``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgMgrMonitorRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePgMgrPluginUpgradeMonitor

> DeletePgMgrPluginUpgradeMonitor(ctx, pluginId, operationId).Execute()

Delete the Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.DeletePgMgrPluginUpgradeMonitor(context.Background(), pluginId, operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePgMgrPluginUpgradeMonitor``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgMgrPluginUpgradeMonitorRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePgMgrUpgradeMonitor

> DeletePgMgrUpgradeMonitor(ctx, operationId).Execute()

Delete the Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.DeletePgMgrUpgradeMonitor(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.DeletePgMgrUpgradeMonitor``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePgMgrUpgradeMonitorRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgr

> LinksObject GetPgMgr(ctx).Execute()

Get the Plugin Manager resource

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgr(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgr``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgr`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrRequest struct via the builder pattern

### Return type

[**LinksObject**](LinksObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, links

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstall

> LinksObject GetPgMgrInstall(ctx).Execute()

Get the installation service resources

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstall(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstall`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstall`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstallRequest struct via the builder pattern

### Return type

[**LinksObject**](LinksObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, links

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstallStatus

> OperationInProgressObject GetPgMgrInstallStatus(ctx, operationId).Execute()

Get the status of a plugin installation Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstallStatus(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstallStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstallStatus`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstallStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstallStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**OperationInProgressObject**](OperationInProgressObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstallable

> PackageList GetPgMgrInstallable(ctx).Execute()

Get the installable plugins list

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstallable(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstallable``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstallable`: PackageList
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstallable`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstallableRequest struct via the builder pattern

### Return type

[**PackageList**](PackageList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstallableList

> PackageList GetPgMgrInstallableList(ctx, pluginId).Execute()

Get the installable packages list

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstallableList(context.Background(), pluginId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstallableList``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstallableList`: PackageList
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstallableList`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstallableListRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PackageList**](PackageList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstalled

> PackageList GetPgMgrInstalled(ctx).Execute()

Get the installed plugins list

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstalled(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstalled``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstalled`: PackageList
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstalled`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstalledRequest struct via the builder pattern

### Return type

[**PackageList**](PackageList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrInstalledList

> PackageList GetPgMgrInstalledList(ctx, pluginId).Execute()

Get the installed packages list

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrInstalledList(context.Background(), pluginId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrInstalledList``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrInstalledList`: PackageList
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrInstalledList`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrInstalledListRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PackageList**](PackageList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPlugin

> LinksObject GetPgMgrPlugin(ctx, pluginId).Execute()

Get the resources of a specific plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPlugin(context.Background(), pluginId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPlugin`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPlugin`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**LinksObject**](LinksObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, links

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPluginInfo

> PluginInfo GetPgMgrPluginInfo(ctx, pluginId).Execute()

Get the information of a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPluginInfo(context.Background(), pluginId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPluginInfo``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPluginInfo`: PluginInfo
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPluginInfo`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginInfoRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**PluginInfo**](PluginInfo.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, plugin_info

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPluginInstall

> LinksObject GetPgMgrPluginInstall(ctx, pluginId).Execute()

Get the package installation service resources

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPluginInstall(context.Background(), pluginId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPluginInstall``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPluginInstall`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPluginInstall`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginInstallRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**LinksObject**](LinksObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, links

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPluginInstallStatus

> OperationInProgressObject GetPgMgrPluginInstallStatus(ctx, pluginId, operationId).Execute()

Get the status of a package installation Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPluginInstallStatus(context.Background(), pluginId, operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPluginInstallStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPluginInstallStatus`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPluginInstallStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginInstallStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**OperationInProgressObject**](OperationInProgressObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPluginUpgradeStatus

> OperationInProgressObject GetPgMgrPluginUpgradeStatus(ctx, pluginId, operationId).Execute()

Get the status of a package upgrade Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPluginUpgradeStatus(context.Background(), pluginId, operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPluginUpgradeStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPluginUpgradeStatus`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPluginUpgradeStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginUpgradeStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**OperationInProgressObject**](OperationInProgressObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrPlugins

> PluginsObject GetPgMgrPlugins(ctx).Execute()

List the installed plugins

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrPlugins(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrPlugins``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrPlugins`: PluginsObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrPlugins`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrPluginsRequest struct via the builder pattern

### Return type

[**PluginsObject**](PluginsObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json, plugins

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrUpdateStatus

> OperationInProgressObject GetPgMgrUpdateStatus(ctx, operationId).Execute()

Get the status of a plugin database update Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrUpdateStatus(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrUpdateStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrUpdateStatus`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrUpdateStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrUpdateStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**OperationInProgressObject**](OperationInProgressObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetPgMgrUpgradeStatus

> OperationInProgressObject GetPgMgrUpgradeStatus(ctx, operationId).Execute()

Get the status of a plugin upgrade Operation In Progress

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 operationId := "operationId_example" // string | Operation In Progress ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PluginsAPI.GetPgMgrUpgradeStatus(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.GetPgMgrUpgradeStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPgMgrUpgradeStatus`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `PluginsAPI.GetPgMgrUpgradeStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPgMgrUpgradeStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**OperationInProgressObject**](OperationInProgressObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrInstallPlugin

> PostPgMgrInstallPlugin(ctx).Body(body).Execute()

Install a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 body := *openapiclient.NewIdObject() // IdObject | Package ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrInstallPlugin(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrInstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrInstallPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Package ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrPluginInstallPlugin

> PostPgMgrPluginInstallPlugin(ctx, pluginId).Body(body).Execute()

Install a package

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 body := *openapiclient.NewIdObject() // IdObject | Package ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrPluginInstallPlugin(context.Background(), pluginId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrPluginInstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrPluginInstallPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdObject**](IdObject.md) | Package ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrPluginUninstallPlugin

> PostPgMgrPluginUninstallPlugin(ctx, pluginId).Body(body).Execute()

Uninstall a package

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 pluginId := "pluginId_example" // string | Plugin ID
 body := *openapiclient.NewIdObject() // IdObject | Package ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrPluginUninstallPlugin(context.Background(), pluginId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrPluginUninstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pluginId** | **string** | Plugin ID |

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrPluginUninstallPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**IdObject**](IdObject.md) | Package ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrReload

> PostPgMgrReload(ctx).Body(body).Execute()

Reload a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 body := *openapiclient.NewIdObject() // IdObject | Plugin ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrReload(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrReload``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrReloadRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Plugin ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrUninstallPlugin

> PostPgMgrUninstallPlugin(ctx).Body(body).Execute()

Uninstall a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 body := *openapiclient.NewIdObject() // IdObject | Package ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrUninstallPlugin(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrUninstallPlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrUninstallPluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Package ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrUpdateList

> PostPgMgrUpdateList(ctx).Body(body).Execute()

Update the List of installable plugins

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 body := map[string]interface{}{ ... } // map[string]interface{} | Empty object body (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrUpdateList(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrUpdateList``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrUpdateListRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Empty object body |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostPgMgrUpgradePlugin

> PostPgMgrUpgradePlugin(ctx).Body(body).Execute()

Upgrade a plugin

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/provd"
)

func main() {
 body := *openapiclient.NewIdObject() // IdObject | Package ID body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PluginsAPI.PostPgMgrUpgradePlugin(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PluginsAPI.PostPgMgrUpgradePlugin``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostPgMgrUpgradePluginRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Package ID body definition |

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
