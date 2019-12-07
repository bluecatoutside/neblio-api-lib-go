/*
 * Neblio REST API Suite
 *
 * APIs for Interacting with NTP1 Tokens & The Neblio Blockchain
 *
 * API version: 1.3.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package neblioapi

type GetAddressResponse struct {
	// Address in string form
	AddrStr string `json:"addrStr,omitempty"`
	// NEBL balance
	Balance float32 `json:"balance,omitempty"`
	// NEBL balance in satoshis
	BalanceSat float32 `json:"balanceSat,omitempty"`
	// Total NEBL received
	TotalReceived float32 `json:"totalReceived,omitempty"`
	// Total NEBL received in satoshis
	TotalReceivedSat float32 `json:"totalReceivedSat,omitempty"`
	// Total NEBL sent
	TotalSent float32 `json:"totalSent,omitempty"`
	// Total NEBL sent satoshis
	TotalSentSat float32 `json:"totalSentSat,omitempty"`
	// Unconfirmed NEBL balance
	UnconfirmedBalance float32 `json:"unconfirmedBalance,omitempty"`
	// Unconfirmed NEBL balance in satoshis
	UnconfirmedBalanceSat float32 `json:"unconfirmedBalanceSat,omitempty"`
	// Number of unconfirmed transactions for this address
	UnconfirmedTxAppearances float32 `json:"unconfirmedTxAppearances,omitempty"`
	// Number of transactions for this address
	TxAppearances float32 `json:"txAppearances,omitempty"`
	// Array of transaction ids for this address
	Transactions []string `json:"transactions,omitempty"`
}
