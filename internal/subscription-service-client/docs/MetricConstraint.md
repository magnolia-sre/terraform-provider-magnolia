# MetricConstraint

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**MetricId** | Pointer to **string** |  | [optional] 
**Type** | Pointer to **string** |  | [optional] 
**Limit** | Pointer to **int32** |  | [optional] 

## Methods

### NewMetricConstraint

`func NewMetricConstraint() *MetricConstraint`

NewMetricConstraint instantiates a new MetricConstraint object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricConstraintWithDefaults

`func NewMetricConstraintWithDefaults() *MetricConstraint`

NewMetricConstraintWithDefaults instantiates a new MetricConstraint object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetricId

`func (o *MetricConstraint) GetMetricId() string`

GetMetricId returns the MetricId field if non-nil, zero value otherwise.

### GetMetricIdOk

`func (o *MetricConstraint) GetMetricIdOk() (*string, bool)`

GetMetricIdOk returns a tuple with the MetricId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetricId

`func (o *MetricConstraint) SetMetricId(v string)`

SetMetricId sets MetricId field to given value.

### HasMetricId

`func (o *MetricConstraint) HasMetricId() bool`

HasMetricId returns a boolean if a field has been set.

### GetType

`func (o *MetricConstraint) GetType() string`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricConstraint) GetTypeOk() (*string, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricConstraint) SetType(v string)`

SetType sets Type field to given value.

### HasType

`func (o *MetricConstraint) HasType() bool`

HasType returns a boolean if a field has been set.

### GetLimit

`func (o *MetricConstraint) GetLimit() int32`

GetLimit returns the Limit field if non-nil, zero value otherwise.

### GetLimitOk

`func (o *MetricConstraint) GetLimitOk() (*int32, bool)`

GetLimitOk returns a tuple with the Limit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLimit

`func (o *MetricConstraint) SetLimit(v int32)`

SetLimit sets Limit field to given value.

### HasLimit

`func (o *MetricConstraint) HasLimit() bool`

HasLimit returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


