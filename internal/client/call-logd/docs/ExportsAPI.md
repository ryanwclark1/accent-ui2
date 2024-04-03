# \ExportsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCDRRecordingsMediaExport**](ExportsAPI.md#CreateCDRRecordingsMediaExport) | **Post** /cdr/recordings/media/export | Create an export for the recording media of multiple CDRs
[**GetExport**](ExportsAPI.md#GetExport) | **Get** /exports/{export_uuid} | Get an export by the given UUID
[**GetExportDownload**](ExportsAPI.md#GetExportDownload) | **Get** /exports/{export_uuid}/download | Download an export as a ZIP archive by the given UUID

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
 resp, r, err := apiClient.ExportsAPI.CreateCDRRecordingsMediaExport(context.Background()).Body(body).From(from).Until(until).Search(search).CallDirection(callDirection).Number(number).Tags(tags).UserUuid(userUuid).FromId(fromId).Recurse(recurse).AccentTenant(accentTenant).Email(email).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.CreateCDRRecordingsMediaExport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateCDRRecordingsMediaExport`: CreateCDRRecordingsMediaExport202Response
 fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.CreateCDRRecordingsMediaExport`: %v\n", resp)
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

## GetExport

> Export GetExport(ctx, exportUuid).Execute()

Get an export by the given UUID

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
 exportUuid := "exportUuid_example" // string | UUID of the given export

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExportsAPI.GetExport(context.Background(), exportUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.GetExport``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExport`: Export
 fmt.Fprintf(os.Stdout, "Response from `ExportsAPI.GetExport`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportUuid** | **string** | UUID of the given export |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Export**](Export.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetExportDownload

> GetExportDownload(ctx, exportUuid).Execute()

Download an export as a ZIP archive by the given UUID

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
 exportUuid := "exportUuid_example" // string | UUID of the given export

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExportsAPI.GetExportDownload(context.Background(), exportUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExportsAPI.GetExportDownload``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**exportUuid** | **string** | UUID of the given export |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExportDownloadRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

 (empty response body)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/zip

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
