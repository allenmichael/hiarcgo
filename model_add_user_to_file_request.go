/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc
// AddUserToFileRequest struct for AddUserToFileRequest
type AddUserToFileRequest struct {
	UserKey string `json:"userKey,omitempty"`
	AccessLevel AccessLevel `json:"accessLevel,omitempty"`
}
