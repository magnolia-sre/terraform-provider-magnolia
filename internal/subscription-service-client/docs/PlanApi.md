# \PlanApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivatePlan**](PlanApi.md#ActivatePlan) | **Post** /admin/plan/{planId}/activate | Activate plan
[**CreatePlan**](PlanApi.md#CreatePlan) | **Post** /admin/plan | Create new plan
[**DeactivatePlan**](PlanApi.md#DeactivatePlan) | **Post** /admin/plan/{planId}/deactivate | Deactivate plan
[**DeletePlan**](PlanApi.md#DeletePlan) | **Delete** /admin/plan/{planId} | Delete plan by Id
[**FindAllPlans**](PlanApi.md#FindAllPlans) | **Get** /admin/plan | Find all plans
[**FindById**](PlanApi.md#FindById) | **Get** /admin/plan/{planId} | Find plan by Id
[**UpdatePlan**](PlanApi.md#UpdatePlan) | **Put** /admin/plan/{planId} | Update plan
[**UpdateToDefaultForNewSubscription**](PlanApi.md#UpdateToDefaultForNewSubscription) | **Post** /admin/plan/{planId}/default | Update plan to be default plan



## ActivatePlan

> Plan ActivatePlan(ctx, planId).Execute()

Activate plan



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
    planId := "planId_example" // string | Id of subscription plan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.ActivatePlan(context.Background(), planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.ActivatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.ActivatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CreatePlan

> Plan CreatePlan(ctx).PlanRequest(planRequest).Execute()

Create new plan



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
    planRequest := *openapiclient.NewPlanRequest("Name_example") // PlanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.CreatePlan(context.Background()).PlanRequest(planRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.CreatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.CreatePlan`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **planRequest** | [**PlanRequest**](PlanRequest.md) |  | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeactivatePlan

> Plan DeactivatePlan(ctx, planId).Execute()

Deactivate plan



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
    planId := "planId_example" // string | Id of subscription plan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DeactivatePlan(context.Background(), planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DeactivatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `DeactivatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.DeactivatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeactivatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## DeletePlan

> DeletePlan(ctx, planId).Execute()

Delete plan by Id



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
    planId := "planId_example" // string | Id of subscription plan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.DeletePlan(context.Background(), planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.DeletePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeletePlanRequest struct via the builder pattern


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


## FindAllPlans

> []Plan FindAllPlans(ctx).Execute()

Find all plans



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
    resp, r, err := apiClient.PlanApi.FindAllPlans(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.FindAllPlans``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindAllPlans`: []Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.FindAllPlans`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindAllPlansRequest struct via the builder pattern


### Return type

[**[]Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindById

> Plan FindById(ctx, planId).Execute()

Find plan by Id



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
    planId := "planId_example" // string | Id of subscription plan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.FindById(context.Background(), planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.FindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindById`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.FindById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiFindByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdatePlan

> Plan UpdatePlan(ctx, planId).PlanRequest(planRequest).Execute()

Update plan



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
    planId := "planId_example" // string | Id of subscription plan
    planRequest := *openapiclient.NewPlanRequest("Name_example") // PlanRequest |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.UpdatePlan(context.Background(), planId).PlanRequest(planRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.UpdatePlan``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdatePlan`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.UpdatePlan`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdatePlanRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **planRequest** | [**PlanRequest**](PlanRequest.md) |  | 

### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateToDefaultForNewSubscription

> Plan UpdateToDefaultForNewSubscription(ctx, planId).Execute()

Update plan to be default plan



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
    planId := "planId_example" // string | Id of subscription plan

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PlanApi.UpdateToDefaultForNewSubscription(context.Background(), planId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PlanApi.UpdateToDefaultForNewSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateToDefaultForNewSubscription`: Plan
    fmt.Fprintf(os.Stdout, "Response from `PlanApi.UpdateToDefaultForNewSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**planId** | **string** | Id of subscription plan | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateToDefaultForNewSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Plan**](Plan.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

