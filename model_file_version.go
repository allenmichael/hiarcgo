/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc
import (
	"time"
)
// FileVersion struct for FileVersion
type FileVersion struct {
	CreatedBy string `json:"createdBy,omitempty"`
	CreatedAt time.Time `json:"createdAt,omitempty"`
	StorageService string `json:"storageService,omitempty"`
	StorageId string `json:"storageId,omitempty"`
}
