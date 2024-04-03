# \AsteriskAPI

All URIs are relative to *<http://api.accentvoice.io/1.1>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ListAsteriskConfbridgeAccentDefaultBridge**](AsteriskAPI.md#ListAsteriskConfbridgeAccentDefaultBridge) | **Get** /asterisk/confbridge/accent_default_bridge | List ConfBridge accent_default_bridge options
[**ListAsteriskConfbridgeAccentDefaultUser**](AsteriskAPI.md#ListAsteriskConfbridgeAccentDefaultUser) | **Get** /asterisk/confbridge/accent_default_user | List ConfBridge accent_default_user options
[**ListAsteriskFeaturesApplicationmap**](AsteriskAPI.md#ListAsteriskFeaturesApplicationmap) | **Get** /asterisk/features/applicationmap | List Features applicationmap options
[**ListAsteriskFeaturesFeaturemap**](AsteriskAPI.md#ListAsteriskFeaturesFeaturemap) | **Get** /asterisk/features/featuremap | List Features featuremap options
[**ListAsteriskFeaturesGeneral**](AsteriskAPI.md#ListAsteriskFeaturesGeneral) | **Get** /asterisk/features/general | List Features general options
[**ListAsteriskHepGeneral**](AsteriskAPI.md#ListAsteriskHepGeneral) | **Get** /asterisk/hep/general | List HEP general options
[**ListAsteriskIaxCallnumberlimits**](AsteriskAPI.md#ListAsteriskIaxCallnumberlimits) | **Get** /asterisk/iax/callnumberlimits | List IAX callnumberlimits options
[**ListAsteriskIaxGeneral**](AsteriskAPI.md#ListAsteriskIaxGeneral) | **Get** /asterisk/iax/general | List IAX general options
[**ListAsteriskPjsipGlobal**](AsteriskAPI.md#ListAsteriskPjsipGlobal) | **Get** /asterisk/pjsip/global | List of PJSIP options for the &#x60;global&#x60; section
[**ListAsteriskPjsipSystem**](AsteriskAPI.md#ListAsteriskPjsipSystem) | **Get** /asterisk/pjsip/system | List of PJSIP options for the &#x60;system&#x60; section
[**ListAsteriskQueueGeneral**](AsteriskAPI.md#ListAsteriskQueueGeneral) | **Get** /asterisk/queues/general | List Queue general options
[**ListAsteriskRtpGeneral**](AsteriskAPI.md#ListAsteriskRtpGeneral) | **Get** /asterisk/rtp/general | List RTP general options
[**ListAsteriskRtpIceHostCandidates**](AsteriskAPI.md#ListAsteriskRtpIceHostCandidates) | **Get** /asterisk/rtp/ice_host_candidates | List RTP ice_host_candidates options
[**ListAsteriskSccpGeneral**](AsteriskAPI.md#ListAsteriskSccpGeneral) | **Get** /asterisk/sccp/general | List SCCP general options
[**ListAsteriskVoicemailGeneral**](AsteriskAPI.md#ListAsteriskVoicemailGeneral) | **Get** /asterisk/voicemail/general | List Voicemail general options
[**ListAsteriskVoicemailZonemessages**](AsteriskAPI.md#ListAsteriskVoicemailZonemessages) | **Get** /asterisk/voicemail/zonemessages | List Voicemail zonemessages options
[**ShowPjsipDoc**](AsteriskAPI.md#ShowPjsipDoc) | **Get** /asterisk/pjsip/doc | List all PJSIP configuration options
[**UpdateAsteriskConfbridgeAccentDefaultBridge**](AsteriskAPI.md#UpdateAsteriskConfbridgeAccentDefaultBridge) | **Put** /asterisk/confbridge/accent_default_bridge | Update ConfBridge accent_default_bridge option
[**UpdateAsteriskConfbridgeAccentDefaultUser**](AsteriskAPI.md#UpdateAsteriskConfbridgeAccentDefaultUser) | **Put** /asterisk/confbridge/accent_default_user | Update ConfBridge accent_default_user option
[**UpdateAsteriskFeaturesApplicationmap**](AsteriskAPI.md#UpdateAsteriskFeaturesApplicationmap) | **Put** /asterisk/features/applicationmap | Update Features applicationmap option
[**UpdateAsteriskFeaturesFeaturemap**](AsteriskAPI.md#UpdateAsteriskFeaturesFeaturemap) | **Put** /asterisk/features/featuremap | Update Features featuremap option
[**UpdateAsteriskFeaturesGeneral**](AsteriskAPI.md#UpdateAsteriskFeaturesGeneral) | **Put** /asterisk/features/general | Update Features general option
[**UpdateAsteriskHepGeneral**](AsteriskAPI.md#UpdateAsteriskHepGeneral) | **Put** /asterisk/hep/general | Update HEP general option
[**UpdateAsteriskIaxCallnumberlimits**](AsteriskAPI.md#UpdateAsteriskIaxCallnumberlimits) | **Put** /asterisk/iax/callnumberlimits | Update IAX callnumberlimits option
[**UpdateAsteriskIaxGeneral**](AsteriskAPI.md#UpdateAsteriskIaxGeneral) | **Put** /asterisk/iax/general | Update IAX general option
[**UpdateAsteriskPjsipGlobal**](AsteriskAPI.md#UpdateAsteriskPjsipGlobal) | **Put** /asterisk/pjsip/global | Update PJSIP section options
[**UpdateAsteriskPjsipSystem**](AsteriskAPI.md#UpdateAsteriskPjsipSystem) | **Put** /asterisk/pjsip/system | Update PJSIP section options
[**UpdateAsteriskQueueGeneral**](AsteriskAPI.md#UpdateAsteriskQueueGeneral) | **Put** /asterisk/queues/general | Update Queue general option
[**UpdateAsteriskRtpGeneral**](AsteriskAPI.md#UpdateAsteriskRtpGeneral) | **Put** /asterisk/rtp/general | Update RTP general option
[**UpdateAsteriskRtpIceHostCandidates**](AsteriskAPI.md#UpdateAsteriskRtpIceHostCandidates) | **Put** /asterisk/rtp/ice_host_candidates | Update RTP ice_host_candidates option
[**UpdateAsteriskSccpGeneral**](AsteriskAPI.md#UpdateAsteriskSccpGeneral) | **Put** /asterisk/sccp/general | Update SCCP general option
[**UpdateAsteriskVoicemailGeneral**](AsteriskAPI.md#UpdateAsteriskVoicemailGeneral) | **Put** /asterisk/voicemail/general | Update Voicemail general option
[**UpdateAsteriskVoicemailZonemessages**](AsteriskAPI.md#UpdateAsteriskVoicemailZonemessages) | **Put** /asterisk/voicemail/zonemessages | Update Voicemail zonemessages option

## ListAsteriskConfbridgeAccentDefaultBridge

> ConfBridgeConfiguration ListAsteriskConfbridgeAccentDefaultBridge(ctx).Execute()

List ConfBridge accent_default_bridge options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskConfbridgeAccentDefaultBridge(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskConfbridgeAccentDefaultBridge``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskConfbridgeAccentDefaultBridge`: ConfBridgeConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskConfbridgeAccentDefaultBridge`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskConfbridgeAccentDefaultBridgeRequest struct via the builder pattern

### Return type

[**ConfBridgeConfiguration**](ConfBridgeConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskConfbridgeAccentDefaultUser

> ConfBridgeConfiguration ListAsteriskConfbridgeAccentDefaultUser(ctx).Execute()

List ConfBridge accent_default_user options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskConfbridgeAccentDefaultUser(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskConfbridgeAccentDefaultUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskConfbridgeAccentDefaultUser`: ConfBridgeConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskConfbridgeAccentDefaultUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskConfbridgeAccentDefaultUserRequest struct via the builder pattern

### Return type

[**ConfBridgeConfiguration**](ConfBridgeConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskFeaturesApplicationmap

> FeaturesConfiguration ListAsteriskFeaturesApplicationmap(ctx).Execute()

List Features applicationmap options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskFeaturesApplicationmap(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskFeaturesApplicationmap``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskFeaturesApplicationmap`: FeaturesConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskFeaturesApplicationmap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskFeaturesApplicationmapRequest struct via the builder pattern

### Return type

[**FeaturesConfiguration**](FeaturesConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskFeaturesFeaturemap

> FeaturesConfiguration ListAsteriskFeaturesFeaturemap(ctx).Execute()

List Features featuremap options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskFeaturesFeaturemap(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskFeaturesFeaturemap``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskFeaturesFeaturemap`: FeaturesConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskFeaturesFeaturemap`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskFeaturesFeaturemapRequest struct via the builder pattern

### Return type

[**FeaturesConfiguration**](FeaturesConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskFeaturesGeneral

> FeaturesConfiguration ListAsteriskFeaturesGeneral(ctx).Execute()

List Features general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskFeaturesGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskFeaturesGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskFeaturesGeneral`: FeaturesConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskFeaturesGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskFeaturesGeneralRequest struct via the builder pattern

### Return type

[**FeaturesConfiguration**](FeaturesConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskHepGeneral

> HEPConfiguration ListAsteriskHepGeneral(ctx).Execute()

List HEP general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskHepGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskHepGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskHepGeneral`: HEPConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskHepGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskHepGeneralRequest struct via the builder pattern

### Return type

[**HEPConfiguration**](HEPConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskIaxCallnumberlimits

> IAXCallNumberLimitss ListAsteriskIaxCallnumberlimits(ctx).Execute()

List IAX callnumberlimits options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskIaxCallnumberlimits(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskIaxCallnumberlimits``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskIaxCallnumberlimits`: IAXCallNumberLimitss
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskIaxCallnumberlimits`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskIaxCallnumberlimitsRequest struct via the builder pattern

### Return type

[**IAXCallNumberLimitss**](IAXCallNumberLimitss.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskIaxGeneral

> IAXGeneral ListAsteriskIaxGeneral(ctx).Execute()

List IAX general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskIaxGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskIaxGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskIaxGeneral`: IAXGeneral
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskIaxGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskIaxGeneralRequest struct via the builder pattern

### Return type

[**IAXGeneral**](IAXGeneral.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskPjsipGlobal

> PJSIPGlobal ListAsteriskPjsipGlobal(ctx).Execute()

List of PJSIP options for the `global` section

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskPjsipGlobal(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskPjsipGlobal``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskPjsipGlobal`: PJSIPGlobal
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskPjsipGlobal`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskPjsipGlobalRequest struct via the builder pattern

### Return type

[**PJSIPGlobal**](PJSIPGlobal.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskPjsipSystem

> PJSIPSystem ListAsteriskPjsipSystem(ctx).Execute()

List of PJSIP options for the `system` section

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskPjsipSystem(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskPjsipSystem``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskPjsipSystem`: PJSIPSystem
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskPjsipSystem`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskPjsipSystemRequest struct via the builder pattern

### Return type

[**PJSIPSystem**](PJSIPSystem.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskQueueGeneral

> QueueGeneral ListAsteriskQueueGeneral(ctx).Execute()

List Queue general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskQueueGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskQueueGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskQueueGeneral`: QueueGeneral
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskQueueGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskQueueGeneralRequest struct via the builder pattern

### Return type

[**QueueGeneral**](QueueGeneral.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskRtpGeneral

> RTPConfiguration ListAsteriskRtpGeneral(ctx).Execute()

List RTP general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskRtpGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskRtpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskRtpGeneral`: RTPConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskRtpGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskRtpGeneralRequest struct via the builder pattern

### Return type

[**RTPConfiguration**](RTPConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskRtpIceHostCandidates

> RTPConfiguration ListAsteriskRtpIceHostCandidates(ctx).Execute()

List RTP ice_host_candidates options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskRtpIceHostCandidates(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskRtpIceHostCandidates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskRtpIceHostCandidates`: RTPConfiguration
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskRtpIceHostCandidates`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskRtpIceHostCandidatesRequest struct via the builder pattern

### Return type

[**RTPConfiguration**](RTPConfiguration.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskSccpGeneral

> SCCPGeneral ListAsteriskSccpGeneral(ctx).Execute()

List SCCP general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskSccpGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskSccpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskSccpGeneral`: SCCPGeneral
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskSccpGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskSccpGeneralRequest struct via the builder pattern

### Return type

[**SCCPGeneral**](SCCPGeneral.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskVoicemailGeneral

> VoicemailGeneral ListAsteriskVoicemailGeneral(ctx).Execute()

List Voicemail general options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskVoicemailGeneral(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskVoicemailGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskVoicemailGeneral`: VoicemailGeneral
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskVoicemailGeneral`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskVoicemailGeneralRequest struct via the builder pattern

### Return type

[**VoicemailGeneral**](VoicemailGeneral.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ListAsteriskVoicemailZonemessages

> VoicemailZoneMessages ListAsteriskVoicemailZonemessages(ctx).Execute()

List Voicemail zonemessages options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ListAsteriskVoicemailZonemessages(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ListAsteriskVoicemailZonemessages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ListAsteriskVoicemailZonemessages`: VoicemailZoneMessages
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ListAsteriskVoicemailZonemessages`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiListAsteriskVoicemailZonemessagesRequest struct via the builder pattern

### Return type

[**VoicemailZoneMessages**](VoicemailZoneMessages.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## ShowPjsipDoc

> PJSIPConfigurationOptions ShowPjsipDoc(ctx).Execute()

List all PJSIP configuration options

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

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.AsteriskAPI.ShowPjsipDoc(context.Background()).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.ShowPjsipDoc``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `ShowPjsipDoc`: PJSIPConfigurationOptions
 fmt.Fprintf(os.Stdout, "Response from `AsteriskAPI.ShowPjsipDoc`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiShowPjsipDocRequest struct via the builder pattern

### Return type

[**PJSIPConfigurationOptions**](PJSIPConfigurationOptions.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## UpdateAsteriskConfbridgeAccentDefaultBridge

> UpdateAsteriskConfbridgeAccentDefaultBridge(ctx).Body(body).Execute()

Update ConfBridge accent_default_bridge option

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
 body := *openapiclient.NewConfBridgeConfiguration() // ConfBridgeConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultBridge(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultBridge``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskConfbridgeAccentDefaultBridgeRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfBridgeConfiguration**](ConfBridgeConfiguration.md) |  |

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

## UpdateAsteriskConfbridgeAccentDefaultUser

> UpdateAsteriskConfbridgeAccentDefaultUser(ctx).Body(body).Execute()

Update ConfBridge accent_default_user option

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
 body := *openapiclient.NewConfBridgeConfiguration() // ConfBridgeConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultUser(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskConfbridgeAccentDefaultUser``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskConfbridgeAccentDefaultUserRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**ConfBridgeConfiguration**](ConfBridgeConfiguration.md) |  |

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

## UpdateAsteriskFeaturesApplicationmap

> UpdateAsteriskFeaturesApplicationmap(ctx).Body(body).Execute()

Update Features applicationmap option

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
 body := *openapiclient.NewFeaturesConfiguration() // FeaturesConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesApplicationmap(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskFeaturesApplicationmap``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskFeaturesApplicationmapRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FeaturesConfiguration**](FeaturesConfiguration.md) |  |

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

## UpdateAsteriskFeaturesFeaturemap

> UpdateAsteriskFeaturesFeaturemap(ctx).Body(body).Execute()

Update Features featuremap option

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
 body := *openapiclient.NewFeaturesConfiguration() // FeaturesConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesFeaturemap(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskFeaturesFeaturemap``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskFeaturesFeaturemapRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FeaturesConfiguration**](FeaturesConfiguration.md) |  |

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

## UpdateAsteriskFeaturesGeneral

> UpdateAsteriskFeaturesGeneral(ctx).Body(body).Execute()

Update Features general option

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
 body := *openapiclient.NewFeaturesConfiguration() // FeaturesConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskFeaturesGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskFeaturesGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskFeaturesGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**FeaturesConfiguration**](FeaturesConfiguration.md) |  |

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

## UpdateAsteriskHepGeneral

> UpdateAsteriskHepGeneral(ctx).Body(body).Execute()

Update HEP general option

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
 body := *openapiclient.NewHEPConfiguration() // HEPConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskHepGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskHepGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskHepGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**HEPConfiguration**](HEPConfiguration.md) |  |

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

## UpdateAsteriskIaxCallnumberlimits

> UpdateAsteriskIaxCallnumberlimits(ctx).Body(body).Execute()

Update IAX callnumberlimits option

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
 body := *openapiclient.NewIAXCallNumberLimitss() // IAXCallNumberLimitss | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskIaxCallnumberlimits(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskIaxCallnumberlimits``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskIaxCallnumberlimitsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IAXCallNumberLimitss**](IAXCallNumberLimitss.md) |  |

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

## UpdateAsteriskIaxGeneral

> UpdateAsteriskIaxGeneral(ctx).Body(body).Execute()

Update IAX general option

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
 body := *openapiclient.NewIAXGeneral() // IAXGeneral | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskIaxGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskIaxGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskIaxGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**IAXGeneral**](IAXGeneral.md) |  |

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

## UpdateAsteriskPjsipGlobal

> UpdateAsteriskPjsipGlobal(ctx).Body(body).Execute()

Update PJSIP section options

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
 body := *openapiclient.NewPJSIPGlobal() // PJSIPGlobal | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskPjsipGlobal(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskPjsipGlobal``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskPjsipGlobalRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PJSIPGlobal**](PJSIPGlobal.md) |  |

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

## UpdateAsteriskPjsipSystem

> UpdateAsteriskPjsipSystem(ctx).Body(body).Execute()

Update PJSIP section options

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
 body := *openapiclient.NewPJSIPSystem() // PJSIPSystem | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskPjsipSystem(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskPjsipSystem``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskPjsipSystemRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**PJSIPSystem**](PJSIPSystem.md) |  |

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

## UpdateAsteriskQueueGeneral

> UpdateAsteriskQueueGeneral(ctx).Body(body).Execute()

Update Queue general option

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
 body := *openapiclient.NewQueueGeneral() // QueueGeneral | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskQueueGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskQueueGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskQueueGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**QueueGeneral**](QueueGeneral.md) |  |

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

## UpdateAsteriskRtpGeneral

> UpdateAsteriskRtpGeneral(ctx).Body(body).Execute()

Update RTP general option

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
 body := *openapiclient.NewRTPConfiguration() // RTPConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskRtpGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskRtpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskRtpGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RTPConfiguration**](RTPConfiguration.md) |  |

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

## UpdateAsteriskRtpIceHostCandidates

> UpdateAsteriskRtpIceHostCandidates(ctx).Body(body).Execute()

Update RTP ice_host_candidates option

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
 body := *openapiclient.NewRTPConfiguration() // RTPConfiguration | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskRtpIceHostCandidates(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskRtpIceHostCandidates``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskRtpIceHostCandidatesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**RTPConfiguration**](RTPConfiguration.md) |  |

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

## UpdateAsteriskSccpGeneral

> UpdateAsteriskSccpGeneral(ctx).Body(body).Execute()

Update SCCP general option

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
 body := *openapiclient.NewSCCPGeneral() // SCCPGeneral | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskSccpGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskSccpGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskSccpGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**SCCPGeneral**](SCCPGeneral.md) |  |

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

## UpdateAsteriskVoicemailGeneral

> UpdateAsteriskVoicemailGeneral(ctx).Body(body).Execute()

Update Voicemail general option

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
 body := *openapiclient.NewVoicemailGeneral() // VoicemailGeneral | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskVoicemailGeneral(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskVoicemailGeneral``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskVoicemailGeneralRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailGeneral**](VoicemailGeneral.md) |  |

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

## UpdateAsteriskVoicemailZonemessages

> UpdateAsteriskVoicemailZonemessages(ctx).Body(body).Execute()

Update Voicemail zonemessages option

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
 body := *openapiclient.NewVoicemailZoneMessages() // VoicemailZoneMessages | 

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 r, err := apiClient.AsteriskAPI.UpdateAsteriskVoicemailZonemessages(context.Background()).Body(body).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `AsteriskAPI.UpdateAsteriskVoicemailZonemessages``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateAsteriskVoicemailZonemessagesRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **body** | [**VoicemailZoneMessages**](VoicemailZoneMessages.md) |  |

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
