# \PresencesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserPresence**](PresencesAPI.md#GetUserPresence) | **Get** /users/{user_uuid}/presences | Get user presence
[**ListPresences**](PresencesAPI.md#ListPresences) | **Get** /users/presences | List presences
[**UpdateUserPresence**](PresencesAPI.md#UpdateUserPresence) | **Put** /users/{user_uuid}/presences | Update user presence

## GetUserPresence

> Presence GetUserPresence(ctx, userUuid).AccentTenant(accentTenant).Execute()

Get user presence

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 userUuid := "userUuid_example" // string | The UUID of the user
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PresencesAPI.GetUserPresence(context.Background(), userUuid).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PresencesAPI.GetUserPresence``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserPresence`: Presence
 fmt.Fprintf(os.Stdout, "Response from `PresencesAPI.GetUserPresence`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserPresenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Presence**](Presence.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListPresences

> PresenceList ListPresences(ctx).AccentTenant(accentTenant).Recurse(recurse).UserUuid(userUuid).Execute()

List presences

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 userUuid := []string{"Inner_example"} // []string | Filter by user_uuid. Many uuid can be specified. A logical AND is used for filtering. Each uuid MUST be separated by a comma (,). (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.PresencesAPI.ListPresences(context.Background()).AccentTenant(accentTenant).Recurse(recurse).UserUuid(userUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PresencesAPI.ListPresences``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListPresences`: PresenceList
 fmt.Fprintf(os.Stdout, "Response from `PresencesAPI.ListPresences`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListPresencesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **userUuid** | **[]string** | Filter by user_uuid. Many uuid can be specified. A logical AND is used for filtering. Each uuid MUST be separated by a comma (,). |

### Return type

[**PresenceList**](PresenceList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateUserPresence

> UpdateUserPresence(ctx, userUuid).Body(body).AccentTenant(accentTenant).Execute()

Update user presence

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 body := *openapiclient.NewPresence("State_example") // Presence | 
 userUuid := "userUuid_example" // string | The UUID of the user
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.PresencesAPI.UpdateUserPresence(context.Background(), userUuid).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `PresencesAPI.UpdateUserPresence``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | The UUID of the user |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserPresenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Presence**](Presence.md) |  |

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
