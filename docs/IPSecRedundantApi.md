# \IPSecRedundantAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardCreateIPSecRedundantTunnel**](IPSecRedundantAPI.md#StandardCreateIPSecRedundantTunnel) | **Post** /v2.3/networks/standard/{networkId}/tunnels/ipsec/redundant | Create a new IPSec Redundant tunnel
[**StandardDeleteIPSecRedundantTunnel**](IPSecRedundantAPI.md#StandardDeleteIPSecRedundantTunnel) | **Delete** /v2.3/networks/standard/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Delete IPSec Redundant tunnel
[**StandardGetIPSecRedundantTunnel**](IPSecRedundantAPI.md#StandardGetIPSecRedundantTunnel) | **Get** /v2.3/networks/standard/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Get one IPSec Redundant tunnel
[**StandardUpdateIPSecRedundantTunnel**](IPSecRedundantAPI.md#StandardUpdateIPSecRedundantTunnel) | **Put** /v2.3/networks/standard/{networkId}/tunnels/ipsec/redundant/{haTunnelId} | Update a new IPSec Redundant tunnel



## StandardCreateIPSecRedundantTunnel

> AsyncOperationResponse StandardCreateIPSecRedundantTunnel(ctx, networkId).CreateIPSecRedundantPayload(createIPSecRedundantPayload).Execute()

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
	resp, r, err := apiClient.IPSecRedundantAPI.StandardCreateIPSecRedundantTunnel(context.Background(), networkId).CreateIPSecRedundantPayload(createIPSecRedundantPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecRedundantAPI.StandardCreateIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardCreateIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecRedundantAPI.StandardCreateIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardCreateIPSecRedundantTunnelRequest struct via the builder pattern


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


## StandardDeleteIPSecRedundantTunnel

> AsyncOperationResponse StandardDeleteIPSecRedundantTunnel(ctx, networkId, haTunnelId).Execute()

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
	resp, r, err := apiClient.IPSecRedundantAPI.StandardDeleteIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecRedundantAPI.StandardDeleteIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardDeleteIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecRedundantAPI.StandardDeleteIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardDeleteIPSecRedundantTunnelRequest struct via the builder pattern


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


## StandardGetIPSecRedundantTunnel

> IPSecRedundantTunnels StandardGetIPSecRedundantTunnel(ctx, networkId, haTunnelId).Execute()

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
	resp, r, err := apiClient.IPSecRedundantAPI.StandardGetIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecRedundantAPI.StandardGetIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetIPSecRedundantTunnel`: IPSecRedundantTunnels
	fmt.Fprintf(os.Stdout, "Response from `IPSecRedundantAPI.StandardGetIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetIPSecRedundantTunnelRequest struct via the builder pattern


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


## StandardUpdateIPSecRedundantTunnel

> AsyncOperationResponse StandardUpdateIPSecRedundantTunnel(ctx, networkId, haTunnelId).UpdateIPSecRedundantPayload(updateIPSecRedundantPayload).Execute()

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
	resp, r, err := apiClient.IPSecRedundantAPI.StandardUpdateIPSecRedundantTunnel(context.Background(), networkId, haTunnelId).UpdateIPSecRedundantPayload(updateIPSecRedundantPayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecRedundantAPI.StandardUpdateIPSecRedundantTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardUpdateIPSecRedundantTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecRedundantAPI.StandardUpdateIPSecRedundantTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**haTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardUpdateIPSecRedundantTunnelRequest struct via the builder pattern


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

