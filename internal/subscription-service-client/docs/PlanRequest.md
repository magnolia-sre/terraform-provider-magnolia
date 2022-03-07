# PlanRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**MetricConstraints** | Pointer to [**[]MetricConstraint**](MetricConstraint.md) |  | [optional] 

## Methods

### NewPlanRequest

`func NewPlanRequest(name string, ) *PlanRequest`

NewPlanRequest instantiates a new PlanRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewPlanRequestWithDefaults

`func NewPlanRequestWithDefaults() *PlanRequest`

NewPlanRequestWithDefaults instantiates a new PlanRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *PlanRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *PlanRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *PlanRequest) SetName(v string)`

SetName sets Name field to given value.


### GetMetricConstraints

`func (o *PlanRequest) GetMetricConstraints() []MetricConstraint`

GetMetricConstraints returns the MetricConstraints field if non-nil, zero value otherwise.

### GetMetricConstraintsOk

`func (o *PlanRequest) GetMetricConstraintsOk() (*[]MetricConstraint, bool)`

GetMetricConstraintsOk returns a tuple with the MetricConstraints field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricConstraints

`func (o *PlanRequest) SetMetricConstraints(v []MetricConstraint)`

SetMetricConstraints sets MetricConstraints field to given value.

### HasMetricConstraints

`func (o *PlanRequest) HasMetricConstraints() bool`

HasMetricConstraints returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


