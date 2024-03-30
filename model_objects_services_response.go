/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type ObjectsServicesResponse struct {
	Page       float64                      `json:"page"`
	TotalPage  float64                      `json:"totalPage"`
	ItemsTotal float64                      `json:"itemsTotal"`
	Data       []ObjectsServicesResponseObj `json:"data"`
}