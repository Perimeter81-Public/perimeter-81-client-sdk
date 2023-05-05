/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type IpSecAdvancedSettings struct {
	KeyExchange string      `json:"keyExchange,omitempty"`
	IkeLifeTime string      `json:"ikeLifeTime,omitempty"`
	Lifetime    string      `json:"lifetime,omitempty"`
	DpdDelay    string      `json:"dpdDelay,omitempty"`
	DpdTimeout  string      `json:"dpdTimeout,omitempty"`
	Phase1      *IpSecPhase `json:"phase1,omitempty"`
	Phase2      *IpSecPhase `json:"phase2,omitempty"`
}
