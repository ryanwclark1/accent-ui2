# \ConfigsAPI

All URIs are relative to *<http://api.accentvoice.io/0.2>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteCfgMgrConfigsConfigId**](ConfigsAPI.md#DeleteCfgMgrConfigsConfigId) | **Delete** /cfg_mgr/configs/{config_id} | Delete a configuration
[**GetCfgMgr**](ConfigsAPI.md#GetCfgMgr) | **Get** /cfg_mgr | Get the Config Manager resource
[**GetCfgMgrConfig**](ConfigsAPI.md#GetCfgMgrConfig) | **Get** /cfg_mgr/configs/{config_id} | Get a configuration
[**GetCfgMgrConfigs**](ConfigsAPI.md#GetCfgMgrConfigs) | **Get** /cfg_mgr/configs | List and find configurations
[**GetCfgMgrRawConfig**](ConfigsAPI.md#GetCfgMgrRawConfig) | **Get** /cfg_mgr/configs/{config_id}/raw | Get a raw configuration
[**PostCfgMgrAutocreate**](ConfigsAPI.md#PostCfgMgrAutocreate) | **Post** /cfg_mgr/autocreate | Create an autocreate configuration
[**PostCfgMgrConfigs**](ConfigsAPI.md#PostCfgMgrConfigs) | **Post** /cfg_mgr/configs | Create a configuration
[**PutCfgMgrConfig**](ConfigsAPI.md#PutCfgMgrConfig) | **Put** /cfg_mgr/configs/{config_id} | Update a configuration

## DeleteCfgMgrConfigsConfigId

> DeleteCfgMgrConfigsConfigId(ctx, configId).Execute()

Delete a configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 configId := "configId_example" // string | Configuration ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConfigsAPI.DeleteCfgMgrConfigsConfigId(context.Background(), configId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.DeleteCfgMgrConfigsConfigId``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCfgMgrConfigsConfigIdRequest struct via the builder pattern

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

## GetCfgMgr

> LinksObject GetCfgMgr(ctx).Execute()

Get the Config Manager resource

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.GetCfgMgr(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetCfgMgr``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCfgMgr`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetCfgMgr`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetCfgMgrRequest struct via the builder pattern

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

## GetCfgMgrConfig

> ConfigObject GetCfgMgrConfig(ctx, configId).Execute()

Get a configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 configId := "configId_example" // string | Configuration ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.GetCfgMgrConfig(context.Background(), configId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetCfgMgrConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCfgMgrConfig`: ConfigObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetCfgMgrConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCfgMgrConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ConfigObject**](ConfigObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCfgMgrConfigs

> ConfigsObject GetCfgMgrConfigs(ctx).Q(q).Fields(fields).Skip(skip).Sort(sort).SortOrd(sortOrd).Execute()

List and find configurations

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 q := "q_example" // string | A selector, encoded in JSON, describing which entries should be returned. All entries are returned if not specified.  Example: `{\"ip\":\"10.34.1.110\"}`  (optional)
 fields := "fields_example" // string | A list of fields, separated by comma.  Example: `mac,ip`  (optional)
 skip := int32(56) // int32 | An integer specifing the number of entries to skip.  Example: 10  (optional)
 sort := "sort_example" // string | The key on which to sort the results.  Example: `id`  (optional)
 sortOrd := "sortOrd_example" // string | The order of sort (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.GetCfgMgrConfigs(context.Background()).Q(q).Fields(fields).Skip(skip).Sort(sort).SortOrd(sortOrd).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetCfgMgrConfigs``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCfgMgrConfigs`: ConfigsObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetCfgMgrConfigs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetCfgMgrConfigsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **q** | **string** | A selector, encoded in JSON, describing which entries should be returned. All entries are returned if not specified.  Example: &#x60;{\&quot;ip\&quot;:\&quot;10.34.1.110\&quot;}&#x60;  |
 **fields** | **string** | A list of fields, separated by comma.  Example: &#x60;mac,ip&#x60;  |
 **skip** | **int32** | An integer specifing the number of entries to skip.  Example: 10  |
 **sort** | **string** | The key on which to sort the results.  Example: &#x60;id&#x60;  |
 **sortOrd** | **string** | The order of sort |

### Return type

[**ConfigsObject**](ConfigsObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCfgMgrRawConfig

> RawConfigurationObject GetCfgMgrRawConfig(ctx, configId).Execute()

Get a raw configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 configId := "configId_example" // string | Configuration ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.GetCfgMgrRawConfig(context.Background(), configId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.GetCfgMgrRawConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCfgMgrRawConfig`: RawConfigurationObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.GetCfgMgrRawConfig`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCfgMgrRawConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**RawConfigurationObject**](RawConfigurationObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PostCfgMgrAutocreate

> IdObject PostCfgMgrAutocreate(ctx).Body(body).Execute()

Create an autocreate configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 body := map[string]interface{}{ ... } // map[string]interface{} | Empty object body (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.PostCfgMgrAutocreate(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.PostCfgMgrAutocreate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PostCfgMgrAutocreate`: IdObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.PostCfgMgrAutocreate`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostCfgMgrAutocreateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | **map[string]interface{}** | Empty object body |

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

## PostCfgMgrConfigs

> IdObject PostCfgMgrConfigs(ctx).Body(body).Execute()

Create a configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 body := *openapiclient.NewConfigObject() // ConfigObject | Body of a configuration parameter (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConfigsAPI.PostCfgMgrConfigs(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.PostCfgMgrConfigs``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `PostCfgMgrConfigs`: IdObject
 fmt.Fprintf(os.Stdout, "Response from `ConfigsAPI.PostCfgMgrConfigs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPostCfgMgrConfigsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfigObject**](ConfigObject.md) | Body of a configuration parameter |

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

## PutCfgMgrConfig

> PutCfgMgrConfig(ctx, configId).Body(body).Execute()

Update a configuration

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/provd"
)

func main() {
 configId := "configId_example" // string | Configuration ID
 body := *openapiclient.NewConfigObject() // ConfigObject | Body of a configuration parameter (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConfigsAPI.PutCfgMgrConfig(context.Background(), configId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConfigsAPI.PutCfgMgrConfig``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**configId** | **string** | Configuration ID |

### Other Parameters

Other parameters are passed through a pointer to a apiPutCfgMgrConfigRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**ConfigObject**](ConfigObject.md) | Body of a configuration parameter |

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
