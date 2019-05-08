/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type Error struct {
	Code int32 `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
	Fields string `json:"fields,omitempty"`
}
