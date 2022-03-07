# \SubscriptionGitApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetGitConfiguration**](SubscriptionGitApi.md#GetGitConfiguration) | **Get** /subscriptions/{subscriptionId}/git/configuration | Get Git configuration by subscription Id
[**SyncGitRepository**](SubscriptionGitApi.md#SyncGitRepository) | **Post** /subscriptions/{subscriptionId}/git/sync | Sync Git repository
[**UpdateGitConfiguration**](SubscriptionGitApi.md#UpdateGitConfiguration) | **Post** /subscriptions/{subscriptionId}/git/configuration | Update Git configuration for the subscription



## GetGitConfiguration

> GitConfigurationResponse GetGitConfiguration(ctx, subscriptionId).Execute()

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
    resp, r, err := apiClient.SubscriptionGitApi.GetGitConfiguration(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGitApi.GetGitConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetGitConfiguration`: GitConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGitApi.GetGitConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetGitConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GitConfigurationResponse**](GitConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## SyncGitRepository

> SyncGitRepository(ctx, subscriptionId).Execute()

Sync Git repository



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
    resp, r, err := apiClient.SubscriptionGitApi.SyncGitRepository(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGitApi.SyncGitRepository``: %v\n", err)
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

Other parameters are passed through a pointer to a apiSyncGitRepositoryRequest struct via the builder pattern


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


## UpdateGitConfiguration

> GitConfigurationResponse UpdateGitConfiguration(ctx, subscriptionId).UpdateGitConfigurationRequest(updateGitConfigurationRequest).Execute()

Update Git configuration for the subscription



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
    updateGitConfigurationRequest := *openapiclient.NewUpdateGitConfigurationRequest("GitCloneUrl_example", openapiclient.GitProvider("GITHUB")) // UpdateGitConfigurationRequest | Git configuration

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.SubscriptionGitApi.UpdateGitConfiguration(context.Background(), subscriptionId).UpdateGitConfigurationRequest(updateGitConfigurationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `SubscriptionGitApi.UpdateGitConfiguration``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateGitConfiguration`: GitConfigurationResponse
    fmt.Fprintf(os.Stdout, "Response from `SubscriptionGitApi.UpdateGitConfiguration`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateGitConfigurationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateGitConfigurationRequest** | [**UpdateGitConfigurationRequest**](UpdateGitConfigurationRequest.md) | Git configuration | 

### Return type

[**GitConfigurationResponse**](GitConfigurationResponse.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

