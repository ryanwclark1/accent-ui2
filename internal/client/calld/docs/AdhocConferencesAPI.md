# \AdhocConferencesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AddParticipantToAdhocConference**](AdhocConferencesAPI.md#AddParticipantToAdhocConference) | **Put** /users/me/conferences/adhoc/{conference_id}/participants/{call_id} | Add a participant into an adhoc conference
[**CreateAdhocConference**](AdhocConferencesAPI.md#CreateAdhocConference) | **Post** /users/me/conferences/adhoc | Create an adhoc conference
[**DeleteAdhocConference**](AdhocConferencesAPI.md#DeleteAdhocConference) | **Delete** /users/me/conferences/adhoc/{conference_id} | Delete an adhoc conference
[**RemoveParticipantFromAdhocConference**](AdhocConferencesAPI.md#RemoveParticipantFromAdhocConference) | **Delete** /users/me/conferences/adhoc/{conference_id}/participants/{call_id} | Remove a participant from an adhoc conference

## AddParticipantToAdhocConference

> AddParticipantToAdhocConference(ctx, conferenceId, callId).Execute()

Add a participant into an adhoc conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 conferenceId := "conferenceId_example" // string | ID of the adhoc conference
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AdhocConferencesAPI.AddParticipantToAdhocConference(context.Background(), conferenceId, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AdhocConferencesAPI.AddParticipantToAdhocConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | ID of the adhoc conference |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiAddParticipantToAdhocConferenceRequest struct via the builder pattern

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

## CreateAdhocConference

> AdhocConference CreateAdhocConference(ctx).Body(body).Execute()

Create an adhoc conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 body := *openapiclient.NewAdhocConferenceCreation() // AdhocConferenceCreation | Parameters of the conference calls

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AdhocConferencesAPI.CreateAdhocConference(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AdhocConferencesAPI.CreateAdhocConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateAdhocConference`: AdhocConference
 fmt.Fprintf(os.Stdout, "Response from `AdhocConferencesAPI.CreateAdhocConference`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateAdhocConferenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**AdhocConferenceCreation**](AdhocConferenceCreation.md) | Parameters of the conference calls |

### Return type

[**AdhocConference**](AdhocConference.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteAdhocConference

> AdhocConference DeleteAdhocConference(ctx, conferenceId).Execute()

Delete an adhoc conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 conferenceId := "conferenceId_example" // string | ID of the adhoc conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AdhocConferencesAPI.DeleteAdhocConference(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AdhocConferencesAPI.DeleteAdhocConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `DeleteAdhocConference`: AdhocConference
 fmt.Fprintf(os.Stdout, "Response from `AdhocConferencesAPI.DeleteAdhocConference`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | ID of the adhoc conference |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteAdhocConferenceRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**AdhocConference**](AdhocConference.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## RemoveParticipantFromAdhocConference

> RemoveParticipantFromAdhocConference(ctx, conferenceId, callId).Execute()

Remove a participant from an adhoc conference

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/calld"
)

func main() {
 conferenceId := "conferenceId_example" // string | ID of the adhoc conference
 callId := "callId_example" // string | ID of the call

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AdhocConferencesAPI.RemoveParticipantFromAdhocConference(context.Background(), conferenceId, callId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AdhocConferencesAPI.RemoveParticipantFromAdhocConference``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | ID of the adhoc conference |
**callId** | **string** | ID of the call |

### Other Parameters

Other parameters are passed through a pointer to a apiRemoveParticipantFromAdhocConferenceRequest struct via the builder pattern

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
