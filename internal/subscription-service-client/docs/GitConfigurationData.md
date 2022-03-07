# GitConfigurationData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**GitSecret** | Pointer to **string** |  | [optional] 
**GitCloneUrl** | Pointer to **string** |  | [optional] 
**Provider** | Pointer to [**GitProvider**](GitProvider.md) |  | [optional] 

## Methods

### NewGitConfigurationData

`func NewGitConfigurationData() *GitConfigurationData`

NewGitConfigurationData instantiates a new GitConfigurationData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitConfigurationDataWithDefaults

`func NewGitConfigurationDataWithDefaults() *GitConfigurationData`

NewGitConfigurationDataWithDefaults instantiates a new GitConfigurationData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetGitSecret

`func (o *GitConfigurationData) GetGitSecret() string`

GetGitSecret returns the GitSecret field if non-nil, zero value otherwise.

### GetGitSecretOk

`func (o *GitConfigurationData) GetGitSecretOk() (*string, bool)`

GetGitSecretOk returns a tuple with the GitSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSecret

`func (o *GitConfigurationData) SetGitSecret(v string)`

SetGitSecret sets GitSecret field to given value.

### HasGitSecret

`func (o *GitConfigurationData) HasGitSecret() bool`

HasGitSecret returns a boolean if a field has been set.

### GetGitCloneUrl

`func (o *GitConfigurationData) GetGitCloneUrl() string`

GetGitCloneUrl returns the GitCloneUrl field if non-nil, zero value otherwise.

### GetGitCloneUrlOk

`func (o *GitConfigurationData) GetGitCloneUrlOk() (*string, bool)`

GetGitCloneUrlOk returns a tuple with the GitCloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneUrl

`func (o *GitConfigurationData) SetGitCloneUrl(v string)`

SetGitCloneUrl sets GitCloneUrl field to given value.

### HasGitCloneUrl

`func (o *GitConfigurationData) HasGitCloneUrl() bool`

HasGitCloneUrl returns a boolean if a field has been set.

### GetProvider

`func (o *GitConfigurationData) GetProvider() GitProvider`

GetProvider returns the Provider field if non-nil, zero value otherwise.

### GetProviderOk

`func (o *GitConfigurationData) GetProviderOk() (*GitProvider, bool)`

GetProviderOk returns a tuple with the Provider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetProvider

`func (o *GitConfigurationData) SetProvider(v GitProvider)`

SetProvider sets Provider field to given value.

### HasProvider

`func (o *GitConfigurationData) HasProvider() bool`

HasProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


