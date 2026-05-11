# \ObjectsServicesAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectsServices**](ObjectsServicesAPI.md#DeleteObjectsServices) | **Delete** /v2.3/objects/services/{objectId} | Delete Service object
[**GetObjectsServices**](ObjectsServicesAPI.md#GetObjectsServices) | **Get** /v2.3/objects/services | Get object services
[**PostObjectsServices**](ObjectsServicesAPI.md#PostObjectsServices) | **Post** /v2.3/objects/services | Create new Service object
[**PutObjectsServices**](ObjectsServicesAPI.md#PutObjectsServices) | **Put** /v2.3/objects/services/{objectId} | Update Service object



## DeleteObjectsServices

> DeleteObjectsServices(ctx, objectId).Execute()

Delete Service object



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
	objectId := "objectId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	r, err := apiClient.ObjectsServicesAPI.DeleteObjectsServices(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsServicesAPI.DeleteObjectsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteObjectsServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

 (empty response body)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: Not defined

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetObjectsServices

> ObjectsServicesResponse GetObjectsServices(ctx).Execute()

Get object services



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

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsServicesAPI.GetObjectsServices(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsServicesAPI.GetObjectsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectsServices`: ObjectsServicesResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectsServicesAPI.GetObjectsServices`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectsServicesRequest struct via the builder pattern


### Return type

[**ObjectsServicesResponse**](ObjectsServicesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectsServices

> PostObjectsServices201Response PostObjectsServices(ctx).ObjectsServicesRequestObj(objectsServicesRequestObj).Execute()

Create new Service object



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
	objectsServicesRequestObj := *openapiclient.NewObjectsServicesRequestObj("Name_example", []openapiclient.ObjectsServicesProtocolRequestObj{*openapiclient.NewObjectsServicesProtocolRequestObj("Protocol_example", "ValueType_example", []int32{int32(123)}, openapiclient.ObjectServiceProtocolOptionsICMPrequest(-1))}) // ObjectsServicesRequestObj | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsServicesAPI.PostObjectsServices(context.Background()).ObjectsServicesRequestObj(objectsServicesRequestObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsServicesAPI.PostObjectsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectsServices`: PostObjectsServices201Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsServicesAPI.PostObjectsServices`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectsServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectsServicesRequestObj** | [**ObjectsServicesRequestObj**](ObjectsServicesRequestObj.md) |  | 

### Return type

[**PostObjectsServices201Response**](PostObjectsServices201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectsServices

> ObjectsServicesResponseObj PutObjectsServices(ctx, objectId).ObjectsServicesRequestObj(objectsServicesRequestObj).Execute()

Update Service object



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
	objectId := "objectId_example" // string | 
	objectsServicesRequestObj := *openapiclient.NewObjectsServicesRequestObj("Name_example", []openapiclient.ObjectsServicesProtocolRequestObj{*openapiclient.NewObjectsServicesProtocolRequestObj("Protocol_example", "ValueType_example", []int32{int32(123)}, openapiclient.ObjectServiceProtocolOptionsICMPrequest(-1))}) // ObjectsServicesRequestObj | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsServicesAPI.PutObjectsServices(context.Background(), objectId).ObjectsServicesRequestObj(objectsServicesRequestObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsServicesAPI.PutObjectsServices``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutObjectsServices`: ObjectsServicesResponseObj
	fmt.Fprintf(os.Stdout, "Response from `ObjectsServicesAPI.PutObjectsServices`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectsServicesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectsServicesRequestObj** | [**ObjectsServicesRequestObj**](ObjectsServicesRequestObj.md) |  | 

### Return type

[**ObjectsServicesResponseObj**](ObjectsServicesResponseObj.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

