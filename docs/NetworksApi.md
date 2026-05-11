# \NetworksAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateIPSecRedundantTunnel**](NetworksAPI.md#CreateIPSecRedundantTunnel) | **Post** /v2.3/networks/{networkId}/tunnels/ipsec/redundant | Create a new IPSec Redundant tunnel
[**CreateIPSecSingleTunnel**](NetworksAPI.md#CreateIPSecSingleTunnel) | **Post** /v2.3/networks/{networkId}/tunnels/ipsec/single | Create a new IPSec Single tunnel
[**CreateOpenVPNTunnel**](NetworksAPI.md#CreateOpenVPNTunnel) | **Post** /v2.3/networks/{networkId}/tunnels/openvpn | Create a new OpenVPN tunnel
[**CreateWireguardTunnel**](NetworksAPI.md#CreateWireguardTunnel) | **Post** /v2.3/networks/{networkId}/tunnels/wireguard | Create a new Wireguard tunnel
[**DeleteIPSecRedundantTunnel**](NetworksAPI.md#DeleteIPSecRedundantTunnel) | **Delete** /v2.3/networks/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Delete IPSec Redundant tunnel
[**DeleteIPSecSingleTunnel**](NetworksAPI.md#DeleteIPSecSingleTunnel) | **Delete** /v2.3/networks/{networkId}/tunnels/ipsec/single/{tunnelId} | Delete IPSec Single tunnel
[**DeleteOpenVPNTunnel**](NetworksAPI.md#DeleteOpenVPNTunnel) | **Delete** /v2.3/networks/{networkId}/tunnels/openvpn/{tunnelId} | Delete OpenVPN tunnel
[**DeleteWireguardTunnel**](NetworksAPI.md#DeleteWireguardTunnel) | **Delete** /v2.3/networks/{networkId}/tunnels/wireguard/{tunnelId} | Delete Wireguard tunnel
[**GetIPSecRedundantTunnel**](NetworksAPI.md#GetIPSecRedundantTunnel) | **Get** /v2.3/networks/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Get one IPSec Redundant tunnel
[**GetIPSecSingleTunnel**](NetworksAPI.md#GetIPSecSingleTunnel) | **Get** /v2.3/networks/{networkId}/tunnels/ipsec/single/{tunnelId} | Get one IPSec Single tunnel
[**GetInstance**](NetworksAPI.md#GetInstance) | **Get** /v2.3/networks/{networkId}/instances/{gatewayId} | Get Instance of a GW by ID
[**GetNetworks**](NetworksAPI.md#GetNetworks) | **Get** /v2.3/networks | Get all Networks
[**GetOpenVPNTunnel**](NetworksAPI.md#GetOpenVPNTunnel) | **Get** /v2.3/networks/{networkId}/tunnels/openvpn/{tunnelId} | Get one openVPN tunnel
[**GetRegion**](NetworksAPI.md#GetRegion) | **Get** /v2.3/networks/{networkId}/regions/{regionId} | Get region by ID
[**GetRouteTable**](NetworksAPI.md#GetRouteTable) | **Get** /v2.3/networks/{networkId}/route-table | Get route table
[**GetStatus**](NetworksAPI.md#GetStatus) | **Get** /v2.3/status | Get system status
[**GetWireguardTunnel**](NetworksAPI.md#GetWireguardTunnel) | **Get** /v2.3/networks/{networkId}/tunnels/wireguard/{tunnelId} | Get a Wireguard tunnel
[**NetworksControllerV2AddNetworkInstance**](NetworksAPI.md#NetworksControllerV2AddNetworkInstance) | **Post** /v2.3/networks/{networkId}/instances | Add gateway
[**NetworksControllerV2AddNetworkRegion**](NetworksAPI.md#NetworksControllerV2AddNetworkRegion) | **Put** /v2.3/networks/{networkId}/regions | Add regions to a network
[**NetworksControllerV2DeleteNetworkInstance**](NetworksAPI.md#NetworksControllerV2DeleteNetworkInstance) | **Delete** /v2.3/networks/{networkId}/instances | Remove Gateways from Network
[**NetworksControllerV2DeleteNetworkRegion**](NetworksAPI.md#NetworksControllerV2DeleteNetworkRegion) | **Delete** /v2.3/networks/{networkId}/regions | Remove regions from network
[**NetworksControllerV2GetNetworkHealth**](NetworksAPI.md#NetworksControllerV2GetNetworkHealth) | **Get** /v2.3/networks/{networkId}/health | Get network health status
[**NetworksControllerV2GetRegions**](NetworksAPI.md#NetworksControllerV2GetRegions) | **Get** /v2.3/networks/harmony-sase-regions | List of available regions
[**NetworksControllerV2NetworkCreate**](NetworksAPI.md#NetworksControllerV2NetworkCreate) | **Post** /v2.3/networks | Create network
[**NetworksControllerV2NetworkDelete**](NetworksAPI.md#NetworksControllerV2NetworkDelete) | **Delete** /v2.3/networks/{networkId} | Delete network
[**NetworksControllerV2NetworkFind**](NetworksAPI.md#NetworksControllerV2NetworkFind) | **Get** /v2.3/networks/{networkId} | Get network by Id
[**NetworksControllerV2NetworkUpdate**](NetworksAPI.md#NetworksControllerV2NetworkUpdate) | **Put** /v2.3/networks/{networkId} | Update network
[**NetworksControllerV2Status**](NetworksAPI.md#NetworksControllerV2Status) | **Get** /v2.3/networks/status/{statusId} | Get status of asynchronous operations.
[**UpdateIPSecRedundantTunnel**](NetworksAPI.md#UpdateIPSecRedundantTunnel) | **Put** /v2.3/networks/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Update a new IPSec Redundant tunnel
[**UpdateIPSecSingleTunnel**](NetworksAPI.md#UpdateIPSecSingleTunnel) | **Put** /v2.3/networks/{networkId}/tunnels/ipsec/single/{tunnelId} | Update IPSec Single Tunnel
[**UpdateOpenVPNTunnel**](NetworksAPI.md#UpdateOpenVPNTunnel) | **Put** /v2.3/networks/{networkId}/tunnels/openvpn/{tunnelId} | Update openVPN Tunnel
[**UpdateWireguardTunnel**](NetworksAPI.md#UpdateWireguardTunnel) | **Put** /v2.3/networks/{networkId}/tunnels/wireguard/{tunnelId} | Update a Wireguard tunnel



## CreateIPSecRedundantTunnel

> AsyncOperationResponse CreateIPSecRedundantTunnel(ctx, networkId).CreateIPSecRedundantPayload(createIPSecRedundantPayload).Execute()

Create a new IPSec Redundant tunnel



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
	createIPSecRedundantPayload := *openapiclient.NewCreateIPSecRedundantPayload("TunnelName_example", "RegionID_example", *openapiclient.NewIPSecRedundantTunnelPayload("Passphrase_example", "P81GWInternalIP_example", "RemoteGWInternalIP_example", "RemotePublicIP_example", *openapiclient.NewRemoteASN(), openapiclient.remoteID{String: new(string)}, "GatewayID_example"), *openapiclient.NewIPSecRedundantTunnelPayload("Passphrase_example", "P81GWInternalIP_example", "RemoteGWInternalIP_example", "RemotePublicIP_example", *openapiclient.NewRemoteASN(), openapiclient.remoteID{String: new(string)}, "GatewayID_example"), *openapiclient.NewIPSecSharedSettingsCreate([]string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}, *openapiclient.NewRemoteASN()), *openapiclient.NewIPSecAdvancedSettings("KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}))) // CreateIPSecRedundantPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateIPSecRedundantTunnel(context.Background(), networkId).CreateIPSecRedundantPayload(createIPSecRedundantPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPSecRedundantTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIPSecRedundantPayload** | [**CreateIPSecRedundantPayload**](CreateIPSecRedundantPayload.md) |  | 

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


## CreateIPSecSingleTunnel

> AsyncOperationResponse CreateIPSecSingleTunnel(ctx, networkId).CreateIPSecSinglePayload(createIPSecSinglePayload).Execute()

Create a new IPSec Single tunnel



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
	createIPSecSinglePayload := *openapiclient.NewCreateIPSecSinglePayload("RegionID_example", "GatewayID_example", "TunnelName_example", []string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}, "KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), "Passphrase_example", "RemotePublicIP_example") // CreateIPSecSinglePayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.CreateIPSecSingleTunnel(context.Background(), networkId).CreateIPSecSinglePayload(createIPSecSinglePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateIPSecSingleTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **createIPSecSinglePayload** | [**CreateIPSecSinglePayload**](CreateIPSecSinglePayload.md) |  | 

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


## CreateOpenVPNTunnel

> OpenVPNAsyncOperationResponse CreateOpenVPNTunnel(ctx, networkId).BaseTunnelValues(baseTunnelValues).Execute()

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
	resp, r, err := apiClient.NetworksAPI.CreateOpenVPNTunnel(context.Background(), networkId).BaseTunnelValues(baseTunnelValues).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateOpenVPNTunnel`: OpenVPNAsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateOpenVPNTunnelRequest struct via the builder pattern


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


## CreateWireguardTunnel

> AsyncOperationResponse CreateWireguardTunnel(ctx, networkId).CreateWireguardTunnelPayload(createWireguardTunnelPayload).Execute()

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
	resp, r, err := apiClient.NetworksAPI.CreateWireguardTunnel(context.Background(), networkId).CreateWireguardTunnelPayload(createWireguardTunnelPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.CreateWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.CreateWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateWireguardTunnelRequest struct via the builder pattern


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


## DeleteIPSecRedundantTunnel

> AsyncOperationResponse DeleteIPSecRedundantTunnel(ctx, networkId, haTunnelId).Execute()

Delete IPSec Redundant tunnel



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
	haTunnelId := "haTunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.DeleteIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPSecRedundantTunnelRequest struct via the builder pattern


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


## DeleteIPSecSingleTunnel

> AsyncOperationResponse DeleteIPSecSingleTunnel(ctx, networkId, tunnelId).Execute()

Delete IPSec Single tunnel



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
	resp, r, err := apiClient.NetworksAPI.DeleteIPSecSingleTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteIPSecSingleTunnelRequest struct via the builder pattern


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


## DeleteOpenVPNTunnel

> AsyncOperationResponse DeleteOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.DeleteOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteOpenVPNTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteOpenVPNTunnelRequest struct via the builder pattern


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


## DeleteWireguardTunnel

> AsyncOperationResponse DeleteWireguardTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.DeleteWireguardTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.DeleteWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.DeleteWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteWireguardTunnelRequest struct via the builder pattern


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


## GetIPSecRedundantTunnel

> IPSecRedundantTunnels GetIPSecRedundantTunnel(ctx, networkId, haTunnelId).Execute()

Get one IPSec Redundant tunnel



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
	haTunnelId := "haTunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.GetIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPSecRedundantTunnel`: IPSecRedundantTunnels
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPSecRedundantTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IPSecRedundantTunnels**](IPSecRedundantTunnels.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetIPSecSingleTunnel

> IPSecSingleTunnel GetIPSecSingleTunnel(ctx, networkId, tunnelId).Execute()

Get one IPSec Single tunnel



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
	resp, r, err := apiClient.NetworksAPI.GetIPSecSingleTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetIPSecSingleTunnel`: IPSecSingleTunnel
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetIPSecSingleTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**IPSecSingleTunnel**](IPSecSingleTunnel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetInstance

> NetworkInstance GetInstance(ctx, networkId, gatewayId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.GetInstance(context.Background(), networkId, gatewayId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetInstance`: NetworkInstance
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**gatewayId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetInstanceRequest struct via the builder pattern


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


## GetNetworks

> []Network GetNetworks(ctx).Execute()

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
	resp, r, err := apiClient.NetworksAPI.GetNetworks(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetNetworks``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetNetworks`: []Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetNetworks`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetNetworksRequest struct via the builder pattern


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


## GetOpenVPNTunnel

> OpenVPNTunnel GetOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.GetOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetOpenVPNTunnel`: OpenVPNTunnel
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetOpenVPNTunnelRequest struct via the builder pattern


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


## GetRegion

> NetworkRegion GetRegion(ctx, networkId, regionId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.GetRegion(context.Background(), networkId, regionId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRegion`: NetworkRegion
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRegionRequest struct via the builder pattern


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


## GetRouteTable

> []GetRouteTable200ResponseInner GetRouteTable(ctx, networkId).Execute()

Get route table



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
	resp, r, err := apiClient.NetworksAPI.GetRouteTable(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteTable`: []GetRouteTable200ResponseInner
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**[]GetRouteTable200ResponseInner**](GetRouteTable200ResponseInner.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStatus

> string GetStatus(ctx).Execute()

Get system status



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
	resp, r, err := apiClient.NetworksAPI.GetStatus(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetStatus``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStatus`: string
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetStatus`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiGetStatusRequest struct via the builder pattern


### Return type

**string**

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: text/plain, application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetWireguardTunnel

> WireguardTunnel GetWireguardTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.GetWireguardTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.GetWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetWireguardTunnel`: WireguardTunnel
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.GetWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetWireguardTunnelRequest struct via the builder pattern


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


## NetworksControllerV2AddNetworkInstance

> AsyncOperationResponse NetworksControllerV2AddNetworkInstance(ctx, networkId).CreateInstancesInNetworkPayload(createInstancesInNetworkPayload).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2AddNetworkInstance(context.Background(), networkId).CreateInstancesInNetworkPayload(createInstancesInNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2AddNetworkInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2AddNetworkInstance`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2AddNetworkInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2AddNetworkInstanceRequest struct via the builder pattern


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


## NetworksControllerV2AddNetworkRegion

> AsyncOperationResult NetworksControllerV2AddNetworkRegion(ctx, networkId).CreateRegionInNetworkPayload(createRegionInNetworkPayload).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2AddNetworkRegion(context.Background(), networkId).CreateRegionInNetworkPayload(createRegionInNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2AddNetworkRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2AddNetworkRegion`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2AddNetworkRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2AddNetworkRegionRequest struct via the builder pattern


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


## NetworksControllerV2DeleteNetworkInstance

> AsyncOperationResult NetworksControllerV2DeleteNetworkInstance(ctx, networkId).RemoveRegionInstance(removeRegionInstance).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2DeleteNetworkInstance(context.Background(), networkId).RemoveRegionInstance(removeRegionInstance).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2DeleteNetworkInstance``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2DeleteNetworkInstance`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2DeleteNetworkInstance`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2DeleteNetworkInstanceRequest struct via the builder pattern


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


## NetworksControllerV2DeleteNetworkRegion

> AsyncOperationResult NetworksControllerV2DeleteNetworkRegion(ctx, networkId).RemoveRegionDTO(removeRegionDTO).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2DeleteNetworkRegion(context.Background(), networkId).RemoveRegionDTO(removeRegionDTO).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2DeleteNetworkRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2DeleteNetworkRegion`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2DeleteNetworkRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2DeleteNetworkRegionRequest struct via the builder pattern


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


## NetworksControllerV2GetNetworkHealth

> StandardHealthResponse NetworksControllerV2GetNetworkHealth(ctx, networkId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2GetNetworkHealth(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2GetNetworkHealth``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2GetNetworkHealth`: StandardHealthResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2GetNetworkHealth`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2GetNetworkHealthRequest struct via the builder pattern


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


## NetworksControllerV2GetRegions

> []Region NetworksControllerV2GetRegions(ctx).Execute()

List of available regions



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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2GetRegions(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2GetRegions``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2GetRegions`: []Region
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2GetRegions`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2GetRegionsRequest struct via the builder pattern


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


## NetworksControllerV2NetworkCreate

> AsyncOperationResponse NetworksControllerV2NetworkCreate(ctx).DeployNetworkPayload(deployNetworkPayload).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2NetworkCreate(context.Background()).DeployNetworkPayload(deployNetworkPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2NetworkCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2NetworkCreate`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2NetworkCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2NetworkCreateRequest struct via the builder pattern


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


## NetworksControllerV2NetworkDelete

> AsyncOperationResult NetworksControllerV2NetworkDelete(ctx, networkId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2NetworkDelete(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2NetworkDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2NetworkDelete`: AsyncOperationResult
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2NetworkDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2NetworkDeleteRequest struct via the builder pattern


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


## NetworksControllerV2NetworkFind

> Network NetworksControllerV2NetworkFind(ctx, networkId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2NetworkFind(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2NetworkFind``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2NetworkFind`: Network
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2NetworkFind`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2NetworkFindRequest struct via the builder pattern


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


## NetworksControllerV2NetworkUpdate

> UpdatedNetwork NetworksControllerV2NetworkUpdate(ctx, networkId).UpdateNetworkDto(updateNetworkDto).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2NetworkUpdate(context.Background(), networkId).UpdateNetworkDto(updateNetworkDto).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2NetworkUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2NetworkUpdate`: UpdatedNetwork
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2NetworkUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2NetworkUpdateRequest struct via the builder pattern


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


## NetworksControllerV2Status

> AsyncOperationStatus NetworksControllerV2Status(ctx, statusId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.NetworksControllerV2Status(context.Background(), statusId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.NetworksControllerV2Status``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `NetworksControllerV2Status`: AsyncOperationStatus
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.NetworksControllerV2Status`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**statusId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiNetworksControllerV2StatusRequest struct via the builder pattern


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


## UpdateIPSecRedundantTunnel

> AsyncOperationResponse UpdateIPSecRedundantTunnel(ctx, networkId, haTunnelId).UpdateIPSecRedundantPayload(updateIPSecRedundantPayload).Execute()

Update a new IPSec Redundant tunnel



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
	haTunnelId := "haTunnelId_example" // string | 
	updateIPSecRedundantPayload := *openapiclient.NewUpdateIPSecRedundantPayload(*openapiclient.NewIPSecSharedSettings([]string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}), *openapiclient.NewIPSecAdvancedSettings("KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}))) // UpdateIPSecRedundantPayload | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).UpdateIPSecRedundantPayload(updateIPSecRedundantPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPSecRedundantTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **updateIPSecRedundantPayload** | [**UpdateIPSecRedundantPayload**](UpdateIPSecRedundantPayload.md) |  | 

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


## UpdateIPSecSingleTunnel

> AsyncOperationResponse UpdateIPSecSingleTunnel(ctx, networkId, tunnelId).IPSecSingleDetails(iPSecSingleDetails).Execute()

Update IPSec Single Tunnel



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
	iPSecSingleDetails := *openapiclient.NewIPSecSingleDetails("KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), *openapiclient.NewIPSecPhaseConfig([]string{"Auth_example"}, []string{"Encryption_example"}, []int32{int32(123)}), []string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}) // IPSecSingleDetails | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.NetworksAPI.UpdateIPSecSingleTunnel(context.Background(), networkId, tunnelId).IPSecSingleDetails(iPSecSingleDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateIPSecSingleTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **iPSecSingleDetails** | [**IPSecSingleDetails**](IPSecSingleDetails.md) |  | 

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


## UpdateOpenVPNTunnel

> OpenVPNAsyncOperationResponse UpdateOpenVPNTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.NetworksAPI.UpdateOpenVPNTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateOpenVPNTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateOpenVPNTunnel`: OpenVPNAsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateOpenVPNTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateOpenVPNTunnelRequest struct via the builder pattern


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


## UpdateWireguardTunnel

> AsyncOperationResponse UpdateWireguardTunnel(ctx, networkId, tunnelId).WireGuradDetails(wireGuradDetails).Execute()

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
	resp, r, err := apiClient.NetworksAPI.UpdateWireguardTunnel(context.Background(), networkId, tunnelId).WireGuradDetails(wireGuradDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `NetworksAPI.UpdateWireguardTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateWireguardTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `NetworksAPI.UpdateWireguardTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateWireguardTunnelRequest struct via the builder pattern


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

