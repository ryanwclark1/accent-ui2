# \MessagesAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateRoomMessage**](MessagesAPI.md#CreateRoomMessage) | **Post** /users/me/rooms/{room_uuid}/messages | Create room messages
[**ListRoomMessage**](MessagesAPI.md#ListRoomMessage) | **Get** /users/me/rooms/{room_uuid}/messages | List room messages
[**ListRoomsMessages**](MessagesAPI.md#ListRoomsMessages) | **Get** /users/me/rooms/messages | List rooms messages

## CreateRoomMessage

> Message CreateRoomMessage(ctx, roomUuid).Body(body).Execute()

Create room messages

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 body := *openapiclient.NewUserMessagePOST("Alias_example", "Content_example") // UserMessagePOST | message to create
 roomUuid := "roomUuid_example" // string | The UUID of the room

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MessagesAPI.CreateRoomMessage(context.Background(), roomUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.CreateRoomMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateRoomMessage`: Message
 fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.CreateRoomMessage`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomUuid** | **string** | The UUID of the room |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateRoomMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**UserMessagePOST**](UserMessagePOST.md) | message to create |

### Return type

[**Message**](Message.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListRoomMessage

> Messages ListRoomMessage(ctx, roomUuid).FromDate(fromDate).Direction(direction).Limit(limit).Order(order).Offset(offset).Search(search).Execute()

List room messages

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 roomUuid := "roomUuid_example" // string | The UUID of the room
 fromDate := time.Now() // time.Time | The date and time from which to retrieve messages. Example: 2019-06-12T10:00:00.000+00:00 (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. Required if `distinct` is not specified. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MessagesAPI.ListRoomMessage(context.Background(), roomUuid).FromDate(fromDate).Direction(direction).Limit(limit).Order(order).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ListRoomMessage``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListRoomMessage`: Messages
 fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ListRoomMessage`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**roomUuid** | **string** | The UUID of the room |

### Other Parameters

Other parameters are passed through a pointer to a apiListRoomMessageRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **fromDate** | **time.Time** | The date and time from which to retrieve messages. Example: 2019-06-12T10:00:00.000+00:00 |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. Required if &#x60;distinct&#x60; is not specified. |

### Return type

[**Messages**](Messages.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListRoomsMessages

> Messages ListRoomsMessages(ctx).Direction(direction).Limit(limit).Order(order).Offset(offset).Search(search).Distinct(distinct).Execute()

List rooms messages

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark/accent-voice/chatd"
)

func main() {
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 distinct := "distinct_example" // string | Distinct list results by field. Always picks the latest entry. Required if `search` is not specified. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.MessagesAPI.ListRoomsMessages(context.Background()).Direction(direction).Limit(limit).Order(order).Offset(offset).Search(search).Distinct(distinct).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `MessagesAPI.ListRoomsMessages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListRoomsMessages`: Messages
 fmt.Fprintf(os.Stdout, "Response from `MessagesAPI.ListRoomsMessages`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListRoomsMessagesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **distinct** | **string** | Distinct list results by field. Always picks the latest entry. Required if &#x60;search&#x60; is not specified. |

### Return type

[**Messages**](Messages.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
