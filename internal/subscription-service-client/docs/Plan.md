# Plan

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Id** | **string** |  | 
**Name** | **string** |  | 
**Creation** | Pointer to **int64** |  | [optional] 
**Update** | Pointer to **int64** |  | [optional] 
**Active** | **bool** |  | 
**DefaultForNewSubscriptions** | **bool** |  | 
**MetricConstraints** | Pointer to [**[]MetricConstraint**](MetricConstraint.md) |  | [optional] 

## Methods

### NewPlan

`func NewPlan(id string, name string, active bool, defaultForNewSubscriptions bool, ) *Plan`

NewPlan instantiates a new Plan object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanWithDefaults

`func NewPlanWithDefaults() *Plan`

NewPlanWithDefaults instantiates a new Plan object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetId

`func (o *Plan) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Plan) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Plan) SetId(v string)`

SetId sets Id field to given value.


### GetName

`func (o *Plan) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Plan) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Plan) SetName(v string)`

SetName sets Name field to given value.


### GetCreation

`func (o *Plan) GetCreation() int64`

GetCreation returns the Creation field if non-nil, zero value otherwise.

### GetCreationOk

`func (o *Plan) GetCreationOk() (*int64, bool)`

GetCreationOk returns a tuple with the Creation field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCreation

`func (o *Plan) SetCreation(v int64)`

SetCreation sets Creation field to given value.

### HasCreation

`func (o *Plan) HasCreation() bool`

HasCreation returns a boolean if a field has been set.

### GetUpdate

`func (o *Plan) GetUpdate() int64`

GetUpdate returns the Update field if non-nil, zero value otherwise.

### GetUpdateOk

`func (o *Plan) GetUpdateOk() (*int64, bool)`

GetUpdateOk returns a tuple with the Update field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdate

`func (o *Plan) SetUpdate(v int64)`

SetUpdate sets Update field to given value.

### HasUpdate

`func (o *Plan) HasUpdate() bool`

HasUpdate returns a boolean if a field has been set.

### GetActive

`func (o *Plan) GetActive() bool`

GetActive returns the Active field if non-nil, zero value otherwise.

### GetActiveOk

`func (o *Plan) GetActiveOk() (*bool, bool)`

GetActiveOk returns a tuple with the Active field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetActive

`func (o *Plan) SetActive(v bool)`

SetActive sets Active field to given value.


### GetDefaultForNewSubscriptions

`func (o *Plan) GetDefaultForNewSubscriptions() bool`

GetDefaultForNewSubscriptions returns the DefaultForNewSubscriptions field if non-nil, zero value otherwise.

### GetDefaultForNewSubscriptionsOk

`func (o *Plan) GetDefaultForNewSubscriptionsOk() (*bool, bool)`

GetDefaultForNewSubscriptionsOk returns a tuple with the DefaultForNewSubscriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultForNewSubscriptions

`func (o *Plan) SetDefaultForNewSubscriptions(v bool)`

SetDefaultForNewSubscriptions sets DefaultForNewSubscriptions field to given value.


### GetMetricConstraints

`func (o *Plan) GetMetricConstraints() []MetricConstraint`

GetMetricConstraints returns the MetricConstraints field if non-nil, zero value otherwise.

### GetMetricConstraintsOk

`func (o *Plan) GetMetricConstraintsOk() (*[]MetricConstraint, bool)`

GetMetricConstraintsOk returns a tuple with the MetricConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricConstraints

`func (o *Plan) SetMetricConstraints(v []MetricConstraint)`

SetMetricConstraints sets MetricConstraints field to given value.

### HasMetricConstraints

`func (o *Plan) HasMetricConstraints() bool`

HasMetricConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


