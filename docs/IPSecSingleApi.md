# \IPSecSingleAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**StandardCreateIPSecSingleTunnel**](IPSecSingleAPI.md#StandardCreateIPSecSingleTunnel) | **Post** /v2.3/networks/standard/{networkId}/tunnels/ipsec/single | Create a new IPSec Single tunnel
[**StandardDeleteIPSecSingleTunnel**](IPSecSingleAPI.md#StandardDeleteIPSecSingleTunnel) | **Delete** /v2.3/networks/standard/{networkId}/tunnels/ipsec/single/{tunnelId} | Delete IPSec Single tunnel
[**StandardGetIPSecSingleTunnel**](IPSecSingleAPI.md#StandardGetIPSecSingleTunnel) | **Get** /v2.3/networks/standard/{networkId}/tunnels/ipsec/single/{tunnelId} | Get one IPSec Single tunnel
[**StandardUpdateIPSecSingleTunnel**](IPSecSingleAPI.md#StandardUpdateIPSecSingleTunnel) | **Put** /v2.3/networks/standard/{networkId}/tunnels/ipsec/single/{tunnelId} | Update IPSec Single Tunnel



## StandardCreateIPSecSingleTunnel

> AsyncOperationResponse StandardCreateIPSecSingleTunnel(ctx, networkId).CreateIPSecSinglePayload(createIPSecSinglePayload).Execute()

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
	resp, r, err := apiClient.IPSecSingleAPI.StandardCreateIPSecSingleTunnel(context.Background(), networkId).CreateIPSecSinglePayload(createIPSecSinglePayload).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecSingleAPI.StandardCreateIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardCreateIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecSingleAPI.StandardCreateIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardCreateIPSecSingleTunnelRequest struct via the builder pattern


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


## StandardDeleteIPSecSingleTunnel

> AsyncOperationResponse StandardDeleteIPSecSingleTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.IPSecSingleAPI.StandardDeleteIPSecSingleTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecSingleAPI.StandardDeleteIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardDeleteIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecSingleAPI.StandardDeleteIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardDeleteIPSecSingleTunnelRequest struct via the builder pattern


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


## StandardGetIPSecSingleTunnel

> IPSecSingleTunnel StandardGetIPSecSingleTunnel(ctx, networkId, tunnelId).Execute()

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
	resp, r, err := apiClient.IPSecSingleAPI.StandardGetIPSecSingleTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecSingleAPI.StandardGetIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardGetIPSecSingleTunnel`: IPSecSingleTunnel
	fmt.Fprintf(os.Stdout, "Response from `IPSecSingleAPI.StandardGetIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardGetIPSecSingleTunnelRequest struct via the builder pattern


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


## StandardUpdateIPSecSingleTunnel

> AsyncOperationResponse StandardUpdateIPSecSingleTunnel(ctx, networkId, tunnelId).IPSecSingleDetails(iPSecSingleDetails).Execute()

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
	resp, r, err := apiClient.IPSecSingleAPI.StandardUpdateIPSecSingleTunnel(context.Background(), networkId, tunnelId).IPSecSingleDetails(iPSecSingleDetails).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `IPSecSingleAPI.StandardUpdateIPSecSingleTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `StandardUpdateIPSecSingleTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `IPSecSingleAPI.StandardUpdateIPSecSingleTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiStandardUpdateIPSecSingleTunnelRequest struct via the builder pattern


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

