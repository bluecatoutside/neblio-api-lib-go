/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.2.3
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetBlockResponse struct {
	// Block hash
	Hash string `json:"hash,omitempty"`
	// Number of confirmations block has
	Confirmations float32 `json:"confirmations,omitempty"`
	// Block size in bytes
	Size float32 `json:"size,omitempty"`
	// Block height
	Height float32 `json:"height,omitempty"`
	// Block version
	Version float32 `json:"version,omitempty"`
	// Merkleroot of block
	Merkleroot string `json:"merkleroot,omitempty"`
	// Array of tx ids in the block
	Tx []string `json:"tx,omitempty"`
	// Block time relative to epoch
	Time float32 `json:"time,omitempty"`
	// Block nonce
	Nonce float32 `json:"nonce,omitempty"`
	// Block bits
	Bits string `json:"bits,omitempty"`
	// Block difficulty
	Difficulty float32 `json:"difficulty,omitempty"`
	// Hash of the previous block on the chain
	Previousblockhash string `json:"previousblockhash,omitempty"`
	// Hash of the next block on the chain
	Nextblockhash string `json:"nextblockhash,omitempty"`
	// Number of NEBL awarded in this block
	Reward float32 `json:"reward,omitempty"`
}
