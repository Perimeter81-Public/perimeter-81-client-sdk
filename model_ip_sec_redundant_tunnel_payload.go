/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type IpSecRedundantTunnelPayload struct {
	Passphrase        string  `json:"passphrase,omitempty"`
	P81GWinternalIP   string  `json:"p81GWinternalIP,omitempty"`
	RemoteGWinernalIP string  `json:"remoteGWinernalIP,omitempty"`
	RemotePublicIP    string  `json:"remotePublicIP,omitempty"`
	RemoteASN         string `json:"remoteASN,omitempty"`
	RemoteID          string  `json:"remoteID,omitempty"`
	GatewayID         string  `json:"gatewayID,omitempty"`
}
