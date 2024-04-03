# \UsersAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetCurrentUserCDR**](UsersAPI.md#GetCurrentUserCDR) | **Get** /users/me/cdr | List CDR of the authenticated user
[**GetUserCDR**](UsersAPI.md#GetUserCDR) | **Get** /users/{user_uuid}/cdr | List CDR of the given user

## GetCurrentUserCDR

> CDRList GetCurrentUserCDR(ctx).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).UserUuid(userUuid).Distinct(distinct).Recorded(recorded).Format(format).Execute()

List CDR of the authenticated user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
)

func main() {
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list. Default to 1000 if not specified. (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. Unsupported values: ``end``. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 search := "search_example" // string | Filter list of items (optional)
 callDirection := "callDirection_example" // string | Filter list of items (optional)
 number := "number_example" // string | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. (optional)
 fromId := int32(56) // int32 | Ignore CDR created before the given CDR ID. (optional)
 userUuid := []string{"Inner_example"} // []string | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. (optional)
 distinct := "distinct_example" // string | Will only return one result for the selected field (optional)
 recorded := true // bool | Filter by recorded status. (optional)
 format := "format_example" // string | Overrides the Content-Type header. This is used to be able to have a downloadable link. Allowed values are \"csv\" and \"json\" (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetCurrentUserCDR(context.Background()).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).UserUuid(userUuid).Distinct(distinct).Recorded(recorded).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetCurrentUserCDR``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCurrentUserCDR`: CDRList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetCurrentUserCDR`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserCDRRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **limit** | **int32** | Maximum number of items to return in the list. Default to 1000 if not specified. |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. Unsupported values: &#x60;&#x60;end&#x60;&#x60;. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **search** | **string** | Filter list of items |
 **callDirection** | **string** | Filter list of items |
 **number** | **string** | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. |
 **fromId** | **int32** | Ignore CDR created before the given CDR ID. |
 **userUuid** | **[]string** | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. |
 **distinct** | **string** | Will only return one result for the selected field |
 **recorded** | **bool** | Filter by recorded status. |
 **format** | **string** | Overrides the Content-Type header. This is used to be able to have a downloadable link. Allowed values are \&quot;csv\&quot; and \&quot;json\&quot; |

### Return type

[**CDRList**](CDRList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetUserCDR

> CDRList GetUserCDR(ctx, userUuid).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).Distinct(distinct).Recorded(recorded).Format(format).Execute()

List CDR of the given user

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
)

func main() {
 userUuid := "userUuid_example" // string | UUID of the given user
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list. Default to 1000 if not specified. (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. Unsupported values: ``end``. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 search := "search_example" // string | Filter list of items (optional)
 callDirection := "callDirection_example" // string | Filter list of items (optional)
 number := "number_example" // string | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. (optional)
 fromId := int32(56) // int32 | Ignore CDR created before the given CDR ID. (optional)
 distinct := "distinct_example" // string | Will only return one result for the selected field (optional)
 recorded := true // bool | Filter by recorded status. (optional)
 format := "format_example" // string | Overrides the Content-Type header. This is used to be able to have a downloadable link. Allowed values are \"csv\" and \"json\" (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.UsersAPI.GetUserCDR(context.Background(), userUuid).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).Distinct(distinct).Recorded(recorded).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `UsersAPI.GetUserCDR``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserCDR`: CDRList
 fmt.Fprintf(os.Stdout, "Response from `UsersAPI.GetUserCDR`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**userUuid** | **string** | UUID of the given user |

### Other Parameters

Other parameters are passed through a pointer to a apiGetUserCDRRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **limit** | **int32** | Maximum number of items to return in the list. Default to 1000 if not specified. |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. Unsupported values: &#x60;&#x60;end&#x60;&#x60;. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **search** | **string** | Filter list of items |
 **callDirection** | **string** | Filter list of items |
 **number** | **string** | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. |
 **fromId** | **int32** | Ignore CDR created before the given CDR ID. |
 **distinct** | **string** | Will only return one result for the selected field |
 **recorded** | **bool** | Filter by recorded status. |
 **format** | **string** | Overrides the Content-Type header. This is used to be able to have a downloadable link. Allowed values are \&quot;csv\&quot; and \&quot;json\&quot; |

### Return type

[**CDRList**](CDRList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
