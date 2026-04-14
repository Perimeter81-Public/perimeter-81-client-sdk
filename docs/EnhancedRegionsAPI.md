# \EnhancedRegionsAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnhancedRegion**](EnhancedRegionsAPI.md#CreateEnhancedRegion) | **Post** /v2.3/networks/enhanced/{networkId}/regions | Create enhanced region
[**DeleteEnhancedRegion**](EnhancedRegionsAPI.md#DeleteEnhancedRegion) | **Delete** /v2.3/networks/enhanced/{networkId}/regions/{regionId} | Delete enhanced region
[**EnhancedNetworksControllerV2GetRegions**](EnhancedRegionsAPI.md#EnhancedNetworksControllerV2GetRegions) | **Get** /v2.3/networks/enhanced/harmony-sase-regions | List of available Harmony SASE regions that supports enhanced networks
[**GetEnhancedRegion**](EnhancedRegionsAPI.md#GetEnhancedRegion) | **Get** /v2.3/networks/enhanced/{networkId}/regions/{regionId} | Get region
[**IncreaseScaleUnit**](EnhancedRegionsAPI.md#IncreaseScaleUnit) | **Put** /v2.3/networks/enhanced/{networkId}/regions/{regionId}/scale-unit/increase | Increase scale units
[**ListEnhancedRegions**](EnhancedRegionsAPI.md#ListEnhancedRegions) | **Get** /v2.3/networks/enhanced/{networkId}/regions | List regions
[**ReduceScaleUnit**](EnhancedRegionsAPI.md#ReduceScaleUnit) | **Put** /v2.3/networks/enhanced/{networkId}/regions/{regionId}/scale-unit/reduce | Reduce scale units



## CreateEnhancedRegion

> AsyncOperationResponse CreateEnhancedRegion(ctx, networkId).EnhancedRegionCreate(enhancedRegionCreate).Execute()

Create enhanced region



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
	enhancedRegionCreate := *openapiclient.NewEnhancedRegionCreate("HarmonySaseRegionId_example") // EnhancedRegionCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRegionsAPI.CreateEnhancedRegion(context.Background(), networkId).EnhancedRegionCreate(enhancedRegionCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.CreateEnhancedRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnhancedRegion`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.CreateEnhancedRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnhancedRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enhancedRegionCreate** | [**EnhancedRegionCreate**](EnhancedRegionCreate.md) |  | 

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


## DeleteEnhancedRegion

> AsyncOperationResponse DeleteEnhancedRegion(ctx, networkId, regionId).Execute()

Delete enhanced region



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
	resp, r, err := apiClient.EnhancedRegionsAPI.DeleteEnhancedRegion(context.Background(), networkId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.DeleteEnhancedRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnhancedRegion`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.DeleteEnhancedRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnhancedRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**AsyncOperationResponse**](AsyncOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## EnhancedNetworksControllerV2GetRegions

> []HarmonySaseRegion EnhancedNetworksControllerV2GetRegions(ctx).Execute()

List of available Harmony SASE regions that supports enhanced networks



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
	resp, r, err := apiClient.EnhancedRegionsAPI.EnhancedNetworksControllerV2GetRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.EnhancedNetworksControllerV2GetRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnhancedNetworksControllerV2GetRegions`: []HarmonySaseRegion
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.EnhancedNetworksControllerV2GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnhancedNetworksControllerV2GetRegionsRequest struct via the builder pattern


### Return type

[**[]HarmonySaseRegion**](HarmonySaseRegion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedRegion

> EnhancedRegion GetEnhancedRegion(ctx, networkId, regionId).Execute()

Get region



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
	resp, r, err := apiClient.EnhancedRegionsAPI.GetEnhancedRegion(context.Background(), networkId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.GetEnhancedRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedRegion`: EnhancedRegion
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.GetEnhancedRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnhancedRegion**](EnhancedRegion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## IncreaseScaleUnit

> AsyncOperationResponse IncreaseScaleUnit(ctx, networkId, regionId).ScaleUnitsOperation(scaleUnitsOperation).Execute()

Increase scale units



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
	scaleUnitsOperation := *openapiclient.NewScaleUnitsOperation("UnitType_example", int32(123)) // ScaleUnitsOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRegionsAPI.IncreaseScaleUnit(context.Background(), networkId, regionId).ScaleUnitsOperation(scaleUnitsOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.IncreaseScaleUnit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `IncreaseScaleUnit`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.IncreaseScaleUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiIncreaseScaleUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scaleUnitsOperation** | [**ScaleUnitsOperation**](ScaleUnitsOperation.md) |  | 

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


## ListEnhancedRegions

> []EnhancedRegion ListEnhancedRegions(ctx, networkId).Execute()

List regions



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
	resp, r, err := apiClient.EnhancedRegionsAPI.ListEnhancedRegions(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.ListEnhancedRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ListEnhancedRegions`: []EnhancedRegion
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.ListEnhancedRegions`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiListEnhancedRegionsRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]EnhancedRegion**](EnhancedRegion.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## ReduceScaleUnit

> AsyncOperationResponse ReduceScaleUnit(ctx, networkId, regionId).ScaleUnitsOperation(scaleUnitsOperation).Execute()

Reduce scale units



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
	scaleUnitsOperation := *openapiclient.NewScaleUnitsOperation("UnitType_example", int32(123)) // ScaleUnitsOperation | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRegionsAPI.ReduceScaleUnit(context.Background(), networkId, regionId).ScaleUnitsOperation(scaleUnitsOperation).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRegionsAPI.ReduceScaleUnit``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `ReduceScaleUnit`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRegionsAPI.ReduceScaleUnit`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiReduceScaleUnitRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **scaleUnitsOperation** | [**ScaleUnitsOperation**](ScaleUnitsOperation.md) |  | 

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

