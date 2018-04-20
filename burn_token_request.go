/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.1.1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type BurnTokenRequest struct {

	// Fee in satoshi to include in the issuance transaction min 10000 (0.0001 NEBL)
	Fee float32 `json:"fee"`

	// Array of addresses to send the token from
	From []string `json:"from,omitempty"`

	Transfer []SendTokenRequestTo `json:"transfer,omitempty"`

	// Array of objects representing tokens to be burned
	Burn []BurnTokenRequestBurn `json:"burn"`
}