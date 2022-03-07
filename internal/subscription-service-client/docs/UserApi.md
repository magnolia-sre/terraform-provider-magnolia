# \UserApi

All URIs are relative to *http://localhost:8086*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateInvitations**](UserApi.md#CreateInvitations) | **Post** /subscriptions/{subscriptionId}/invitations | Invite users to a subscription
[**ListUsersOfSubscription**](UserApi.md#ListUsersOfSubscription) | **Get** /subscriptions/{subscriptionId}/users | List users of a subscription



## CreateInvitations

> []User CreateInvitations(ctx, subscriptionId).InvitationEmails(invitationEmails).Execute()

Invite users to a subscription



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
    invitationEmails := *openapiclient.NewInvitationEmails([]string{"Emails_example"}, "Group_example") // InvitationEmails | List of emails to invite and a group to associate them with.

    configuration := openapiclient.NewConfiguration()
    apiClient := openapiclient.NewAPIClient(configuration)
    resp, r, err := apiClient.UserApi.CreateInvitations(context.Background(), subscriptionId).InvitationEmails(invitationEmails).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.CreateInvitations``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `CreateInvitations`: []User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.CreateInvitations`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateInvitationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **invitationEmails** | [**InvitationEmails**](InvitationEmails.md) | List of emails to invite and a group to associate them with. | 

### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ListUsersOfSubscription

> []User ListUsersOfSubscription(ctx, subscriptionId).Execute()

List users of a subscription



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
    resp, r, err := apiClient.UserApi.ListUsersOfSubscription(context.Background(), subscriptionId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `UserApi.ListUsersOfSubscription``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
    }
    // response from `ListUsersOfSubscription`: []User
    fmt.Fprintf(os.Stdout, "Response from `UserApi.ListUsersOfSubscription`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**subscriptionId** | **string** | Generated id of subscription | 

### Other Parameters

Other parameters are passed through a pointer to a apiListUsersOfSubscriptionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]User**](User.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

