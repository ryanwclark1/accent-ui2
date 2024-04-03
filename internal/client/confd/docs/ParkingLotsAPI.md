# \ParkingLotsAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AssociateParkingLotExtension**](ParkingLotsAPI.md#AssociateParkingLotExtension) | **Put** /parkinglots/{parking_lot_id}/extensions/{extension_id} | Associate parking_lot and extension
[**CreateParkingLot**](ParkingLotsAPI.md#CreateParkingLot) | **Post** /parkinglots | Create parking lot
[**DeleteParkingLot**](ParkingLotsAPI.md#DeleteParkingLot) | **Delete** /parkinglots/{parking_lot_id} | Delete parking lot
[**DissociateParkingLotExtension**](ParkingLotsAPI.md#DissociateParkingLotExtension) | **Delete** /parkinglots/{parking_lot_id}/extensions/{extension_id} | Dissociate parking lot and extension
[**GetParkingLot**](ParkingLotsAPI.md#GetParkingLot) | **Get** /parkinglots/{parking_lot_id} | Get parking lot
[**ListParkingLots**](ParkingLotsAPI.md#ListParkingLots) | **Get** /parkinglots | List parking lots
[**UpdateParkingLot**](ParkingLotsAPI.md#UpdateParkingLot) | **Put** /parkinglots/{parking_lot_id} | Update parking lot

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
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ParkingLotsAPI.AssociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.AssociateParkingLotExtension``: %v\n", err)
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

## CreateParkingLot

> ParkingLot CreateParkingLot(ctx).Body(body).AccentTenant(accentTenant).Execute()

Create parking lot

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
 body := *openapiclient.NewParkingLot() // ParkingLot | Parking lot to create
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ParkingLotsAPI.CreateParkingLot(context.Background()).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.CreateParkingLot``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `CreateParkingLot`: ParkingLot
 fmt.Fprintf(os.Stdout, "Response from `ParkingLotsAPI.CreateParkingLot`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiCreateParkingLotRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ParkingLot**](ParkingLot.md) | Parking lot to create |
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**ParkingLot**](ParkingLot.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## DeleteParkingLot

> DeleteParkingLot(ctx, parkingLotId).AccentTenant(accentTenant).Execute()

Delete parking lot

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
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ParkingLotsAPI.DeleteParkingLot(context.Background(), parkingLotId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.DeleteParkingLot``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parkingLotId** | **int32** | Parking Lot&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteParkingLotRequest struct via the builder pattern

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
 openapiclient "github.com/ryanwclark/accent-voice/confd"
)

func main() {
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 extensionId := int32(56) // int32 | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ParkingLotsAPI.DissociateParkingLotExtension(context.Background(), parkingLotId, extensionId).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.DissociateParkingLotExtension``: %v\n", err)
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

## GetParkingLot

> ParkingLot GetParkingLot(ctx, parkingLotId).AccentTenant(accentTenant).Execute()

Get parking lot

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
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ParkingLotsAPI.GetParkingLot(context.Background(), parkingLotId).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.GetParkingLot``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetParkingLot`: ParkingLot
 fmt.Fprintf(os.Stdout, "Response from `ParkingLotsAPI.GetParkingLot`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parkingLotId** | **int32** | Parking Lot&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiGetParkingLotRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

### Return type

[**ParkingLot**](ParkingLot.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListParkingLots

> ParkingLotItems ListParkingLots(ctx).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()

List parking lots

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
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 recurse := true // bool | Should the query include sub-tenants (optional) (default to false)
 order := "order_example" // string | Name of the field to use for sorting the list of items returned. (optional)
 direction := "direction_example" // string | Sort list of items in 'asc' (ascending) or 'desc' (descending) order (optional)
 limit := int32(56) // int32 | Maximum number of items to return in the list (optional)
 offset := int32(56) // int32 | Number of items to skip over in the list. Useful for pagination. (optional)
 search := "search_example" // string | Search term for filtering a list of items. Only items with a field containing the search term will be returned. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.ParkingLotsAPI.ListParkingLots(context.Background()).AccentTenant(accentTenant).Recurse(recurse).Order(order).Direction(direction).Limit(limit).Offset(offset).Search(search).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.ListParkingLots``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListParkingLots`: ParkingLotItems
 fmt.Fprintf(os.Stdout, "Response from `ParkingLotsAPI.ListParkingLots`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiListParkingLotsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **recurse** | **bool** | Should the query include sub-tenants | [default to false]
 **order** | **string** | Name of the field to use for sorting the list of items returned. |
 **direction** | **string** | Sort list of items in &#39;asc&#39; (ascending) or &#39;desc&#39; (descending) order |
 **limit** | **int32** | Maximum number of items to return in the list |
 **offset** | **int32** | Number of items to skip over in the list. Useful for pagination. |
 **search** | **string** | Search term for filtering a list of items. Only items with a field containing the search term will be returned. |

### Return type

[**ParkingLotItems**](ParkingLotItems.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateParkingLot

> UpdateParkingLot(ctx, parkingLotId).Body(body).AccentTenant(accentTenant).Execute()

Update parking lot

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
 body := *openapiclient.NewParkingLot() // ParkingLot | 
 parkingLotId := int32(56) // int32 | Parking Lot's ID
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.ParkingLotsAPI.UpdateParkingLot(context.Background(), parkingLotId).Body(body).AccentTenant(accentTenant).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `ParkingLotsAPI.UpdateParkingLot``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**parkingLotId** | **int32** | Parking Lot&#39;s ID |

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateParkingLotRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ParkingLot**](ParkingLot.md) |  |

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |

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
