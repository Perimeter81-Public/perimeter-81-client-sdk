/*
 * Perimeter81 Public API
 *
 * The YAML for Preimeter81 Public API.
 *
 * API version: 1.0.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package perimeter81

type CreateUserPayloadModel struct {
	IdpType      string   `json:"idpType,omitempty"`
	AccessGroups []string `json:"accessGroups,omitempty"`
	// User email.
	Email string `json:"email"`
	// Skips email verification if set to `true`.
	NoEmailVerification bool `json:"noEmailVerification,omitempty"`
	// Invitation message that will be sent by email.
	InviteMessage string              `json:"inviteMessage"`
	ProfileData   *UserProfilePayload `json:"profileData,omitempty"`
}
