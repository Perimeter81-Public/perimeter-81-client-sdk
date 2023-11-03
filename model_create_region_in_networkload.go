/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type CreateRegionInNetworkload struct {
	// cpRegion ID.
	CpRegionId string `json:"regionId"`
	// Region ID.
	RegionID string `json:",omitempty"`
	// Create the gateway as disabled if true.
	Idle bool `json:"idle"`
	// Region name.
	Name string `json:",omitempty"`
	// Region dns.
	Dns string `json:",omitempty"`
	// Default gateway ip.
	DefaultGatewayIp string `json:",omitempty"`
}
