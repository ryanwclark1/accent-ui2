# \UsersAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetUserSubscription**](UsersAPI.md#GetUserSubscription) | **Get** /users/me/subscriptions/{subscription_uuid} | Get a user subscription
[**UserCreate**](UsersAPI.md#UserCreate) | **Post** /users/me/subscriptions | Subscribe to a HTTP callback (webhook) as a user
[**UserDelete**](UsersAPI.md#UserDelete) | **Delete** /users/me/subscriptions/{subscription_uuid} | Delete a user subscription
[**UserList**](UsersAPI.md#UserList) | **Get** /users/me/subscriptions | List subscriptions of a user to HTTP callbacks

## GetUserSubscription

> Subscription GetUserSubscription(ctx, subscriptionUuid).Execute()

Get a user subscription

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/webhookd"
)

func main() {
 subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the subscription

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserSubscription(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserSubscription``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserSubscription`: Subscription
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserSubscription`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionUuid** | **string** | The UUID of the subscription |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserSubscriptionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Subscription**](Subscription.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UserCreate

> UserCreate(ctx).Body(body).Execute()

Subscribe to a HTTP callback (webhook) as a user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/webhookd"
)

func main() {
 body := *openapiclient.NewUserSubscriptionRequest(*openapiclient.NewHTTPServiceConfig("Method_example", "Url_example"), []string{"Events_example"}, "Name_example", "Service_example") // UserSubscriptionRequest | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UserCreate(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserCreate``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserCreateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserSubscriptionRequest**](UserSubscriptionRequest.md) |  |

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

## UserDelete

> UserDelete(ctx, subscriptionUuid).Execute()

Delete a user subscription

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/webhookd"
)

func main() {
 subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the subscription

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.UsersAPI.UserDelete(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserDelete``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionUuid** | **string** | The UUID of the subscription |

### Other Parameters

Other parameters are passed through a pointer to a apiUserDeleteRequest struct via the builder pattern

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

## UserList

> SubscriptionList UserList(ctx).SearchMetadata(searchMetadata).Execute()

List subscriptions of a user to HTTP callbacks

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/webhookd"
)

func main() {
 searchMetadata := "searchMetadata_example" // string | A search term formatted like \"key:value\" that will only match subscriptions having a metadata entry \"key=value\". May be given multiple times to filter more precisely on different metadata keys. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.UserList(context.Background()).SearchMetadata(searchMetadata).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.UserList``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UserList`: SubscriptionList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.UserList`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUserListRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **searchMetadata** | **string** | A search term formatted like \&quot;key:value\&quot; that will only match subscriptions having a metadata entry \&quot;key&#x3D;value\&quot;. May be given multiple times to filter more precisely on different metadata keys. |

### Return type

[**SubscriptionList**](SubscriptionList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
