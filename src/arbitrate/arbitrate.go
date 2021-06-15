package arbitrate

import (
	"arb-finder/src/bscconnector"
	oneinchservice "arb-finder/src/one_inch_service"
	"arb-finder/src/util"
	"errors"
	"fmt"
	"log"
	"math/big"
	"os"

	"github.com/ethereum/go-ethereum/common"
)

const ()

func Arbitrate(token0 string, token0Amount int64, token1 string, token1Amount int64, pool string) {
	token0AmountBig := util.ConvertToCryptoValue(token0Amount)
	token1AmountBig := util.ConvertToCryptoValue(token1Amount)
	reservesChan := make(chan *bscconnector.Reserve, 2)

	go bscconnector.Reserves(pool, reservesChan)

	// Call check for token0 to token
	go checkArbitragePossibility(token0, token1, token0AmountBig, reservesChan, 0)

	// Call check for token1 to token0
	go checkArbitragePossibility(token1, token0, token1AmountBig, reservesChan, 1)
}

func checkArbitragePossibility(tokenFrom string, tokenTo string, amount *big.Int, reservesChan chan *bscconnector.Reserve, fromTokenIndex uint8) {
	quote, _ := oneinchservice.Quote(tokenFrom, tokenTo, amount)

	fee := new(big.Float).SetFloat64(1.045)
	amountFloat := new(big.Float).SetInt(amount)
	fromTokenFloat, _, _ := big.ParseFloat(quote.FromTokenAmount, 10, 7, big.ToPositiveInf)
	toTokenFloat, _, _ := big.ParseFloat(quote.ToTokenAmount, 10, 7, big.ToPositiveInf)
	amountRatio := big.NewFloat(0).Quo(fromTokenFloat, toTokenFloat)
	reserves := <-reservesChan
	var netLoanAmount, liquidity *big.Float
	if fromTokenIndex == 0 {
		netLoanAmount = new(big.Float).Quo(amountFloat, reserves.RationFrom0)
		liquidity = new(big.Float).Quo(amountFloat, new(big.Float).SetInt(reserves.Reserve0))
	} else {
		netLoanAmount = new(big.Float).Quo(amountFloat, reserves.RationFrom1)
		liquidity = new(big.Float).Quo(amountFloat, new(big.Float).SetInt(reserves.Reserve1))
	}

	loadAmountWithFee := new(big.Float).Mul(netLoanAmount, fee)
	tradeAmount := new(big.Float).Quo(amountFloat, amountRatio)
	profit := new(big.Float).Sub(tradeAmount, loadAmountWithFee)
	hasProfit := profit.Sign() > 0
	hasLiquidity := liquidity.Cmp(new(big.Float).SetFloat64(0.01)) == -1

	if hasProfit && hasLiquidity {
		routes, path, err := routersAndPath(quote)
		if err != nil {
			log.Println("DEU PAU NAS ROTAS %s", err)
			return
		}

		if os.Getenv("RUN") == "true" {
			bscconnector.StartArbitrage(amount, routes, *path, util.ContractAddress)
		}
		fmt.Printf("CHAMOUUUU StartArbitrage(%s, %s, %s) @ %d\n", amount, routes, path, bscconnector.CurrentBlock())
		log.Printf("netLoanAmount = %s / loadAmountWithFee = %s / tradeAmount = %s / profit = %s / liquidity = %s \n",
			netLoanAmount, loadAmountWithFee, tradeAmount, profit, liquidity)
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
			if err != nil {
				routes = append(routes, big.NewInt(route))
			} else {
				hasError = err
			}
		} else if protocol.Name == "ELLIPSIS_FINANCE" {
			route, err := exchangeForEllipsis(&protocol)
			if err != nil {
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
		return 255, errors.New("TOKEN NOT ALLOWED FOR ACRYPTOS")
	}

	if protocol.FromTokenAddress == util.Tokens["VAI"] || protocol.ToTokenAddress == util.Tokens["VAI"] {
		return util.ExchangesMap["ACRYPTOS_META"], nil
	}
	return util.ExchangesMap["ACRYPTOS_CORE"], nil
}

func exchangeForEllipsis(protocol *oneinchservice.OneInchProtocol) (int64, error) {
	allowedTokenEllipsis := map[string]bool{
		util.Tokens["BUSD"]: true,
		util.Tokens["USDT"]: true,
		util.Tokens["DAI"]:  true,
		util.Tokens["USDC"]: true,
	}

	if allowedTokenEllipsis[protocol.FromTokenAddress] || allowedTokenEllipsis[protocol.ToTokenAddress] {
		return 255, errors.New("TOKEN NOT ALLOWED FOR ELLIPSIS")
	}

	if protocol.FromTokenAddress == util.Tokens["DAI"] || protocol.ToTokenAddress == util.Tokens["DAI"] {
		return util.ExchangesMap["ELLIPSIS_META"], nil
	}
	return util.ExchangesMap["ELLIPSIS_CORE"], nil
}
