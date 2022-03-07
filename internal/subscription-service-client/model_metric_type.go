/*
Subscription Service Endpoints

## Create subscription  Create a new subscription.  ``` curl -v -d '{\\     \"company\": \"My Company\", \\     \"firstName\": \"First\", \\     \"lastName\": \"Last\", \\     \"email\": \"first.last@magnolia-cms.com\", \\     \"password\": \"some1%2Tres\", \\     \"function\": \"CTO\", \\     \"country\": \"Spain\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/public/subscriptions\" ``` ## Update subscription  Update a new subscription. This is an idempotent operation.  ``` curl -v -d '{\\     \"id\": \"my-company\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/admin/subscriptions/{subId}\" ``` ## Invite users  Invite a list of users to a subscription.  ## Validate invitation  Validate an invitation to a subscription.  ## Activate user  Activate a user for a subscription. The user needs an invitation to be activated. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
	"fmt"
)

// MetricType the model 'MetricType'
type MetricType string

// List of MetricType
const (
	GAUGE MetricType = "GAUGE"
	COUNT MetricType = "COUNT"
)

// All allowed values of MetricType enum
var AllowedMetricTypeEnumValues = []MetricType{
	"GAUGE",
	"COUNT",
}

func (v *MetricType) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MetricType(value)
	for _, existing := range AllowedMetricTypeEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MetricType", value)
}

// NewMetricTypeFromValue returns a pointer to a valid MetricType
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewMetricTypeFromValue(v string) (*MetricType, error) {
	ev := MetricType(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for MetricType: valid values are %v", v, AllowedMetricTypeEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v MetricType) IsValid() bool {
	for _, existing := range AllowedMetricTypeEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to MetricType value
func (v MetricType) Ptr() *MetricType {
	return &v
}

type NullableMetricType struct {
	value *MetricType
	isSet bool
}

func (v NullableMetricType) Get() *MetricType {
	return v.value
}

func (v *NullableMetricType) Set(val *MetricType) {
	v.value = val
	v.isSet = true
}

func (v NullableMetricType) IsSet() bool {
	return v.isSet
}

func (v *NullableMetricType) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMetricType(val *MetricType) *NullableMetricType {
	return &NullableMetricType{value: val, isSet: true}
}

func (v NullableMetricType) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMetricType) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
