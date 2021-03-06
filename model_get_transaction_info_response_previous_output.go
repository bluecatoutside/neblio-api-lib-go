/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetTransactionInfoResponsePreviousOutput struct {
	Asm string `json:"asm,omitempty"`
	Hex string `json:"hex,omitempty"`
	ReqSigs float32 `json:"reqSigs,omitempty"`
	Type string `json:"type,omitempty"`
	Addresses []string `json:"addresses,omitempty"`
}
