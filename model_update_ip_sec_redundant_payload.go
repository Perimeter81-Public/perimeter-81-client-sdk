/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type UpdateIpSecRedundantPayload struct {
	Tunnel1          *IpSecRedundantTunnelDetails `json:"tunnel1,omitempty"`
	Tunnel2          *IpSecRedundantTunnelDetails `json:"tunnel2,omitempty"`
	SharedSettings   *IpSecSharedSettings         `json:"sharedSettings,omitempty"`
	AdvancedSettings *IpSecAdvancedSettings       `json:"advancedSettings,omitempty"`
}
