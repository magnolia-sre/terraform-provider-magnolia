# \PrivateSubscriptionApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateCanonicalSubscription**](PrivateSubscriptionApi.md#CreateCanonicalSubscription) | **Post** /private/subscriptions/canonical | Create a new canonical subscription
[**PrivateCreateSubscription**](PrivateSubscriptionApi.md#PrivateCreateSubscription) | **Post** /private/subscriptions | Create a new subscription
[**PrivateDeleteSubscription**](PrivateSubscriptionApi.md#PrivateDeleteSubscription) | **Delete** /private/subscriptions/{subscriptionId} | Delete a subscription by Id
[**PrivateFindAllSubscriptions**](PrivateSubscriptionApi.md#PrivateFindAllSubscriptions) | **Get** /private/subscriptions | Find all subscriptions
[**PrivateFindSubscriptionById**](PrivateSubscriptionApi.md#PrivateFindSubscriptionById) | **Get** /private/subscriptions/{subscriptionId} | Find a subscription by Id
[**PrivateGetGitConfiguration**](PrivateSubscriptionApi.md#PrivateGetGitConfiguration) | **Get** /private/subscriptions/{subscriptionId}/git/configuration | Get Git configuration by subscription Id



## CreateCanonicalSubscription

> Subscription CreateCanonicalSubscription(ctx).CanonicalSubscriptionRequest(canonicalSubscriptionRequest).Execute()

Create a new canonical subscription



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
    canonicalSubscriptionRequest := *openapiclient.NewCanonicalSubscriptionRequest("Id_example") // CanonicalSubscriptionRequest | Canonical subscription request

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateSubscriptionApi.CreateCanonicalSubscription(context.Background()).CanonicalSubscriptionRequest(canonicalSubscriptionRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.CreateCanonicalSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateCanonicalSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionApi.CreateCanonicalSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateCanonicalSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **canonicalSubscriptionRequest** | [**CanonicalSubscriptionRequest**](CanonicalSubscriptionRequest.md) | Canonical subscription request | 

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


## PrivateCreateSubscription

> Subscription PrivateCreateSubscription(ctx).NewSubscription(newSubscription).Execute()

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
    resp, r, err := apiClient.PrivateSubscriptionApi.PrivateCreateSubscription(context.Background()).NewSubscription(newSubscription).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.PrivateCreateSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateCreateSubscription`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionApi.PrivateCreateSubscription`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPrivateCreateSubscriptionRequest struct via the builder pattern


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


## PrivateDeleteSubscription

> PrivateDeleteSubscription(ctx, subscriptionId).Execute()

Delete a subscription by Id



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
    resp, r, err := apiClient.PrivateSubscriptionApi.PrivateDeleteSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.PrivateDeleteSubscription``: %v\n", err)
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

Other parameters are passed through a pointer to a apiPrivateDeleteSubscriptionRequest struct via the builder pattern


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


## PrivateFindAllSubscriptions

> []Subscription PrivateFindAllSubscriptions(ctx).Execute()

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
    resp, r, err := apiClient.PrivateSubscriptionApi.PrivateFindAllSubscriptions(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.PrivateFindAllSubscriptions``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateFindAllSubscriptions`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionApi.PrivateFindAllSubscriptions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateFindAllSubscriptionsRequest struct via the builder pattern


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


## PrivateFindSubscriptionById

> Subscription PrivateFindSubscriptionById(ctx, subscriptionId).Execute()

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
    resp, r, err := apiClient.PrivateSubscriptionApi.PrivateFindSubscriptionById(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.PrivateFindSubscriptionById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateFindSubscriptionById`: Subscription
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionApi.PrivateFindSubscriptionById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateFindSubscriptionByIdRequest struct via the builder pattern


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


## PrivateGetGitConfiguration

> GitConfigurationPrivateResponse PrivateGetGitConfiguration(ctx, subscriptionId).Execute()

Get Git configuration by subscription Id



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
    resp, r, err := apiClient.PrivateSubscriptionApi.PrivateGetGitConfiguration(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateSubscriptionApi.PrivateGetGitConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `PrivateGetGitConfiguration`: GitConfigurationPrivateResponse
    fmt.Fprintf(os.Stdout, "Response from `PrivateSubscriptionApi.PrivateGetGitConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiPrivateGetGitConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitConfigurationPrivateResponse**](GitConfigurationPrivateResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

