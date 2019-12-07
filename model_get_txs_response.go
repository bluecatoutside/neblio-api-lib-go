/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

// Object containing an array of transaction objects
type GetTxsResponse struct {
	// Number of pages of transactions
	PagesTotal float32 `json:"pagesTotal,omitempty"`
	// Array of transaction objects
	Txs []GetTxResponse `json:"txs,omitempty"`
}
