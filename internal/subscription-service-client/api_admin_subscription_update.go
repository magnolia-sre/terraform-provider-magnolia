/*
Subscription Service Endpoints

## Create subscription  Create a new subscription.  ``` curl -v -d '{\\     \"company\": \"My Company\", \\     \"firstName\": \"First\", \\     \"lastName\": \"Last\", \\     \"email\": \"first.last@magnolia-cms.com\", \\     \"password\": \"some1%2Tres\", \\     \"function\": \"CTO\", \\     \"country\": \"Spain\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/public/subscriptions\" ``` ## Update subscription  Update a new subscription. This is an idempotent operation.  ``` curl -v -d '{\\     \"id\": \"my-company\" \\     }' \\ -H \"Content-Type: application/json\" -X POST \"http://localhost:8080/admin/subscriptions/{subId}\" ``` ## Invite users  Invite a list of users to a subscription.  ## Validate invitation  Validate an invitation to a subscription.  ## Activate user  Activate a user for a subscription. The user needs an invitation to be activated. 

API version: 1.1.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package openapi

import (
	"bytes"
	"context"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

// Linger please
var (
	_ context.Context
)

// AdminSubscriptionUpdateApiService AdminSubscriptionUpdateApi service
type AdminSubscriptionUpdateApiService service

type ApiFindPendingUpdatesRequest struct {
	ctx context.Context
	ApiService *AdminSubscriptionUpdateApiService
	subscriptionId string
}


func (r ApiFindPendingUpdatesRequest) Execute() ([]SubscriptionUpdate, *http.Response, error) {
	return r.ApiService.FindPendingUpdatesExecute(r)
}

/*
FindPendingUpdates Find pending updates of given subscription

Find pending updates of given subscriptionId

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Generated id of subscription
 @return ApiFindPendingUpdatesRequest
*/
func (a *AdminSubscriptionUpdateApiService) FindPendingUpdates(ctx context.Context, subscriptionId string) ApiFindPendingUpdatesRequest {
	return ApiFindPendingUpdatesRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return []SubscriptionUpdate
func (a *AdminSubscriptionUpdateApiService) FindPendingUpdatesExecute(r ApiFindPendingUpdatesRequest) ([]SubscriptionUpdate, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodGet
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  []SubscriptionUpdate
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminSubscriptionUpdateApiService.FindPendingUpdates")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/subscriptions/{subscriptionId}/migrations/pending"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiUpdateSubscriptionRequest struct {
	ctx context.Context
	ApiService *AdminSubscriptionUpdateApiService
	subscriptionId string
}


func (r ApiUpdateSubscriptionRequest) Execute() (*Subscription, *http.Response, error) {
	return r.ApiService.UpdateSubscriptionExecute(r)
}

/*
UpdateSubscription Update and deploy subscription

Update and deploy subscription

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @param subscriptionId Generated id of subscription
 @return ApiUpdateSubscriptionRequest
*/
func (a *AdminSubscriptionUpdateApiService) UpdateSubscription(ctx context.Context, subscriptionId string) ApiUpdateSubscriptionRequest {
	return ApiUpdateSubscriptionRequest{
		ApiService: a,
		ctx: ctx,
		subscriptionId: subscriptionId,
	}
}

// Execute executes the request
//  @return Subscription
func (a *AdminSubscriptionUpdateApiService) UpdateSubscriptionExecute(r ApiUpdateSubscriptionRequest) (*Subscription, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *Subscription
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "AdminSubscriptionUpdateApiService.UpdateSubscription")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/admin/subscriptions/{subscriptionId}/migrations"
	localVarPath = strings.Replace(localVarPath, "{"+"subscriptionId"+"}", url.PathEscape(parameterToString(r.subscriptionId, "")), -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = ioutil.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
