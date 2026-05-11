# \GatewaysAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardGetInstance**](GatewaysAPI.md#StandardGetInstance) | **Get** /v2.3/networks/standard/{networkId}/instances/{gatewayId} | Get Instance of a GW by ID
[**StandardNetworksControllerV2AddNetworkInstance**](GatewaysAPI.md#StandardNetworksControllerV2AddNetworkInstance) | **Post** /v2.3/networks/standard/{networkId}/instances | Add gateway
[**StandardNetworksControllerV2DeleteNetworkInstance**](GatewaysAPI.md#StandardNetworksControllerV2DeleteNetworkInstance) | **Delete** /v2.3/networks/standard/{networkId}/instances | Remove Gateways from Network



## StandardGetInstance

> NetworkInstance StandardGetInstance(ctx, networkId, gatewayId).Execute()

Get Instance of a GW by ID



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
	gatewayId := "gatewayId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysAPI.StandardGetInstance(context.Background(), networkId, gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysAPI.StandardGetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetInstance`: NetworkInstance
	fmt.Fprintf(os.Stdout, "Response from `GatewaysAPI.StandardGetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkInstance**](NetworkInstance.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2AddNetworkInstance

> AsyncOperationResponse StandardNetworksControllerV2AddNetworkInstance(ctx, networkId).CreateInstancesInNetworkPayload(createInstancesInNetworkPayload).Execute()

Add gateway



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
	createInstancesInNetworkPayload := *openapiclient.NewCreateInstancesInNetworkPayload("RegionId_example", false) // CreateInstancesInNetworkPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysAPI.StandardNetworksControllerV2AddNetworkInstance(context.Background(), networkId).CreateInstancesInNetworkPayload(createInstancesInNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysAPI.StandardNetworksControllerV2AddNetworkInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2AddNetworkInstance`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `GatewaysAPI.StandardNetworksControllerV2AddNetworkInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2AddNetworkInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createInstancesInNetworkPayload** | [**CreateInstancesInNetworkPayload**](CreateInstancesInNetworkPayload.md) |  | 

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


## StandardNetworksControllerV2DeleteNetworkInstance

> AsyncOperationResult StandardNetworksControllerV2DeleteNetworkInstance(ctx, networkId).RemoveRegionInstance(removeRegionInstance).Execute()

Remove Gateways from Network



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
	removeRegionInstance := *openapiclient.NewRemoveRegionInstance([]openapiclient.RemoveRegionPayload{*openapiclient.NewRemoveRegionPayload()}) // RemoveRegionInstance | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.GatewaysAPI.StandardNetworksControllerV2DeleteNetworkInstance(context.Background(), networkId).RemoveRegionInstance(removeRegionInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `GatewaysAPI.StandardNetworksControllerV2DeleteNetworkInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2DeleteNetworkInstance`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `GatewaysAPI.StandardNetworksControllerV2DeleteNetworkInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2DeleteNetworkInstanceRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeRegionInstance** | [**RemoveRegionInstance**](RemoveRegionInstance.md) |  | 

### Return type

[**AsyncOperationResult**](AsyncOperationResult.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

