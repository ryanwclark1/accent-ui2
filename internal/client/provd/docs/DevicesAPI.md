# \DevicesAPI

All URIs are relative to *<http://api.accentvoice.io/0.2>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteDevMgrDevice**](DevicesAPI.md#DeleteDevMgrDevice) | **Delete** /dev_mgr/devices/{device_id} | Delete a device
[**DeleteDevMgrSynchronize**](DevicesAPI.md#DeleteDevMgrSynchronize) | **Delete** /dev_mgr/synchronize/{operation_id} | Delete the Operation In Progress
[**DevMgrDevicesPost**](DevicesAPI.md#DevMgrDevicesPost) | **Post** /dev_mgr/devices | Create a device
[**GetDevMgr**](DevicesAPI.md#GetDevMgr) | **Get** /dev_mgr | Get the Device Manager resource
[**GetDevMgrDevice**](DevicesAPI.md#GetDevMgrDevice) | **Get** /dev_mgr/devices/{device_id} | Get a device by ID
[**GetDevMgrDevices**](DevicesAPI.md#GetDevMgrDevices) | **Get** /dev_mgr/devices | List and find devices
[**GetDevMgrSynchronize**](DevicesAPI.md#GetDevMgrSynchronize) | **Get** /dev_mgr/synchronize/{operation_id} | Get the status of a synchronize Operation In Progress
[**PostDevMgrDhcpinfo**](DevicesAPI.md#PostDevMgrDhcpinfo) | **Post** /dev_mgr/dhcpinfo | Push DHCP request information
[**PostDevMgrReconfigure**](DevicesAPI.md#PostDevMgrReconfigure) | **Post** /dev_mgr/reconfigure | Reconfigure a device
[**PostDevMgrSynchronize**](DevicesAPI.md#PostDevMgrSynchronize) | **Post** /dev_mgr/synchronize | Synchronize a device
[**PutDevMgrDevice**](DevicesAPI.md#PutDevMgrDevice) | **Put** /dev_mgr/devices/{device_id} | Update a device

## DeleteDevMgrDevice

> DeleteDevMgrDevice(ctx, deviceId).AccentTenant(accentTenant).Execute()

Delete a device

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
 deviceId := "deviceId_example" // string | Device ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.DevicesAPI.DeleteDevMgrDevice(context.Background(), deviceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DeleteDevMgrDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDevMgrDeviceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

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

## DeleteDevMgrSynchronize

> DeleteDevMgrSynchronize(ctx, operationId).Execute()

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
 r, err := apiClient.DevicesAPI.DeleteDevMgrSynchronize(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DeleteDevMgrSynchronize``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteDevMgrSynchronizeRequest struct via the builder pattern

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

## DevMgrDevicesPost

> IdObject DevMgrDevicesPost(ctx).Device(device).AccentTenant(accentTenant).Execute()

Create a device

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
 device := *openapiclient.NewDeviceObject() // DeviceObject | Device to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.DevicesAPI.DevMgrDevicesPost(context.Background()).Device(device).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.DevMgrDevicesPost``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `DevMgrDevicesPost`: IdObject
 fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.DevMgrDevicesPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDevMgrDevicesPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **device** | [**DeviceObject**](DeviceObject.md) | Device to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

### Return type

[**IdObject**](IdObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/vnd.accent.provd+json
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDevMgr

> LinksObject GetDevMgr(ctx).Execute()

Get the Device Manager resource

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
 resp, r, err := apiClient.DevicesAPI.GetDevMgr(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevMgr``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetDevMgr`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevMgr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevMgrRequest struct via the builder pattern

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

## GetDevMgrDevice

> DeviceObject GetDevMgrDevice(ctx, deviceId).AccentTenant(accentTenant).Execute()

Get a device by ID

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
 deviceId := "deviceId_example" // string | Device ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.DevicesAPI.GetDevMgrDevice(context.Background(), deviceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevMgrDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetDevMgrDevice`: DeviceObject
 fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevMgrDevice`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevMgrDeviceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

### Return type

[**DeviceObject**](DeviceObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDevMgrDevices

> DevicesList GetDevMgrDevices(ctx).Q(q).Fields(fields).Skip(skip).Sort(sort).SortOrd(sortOrd).AccentTenant(accentTenant).Recurse(recurse).Execute()

List and find devices

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
 q := "q_example" // string | A selector, encoded in JSON, describing which entries should be returned. All entries are returned if not specified.  Example: `{\"ip\":\"10.34.1.110\"}`  (optional)
 fields := "fields_example" // string | A list of fields, separated by comma.  Example: `mac,ip`  (optional)
 skip := int32(56) // int32 | An integer specifing the number of entries to skip.  Example: 10  (optional)
 sort := "sort_example" // string | The key on which to sort the results.  Example: `id`  (optional)
 sortOrd := "sortOrd_example" // string | The order of sort (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.DevicesAPI.GetDevMgrDevices(context.Background()).Q(q).Fields(fields).Skip(skip).Sort(sort).SortOrd(sortOrd).AccentTenant(accentTenant).Recurse(recurse).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevMgrDevices``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetDevMgrDevices`: DevicesList
 fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevMgrDevices`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevMgrDevicesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | A selector, encoded in JSON, describing which entries should be returned. All entries are returned if not specified.  Example: &#x60;{\&quot;ip\&quot;:\&quot;10.34.1.110\&quot;}&#x60;  |
 **fields** | **string** | A list of fields, separated by comma.  Example: &#x60;mac,ip&#x60;  |
 **skip** | **int32** | An integer specifing the number of entries to skip.  Example: 10  |
 **sort** | **string** | The key on which to sort the results.  Example: &#x60;id&#x60;  |
 **sortOrd** | **string** | The order of sort |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]

### Return type

[**DevicesList**](DevicesList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetDevMgrSynchronize

> OperationInProgressObject GetDevMgrSynchronize(ctx, operationId).Execute()

Get the status of a synchronize Operation In Progress

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
 resp, r, err := apiClient.DevicesAPI.GetDevMgrSynchronize(context.Background(), operationId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.GetDevMgrSynchronize``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetDevMgrSynchronize`: OperationInProgressObject
 fmt.Fprintf(os.Stdout, "Response from `DevicesAPI.GetDevMgrSynchronize`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**operationId** | **string** | Operation In Progress ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetDevMgrSynchronizeRequest struct via the builder pattern

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

## PostDevMgrDhcpinfo

> PostDevMgrDhcpinfo(ctx).Body(body).AccentTenant(accentTenant).Execute()

Push DHCP request information

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
 body := *openapiclient.NewDHCPInfoObject() // DHCPInfoObject | DHCP request information (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.DevicesAPI.PostDevMgrDhcpinfo(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PostDevMgrDhcpinfo``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevMgrDhcpinfoRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DHCPInfoObject**](DHCPInfoObject.md) | DHCP request information |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

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

## PostDevMgrReconfigure

> PostDevMgrReconfigure(ctx).Body(body).AccentTenant(accentTenant).Execute()

Reconfigure a device

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
 body := *openapiclient.NewIdObject() // IdObject | Device ID body definition (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.DevicesAPI.PostDevMgrReconfigure(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PostDevMgrReconfigure``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevMgrReconfigureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Device ID body definition |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

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

## PostDevMgrSynchronize

> PostDevMgrSynchronize(ctx).Body(body).AccentTenant(accentTenant).Execute()

Synchronize a device

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
 body := *openapiclient.NewIdObject() // IdObject | Device ID body definition (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.DevicesAPI.PostDevMgrSynchronize(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PostDevMgrSynchronize``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostDevMgrSynchronizeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IdObject**](IdObject.md) | Device ID body definition |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

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

## PutDevMgrDevice

> PutDevMgrDevice(ctx, deviceId).Device(device).AccentTenant(accentTenant).Execute()

Update a device

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
 deviceId := "deviceId_example" // string | Device ID
 device := *openapiclient.NewDeviceObject() // DeviceObject | Device information to update (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.DevicesAPI.PutDevMgrDevice(context.Background(), deviceId).Device(device).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `DevicesAPI.PutDevMgrDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device ID |

### Other Parameters

Other parameters are passed through a pointer to a apiPutDevMgrDeviceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **device** | [**DeviceObject**](DeviceObject.md) | Device information to update |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

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
