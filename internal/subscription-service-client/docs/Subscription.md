# Subscription

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**FirstName** | **string** |  | 
**LastName** | **string** |  | 
**Email** | **string** |  | 
**Function** | **string** |  | 
**Id** | **string** |  | 
**Company** | **string** |  | 
**Creation** | Pointer to **string** |  | [optional] 
**Country** | **string** |  | 
**Type** | **string** |  | 
**Alias** | Pointer to **string** |  | [optional] 
**PlanId** | Pointer to **string** |  | [optional] 
**Provisioned** | Pointer to **bool** |  | [optional] 
**Status** | Pointer to [**SubscriptionStatus**](SubscriptionStatus.md) |  | [optional] 
**DeploymentInfo** | Pointer to [**SubscriptionAllOfDeploymentInfo**](SubscriptionAllOfDeploymentInfo.md) |  | [optional] 
**Configuration** | Pointer to [**SubscriptionAllOfConfiguration**](SubscriptionAllOfConfiguration.md) |  | [optional] 
**ErrorCause** | Pointer to [**SubscriptionErrorCause**](SubscriptionErrorCause.md) |  | [optional] 

## Methods

### NewSubscription

`func NewSubscription(firstName string, lastName string, email string, function string, id string, company string, country string, type_ string, ) *Subscription`

NewSubscription instantiates a new Subscription object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionWithDefaults

`func NewSubscriptionWithDefaults() *Subscription`

NewSubscriptionWithDefaults instantiates a new Subscription object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetFirstName

`func (o *Subscription) GetFirstName() string`

GetFirstName returns the FirstName field if non-nil, zero value otherwise.

### GetFirstNameOk

`func (o *Subscription) GetFirstNameOk() (*string, bool)`

GetFirstNameOk returns a tuple with the FirstName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFirstName

`func (o *Subscription) SetFirstName(v string)`

SetFirstName sets FirstName field to given value.


### GetLastName

`func (o *Subscription) GetLastName() string`

GetLastName returns the LastName field if non-nil, zero value otherwise.

### GetLastNameOk

`func (o *Subscription) GetLastNameOk() (*string, bool)`

GetLastNameOk returns a tuple with the LastName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLastName

`func (o *Subscription) SetLastName(v string)`

SetLastName sets LastName field to given value.


### GetEmail

`func (o *Subscription) GetEmail() string`

GetEmail returns the Email field if non-nil, zero value otherwise.

### GetEmailOk

`func (o *Subscription) GetEmailOk() (*string, bool)`

GetEmailOk returns a tuple with the Email field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmail

`func (o *Subscription) SetEmail(v string)`

SetEmail sets Email field to given value.


### GetFunction

`func (o *Subscription) GetFunction() string`

GetFunction returns the Function field if non-nil, zero value otherwise.

### GetFunctionOk

`func (o *Subscription) GetFunctionOk() (*string, bool)`

GetFunctionOk returns a tuple with the Function field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetFunction

`func (o *Subscription) SetFunction(v string)`

SetFunction sets Function field to given value.


### GetId

`func (o *Subscription) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Subscription) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Subscription) SetId(v string)`

SetId sets Id field to given value.


### GetCompany

`func (o *Subscription) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *Subscription) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *Subscription) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetCreation

`func (o *Subscription) GetCreation() string`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *Subscription) GetCreationOk() (*string, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *Subscription) SetCreation(v string)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *Subscription) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### GetCountry

`func (o *Subscription) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *Subscription) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *Subscription) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetType

`func (o *Subscription) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Subscription) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Subscription) SetType(v string)`

SetType sets Type field to given value.


### GetAlias

`func (o *Subscription) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *Subscription) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *Subscription) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *Subscription) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPlanId

`func (o *Subscription) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *Subscription) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *Subscription) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *Subscription) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProvisioned

`func (o *Subscription) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *Subscription) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *Subscription) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *Subscription) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetStatus

`func (o *Subscription) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Subscription) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Subscription) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Subscription) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeploymentInfo

`func (o *Subscription) GetDeploymentInfo() SubscriptionAllOfDeploymentInfo`

GetDeploymentInfo returns the DeploymentInfo field if non-nil, zero value otherwise.

### GetDeploymentInfoOk

`func (o *Subscription) GetDeploymentInfoOk() (*SubscriptionAllOfDeploymentInfo, bool)`

GetDeploymentInfoOk returns a tuple with the DeploymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInfo

`func (o *Subscription) SetDeploymentInfo(v SubscriptionAllOfDeploymentInfo)`

SetDeploymentInfo sets DeploymentInfo field to given value.

### HasDeploymentInfo

`func (o *Subscription) HasDeploymentInfo() bool`

HasDeploymentInfo returns a boolean if a field has been set.

### GetConfiguration

`func (o *Subscription) GetConfiguration() SubscriptionAllOfConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *Subscription) GetConfigurationOk() (*SubscriptionAllOfConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *Subscription) SetConfiguration(v SubscriptionAllOfConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *Subscription) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetErrorCause

`func (o *Subscription) GetErrorCause() SubscriptionErrorCause`

GetErrorCause returns the ErrorCause field if non-nil, zero value otherwise.

### GetErrorCauseOk

`func (o *Subscription) GetErrorCauseOk() (*SubscriptionErrorCause, bool)`

GetErrorCauseOk returns a tuple with the ErrorCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCause

`func (o *Subscription) SetErrorCause(v SubscriptionErrorCause)`

SetErrorCause sets ErrorCause field to given value.

### HasErrorCause

`func (o *Subscription) HasErrorCause() bool`

HasErrorCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


