package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	// ABI of the smart contract, used to decode event data
	contractABI = `[{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"},{"indexed":true,"internalType":"address","name":"owner","type":"address"},{"indexed":false,"internalType":"string","name":"tokenURI","type":"string"}],"name":"TokenMinted","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"},{"indexed":false,"internalType":"uint256","name":"price","type":"uint256"},{"indexed":true,"internalType":"address","name":"seller","type":"address"}],"name":"TokenListedForSale","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"uint256","name":"tokenId","type":"uint256"},{"indexed":true,"internalType":"address","name":"buyer","type":"address"},{"indexed":false,"internalType":"uint256","name":"price","type":"uint256"}],"name":"TokenSold","type":"event"}]`
)

func main() {
	// Connect to the Ethereum client using WebSocket
	client, err := ethclient.Dial("ws://localhost:8545")
	if err != nil {
		log.Fatal(err)
	}

	// Address of the deployed smart contract
	contractAddress := common.HexToAddress("0x0165878A594ca255338adfa4d48449f69242Eb8F")
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	// Create a channel to receive logs and subscribe to contract events
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	// Parse the contract ABI to decode event data
	parsedABI, err := abi.JSON(strings.NewReader(contractABI))
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			// Handle subscription errors
			log.Fatal(err)
		case vLog := <-logs:
			// Check if the log corresponds to the "TokenMinted" event
			if vLog.Topics[0].Hex() == parsedABI.Events["TokenMinted"].ID.Hex() {
				var eventData struct {
					TokenId  *big.Int
					Owner    common.Address
					TokenURI string
				}

				// Decode indexed and non-indexed event data
				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes())
				eventData.Owner = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenMinted", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// Print the details of the "TokenMinted" event
				fmt.Printf("Token Minted: TokenId=%s, Owner=%s, TokenURI=%s\n", eventData.TokenId.String(), eventData.Owner.Hex(), eventData.TokenURI)
			}

			// Check if the log corresponds to the "TokenListedForSale" event
			if vLog.Topics[0].Hex() == parsedABI.Events["TokenListedForSale"].ID.Hex() {
				var eventData struct {
					TokenId uint64
					Price   *big.Int
					Seller  common.Address
				}

				// Decode indexed and non-indexed event data
				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
				eventData.Seller = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenListedForSale", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// Print the details of the "TokenListedForSale" event
				fmt.Printf("Token Listed For Sale: TokenId=%d, Price=%s, Seller=%s\n", eventData.TokenId, eventData.Price.String(), eventData.Seller.Hex())
			}

			// Check if the log corresponds to the "TokenSold" event
			if vLog.Topics[0].Hex() == parsedABI.Events["TokenSold"].ID.Hex() {
				var eventData struct {
					TokenId uint64
					Buyer   common.Address
					Price   *big.Int
				}

				// Decode indexed and non-indexed event data
				eventData.TokenId = new(big.Int).SetBytes(vLog.Topics[1].Bytes()).Uint64()
				eventData.Buyer = common.HexToAddress(vLog.Topics[2].Hex())

				err := parsedABI.UnpackIntoInterface(&eventData, "TokenSold", vLog.Data)
				if err != nil {
					log.Fatal(err)
				}

				// Print the details of the "TokenSold" event
				fmt.Printf("Token Sold: TokenId=%d, Buyer=%s, Price=%s\n", eventData.TokenId, eventData.Buyer.Hex(), eventData.Price.String())
			}
		}
	}
}
