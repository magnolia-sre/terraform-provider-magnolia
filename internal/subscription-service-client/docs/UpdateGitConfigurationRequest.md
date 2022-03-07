# UpdateGitConfigurationRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitCloneUrl** | **string** |  | 
**GitProvider** | [**GitProvider**](GitProvider.md) |  | 

## Methods

### NewUpdateGitConfigurationRequest

`func NewUpdateGitConfigurationRequest(gitCloneUrl string, gitProvider GitProvider, ) *UpdateGitConfigurationRequest`

NewUpdateGitConfigurationRequest instantiates a new UpdateGitConfigurationRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateGitConfigurationRequestWithDefaults

`func NewUpdateGitConfigurationRequestWithDefaults() *UpdateGitConfigurationRequest`

NewUpdateGitConfigurationRequestWithDefaults instantiates a new UpdateGitConfigurationRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitCloneUrl

`func (o *UpdateGitConfigurationRequest) GetGitCloneUrl() string`

GetGitCloneUrl returns the GitCloneUrl field if non-nil, zero value otherwise.

### GetGitCloneUrlOk

`func (o *UpdateGitConfigurationRequest) GetGitCloneUrlOk() (*string, bool)`

GetGitCloneUrlOk returns a tuple with the GitCloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneUrl

`func (o *UpdateGitConfigurationRequest) SetGitCloneUrl(v string)`

SetGitCloneUrl sets GitCloneUrl field to given value.


### GetGitProvider

`func (o *UpdateGitConfigurationRequest) GetGitProvider() GitProvider`

GetGitProvider returns the GitProvider field if non-nil, zero value otherwise.

### GetGitProviderOk

`func (o *UpdateGitConfigurationRequest) GetGitProviderOk() (*GitProvider, bool)`

GetGitProviderOk returns a tuple with the GitProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitProvider

`func (o *UpdateGitConfigurationRequest) SetGitProvider(v GitProvider)`

SetGitProvider sets GitProvider field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


