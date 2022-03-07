# \PrivateMetricApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ValidateHardLimit**](PrivateMetricApi.md#ValidateHardLimit) | **Get** /private/subscriptions/{subscriptionId}/metrics/validate | Validate hard limit of given metric against current plan of the subscription



## ValidateHardLimit

> PlanHardLimitValidatorResponse ValidateHardLimit(ctx, subscriptionId).Name(name).Execute()

Validate hard limit of given metric against current plan of the subscription



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
    subscriptionId := "subscriptionId_example" // string | Generated id of subscription
    name := "name_example" // string | Metric name (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateMetricApi.ValidateHardLimit(context.Background(), subscriptionId).Name(name).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateMetricApi.ValidateHardLimit``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateHardLimit`: PlanHardLimitValidatorResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateMetricApi.ValidateHardLimit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiValidateHardLimitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **name** | **string** | Metric name | 

### Return type

[**PlanHardLimitValidatorResponse**](PlanHardLimitValidatorResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

