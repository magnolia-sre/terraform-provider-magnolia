/*
Subscription Service Endpoints

## Create subscription  Create a new subscription.  ``` curl -v -d '{\\     \"company\": \"My Company\", \\     \"firstName\": \"First\", \\     \"lastName\": \"Last\", \\     \"email\": \"first.last@magnolia-cms.com\", \\     \"password\": \"some1%2Tres\", \\     \"function\": \"CTO\", \\     \"country\": \"Spain\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/public/subscriptions\" ``` ## Update subscription  Update a new subscription. This is an idempotent operation.  ``` curl -v -d '{\\     \"id\": \"my-company\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/admin/subscriptions/{subId}\" ``` ## Invite users  Invite a list of users to a subscription.  ## Validate invitation  Validate an invitation to a subscription.  ## Activate user  Activate a user for a subscription. The user needs an invitation to be activated. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"encoding/json"
)

// PlanRequest struct for PlanRequest
type PlanRequest struct {
	Name string `json:"name"`
	MetricConstraints []MetricConstraint `json:"metricConstraints,omitempty"`
}

// NewPlanRequest instantiates a new PlanRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPlanRequest(name string) *PlanRequest {
	this := PlanRequest{}
	this.Name = name
	return &this
}

// NewPlanRequestWithDefaults instantiates a new PlanRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPlanRequestWithDefaults() *PlanRequest {
	this := PlanRequest{}
	return &this
}

// GetName returns the Name field value
func (o *PlanRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *PlanRequest) SetName(v string) {
	o.Name = v
}

// GetMetricConstraints returns the MetricConstraints field value if set, zero value otherwise.
func (o *PlanRequest) GetMetricConstraints() []MetricConstraint {
	if o == nil || o.MetricConstraints == nil {
		var ret []MetricConstraint
		return ret
	}
	return o.MetricConstraints
}

// GetMetricConstraintsOk returns a tuple with the MetricConstraints field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PlanRequest) GetMetricConstraintsOk() ([]MetricConstraint, bool) {
	if o == nil || o.MetricConstraints == nil {
		return nil, false
	}
	return o.MetricConstraints, true
}

// HasMetricConstraints returns a boolean if a field has been set.
func (o *PlanRequest) HasMetricConstraints() bool {
	if o != nil && o.MetricConstraints != nil {
		return true
	}

	return false
}

// SetMetricConstraints gets a reference to the given []MetricConstraint and assigns it to the MetricConstraints field.
func (o *PlanRequest) SetMetricConstraints(v []MetricConstraint) {
	o.MetricConstraints = v
}

func (o PlanRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.MetricConstraints != nil {
		toSerialize["metricConstraints"] = o.MetricConstraints
	}
	return json.Marshal(toSerialize)
}

type NullablePlanRequest struct {
	value *PlanRequest
	isSet bool
}

func (v NullablePlanRequest) Get() *PlanRequest {
	return v.value
}

func (v *NullablePlanRequest) Set(val *PlanRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePlanRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePlanRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePlanRequest(val *PlanRequest) *NullablePlanRequest {
	return &NullablePlanRequest{value: val, isSet: true}
}

func (v NullablePlanRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePlanRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


