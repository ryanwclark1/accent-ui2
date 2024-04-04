# \PagingsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreatePaging**](PagingsAPI.md#CreatePaging) | **Post** /pagings | Create paging
[**DeletePaging**](PagingsAPI.md#DeletePaging) | **Delete** /pagings/{paging_id} | Delete paging
[**GetPaging**](PagingsAPI.md#GetPaging) | **Get** /pagings/{paging_id} | Get paging
[**ListPagings**](PagingsAPI.md#ListPagings) | **Get** /pagings | List paging
[**UpdatePaging**](PagingsAPI.md#UpdatePaging) | **Put** /pagings/{paging_id} | Update paging
[**UpdatePagingCallerUsers**](PagingsAPI.md#UpdatePagingCallerUsers) | **Put** /pagings/{paging_id}/callers/users | Update paging and callers
[**UpdatePagingMemberUsers**](PagingsAPI.md#UpdatePagingMemberUsers) | **Put** /pagings/{paging_id}/members/users | Update paging and members

## CreatePaging

> Paging CreatePaging(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create paging

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
 body := *openapiclient.NewPaging("Number_example") // Paging | Paging to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PagingsAPI.CreatePaging(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.CreatePaging``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreatePaging`: Paging
 fmt.Fprintf(os.Stdout, "Response from `PagingsAPI.CreatePaging`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreatePagingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Paging**](Paging.md) | Paging to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Paging**](Paging.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeletePaging

> DeletePaging(ctx, pagingId).AccentTenant(accentTenant).Execute()

Delete paging

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
 pagingId := int32(56) // int32 | Paging's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PagingsAPI.DeletePaging(context.Background(), pagingId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.DeletePaging``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePagingRequest struct via the builder pattern

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

## GetPaging

> Paging GetPaging(ctx, pagingId).AccentTenant(accentTenant).Execute()

Get paging

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
 pagingId := int32(56) // int32 | Paging's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PagingsAPI.GetPaging(context.Background(), pagingId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.GetPaging``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetPaging`: Paging
 fmt.Fprintf(os.Stdout, "Response from `PagingsAPI.GetPaging`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetPagingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Paging**](Paging.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListPagings

> PagingItems ListPagings(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List paging

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
 resp, r, err := apiClient.PagingsAPI.ListPagings(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.ListPagings``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListPagings`: PagingItems
 fmt.Fprintf(os.Stdout, "Response from `PagingsAPI.ListPagings`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListPagingsRequest struct via the builder pattern

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

[**PagingItems**](PagingItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdatePaging

> UpdatePaging(ctx, pagingId).Body(body).AccentTenant(accentTenant).Execute()

Update paging

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
 body := *openapiclient.NewPaging("Number_example") // Paging | 
 pagingId := int32(56) // int32 | Paging's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PagingsAPI.UpdatePaging(context.Background(), pagingId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.UpdatePaging``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagingRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Paging**](Paging.md) |  |

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

## UpdatePagingCallerUsers

> UpdatePagingCallerUsers(ctx, pagingId).Body(body).Execute()

Update paging and callers

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
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 pagingId := int32(56) // int32 | Paging's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PagingsAPI.UpdatePagingCallerUsers(context.Background(), pagingId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.UpdatePagingCallerUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagingCallerUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

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

## UpdatePagingMemberUsers

> UpdatePagingMemberUsers(ctx, pagingId).Body(body).Execute()

Update paging and members

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
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 pagingId := int32(56) // int32 | Paging's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PagingsAPI.UpdatePagingMemberUsers(context.Background(), pagingId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PagingsAPI.UpdatePagingMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**pagingId** | **int32** | Paging&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePagingMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

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
