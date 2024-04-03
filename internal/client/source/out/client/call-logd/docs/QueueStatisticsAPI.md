# \QueueStatisticsAPI

All URIs are relative to *<http://api.accentvoice.io/1.0>*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetQueueQoSStatistics**](QueueStatisticsAPI.md#GetQueueQoSStatistics) | **Get** /queues/{queue_id}/statistics/qos | QoS statistics for a specific queue
[**GetQueueStatistics**](QueueStatisticsAPI.md#GetQueueStatistics) | **Get** /queues/{queue_id}/statistics | Statistics for a specific queue
[**GetQueuesStatistics**](QueueStatisticsAPI.md#GetQueuesStatistics) | **Get** /queues/statistics | Statistics for all queues

## GetQueueQoSStatistics

> QueueQoSStatistics GetQueueQoSStatistics(ctx, queueId).AccentTenant(accentTenant).From(from).Until(until).Interval(interval).QosThresholds(qosThresholds).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()

QoS statistics for a specific queue

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
)

func main() {
 queueId := int32(56) // int32 | ID of the queue.
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 interval := "interval_example" // string | Aggregation interval. An empty value means no interval, so an aggregation on all values. (optional)
 qosThresholds := []int32{int32(123)} // []int32 | The steps of quality of service times used for the interval generation. (optional)
 dayStartTime := "dayStartTime_example" // string | The time at which a day starts, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 dayEndTime := "dayEndTime_example" // string | The time at which a day ends, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 weekDays := []int32{int32(123)} // []int32 | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). (optional) (default to [1,2,3,4,5,6,7])
 timezone := "timezone_example" // string | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the <a href=\"https://en.wikipedia.org/wiki/Tz_database\">Time Zone Database</a> version installed on the server.  (optional) (default to "UTC")

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.QueueStatisticsAPI.GetQueueQoSStatistics(context.Background(), queueId).AccentTenant(accentTenant).From(from).Until(until).Interval(interval).QosThresholds(qosThresholds).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `QueueStatisticsAPI.GetQueueQoSStatistics``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetQueueQoSStatistics`: QueueQoSStatistics
 fmt.Fprintf(os.Stdout, "Response from `QueueStatisticsAPI.GetQueueQoSStatistics`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | ID of the queue. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueQoSStatisticsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **interval** | **string** | Aggregation interval. An empty value means no interval, so an aggregation on all values. |
 **qosThresholds** | **[]int32** | The steps of quality of service times used for the interval generation. |
 **dayStartTime** | **string** | The time at which a day starts, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **dayEndTime** | **string** | The time at which a day ends, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **weekDays** | **[]int32** | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). | [default to [1,2,3,4,5,6,7]]
 **timezone** | **string** | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/Tz_database\&quot;&gt;Time> Zone Database&lt;/a&gt; version installed on the server.  | [default to &quot;UTC&quot;]

### Return type

[**QueueQoSStatistics**](QueueQoSStatistics.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetQueueStatistics

> QueueStatistics GetQueueStatistics(ctx, queueId).AccentTenant(accentTenant).From(from).Until(until).Interval(interval).QosThreshold(qosThreshold).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()

Statistics for a specific queue

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
)

func main() {
 queueId := int32(56) // int32 | ID of the queue.
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 interval := "interval_example" // string | Aggregation interval. An empty value means no interval, so an aggregation on all values. (optional)
 qosThreshold := int32(56) // int32 | The number of seconds representing a good quality of service. (optional)
 dayStartTime := "dayStartTime_example" // string | The time at which a day starts, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 dayEndTime := "dayEndTime_example" // string | The time at which a day ends, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 weekDays := []int32{int32(123)} // []int32 | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). (optional) (default to [1,2,3,4,5,6,7])
 timezone := "timezone_example" // string | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the <a href=\"https://en.wikipedia.org/wiki/Tz_database\">Time Zone Database</a> version installed on the server.  (optional) (default to "UTC")

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.QueueStatisticsAPI.GetQueueStatistics(context.Background(), queueId).AccentTenant(accentTenant).From(from).Until(until).Interval(interval).QosThreshold(qosThreshold).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `QueueStatisticsAPI.GetQueueStatistics``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetQueueStatistics`: QueueStatistics
 fmt.Fprintf(os.Stdout, "Response from `QueueStatisticsAPI.GetQueueStatistics`: %v\n", resp)
}
```

### Path Parameters

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**queueId** | **int32** | ID of the queue. |

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueueStatisticsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **interval** | **string** | Aggregation interval. An empty value means no interval, so an aggregation on all values. |
 **qosThreshold** | **int32** | The number of seconds representing a good quality of service. |
 **dayStartTime** | **string** | The time at which a day starts, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **dayEndTime** | **string** | The time at which a day ends, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **weekDays** | **[]int32** | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). | [default to [1,2,3,4,5,6,7]]
 **timezone** | **string** | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/Tz_database\&quot;&gt;Time> Zone Database&lt;/a&gt; version installed on the server.  | [default to &quot;UTC&quot;]

### Return type

[**QueueStatistics**](QueueStatistics.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

## GetQueuesStatistics

> QueuesStatistics GetQueuesStatistics(ctx).AccentTenant(accentTenant).From(from).Until(until).QosThreshold(qosThreshold).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()

Statistics for all queues

### Example

```go
package main

import (
 "context"
 "fmt"
 "os"
    "time"
 openapiclient "github.com/ryanwclark1/accent-ui2/internal/client/calllogd"
)

func main() {
 accentTenant := "accentTenant_example" // string | The tenant's UUID, defining the ownership of a given resource. (optional)
 from := time.Now() // time.Time | Ignore calls before the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  (optional)
 until := time.Now() // time.Time | Ignore calls starting at or after the given date. Format is <a href=\"https://en.wikipedia.org/wiki/ISO_8601\">ISO-8601</a>. Timezone will be converted according to the `timezone` parameter. If missing, the statistics will include the current day.  (optional)
 qosThreshold := int32(56) // int32 | The number of seconds representing a good quality of service. (optional)
 dayStartTime := "dayStartTime_example" // string | The time at which a day starts, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 dayEndTime := "dayEndTime_example" // string | The time at which a day ends, inclusively. Accepted format is `HH:MM`, minutes are ignored. (optional)
 weekDays := []int32{int32(123)} // []int32 | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). (optional) (default to [1,2,3,4,5,6,7])
 timezone := "timezone_example" // string | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the <a href=\"https://en.wikipedia.org/wiki/Tz_database\">Time Zone Database</a> version installed on the server.  (optional) (default to "UTC")

 configuration := openapiclient.NewConfiguration()
 apiClient := openapiclient.NewAPIClient(configuration)
 resp, r, err := apiClient.QueueStatisticsAPI.GetQueuesStatistics(context.Background()).AccentTenant(accentTenant).From(from).Until(until).QosThreshold(qosThreshold).DayStartTime(dayStartTime).DayEndTime(dayEndTime).WeekDays(weekDays).Timezone(timezone).Execute()
 if err != nil {
  fmt.Fprintf(os.Stderr, "Error when calling `QueueStatisticsAPI.GetQueuesStatistics``: %v\n", err)
  fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
 }
 // response from `GetQueuesStatistics`: QueuesStatistics
 fmt.Fprintf(os.Stdout, "Response from `QueueStatisticsAPI.GetQueuesStatistics`: %v\n", resp)
}
```

### Path Parameters

### Other Parameters

Other parameters are passed through a pointer to a apiGetQueuesStatisticsRequest struct via the builder pattern

Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **accentTenant** | **string** | The tenant&#39;s UUID, defining the ownership of a given resource. |
 **from** | **time.Time** | Ignore calls before the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will start at the oldest available call with timezone UTC.  |
 **until** | **time.Time** | Ignore calls starting at or after the given date. Format is &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/ISO_8601\&quot;&gt;ISO-8601&lt;/a>&gt;. Timezone will be converted according to the &#x60;timezone&#x60; parameter. If missing, the statistics will include the current day.  |
 **qosThreshold** | **int32** | The number of seconds representing a good quality of service. |
 **dayStartTime** | **string** | The time at which a day starts, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **dayEndTime** | **string** | The time at which a day ends, inclusively. Accepted format is &#x60;HH:MM&#x60;, minutes are ignored. |
 **weekDays** | **[]int32** | The days of the week that should be included. A week starts on Monday (1) and ends on Sunday (7). | [default to [1,2,3,4,5,6,7]]
 **timezone** | **string** | Name of the timezone to use for dates and times. Example: America/New_York. Valid timezones are defined by the &lt;a href&#x3D;\&quot;<https://en.wikipedia.org/wiki/Tz_database\&quot;&gt;Time> Zone Database&lt;/a&gt; version installed on the server.  | [default to &quot;UTC&quot;]

### Return type

[**QueuesStatistics**](QueuesStatistics.md)

### Authorization

[accent_auth_token](../README.md#accent_auth_token)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)
