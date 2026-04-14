# \StandardNetworksAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardGetNetworks**](StandardNetworksAPI.md#StandardGetNetworks) | **Get** /v2.3/networks/standard | Get all Networks
[**StandardNetworksControllerV2GetNetworkHealth**](StandardNetworksAPI.md#StandardNetworksControllerV2GetNetworkHealth) | **Get** /v2.3/networks/standard/{networkId}/health | Get network health status
[**StandardNetworksControllerV2NetworkCreate**](StandardNetworksAPI.md#StandardNetworksControllerV2NetworkCreate) | **Post** /v2.3/networks/standard | Create network
[**StandardNetworksControllerV2NetworkDelete**](StandardNetworksAPI.md#StandardNetworksControllerV2NetworkDelete) | **Delete** /v2.3/networks/standard/{networkId} | Delete network
[**StandardNetworksControllerV2NetworkFind**](StandardNetworksAPI.md#StandardNetworksControllerV2NetworkFind) | **Get** /v2.3/networks/standard/{networkId} | Get network by Id
[**StandardNetworksControllerV2NetworkUpdate**](StandardNetworksAPI.md#StandardNetworksControllerV2NetworkUpdate) | **Put** /v2.3/networks/standard/{networkId} | Update network
[**StandardNetworksControllerV2Status**](StandardNetworksAPI.md#StandardNetworksControllerV2Status) | **Get** /v2.3/networks/standard/status/{statusId} | Get status of asynchronous operations.



## StandardGetNetworks

> []Network StandardGetNetworks(ctx).Execute()

Get all Networks



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
	resp, r, err := apiClient.StandardNetworksAPI.StandardGetNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardGetNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetNetworks`: []Network
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardGetNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetNetworksRequest struct via the builder pattern


### Return type

[**[]Network**](Network.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2GetNetworkHealth

> StandardHealthResponse StandardNetworksControllerV2GetNetworkHealth(ctx, networkId).Execute()

Get network health status



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
	networkId := "networkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2GetNetworkHealth(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2GetNetworkHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2GetNetworkHealth`: StandardHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2GetNetworkHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2GetNetworkHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**StandardHealthResponse**](StandardHealthResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2NetworkCreate

> AsyncOperationResponse StandardNetworksControllerV2NetworkCreate(ctx).DeployNetworkPayload(deployNetworkPayload).Execute()

Create network



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
	deployNetworkPayload := *openapiclient.NewDeployNetworkPayload(*openapiclient.NewCreateNetworkPayload("Name_example"), []openapiclient.CreateRegionInNetworkPayload{*openapiclient.NewCreateRegionInNetworkPayload("HarmonySaseRegionId_example", false)}) // DeployNetworkPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2NetworkCreate(context.Background()).DeployNetworkPayload(deployNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2NetworkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2NetworkCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2NetworkCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2NetworkCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployNetworkPayload** | [**DeployNetworkPayload**](DeployNetworkPayload.md) |  | 

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


## StandardNetworksControllerV2NetworkDelete

> AsyncOperationResult StandardNetworksControllerV2NetworkDelete(ctx, networkId).Execute()

Delete network



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
	networkId := "networkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2NetworkDelete(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2NetworkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2NetworkDelete`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2NetworkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2NetworkDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationResult**](AsyncOperationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2NetworkFind

> Network StandardNetworksControllerV2NetworkFind(ctx, networkId).Execute()

Get network by Id



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
	networkId := "networkId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2NetworkFind(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2NetworkFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2NetworkFind`: Network
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2NetworkFind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2NetworkFindRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**Network**](Network.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2NetworkUpdate

> UpdatedNetwork StandardNetworksControllerV2NetworkUpdate(ctx, networkId).UpdateNetworkDto(updateNetworkDto).Execute()

Update network



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
	networkId := "networkId_example" // string | 
	updateNetworkDto := *openapiclient.NewUpdateNetworkDto(*openapiclient.NewBaseNetworkDto()) // UpdateNetworkDto | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2NetworkUpdate(context.Background(), networkId).UpdateNetworkDto(updateNetworkDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2NetworkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2NetworkUpdate`: UpdatedNetwork
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2NetworkUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2NetworkUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **updateNetworkDto** | [**UpdateNetworkDto**](UpdateNetworkDto.md) |  | 

### Return type

[**UpdatedNetwork**](UpdatedNetwork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2Status

> AsyncOperationStatus StandardNetworksControllerV2Status(ctx, statusId).Execute()

Get status of asynchronous operations.



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
	resp, r, err := apiClient.StandardNetworksAPI.StandardNetworksControllerV2Status(context.Background(), statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `StandardNetworksAPI.StandardNetworksControllerV2Status``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2Status`: AsyncOperationStatus
	fmt.Fprintf(os.Stdout, "Response from `StandardNetworksAPI.StandardNetworksControllerV2Status`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2StatusRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**AsyncOperationStatus**](AsyncOperationStatus.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

