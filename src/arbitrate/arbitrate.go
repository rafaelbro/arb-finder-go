package arbitrate

import (
	"arb-finder/src/bscconnector"
	oneinchservice "arb-finder/src/one_inch_service"
	"arb-finder/src/util"
	"errors"
	"fmt"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/common"
)

const ()

type ContractCallEvent struct {
	PairAddress string
	Quote       *oneinchservice.QuoteResponse
	Amount      *big.Int
}

func Arbitrate(token0 string, token0Amount int64, token1 string, token1Amount int64) uint64 {
	token0AmountBig := util.ConvertToCryptoValue(token0Amount)
	token1AmountBig := util.ConvertToCryptoValue(token1Amount)
	reservesToken0ToToken1 := make(chan *bscconnector.Reserve, 2)
	reservesToken1ToToken0 := make(chan *bscconnector.Reserve, 2)
	callContractChan := make(chan *ContractCallEvent, 2)

	routerPoolMap := util.Pairs[token0][token1]
	routersCount := len(routerPoolMap)

	for router, poolAddress := range routerPoolMap {
		go bscconnector.Reserves(poolAddress, router, reservesToken0ToToken1, reservesToken1ToToken0)
	}

	// Call check for token0 to token
	go checkArbitragePossibility(token0, token1, token0AmountBig, reservesToken0ToToken1, 0, callContractChan, routersCount)

	// Call check for token1 to token0
	go checkArbitragePossibility(token1, token0, token1AmountBig, reservesToken1ToToken0, 1, callContractChan, routersCount)

	return CallContract(callContractChan)
}

func checkArbitragePossibility(tokenFrom string, tokenTo string, amount *big.Int, reservesChan chan *bscconnector.Reserve, fromTokenIndex uint8, callContractChan chan *ContractCallEvent, routerCount int) {
	quote, err := oneinchservice.Quote(tokenFrom, tokenTo, amount)
	if err != nil {
		fmt.Printf("1Inch Call Failed - Timed out - %d\n", fromTokenIndex)
		callContractChan <- nil
		return
	}

	toTokenAmount, valid := new(big.Int).SetString(quote.ToTokenAmount, 10)
	if !valid {
		fmt.Println("Fail to convert 1Inch ToTokenAmount")
		callContractChan <- nil
		return
	}

	minPayableAmount, reserveInfo := MinAmountIn(amount, fromTokenIndex, reservesChan, routerCount)

	profit := new(big.Int).Sub(toTokenAmount, minPayableAmount)
	hasProfit := profit.Sign() > 0

	if hasProfit {
		var contractCall ContractCallEvent
		contractCall.Amount = amount
		contractCall.Quote = quote
		contractCall.PairAddress = reserveInfo.PairAddress

		callContractChan <- &contractCall

		fmt.Printf("payableAmount = %s / toTokenAmount = %s / profit = %s \n",
			minPayableAmount, toTokenAmount, profit)
		// fmt.Printf("QUOTE: %s\n\n", quote)
	} else {
		callContractChan <- nil
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

func AmountIn(amountOut *big.Int, fee int64, reserveIn *big.Int, reserveOut *big.Int, fromTokenIndex uint8) *big.Int {
	feeBig := big.NewInt(10000 - (fee + util.Spread))

	numerator := new(big.Int)
	denominator := new(big.Int)
	if fromTokenIndex == 0 {
		numerator.Mul(reserveIn, amountOut)
		numerator.Mul(numerator, big.NewInt(10000))

		denominator.Sub(reserveOut, amountOut)
		denominator.Mul(denominator, feeBig)
	} else {
		numerator.Mul(reserveOut, amountOut)
		numerator.Mul(numerator, big.NewInt(10000))

		denominator.Sub(reserveIn, amountOut)
		denominator.Mul(denominator, feeBig)
	}
	payableAmount := new(big.Int).Div(numerator, denominator)
	payableAmount.Add(payableAmount, big.NewInt(1))

	return payableAmount
}

func MinAmountIn(amountOut *big.Int, fromTokenIndex uint8, reservesChan chan *bscconnector.Reserve, routerCount int) (*big.Int, *bscconnector.Reserve) {
	minAmount := big.NewInt(-1)
	var minReserve *bscconnector.Reserve

	var reserve *bscconnector.Reserve
	for i := 0; i < int(routerCount); i++ {
		select {
		case chanValue := <-reservesChan:
			reserve = chanValue
		case <-time.After(700 * time.Millisecond):
			fmt.Println("GetReserve Timed out")
			continue
		}

		fee := util.RouterFeeMap[reserve.PairRouter]
		currentValue := AmountIn(amountOut, fee, reserve.Reserve1, reserve.Reserve0, fromTokenIndex)
		if minAmount.Sign() == -1 || currentValue.Cmp(minAmount) < 0 {
			minAmount = currentValue
			minReserve = reserve
		}
	}

	return minAmount, minReserve
}

func CallContract(callContractChan chan *ContractCallEvent) uint64 {
	curTime := time.Now()
	var event *ContractCallEvent = nil
	select {
	case e := <-callContractChan:
		event = e
	case <-time.After(2 * time.Second):
		fmt.Println("Waiting arbitrage response [0]")
	}

	if event == nil {
		select {
		case e := <-callContractChan:
			event = e
		case <-time.After((2010 - time.Duration((time.Since(curTime)).Milliseconds())) * time.Millisecond):
			fmt.Println("Waiting arbitrage response [1]")
			return 0
		}
	}

	if event == nil {
		return 0
	}

	routes, path, err := routersAndPath(event.Quote)
	if err != nil {
		fmt.Printf("DEU PAU NAS ROTAS %s\n", err)
		return 0
	}

	if os.Getenv("RUN") == "true" {
		bscconnector.StartArbitrage(event.PairAddress, event.Amount, routes, *path, util.ContractAddress)
	}

	currentBlock := bscconnector.CurrentBlock()
	fmt.Printf("CHAMOUUUU StartArbitrage(%s, %s, %s, %s) @ %d\n", event.PairAddress, event.Amount, routes, path, currentBlock)

	return currentBlock
}
