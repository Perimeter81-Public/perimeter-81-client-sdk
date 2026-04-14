# \EnhancedTunnelsAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicTunnel**](EnhancedTunnelsAPI.md#CreateDynamicTunnel) | **Post** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/dynamic | Create dynamic IPSec tunnel
[**CreateStaticTunnel**](EnhancedTunnelsAPI.md#CreateStaticTunnel) | **Post** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/static | Create static IPSec tunnel
[**DeleteDynamicTunnel**](EnhancedTunnelsAPI.md#DeleteDynamicTunnel) | **Delete** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/dynamic/{dynamicTunnelId} | Delete dynamic IPSec tunnel
[**DeleteStaticTunnel**](EnhancedTunnelsAPI.md#DeleteStaticTunnel) | **Delete** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/static/{tunnelId} | Delete static IPSec tunnel
[**GetDynamicTunnel**](EnhancedTunnelsAPI.md#GetDynamicTunnel) | **Get** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/dynamic/{dynamicTunnelId} | Get dynamic IPSec tunnel
[**GetEnhancedRegionTunnelsPerNetwork**](EnhancedTunnelsAPI.md#GetEnhancedRegionTunnelsPerNetwork) | **Get** /v2.3/networks/enhanced/{networkId}/tunnels | Get tunnels for a network
[**GetEnhancedRegionTunnelsPerRegion**](EnhancedTunnelsAPI.md#GetEnhancedRegionTunnelsPerRegion) | **Get** /v2.3/networks/enhanced/{networkId}/regions/{regionId}/tunnels | Get tunnels for a region
[**GetStaticTunnel**](EnhancedTunnelsAPI.md#GetStaticTunnel) | **Get** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/static/{tunnelId} | Get static IPSec tunnel
[**UpdateDynamicTunnel**](EnhancedTunnelsAPI.md#UpdateDynamicTunnel) | **Put** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/dynamic/{dynamicTunnelId} | Update dynamic IPSec tunnel
[**UpdateStaticTunnel**](EnhancedTunnelsAPI.md#UpdateStaticTunnel) | **Put** /v2.3/networks/enhanced/{networkId}/tunnels/ipsec/static/{tunnelId} | Update static IPSec tunnel



## CreateDynamicTunnel

> AsyncOperationResponse CreateDynamicTunnel(ctx, networkId).DynamicTunnelCreate(dynamicTunnelCreate).Execute()

Create dynamic IPSec tunnel



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
	dynamicTunnelCreate := *openapiclient.NewDynamicTunnelCreate("TunnelName_example", []openapiclient.DynamicTunnelDetails{*openapiclient.NewDynamicTunnelDetails("RegionID_example")}, *openapiclient.NewEnhancedIPSecSharedSettingsCreate([]string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}, *openapiclient.NewNetworkFeaturesCreate(), *openapiclient.NewRemoteASN()), *openapiclient.NewIPSecAdvancedSettingsV23("KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfigV23([]string{"Auth_example"}, []string{"Encryption_example"}, []string{"KeyExchangeMethod_example"}), *openapiclient.NewIPSecPhaseConfigV23([]string{"Auth_example"}, []string{"Encryption_example"}, []string{"KeyExchangeMethod_example"}))) // DynamicTunnelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.CreateDynamicTunnel(context.Background(), networkId).DynamicTunnelCreate(dynamicTunnelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.CreateDynamicTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.CreateDynamicTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **dynamicTunnelCreate** | [**DynamicTunnelCreate**](DynamicTunnelCreate.md) |  | 

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


## CreateStaticTunnel

> AsyncOperationResponse CreateStaticTunnel(ctx, networkId).StaticTunnelCreate(staticTunnelCreate).Execute()

Create static IPSec tunnel



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
	staticTunnelCreate := *openapiclient.NewStaticTunnelCreate("RegionID_example", "TunnelName_example", []string{"P81GatewaySubnets_example"}, []string{"RemoteGatewaySubnets_example"}, "KeyExchange_example", "IkeLifeTime_example", "Lifetime_example", "DpdDelay_example", "DpdTimeout_example", *openapiclient.NewIPSecPhaseConfigV23([]string{"Auth_example"}, []string{"Encryption_example"}, []string{"KeyExchangeMethod_example"}), *openapiclient.NewIPSecPhaseConfigV23([]string{"Auth_example"}, []string{"Encryption_example"}, []string{"KeyExchangeMethod_example"})) // StaticTunnelCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.CreateStaticTunnel(context.Background(), networkId).StaticTunnelCreate(staticTunnelCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.CreateStaticTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStaticTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.CreateStaticTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaticTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **staticTunnelCreate** | [**StaticTunnelCreate**](StaticTunnelCreate.md) |  | 

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


## DeleteDynamicTunnel

> AsyncOperationResponse DeleteDynamicTunnel(ctx, networkId, dynamicTunnelId).Execute()

Delete dynamic IPSec tunnel



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
	dynamicTunnelId := "dynamicTunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.DeleteDynamicTunnel(context.Background(), networkId, dynamicTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.DeleteDynamicTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteDynamicTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.DeleteDynamicTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**dynamicTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteDynamicTunnelRequest struct via the builder pattern


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


## DeleteStaticTunnel

> AsyncOperationResponse DeleteStaticTunnel(ctx, networkId, tunnelId).Execute()

Delete static IPSec tunnel



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
	resp, r, err := apiClient.EnhancedTunnelsAPI.DeleteStaticTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.DeleteStaticTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteStaticTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.DeleteStaticTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteStaticTunnelRequest struct via the builder pattern


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


## GetDynamicTunnel

> []EnhancedTunnel GetDynamicTunnel(ctx, networkId, dynamicTunnelId).Execute()

Get dynamic IPSec tunnel



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
	dynamicTunnelId := "dynamicTunnelId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.GetDynamicTunnel(context.Background(), networkId, dynamicTunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.GetDynamicTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetDynamicTunnel`: []EnhancedTunnel
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.GetDynamicTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**dynamicTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetDynamicTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**[]EnhancedTunnel**](EnhancedTunnel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedRegionTunnelsPerNetwork

> GetEnhancedRegionTunnelsPerNetwork200Response GetEnhancedRegionTunnelsPerNetwork(ctx, networkId).Page(page).Limit(limit).Execute()

Get tunnels for a network



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
	page := int32(56) // int32 | Specifies the current page of the paginated result set (optional)
	limit := int32(56) // int32 | Defines the number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerNetwork(context.Background(), networkId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerNetwork``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedRegionTunnelsPerNetwork`: GetEnhancedRegionTunnelsPerNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerNetwork`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedRegionTunnelsPerNetworkRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specifies the current page of the paginated result set | 
 **limit** | **int32** | Defines the number of results per page | 

### Return type

[**GetEnhancedRegionTunnelsPerNetwork200Response**](GetEnhancedRegionTunnelsPerNetwork200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetEnhancedRegionTunnelsPerRegion

> GetEnhancedRegionTunnelsPerNetwork200Response GetEnhancedRegionTunnelsPerRegion(ctx, networkId, regionId).Page(page).Limit(limit).Execute()

Get tunnels for a region



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
	page := int32(56) // int32 | Specifies the current page of the paginated result set (optional)
	limit := int32(56) // int32 | Defines the number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerRegion(context.Background(), networkId, regionId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerRegion``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedRegionTunnelsPerRegion`: GetEnhancedRegionTunnelsPerNetwork200Response
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.GetEnhancedRegionTunnelsPerRegion`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**regionId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedRegionTunnelsPerRegionRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **page** | **int32** | Specifies the current page of the paginated result set | 
 **limit** | **int32** | Defines the number of results per page | 

### Return type

[**GetEnhancedRegionTunnelsPerNetwork200Response**](GetEnhancedRegionTunnelsPerNetwork200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetStaticTunnel

> EnhancedTunnel GetStaticTunnel(ctx, networkId, tunnelId).Execute()

Get static IPSec tunnel



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
	resp, r, err := apiClient.EnhancedTunnelsAPI.GetStaticTunnel(context.Background(), networkId, tunnelId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.GetStaticTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetStaticTunnel`: EnhancedTunnel
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.GetStaticTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetStaticTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnhancedTunnel**](EnhancedTunnel.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateDynamicTunnel

> AsyncOperationResponse UpdateDynamicTunnel(ctx, networkId, dynamicTunnelId).DynamicTunnelUpdate(dynamicTunnelUpdate).Execute()

Update dynamic IPSec tunnel



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
	dynamicTunnelId := "dynamicTunnelId_example" // string | 
	dynamicTunnelUpdate := *openapiclient.NewDynamicTunnelUpdate("TunnelName_example") // DynamicTunnelUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.UpdateDynamicTunnel(context.Background(), networkId, dynamicTunnelId).DynamicTunnelUpdate(dynamicTunnelUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.UpdateDynamicTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateDynamicTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.UpdateDynamicTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**dynamicTunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateDynamicTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **dynamicTunnelUpdate** | [**DynamicTunnelUpdate**](DynamicTunnelUpdate.md) |  | 

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


## UpdateStaticTunnel

> AsyncOperationResponse UpdateStaticTunnel(ctx, networkId, tunnelId).StaticTunnelUpdate(staticTunnelUpdate).Execute()

Update static IPSec tunnel



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
	staticTunnelUpdate := *openapiclient.NewStaticTunnelUpdate() // StaticTunnelUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedTunnelsAPI.UpdateStaticTunnel(context.Background(), networkId, tunnelId).StaticTunnelUpdate(staticTunnelUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedTunnelsAPI.UpdateStaticTunnel``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateStaticTunnel`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedTunnelsAPI.UpdateStaticTunnel`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**tunnelId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateStaticTunnelRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **staticTunnelUpdate** | [**StaticTunnelUpdate**](StaticTunnelUpdate.md) |  | 

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

