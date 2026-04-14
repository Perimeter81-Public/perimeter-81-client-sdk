# \EnhancedNetworksAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateEnhancedNetwork**](EnhancedNetworksAPI.md#CreateEnhancedNetwork) | **Post** /v2.3/networks/enhanced | Create enhanced network
[**DeleteEnhancedNetwork**](EnhancedNetworksAPI.md#DeleteEnhancedNetwork) | **Delete** /v2.3/networks/enhanced/{networkId} | Delete enhanced network
[**EnhancedNetworksControllerV23GetNetworkCustomerCertificate**](EnhancedNetworksAPI.md#EnhancedNetworksControllerV23GetNetworkCustomerCertificate) | **Get** /v2.3/networks/enhanced/network-customer-certificate | Get network customer certificate
[**GetEnhancedNetwork**](EnhancedNetworksAPI.md#GetEnhancedNetwork) | **Get** /v2.3/networks/enhanced/{networkId} | Get enhanced network by ID
[**GetEnhancedNetworkHealth**](EnhancedNetworksAPI.md#GetEnhancedNetworkHealth) | **Get** /v2.3/networks/enhanced/{networkId}/health | Get enhanced network health status
[**GetEnhancedNetworks**](EnhancedNetworksAPI.md#GetEnhancedNetworks) | **Get** /v2.3/networks/enhanced | Get all enhanced networks
[**UpdateEnhancedNetwork**](EnhancedNetworksAPI.md#UpdateEnhancedNetwork) | **Put** /v2.3/networks/enhanced/{networkId} | Update enhanced network



## CreateEnhancedNetwork

> AsyncOperationResponse CreateEnhancedNetwork(ctx).DeployEnhancedNetwork(deployEnhancedNetwork).Execute()

Create enhanced network



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
	deployEnhancedNetwork := *openapiclient.NewDeployEnhancedNetwork(*openapiclient.NewDeployEnhancedNetworkNetwork("Name_example"), []openapiclient.EnhancedRegionCreate{*openapiclient.NewEnhancedRegionCreate("HarmonySaseRegionId_example")}) // DeployEnhancedNetwork | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedNetworksAPI.CreateEnhancedNetwork(context.Background()).DeployEnhancedNetwork(deployEnhancedNetwork).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.CreateEnhancedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateEnhancedNetwork`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.CreateEnhancedNetwork`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateEnhancedNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **deployEnhancedNetwork** | [**DeployEnhancedNetwork**](DeployEnhancedNetwork.md) |  | 

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


## DeleteEnhancedNetwork

> AsyncOperationResponse DeleteEnhancedNetwork(ctx, networkId).Execute()

Delete enhanced network



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
	resp, r, err := apiClient.EnhancedNetworksAPI.DeleteEnhancedNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.DeleteEnhancedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteEnhancedNetwork`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.DeleteEnhancedNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteEnhancedNetworkRequest struct via the builder pattern


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


## EnhancedNetworksControllerV23GetNetworkCustomerCertificate

> []NetworkCustomerCertificateResponseInner EnhancedNetworksControllerV23GetNetworkCustomerCertificate(ctx).Execute()

Get network customer certificate



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
	resp, r, err := apiClient.EnhancedNetworksAPI.EnhancedNetworksControllerV23GetNetworkCustomerCertificate(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.EnhancedNetworksControllerV23GetNetworkCustomerCertificate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `EnhancedNetworksControllerV23GetNetworkCustomerCertificate`: []NetworkCustomerCertificateResponseInner
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.EnhancedNetworksControllerV23GetNetworkCustomerCertificate`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiEnhancedNetworksControllerV23GetNetworkCustomerCertificateRequest struct via the builder pattern


### Return type

[**[]NetworkCustomerCertificateResponseInner**](NetworkCustomerCertificateResponseInner.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedNetwork

> EnhancedNetwork GetEnhancedNetwork(ctx, networkId).Execute()

Get enhanced network by ID



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
	resp, r, err := apiClient.EnhancedNetworksAPI.GetEnhancedNetwork(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.GetEnhancedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedNetwork`: EnhancedNetwork
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.GetEnhancedNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnhancedNetwork**](EnhancedNetwork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedNetworkHealth

> EnhancedHealthResponse GetEnhancedNetworkHealth(ctx, networkId).Execute()

Get enhanced network health status



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
	resp, r, err := apiClient.EnhancedNetworksAPI.GetEnhancedNetworkHealth(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.GetEnhancedNetworkHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedNetworkHealth`: EnhancedHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.GetEnhancedNetworkHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedNetworkHealthRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**EnhancedHealthResponse**](EnhancedHealthResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedNetworks

> []EnhancedNetwork GetEnhancedNetworks(ctx).Execute()

Get all enhanced networks



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
	resp, r, err := apiClient.EnhancedNetworksAPI.GetEnhancedNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.GetEnhancedNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedNetworks`: []EnhancedNetwork
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.GetEnhancedNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedNetworksRequest struct via the builder pattern


### Return type

[**[]EnhancedNetwork**](EnhancedNetwork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateEnhancedNetwork

> EnhancedNetwork UpdateEnhancedNetwork(ctx, networkId).EnhancedNetworkUpdate(enhancedNetworkUpdate).Execute()

Update enhanced network



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
	enhancedNetworkUpdate := *openapiclient.NewEnhancedNetworkUpdate() // EnhancedNetworkUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedNetworksAPI.UpdateEnhancedNetwork(context.Background(), networkId).EnhancedNetworkUpdate(enhancedNetworkUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedNetworksAPI.UpdateEnhancedNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateEnhancedNetwork`: EnhancedNetwork
	fmt.Fprintf(os.Stdout, "Response from `EnhancedNetworksAPI.UpdateEnhancedNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateEnhancedNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enhancedNetworkUpdate** | [**EnhancedNetworkUpdate**](EnhancedNetworkUpdate.md) |  | 

### Return type

[**EnhancedNetwork**](EnhancedNetwork.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

