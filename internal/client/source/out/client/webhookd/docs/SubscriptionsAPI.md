# \SubscriptionsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**Create**](SubscriptionsAPI.md#Create) | **Post** /subscriptions | Subscribe to a HTTP callback (webhook)
[**Delete**](SubscriptionsAPI.md#Delete) | **Delete** /subscriptions/{subscription_uuid} | Delete a subscription
[**Edit**](SubscriptionsAPI.md#Edit) | **Put** /subscriptions/{subscription_uuid} | Edit a subscription
[**Get**](SubscriptionsAPI.md#Get) | **Get** /subscriptions/{subscription_uuid} | Get a subscription
[**GetLogs**](SubscriptionsAPI.md#GetLogs) | **Get** /subscriptions/{subscription_uuid}/logs | Get hook logs
[**GetSubscriptionsServices**](SubscriptionsAPI.md#GetSubscriptionsServices) | **Get** /subscriptions/services | Show the available subscription services
[**GetUserSubscription**](SubscriptionsAPI.md#GetUserSubscription) | **Get** /users/me/subscriptions/{subscription_uuid} | Get a user subscription
[**List**](SubscriptionsAPI.md#List) | **Get** /subscriptions | List subscriptions to HTTP callbacks
[**UserCreate**](SubscriptionsAPI.md#UserCreate) | **Post** /users/me/subscriptions | Subscribe to a HTTP callback (webhook) as a user
[**UserDelete**](SubscriptionsAPI.md#UserDelete) | **Delete** /users/me/subscriptions/{subscription_uuid} | Delete a user subscription
[**UserList**](SubscriptionsAPI.md#UserList) | **Get** /users/me/subscriptions | List subscriptions of a user to HTTP callbacks

## Create

> Create(ctx).Body(body).Execute()

Subscribe to a HTTP callback (webhook)

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
 body := *openapiclient.NewSubscriptionRequest(*openapiclient.NewHTTPServiceConfig("Method_example", "Url_example"), []string{"Events_example"}, "Name_example", "Service_example") // SubscriptionRequest | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.SubscriptionsAPI.Create(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.Create``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubscriptionRequest**](SubscriptionRequest.md) |  |

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

## Delete

> Delete(ctx, subscriptionUuid).Execute()

Delete a subscription

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
 r, err := apiClient.SubscriptionsAPI.Delete(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.Delete``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteRequest struct via the builder pattern

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

## Edit

> Subscription Edit(ctx, subscriptionUuid).Body(body).Execute()

Edit a subscription

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
 body := *openapiclient.NewSubscriptionRequest(*openapiclient.NewHTTPServiceConfig("Method_example", "Url_example"), []string{"Events_example"}, "Name_example", "Service_example") // SubscriptionRequest | 
 subscriptionUuid := "subscriptionUuid_example" // string | The UUID of the subscription

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SubscriptionsAPI.Edit(context.Background(), subscriptionUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.Edit``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `Edit`: Subscription
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.Edit`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionUuid** | **string** | The UUID of the subscription |

### Other Parameters

Other parameters are passed through a pointer to a apiEditRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SubscriptionRequest**](SubscriptionRequest.md) |  |

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

## Get

> Subscription Get(ctx, subscriptionUuid).Execute()

Get a subscription

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
 resp, r, err := apiClient.SubscriptionsAPI.Get(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.Get``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `Get`: Subscription
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.Get`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionUuid** | **string** | The UUID of the subscription |

### Other Parameters

Other parameters are passed through a pointer to a apiGetRequest struct via the builder pattern

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

## GetLogs

> SubscriptionLog GetLogs(ctx, subscriptionUuid).Execute()

Get hook logs

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
 resp, r, err := apiClient.SubscriptionsAPI.GetLogs(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetLogs``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetLogs`: SubscriptionLog
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetLogs`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionUuid** | **string** | The UUID of the subscription |

### Other Parameters

Other parameters are passed through a pointer to a apiGetLogsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**SubscriptionLog**](SubscriptionLog.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetSubscriptionsServices

> Services GetSubscriptionsServices(ctx).Execute()

Show the available subscription services

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SubscriptionsAPI.GetSubscriptionsServices(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetSubscriptionsServices``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetSubscriptionsServices`: Services
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetSubscriptionsServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetSubscriptionsServicesRequest struct via the builder pattern

### Return type

[**Services**](Services.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

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
 resp, r, err := apiClient.SubscriptionsAPI.GetUserSubscription(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.GetUserSubscription``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserSubscription`: Subscription
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.GetUserSubscription`: %v\n", resp)
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

## List

> SubscriptionList List(ctx).AccentTenant(accentTenant).Recurse(recurse).SearchMetadata(searchMetadata).Execute()

List subscriptions to HTTP callbacks

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 searchMetadata := "searchMetadata_example" // string | A search term formatted like \"key:value\" that will only match subscriptions having a metadata entry \"key=value\". May be given multiple times to filter more precisely on different metadata keys. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.SubscriptionsAPI.List(context.Background()).AccentTenant(accentTenant).Recurse(recurse).SearchMetadata(searchMetadata).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.List``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `List`: SubscriptionList
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.List`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
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
 r, err := apiClient.SubscriptionsAPI.UserCreate(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UserCreate``: %v\n", err)
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
 r, err := apiClient.SubscriptionsAPI.UserDelete(context.Background(), subscriptionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UserDelete``: %v\n", err)
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
 resp, r, err := apiClient.SubscriptionsAPI.UserList(context.Background()).SearchMetadata(searchMetadata).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionsAPI.UserList``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `UserList`: SubscriptionList
 fmt.Fprintf(os.Stdout, "Response from `SubscriptionsAPI.UserList`: %v\n", resp)
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
