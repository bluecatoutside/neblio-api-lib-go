/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type BurnTokenRequest struct {
	// Fee in satoshi to include in the issuance transaction min 10000 (0.0001 NEBL)
	Fee float32 `json:"fee"`
	// Array of addresses to send the token from
	From []string `json:"from,omitempty"`
	Transfer []SendTokenRequestTo `json:"transfer,omitempty"`
	// Array of objects representing tokens to be burned
	Burn []BurnTokenRequestBurn `json:"burn"`
}
