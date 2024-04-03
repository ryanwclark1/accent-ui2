# \ProvdAPI

All URIs are relative to *<http://api.accentvoice.io/0.2>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetConfigure**](ProvdAPI.md#GetConfigure) | **Get** /configure | Get the general provd configuration
[**GetConfigureParamId**](ProvdAPI.md#GetConfigureParamId) | **Get** /configure/{param_id} | Get the configuration parameter value
[**GetProvd**](ProvdAPI.md#GetProvd) | **Get** / | Get the Provd Manager resource
[**PutConfigureFtpProxy**](ProvdAPI.md#PutConfigureFtpProxy) | **Put** /configure/ftp_proxy | Update the configuration&#39;s ftp_proxy
[**PutConfigureHttpProxy**](ProvdAPI.md#PutConfigureHttpProxy) | **Put** /configure/http_proxy | Update the configuration&#39;s http_proxy
[**PutConfigureHttpsProxy**](ProvdAPI.md#PutConfigureHttpsProxy) | **Put** /configure/https_proxy | Update the configuration&#39;s https_proxy
[**PutConfigureLocale**](ProvdAPI.md#PutConfigureLocale) | **Put** /configure/locale | Update the configuration&#39;s locale
[**PutConfigureNAT**](ProvdAPI.md#PutConfigureNAT) | **Put** /configure/NAT | Update the configuration&#39;s NAT
[**PutConfigureParamId**](ProvdAPI.md#PutConfigureParamId) | **Put** /configure/{param_id} | Set the value of a parameter
[**PutConfigurePluginServer**](ProvdAPI.md#PutConfigurePluginServer) | **Put** /configure/plugin_server | Update the configuration&#39;s plugin_server
[**PutConfigureProvisioningKey**](ProvdAPI.md#PutConfigureProvisioningKey) | **Put** /configure/provisioning_key | Update the tenant provisioning key

## GetConfigure

> GeneralConfigsObject GetConfigure(ctx).Execute()

Get the general provd configuration

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
 resp, r, err := apiClient.ProvdAPI.GetConfigure(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.GetConfigure``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetConfigure`: GeneralConfigsObject
 fmt.Fprintf(os.Stdout, "Response from `ProvdAPI.GetConfigure`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigureRequest struct via the builder pattern

### Return type

[**GeneralConfigsObject**](GeneralConfigsObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetConfigureParamId

> Param GetConfigureParamId(ctx, paramId).AccentTenant(accentTenant).Execute()

Get the configuration parameter value

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
 paramId := "paramId_example" // string | Configuration parameter ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ProvdAPI.GetConfigureParamId(context.Background(), paramId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.GetConfigureParamId``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetConfigureParamId`: Param
 fmt.Fprintf(os.Stdout, "Response from `ProvdAPI.GetConfigureParamId`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paramId** | **string** | Configuration parameter ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetConfigureParamIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource |

### Return type

[**Param**](Param.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetProvd

> LinksObject GetProvd(ctx).Execute()

Get the Provd Manager resource

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
 resp, r, err := apiClient.ProvdAPI.GetProvd(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.GetProvd``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetProvd`: LinksObject
 fmt.Fprintf(os.Stdout, "Response from `ProvdAPI.GetProvd`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetProvdRequest struct via the builder pattern

### Return type

[**LinksObject**](LinksObject.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/vnd.accent.provd+json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## PutConfigureFtpProxy

> PutConfigureFtpProxy(ctx).Body(body).Execute()

Update the configuration's ftp_proxy

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
 body := *openapiclient.NewFtpProxy() // FtpProxy | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureFtpProxy(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureFtpProxy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureFtpProxyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FtpProxy**](FtpProxy.md) | Configuration parameter body definition |

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

## PutConfigureHttpProxy

> PutConfigureHttpProxy(ctx).Body(body).Execute()

Update the configuration's http_proxy

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
 body := *openapiclient.NewHttpProxy() // HttpProxy | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureHttpProxy(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureHttpProxy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureHttpProxyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HttpProxy**](HttpProxy.md) | Configuration parameter body definition |

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

## PutConfigureHttpsProxy

> PutConfigureHttpsProxy(ctx).Body(body).Execute()

Update the configuration's https_proxy

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
 body := *openapiclient.NewHttpsProxy() // HttpsProxy | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureHttpsProxy(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureHttpsProxy``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureHttpsProxyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HttpsProxy**](HttpsProxy.md) | Configuration parameter body definition |

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

## PutConfigureLocale

> PutConfigureLocale(ctx).Body(body).Execute()

Update the configuration's locale

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
 body := *openapiclient.NewLocale() // Locale | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureLocale(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureLocale``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureLocaleRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Locale**](Locale.md) | Configuration parameter body definition |

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

## PutConfigureNAT

> PutConfigureNAT(ctx).Body(body).Execute()

Update the configuration's NAT

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
 body := *openapiclient.NewNat() // Nat | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureNAT(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureNAT``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureNATRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Nat**](Nat.md) | Configuration parameter body definition |

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

## PutConfigureParamId

> PutConfigureParamId(ctx, paramId).Body(body).Execute()

Set the value of a parameter

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
 paramId := "paramId_example" // string | Configuration parameter ID
 body := *openapiclient.NewParam() // Param | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureParamId(context.Background(), paramId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureParamId``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**paramId** | **string** | Configuration parameter ID |

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureParamIdRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**Param**](Param.md) | Configuration parameter body definition |

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

## PutConfigurePluginServer

> PutConfigurePluginServer(ctx).Body(body).Execute()

Update the configuration's plugin_server

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
 body := *openapiclient.NewPluginServer() // PluginServer | Configuration parameter body definition (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigurePluginServer(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigurePluginServer``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigurePluginServerRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PluginServer**](PluginServer.md) | Configuration parameter body definition |

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

## PutConfigureProvisioningKey

> PutConfigureProvisioningKey(ctx).Body(body).AccentTenant(accentTenant).Execute()

Update the tenant provisioning key

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
 body := *openapiclient.NewProvisioningKey() // ProvisioningKey | Configuration parameter body definition (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ProvdAPI.PutConfigureProvisioningKey(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ProvdAPI.PutConfigureProvisioningKey``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiPutConfigureProvisioningKeyRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ProvisioningKey**](ProvisioningKey.md) | Configuration parameter body definition |
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
