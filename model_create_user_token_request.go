/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc
// CreateUserTokenRequest struct for CreateUserTokenRequest
type CreateUserTokenRequest struct {
	Key string `json:"key,omitempty"`
	ExpirationMinues float32 `json:"expirationMinues,omitempty"`
}
