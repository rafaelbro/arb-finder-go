package main

import (
	"arb-finder/src/arbitrate"
	"arb-finder/src/bscconnector"
	"arb-finder/src/util"
	"log"

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

	for {
		select {
		case err := <-subscription.Err():
			log.Fatal(err)
		case <-blockHeadersChan:
			token0 := util.Tokens[token0Keys[token0Idx]]
			token1 := util.Tokens[token1Keys[token1Idx]]
			pairPool := util.Pairs[token0Keys[token0Idx]][token1Keys[token1Idx]]
			token0Amount := util.TradeQuantity[token0Keys[token0Idx]]
			token1Amount := util.TradeQuantity[token1Keys[token1Idx]]

			arbitrate.Arbitrate(token0, token0Amount, token1, token1Amount, pairPool)

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
