# \ExternalAppsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateExternalApp**](ExternalAppsAPI.md#CreateExternalApp) | **Post** /external/apps/{app_name} | Create external app
[**CreateUserExternalApp**](ExternalAppsAPI.md#CreateUserExternalApp) | **Post** /users/{user_uuid}/external/apps/{app_name} | Create user external app
[**DeleteExternalApp**](ExternalAppsAPI.md#DeleteExternalApp) | **Delete** /external/apps/{app_name} | Delete external app
[**DeleteUserExternalApp**](ExternalAppsAPI.md#DeleteUserExternalApp) | **Delete** /users/{user_uuid}/external/apps/{app_name} | Delete user external app
[**GetExternalApp**](ExternalAppsAPI.md#GetExternalApp) | **Get** /external/apps/{app_name} | Get external app
[**GetUserExternalApp**](ExternalAppsAPI.md#GetUserExternalApp) | **Get** /users/{user_uuid}/external/apps/{app_name} | Get user external app
[**ListExternalApps**](ExternalAppsAPI.md#ListExternalApps) | **Get** /external/apps | List external apps
[**ListUserExternalApps**](ExternalAppsAPI.md#ListUserExternalApps) | **Get** /users/{user_uuid}/external/apps | List user external apps
[**UpdateExternalApp**](ExternalAppsAPI.md#UpdateExternalApp) | **Put** /external/apps/{app_name} | Update external app
[**UpdateUserExternalApp**](ExternalAppsAPI.md#UpdateUserExternalApp) | **Put** /users/{user_uuid}/external/apps/{app_name} | Update user external app

## CreateExternalApp

> ExternalApp CreateExternalApp(ctx, appName).Body(body).AccentTenant(accentTenant).Execute()

Create external app

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
 body := *openapiclient.NewExternalApp() // ExternalApp | External app to create
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAppsAPI.CreateExternalApp(context.Background(), appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.CreateExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateExternalApp`: ExternalApp
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.CreateExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExternalApp**](ExternalApp.md) | External app to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**ExternalApp**](ExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateUserExternalApp

> UserExternalApp CreateUserExternalApp(ctx, userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()

Create user external app

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
 body := *openapiclient.NewUserExternalApp() // UserExternalApp | External app to create
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAppsAPI.CreateUserExternalApp(context.Background(), userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.CreateUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateUserExternalApp`: UserExternalApp
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.CreateUserExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserExternalApp**](UserExternalApp.md) | External app to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**UserExternalApp**](UserExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteExternalApp

> DeleteExternalApp(ctx, appName).AccentTenant(accentTenant).Execute()

Delete external app

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
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAppsAPI.DeleteExternalApp(context.Background(), appName).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.DeleteExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExternalAppRequest struct via the builder pattern

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

## DeleteUserExternalApp

> DeleteUserExternalApp(ctx, userUuid, appName).AccentTenant(accentTenant).Execute()

Delete user external app

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
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAppsAPI.DeleteUserExternalApp(context.Background(), userUuid, appName).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.DeleteUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteUserExternalAppRequest struct via the builder pattern

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

## GetExternalApp

> ExternalApp GetExternalApp(ctx, appName).AccentTenant(accentTenant).Execute()

Get external app

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
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAppsAPI.GetExternalApp(context.Background(), appName).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.GetExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExternalApp`: ExternalApp
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.GetExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**ExternalApp**](ExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserExternalApp

> UserExternalApp GetUserExternalApp(ctx, userUuid, appName).AccentTenant(accentTenant).View(view).Execute()

Get user external app

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
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 view := "view_example" // string | Different view of the external app endpoint  The `default` view, when the argument is omitted, is to return the user value for this external app  The `fallback` view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAppsAPI.GetUserExternalApp(context.Background(), userUuid, appName).AccentTenant(accentTenant).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.GetUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserExternalApp`: UserExternalApp
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.GetUserExternalApp`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **view** | **string** | Different view of the external app endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to return the user value for this external app  The &#x60;fallback&#x60; view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  |

### Return type

[**UserExternalApp**](UserExternalApp.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListExternalApps

> ExternalAppItems ListExternalApps(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List external apps

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
 resp, r, err := apiClient.ExternalAppsAPI.ListExternalApps(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.ListExternalApps``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListExternalApps`: ExternalAppItems
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.ListExternalApps`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListExternalAppsRequest struct via the builder pattern

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

[**ExternalAppItems**](ExternalAppItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserExternalApps

> UserExternalAppItems ListUserExternalApps(ctx, userUuid).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()

List user external apps

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
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 view := "view_example" // string | Different view of the external app endpoint  The `default` view, when the argument is omitted, is to return the user value for this external app  The `fallback` view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExternalAppsAPI.ListUserExternalApps(context.Background(), userUuid).AccentTenant(accentTenant).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).View(view).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.ListUserExternalApps``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserExternalApps`: UserExternalAppItems
 fmt.Fprintf(os.Stdout, "Response from `ExternalAppsAPI.ListUserExternalApps`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserExternalAppsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **view** | **string** | Different view of the external app endpoint  The &#x60;default&#x60; view, when the argument is omitted, is to return the user value for this external app  The &#x60;fallback&#x60; view return the user value for this external app, but if not found, will fallback to the tenant configured value  **WARNING**: Using fallback view on list will disabled all pagination and search features  |

### Return type

[**UserExternalAppItems**](UserExternalAppItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateExternalApp

> UpdateExternalApp(ctx, appName).Body(body).AccentTenant(accentTenant).Execute()

Update external app

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
 body := *openapiclient.NewExternalApp() // ExternalApp | 
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAppsAPI.UpdateExternalApp(context.Background(), appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.UpdateExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExternalApp**](ExternalApp.md) |  |

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

## UpdateUserExternalApp

> UpdateUserExternalApp(ctx, userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()

Update user external app

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
 body := *openapiclient.NewUserExternalApp() // UserExternalApp | 
 userUuid := "38400000-8cf0-11bd-b23e-10b96e4ef00d" // string | the user's UUID
 appName := "appName_example" // string | External App's name
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExternalAppsAPI.UpdateUserExternalApp(context.Background(), userUuid, appName).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExternalAppsAPI.UpdateUserExternalApp``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | the user&#39;s UUID |
**appName** | **string** | External App&#39;s name |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateUserExternalAppRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserExternalApp**](UserExternalApp.md) |  |

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
