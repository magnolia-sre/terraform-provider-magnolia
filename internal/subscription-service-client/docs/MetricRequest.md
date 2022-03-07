# MetricRequest

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

## Methods

### NewMetricRequest

`func NewMetricRequest(name string, label string, type_ MetricType, source MetricSource, ) *MetricRequest`

NewMetricRequest instantiates a new MetricRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricRequestWithDefaults

`func NewMetricRequestWithDefaults() *MetricRequest`

NewMetricRequestWithDefaults instantiates a new MetricRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricRequest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricRequest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricRequest) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *MetricRequest) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricRequest) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricRequest) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *MetricRequest) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricRequest) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricRequest) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricRequest) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *MetricRequest) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricRequest) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricRequest) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricRequest) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStatus

`func (o *MetricRequest) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricRequest) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricRequest) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetricRequest) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MetricRequest) GetType() MetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricRequest) GetTypeOk() (*MetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricRequest) SetType(v MetricType)`

SetType sets Type field to given value.


### GetSource

`func (o *MetricRequest) GetSource() MetricSource`

GetSource returns the Source field if non-nil, zero value otherwise.

### GetSourceOk

`func (o *MetricRequest) GetSourceOk() (*MetricSource, bool)`

GetSourceOk returns a tuple with the Source field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSource

`func (o *MetricRequest) SetSource(v MetricSource)`

SetSource sets Source field to given value.


### GetEndpoint

`func (o *MetricRequest) GetEndpoint() string`

GetEndpoint returns the Endpoint field if non-nil, zero value otherwise.

### GetEndpointOk

`func (o *MetricRequest) GetEndpointOk() (*string, bool)`

GetEndpointOk returns a tuple with the Endpoint field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetEndpoint

`func (o *MetricRequest) SetEndpoint(v string)`

SetEndpoint sets Endpoint field to given value.

### HasEndpoint

`func (o *MetricRequest) HasEndpoint() bool`

HasEndpoint returns a boolean if a field has been set.

### GetPath

`func (o *MetricRequest) GetPath() string`

GetPath returns the Path field if non-nil, zero value otherwise.

### GetPathOk

`func (o *MetricRequest) GetPathOk() (*string, bool)`

GetPathOk returns a tuple with the Path field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPath

`func (o *MetricRequest) SetPath(v string)`

SetPath sets Path field to given value.

### HasPath

`func (o *MetricRequest) HasPath() bool`

HasPath returns a boolean if a field has been set.

### GetKey

`func (o *MetricRequest) GetKey() string`

GetKey returns the Key field if non-nil, zero value otherwise.

### GetKeyOk

`func (o *MetricRequest) GetKeyOk() (*string, bool)`

GetKeyOk returns a tuple with the Key field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetKey

`func (o *MetricRequest) SetKey(v string)`

SetKey sets Key field to given value.

### HasKey

`func (o *MetricRequest) HasKey() bool`

HasKey returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


