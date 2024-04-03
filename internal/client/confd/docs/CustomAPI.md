# \CustomAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateLineEndpointCustom**](CustomAPI.md#AssociateLineEndpointCustom) | **Put** /lines/{line_id}/endpoints/custom/{custom_id} | Associate line and Custom endpoint
[**AssociateTrunkEndpointCustom**](CustomAPI.md#AssociateTrunkEndpointCustom) | **Put** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Associate trunk and Custom endpoint
[**CreateEndpointCustom**](CustomAPI.md#CreateEndpointCustom) | **Post** /endpoints/custom | Create Custom endpoint
[**DeleteEndpointCustom**](CustomAPI.md#DeleteEndpointCustom) | **Delete** /endpoints/custom/{custom_id} | Delete Custom Endpoint
[**DissociateLineEndpointCustom**](CustomAPI.md#DissociateLineEndpointCustom) | **Delete** /lines/{line_id}/endpoints/custom/{custom_id} | Dissociate line and Custom endpoint
[**DissociateTrunkEndpointCustom**](CustomAPI.md#DissociateTrunkEndpointCustom) | **Delete** /trunks/{trunk_id}/endpoints/custom/{custom_id} | Dissociate trunk and Custom endpoint
[**GetEndpointCustom**](CustomAPI.md#GetEndpointCustom) | **Get** /endpoints/custom/{custom_id} | Get Custom Endpoint
[**ListEndpointsCustom**](CustomAPI.md#ListEndpointsCustom) | **Get** /endpoints/custom | List Custom endpoints
[**UpdateEndpointCustom**](CustomAPI.md#UpdateEndpointCustom) | **Put** /endpoints/custom/{custom_id} | Update Custom Endpoint

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.AssociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.AssociateLineEndpointCustom``: %v\n", err)
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

## AssociateTrunkEndpointCustom

> AssociateTrunkEndpointCustom(ctx, trunkId, customId).Execute()

Associate trunk and Custom endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.AssociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.AssociateTrunkEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkEndpointCustomRequest struct via the builder pattern

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

## CreateEndpointCustom

> EndpointCustom CreateEndpointCustom(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create Custom endpoint

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
 body := *openapiclient.NewEndpointCustom() // EndpointCustom | Custom Endpoint to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CustomAPI.CreateEndpointCustom(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.CreateEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointCustom`: EndpointCustom
 fmt.Fprintf(os.Stdout, "Response from `CustomAPI.CreateEndpointCustom`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointCustom**](EndpointCustom.md) | Custom Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointCustom**](EndpointCustom.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteEndpointCustom

> DeleteEndpointCustom(ctx, customId).AccentTenant(accentTenant).Execute()

Delete Custom Endpoint

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
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.DeleteEndpointCustom(context.Background(), customId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.DeleteEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointCustomRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.DissociateLineEndpointCustom(context.Background(), lineId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.DissociateLineEndpointCustom``: %v\n", err)
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

## DissociateTrunkEndpointCustom

> DissociateTrunkEndpointCustom(ctx, trunkId, customId).Execute()

Dissociate trunk and Custom endpoint

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
 trunkId := int32(56) // int32 | Trunk's ID
 customId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.DissociateTrunkEndpointCustom(context.Background(), trunkId, customId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.DissociateTrunkEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkEndpointCustomRequest struct via the builder pattern

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

## GetEndpointCustom

> EndpointCustom GetEndpointCustom(ctx, customId).AccentTenant(accentTenant).Execute()

Get Custom Endpoint

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
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CustomAPI.GetEndpointCustom(context.Background(), customId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.GetEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointCustom`: EndpointCustom
 fmt.Fprintf(os.Stdout, "Response from `CustomAPI.GetEndpointCustom`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointCustom**](EndpointCustom.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsCustom

> EndpointCustomItems ListEndpointsCustom(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List Custom endpoints

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CustomAPI.ListEndpointsCustom(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.ListEndpointsCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsCustom`: EndpointCustomItems
 fmt.Fprintf(os.Stdout, "Response from `CustomAPI.ListEndpointsCustom`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsCustomRequest struct via the builder pattern

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

[**EndpointCustomItems**](EndpointCustomItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateEndpointCustom

> UpdateEndpointCustom(ctx, customId).Body(body).AccentTenant(accentTenant).Execute()

Update Custom Endpoint

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
 body := *openapiclient.NewEndpointCustom() // EndpointCustom | 
 customId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CustomAPI.UpdateEndpointCustom(context.Background(), customId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CustomAPI.UpdateEndpointCustom``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**customId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointCustomRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointCustom**](EndpointCustom.md) |  |

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
