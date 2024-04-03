# \FaxesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**SendFax**](FaxesAPI.md#SendFax) | **Post** /faxes | Send a fax
[**SendUserFax**](FaxesAPI.md#SendUserFax) | **Post** /users/me/faxes | Send a fax as the user detected from the token

## SendFax

> Fax SendFax(ctx).FaxContent(faxContent).Context(context).Extension(extension).CallerId(callerId).IvrExtension(ivrExtension).WaitTime(waitTime).Execute()

Send a fax

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 faxContent := os.NewFile(1234, "some_file") // *os.File | The fax file content, in PDF format
 context := "context_example" // string | Context of the recipient of the fax
 extension := "extension_example" // string | Extension of the recipient of the fax
 callerId := "callerId_example" // string | Caller ID that will be presented to the recipient of the fax. Example: \"my-name <+15551112222>\" (optional) (default to "Accent Fax")
 ivrExtension := "ivrExtension_example" // string | Extension to compose before sending fax. Useful for fax in IVR (optional)
 waitTime := int32(56) // int32 | Time waiting before sending fax when destination has answered (in seconds) (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FaxesAPI.SendFax(context.Background()).FaxContent(faxContent).Context(context).Extension(extension).CallerId(callerId).IvrExtension(ivrExtension).WaitTime(waitTime).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FaxesAPI.SendFax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SendFax`: Fax
 fmt.Fprintf(os.Stdout, "Response from `FaxesAPI.SendFax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSendFaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **faxContent** | ***os.File** | The fax file content, in PDF format |
 **context** | **string** | Context of the recipient of the fax |
 **extension** | **string** | Extension of the recipient of the fax |
 **callerId** | **string** | Caller ID that will be presented to the recipient of the fax. Example: \&quot;my-name &lt;+15551112222&gt;\&quot; | [default to &quot;Accent Fax&quot;]
 **ivrExtension** | **string** | Extension to compose before sending fax. Useful for fax in IVR |
 **waitTime** | **int32** | Time waiting before sending fax when destination has answered (in seconds) |

### Return type

[**Fax**](Fax.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/pdf
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## SendUserFax

> Fax SendUserFax(ctx).FaxContent(faxContent).Extension(extension).CallerId(callerId).IvrExtension(ivrExtension).WaitTime(waitTime).Execute()

Send a fax as the user detected from the token

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calld"
)

func main() {
 faxContent := os.NewFile(1234, "some_file") // *os.File | The fax file content, in PDF format
 extension := "extension_example" // string | Extension of the recipient of the fax
 callerId := "callerId_example" // string | Caller ID that will be presented to the recipient of the fax. Example: \"my-name <+15551112222>\" (optional) (default to "Accent Fax")
 ivrExtension := "ivrExtension_example" // string | Extension to compose before sending fax. Useful for fax in IVR (optional)
 waitTime := int32(56) // int32 | Time waiting before sending fax when destination has answered (in seconds) (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.FaxesAPI.SendUserFax(context.Background()).FaxContent(faxContent).Extension(extension).CallerId(callerId).IvrExtension(ivrExtension).WaitTime(waitTime).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `FaxesAPI.SendUserFax``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `SendUserFax`: Fax
 fmt.Fprintf(os.Stdout, "Response from `FaxesAPI.SendUserFax`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiSendUserFaxRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **faxContent** | ***os.File** | The fax file content, in PDF format |
 **extension** | **string** | Extension of the recipient of the fax |
 **callerId** | **string** | Caller ID that will be presented to the recipient of the fax. Example: \&quot;my-name &lt;+15551112222&gt;\&quot; | [default to &quot;Accent Fax&quot;]
 **ivrExtension** | **string** | Extension to compose before sending fax. Useful for fax in IVR |
 **waitTime** | **int32** | Time waiting before sending fax when destination has answered (in seconds) |

### Return type

[**Fax**](Fax.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/pdf
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
