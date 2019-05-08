/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

// Metadata set by user on token for UTXO
type GetTokenMetadataResponseMetadataOfUtxoUserData struct {
	Meta []GetTokenMetadataResponseMetadataOfIssuenceDataUserDataMeta `json:"meta,omitempty"`
}
