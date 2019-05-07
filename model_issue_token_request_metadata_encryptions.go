/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.2
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type IssueTokenRequestMetadataEncryptions struct {
	// userData key to encrypt
	Key string `json:"key,omitempty"`
	// RSA public key used for encryption
	Pubkey string `json:"pubkey,omitempty"`
	// key format (pem or der)
	Format string `json:"format,omitempty"`
	// pkcs1 or pkcs8
	Type string `json:"type,omitempty"`
}
