/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.2
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81sdk

type FirewallPolicyRule struct {
	Id string `json:"id,omitempty"`
	// Policy rule ID.
	Name string `json:"name"`
	// wether the rule is enabled.
	Enabled bool `json:"enabled"`
	// wether the rule is enabled.
	Allowed bool `json:"allowed"`
	Sources *SourcesAndDestinations `json:"sources"`
	Destinations *SourcesAndDestinations `json:"destinations"`
	// List of services objects IDs.
	Services []string `json:"services,omitempty"`
}