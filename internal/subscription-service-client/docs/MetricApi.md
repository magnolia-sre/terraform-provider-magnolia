# \MetricApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**FindAllMetricsPublic**](MetricApi.md#FindAllMetricsPublic) | **Get** /metrics | Find all subscriptions



## FindAllMetricsPublic

> []Metric FindAllMetricsPublic(ctx).Execute()

Find all subscriptions



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"
    openapiclient "./openapi"
)

func main() {

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.MetricApi.FindAllMetricsPublic(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `MetricApi.FindAllMetricsPublic``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllMetricsPublic`: []Metric
    fmt.Fprintf(os.Stdout, "Response from `MetricApi.FindAllMetricsPublic`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllMetricsPublicRequest struct via the builder pattern


### Return type

[**[]Metric**](Metric.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

