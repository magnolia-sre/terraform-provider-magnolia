# MetricDetails

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Source** | [**MetricSource**](MetricSource.md) |  | 
**Endpoint** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 

## Methods

### NewMetricDetails

`func NewMetricDetails(source MetricSource, ) *MetricDetails`

NewMetricDetails instantiates a new MetricDetails object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDetailsWithDefaults

`func NewMetricDetailsWithDefaults() *MetricDetails`

NewMetricDetailsWithDefaults instantiates a new MetricDetails object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetSource

`func (o *MetricDetails) GetSource() MetricSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MetricDetails) GetSourceOk() (*MetricSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MetricDetails) SetSource(v MetricSource)`

SetSource sets Source field to given value.


### GetEndpoint

`func (o *MetricDetails) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetricDetails) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetricDetails) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetricDetails) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPath

`func (o *MetricDetails) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MetricDetails) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MetricDetails) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MetricDetails) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetKey

`func (o *MetricDetails) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricDetails) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricDetails) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricDetails) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


