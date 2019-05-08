# Go API client for neblioapi

APIs for Interacting with NTP1 Tokens & The Neblio Blockchain

## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 1.2.3
- Package version: 1.2.1
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:
```
go get github.com/stretchr/testify/assert
go get golang.org/x/oauth2
go get golang.org/x/net/context
go get github.com/antihax/optional
```

Put the package under your project folder and add the following in import:
```golang
import "./neblioapi"
```

## Documentation for API Endpoints

All URIs are relative to *https://ntp1node.nebl.io*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*InsightApi* | [**GetAddress**](docs/InsightApi.md#getaddress) | **Get** /ins/addr/{address} | Returns address object
*InsightApi* | [**GetAddressBalance**](docs/InsightApi.md#getaddressbalance) | **Get** /ins/addr/{address}/balance | Returns address balance in sats
*InsightApi* | [**GetAddressTotalReceived**](docs/InsightApi.md#getaddresstotalreceived) | **Get** /ins/addr/{address}/totalReceived | Returns total received by address in sats
*InsightApi* | [**GetAddressTotalSent**](docs/InsightApi.md#getaddresstotalsent) | **Get** /ins/addr/{address}/totalSent | Returns total sent by address in sats
*InsightApi* | [**GetAddressUnconfirmedBalance**](docs/InsightApi.md#getaddressunconfirmedbalance) | **Get** /ins/addr/{address}/unconfirmedBalance | Returns address unconfirmed balance in sats
*InsightApi* | [**GetAddressUtxos**](docs/InsightApi.md#getaddressutxos) | **Get** /ins/addr/{address}/utxo | Returns all UTXOs at a given address
*InsightApi* | [**GetBlock**](docs/InsightApi.md#getblock) | **Get** /ins/block/{blockhash} | Returns information regarding a Neblio block
*InsightApi* | [**GetBlockIndex**](docs/InsightApi.md#getblockindex) | **Get** /ins/block-index/{blockindex} | Returns block hash of block
*InsightApi* | [**GetRawTx**](docs/InsightApi.md#getrawtx) | **Get** /ins/rawtx/{txid} | Returns raw transaction hex
*InsightApi* | [**GetStatus**](docs/InsightApi.md#getstatus) | **Get** /ins/status | Utility API for calling several blockchain node functions
*InsightApi* | [**GetSync**](docs/InsightApi.md#getsync) | **Get** /ins/sync | Get node sync status
*InsightApi* | [**GetTx**](docs/InsightApi.md#gettx) | **Get** /ins/tx/{txid} | Returns transaction object
*InsightApi* | [**GetTxs**](docs/InsightApi.md#gettxs) | **Get** /ins/txs | Get transactions by block or address
*InsightApi* | [**SendTx**](docs/InsightApi.md#sendtx) | **Post** /ins/tx/send | Broadcasts a signed raw transaction to the network (not NTP1 specific)
*JSONRPCApi* | [**JsonRpc**](docs/JSONRPCApi.md#jsonrpc) | **Post** / | Send a JSON-RPC call to a localhost neblio-Qt or nebliod node
*NTP1Api* | [**BroadcastTx**](docs/NTP1Api.md#broadcasttx) | **Post** /ntp1/broadcast | Broadcasts a signed raw transaction to the network
*NTP1Api* | [**BurnToken**](docs/NTP1Api.md#burntoken) | **Post** /ntp1/burntoken | Builds a transaction that burns an NTP1 Token
*NTP1Api* | [**GetAddressInfo**](docs/NTP1Api.md#getaddressinfo) | **Get** /ntp1/addressinfo/{address} | Information On a Neblio Address
*NTP1Api* | [**GetTokenHolders**](docs/NTP1Api.md#gettokenholders) | **Get** /ntp1/stakeholders/{tokenid} | Get Addresses Holding a Token
*NTP1Api* | [**GetTokenId**](docs/NTP1Api.md#gettokenid) | **Get** /ntp1/tokenid/{tokensymbol} | Returns the tokenId representing a token
*NTP1Api* | [**GetTokenMetadata**](docs/NTP1Api.md#gettokenmetadata) | **Get** /ntp1/tokenmetadata/{tokenid} | Get Metadata of Token
*NTP1Api* | [**GetTokenMetadataOfUtxo**](docs/NTP1Api.md#gettokenmetadataofutxo) | **Get** /ntp1/tokenmetadata/{tokenid}/{utxo} | Get UTXO Metadata of Token
*NTP1Api* | [**GetTransactionInfo**](docs/NTP1Api.md#gettransactioninfo) | **Get** /ntp1/transactioninfo/{txid} | Information On an NTP1 Transaction
*NTP1Api* | [**IssueToken**](docs/NTP1Api.md#issuetoken) | **Post** /ntp1/issue | Builds a transaction that issues a new NTP1 Token
*NTP1Api* | [**SendToken**](docs/NTP1Api.md#sendtoken) | **Post** /ntp1/sendtoken | Builds a transaction that sends an NTP1 Token
*TestnetFaucetApi* | [**TestnetGetFaucet**](docs/TestnetFaucetApi.md#testnetgetfaucet) | **Get** /testnet/faucet | Withdraws testnet NEBL to the specified address
*TestnetInsightApi* | [**TestnetGetAddress**](docs/TestnetInsightApi.md#testnetgetaddress) | **Get** /testnet/ins/addr/{address} | Returns address object
*TestnetInsightApi* | [**TestnetGetAddressBalance**](docs/TestnetInsightApi.md#testnetgetaddressbalance) | **Get** /testnet/ins/addr/{address}/balance | Returns address balance in sats
*TestnetInsightApi* | [**TestnetGetAddressTotalReceived**](docs/TestnetInsightApi.md#testnetgetaddresstotalreceived) | **Get** /testnet/ins/addr/{address}/totalReceived | Returns total received by address in sats
*TestnetInsightApi* | [**TestnetGetAddressTotalSent**](docs/TestnetInsightApi.md#testnetgetaddresstotalsent) | **Get** /testnet/ins/addr/{address}/totalSent | Returns total sent by address in sats
*TestnetInsightApi* | [**TestnetGetAddressUnconfirmedBalance**](docs/TestnetInsightApi.md#testnetgetaddressunconfirmedbalance) | **Get** /testnet/ins/addr/{address}/unconfirmedBalance | Returns address unconfirmed balance in sats
*TestnetInsightApi* | [**TestnetGetAddressUtxos**](docs/TestnetInsightApi.md#testnetgetaddressutxos) | **Get** /testnet/ins/addr/{address}/utxo | Returns all UTXOs at a given address
*TestnetInsightApi* | [**TestnetGetBlock**](docs/TestnetInsightApi.md#testnetgetblock) | **Get** /testnet/ins/block/{blockhash} | Returns information regarding a Neblio block
*TestnetInsightApi* | [**TestnetGetBlockIndex**](docs/TestnetInsightApi.md#testnetgetblockindex) | **Get** /testnet/ins/block-index/{blockindex} | Returns block hash of block
*TestnetInsightApi* | [**TestnetGetRawTx**](docs/TestnetInsightApi.md#testnetgetrawtx) | **Get** /testnet/ins/rawtx/{txid} | Returns raw transaction hex
*TestnetInsightApi* | [**TestnetGetStatus**](docs/TestnetInsightApi.md#testnetgetstatus) | **Get** /testnet/ins/status | Utility API for calling several blockchain node functions
*TestnetInsightApi* | [**TestnetGetSync**](docs/TestnetInsightApi.md#testnetgetsync) | **Get** /testnet/ins/sync | Get node sync status
*TestnetInsightApi* | [**TestnetGetTx**](docs/TestnetInsightApi.md#testnetgettx) | **Get** /testnet/ins/tx/{txid} | Returns transaction object
*TestnetInsightApi* | [**TestnetGetTxs**](docs/TestnetInsightApi.md#testnetgettxs) | **Get** /testnet/ins/txs | Get transactions by block or address
*TestnetInsightApi* | [**TestnetSendTx**](docs/TestnetInsightApi.md#testnetsendtx) | **Post** /testnet/ins/tx/send | Broadcasts a signed raw transaction to the network (not NTP1 specific)
*TestnetNTP1Api* | [**TestnetBroadcastTx**](docs/TestnetNTP1Api.md#testnetbroadcasttx) | **Post** /testnet/ntp1/broadcast | Broadcasts a signed raw transaction to the network
*TestnetNTP1Api* | [**TestnetBurnToken**](docs/TestnetNTP1Api.md#testnetburntoken) | **Post** /testnet/ntp1/burntoken | Builds a transaction that burns an NTP1 Token
*TestnetNTP1Api* | [**TestnetGetAddressInfo**](docs/TestnetNTP1Api.md#testnetgetaddressinfo) | **Get** /testnet/ntp1/addressinfo/{address} | Information On a Neblio Address
*TestnetNTP1Api* | [**TestnetGetTokenHolders**](docs/TestnetNTP1Api.md#testnetgettokenholders) | **Get** /testnet/ntp1/stakeholders/{tokenid} | Get Addresses Holding a Token
*TestnetNTP1Api* | [**TestnetGetTokenId**](docs/TestnetNTP1Api.md#testnetgettokenid) | **Get** /testnet/ntp1/tokenid/{tokensymbol} | Returns the tokenId representing a token
*TestnetNTP1Api* | [**TestnetGetTokenMetadata**](docs/TestnetNTP1Api.md#testnetgettokenmetadata) | **Get** /testnet/ntp1/tokenmetadata/{tokenid} | Get Metadata of Token
*TestnetNTP1Api* | [**TestnetGetTokenMetadataOfUtxo**](docs/TestnetNTP1Api.md#testnetgettokenmetadataofutxo) | **Get** /testnet/ntp1/tokenmetadata/{tokenid}/{utxo} | Get UTXO Metadata of Token
*TestnetNTP1Api* | [**TestnetGetTransactionInfo**](docs/TestnetNTP1Api.md#testnetgettransactioninfo) | **Get** /testnet/ntp1/transactioninfo/{txid} | Information On an NTP1 Transaction
*TestnetNTP1Api* | [**TestnetIssueToken**](docs/TestnetNTP1Api.md#testnetissuetoken) | **Post** /testnet/ntp1/issue | Builds a transaction that issues a new NTP1 Token
*TestnetNTP1Api* | [**TestnetSendToken**](docs/TestnetNTP1Api.md#testnetsendtoken) | **Post** /testnet/ntp1/sendtoken | Builds a transaction that sends an NTP1 Token


## Documentation For Models

 - [BroadcastTxRequest](docs/BroadcastTxRequest.md)
 - [BroadcastTxResponse](docs/BroadcastTxResponse.md)
 - [BurnTokenRequest](docs/BurnTokenRequest.md)
 - [BurnTokenRequestBurn](docs/BurnTokenRequestBurn.md)
 - [BurnTokenResponse](docs/BurnTokenResponse.md)
 - [Error](docs/Error.md)
 - [GetAddressInfoResponse](docs/GetAddressInfoResponse.md)
 - [GetAddressInfoResponseTokens](docs/GetAddressInfoResponseTokens.md)
 - [GetAddressInfoResponseUtxos](docs/GetAddressInfoResponseUtxos.md)
 - [GetAddressResponse](docs/GetAddressResponse.md)
 - [GetBlockIndexResponse](docs/GetBlockIndexResponse.md)
 - [GetBlockResponse](docs/GetBlockResponse.md)
 - [GetFaucetResponse](docs/GetFaucetResponse.md)
 - [GetFaucetResponseData](docs/GetFaucetResponseData.md)
 - [GetRawTxResponse](docs/GetRawTxResponse.md)
 - [GetSyncResponse](docs/GetSyncResponse.md)
 - [GetTokenHoldersResponse](docs/GetTokenHoldersResponse.md)
 - [GetTokenHoldersResponseHolders](docs/GetTokenHoldersResponseHolders.md)
 - [GetTokenIdResponse](docs/GetTokenIdResponse.md)
 - [GetTokenMetadataResponse](docs/GetTokenMetadataResponse.md)
 - [GetTokenMetadataResponseMetadataOfIssuence](docs/GetTokenMetadataResponseMetadataOfIssuence.md)
 - [GetTokenMetadataResponseMetadataOfIssuenceData](docs/GetTokenMetadataResponseMetadataOfIssuenceData.md)
 - [GetTokenMetadataResponseMetadataOfIssuenceDataUserData](docs/GetTokenMetadataResponseMetadataOfIssuenceDataUserData.md)
 - [GetTokenMetadataResponseMetadataOfIssuenceDataUserDataMeta](docs/GetTokenMetadataResponseMetadataOfIssuenceDataUserDataMeta.md)
 - [GetTokenMetadataResponseMetadataOfUtxo](docs/GetTokenMetadataResponseMetadataOfUtxo.md)
 - [GetTokenMetadataResponseMetadataOfUtxoUserData](docs/GetTokenMetadataResponseMetadataOfUtxoUserData.md)
 - [GetTransactionInfoResponse](docs/GetTransactionInfoResponse.md)
 - [GetTransactionInfoResponsePreviousOutput](docs/GetTransactionInfoResponsePreviousOutput.md)
 - [GetTransactionInfoResponseScriptSig](docs/GetTransactionInfoResponseScriptSig.md)
 - [GetTransactionInfoResponseTokens](docs/GetTransactionInfoResponseTokens.md)
 - [GetTransactionInfoResponseVin](docs/GetTransactionInfoResponseVin.md)
 - [GetTransactionInfoResponseVout](docs/GetTransactionInfoResponseVout.md)
 - [GetTxResponse](docs/GetTxResponse.md)
 - [GetTxResponseVin](docs/GetTxResponseVin.md)
 - [GetTxResponseVout](docs/GetTxResponseVout.md)
 - [GetTxsResponse](docs/GetTxsResponse.md)
 - [IssueTokenRequest](docs/IssueTokenRequest.md)
 - [IssueTokenRequestFlags](docs/IssueTokenRequestFlags.md)
 - [IssueTokenRequestMetadata](docs/IssueTokenRequestMetadata.md)
 - [IssueTokenRequestMetadataEncryptions](docs/IssueTokenRequestMetadataEncryptions.md)
 - [IssueTokenRequestMetadataRules](docs/IssueTokenRequestMetadataRules.md)
 - [IssueTokenRequestMetadataRulesExpiration](docs/IssueTokenRequestMetadataRulesExpiration.md)
 - [IssueTokenRequestMetadataRulesFees](docs/IssueTokenRequestMetadataRulesFees.md)
 - [IssueTokenRequestMetadataRulesFeesItems](docs/IssueTokenRequestMetadataRulesFeesItems.md)
 - [IssueTokenRequestMetadataRulesHolders](docs/IssueTokenRequestMetadataRulesHolders.md)
 - [IssueTokenRequestMetadataUrls](docs/IssueTokenRequestMetadataUrls.md)
 - [IssueTokenRequestTransfer](docs/IssueTokenRequestTransfer.md)
 - [IssueTokenResponse](docs/IssueTokenResponse.md)
 - [RpcRequest](docs/RpcRequest.md)
 - [RpcResponse](docs/RpcResponse.md)
 - [SendTokenRequest](docs/SendTokenRequest.md)
 - [SendTokenRequestTo](docs/SendTokenRequestTo.md)
 - [SendTokenResponse](docs/SendTokenResponse.md)
 - [SendTxRequest](docs/SendTxRequest.md)


## Documentation For Authorization

## rpcAuth
- **Type**: HTTP basic authentication

Example
```golang
auth := context.WithValue(context.Background(), sw.ContextBasicAuth, sw.BasicAuth{
	UserName: "username",
	Password: "password",
})
r, err := client.Service.Operation(auth, args)
```

## Author



