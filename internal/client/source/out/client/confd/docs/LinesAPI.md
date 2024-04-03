# \LinesAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateLineApplication**](LinesAPI.md#AssociateLineApplication) | **Put** /lines/{line_id}/applications/{application_uuid} | Associate line and application
[**AssociateLineDevice**](LinesAPI.md#AssociateLineDevice) | **Put** /lines/{line_id}/devices/{device_id} | Associate line and device
[**AssociateLineEndpointCustom**](LinesAPI.md#AssociateLineEndpointCustom) | **Put** /lines/{line_id}/endpoints/custom/{custom_id} | Associate line and Custom endpoint
[**AssociateLineEndpointSccp**](LinesAPI.md#AssociateLineEndpointSccp) | **Put** /lines/{line_id}/endpoints/sccp/{sccp_id} | Associate line and SCCP endpoint
[**AssociateLineEndpointSip**](LinesAPI.md#AssociateLineEndpointSip) | **Put** /lines/{line_id}/endpoints/sip/{sip_uuid} | Associate line and SIP endpoint
[**AssociateLineExtension**](LinesAPI.md#AssociateLineExtension) | **Put** /lines/{line_id}/extensions/{extension_id} | Associate line and extension
[**AssociateUserLine**](LinesAPI.md#AssociateUserLine) | **Put** /users/{user_id}/lines/{line_id} | Associate user and line
[**AssociateUserLines**](LinesAPI.md#AssociateUserLines) | **Put** /users/{user_id}/lines | Associate user and lines
[**CreateLine**](LinesAPI.md#CreateLine) | **Post** /lines | Create line
[**CreateLineExtension**](LinesAPI.md#CreateLineExtension) | **Post** /lines/{line_id}/extensions | Create extension
[**DeleteLine**](LinesAPI.md#DeleteLine) | **Delete** /lines/{line_id} | Delete line
[**DissociateLineApplication**](LinesAPI.md#DissociateLineApplication) | **Delete** /lines/{line_id}/applications/{application_uuid} | Dissociate line and application
[**DissociateLineDevice**](LinesAPI.md#DissociateLineDevice) | **Delete** /lines/{line_id}/devices/{device_id} | Dissociate line and device
[**DissociateLineEndpointCustom**](LinesAPI.md#DissociateLineEndpointCustom) | **Delete** /lines/{line_id}/endpoints/custom/{custom_id} | Dissociate line and Custom endpoint
[**DissociateLineEndpointSccp**](LinesAPI.md#DissociateLineEndpointSccp) | **Delete** /lines/{line_id}/endpoints/sccp/{sccp_id} | Dissociate line and SCCP endpoint
[**DissociateLineEndpointSip**](LinesAPI.md#DissociateLineEndpointSip) | **Delete** /lines/{line_id}/endpoints/sip/{sip_uuid} | Dissociate line and SIP endpoint
[**DissociateLineExtension**](LinesAPI.md#DissociateLineExtension) | **Delete** /lines/{line_id}/extensions/{extension_id} | Dissociate line and extension
[**DissociateUserLine**](LinesAPI.md#DissociateUserLine) | **Delete** /users/{user_id}/lines/{line_id} | Dissociate user and line
[**GetDeviceLineAssociation**](LinesAPI.md#GetDeviceLineAssociation) | **Get** /devices/{device_id}/lines | List lines associated to device
[**GetLine**](LinesAPI.md#GetLine) | **Get** /lines/{line_id} | Get line
[**GetLineDevice**](LinesAPI.md#GetLineDevice) | **Get** /lines/{line_id}/devices | Get Device associated to Line
[**GetUserLineAssociatedEndpointsSip**](LinesAPI.md#GetUserLineAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/{line_id}/associated/endpoints/sip | Get SIP endpoint of a line for a user
[**GetUserLineMainAssociatedEndpointsSip**](LinesAPI.md#GetUserLineMainAssociatedEndpointsSip) | **Get** /users/{user_uuid}/lines/main/associated/endpoints/sip | Get SIP endpoint of main line for a user
[**ListLines**](LinesAPI.md#ListLines) | **Get** /lines | List lines
[**UpdateLine**](LinesAPI.md#UpdateLine) | **Put** /lines/{line_id} | Update line

## AssociateLineApplication

> AssociateLineApplication(ctx, lineId, applicationUuid).AccentTenant(accentTenant).Execute()

Associate line and application

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 applicationUuid := int32(56) // int32 | Application's UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineApplication(context.Background(), lineId, applicationUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineApplication``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**applicationUuid** | **int32** | Application&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineApplicationRequest struct via the builder pattern

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

## AssociateLineDevice

> AssociateLineDevice(ctx, lineId, deviceId).AccentTenant(accentTenant).Execute()

Associate line and device

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 deviceId := "deviceId_example" // string | Device's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineDevice(context.Background(), lineId, deviceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**deviceId** | **string** | Device&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineDeviceRequest struct via the builder pattern

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

## AssociateLineEndpointCustom

> AssociateLineEndpointCustom(ctx, lineId, customId).Execute()

Associate line and Custom endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointCustomRequest struct via the builder pattern

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

## AssociateLineEndpointSccp

> AssociateLineEndpointSccp(ctx, lineId, sccpId).Execute()

Associate line and SCCP endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 sccpId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointSccpRequest struct via the builder pattern

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

## AssociateLineEndpointSip

> AssociateLineEndpointSip(ctx, lineId, sipUuid).Execute()

Associate line and SIP endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineEndpointSipRequest struct via the builder pattern

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

## AssociateLineExtension

> AssociateLineExtension(ctx, lineId, extensionId).Execute()

Associate line and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateLineExtension(context.Background(), lineId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineExtensionRequest struct via the builder pattern

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

## AssociateUserLine

> AssociateUserLine(ctx, userId, lineId).Execute()

Associate user and line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 lineId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateUserLine(context.Background(), userId, lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateUserLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserLineRequest struct via the builder pattern

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

## AssociateUserLines

> AssociateUserLines(ctx, userId).Body(body).Execute()

Associate user and lines

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewLinesID() // LinesID | 
 userId := "userId_example" // string | the user's ID or UUID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.AssociateUserLines(context.Background(), userId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.AssociateUserLines``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateUserLinesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LinesID**](LinesID.md) |  |

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

## CreateLine

> LineView CreateLine(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewLineCreate("Context_example") // LineCreate | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.CreateLine(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.CreateLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateLine`: LineView
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.CreateLine`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLineRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**LineCreate**](LineCreate.md) |  |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LineView**](LineView.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateLineExtension

> Extension CreateLineExtension(ctx, lineId).Body(body).AccentTenant(accentTenant).Execute()

Create extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewExtension() // Extension | Extension to create
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.CreateLineExtension(context.Background(), lineId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.CreateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateLineExtension`: Extension
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.CreateLineExtension`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLineExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Extension**](Extension.md) | Extension to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Extension**](Extension.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteLine

> DeleteLine(ctx, lineId).AccentTenant(accentTenant).Execute()

Delete line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DeleteLine(context.Background(), lineId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DeleteLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteLineRequest struct via the builder pattern

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

## DissociateLineApplication

> DissociateLineApplication(ctx, lineId, applicationUuid).AccentTenant(accentTenant).Execute()

Dissociate line and application

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 applicationUuid := int32(56) // int32 | Application's UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineApplication(context.Background(), lineId, applicationUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineApplication``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**applicationUuid** | **int32** | Application&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineApplicationRequest struct via the builder pattern

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

## DissociateLineDevice

> DissociateLineDevice(ctx, lineId, deviceId).AccentTenant(accentTenant).Execute()

Dissociate line and device

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 deviceId := "deviceId_example" // string | Device's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineDevice(context.Background(), lineId, deviceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**deviceId** | **string** | Device&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineDeviceRequest struct via the builder pattern

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

## DissociateLineEndpointCustom

> DissociateLineEndpointCustom(ctx, lineId, customId).Execute()

Dissociate line and Custom endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointCustomRequest struct via the builder pattern

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

## DissociateLineEndpointSccp

> DissociateLineEndpointSccp(ctx, lineId, sccpId).Execute()

Dissociate line and SCCP endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 sccpId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointSccpRequest struct via the builder pattern

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

## DissociateLineEndpointSip

> DissociateLineEndpointSip(ctx, lineId, sipUuid).Execute()

Dissociate line and SIP endpoint

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 sipUuid := "sipUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineEndpointSip(context.Background(), lineId, sipUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineEndpointSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**sipUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineEndpointSipRequest struct via the builder pattern

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

## DissociateLineExtension

> DissociateLineExtension(ctx, lineId, extensionId).Execute()

Dissociate line and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateLineExtension(context.Background(), lineId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineExtensionRequest struct via the builder pattern

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

## DissociateUserLine

> DissociateUserLine(ctx, userId, lineId).Execute()

Dissociate user and line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userId := "userId_example" // string | the user's ID or UUID
 lineId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.DissociateUserLine(context.Background(), userId, lineId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.DissociateUserLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userId** | **string** | the user&#39;s ID or UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateUserLineRequest struct via the builder pattern

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

## GetDeviceLineAssociation

> LineDeviceItems GetDeviceLineAssociation(ctx, deviceId).AccentTenant(accentTenant).Execute()

List lines associated to device

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 deviceId := "deviceId_example" // string | Device's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.GetDeviceLineAssociation(context.Background(), deviceId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetDeviceLineAssociation``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetDeviceLineAssociation`: LineDeviceItems
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetDeviceLineAssociation`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**deviceId** | **string** | Device&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetDeviceLineAssociationRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LineDeviceItems**](LineDeviceItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLine

> LineView GetLine(ctx, lineId).AccentTenant(accentTenant).Execute()

Get line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.GetLine(context.Background(), lineId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetLine`: LineView
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetLine`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLineRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LineView**](LineView.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetLineDevice

> LineDevice GetLineDevice(ctx, lineId).AccentTenant(accentTenant).Execute()

Get Device associated to Line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.GetLineDevice(context.Background(), lineId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetLineDevice``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetLineDevice`: LineDevice
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetLineDevice`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLineDeviceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**LineDevice**](LineDevice.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineAssociatedEndpointsSip

> EndpointSIP GetUserLineAssociatedEndpointsSip(ctx, userUuid, lineId).View(view).Execute()

Get SIP endpoint of a line for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 lineId := int32(56) // int32 | 
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.GetUserLineAssociatedEndpointsSip(context.Background(), userUuid, lineId).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetUserLineAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetUserLineAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserLineMainAssociatedEndpointsSip

> EndpointSIP GetUserLineMainAssociatedEndpointsSip(ctx, userUuid).View(view).Execute()

Get SIP endpoint of main line for a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 view := "view_example" // string | Different view of the SIP endpoint  The `default` view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The `merged` view includes all options from included templates.  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.GetUserLineMainAssociatedEndpointsSip(context.Background(), userUuid).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.GetUserLineMainAssociatedEndpointsSip``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserLineMainAssociatedEndpointsSip`: EndpointSIP
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.GetUserLineMainAssociatedEndpointsSip`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserLineMainAssociatedEndpointsSipRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **view** | **string** | Different view of the SIP endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to include only options that are defined on the specified endpoint.  The &#x60;merged&#x60; view includes all options from included templates.  |

### Return type

[**EndpointSIP**](EndpointSIP.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListLines

> LineItems ListLines(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List lines

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.LinesAPI.ListLines(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.ListLines``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListLines`: LineItems
 fmt.Fprintf(os.Stdout, "Response from `LinesAPI.ListLines`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListLinesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**LineItems**](LineItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateLine

> UpdateLine(ctx, lineId).Body(body).AccentTenant(accentTenant).Execute()

Update line

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 body := *openapiclient.NewLine("Context_example") // Line | 
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.LinesAPI.UpdateLine(context.Background(), lineId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `LinesAPI.UpdateLine``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateLineRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Line**](Line.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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
