/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type IssueTokenRequest struct {
	// Address issuing the token
	IssueAddress string `json:"issueAddress"`
	// Number of tokens to issue
	Amount float32 `json:"amount"`
	// Number of decimal places the token should be divisble by (0-7)
	Divisibility float32 `json:"divisibility"`
	// Fee in satoshi to include in the issuance transaction min 1000000000 (10 NEBL)
	Fee float32 `json:"fee"`
	// whether the token should be reissuable
	Reissuable bool `json:"reissuable"`
	Flags IssueTokenRequestFlags `json:"flags,omitempty"`
	Transfer []IssueTokenRequestTransfer `json:"transfer"`
	Metadata IssueTokenRequestMetadata `json:"metadata,omitempty"`
}
