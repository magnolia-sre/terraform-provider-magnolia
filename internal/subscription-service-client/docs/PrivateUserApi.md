# \PrivateUserApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ActivateUser**](PrivateUserApi.md#ActivateUser) | **Post** /private/activation/{externalId} | 
[**CountUsers**](PrivateUserApi.md#CountUsers) | **Get** /private/subscriptions/{subscriptionId}/users/count | Count users
[**ValidateInvitation**](PrivateUserApi.md#ValidateInvitation) | **Post** /private/invitations | 



## ActivateUser

> User ActivateUser(ctx, externalId).ActivationRequest(activationRequest).Execute()





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
    externalId := "externalId_example" // string | 
    activationRequest := *openapiclient.NewActivationRequest("FirstName_example", "LastName_example", "Function_example", "Password_example") // ActivationRequest | User information

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateUserApi.ActivateUser(context.Background(), externalId).ActivationRequest(activationRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateUserApi.ActivateUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ActivateUser`: User
    fmt.Fprintf(os.Stdout, "Response from `PrivateUserApi.ActivateUser`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**externalId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiActivateUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **activationRequest** | [**ActivationRequest**](ActivationRequest.md) | User information | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## CountUsers

> UserCount CountUsers(ctx, subscriptionId).Execute()

Count users



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
    resp, r, err := apiClient.PrivateUserApi.CountUsers(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateUserApi.CountUsers``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CountUsers`: UserCount
    fmt.Fprintf(os.Stdout, "Response from `PrivateUserApi.CountUsers`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiCountUsersRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**UserCount**](UserCount.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ValidateInvitation

> User ValidateInvitation(ctx).Token(token).Execute()





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
    token := "token_example" // string | Token sent by email in the user invitation.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.PrivateUserApi.ValidateInvitation(context.Background()).Token(token).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `PrivateUserApi.ValidateInvitation``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ValidateInvitation`: User
    fmt.Fprintf(os.Stdout, "Response from `PrivateUserApi.ValidateInvitation`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiValidateInvitationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **token** | **string** | Token sent by email in the user invitation. | 

### Return type

[**User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

