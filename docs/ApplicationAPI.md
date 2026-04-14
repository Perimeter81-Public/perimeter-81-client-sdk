# \ApplicationAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateApplication**](ApplicationAPI.md#CreateApplication) | **Post** /v2.3/applications | Create Application
[**GetApplicationById**](ApplicationAPI.md#GetApplicationById) | **Get** /v2.3/applications/{applicationId} | Get Application by Id
[**GetApplicationStatus**](ApplicationAPI.md#GetApplicationStatus) | **Get** /v2.3/applications/status/{statusId} | Get status of Application creation process by statusId
[**GetApplications**](ApplicationAPI.md#GetApplications) | **Get** /v2.3/applications | Get Applications List



## CreateApplication

> AsyncOperationResponse CreateApplication(ctx).CreateApplicationRequest(createApplicationRequest).Execute()

Create Application



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createApplicationRequest := openapiclient.createApplication_request{HttpCreateApplication: openapiclient.NewHttpCreateApplication("MyApplication", "Type_example", "Network_example", "TODO", "TODO", []string{"Users_example"}, []string{"Groups_example"}, map[string]interface{}({"X-Custom-Header":"value"}), *openapiclient.NewHttpAttributes())} // CreateApplicationRequest | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.CreateApplication(context.Background()).CreateApplicationRequest(createApplicationRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.CreateApplication``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateApplication`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.CreateApplication`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateApplicationRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createApplicationRequest** | [**CreateApplicationRequest**](CreateApplicationRequest.md) |  | 

### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationById

> GetApplicationById200Response GetApplicationById(ctx, applicationId).Execute()

Get Application by Id



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	applicationId := "applicationId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationById(context.Background(), applicationId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationById``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationById`: GetApplicationById200Response
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationById`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**applicationId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationByIdRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**GetApplicationById200Response**](GetApplicationById200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplicationStatus

> ApplicationStatusResponse GetApplicationStatus(ctx, statusId).Execute()

Get status of Application creation process by statusId



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	statusId := "statusId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplicationStatus(context.Background(), statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplicationStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplicationStatus`: ApplicationStatusResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplicationStatus`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationStatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ApplicationStatusResponse**](ApplicationStatusResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetApplications

> ApplicationsListPaginatedResponse GetApplications(ctx).Name(name).Host(host).Type_(type_).Page(page).Limit(limit).Sort(sort).Execute()

Get Applications List



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	name := "name_example" // string | Application name to filter (optional)
	host := "host_example" // string | Application host to filter (optional)
	type_ := "http" // string | Application type to filter (optional)
	page := int32(56) // int32 | Specifies the current page of the paginated result set (optional)
	limit := int32(56) // int32 | Defines the number of results per page (optional)
	sort := "sort[name]=asc" // string | Sorts the results based on a specified field and order. Available fields for sorting are name, host and type. (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ApplicationAPI.GetApplications(context.Background()).Name(name).Host(host).Type_(type_).Page(page).Limit(limit).Sort(sort).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ApplicationAPI.GetApplications``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetApplications`: ApplicationsListPaginatedResponse
	fmt.Fprintf(os.Stdout, "Response from `ApplicationAPI.GetApplications`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiGetApplicationsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **name** | **string** | Application name to filter | 
 **host** | **string** | Application host to filter | 
 **type_** | **string** | Application type to filter | 
 **page** | **int32** | Specifies the current page of the paginated result set | 
 **limit** | **int32** | Defines the number of results per page | 
 **sort** | **string** | Sorts the results based on a specified field and order. Available fields for sorting are name, host and type. | 

### Return type

[**ApplicationsListPaginatedResponse**](ApplicationsListPaginatedResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

