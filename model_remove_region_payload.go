/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type RemoveRegionPayload struct {
	// Region ID.
	RegionId  string                  `json:"regionId,omitempty"`
	Instances []RemoveInstancePayload `json:"instances,omitempty"`
}
