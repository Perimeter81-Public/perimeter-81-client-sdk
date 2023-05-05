/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type IpSecPhase struct {
	Auth       []string `json:"auth,omitempty"`
	Encryption []string `json:"encryption,omitempty"`
	// Diffie Helman encryption
	Dh []int32 `json:"dh,omitempty"`
}
