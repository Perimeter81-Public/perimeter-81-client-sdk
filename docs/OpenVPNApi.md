# \OpenVPNAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardCreateOpenVPNTunnel**](OpenVPNAPI.md#StandardCreateOpenVPNTunnel) | **Post** /v2.3/networks/standard/{networkId}/tunnels/openvpn | Create a new OpenVPN tunnel
[**StandardDeleteOpenVPNTunnel**](OpenVPNAPI.md#StandardDeleteOpenVPNTunnel) | **Delete** /v2.3/networks/standard/{networkId}/tunnels/openvpn/{tunnelId} | Delete OpenVPN tunnel
[**StandardGetOpenVPNTunnel**](OpenVPNAPI.md#StandardGetOpenVPNTunnel) | **Get** /v2.3/networks/standard/{networkId}/tunnels/openvpn/{tunnelId} | Get one openVPN tunnel
[**StandardUpdateOpenVPNTunnel**](OpenVPNAPI.md#StandardUpdateOpenVPNTunnel) | **Put** /v2.3/networks/standard/{networkId}/tunnels/openvpn/{tunnelId} | Update openVPN Tunnel



## StandardCreateOpenVPNTunnel

> OpenVPNAsyncOperationResponse StandardCreateOpenVPNTunnel(ctx, networkId).BaseTunnelValues(baseTunnelValues).Execute()

Create a new OpenVPN tunnel



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
	baseTunnelValues := *openapiclient.NewBaseTunnelValues("RegionID_example", "GatewayID_example", "TunnelName_example") // BaseTunnelValues | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenVPNAPI.StandardCreateOpenVPNTunnel(context.Background(), networkId).BaseTunnelValues(baseTunnelValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenVPNAPI.StandardCreateOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardCreateOpenVPNTunnel`: OpenVPNAsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenVPNAPI.StandardCreateOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardCreateOpenVPNTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **baseTunnelValues** | [**BaseTunnelValues**](BaseTunnelValues.md) |  | 

### Return type

[**OpenVPNAsyncOperationResponse**](OpenVPNAsyncOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardDeleteOpenVPNTunnel

> AsyncOperationResponse StandardDeleteOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

Delete OpenVPN tunnel



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
	tunnelId := "tunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenVPNAPI.StandardDeleteOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenVPNAPI.StandardDeleteOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardDeleteOpenVPNTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenVPNAPI.StandardDeleteOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardDeleteOpenVPNTunnelRequest struct via the builder pattern


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


## StandardGetOpenVPNTunnel

> OpenVPNTunnel StandardGetOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

Get one openVPN tunnel



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
	tunnelId := "tunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenVPNAPI.StandardGetOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenVPNAPI.StandardGetOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetOpenVPNTunnel`: OpenVPNTunnel
	fmt.Fprintf(os.Stdout, "Response from `OpenVPNAPI.StandardGetOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetOpenVPNTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenVPNTunnel**](OpenVPNTunnel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardUpdateOpenVPNTunnel

> OpenVPNAsyncOperationResponse StandardUpdateOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

Update openVPN Tunnel



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
	tunnelId := "tunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.OpenVPNAPI.StandardUpdateOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `OpenVPNAPI.StandardUpdateOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardUpdateOpenVPNTunnel`: OpenVPNAsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `OpenVPNAPI.StandardUpdateOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardUpdateOpenVPNTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**OpenVPNAsyncOperationResponse**](OpenVPNAsyncOperationResponse.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

