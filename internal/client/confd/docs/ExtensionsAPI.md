# \ExtensionsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateConferenceExtension**](ExtensionsAPI.md#AssociateConferenceExtension) | **Put** /conferences/{conference_id}/extensions/{extension_id} | Associate conference and extension
[**AssociateGroupExtension**](ExtensionsAPI.md#AssociateGroupExtension) | **Put** /groups/{group_uuid}/extensions/{extension_id} | Associate group and extension
[**AssociateIncallExtension**](ExtensionsAPI.md#AssociateIncallExtension) | **Put** /incalls/{incall_id}/extensions/{extension_id} | Associate incall and extension
[**AssociateLineExtension**](ExtensionsAPI.md#AssociateLineExtension) | **Put** /lines/{line_id}/extensions/{extension_id} | Associate line and extension
[**AssociateOutcallExtension**](ExtensionsAPI.md#AssociateOutcallExtension) | **Put** /outcalls/{outcall_id}/extensions/{extension_id} | Associate outcall and extension
[**AssociateParkingLotExtension**](ExtensionsAPI.md#AssociateParkingLotExtension) | **Put** /parkinglots/{parking_lot_id}/extensions/{extension_id} | Associate parking_lot and extension
[**AssociateQueueExtension**](ExtensionsAPI.md#AssociateQueueExtension) | **Put** /queues/{queue_id}/extensions/{extension_id} | Associate queue and extension
[**CreateExtension**](ExtensionsAPI.md#CreateExtension) | **Post** /extensions | Create extension
[**CreateLineExtension**](ExtensionsAPI.md#CreateLineExtension) | **Post** /lines/{line_id}/extensions | Create extension
[**DeleteExtension**](ExtensionsAPI.md#DeleteExtension) | **Delete** /extensions/{extension_id} | Delete extension
[**DissociateConferenceExtension**](ExtensionsAPI.md#DissociateConferenceExtension) | **Delete** /conferences/{conference_id}/extensions/{extension_id} | Dissociate conference and extension
[**DissociateGroupExtension**](ExtensionsAPI.md#DissociateGroupExtension) | **Delete** /groups/{group_uuid}/extensions/{extension_id} | Dissociate group and extension
[**DissociateIncallExtension**](ExtensionsAPI.md#DissociateIncallExtension) | **Delete** /incalls/{incall_id}/extensions/{extension_id} | Dissociate incall and extension
[**DissociateLineExtension**](ExtensionsAPI.md#DissociateLineExtension) | **Delete** /lines/{line_id}/extensions/{extension_id} | Dissociate line and extension
[**DissociateOutcallExtension**](ExtensionsAPI.md#DissociateOutcallExtension) | **Delete** /outcalls/{outcall_id}/extensions/{extension_id} | Dissociate outcall and extension
[**DissociateParkingLotExtension**](ExtensionsAPI.md#DissociateParkingLotExtension) | **Delete** /parkinglots/{parking_lot_id}/extensions/{extension_id} | Dissociate parking lot and extension
[**DissociateQueueExtension**](ExtensionsAPI.md#DissociateQueueExtension) | **Delete** /queues/{queue_id}/extensions/{extension_id} | Dissociate queue and extension
[**GetExtension**](ExtensionsAPI.md#GetExtension) | **Get** /extensions/{extension_id} | Get extension
[**GetExtensionFeature**](ExtensionsAPI.md#GetExtensionFeature) | **Get** /extensions/features/{extension_uuid} | Get extension feature
[**ListExtensions**](ExtensionsAPI.md#ListExtensions) | **Get** /extensions | List extensions
[**ListExtensionsFeatures**](ExtensionsAPI.md#ListExtensionsFeatures) | **Get** /extensions/features | List extensions features
[**UpdateExtension**](ExtensionsAPI.md#UpdateExtension) | **Put** /extensions/{extension_id} | Update extension
[**UpdateExtensionFeature**](ExtensionsAPI.md#UpdateExtensionFeature) | **Put** /extensions/features/{extension_uuid} | Update extension

## AssociateConferenceExtension

> AssociateConferenceExtension(ctx, conferenceId, extensionId).Execute()

Associate conference and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateConferenceExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateConferenceExtensionRequest struct via the builder pattern

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

## AssociateGroupExtension

> AssociateGroupExtension(ctx, groupUuid, extensionId).Execute()

Associate group and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 groupUuid := "groupUuid_example" // string | the group's UUID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateGroupExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateGroupExtensionRequest struct via the builder pattern

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

## AssociateIncallExtension

> AssociateIncallExtension(ctx, incallId, extensionId).Execute()

Associate incall and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 incallId := int32(56) // int32 | Incoming call's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateIncallExtension(context.Background(), incallId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateIncallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incallId** | **int32** | Incoming call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateIncallExtensionRequest struct via the builder pattern

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

## AssociateLineExtension

> AssociateLineExtension(ctx, lineId, extensionId).Execute()

Associate line and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateLineExtension(context.Background(), lineId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateLineExtensionRequest struct via the builder pattern

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

## AssociateOutcallExtension

> AssociateOutcallExtension(ctx, outcallId, extensionId).Body(body).Execute()

Associate outcall and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 outcallId := int32(56) // int32 | Outgoing call's ID
 extensionId := int32(56) // int32 | 
 body := *openapiclient.NewOutcallExtension() // OutcallExtension | Outgoing Extension (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateOutcallExtension(context.Background(), outcallId, extensionId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateOutcallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateOutcallExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **body** | [**OutcallExtension**](OutcallExtension.md) | Outgoing Extension |

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

## AssociateParkingLotExtension

> AssociateParkingLotExtension(ctx, parkingLotId, extensionId).Execute()

Associate parking_lot and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateParkingLotExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parkingLotId** | **int32** | Parking Lot&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateParkingLotExtensionRequest struct via the builder pattern

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

## AssociateQueueExtension

> AssociateQueueExtension(ctx, queueId, extensionId).AccentTenant(accentTenant).Execute()

Associate queue and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 queueId := int32(56) // int32 | queue's ID
 extensionId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.AssociateQueueExtension(context.Background(), queueId, extensionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.AssociateQueueExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | queue&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiAssociateQueueExtensionRequest struct via the builder pattern

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

## CreateExtension

> Extension CreateExtension(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewExtension() // Extension | Extension to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.CreateExtension(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.CreateExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateExtension`: Extension
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.CreateExtension`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Extension**](Extension.md) | Extension to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Extension**](Extension.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## CreateLineExtension

> Extension CreateLineExtension(ctx, lineId).Body(body).AccentTenant(accentTenant).Execute()

Create extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewExtension() // Extension | Extension to create
 lineId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.CreateLineExtension(context.Background(), lineId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.CreateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateLineExtension`: Extension
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.CreateLineExtension`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiCreateLineExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Extension**](Extension.md) | Extension to create |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**Extension**](Extension.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteExtension

> DeleteExtension(ctx, extensionId).Execute()

Delete extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DeleteExtension(context.Background(), extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DeleteExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteExtensionRequest struct via the builder pattern

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

## DissociateConferenceExtension

> DissociateConferenceExtension(ctx, conferenceId, extensionId).Execute()

Dissociate conference and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 conferenceId := int32(56) // int32 | Conference's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateConferenceExtension(context.Background(), conferenceId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateConferenceExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**conferenceId** | **int32** | Conference&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateConferenceExtensionRequest struct via the builder pattern

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

## DissociateGroupExtension

> DissociateGroupExtension(ctx, groupUuid, extensionId).Execute()

Dissociate group and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 groupUuid := "groupUuid_example" // string | the group's UUID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateGroupExtension(context.Background(), groupUuid, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateGroupExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**groupUuid** | **string** | the group&#39;s UUID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateGroupExtensionRequest struct via the builder pattern

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

## DissociateIncallExtension

> DissociateIncallExtension(ctx, incallId, extensionId).Execute()

Dissociate incall and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 incallId := int32(56) // int32 | Incoming call's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateIncallExtension(context.Background(), incallId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateIncallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**incallId** | **int32** | Incoming call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateIncallExtensionRequest struct via the builder pattern

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

## DissociateLineExtension

> DissociateLineExtension(ctx, lineId, extensionId).Execute()

Dissociate line and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 lineId := int32(56) // int32 | 
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateLineExtension(context.Background(), lineId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateLineExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**lineId** | **int32** |  |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateLineExtensionRequest struct via the builder pattern

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

## DissociateOutcallExtension

> DissociateOutcallExtension(ctx, outcallId, extensionId).Execute()

Dissociate outcall and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 outcallId := int32(56) // int32 | Outgoing call's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateOutcallExtension(context.Background(), outcallId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateOutcallExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**outcallId** | **int32** | Outgoing call&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateOutcallExtensionRequest struct via the builder pattern

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

## DissociateParkingLotExtension

> DissociateParkingLotExtension(ctx, parkingLotId, extensionId).Execute()

Dissociate parking lot and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateParkingLotExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parkingLotId** | **int32** | Parking Lot&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateParkingLotExtensionRequest struct via the builder pattern

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

## DissociateQueueExtension

> DissociateQueueExtension(ctx, queueId, extensionId).AccentTenant(accentTenant).Execute()

Dissociate queue and extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 queueId := int32(56) // int32 | queue's ID
 extensionId := int32(56) // int32 | 
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.DissociateQueueExtension(context.Background(), queueId, extensionId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.DissociateQueueExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | queue&#39;s ID |
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiDissociateQueueExtensionRequest struct via the builder pattern

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

## GetExtension

> Extension GetExtension(ctx, extensionId).Execute()

Get extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.GetExtension(context.Background(), extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExtension`: Extension
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtension`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**Extension**](Extension.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetExtensionFeature

> ExtensionFeature GetExtensionFeature(ctx, extensionUuid).Execute()

Get extension feature

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 extensionUuid := "extensionUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.GetExtensionFeature(context.Background(), extensionUuid).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.GetExtensionFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetExtensionFeature`: ExtensionFeature
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.GetExtensionFeature`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiGetExtensionFeatureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

### Return type

[**ExtensionFeature**](ExtensionFeature.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListExtensions

> ExtensionItems ListExtensions(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Type_(type_).Exten(exten).Context(context).Execute()

List extensions

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)
 type_ := "type__example" // string | Filter extensions of a certain type. Internal: Used for calling a line with an internal number (e.g. “1000@default”). Incall: Used for calling a line from the outside (e.g. “from-extern” with a DID) (optional)
 exten := "exten_example" // string | Filter extensions by exten number (optional)
 context := "context_example" // string | Filter extensions by context name (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.ListExtensions(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Type_(type_).Exten(exten).Context(context).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.ListExtensions``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListExtensions`: ExtensionItems
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.ListExtensions`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |
 **type_** | **string** | Filter extensions of a certain type. Internal: Used for calling a line with an internal number (e.g. “1000@default”). Incall: Used for calling a line from the outside (e.g. “from-extern” with a DID) |
 **exten** | **string** | Filter extensions by exten number |
 **context** | **string** | Filter extensions by context name |

### Return type

[**ExtensionItems**](ExtensionItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListExtensionsFeatures

> ExtensionFeatureItems ListExtensionsFeatures(ctx).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List extensions features

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ExtensionsAPI.ListExtensionsFeatures(context.Background()).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.ListExtensionsFeatures``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListExtensionsFeatures`: ExtensionFeatureItems
 fmt.Fprintf(os.Stdout, "Response from `ExtensionsAPI.ListExtensionsFeatures`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListExtensionsFeaturesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**ExtensionFeatureItems**](ExtensionFeatureItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateExtension

> UpdateExtension(ctx, extensionId).Body(body).Execute()

Update extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewExtension() // Extension | 
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.UpdateExtension(context.Background(), extensionId).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.UpdateExtension``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionId** | **int32** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**Extension**](Extension.md) |  |

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

## UpdateExtensionFeature

> UpdateExtensionFeature(ctx, extensionUuid).Body(body).Execute()

Update extension

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/confd"
)

func main() {
 body := *openapiclient.NewExtensionFeature("Exten_example") // ExtensionFeature | 
 extensionUuid := "extensionUuid_example" // string | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ExtensionsAPI.UpdateExtensionFeature(context.Background(), extensionUuid).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ExtensionsAPI.UpdateExtensionFeature``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**extensionUuid** | **string** |  |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateExtensionFeatureRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ExtensionFeature**](ExtensionFeature.md) |  |

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
