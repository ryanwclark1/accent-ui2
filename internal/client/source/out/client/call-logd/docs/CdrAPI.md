# \CdrAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCDRRecordingsMediaExport**](CdrAPI.md#CreateCDRRecordingsMediaExport) | **Post** /cdr/recordings/media/export | Create an export for the recording media of multiple CDRs
[**DeleteCDRRecordingMedia**](CdrAPI.md#DeleteCDRRecordingMedia) | **Delete** /cdr/{cdr_id}/recordings/{recording_uuid}/media | Delete a recording media
[**DeleteCDRRecordingsMedia**](CdrAPI.md#DeleteCDRRecordingsMedia) | **Delete** /cdr/recordings/media | Delete multiple CDRs recording media
[**GetCDR**](CdrAPI.md#GetCDR) | **Get** /cdr/{cdr_id} | Get a CDR by ID
[**GetCDRRecordingMedia**](CdrAPI.md#GetCDRRecordingMedia) | **Get** /cdr/{cdr_id}/recordings/{recording_uuid}/media | Get a recording media
[**GetCDRs**](CdrAPI.md#GetCDRs) | **Get** /cdr | List CDR
[**GetCurrentUserCDR**](CdrAPI.md#GetCurrentUserCDR) | **Get** /users/me/cdr | List CDR of the authenticated user
[**GetUserCDR**](CdrAPI.md#GetUserCDR) | **Get** /users/{user_uuid}/cdr | List CDR of the given user

## CreateCDRRecordingsMediaExport

> CreateCDRRecordingsMediaExport202Response CreateCDRRecordingsMediaExport(ctx).Body(body).From(from).Until(until).Search(search).CallDirection(callDirection).Number(number).Tags(tags).UserUuid(userUuid).FromId(fromId).Recurse(recurse).AccentTenant(accentTenant).Email(email).Execute()

Create an export for the recording media of multiple CDRs

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 body := *openapiclient.NewCreateCDRRecordingsMediaExportRequest() // CreateCDRRecordingsMediaExportRequest | The CDR IDs list from which to create an export (optional)
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 search := "search_example" // string | Filter list of items (optional)
 callDirection := "callDirection_example" // string | Filter list of items (optional)
 number := "number_example" // string | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. (optional)
 tags := []string{"Inner_example"} // []string | Filter by tags. Each tag MUST be separated by a coma (,). Many tag will perform a logical AND. (optional)
 userUuid := []string{"Inner_example"} // []string | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. (optional)
 fromId := int32(56) // int32 | Ignore CDR created before the given CDR ID. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 email := "email_example" // string | E-mail address (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CdrAPI.CreateCDRRecordingsMediaExport(context.Background()).Body(body).From(from).Until(until).Search(search).CallDirection(callDirection).Number(number).Tags(tags).UserUuid(userUuid).FromId(fromId).Recurse(recurse).AccentTenant(accentTenant).Email(email).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.CreateCDRRecordingsMediaExport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCDRRecordingsMediaExport`: CreateCDRRecordingsMediaExport202Response
 fmt.Fprintf(os.Stdout, "Response from `CdrAPI.CreateCDRRecordingsMediaExport`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateCDRRecordingsMediaExportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**CreateCDRRecordingsMediaExportRequest**](CreateCDRRecordingsMediaExportRequest.md) | The CDR IDs list from which to create an export |
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **search** | **string** | Filter list of items |
 **callDirection** | **string** | Filter list of items |
 **number** | **string** | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. |
 **tags** | **[]string** | Filter by tags. Each tag MUST be separated by a coma (,). Many tag will perform a logical AND. |
 **userUuid** | **[]string** | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. |
 **fromId** | **int32** | Ignore CDR created before the given CDR ID. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **email** | **string** | E-mail address |

### Return type

[**CreateCDRRecordingsMediaExport202Response**](CreateCDRRecordingsMediaExport202Response.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteCDRRecordingMedia

> DeleteCDRRecordingMedia(ctx, cdrId, recordingUuid).Execute()

Delete a recording media

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 cdrId := int32(56) // int32 | ID of the CDR
 recordingUuid := int32(56) // int32 | UUID of the recording

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CdrAPI.DeleteCDRRecordingMedia(context.Background(), cdrId, recordingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.DeleteCDRRecordingMedia``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdrId** | **int32** | ID of the CDR |
**recordingUuid** | **int32** | UUID of the recording |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCDRRecordingMediaRequest struct via the builder pattern

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

## DeleteCDRRecordingsMedia

> DeleteCDRRecordingsMedia(ctx).Body(body).Execute()

Delete multiple CDRs recording media

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 body := *openapiclient.NewDeleteCDRRecordingsMediaRequest() // DeleteCDRRecordingsMediaRequest | The CDR IDs list from which to delete recording media

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CdrAPI.DeleteCDRRecordingsMedia(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.DeleteCDRRecordingsMedia``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteCDRRecordingsMediaRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**DeleteCDRRecordingsMediaRequest**](DeleteCDRRecordingsMediaRequest.md) | The CDR IDs list from which to delete recording media |

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

## GetCDR

> CDR GetCDR(ctx, cdrId).Execute()

Get a CDR by ID

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 cdrId := int32(56) // int32 | ID of the CDR

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CdrAPI.GetCDR(context.Background(), cdrId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.GetCDR``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCDR`: CDR
 fmt.Fprintf(os.Stdout, "Response from `CdrAPI.GetCDR`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdrId** | **int32** | ID of the CDR |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCDRRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**CDR**](CDR.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json, text/csv; charset=utf-8

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCDRRecordingMedia

> GetCDRRecordingMedia(ctx, cdrId, recordingUuid).Execute()

Get a recording media

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 cdrId := int32(56) // int32 | ID of the CDR
 recordingUuid := int32(56) // int32 | UUID of the recording

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.CdrAPI.GetCDRRecordingMedia(context.Background(), cdrId, recordingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.GetCDRRecordingMedia``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**cdrId** | **int32** | ID of the CDR |
**recordingUuid** | **int32** | UUID of the recording |

### Other Parameters

Other parameters are passed through a pointer to a apiGetCDRRecordingMediaRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: audio/wav

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetCDRs

> CDRList GetCDRs(ctx).AccentTenant(accentTenant).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).Tags(tags).UserUuid(userUuid).FromId(fromId).Recurse(recurse).Distinct(distinct).Recorded(recorded).Format(format).Execute()

List CDR

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list. Default to 1000 if not specified. (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. Unsupported values: ``end``. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 search := "search_example" // string | Filter list of items (optional)
 callDirection := "callDirection_example" // string | Filter list of items (optional)
 number := "number_example" // string | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. (optional)
 tags := []string{"Inner_example"} // []string | Filter by tags. Each tag MUST be separated by a coma (,). Many tag will perform a logical AND. (optional)
 userUuid := []string{"Inner_example"} // []string | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. (optional)
 fromId := int32(56) // int32 | Ignore CDR created before the given CDR ID. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 distinct := "distinct_example" // string | Will only return one result for the selected field (optional)
 recorded := true // bool | Filter by recorded status. (optional)
 format := "format_example" // string | Overrides the Content-Type header. This is used to be able to have a downloadable link. Allowed values are \"csv\" and \"json\" (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.CdrAPI.GetCDRs(context.Background()).AccentTenant(accentTenant).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).Tags(tags).UserUuid(userUuid).FromId(fromId).Recurse(recurse).Distinct(distinct).Recorded(recorded).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.GetCDRs``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCDRs`: CDRList
 fmt.Fprintf(os.Stdout, "Response from `CdrAPI.GetCDRs`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetCDRsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **limit** | **int32** | Maximum number of items to return in the list. Default to 1000 if not specified. |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **order** | **string** | Name of the field to use for sorting the list of items returned. Unsupported values: &#x60;&#x60;end&#x60;&#x60;. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **search** | **string** | Filter list of items |
 **callDirection** | **string** | Filter list of items |
 **number** | **string** | Filter by source_extension and destination_extension. A wildcard (underscore) can be used at the start and/or the end of the number. |
 **tags** | **[]string** | Filter by tags. Each tag MUST be separated by a coma (,). Many tag will perform a logical AND. |
 **userUuid** | **[]string** | Filter by user_uuid. Many uuid can be specified. Each uuid MUST be separated by a comma (,). Many uuid will perform a logical OR. |
 **fromId** | **int32** | Ignore CDR created before the given CDR ID. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
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
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
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
 resp, r, err := apiClient.CdrAPI.GetCurrentUserCDR(context.Background()).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).UserUuid(userUuid).Distinct(distinct).Recorded(recorded).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.GetCurrentUserCDR``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetCurrentUserCDR`: CDRList
 fmt.Fprintf(os.Stdout, "Response from `CdrAPI.GetCurrentUserCDR`: %v\n", resp)
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
 openapiclient "github.com/ryanwclark/accent-voice/calllogd"
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
 resp, r, err := apiClient.CdrAPI.GetUserCDR(context.Background(), userUuid).From(from).Until(until).Limit(limit).Offset(offset).Order(order).Direction(direction).Search(search).CallDirection(callDirection).Number(number).FromId(fromId).Distinct(distinct).Recorded(recorded).Format(format).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `CdrAPI.GetUserCDR``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetUserCDR`: CDRList
 fmt.Fprintf(os.Stdout, "Response from `CdrAPI.GetUserCDR`: %v\n", resp)
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
