# \SccpAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateLineEndpointSccp**](SccpAPI.md#AssociateLineEndpointSccp) | **Put** /lines/{line_id}/endpoints/sccp/{sccp_id} | Associate line and SCCP endpoint
[**CreateEndpointSccp**](SccpAPI.md#CreateEndpointSccp) | **Post** /endpoints/sccp | Create SCCP endpoint
[**DeleteEndpointSccp**](SccpAPI.md#DeleteEndpointSccp) | **Delete** /endpoints/sccp/{sccp_id} | Delete SCCP Endpoint
[**DissociateLineEndpointSccp**](SccpAPI.md#DissociateLineEndpointSccp) | **Delete** /lines/{line_id}/endpoints/sccp/{sccp_id} | Dissociate line and SCCP endpoint
[**GetEndpointSccp**](SccpAPI.md#GetEndpointSccp) | **Get** /endpoints/sccp/{sccp_id} | Get SCCP Endpoint
[**ListAsteriskSccpGeneral**](SccpAPI.md#ListAsteriskSccpGeneral) | **Get** /asterisk/sccp/general | List SCCP general options
[**ListEndpointsSccp**](SccpAPI.md#ListEndpointsSccp) | **Get** /endpoints/sccp | List SCCP endpoints
[**UpdateAsteriskSccpGeneral**](SccpAPI.md#UpdateAsteriskSccpGeneral) | **Put** /asterisk/sccp/general | Update SCCP general option
[**UpdateEndpointSccp**](SccpAPI.md#UpdateEndpointSccp) | **Put** /endpoints/sccp/{sccp_id} | Update SCCP Endpoint

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
 r, err := apiClient.SccpAPI.AssociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.AssociateLineEndpointSccp``: %v\n", err)
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

## CreateEndpointSccp

> EndpointSccp CreateEndpointSccp(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create SCCP endpoint

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
 body := *openapiclient.NewEndpointSccp() // EndpointSccp | SCCP Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SccpAPI.CreateEndpointSccp(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.CreateEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointSccp`: EndpointSccp
 fmt.Fprintf(os.Stdout, "Response from `SccpAPI.CreateEndpointSccp`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSccp**](EndpointSccp.md) | SCCP Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSccp**](EndpointSccp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteEndpointSccp

> DeleteEndpointSccp(ctx, sccpId).AccentTenant(accentTenant).Execute()

Delete SCCP Endpoint

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
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SccpAPI.DeleteEndpointSccp(context.Background(), sccpId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.DeleteEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointSccpRequest struct via the builder pattern

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
 r, err := apiClient.SccpAPI.DissociateLineEndpointSccp(context.Background(), lineId, sccpId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.DissociateLineEndpointSccp``: %v\n", err)
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

## GetEndpointSccp

> EndpointSccp GetEndpointSccp(ctx, sccpId).AccentTenant(accentTenant).Execute()

Get SCCP Endpoint

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
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SccpAPI.GetEndpointSccp(context.Background(), sccpId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.GetEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointSccp`: EndpointSccp
 fmt.Fprintf(os.Stdout, "Response from `SccpAPI.GetEndpointSccp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointSccp**](EndpointSccp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskSccpGeneral

> SCCPGeneral ListAsteriskSccpGeneral(ctx).Execute()

List SCCP general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SccpAPI.ListAsteriskSccpGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.ListAsteriskSccpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskSccpGeneral`: SCCPGeneral
 fmt.Fprintf(os.Stdout, "Response from `SccpAPI.ListAsteriskSccpGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskSccpGeneralRequest struct via the builder pattern

### Return type

[**SCCPGeneral**](SCCPGeneral.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsSccp

> EndpointSccpItems ListEndpointsSccp(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List SCCP endpoints

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
 resp, r, err := apiClient.SccpAPI.ListEndpointsSccp(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.ListEndpointsSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsSccp`: EndpointSccpItems
 fmt.Fprintf(os.Stdout, "Response from `SccpAPI.ListEndpointsSccp`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsSccpRequest struct via the builder pattern

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

[**EndpointSccpItems**](EndpointSccpItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskSccpGeneral

> UpdateAsteriskSccpGeneral(ctx).Body(body).Execute()

Update SCCP general option

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
 body := *openapiclient.NewSCCPGeneral() // SCCPGeneral | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SccpAPI.UpdateAsteriskSccpGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.UpdateAsteriskSccpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskSccpGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SCCPGeneral**](SCCPGeneral.md) |  |

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

## UpdateEndpointSccp

> UpdateEndpointSccp(ctx, sccpId).Body(body).AccentTenant(accentTenant).Execute()

Update SCCP Endpoint

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
 body := *openapiclient.NewEndpointSccp() // EndpointSccp | 
 sccpId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SccpAPI.UpdateEndpointSccp(context.Background(), sccpId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SccpAPI.UpdateEndpointSccp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**sccpId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointSccpRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointSccp**](EndpointSccp.md) |  |

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
