# SubscriptionAllOf

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
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

### NewSubscriptionAllOf

`func NewSubscriptionAllOf(id string, company string, country string, type_ string, ) *SubscriptionAllOf`

NewSubscriptionAllOf instantiates a new SubscriptionAllOf object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionAllOfWithDefaults

`func NewSubscriptionAllOfWithDefaults() *SubscriptionAllOf`

NewSubscriptionAllOfWithDefaults instantiates a new SubscriptionAllOf object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *SubscriptionAllOf) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *SubscriptionAllOf) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *SubscriptionAllOf) SetId(v string)`

SetId sets Id field to given value.


### GetCompany

`func (o *SubscriptionAllOf) GetCompany() string`

GetCompany returns the Company field if non-nil, zero value otherwise.

### GetCompanyOk

`func (o *SubscriptionAllOf) GetCompanyOk() (*string, bool)`

GetCompanyOk returns a tuple with the Company field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCompany

`func (o *SubscriptionAllOf) SetCompany(v string)`

SetCompany sets Company field to given value.


### GetCreation

`func (o *SubscriptionAllOf) GetCreation() string`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *SubscriptionAllOf) GetCreationOk() (*string, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *SubscriptionAllOf) SetCreation(v string)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *SubscriptionAllOf) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### GetCountry

`func (o *SubscriptionAllOf) GetCountry() string`

GetCountry returns the Country field if non-nil, zero value otherwise.

### GetCountryOk

`func (o *SubscriptionAllOf) GetCountryOk() (*string, bool)`

GetCountryOk returns a tuple with the Country field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCountry

`func (o *SubscriptionAllOf) SetCountry(v string)`

SetCountry sets Country field to given value.


### GetType

`func (o *SubscriptionAllOf) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *SubscriptionAllOf) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *SubscriptionAllOf) SetType(v string)`

SetType sets Type field to given value.


### GetAlias

`func (o *SubscriptionAllOf) GetAlias() string`

GetAlias returns the Alias field if non-nil, zero value otherwise.

### GetAliasOk

`func (o *SubscriptionAllOf) GetAliasOk() (*string, bool)`

GetAliasOk returns a tuple with the Alias field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAlias

`func (o *SubscriptionAllOf) SetAlias(v string)`

SetAlias sets Alias field to given value.

### HasAlias

`func (o *SubscriptionAllOf) HasAlias() bool`

HasAlias returns a boolean if a field has been set.

### GetPlanId

`func (o *SubscriptionAllOf) GetPlanId() string`

GetPlanId returns the PlanId field if non-nil, zero value otherwise.

### GetPlanIdOk

`func (o *SubscriptionAllOf) GetPlanIdOk() (*string, bool)`

GetPlanIdOk returns a tuple with the PlanId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPlanId

`func (o *SubscriptionAllOf) SetPlanId(v string)`

SetPlanId sets PlanId field to given value.

### HasPlanId

`func (o *SubscriptionAllOf) HasPlanId() bool`

HasPlanId returns a boolean if a field has been set.

### GetProvisioned

`func (o *SubscriptionAllOf) GetProvisioned() bool`

GetProvisioned returns the Provisioned field if non-nil, zero value otherwise.

### GetProvisionedOk

`func (o *SubscriptionAllOf) GetProvisionedOk() (*bool, bool)`

GetProvisionedOk returns a tuple with the Provisioned field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvisioned

`func (o *SubscriptionAllOf) SetProvisioned(v bool)`

SetProvisioned sets Provisioned field to given value.

### HasProvisioned

`func (o *SubscriptionAllOf) HasProvisioned() bool`

HasProvisioned returns a boolean if a field has been set.

### GetStatus

`func (o *SubscriptionAllOf) GetStatus() SubscriptionStatus`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *SubscriptionAllOf) GetStatusOk() (*SubscriptionStatus, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *SubscriptionAllOf) SetStatus(v SubscriptionStatus)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *SubscriptionAllOf) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetDeploymentInfo

`func (o *SubscriptionAllOf) GetDeploymentInfo() SubscriptionAllOfDeploymentInfo`

GetDeploymentInfo returns the DeploymentInfo field if non-nil, zero value otherwise.

### GetDeploymentInfoOk

`func (o *SubscriptionAllOf) GetDeploymentInfoOk() (*SubscriptionAllOfDeploymentInfo, bool)`

GetDeploymentInfoOk returns a tuple with the DeploymentInfo field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDeploymentInfo

`func (o *SubscriptionAllOf) SetDeploymentInfo(v SubscriptionAllOfDeploymentInfo)`

SetDeploymentInfo sets DeploymentInfo field to given value.

### HasDeploymentInfo

`func (o *SubscriptionAllOf) HasDeploymentInfo() bool`

HasDeploymentInfo returns a boolean if a field has been set.

### GetConfiguration

`func (o *SubscriptionAllOf) GetConfiguration() SubscriptionAllOfConfiguration`

GetConfiguration returns the Configuration field if non-nil, zero value otherwise.

### GetConfigurationOk

`func (o *SubscriptionAllOf) GetConfigurationOk() (*SubscriptionAllOfConfiguration, bool)`

GetConfigurationOk returns a tuple with the Configuration field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetConfiguration

`func (o *SubscriptionAllOf) SetConfiguration(v SubscriptionAllOfConfiguration)`

SetConfiguration sets Configuration field to given value.

### HasConfiguration

`func (o *SubscriptionAllOf) HasConfiguration() bool`

HasConfiguration returns a boolean if a field has been set.

### GetErrorCause

`func (o *SubscriptionAllOf) GetErrorCause() SubscriptionErrorCause`

GetErrorCause returns the ErrorCause field if non-nil, zero value otherwise.

### GetErrorCauseOk

`func (o *SubscriptionAllOf) GetErrorCauseOk() (*SubscriptionErrorCause, bool)`

GetErrorCauseOk returns a tuple with the ErrorCause field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorCause

`func (o *SubscriptionAllOf) SetErrorCause(v SubscriptionErrorCause)`

SetErrorCause sets ErrorCause field to given value.

### HasErrorCause

`func (o *SubscriptionAllOf) HasErrorCause() bool`

HasErrorCause returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


