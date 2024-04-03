# \IaxAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateTrunkEndpointIax**](IaxAPI.md#AssociateTrunkEndpointIax) | **Put** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Associate trunk and IAX endpoint
[**AssociateTrunkRegisterIax**](IaxAPI.md#AssociateTrunkRegisterIax) | **Put** /trunks/{trunk_id}/registers/iax/{iax_id} | Associate trunk and IAX register
[**CreateEndpointIax**](IaxAPI.md#CreateEndpointIax) | **Post** /endpoints/iax | Create IAX endpoint
[**CreateRegisterIax**](IaxAPI.md#CreateRegisterIax) | **Post** /registers/iax | Create register_iax
[**DeleteEndpointIax**](IaxAPI.md#DeleteEndpointIax) | **Delete** /endpoints/iax/{iax_id} | Delete IAX Endpoint
[**DeleteRegisterIax**](IaxAPI.md#DeleteRegisterIax) | **Delete** /registers/iax/{register_iax_id} | Delete register IAX
[**DissociateTrunkEndpointIax**](IaxAPI.md#DissociateTrunkEndpointIax) | **Delete** /trunks/{trunk_id}/endpoints/iax/{iax_id} | Dissociate trunk and IAX endpoint
[**DissociateTrunkRegisterIax**](IaxAPI.md#DissociateTrunkRegisterIax) | **Delete** /trunks/{trunk_id}/registers/iax/{iax_id} | Dissociate trunk and IAX register
[**GetEndpointIax**](IaxAPI.md#GetEndpointIax) | **Get** /endpoints/iax/{iax_id} | Get IAX Endpoint
[**GetRegisterIax**](IaxAPI.md#GetRegisterIax) | **Get** /registers/iax/{register_iax_id} | Get register IAX
[**ListAsteriskIaxCallnumberlimits**](IaxAPI.md#ListAsteriskIaxCallnumberlimits) | **Get** /asterisk/iax/callnumberlimits | List IAX callnumberlimits options
[**ListEndpointsIax**](IaxAPI.md#ListEndpointsIax) | **Get** /endpoints/iax | List IAX endpoints
[**ListRegistersIax**](IaxAPI.md#ListRegistersIax) | **Get** /registers/iax | List registers iax
[**UpdateAsteriskIaxCallnumberlimits**](IaxAPI.md#UpdateAsteriskIaxCallnumberlimits) | **Put** /asterisk/iax/callnumberlimits | Update IAX callnumberlimits option
[**UpdateEndpointIax**](IaxAPI.md#UpdateEndpointIax) | **Put** /endpoints/iax/{iax_id} | Update IAX Endpoint
[**UpdateRegisterIax**](IaxAPI.md#UpdateRegisterIax) | **Put** /registers/iax/{register_iax_id} | Update register IAX

## AssociateTrunkEndpointIax

> AssociateTrunkEndpointIax(ctx, trunkId, iaxId).Execute()

Associate trunk and IAX endpoint

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
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.AssociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.AssociateTrunkEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkEndpointIaxRequest struct via the builder pattern

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

## AssociateTrunkRegisterIax

> AssociateTrunkRegisterIax(ctx, trunkId, iaxId).Execute()

Associate trunk and IAX register

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
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.AssociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.AssociateTrunkRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateTrunkRegisterIaxRequest struct via the builder pattern

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

## CreateEndpointIax

> EndpointIAX CreateEndpointIax(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create IAX endpoint

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
 body := *openapiclient.NewEndpointIAX() // EndpointIAX | IAX Endpoint to create (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.IaxAPI.CreateEndpointIax(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.CreateEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateEndpointIax`: EndpointIAX
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.CreateEndpointIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointIAX**](EndpointIAX.md) | IAX Endpoint to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointIAX**](EndpointIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateRegisterIax

> RegisterIAX CreateRegisterIax(ctx).Body(body).Execute()

Create register_iax

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
 body := *openapiclient.NewRegisterIAX("RemoteHost_example") // RegisterIAX | Register iax to create

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.IaxAPI.CreateRegisterIax(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.CreateRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateRegisterIax`: RegisterIAX
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.CreateRegisterIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRegisterIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterIAX**](RegisterIAX.md) | Register iax to create |

### Return type

[**RegisterIAX**](RegisterIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteEndpointIax

> DeleteEndpointIax(ctx, iaxId).AccentTenant(accentTenant).Execute()

Delete IAX Endpoint

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
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.DeleteEndpointIax(context.Background(), iaxId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.DeleteEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEndpointIaxRequest struct via the builder pattern

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

## DeleteRegisterIax

> DeleteRegisterIax(ctx, registerIaxId).Execute()

Delete register IAX

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
 registerIaxId := int32(56) // int32 | Register IAX's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.DeleteRegisterIax(context.Background(), registerIaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.DeleteRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerIaxId** | **int32** | Register IAX&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRegisterIaxRequest struct via the builder pattern

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

## DissociateTrunkEndpointIax

> DissociateTrunkEndpointIax(ctx, trunkId, iaxId).Execute()

Dissociate trunk and IAX endpoint

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
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.DissociateTrunkEndpointIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.DissociateTrunkEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkEndpointIaxRequest struct via the builder pattern

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

## DissociateTrunkRegisterIax

> DissociateTrunkRegisterIax(ctx, trunkId, iaxId).Execute()

Dissociate trunk and IAX register

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
 iaxId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.DissociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.DissociateTrunkRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**trunkId** | **int32** | Trunk&#39;s ID |
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateTrunkRegisterIaxRequest struct via the builder pattern

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

## GetEndpointIax

> EndpointIAX GetEndpointIax(ctx, iaxId).AccentTenant(accentTenant).Execute()

Get IAX Endpoint

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
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.IaxAPI.GetEndpointIax(context.Background(), iaxId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.GetEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetEndpointIax`: EndpointIAX
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.GetEndpointIax`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**EndpointIAX**](EndpointIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetRegisterIax

> RegisterIAX GetRegisterIax(ctx, registerIaxId).Execute()

Get register IAX

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
 registerIaxId := int32(56) // int32 | Register IAX's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.IaxAPI.GetRegisterIax(context.Background(), registerIaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.GetRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetRegisterIax`: RegisterIAX
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.GetRegisterIax`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerIaxId** | **int32** | Register IAX&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegisterIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**RegisterIAX**](RegisterIAX.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskIaxCallnumberlimits

> IAXCallNumberLimitss ListAsteriskIaxCallnumberlimits(ctx).Execute()

List IAX callnumberlimits options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.IaxAPI.ListAsteriskIaxCallnumberlimits(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.ListAsteriskIaxCallnumberlimits``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskIaxCallnumberlimits`: IAXCallNumberLimitss
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.ListAsteriskIaxCallnumberlimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskIaxCallnumberlimitsRequest struct via the builder pattern

### Return type

[**IAXCallNumberLimitss**](IAXCallNumberLimitss.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListEndpointsIax

> EndpointIAXItems ListEndpointsIax(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List IAX endpoints

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
 resp, r, err := apiClient.IaxAPI.ListEndpointsIax(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.ListEndpointsIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListEndpointsIax`: EndpointIAXItems
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.ListEndpointsIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListEndpointsIaxRequest struct via the builder pattern

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

[**EndpointIAXItems**](EndpointIAXItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListRegistersIax

> RegisterIAXItems ListRegistersIax(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List registers iax

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
 resp, r, err := apiClient.IaxAPI.ListRegistersIax(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.ListRegistersIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListRegistersIax`: RegisterIAXItems
 fmt.Fprintf(os.Stdout, "Response from `IaxAPI.ListRegistersIax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListRegistersIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**RegisterIAXItems**](RegisterIAXItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskIaxCallnumberlimits

> UpdateAsteriskIaxCallnumberlimits(ctx).Body(body).Execute()

Update IAX callnumberlimits option

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
 body := *openapiclient.NewIAXCallNumberLimitss() // IAXCallNumberLimitss | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.UpdateAsteriskIaxCallnumberlimits(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.UpdateAsteriskIaxCallnumberlimits``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskIaxCallnumberlimitsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IAXCallNumberLimitss**](IAXCallNumberLimitss.md) |  |

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

## UpdateEndpointIax

> UpdateEndpointIax(ctx, iaxId).Body(body).AccentTenant(accentTenant).Execute()

Update IAX Endpoint

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
 body := *openapiclient.NewEndpointIAX() // EndpointIAX | 
 iaxId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.UpdateEndpointIax(context.Background(), iaxId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.UpdateEndpointIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**iaxId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEndpointIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**EndpointIAX**](EndpointIAX.md) |  |

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

## UpdateRegisterIax

> UpdateRegisterIax(ctx, registerIaxId).Body(body).Execute()

Update register IAX

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
 body := *openapiclient.NewRegisterIAX("RemoteHost_example") // RegisterIAX | 
 registerIaxId := int32(56) // int32 | Register IAX's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.IaxAPI.UpdateRegisterIax(context.Background(), registerIaxId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `IaxAPI.UpdateRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**registerIaxId** | **int32** | Register IAX&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRegisterIaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RegisterIAX**](RegisterIAX.md) |  |

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
