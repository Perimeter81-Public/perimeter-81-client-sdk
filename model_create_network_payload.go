/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type CreateNetworkPayload struct {
	// Subnet of the network.
	Subnet string `json:"subnet,omitempty"`
	// Network name.
	Name string `json:"name"`
	// List of network tags.
	Tags []string `json:"tags"`
	// Network dns.
	Dns string `json:",omitempty"`
}
