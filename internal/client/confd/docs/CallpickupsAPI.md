# \CallpickupsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCallpickup**](CallpickupsAPI.md#CreateCallpickup) | **Post** /callpickups | Create call pickup
[**DeleteCallpickup**](CallpickupsAPI.md#DeleteCallpickup) | **Delete** /callpickups/{callpickup_id} | Delete call pickup
[**GetCallpickup**](CallpickupsAPI.md#GetCallpickup) | **Get** /callpickups/{callpickup_id} | Get call pickup
[**ListCallPickups**](CallpickupsAPI.md#ListCallPickups) | **Get** /callpickups | List call pickups
[**UpdateCallPickupInterceptorGroups**](CallpickupsAPI.md#UpdateCallPickupInterceptorGroups) | **Put** /callpickups/{callpickup_id}/interceptors/groups | Update call pickup and interceptors
[**UpdateCallPickupInterceptorUsers**](CallpickupsAPI.md#UpdateCallPickupInterceptorUsers) | **Put** /callpickups/{callpickup_id}/interceptors/users | Update call pickup and interceptors
[**UpdateCallPickupTargetGroups**](CallpickupsAPI.md#UpdateCallPickupTargetGroups) | **Put** /callpickups/{callpickup_id}/targets/groups | Update call pickup and targets
[**UpdateCallPickupTargetUsers**](CallpickupsAPI.md#UpdateCallPickupTargetUsers) | **Put** /callpickups/{callpickup_id}/targets/users | Update call pickup and targets
[**UpdateCallpickup**](CallpickupsAPI.md#UpdateCallpickup) | **Put** /callpickups/{callpickup_id} | Update call pickup

## CreateCallpickup

> CallPickup CreateCallpickup(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create call pickup

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
 body := *openapiclient.NewCallPickup("Name_example") // CallPickup | Call Pickup to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallpickupsAPI.CreateCallpickup(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.CreateCallpickup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCallpickup`: CallPickup
 fmt.Fprintf(os.Stdout, "Response from `CallpickupsAPI.CreateCallpickup`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCallpickupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallPickup**](CallPickup.md) | Call Pickup to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallPickup**](CallPickup.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCallpickup

> DeleteCallpickup(ctx, callpickupId).AccentTenant(accentTenant).Execute()

Delete call pickup

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
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.DeleteCallpickup(context.Background(), callpickupId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.DeleteCallpickup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCallpickupRequest struct via the builder pattern

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

## GetCallpickup

> CallPickup GetCallpickup(ctx, callpickupId).AccentTenant(accentTenant).Execute()

Get call pickup

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
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CallpickupsAPI.GetCallpickup(context.Background(), callpickupId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.GetCallpickup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCallpickup`: CallPickup
 fmt.Fprintf(os.Stdout, "Response from `CallpickupsAPI.GetCallpickup`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCallpickupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**CallPickup**](CallPickup.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListCallPickups

> CallPickupItems ListCallPickups(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List call pickups

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
 resp, r, err := apiClient.CallpickupsAPI.ListCallPickups(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.ListCallPickups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListCallPickups`: CallPickupItems
 fmt.Fprintf(os.Stdout, "Response from `CallpickupsAPI.ListCallPickups`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListCallPickupsRequest struct via the builder pattern

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

[**CallPickupItems**](CallPickupItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateCallPickupInterceptorGroups

> UpdateCallPickupInterceptorGroups(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and interceptors

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
 body := *openapiclient.NewGroupsID() // GroupsID | Groups to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.UpdateCallPickupInterceptorGroups(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.UpdateCallPickupInterceptorGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupInterceptorGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupsID**](GroupsID.md) | Groups to associated |

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

## UpdateCallPickupInterceptorUsers

> UpdateCallPickupInterceptorUsers(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and interceptors

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
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.UpdateCallPickupInterceptorUsers(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.UpdateCallPickupInterceptorUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupInterceptorUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

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

## UpdateCallPickupTargetGroups

> UpdateCallPickupTargetGroups(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and targets

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
 body := *openapiclient.NewGroupsID() // GroupsID | Groups to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.UpdateCallPickupTargetGroups(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.UpdateCallPickupTargetGroups``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupTargetGroupsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**GroupsID**](GroupsID.md) | Groups to associated |

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

## UpdateCallPickupTargetUsers

> UpdateCallPickupTargetUsers(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup and targets

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
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associated
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.UpdateCallPickupTargetUsers(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.UpdateCallPickupTargetUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallPickupTargetUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associated |

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

## UpdateCallpickup

> UpdateCallpickup(ctx, callpickupId).Body(body).AccentTenant(accentTenant).Execute()

Update call pickup

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
 body := *openapiclient.NewCallPickup("Name_example") // CallPickup | 
 callpickupId := int32(56) // int32 | Call Pickup's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CallpickupsAPI.UpdateCallpickup(context.Background(), callpickupId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CallpickupsAPI.UpdateCallpickup``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**callpickupId** | **int32** | Call Pickup&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateCallpickupRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CallPickup**](CallPickup.md) |  |

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
