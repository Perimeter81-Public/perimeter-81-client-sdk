# \WireguardAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardCreateWireguardTunnel**](WireguardAPI.md#StandardCreateWireguardTunnel) | **Post** /v2.3/networks/standard/{networkId}/tunnels/wireguard | Create a new Wireguard tunnel
[**StandardDeleteWireguardTunnel**](WireguardAPI.md#StandardDeleteWireguardTunnel) | **Delete** /v2.3/networks/standard/{networkId}/tunnels/wireguard/{tunnelId} | Delete Wireguard tunnel
[**StandardGetWireguardTunnel**](WireguardAPI.md#StandardGetWireguardTunnel) | **Get** /v2.3/networks/standard/{networkId}/tunnels/wireguard/{tunnelId} | Get a Wireguard tunnel
[**StandardUpdateWireguardTunnel**](WireguardAPI.md#StandardUpdateWireguardTunnel) | **Put** /v2.3/networks/standard/{networkId}/tunnels/wireguard/{tunnelId} | Update a Wireguard tunnel



## StandardCreateWireguardTunnel

> AsyncOperationResponse StandardCreateWireguardTunnel(ctx, networkId).CreateWireguardTunnelPayload(createWireguardTunnelPayload).Execute()

Create a new Wireguard tunnel



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
	createWireguardTunnelPayload := *openapiclient.NewCreateWireguardTunnelPayload("RemoteEndpoint_example", []string{"RemoteSubnets_example"}, "RegionID_example", "GatewayID_example", "TunnelName_example") // CreateWireguardTunnelPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardAPI.StandardCreateWireguardTunnel(context.Background(), networkId).CreateWireguardTunnelPayload(createWireguardTunnelPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardAPI.StandardCreateWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardCreateWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WireguardAPI.StandardCreateWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardCreateWireguardTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createWireguardTunnelPayload** | [**CreateWireguardTunnelPayload**](CreateWireguardTunnelPayload.md) |  | 

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


## StandardDeleteWireguardTunnel

> AsyncOperationResponse StandardDeleteWireguardTunnel(ctx, networkId, tunnelId).Execute()

Delete Wireguard tunnel



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
	resp, r, err := apiClient.WireguardAPI.StandardDeleteWireguardTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardAPI.StandardDeleteWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardDeleteWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WireguardAPI.StandardDeleteWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardDeleteWireguardTunnelRequest struct via the builder pattern


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


## StandardGetWireguardTunnel

> WireguardTunnel StandardGetWireguardTunnel(ctx, networkId, tunnelId).Execute()

Get a Wireguard tunnel



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
	resp, r, err := apiClient.WireguardAPI.StandardGetWireguardTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardAPI.StandardGetWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetWireguardTunnel`: WireguardTunnel
	fmt.Fprintf(os.Stdout, "Response from `WireguardAPI.StandardGetWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetWireguardTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**WireguardTunnel**](WireguardTunnel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## StandardUpdateWireguardTunnel

> AsyncOperationResponse StandardUpdateWireguardTunnel(ctx, networkId, tunnelId).WireGuradDetails(wireGuradDetails).Execute()

Update a Wireguard tunnel



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
	wireGuradDetails := *openapiclient.NewWireGuradDetails("RemoteEndpoint_example", []string{"RemoteSubnets_example"}) // WireGuradDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.WireguardAPI.StandardUpdateWireguardTunnel(context.Background(), networkId, tunnelId).WireGuradDetails(wireGuradDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `WireguardAPI.StandardUpdateWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardUpdateWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `WireguardAPI.StandardUpdateWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardUpdateWireguardTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **wireGuradDetails** | [**WireGuradDetails**](WireGuradDetails.md) |  | 

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

