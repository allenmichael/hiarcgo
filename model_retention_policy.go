/*
 * Hiarc API
 *
 * Welcome to the Hiarc API documentation.
 *
 * API version: 0.1.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package gohiarc

// RetentionPolicy struct for RetentionPolicy
import (
	"time"
)

type RetentionPolicy struct {
	Key         string                            `json:"key,omitempty"`
	Type        string                            `json:"type,omitempty"`
	Name        string                            `json:"name,omitempty"`
	Description string                            `json:"description,omitempty"`
	Metadata    map[string]map[string]interface{} `json:"metadata,omitempty"`
	CreatedBy   string                            `json:"createdBy,omitempty"`
	CreatedAt   time.Time                         `json:"createdAt,omitempty"`
	ModifiedAt  time.Time                         `json:"modifiedAt,omitempty"`
	Seconds     int32                             `json:"seconds,omitempty"`
}
