/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetAddressInfoResponseTokens struct {
	// Unique NTP1 identifier for this token
	TokenId string `json:"tokenId,omitempty"`
	// Number of Tokens
	Amount float32 `json:"amount,omitempty"`
	// TXID the token originally was issued in
	IssueTxid string `json:"issueTxid,omitempty"`
	// Decimal places the token is divisible to
	Divisibility float32 `json:"divisibility,omitempty"`
	// Whether the token is locked, preventing more from being issued
	LockStatus bool `json:"lockStatus,omitempty"`
	// Whether the tokens can be aggregated together
	AggregationPolicy string `json:"aggregationPolicy,omitempty"`
}
