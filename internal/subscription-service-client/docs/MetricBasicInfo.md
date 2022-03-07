# MetricBasicInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Label** | **string** |  | 
**Description** | Pointer to **string** |  | [optional] 
**Unit** | Pointer to **string** |  | [optional] 
**Status** | Pointer to **bool** |  | [optional] 
**Type** | [**MetricType**](MetricType.md) |  | 

## Methods

### NewMetricBasicInfo

`func NewMetricBasicInfo(name string, label string, type_ MetricType, ) *MetricBasicInfo`

NewMetricBasicInfo instantiates a new MetricBasicInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewMetricBasicInfoWithDefaults

`func NewMetricBasicInfoWithDefaults() *MetricBasicInfo`

NewMetricBasicInfoWithDefaults instantiates a new MetricBasicInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *MetricBasicInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *MetricBasicInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *MetricBasicInfo) SetName(v string)`

SetName sets Name field to given value.


### GetLabel

`func (o *MetricBasicInfo) GetLabel() string`

GetLabel returns the Label field if non-nil, zero value otherwise.

### GetLabelOk

`func (o *MetricBasicInfo) GetLabelOk() (*string, bool)`

GetLabelOk returns a tuple with the Label field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLabel

`func (o *MetricBasicInfo) SetLabel(v string)`

SetLabel sets Label field to given value.


### GetDescription

`func (o *MetricBasicInfo) GetDescription() string`

GetDescription returns the Description field if non-nil, zero value otherwise.

### GetDescriptionOk

`func (o *MetricBasicInfo) GetDescriptionOk() (*string, bool)`

GetDescriptionOk returns a tuple with the Description field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescription

`func (o *MetricBasicInfo) SetDescription(v string)`

SetDescription sets Description field to given value.

### HasDescription

`func (o *MetricBasicInfo) HasDescription() bool`

HasDescription returns a boolean if a field has been set.

### GetUnit

`func (o *MetricBasicInfo) GetUnit() string`

GetUnit returns the Unit field if non-nil, zero value otherwise.

### GetUnitOk

`func (o *MetricBasicInfo) GetUnitOk() (*string, bool)`

GetUnitOk returns a tuple with the Unit field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUnit

`func (o *MetricBasicInfo) SetUnit(v string)`

SetUnit sets Unit field to given value.

### HasUnit

`func (o *MetricBasicInfo) HasUnit() bool`

HasUnit returns a boolean if a field has been set.

### GetStatus

`func (o *MetricBasicInfo) GetStatus() bool`

GetStatus returns the Status field if non-nil, zero value otherwise.

### GetStatusOk

`func (o *MetricBasicInfo) GetStatusOk() (*bool, bool)`

GetStatusOk returns a tuple with the Status field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetStatus

`func (o *MetricBasicInfo) SetStatus(v bool)`

SetStatus sets Status field to given value.

### HasStatus

`func (o *MetricBasicInfo) HasStatus() bool`

HasStatus returns a boolean if a field has been set.

### GetType

`func (o *MetricBasicInfo) GetType() MetricType`

GetType returns the Type field if non-nil, zero value otherwise.

### GetTypeOk

`func (o *MetricBasicInfo) GetTypeOk() (*MetricType, bool)`

GetTypeOk returns a tuple with the Type field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetType

`func (o *MetricBasicInfo) SetType(v MetricType)`

SetType sets Type field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


