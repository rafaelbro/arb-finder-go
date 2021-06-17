package main

import (
	"arb-finder/src/arbitrate"
	"arb-finder/src/bscconnector"
	"arb-finder/src/util"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func main() {
	token0Idx := 0
	token1Idx := 0
	token0Keys := make([]string, 0, len(util.Pairs))
	for k := range util.Pairs {
		token0Keys = append(token0Keys, k)
	}
	token1Keys := genToken1Keys(token0Keys, uint8(token0Idx))

	blockHeadersChan := make(chan *types.Header)
	subscription := bscconnector.SubscribeNewBlock(blockHeadersChan)
	fmt.Println("Listening to BSC New Blocks...")

	executedBlock := int64(0)

	for {
		select {
		case err := <-subscription.Err():
			fmt.Println(err)
		case header := <-blockHeadersChan:
			if big.NewInt(executedBlock+5).Cmp(header.Number) < 0 {

				token0 := token0Keys[token0Idx]
				token1 := token1Keys[token1Idx]
				token0Amount := util.TradeQuantity[token0Keys[token0Idx]]
				token1Amount := util.TradeQuantity[token1Keys[token1Idx]]

				executedBlock = arbitrate.Arbitrate(token0, token0Amount, token1, token1Amount)

				if token0Idx+1 == len(token0Keys) {
					token0Idx = 0
					token1Idx = 0
					token1Keys = genToken1Keys(token0Keys, uint8(token0Idx))
				} else if token1Idx+1 == len(token1Keys) {
					token0Idx += 1
					token1Idx = 0
					token1Keys = genToken1Keys(token0Keys, uint8(token0Idx))
				} else {
					token1Idx += 1
				}
			} else {
				fmt.Println("Contract Called - Skipping arbitrate")
			}

		}
	}
}

func genToken1Keys(token0Keys []string, idx uint8) []string {
	token1Keys := make([]string, 0, len(util.Pairs[token0Keys[idx]]))
	for k := range util.Pairs[token0Keys[idx]] {
		token1Keys = append(token1Keys, k)
	}
	return token1Keys
}
