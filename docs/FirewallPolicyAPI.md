# \FirewallPolicyAPI

All URIs are relative to *https://virtserver.swaggerhub.com/perimeter81/public-api-yaml/1.0.0*

Method | HTTP request | Description
------------- | ------------- | -------------
[**GetFirewallPolicy**](FirewallPolicyAPI.md#GetFirewallPolicy) | **Get** /v2.3/networks/{networkId}/policy | Get firewall policy by network Id
[**UpdateFirewallPolicy**](FirewallPolicyAPI.md#UpdateFirewallPolicy) | **Put** /v2.3/networks/{networkId}/policy | Update firewall policy by network ID



## GetFirewallPolicy

> FirewallPolicy GetFirewallPolicy(ctx, networkId).Execute()

Get firewall policy by network Id



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
	resp, r, err := apiClient.FirewallPolicyAPI.GetFirewallPolicy(context.Background(), networkId).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallPolicyAPI.GetFirewallPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `GetFirewallPolicy`: FirewallPolicy
	fmt.Fprintf(os.Stdout, "Response from `FirewallPolicyAPI.GetFirewallPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiGetFirewallPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**FirewallPolicy**](FirewallPolicy.md)

### Authorization

[bearer](../README.md#bearer)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## UpdateFirewallPolicy

> AsyncOperationResponse UpdateFirewallPolicy(ctx, networkId).FirewallPolicy(firewallPolicy).Execute()

Update firewall policy by network ID



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
	firewallPolicy := *openapiclient.NewFirewallPolicy(false, false, "Id_example", []openapiclient.FirewallPolicyRule{*openapiclient.NewFirewallPolicyRule("Name_example", false, false, openapiclient.SourcesAndDestinations{Addresses: openapiclient.NewAddresses([]string{"Addresses_example"})}, openapiclient.SourcesAndDestinations{Addresses: openapiclient.NewAddresses([]string{"Addresses_example"})})}) // FirewallPolicy | 

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.FirewallPolicyAPI.UpdateFirewallPolicy(context.Background(), networkId).FirewallPolicy(firewallPolicy).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `FirewallPolicyAPI.UpdateFirewallPolicy``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `UpdateFirewallPolicy`: AsyncOperationResponse
	fmt.Fprintf(os.Stdout, "Response from `FirewallPolicyAPI.UpdateFirewallPolicy`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**networkId** | **string** |  | 

### Other Parameters

Other parameters are passed through a pointer to a apiUpdateFirewallPolicyRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **firewallPolicy** | [**FirewallPolicy**](FirewallPolicy.md) |  | 

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

