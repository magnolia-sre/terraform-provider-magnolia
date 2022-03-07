# GitConfigurationResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**PublicKey** | Pointer to **string** |  | [optional] 
**GitSecret** | Pointer to **string** |  | [optional] 
**WebhookUrl** | Pointer to **string** |  | [optional] 
**GitCloneUrl** | Pointer to **string** |  | [optional] 
**GitProvider** | Pointer to [**GitProvider**](GitProvider.md) |  | [optional] 

## Methods

### NewGitConfigurationResponse

`func NewGitConfigurationResponse() *GitConfigurationResponse`

NewGitConfigurationResponse instantiates a new GitConfigurationResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewGitConfigurationResponseWithDefaults

`func NewGitConfigurationResponseWithDefaults() *GitConfigurationResponse`

NewGitConfigurationResponseWithDefaults instantiates a new GitConfigurationResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetPublicKey

`func (o *GitConfigurationResponse) GetPublicKey() string`

GetPublicKey returns the PublicKey field if non-nil, zero value otherwise.

### GetPublicKeyOk

`func (o *GitConfigurationResponse) GetPublicKeyOk() (*string, bool)`

GetPublicKeyOk returns a tuple with the PublicKey field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPublicKey

`func (o *GitConfigurationResponse) SetPublicKey(v string)`

SetPublicKey sets PublicKey field to given value.

### HasPublicKey

`func (o *GitConfigurationResponse) HasPublicKey() bool`

HasPublicKey returns a boolean if a field has been set.

### GetGitSecret

`func (o *GitConfigurationResponse) GetGitSecret() string`

GetGitSecret returns the GitSecret field if non-nil, zero value otherwise.

### GetGitSecretOk

`func (o *GitConfigurationResponse) GetGitSecretOk() (*string, bool)`

GetGitSecretOk returns a tuple with the GitSecret field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitSecret

`func (o *GitConfigurationResponse) SetGitSecret(v string)`

SetGitSecret sets GitSecret field to given value.

### HasGitSecret

`func (o *GitConfigurationResponse) HasGitSecret() bool`

HasGitSecret returns a boolean if a field has been set.

### GetWebhookUrl

`func (o *GitConfigurationResponse) GetWebhookUrl() string`

GetWebhookUrl returns the WebhookUrl field if non-nil, zero value otherwise.

### GetWebhookUrlOk

`func (o *GitConfigurationResponse) GetWebhookUrlOk() (*string, bool)`

GetWebhookUrlOk returns a tuple with the WebhookUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetWebhookUrl

`func (o *GitConfigurationResponse) SetWebhookUrl(v string)`

SetWebhookUrl sets WebhookUrl field to given value.

### HasWebhookUrl

`func (o *GitConfigurationResponse) HasWebhookUrl() bool`

HasWebhookUrl returns a boolean if a field has been set.

### GetGitCloneUrl

`func (o *GitConfigurationResponse) GetGitCloneUrl() string`

GetGitCloneUrl returns the GitCloneUrl field if non-nil, zero value otherwise.

### GetGitCloneUrlOk

`func (o *GitConfigurationResponse) GetGitCloneUrlOk() (*string, bool)`

GetGitCloneUrlOk returns a tuple with the GitCloneUrl field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitCloneUrl

`func (o *GitConfigurationResponse) SetGitCloneUrl(v string)`

SetGitCloneUrl sets GitCloneUrl field to given value.

### HasGitCloneUrl

`func (o *GitConfigurationResponse) HasGitCloneUrl() bool`

HasGitCloneUrl returns a boolean if a field has been set.

### GetGitProvider

`func (o *GitConfigurationResponse) GetGitProvider() GitProvider`

GetGitProvider returns the GitProvider field if non-nil, zero value otherwise.

### GetGitProviderOk

`func (o *GitConfigurationResponse) GetGitProviderOk() (*GitProvider, bool)`

GetGitProviderOk returns a tuple with the GitProvider field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGitProvider

`func (o *GitConfigurationResponse) SetGitProvider(v GitProvider)`

SetGitProvider sets GitProvider field to given value.

### HasGitProvider

`func (o *GitConfigurationResponse) HasGitProvider() bool`

HasGitProvider returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


