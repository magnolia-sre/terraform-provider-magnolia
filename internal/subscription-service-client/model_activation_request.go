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

// ActivationRequest struct for ActivationRequest
type ActivationRequest struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Function string `json:"function"`
	Password string `json:"password"`
}

// NewActivationRequest instantiates a new ActivationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewActivationRequest(firstName string, lastName string, function string, password string) *ActivationRequest {
	this := ActivationRequest{}
	this.FirstName = firstName
	this.LastName = lastName
	this.Function = function
	this.Password = password
	return &this
}

// NewActivationRequestWithDefaults instantiates a new ActivationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewActivationRequestWithDefaults() *ActivationRequest {
	this := ActivationRequest{}
	return &this
}

// GetFirstName returns the FirstName field value
func (o *ActivationRequest) GetFirstName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.FirstName
}

// GetFirstNameOk returns a tuple with the FirstName field value
// and a boolean to check if the value has been set.
func (o *ActivationRequest) GetFirstNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.FirstName, true
}

// SetFirstName sets field value
func (o *ActivationRequest) SetFirstName(v string) {
	o.FirstName = v
}

// GetLastName returns the LastName field value
func (o *ActivationRequest) GetLastName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.LastName
}

// GetLastNameOk returns a tuple with the LastName field value
// and a boolean to check if the value has been set.
func (o *ActivationRequest) GetLastNameOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.LastName, true
}

// SetLastName sets field value
func (o *ActivationRequest) SetLastName(v string) {
	o.LastName = v
}

// GetFunction returns the Function field value
func (o *ActivationRequest) GetFunction() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Function
}

// GetFunctionOk returns a tuple with the Function field value
// and a boolean to check if the value has been set.
func (o *ActivationRequest) GetFunctionOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Function, true
}

// SetFunction sets field value
func (o *ActivationRequest) SetFunction(v string) {
	o.Function = v
}

// GetPassword returns the Password field value
func (o *ActivationRequest) GetPassword() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Password
}

// GetPasswordOk returns a tuple with the Password field value
// and a boolean to check if the value has been set.
func (o *ActivationRequest) GetPasswordOk() (*string, bool) {
	if o == nil  {
		return nil, false
	}
	return &o.Password, true
}

// SetPassword sets field value
func (o *ActivationRequest) SetPassword(v string) {
	o.Password = v
}

func (o ActivationRequest) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["firstName"] = o.FirstName
	}
	if true {
		toSerialize["lastName"] = o.LastName
	}
	if true {
		toSerialize["function"] = o.Function
	}
	if true {
		toSerialize["password"] = o.Password
	}
	return json.Marshal(toSerialize)
}

type NullableActivationRequest struct {
	value *ActivationRequest
	isSet bool
}

func (v NullableActivationRequest) Get() *ActivationRequest {
	return v.value
}

func (v *NullableActivationRequest) Set(val *ActivationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableActivationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableActivationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableActivationRequest(val *ActivationRequest) *NullableActivationRequest {
	return &NullableActivationRequest{value: val, isSet: true}
}

func (v NullableActivationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableActivationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}


