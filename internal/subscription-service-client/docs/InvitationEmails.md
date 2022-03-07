# InvitationEmails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Emails** | **[]string** | A set of emails, used e.g. for invitations or reactivation. | 
**Group** | **string** |  | 

## Methods

### NewInvitationEmails

`func NewInvitationEmails(emails []string, group string, ) *InvitationEmails`

NewInvitationEmails instantiates a new InvitationEmails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewInvitationEmailsWithDefaults

`func NewInvitationEmailsWithDefaults() *InvitationEmails`

NewInvitationEmailsWithDefaults instantiates a new InvitationEmails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetEmails

`func (o *InvitationEmails) GetEmails() []string`

GetEmails returns the Emails field if non-nil, zero value otherwise.

### GetEmailsOk

`func (o *InvitationEmails) GetEmailsOk() (*[]string, bool)`

GetEmailsOk returns a tuple with the Emails field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEmails

`func (o *InvitationEmails) SetEmails(v []string)`

SetEmails sets Emails field to given value.


### GetGroup

`func (o *InvitationEmails) GetGroup() string`

GetGroup returns the Group field if non-nil, zero value otherwise.

### GetGroupOk

`func (o *InvitationEmails) GetGroupOk() (*string, bool)`

GetGroupOk returns a tuple with the Group field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetGroup

`func (o *InvitationEmails) SetGroup(v string)`

SetGroup sets Group field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


