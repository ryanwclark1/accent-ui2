# \CallfiltersAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallfilter**](CallfiltersAPI.md#CreateCallfilter) | **Post** /callfilters | Create call filter
[**DeleteCallfilter**](CallfiltersAPI.md#DeleteCallfilter) | **Delete** /callfilters/{callfilter_id} | Delete call filter
[**GetCallFilterFallback**](CallfiltersAPI.md#GetCallFilterFallback) | **Get** /callfilters/{callfilter_id}/fallbacks | List all fallbacks for call filter
[**GetCallfilter**](CallfiltersAPI.md#GetCallfilter) | **Get** /callfilters/{callfilter_id} | Get call filter
[**ListCallFilters**](CallfiltersAPI.md#ListCallFilters) | **Get** /callfilters | List call filters
[**UpdateCallFilterCallerUsers**](CallfiltersAPI.md#UpdateCallFilterCallerUsers) | **Put** /callfilters/{callfilter_id}/recipients/users | Update call filter and recipients
[**UpdateCallFilterFallback**](CallfiltersAPI.md#UpdateCallFilterFallback) | **Put** /callfilters/{callfilter_id}/fallbacks | Update call filter&#39;s fallbacks
[**UpdateCallFilterMemberUsers**](CallfiltersAPI.md#UpdateCallFilterMemberUsers) | **Put** /callfilters/{callfilter_id}/surrogates/users | Update call filter and surrogates
[**UpdateCallfilter**](CallfiltersAPI.md#UpdateCallfilter) | **Put** /callfilters/{callfilter_id} | Update call filter

## CreateCallfilter

> CallFilter CreateCallfilter(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create call filter

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
 body := *openapiclient.NewCallFilter("Name_example", "Source_example", "Strategy_example") // CallFilter | Call Filter to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallfiltersAPI.CreateCallfilter(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.CreateCallfilter``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCallfilter`: CallFilter
 fmt.Fprintf(os.Stdout, "Response from `CallfiltersAPI.CreateCallfilter`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallfilterRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallFilter**](CallFilter.md) | Call Filter to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallFilter**](CallFilter.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCallfilter

> DeleteCallfilter(ctx, callfilterId).AccentTenant(accentTenant).Execute()

Delete call filter

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
 callfilterId := int32(56) // int32 | Call Filter's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallfiltersAPI.DeleteCallfilter(context.Background(), callfilterId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.DeleteCallfilter``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallfilterRequest struct via the builder pattern

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

## GetCallFilterFallback

> CallFilterFallbacks GetCallFilterFallback(ctx, callfilterId).Execute()

List all fallbacks for call filter

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
 callfilterId := int32(56) // int32 | Call Filter's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallfiltersAPI.GetCallFilterFallback(context.Background(), callfilterId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.GetCallFilterFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCallFilterFallback`: CallFilterFallbacks
 fmt.Fprintf(os.Stdout, "Response from `CallfiltersAPI.GetCallFilterFallback`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallFilterFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**CallFilterFallbacks**](CallFilterFallbacks.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCallfilter

> CallFilter GetCallfilter(ctx, callfilterId).AccentTenant(accentTenant).Execute()

Get call filter

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
 callfilterId := int32(56) // int32 | Call Filter's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallfiltersAPI.GetCallfilter(context.Background(), callfilterId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.GetCallfilter``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCallfilter`: CallFilter
 fmt.Fprintf(os.Stdout, "Response from `CallfiltersAPI.GetCallfilter`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallfilterRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallFilter**](CallFilter.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCallFilters

> CallFilterItems ListCallFilters(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List call filters

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
 resp, r, err := apiClient.CallfiltersAPI.ListCallFilters(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.ListCallFilters``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListCallFilters`: CallFilterItems
 fmt.Fprintf(os.Stdout, "Response from `CallfiltersAPI.ListCallFilters`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCallFiltersRequest struct via the builder pattern

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

[**CallFilterItems**](CallFilterItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallFilterCallerUsers

> UpdateCallFilterCallerUsers(ctx, callfilterId).Body(body).Execute()

Update call filter and recipients

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
 body := *openapiclient.NewCallFilterRecipientUsersUuid() // CallFilterRecipientUsersUuid | Users to associated
 callfilterId := int32(56) // int32 | Call Filter's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallfiltersAPI.UpdateCallFilterCallerUsers(context.Background(), callfilterId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.UpdateCallFilterCallerUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFilterCallerUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallFilterRecipientUsersUuid**](CallFilterRecipientUsersUuid.md) | Users to associated |

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

## UpdateCallFilterFallback

> UpdateCallFilterFallback(ctx, callfilterId).Body(body).Execute()

Update call filter's fallbacks

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
 callfilterId := int32(56) // int32 | Call Filter's ID
 body := *openapiclient.NewCallFilterFallbacks() // CallFilterFallbacks | Fallbacks for call filter (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallfiltersAPI.UpdateCallFilterFallback(context.Background(), callfilterId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.UpdateCallFilterFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFilterFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**CallFilterFallbacks**](CallFilterFallbacks.md) | Fallbacks for call filter |

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

## UpdateCallFilterMemberUsers

> UpdateCallFilterMemberUsers(ctx, callfilterId).Body(body).Execute()

Update call filter and surrogates

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
 callfilterId := int32(56) // int32 | Call Filter's ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallfiltersAPI.UpdateCallFilterMemberUsers(context.Background(), callfilterId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.UpdateCallFilterMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallFilterMemberUsersRequest struct via the builder pattern

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

## UpdateCallfilter

> UpdateCallfilter(ctx, callfilterId).Body(body).AccentTenant(accentTenant).Execute()

Update call filter

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
 body := *openapiclient.NewCallFilter("Name_example", "Source_example", "Strategy_example") // CallFilter | 
 callfilterId := int32(56) // int32 | Call Filter's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallfiltersAPI.UpdateCallfilter(context.Background(), callfilterId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallfiltersAPI.UpdateCallfilter``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callfilterId** | **int32** | Call Filter&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallfilterRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallFilter**](CallFilter.md) |  |

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
