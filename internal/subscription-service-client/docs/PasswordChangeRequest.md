# PasswordChangeRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Strict** | Pointer to **string** |  | [optional] 
**OldPassword** | Pointer to **string** |  | [optional] 
**NewPassword** | Pointer to **string** |  | [optional] 

## Methods

### NewPasswordChangeRequest

`func NewPasswordChangeRequest() *PasswordChangeRequest`

NewPasswordChangeRequest instantiates a new PasswordChangeRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPasswordChangeRequestWithDefaults

`func NewPasswordChangeRequestWithDefaults() *PasswordChangeRequest`

NewPasswordChangeRequestWithDefaults instantiates a new PasswordChangeRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetStrict

`func (o *PasswordChangeRequest) GetStrict() string`

GetStrict returns the Strict field if non-nil, zero value otherwise.

### GetStrictOk

`func (o *PasswordChangeRequest) GetStrictOk() (*string, bool)`

GetStrictOk returns a tuple with the Strict field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStrict

`func (o *PasswordChangeRequest) SetStrict(v string)`

SetStrict sets Strict field to given value.

### HasStrict

`func (o *PasswordChangeRequest) HasStrict() bool`

HasStrict returns a boolean if a field has been set.

### GetOldPassword

`func (o *PasswordChangeRequest) GetOldPassword() string`

GetOldPassword returns the OldPassword field if non-nil, zero value otherwise.

### GetOldPasswordOk

`func (o *PasswordChangeRequest) GetOldPasswordOk() (*string, bool)`

GetOldPasswordOk returns a tuple with the OldPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetOldPassword

`func (o *PasswordChangeRequest) SetOldPassword(v string)`

SetOldPassword sets OldPassword field to given value.

### HasOldPassword

`func (o *PasswordChangeRequest) HasOldPassword() bool`

HasOldPassword returns a boolean if a field has been set.

### GetNewPassword

`func (o *PasswordChangeRequest) GetNewPassword() string`

GetNewPassword returns the NewPassword field if non-nil, zero value otherwise.

### GetNewPasswordOk

`func (o *PasswordChangeRequest) GetNewPasswordOk() (*string, bool)`

GetNewPasswordOk returns a tuple with the NewPassword field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetNewPassword

`func (o *PasswordChangeRequest) SetNewPassword(v string)`

SetNewPassword sets NewPassword field to given value.

### HasNewPassword

`func (o *PasswordChangeRequest) HasNewPassword() bool`

HasNewPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


