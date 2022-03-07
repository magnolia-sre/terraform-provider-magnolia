# SubscriptionUpdate

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Version** | **int32** |  | 

## Methods

### NewSubscriptionUpdate

`func NewSubscriptionUpdate(name string, version int32, ) *SubscriptionUpdate`

NewSubscriptionUpdate instantiates a new SubscriptionUpdate object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSubscriptionUpdateWithDefaults

`func NewSubscriptionUpdateWithDefaults() *SubscriptionUpdate`

NewSubscriptionUpdateWithDefaults instantiates a new SubscriptionUpdate object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *SubscriptionUpdate) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *SubscriptionUpdate) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *SubscriptionUpdate) SetName(v string)`

SetName sets Name field to given value.


### GetDescription

`func (o *SubscriptionUpdate) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *SubscriptionUpdate) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *SubscriptionUpdate) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *SubscriptionUpdate) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetVersion

`func (o *SubscriptionUpdate) GetVersion() int32`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *SubscriptionUpdate) GetVersionOk() (*int32, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *SubscriptionUpdate) SetVersion(v int32)`

SetVersion sets Version field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


