# MetricData

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Type** | [**MetricType**](MetricType.md) |  | 
**Value** | Pointer to **float64** |  | [optional] 

## Methods

### NewMetricData

`func NewMetricData(name string, label string, type_ MetricType, ) *MetricData`

NewMetricData instantiates a new MetricData object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricDataWithDefaults

`func NewMetricDataWithDefaults() *MetricData`

NewMetricDataWithDefaults instantiates a new MetricData object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricData) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricData) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricData) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *MetricData) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricData) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricData) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *MetricData) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricData) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricData) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricData) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *MetricData) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricData) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricData) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricData) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStatus

`func (o *MetricData) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricData) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricData) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetricData) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MetricData) GetType() MetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricData) GetTypeOk() (*MetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricData) SetType(v MetricType)`

SetType sets Type field to given value.


### GetValue

`func (o *MetricData) GetValue() float64`

GetValue returns the Value field if non-nil, zero value otherwise.

### GetValueOk

`func (o *MetricData) GetValueOk() (*float64, bool)`

GetValueOk returns a tuple with the Value field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetValue

`func (o *MetricData) SetValue(v float64)`

SetValue sets Value field to given value.

### HasValue

`func (o *MetricData) HasValue() bool`

HasValue returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


