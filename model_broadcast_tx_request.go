/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type BroadcastTxRequest struct {
	// Signed raw tx hex to broadcast
	TxHex string `json:"txHex"`
}
