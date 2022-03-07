# \PrivateSubscriptionUpdateApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**PrivateUpdateSubscription**](PrivateSubscriptionUpdateApi.md#PrivateUpdateSubscription) | **Post** /private/subscriptions/{subscriptionId}/migrations | Update and deploy subscription



## PrivateUpdateSubscription

> Subscription PrivateUpdateSubscription(ctx, subscriptionId).Execute()

Update and deploy subscription



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

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateSubscriptionUpdateApi.PrivateUpdateSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionUpdateApi.PrivateUpdateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateUpdateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionUpdateApi.PrivateUpdateSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateUpdateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

