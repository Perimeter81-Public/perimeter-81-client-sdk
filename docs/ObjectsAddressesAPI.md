# \ObjectsAddressesAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**DeleteObjectsAddresses**](ObjectsAddressesAPI.md#DeleteObjectsAddresses) | **Delete** /v2.3/objects/addresses/{objectId} | Delete Address object
[**GetObjectsAddresses**](ObjectsAddressesAPI.md#GetObjectsAddresses) | **Get** /v2.3/objects/addresses | Get object addresses
[**PostObjectsAddresses**](ObjectsAddressesAPI.md#PostObjectsAddresses) | **Post** /v2.3/objects/addresses | Create new Address object
[**PutObjectsAddresses**](ObjectsAddressesAPI.md#PutObjectsAddresses) | **Put** /v2.3/objects/addresses/{objectId} | Update Address object



## DeleteObjectsAddresses

> DeleteObjectsAddresses(ctx, objectId).Execute()

Delete Address object



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
	r, err := apiClient.ObjectsAddressesAPI.DeleteObjectsAddresses(context.Background(), objectId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAddressesAPI.DeleteObjectsAddresses``: %v\n", err)
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

Other parameters are passed through a pointer to a apiDeleteObjectsAddressesRequest struct via the builder pattern


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


## GetObjectsAddresses

> ObjectsAddressesResponse GetObjectsAddresses(ctx).Execute()

Get object addresses



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
	resp, r, err := apiClient.ObjectsAddressesAPI.GetObjectsAddresses(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAddressesAPI.GetObjectsAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetObjectsAddresses`: ObjectsAddressesResponse
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAddressesAPI.GetObjectsAddresses`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetObjectsAddressesRequest struct via the builder pattern


### Return type

[**ObjectsAddressesResponse**](ObjectsAddressesResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PostObjectsAddresses

> PostObjectsAddresses201Response PostObjectsAddresses(ctx).ObjectsAddressObj(objectsAddressObj).Execute()

Create new Address object



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
	objectsAddressObj := *openapiclient.NewObjectsAddressObj("Name_example", "ValueType_example", []string{"Value_example"}) // ObjectsAddressObj | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAddressesAPI.PostObjectsAddresses(context.Background()).ObjectsAddressObj(objectsAddressObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAddressesAPI.PostObjectsAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PostObjectsAddresses`: PostObjectsAddresses201Response
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAddressesAPI.PostObjectsAddresses`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiPostObjectsAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **objectsAddressObj** | [**ObjectsAddressObj**](ObjectsAddressObj.md) |  | 

### Return type

[**PostObjectsAddresses201Response**](PostObjectsAddresses201Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## PutObjectsAddresses

> ObjectsAddressObj PutObjectsAddresses(ctx, objectId).ObjectsAddressObj(objectsAddressObj).Execute()

Update Address object



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
	objectsAddressObj := *openapiclient.NewObjectsAddressObj("Name_example", "ValueType_example", []string{"Value_example"}) // ObjectsAddressObj | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.ObjectsAddressesAPI.PutObjectsAddresses(context.Background(), objectId).ObjectsAddressObj(objectsAddressObj).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `ObjectsAddressesAPI.PutObjectsAddresses``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `PutObjectsAddresses`: ObjectsAddressObj
	fmt.Fprintf(os.Stdout, "Response from `ObjectsAddressesAPI.PutObjectsAddresses`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**objectId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiPutObjectsAddressesRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **objectsAddressObj** | [**ObjectsAddressObj**](ObjectsAddressObj.md) |  | 

### Return type

[**ObjectsAddressObj**](ObjectsAddressObj.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

