# \MeetingsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGuestMeetingStatus**](MeetingsAPI.md#GetGuestMeetingStatus) | **Get** /guests/me/meetings/{meeting_uuid}/status | Get the status of a meeting
[**KickMeetingParticipant**](MeetingsAPI.md#KickMeetingParticipant) | **Delete** /meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting
[**KickUserMeetingParticipant**](MeetingsAPI.md#KickUserMeetingParticipant) | **Delete** /users/me/meetings/{meeting_uuid}/participants/{participant_id} | Kick a participant from a meeting as a user
[**ListMeetingParticipants**](MeetingsAPI.md#ListMeetingParticipants) | **Get** /meetings/{meeting_uuid}/participants | List participants of a meeting
[**ListUserMeetingParticipants**](MeetingsAPI.md#ListUserMeetingParticipants) | **Get** /users/me/meetings/{meeting_uuid}/participants | List participants of a meeting as a user

## GetGuestMeetingStatus

> MeetingStatus GetGuestMeetingStatus(ctx, meetingUuid).Execute()

Get the status of a meeting

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
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MeetingsAPI.GetGuestMeetingStatus(context.Background(), meetingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.GetGuestMeetingStatus``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetGuestMeetingStatus`: MeetingStatus
 fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.GetGuestMeetingStatus`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |

### Other Parameters

Other parameters are passed through a pointer to a apiGetGuestMeetingStatusRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**MeetingStatus**](MeetingStatus.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## KickMeetingParticipant

> KickMeetingParticipant(ctx, meetingUuid, participantId).Execute()

Kick a participant from a meeting

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
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.MeetingsAPI.KickMeetingParticipant(context.Background(), meetingUuid, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.KickMeetingParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiKickMeetingParticipantRequest struct via the builder pattern

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

## KickUserMeetingParticipant

> KickUserMeetingParticipant(ctx, meetingUuid, participantId).Execute()

Kick a participant from a meeting as a user

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
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting
 participantId := "participantId_example" // string | Unique identifier of the participant

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.MeetingsAPI.KickUserMeetingParticipant(context.Background(), meetingUuid, participantId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.KickUserMeetingParticipant``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |
**participantId** | **string** | Unique identifier of the participant |

### Other Parameters

Other parameters are passed through a pointer to a apiKickUserMeetingParticipantRequest struct via the builder pattern

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

## ListMeetingParticipants

> ParticipantList ListMeetingParticipants(ctx, meetingUuid).Execute()

List participants of a meeting

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
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MeetingsAPI.ListMeetingParticipants(context.Background(), meetingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.ListMeetingParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListMeetingParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.ListMeetingParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |

### Other Parameters

Other parameters are passed through a pointer to a apiListMeetingParticipantsRequest struct via the builder pattern

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

## ListUserMeetingParticipants

> ParticipantList ListUserMeetingParticipants(ctx, meetingUuid).Execute()

List participants of a meeting as a user

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
 meetingUuid := "meetingUuid_example" // string | Unique identifier of the meeting

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MeetingsAPI.ListUserMeetingParticipants(context.Background(), meetingUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MeetingsAPI.ListUserMeetingParticipants``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListUserMeetingParticipants`: ParticipantList
 fmt.Fprintf(os.Stdout, "Response from `MeetingsAPI.ListUserMeetingParticipants`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**meetingUuid** | **string** | Unique identifier of the meeting |

### Other Parameters

Other parameters are passed through a pointer to a apiListUserMeetingParticipantsRequest struct via the builder pattern

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
