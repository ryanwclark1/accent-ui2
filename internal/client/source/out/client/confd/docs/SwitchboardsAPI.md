# \SwitchboardsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetSwitchboardFallback**](SwitchboardsAPI.md#GetSwitchboardFallback) | **Get** /switchboards/{switchboard_uuid}/fallbacks | List all fallbacks for switchboard
[**SwitchboardsGet**](SwitchboardsAPI.md#SwitchboardsGet) | **Get** /switchboards | List switchboards
[**SwitchboardsPost**](SwitchboardsAPI.md#SwitchboardsPost) | **Post** /switchboards | Create a switchboard
[**SwitchboardsSwitchboardUuidDelete**](SwitchboardsAPI.md#SwitchboardsSwitchboardUuidDelete) | **Delete** /switchboards/{switchboard_uuid} | Delete a switchboard
[**SwitchboardsSwitchboardUuidGet**](SwitchboardsAPI.md#SwitchboardsSwitchboardUuidGet) | **Get** /switchboards/{switchboard_uuid} | Get a switchboard
[**SwitchboardsSwitchboardUuidPut**](SwitchboardsAPI.md#SwitchboardsSwitchboardUuidPut) | **Put** /switchboards/{switchboard_uuid} | Update a switchboard
[**UpdateSwitchboardFallback**](SwitchboardsAPI.md#UpdateSwitchboardFallback) | **Put** /switchboards/{switchboard_uuid}/fallbacks | Update switchboard&#39;s fallbacks
[**UpdateSwitchboardMemberUsers**](SwitchboardsAPI.md#UpdateSwitchboardMemberUsers) | **Put** /switchboards/{switchboard_uuid}/members/users | Update switchboard and members

## GetSwitchboardFallback

> SwitchboardFallbacks GetSwitchboardFallback(ctx, switchboardUuid).AccentTenant(accentTenant).Execute()

List all fallbacks for switchboard

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
 switchboardUuid := "switchboardUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.GetSwitchboardFallback(context.Background(), switchboardUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.GetSwitchboardFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSwitchboardFallback`: SwitchboardFallbacks
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.GetSwitchboardFallback`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetSwitchboardFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**SwitchboardFallbacks**](SwitchboardFallbacks.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SwitchboardsGet

> SwitchboardsGet200Response SwitchboardsGet(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List switchboards

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
 resp, r, err := apiClient.SwitchboardsAPI.SwitchboardsGet(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.SwitchboardsGet``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SwitchboardsGet`: SwitchboardsGet200Response
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.SwitchboardsGet`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchboardsGetRequest struct via the builder pattern

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

[**SwitchboardsGet200Response**](SwitchboardsGet200Response.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SwitchboardsPost

> Switchboard SwitchboardsPost(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create a switchboard

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
 body := *openapiclient.NewSwitchboard() // Switchboard | Switchboard parameters
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.SwitchboardsPost(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.SwitchboardsPost``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SwitchboardsPost`: Switchboard
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.SwitchboardsPost`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchboardsPostRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Switchboard**](Switchboard.md) | Switchboard parameters |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Switchboard**](Switchboard.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SwitchboardsSwitchboardUuidDelete

> SwitchboardsSwitchboardUuidDelete(ctx, switchboardUuid).AccentTenant(accentTenant).Execute()

Delete a switchboard

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
 switchboardUuid := "switchboardUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SwitchboardsAPI.SwitchboardsSwitchboardUuidDelete(context.Background(), switchboardUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.SwitchboardsSwitchboardUuidDelete``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchboardsSwitchboardUuidDeleteRequest struct via the builder pattern

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

## SwitchboardsSwitchboardUuidGet

> Switchboard SwitchboardsSwitchboardUuidGet(ctx, switchboardUuid).AccentTenant(accentTenant).Execute()

Get a switchboard

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
 switchboardUuid := "switchboardUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SwitchboardsAPI.SwitchboardsSwitchboardUuidGet(context.Background(), switchboardUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.SwitchboardsSwitchboardUuidGet``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SwitchboardsSwitchboardUuidGet`: Switchboard
 fmt.Fprintf(os.Stdout, "Response from `SwitchboardsAPI.SwitchboardsSwitchboardUuidGet`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchboardsSwitchboardUuidGetRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Switchboard**](Switchboard.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SwitchboardsSwitchboardUuidPut

> SwitchboardsSwitchboardUuidPut(ctx, switchboardUuid).Body(body).AccentTenant(accentTenant).Execute()

Update a switchboard

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
 body := *openapiclient.NewSwitchboard() // Switchboard | 
 switchboardUuid := "switchboardUuid_example" // string | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SwitchboardsAPI.SwitchboardsSwitchboardUuidPut(context.Background(), switchboardUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.SwitchboardsSwitchboardUuidPut``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiSwitchboardsSwitchboardUuidPutRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Switchboard**](Switchboard.md) |  |

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

## UpdateSwitchboardFallback

> UpdateSwitchboardFallback(ctx, switchboardUuid).Body(body).AccentTenant(accentTenant).Execute()

Update switchboard's fallbacks

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
 switchboardUuid := "switchboardUuid_example" // string | 
 body := *openapiclient.NewSwitchboardFallbacks() // SwitchboardFallbacks | Fallbacks for switchboard (optional)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SwitchboardsAPI.UpdateSwitchboardFallback(context.Background(), switchboardUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.UpdateSwitchboardFallback``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchboardFallbackRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**SwitchboardFallbacks**](SwitchboardFallbacks.md) | Fallbacks for switchboard |
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

## UpdateSwitchboardMemberUsers

> UpdateSwitchboardMemberUsers(ctx, switchboardUuid).Body(body).Execute()

Update switchboard and members

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
 body := *openapiclient.NewUsersUuid() // UsersUuid | Users to associate with the switchboard
 switchboardUuid := "switchboardUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SwitchboardsAPI.UpdateSwitchboardMemberUsers(context.Background(), switchboardUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SwitchboardsAPI.UpdateSwitchboardMemberUsers``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**switchboardUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSwitchboardMemberUsersRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UsersUuid**](UsersUuid.md) | Users to associate with the switchboard |

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
