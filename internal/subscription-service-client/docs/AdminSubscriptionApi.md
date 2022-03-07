# \AdminSubscriptionApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**AdminFindSubscriptionById**](AdminSubscriptionApi.md#AdminFindSubscriptionById) | **Get** /admin/subscriptions/{subscriptionId} | Find a subscription by Id
[**CreateSubscription**](AdminSubscriptionApi.md#CreateSubscription) | **Post** /admin/subscriptions | Create a new subscription
[**DeleteSubscription**](AdminSubscriptionApi.md#DeleteSubscription) | **Delete** /admin/subscriptions/{subscriptionId} | Deletes a subscription
[**FindAllSubscriptions**](AdminSubscriptionApi.md#FindAllSubscriptions) | **Get** /admin/subscriptions | Find all subscriptions



## AdminFindSubscriptionById

> Subscription AdminFindSubscriptionById(ctx, subscriptionId).Execute()

Find a subscription by Id



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
    resp, r, err := apiClient.AdminSubscriptionApi.AdminFindSubscriptionById(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionApi.AdminFindSubscriptionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `AdminFindSubscriptionById`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionApi.AdminFindSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiAdminFindSubscriptionByIdRequest struct via the builder pattern


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


## CreateSubscription

> Subscription CreateSubscription(ctx).NewSubscription(newSubscription).Execute()

Create a new subscription



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
    newSubscription := *openapiclient.NewNewSubscription("FirstName_example", "LastName_example", "Email_example", "Function_example", "Company_example", "Password_example", "Country_example") // NewSubscription | Subscription and owner information

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.AdminSubscriptionApi.CreateSubscription(context.Background()).NewSubscription(newSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionApi.CreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionApi.CreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **newSubscription** | [**NewSubscription**](NewSubscription.md) | Subscription and owner information | 

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


## DeleteSubscription

> DeleteSubscription(ctx, subscriptionId).Execute()

Deletes a subscription



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
    resp, r, err := apiClient.AdminSubscriptionApi.DeleteSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionApi.DeleteSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindAllSubscriptions

> []Subscription FindAllSubscriptions(ctx).Execute()

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
    resp, r, err := apiClient.AdminSubscriptionApi.FindAllSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `AdminSubscriptionApi.FindAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `AdminSubscriptionApi.FindAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllSubscriptionsRequest struct via the builder pattern


### Return type

[**[]Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

