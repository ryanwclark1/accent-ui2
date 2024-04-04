# \ConferencesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**KickParticipant**](ConferencesAPI.md#KickParticipant) | **Delete** /conferences/{conference_id}/participants/{participant_id} | Kick participant from a conference
[**ListConferenceParticipants**](ConferencesAPI.md#ListConferenceParticipants) | **Get** /conferences/{conference_id}/participants | List participants of a conference
[**ListUserConferenceParticipants**](ConferencesAPI.md#ListUserConferenceParticipants) | **Get** /users/me/conferences/{conference_id}/participants | List participants of a conference as a user
[**MuteParticipant**](ConferencesAPI.md#MuteParticipant) | **Put** /conferences/{conference_id}/participants/{participant_id}/mute | Mute a participant in a conference
[**StartConferenceRecording**](ConferencesAPI.md#StartConferenceRecording) | **Post** /conferences/{conference_id}/record | Record a conference
[**StopConferenceRecording**](ConferencesAPI.md#StopConferenceRecording) | **Delete** /conferences/{conference_id}/record | Stop recording a conference
[**UnmuteParticipant**](ConferencesAPI.md#UnmuteParticipant) | **Put** /conferences/{conference_id}/participants/{participant_id}/unmute | Unmute a participant in a conference

## KickParticipant

> KickParticipant(ctx, conferenceId, participantId).Execute()

Kick participant from a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.KickParticipant(context.Background(), conferenceId, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.KickParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiKickParticipantRequest struct via the builder pattern

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

## ListConferenceParticipants

> ParticipantList ListConferenceParticipants(ctx, conferenceId).Execute()

List participants of a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.ListConferenceParticipants(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListConferenceParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListConferenceParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |

### Other Parameters

Other parameters are passed through a pointer to a apiListConferenceParticipantsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ParticipantList**](ParticipantList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListUserConferenceParticipants

> ParticipantList ListUserConferenceParticipants(ctx, conferenceId).Execute()

List participants of a conference as a user

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ConferencesAPI.ListUserConferenceParticipants(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.ListUserConferenceParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserConferenceParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `ConferencesAPI.ListUserConferenceParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserConferenceParticipantsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ParticipantList**](ParticipantList.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## MuteParticipant

> MuteParticipant(ctx, conferenceId, participantId).Execute()

Mute a participant in a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.MuteParticipant(context.Background(), conferenceId, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.MuteParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiMuteParticipantRequest struct via the builder pattern

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

## StartConferenceRecording

> StartConferenceRecording(ctx, conferenceId).Execute()

Record a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.StartConferenceRecording(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.StartConferenceRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |

### Other Parameters

Other parameters are passed through a pointer to a apiStartConferenceRecordingRequest struct via the builder pattern

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

## StopConferenceRecording

> StopConferenceRecording(ctx, conferenceId).Execute()

Stop recording a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.StopConferenceRecording(context.Background(), conferenceId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.StopConferenceRecording``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |

### Other Parameters

Other parameters are passed through a pointer to a apiStopConferenceRecordingRequest struct via the builder pattern

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

## UnmuteParticipant

> UnmuteParticipant(ctx, conferenceId, participantId).Execute()

Unmute a participant in a conference

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
 conferenceId := "conferenceId_example" // string | Unique identifier of the conference
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ConferencesAPI.UnmuteParticipant(context.Background(), conferenceId, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ConferencesAPI.UnmuteParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **string** | Unique identifier of the conference |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiUnmuteParticipantRequest struct via the builder pattern

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
