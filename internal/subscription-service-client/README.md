# Go API client for openapi

## Create subscription

Create a new subscription.

```
curl -v -d '{\\
    \"company\": \"My Company\", \\
    \"firstName\": \"First\", \\
    \"lastName\": \"Last\", \\
    \"email\": \"first.last@magnolia-cms.com\", \\
    \"password\": \"some1%2Tres\", \\
    \"function\": \"CTO\", \\
    \"country\": \"Spain\" \\
    }' \\
-H \"Content-Type: application/json\" -X POST \"http://localhost:8080/public/subscriptions\"
```
## Update subscription

Update a new subscription. This is an idempotent operation.

```
curl -v -d '{\\
    \"id\": \"my-company\" \\
    }' \\
-H \"Content-Type: application/json\" -X POST \"http://localhost:8080/admin/subscriptions/{subId}\"
```
## Invite users

Invite a list of users to a subscription.

## Validate invitation

Validate an invitation to a subscription.

## Activate user

Activate a user for a subscription. The user needs an invitation to be activated.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.1.0
- Package version: 1.0.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen
For more information, please visit [https://docs.magnolia-cms.com/](https://docs.magnolia-cms.com/)

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import openapi "github.com/GIT_USER_ID/GIT_REPO_ID"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), openapi.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

