/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type CreateIpSecSinglePayload struct {
	P81GatewaySubnets    []string `json:"p81GatewaySubnets"`
	RemoteGatewaySubnets []string `json:"remoteGatewaySubnets"`
	Passphrase           string   `json:"passphrase"`
	RemotePublicIP       string   `json:"remotePublicIP"`
	// RemoteID *Object `json:"remoteID,omitempty"`
	Phase1 *IpSecPhase `json:"phase1"`
	Phase2 *IpSecPhase `json:"phase2"`
	// Region ID
	RegionID string `json:"regionID"`
	// Gatwway ID
	GatewayID   string `json:"gatewayID"`
	TunnelName  string `json:"tunnelName"`
	TunnelID    string `json:"tunnelID,omitempty"`
	KeyExchange string `json:"keyExchange"`
	IkeLifeTime string `json:"ikeLifeTime"`
	Lifetime    string `json:"lifetime"`
	DpdDelay    string `json:"dpdDelay"`
	DpdTimeout  string `json:"dpdTimeout"`
}
