# \AccessFeaturesAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateAccessFeature**](AccessFeaturesAPI.md#CreateAccessFeature) | **Post** /access_features | Create access_feature
[**DeleteAccessFeature**](AccessFeaturesAPI.md#DeleteAccessFeature) | **Delete** /access_features/{access_feature_id} | Delete access feature
[**GetAccessFeature**](AccessFeaturesAPI.md#GetAccessFeature) | **Get** /access_features/{access_feature_id} | Get access_feature
[**ListAccessFeatures**](AccessFeaturesAPI.md#ListAccessFeatures) | **Get** /access_features | List access features
[**UpdateAccessFeature**](AccessFeaturesAPI.md#UpdateAccessFeature) | **Put** /access_features/{access_feature_id} | Update access_feature

## CreateAccessFeature

> AccessFeature CreateAccessFeature(ctx).Body(body).Execute()

Create access_feature

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
 body := *openapiclient.NewAccessFeature("Feature_example", "Host_example") // AccessFeature | Access feature to create

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AccessFeaturesAPI.CreateAccessFeature(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AccessFeaturesAPI.CreateAccessFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateAccessFeature`: AccessFeature
 fmt.Fprintf(os.Stdout, "Response from `AccessFeaturesAPI.CreateAccessFeature`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAccessFeatureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessFeature**](AccessFeature.md) | Access feature to create |

### Return type

[**AccessFeature**](AccessFeature.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteAccessFeature

> DeleteAccessFeature(ctx, accessFeatureId).Execute()

Delete access feature

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
 accessFeatureId := int32(56) // int32 | Access feature ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AccessFeaturesAPI.DeleteAccessFeature(context.Background(), accessFeatureId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AccessFeaturesAPI.DeleteAccessFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessFeatureId** | **int32** | Access feature ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAccessFeatureRequest struct via the builder pattern

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

## GetAccessFeature

> AccessFeature GetAccessFeature(ctx, accessFeatureId).Execute()

Get access_feature

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
 accessFeatureId := int32(56) // int32 | Access feature ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AccessFeaturesAPI.GetAccessFeature(context.Background(), accessFeatureId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AccessFeaturesAPI.GetAccessFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetAccessFeature`: AccessFeature
 fmt.Fprintf(os.Stdout, "Response from `AccessFeaturesAPI.GetAccessFeature`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessFeatureId** | **int32** | Access feature ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetAccessFeatureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**AccessFeature**](AccessFeature.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAccessFeatures

> AccessFeatureItems ListAccessFeatures(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List access features

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
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AccessFeaturesAPI.ListAccessFeatures(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AccessFeaturesAPI.ListAccessFeatures``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAccessFeatures`: AccessFeatureItems
 fmt.Fprintf(os.Stdout, "Response from `AccessFeaturesAPI.ListAccessFeatures`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListAccessFeaturesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**AccessFeatureItems**](AccessFeatureItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAccessFeature

> UpdateAccessFeature(ctx, accessFeatureId).Body(body).Execute()

Update access_feature

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
 body := *openapiclient.NewAccessFeature("Feature_example", "Host_example") // AccessFeature | 
 accessFeatureId := int32(56) // int32 | Access feature ID

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AccessFeaturesAPI.UpdateAccessFeature(context.Background(), accessFeatureId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AccessFeaturesAPI.UpdateAccessFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**accessFeatureId** | **int32** | Access feature ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAccessFeatureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AccessFeature**](AccessFeature.md) |  |

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
