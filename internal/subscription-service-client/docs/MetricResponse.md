# MetricResponse

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Metrics** | Pointer to [**[]MetricData**](MetricData.md) |  | [optional] 
**Updated** | Pointer to **int64** |  | [optional] 

## Methods

### NewMetricResponse

`func NewMetricResponse() *MetricResponse`

NewMetricResponse instantiates a new MetricResponse object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricResponseWithDefaults

`func NewMetricResponseWithDefaults() *MetricResponse`

NewMetricResponseWithDefaults instantiates a new MetricResponse object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetMetrics

`func (o *MetricResponse) GetMetrics() []MetricData`

GetMetrics returns the Metrics field if non-nil, zero value otherwise.

### GetMetricsOk

`func (o *MetricResponse) GetMetricsOk() (*[]MetricData, bool)`

GetMetricsOk returns a tuple with the Metrics field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetMetrics

`func (o *MetricResponse) SetMetrics(v []MetricData)`

SetMetrics sets Metrics field to given value.

### HasMetrics

`func (o *MetricResponse) HasMetrics() bool`

HasMetrics returns a boolean if a field has been set.

### GetUpdated

`func (o *MetricResponse) GetUpdated() int64`

GetUpdated returns the Updated field if non-nil, zero value otherwise.

### GetUpdatedOk

`func (o *MetricResponse) GetUpdatedOk() (*int64, bool)`

GetUpdatedOk returns a tuple with the Updated field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdated

`func (o *MetricResponse) SetUpdated(v int64)`

SetUpdated sets Updated field to given value.

### HasUpdated

`func (o *MetricResponse) HasUpdated() bool`

HasUpdated returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


