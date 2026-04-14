# \RegionsAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardGetRegion**](RegionsAPI.md#StandardGetRegion) | **Get** /v2.3/networks/standard/{networkId}/regions/{regionId} | Get region by ID
[**StandardNetworksControllerV2AddNetworkRegion**](RegionsAPI.md#StandardNetworksControllerV2AddNetworkRegion) | **Put** /v2.3/networks/standard/{networkId}/regions | Add regions to a network
[**StandardNetworksControllerV2DeleteNetworkRegion**](RegionsAPI.md#StandardNetworksControllerV2DeleteNetworkRegion) | **Delete** /v2.3/networks/standard/{networkId}/regions | Remove regions from network
[**StandardNetworksControllerV2GetRegions**](RegionsAPI.md#StandardNetworksControllerV2GetRegions) | **Get** /v2.3/networks/standard/harmony-sase-regions | List of available Harmony SASE regions that supports standard networks



## StandardGetRegion

> NetworkRegion StandardGetRegion(ctx, networkId, regionId).Execute()

Get region by ID



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
	regionId := "regionId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.StandardGetRegion(context.Background(), networkId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.StandardGetRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetRegion`: NetworkRegion
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.StandardGetRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**NetworkRegion**](NetworkRegion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardNetworksControllerV2AddNetworkRegion

> AsyncOperationResult StandardNetworksControllerV2AddNetworkRegion(ctx, networkId).CreateRegionInNetworkPayload(createRegionInNetworkPayload).Execute()

Add regions to a network



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
	createRegionInNetworkPayload := *openapiclient.NewCreateRegionInNetworkPayload("HarmonySaseRegionId_example", false) // CreateRegionInNetworkPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.StandardNetworksControllerV2AddNetworkRegion(context.Background(), networkId).CreateRegionInNetworkPayload(createRegionInNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.StandardNetworksControllerV2AddNetworkRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2AddNetworkRegion`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.StandardNetworksControllerV2AddNetworkRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2AddNetworkRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createRegionInNetworkPayload** | [**CreateRegionInNetworkPayload**](CreateRegionInNetworkPayload.md) |  | 

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


## StandardNetworksControllerV2DeleteNetworkRegion

> AsyncOperationResult StandardNetworksControllerV2DeleteNetworkRegion(ctx, networkId).RemoveRegionDTO(removeRegionDTO).Execute()

Remove regions from network



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
	removeRegionDTO := *openapiclient.NewRemoveRegionDTO("RegionId_example") // RemoveRegionDTO | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.RegionsAPI.StandardNetworksControllerV2DeleteNetworkRegion(context.Background(), networkId).RemoveRegionDTO(removeRegionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.StandardNetworksControllerV2DeleteNetworkRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2DeleteNetworkRegion`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.StandardNetworksControllerV2DeleteNetworkRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2DeleteNetworkRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **removeRegionDTO** | [**RemoveRegionDTO**](RemoveRegionDTO.md) |  | 

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


## StandardNetworksControllerV2GetRegions

> []Region StandardNetworksControllerV2GetRegions(ctx).Execute()

List of available Harmony SASE regions that supports standard networks



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
	resp, r, err := apiClient.RegionsAPI.StandardNetworksControllerV2GetRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `RegionsAPI.StandardNetworksControllerV2GetRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardNetworksControllerV2GetRegions`: []Region
	fmt.Fprintf(os.Stdout, "Response from `RegionsAPI.StandardNetworksControllerV2GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiStandardNetworksControllerV2GetRegionsRequest struct via the builder pattern


### Return type

[**[]Region**](Region.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