### URLs Configuration per Operation

Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"{classname}Service.{nickname}"` string.
Similar rules for overriding default operation server index and variables applies by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```
ctx := context.WithValue(context.Background(), openapi.ContextOperationServerIndices, map[string]int{
	"{classname}Service.{nickname}": 2,
})
ctx = context.WithValue(context.Background(), openapi.ContextOperationServerVariables, map[string]map[string]string{
	"{classname}Service.{nickname}": {
		"port": "8443",
	},
})
```

## Documentation for API Endpoints

All URIs are relative to *http://localhost:8086*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*GroupApi* | [**ListGroupsOfSubscription**](docs/GroupApi.md#listgroupsofsubscription) | **Get** /subscriptions/{subscriptionId}/groups | List groups of a subscription
*MetricApi* | [**FindAllMetricsPublic**](docs/MetricApi.md#findallmetricspublic) | **Get** /metrics | Find all subscriptions
*MetricDataApi* | [**GetSubscriptionMetrics**](docs/MetricDataApi.md#getsubscriptionmetrics) | **Get** /subscriptions/{subscriptionId}/metrics | Get subscription metrics
*SubscriptionApi* | [**FindSubscriptionById**](docs/SubscriptionApi.md#findsubscriptionbyid) | **Get** /subscriptions/{subscriptionId} | Find a subscription by Id
*SubscriptionApi* | [**GetPlanForSubscription**](docs/SubscriptionApi.md#getplanforsubscription) | **Get** /subscriptions/{subscriptionId}/plan | Get current plan of the subscription
*SubscriptionApi* | [**UpdateOrganization**](docs/SubscriptionApi.md#updateorganization) | **Post** /subscriptions/{subscriptionId}/configuration/organization | Update organization of the subscription
*SubscriptionGitApi* | [**GetGitConfiguration**](docs/SubscriptionGitApi.md#getgitconfiguration) | **Get** /subscriptions/{subscriptionId}/git/configuration | Get Git configuration by subscription Id
*SubscriptionGitApi* | [**SyncGitRepository**](docs/SubscriptionGitApi.md#syncgitrepository) | **Post** /subscriptions/{subscriptionId}/git/sync | Sync Git repository
*SubscriptionGitApi* | [**UpdateGitConfiguration**](docs/SubscriptionGitApi.md#updategitconfiguration) | **Post** /subscriptions/{subscriptionId}/git/configuration | Update Git configuration for the subscription
*UserApi* | [**CreateInvitations**](docs/UserApi.md#createinvitations) | **Post** /subscriptions/{subscriptionId}/invitations | Invite users to a subscription
*UserApi* | [**ListUsersOfSubscription**](docs/UserApi.md#listusersofsubscription) | **Get** /subscriptions/{subscriptionId}/users | List users of a subscription
*UserSelfApi* | [**ChangePassword**](docs/UserSelfApi.md#changepassword) | **Post** /users/me/changePassword | Change current user password
*UserSelfApi* | [**FindSubscriptionsByCurrentUser**](docs/UserSelfApi.md#findsubscriptionsbycurrentuser) | **Get** /users/me/subscriptions | Find all subscriptions by current authenticated user
*UserSelfApi* | [**GetCurrentUser**](docs/UserSelfApi.md#getcurrentuser) | **Get** /users/me | Get current user
*UserSelfApi* | [**UpdateProfile**](docs/UserSelfApi.md#updateprofile) | **Put** /users/me/profile | Update user profile


## Documentation For Models

 - [ActivationRequest](docs/ActivationRequest.md)
 - [CanonicalSubscriptionRequest](docs/CanonicalSubscriptionRequest.md)
 - [CreatedSubscription](docs/CreatedSubscription.md)
 - [CreatedSubscriptionAllOf](docs/CreatedSubscriptionAllOf.md)
 - [GitConfigurationData](docs/GitConfigurationData.md)
 - [GitConfigurationPrivateResponse](docs/GitConfigurationPrivateResponse.md)
 - [GitConfigurationResponse](docs/GitConfigurationResponse.md)
 - [GitProvider](docs/GitProvider.md)
 - [Group](docs/Group.md)
 - [IdentityProviderData](docs/IdentityProviderData.md)
 - [InlineObject](docs/InlineObject.md)
 - [InvitationEmails](docs/InvitationEmails.md)
 - [Metric](docs/Metric.md)
 - [MetricAllOf](docs/MetricAllOf.md)
 - [MetricBasicInfo](docs/MetricBasicInfo.md)
 - [MetricConstraint](docs/MetricConstraint.md)
 - [MetricData](docs/MetricData.md)
 - [MetricDataAllOf](docs/MetricDataAllOf.md)
 - [MetricDetails](docs/MetricDetails.md)
 - [MetricRequest](docs/MetricRequest.md)
 - [MetricResponse](docs/MetricResponse.md)
 - [MetricSource](docs/MetricSource.md)
 - [MetricType](docs/MetricType.md)
 - [PasswordChangeRequest](docs/PasswordChangeRequest.md)
 - [Plan](docs/Plan.md)
 - [PlanHardLimitValidatorResponse](docs/PlanHardLimitValidatorResponse.md)
 - [PlanRequest](docs/PlanRequest.md)
 - [Subscription](docs/Subscription.md)
 - [SubscriptionAllOf](docs/SubscriptionAllOf.md)
 - [SubscriptionAllOfConfiguration](docs/SubscriptionAllOfConfiguration.md)
 - [SubscriptionAllOfDeploymentInfo](docs/SubscriptionAllOfDeploymentInfo.md)
 - [SubscriptionErrorCause](docs/SubscriptionErrorCause.md)
 - [SubscriptionStatus](docs/SubscriptionStatus.md)
 - [SubscriptionUpdate](docs/SubscriptionUpdate.md)
 - [UpdateGitConfigurationRequest](docs/UpdateGitConfigurationRequest.md)
 - [UpdateOrganizationRequest](docs/UpdateOrganizationRequest.md)
 - [UpdatePlanRequest](docs/UpdatePlanRequest.md)
 - [UpdateProfileRequest](docs/UpdateProfileRequest.md)
 - [UpdateUserRequest](docs/UpdateUserRequest.md)
 - [User](docs/User.md)
 - [UserCount](docs/UserCount.md)
 - [UserDetails](docs/UserDetails.md)


## Documentation For Authorization



### SubscriptionSecurity

- **Type**: HTTP basic authentication

Example

```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
    UserName: "username",
    Password: "password",
})
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author


