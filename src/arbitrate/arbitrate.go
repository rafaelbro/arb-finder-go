package arbitrate

import (
	"arb-finder/src/bscconnector"
	oneinchservice "arb-finder/src/one_inch_service"
	"arb-finder/src/util"
	"errors"
	"fmt"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const ()

func Arbitrate(token0 string, token0Amount int64, token1 string, token1Amount int64, pool string) int64 {
	token0AmountBig := util.ConvertToCryptoValue(token0Amount)
	token1AmountBig := util.ConvertToCryptoValue(token1Amount)
	reservesChan := make(chan *bscconnector.Reserve, 2)
	fromToken0Chan := make(chan int64, 1)
	fromToken1Chan := make(chan int64, 1)

	go bscconnector.Reserves(pool, reservesChan)

	// Call check for token0 to token
	go checkArbitragePossibility(token0, token1, token0AmountBig, reservesChan, 0, fromToken0Chan)

	// Call check for token1 to token0
	go checkArbitragePossibility(token1, token0, token1AmountBig, reservesChan, 1, fromToken1Chan)

	fromToken0Block := <-fromToken0Chan
	fromToken1Block := <-fromToken1Chan

	if fromToken0Block != 0 {
		return fromToken0Block
	}

	if fromToken1Block != 0 {
		return fromToken1Block
	}

	return 0
}

func checkArbitragePossibility(tokenFrom string, tokenTo string, amount *big.Int, reservesChan chan *bscconnector.Reserve, fromTokenIndex uint8, responseChan chan int64) {
	quote, err := oneinchservice.Quote(tokenFrom, tokenTo, amount)
	if err != nil {
		fmt.Println("1Inch Call Failed - Timed out")
		responseChan <- int64(0)
		return
	}

	fee := big.NewInt(10000 - (25 + 3)) // 0.25 + 0.03 = 0.28% fee
	toTokenAmount, _ := new(big.Int).SetString(quote.ToTokenAmount, 10)
	amountFloat := new(big.Float).SetInt(amount)
	reserves := <-reservesChan

	numerator := new(big.Int)
	denominator := new(big.Int)
	var liquidity *big.Float
	if fromTokenIndex == 0 {
		numerator.Mul(reserves.Reserve1, amount)
		numerator.Mul(numerator, big.NewInt(10000))

		denominator.Sub(reserves.Reserve0, amount)
		denominator.Mul(denominator, fee)
		liquidity = new(big.Float).Quo(amountFloat, new(big.Float).SetInt(reserves.Reserve0))
	} else {
		numerator.Mul(reserves.Reserve0, amount)
		numerator.Mul(numerator, big.NewInt(10000))

		denominator.Sub(reserves.Reserve1, amount)
		denominator.Mul(denominator, fee)
		liquidity = new(big.Float).Quo(amountFloat, new(big.Float).SetInt(reserves.Reserve1))
	}
	payableAmount := new(big.Int).Div(numerator, denominator)
	payableAmount.Add(payableAmount, big.NewInt(1))

	profit := new(big.Int).Sub(toTokenAmount, payableAmount)
	hasProfit := profit.Sign() > 0
	hasLiquidity := liquidity.Cmp(new(big.Float).SetFloat64(0.01)) == -1

	if hasProfit && hasLiquidity {
		routes, path, err := routersAndPath(quote)
		if err != nil {
			fmt.Printf("DEU PAU NAS ROTAS %s\n", err)
			responseChan <- int64(0)
			return
		}

		if os.Getenv("RUN") == "true" {
			bscconnector.StartArbitrage(amount, routes, *path, util.ContractAddress)

			currentBlock := bscconnector.CurrentBlock()
			responseChan <- int64(currentBlock)
			fmt.Printf("CHAMOUUUU StartArbitrage(%s, %s, %s) @ %d\n", amount, routes, path, currentBlock)
		}
		fmt.Printf("payableAmount = %s / toTokenAmount = %s / profit = %s / liquidity = %s \n",
			payableAmount, toTokenAmount, profit, liquidity)
		fmt.Printf("QUOTE: %s\n\n", quote)

	} else {
		responseChan <- int64(0)
	}
}

func routersAndPath(quote *oneinchservice.QuoteResponse) (*[]*big.Int, *[]common.Address, error) {
	var routes []*big.Int
	var path []common.Address
	var hasError error

	for i, cur := range quote.Protocols[0] {
		protocol := cur[0]
		if i == 0 {
			path = append(path, common.HexToAddress(protocol.FromTokenAddress))
		}

		path = append(path, common.HexToAddress(protocol.ToTokenAddress))
		if protocol.Name == "ACRYPTOS" {
			route, err := exchangeForACryptos(&protocol)
			if err == nil {
				routes = append(routes, big.NewInt(route))
			} else {
				hasError = err
			}
		} else if protocol.Name == "ELLIPSIS_FINANCE" {
			route, err := exchangeForEllipsis(&protocol)
			if err == nil {
				routes = append(routes, big.NewInt(route))
			} else {
				hasError = err
			}
		} else {
			routes = append(routes, big.NewInt(util.ExchangesMap[protocol.Name]))
		}
	}

	return &routes, &path, hasError
}

func exchangeForACryptos(protocol *oneinchservice.OneInchProtocol) (int64, error) {
	allowedTokenACryptos := map[string]bool{
		util.Tokens["BUSD"]: true,
		util.Tokens["USDT"]: true,
		util.Tokens["DAI"]:  true,
		util.Tokens["USDC"]: true,
		util.Tokens["VAI"]:  true,
	}

	if allowedTokenACryptos[protocol.FromTokenAddress] || allowedTokenACryptos[protocol.ToTokenAddress] {
		if protocol.FromTokenAddress == util.Tokens["VAI"] || protocol.ToTokenAddress == util.Tokens["VAI"] {
			return util.ExchangesMap["ACRYPTOS_META"], nil
		}
		return util.ExchangesMap["ACRYPTOS_CORE"], nil
	}

	return 254, errors.New("TOKEN NOT ALLOWED FOR ACRYPTOS")
}

func exchangeForEllipsis(protocol *oneinchservice.OneInchProtocol) (int64, error) {
	allowedTokenEllipsis := map[string]bool{
		util.Tokens["BUSD"]: true,
		util.Tokens["USDT"]: true,
		util.Tokens["DAI"]:  true,
		util.Tokens["USDC"]: true,
	}

	if allowedTokenEllipsis[protocol.FromTokenAddress] || allowedTokenEllipsis[protocol.ToTokenAddress] {
		if protocol.FromTokenAddress == util.Tokens["DAI"] || protocol.ToTokenAddress == util.Tokens["DAI"] {
			return util.ExchangesMap["ELLIPSIS_META"], nil
		}
		return util.ExchangesMap["ELLIPSIS_CORE"], nil
	}

	return 255, errors.New("TOKEN NOT ALLOWED FOR ELLIPSIS")
}
