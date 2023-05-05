/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type NetworkTunnel struct {
	// ID of the network instance.
	Instance      string   `json:"instance"`
	InterfaceName string   `json:"interfaceName"`
	LeftAllowedIP []string `json:"leftAllowedIP"`
	LeftEndpoint  string   `json:"leftEndpoint"`
	// ID of the network.
	Network string `json:"network"`
	// ID of the network region.
	Region             string `json:"region"`
	RequestConfigToken string `json:"requestConfigToken"`
	Type_              string `json:"type"`
	// Unique ID.
	Id string `json:"id"`
	// ID of the tenant.
	TenantId string `json:"tenantId"`
	// The date when this record was created.
	CreatedAt string `json:"createdAt"`
	// The date of last update of the record.
	UpdatedAt string `json:"updatedAt,omitempty"`
}
