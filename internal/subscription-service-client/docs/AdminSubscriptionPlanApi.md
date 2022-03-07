# \AdminSubscriptionPlanApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**UpdateSubscriptionPlan**](AdminSubscriptionPlanApi.md#UpdateSubscriptionPlan) | **Post** /admin/subscriptions/{subscriptionId}/plan | Update plan of the subscription



## UpdateSubscriptionPlan

> Subscription UpdateSubscriptionPlan(ctx, subscriptionId).UpdatePlanRequest(updatePlanRequest).Execute()

Update plan of the subscription



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
    updatePlanRequest := *openapiclient.NewUpdatePlanRequest("PlanId_example") // UpdatePlanRequest | Subscription plan info (e.g planId)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminSubscriptionPlanApi.UpdateSubscriptionPlan(context.Background(), subscriptionId).UpdatePlanRequest(updatePlanRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionPlanApi.UpdateSubscriptionPlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateSubscriptionPlan`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionPlanApi.UpdateSubscriptionPlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateSubscriptionPlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updatePlanRequest** | [**UpdatePlanRequest**](UpdatePlanRequest.md) | Subscription plan info (e.g planId) | 

### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

