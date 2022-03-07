# Metric

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Type** | [**MetricType**](MetricType.md) |  | 
**Source** | [**MetricSource**](MetricSource.md) |  | 
**Endpoint** | Pointer to **string** |  | [optional] 
**Path** | Pointer to **string** |  | [optional] 
**Key** | Pointer to **string** |  | [optional] 
**Id** | **string** |  | 

## Methods

### NewMetric

`func NewMetric(name string, label string, type_ MetricType, source MetricSource, id string, ) *Metric`

NewMetric instantiates a new Metric object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricWithDefaults

`func NewMetricWithDefaults() *Metric`

NewMetricWithDefaults instantiates a new Metric object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *Metric) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *Metric) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *Metric) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *Metric) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *Metric) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *Metric) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *Metric) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *Metric) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *Metric) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *Metric) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *Metric) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *Metric) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *Metric) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *Metric) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStatus

`func (o *Metric) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *Metric) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *Metric) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *Metric) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *Metric) GetType() MetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *Metric) GetTypeOk() (*MetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *Metric) SetType(v MetricType)`

SetType sets Type field to given value.


### GetSource

`func (o *Metric) GetSource() MetricSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *Metric) GetSourceOk() (*MetricSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *Metric) SetSource(v MetricSource)`

SetSource sets Source field to given value.


### GetEndpoint

`func (o *Metric) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *Metric) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *Metric) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *Metric) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPath

`func (o *Metric) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *Metric) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *Metric) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *Metric) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetKey

`func (o *Metric) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *Metric) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *Metric) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *Metric) HasKey() bool`

HasKey returns a boolean if a field has been set.

### GetId

`func (o *Metric) GetId() string`

GetId returns the Id field if non-nil, zero value otherwise.

### GetIdOk

`func (o *Metric) GetIdOk() (*string, bool)`

GetIdOk returns a tuple with the Id field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetId

`func (o *Metric) SetId(v string)`

SetId sets Id field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


