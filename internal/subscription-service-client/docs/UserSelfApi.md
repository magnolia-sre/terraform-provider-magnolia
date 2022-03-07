# \UserSelfApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**ChangePassword**](UserSelfApi.md#ChangePassword) | **Post** /users/me/changePassword | Change current user password
[**FindSubscriptionsByCurrentUser**](UserSelfApi.md#FindSubscriptionsByCurrentUser) | **Get** /users/me/subscriptions | Find all subscriptions by current authenticated user
[**GetCurrentUser**](UserSelfApi.md#GetCurrentUser) | **Get** /users/me | Get current user
[**UpdateProfile**](UserSelfApi.md#UpdateProfile) | **Put** /users/me/profile | Update user profile
[**UploadAvatar**](UserSelfApi.md#UploadAvatar) | **Put** /users/me/avatar | Upload user avatar



## ChangePassword

> ChangePassword(ctx).PasswordChangeRequest(passwordChangeRequest).Execute()

Change current user password



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
    passwordChangeRequest := *openapiclient.NewPasswordChangeRequest() // PasswordChangeRequest | Old and New Password values

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSelfApi.ChangePassword(context.Background()).PasswordChangeRequest(passwordChangeRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSelfApi.ChangePassword``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiChangePasswordRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **passwordChangeRequest** | [**PasswordChangeRequest**](PasswordChangeRequest.md) | Old and New Password values | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## FindSubscriptionsByCurrentUser

> []Subscription FindSubscriptionsByCurrentUser(ctx).Execute()

Find all subscriptions by current authenticated user



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
    resp, r, err := apiClient.UserSelfApi.FindSubscriptionsByCurrentUser(context.Background()).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSelfApi.FindSubscriptionsByCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `FindSubscriptionsByCurrentUser`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `UserSelfApi.FindSubscriptionsByCurrentUser`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiFindSubscriptionsByCurrentUserRequest struct via the builder pattern


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


## GetCurrentUser

> User GetCurrentUser(ctx).SubscriptionId(subscriptionId).Execute()

Get current user



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
    subscriptionId := "subscriptionId_example" // string | The id of the subscription the user is signed in.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSelfApi.GetCurrentUser(context.Background()).SubscriptionId(subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSelfApi.GetCurrentUser``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `GetCurrentUser`: User
    fmt.Fprintf(os.Stdout, "Response from `UserSelfApi.GetCurrentUser`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetCurrentUserRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **subscriptionId** | **string** | The id of the subscription the user is signed in. | 

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


## UpdateProfile

> []Subscription UpdateProfile(ctx).UpdateProfileRequest(updateProfileRequest).Execute()

Update user profile



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
    updateProfileRequest := *openapiclient.NewUpdateProfileRequest("FirstName_example", "LastName_example") // UpdateProfileRequest | User profile information

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSelfApi.UpdateProfile(context.Background()).UpdateProfileRequest(updateProfileRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSelfApi.UpdateProfile``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `UpdateProfile`: []Subscription
    fmt.Fprintf(os.Stdout, "Response from `UserSelfApi.UpdateProfile`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUpdateProfileRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **updateProfileRequest** | [**UpdateProfileRequest**](UpdateProfileRequest.md) | User profile information | 

### Return type

[**[]Subscription**](Subscription.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UploadAvatar

> UploadAvatar(ctx).File(file).Execute()

Upload user avatar



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
    file := os.NewFile(1234, "some_file") // *os.File |  (optional)

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserSelfApi.UploadAvatar(context.Background()).File(file).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserSelfApi.UploadAvatar``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiUploadAvatarRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **file** | ***os.File** |  | 

### Return type

 (empty response body)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: multipart/form-data
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

