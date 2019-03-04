/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.1
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetFaucetResponse struct {
	// Whether the withdrawal was successful
	Status string `json:"status,omitempty"`
	Data GetFaucetResponseData `json:"data,omitempty"`
}