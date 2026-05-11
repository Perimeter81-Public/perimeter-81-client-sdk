# \EnhancedRouteTablesAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateDynamicRoute**](EnhancedRouteTablesAPI.md#CreateDynamicRoute) | **Post** /v2.3/networks/enhanced/{networkId}/route-table/dynamic | Create dynamic route
[**CreateStaticRoute**](EnhancedRouteTablesAPI.md#CreateStaticRoute) | **Post** /v2.3/networks/enhanced/{networkId}/route-table/static | Create static route
[**DeleteRouteEntry**](EnhancedRouteTablesAPI.md#DeleteRouteEntry) | **Delete** /v2.3/networks/enhanced/{networkId}/route-table/{routeId} | Delete route entry
[**GetEnhancedRouteTable**](EnhancedRouteTablesAPI.md#GetEnhancedRouteTable) | **Get** /v2.3/networks/enhanced/{networkId}/route-table | Get route table
[**GetRouteEntry**](EnhancedRouteTablesAPI.md#GetRouteEntry) | **Get** /v2.3/networks/enhanced/{networkId}/route-table/{routeId} | Get route entry
[**UpdateRouteEntry**](EnhancedRouteTablesAPI.md#UpdateRouteEntry) | **Put** /v2.3/networks/enhanced/{networkId}/route-table/{routeId} | Update route entry



## CreateDynamicRoute

> AsyncOperationResponse CreateDynamicRoute(ctx, networkId).EnhancedRouteTableDynamicCreate(enhancedRouteTableDynamicCreate).Execute()

Create dynamic route



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
	enhancedRouteTableDynamicCreate := *openapiclient.NewEnhancedRouteTableDynamicCreate([]string{"TunnelIds_example"}, []string{"Subnets_example"}) // EnhancedRouteTableDynamicCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.CreateDynamicRoute(context.Background(), networkId).EnhancedRouteTableDynamicCreate(enhancedRouteTableDynamicCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.CreateDynamicRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateDynamicRoute`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.CreateDynamicRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateDynamicRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enhancedRouteTableDynamicCreate** | [**EnhancedRouteTableDynamicCreate**](EnhancedRouteTableDynamicCreate.md) |  | 

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


## CreateStaticRoute

> AsyncOperationResponse CreateStaticRoute(ctx, networkId).EnhancedRouteTableStaticCreate(enhancedRouteTableStaticCreate).Execute()

Create static route



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
	enhancedRouteTableStaticCreate := *openapiclient.NewEnhancedRouteTableStaticCreate("TunnelId_example", []string{"Subnets_example"}) // EnhancedRouteTableStaticCreate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.CreateStaticRoute(context.Background(), networkId).EnhancedRouteTableStaticCreate(enhancedRouteTableStaticCreate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.CreateStaticRoute``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateStaticRoute`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.CreateStaticRoute`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiCreateStaticRouteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **enhancedRouteTableStaticCreate** | [**EnhancedRouteTableStaticCreate**](EnhancedRouteTableStaticCreate.md) |  | 

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


## DeleteRouteEntry

> AsyncOperationResponse DeleteRouteEntry(ctx, networkId, routeId).Execute()

Delete route entry



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
	routeId := "routeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.DeleteRouteEntry(context.Background(), networkId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.DeleteRouteEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `DeleteRouteEntry`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.DeleteRouteEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiDeleteRouteEntryRequest struct via the builder pattern


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


## GetEnhancedRouteTable

> GetEnhancedRouteTable200Response GetEnhancedRouteTable(ctx, networkId).Page(page).Limit(limit).Execute()

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
	page := int32(56) // int32 | Specifies the current page of the paginated result set (optional)
	limit := int32(56) // int32 | Defines the number of results per page (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.GetEnhancedRouteTable(context.Background(), networkId).Page(page).Limit(limit).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.GetEnhancedRouteTable``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetEnhancedRouteTable`: GetEnhancedRouteTable200Response
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.GetEnhancedRouteTable`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetEnhancedRouteTableRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **page** | **int32** | Specifies the current page of the paginated result set | 
 **limit** | **int32** | Defines the number of results per page | 

### Return type

[**GetEnhancedRouteTable200Response**](GetEnhancedRouteTable200Response.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## GetRouteEntry

> EnhancedRouteTable GetRouteEntry(ctx, networkId, routeId).Execute()

Get route entry



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
	routeId := "routeId_example" // string | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.GetRouteEntry(context.Background(), networkId, routeId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.GetRouteEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetRouteEntry`: EnhancedRouteTable
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.GetRouteEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetRouteEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------



### Return type

[**EnhancedRouteTable**](EnhancedRouteTable.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateRouteEntry

> AsyncOperationResponse UpdateRouteEntry(ctx, networkId, routeId).RouteTableUpdate(routeTableUpdate).Execute()

Update route entry



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
	routeId := "routeId_example" // string | 
	routeTableUpdate := *openapiclient.NewRouteTableUpdate([]string{"Subnets_example"}) // RouteTableUpdate | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.EnhancedRouteTablesAPI.UpdateRouteEntry(context.Background(), networkId, routeId).RouteTableUpdate(routeTableUpdate).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `EnhancedRouteTablesAPI.UpdateRouteEntry``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateRouteEntry`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `EnhancedRouteTablesAPI.UpdateRouteEntry`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 
**routeId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateRouteEntryRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


 **routeTableUpdate** | [**RouteTableUpdate**](RouteTableUpdate.md) |  | 

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

