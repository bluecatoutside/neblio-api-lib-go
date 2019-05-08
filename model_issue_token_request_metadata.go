/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

// Object representing all metadata at token issuance
type IssueTokenRequestMetadata struct {
	// Token Symbol it will be identified by (ex. NIBBL)
	TokenName string `json:"tokenName,omitempty"`
	// Name of token issuer
	Issuer string `json:"issuer,omitempty"`
	// Long name or description of token (ex. Nibble)
	Description string `json:"description,omitempty"`
	Urls []IssueTokenRequestMetadataUrls `json:"urls,omitempty"`
	UserData GetTokenMetadataResponseMetadataOfIssuenceDataUserData `json:"userData,omitempty"`
	// Array of encryption instruction objects for encrypting userData
	Encryptions []IssueTokenRequestMetadataEncryptions `json:"encryptions,omitempty"`
	Rules IssueTokenRequestMetadataRules `json:"rules,omitempty"`
}
