# \RegistersAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateTrunkRegisterIax**](RegistersAPI.md#AssociateTrunkRegisterIax) | **Put** /trunks/{trunk_id}/registers/iax/{iax_id} | Associate trunk and IAX register
[**CreateRegisterIax**](RegistersAPI.md#CreateRegisterIax) | **Post** /registers/iax | Create register_iax
[**DeleteRegisterIax**](RegistersAPI.md#DeleteRegisterIax) | **Delete** /registers/iax/{register_iax_id} | Delete register IAX
[**DissociateTrunkRegisterIax**](RegistersAPI.md#DissociateTrunkRegisterIax) | **Delete** /trunks/{trunk_id}/registers/iax/{iax_id} | Dissociate trunk and IAX register
[**GetRegisterIax**](RegistersAPI.md#GetRegisterIax) | **Get** /registers/iax/{register_iax_id} | Get register IAX
[**ListRegistersIax**](RegistersAPI.md#ListRegistersIax) | **Get** /registers/iax | List registers iax
[**UpdateRegisterIax**](RegistersAPI.md#UpdateRegisterIax) | **Put** /registers/iax/{register_iax_id} | Update register IAX

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
 r, err := apiClient.RegistersAPI.AssociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.AssociateTrunkRegisterIax``: %v\n", err)
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
 resp, r, err := apiClient.RegistersAPI.CreateRegisterIax(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.CreateRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateRegisterIax`: RegisterIAX
 fmt.Fprintf(os.Stdout, "Response from `RegistersAPI.CreateRegisterIax`: %v\n", resp)
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
 r, err := apiClient.RegistersAPI.DeleteRegisterIax(context.Background(), registerIaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.DeleteRegisterIax``: %v\n", err)
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
 r, err := apiClient.RegistersAPI.DissociateTrunkRegisterIax(context.Background(), trunkId, iaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.DissociateTrunkRegisterIax``: %v\n", err)
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
 resp, r, err := apiClient.RegistersAPI.GetRegisterIax(context.Background(), registerIaxId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.GetRegisterIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetRegisterIax`: RegisterIAX
 fmt.Fprintf(os.Stdout, "Response from `RegistersAPI.GetRegisterIax`: %v\n", resp)
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
 resp, r, err := apiClient.RegistersAPI.ListRegistersIax(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.ListRegistersIax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListRegistersIax`: RegisterIAXItems
 fmt.Fprintf(os.Stdout, "Response from `RegistersAPI.ListRegistersIax`: %v\n", resp)
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
 r, err := apiClient.RegistersAPI.UpdateRegisterIax(context.Background(), registerIaxId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `RegistersAPI.UpdateRegisterIax``: %v\n", err)
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
