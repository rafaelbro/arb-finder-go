package bscconnector

import (
	arbitas "arb-finder/src/bscconnector/arbitas_contract"
	pancakepaircontract "arb-finder/src/bscconnector/pancake_pair_contract"
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

const (
	network_url       = "https://bsc-dataseed1.binance.org:443"
	network_websocket = "wss://bsc-ws-node.nariox.org:443"
)

type Reserve struct {
	PairRouter         string
	PairAddress        string
	Reserve0           *big.Int
	Reserve1           *big.Int
	BlockTimestampLast uint32
}

func Reserves(poolId string, router string, reservesToken0Chan chan *Reserve, reservesToken1Chan chan *Reserve) {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		fmt.Println(err)
	}

	poolAddress := common.HexToAddress(poolId)
	instance, err := pancakepaircontract.NewPancakepaircontract(poolAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	result, err := instance.GetReserves(nil)
	if err != nil {
		fmt.Println(err)
	}

	var reserves Reserve
	reserves.PairRouter = router
	reserves.PairAddress = poolId
	reserves.Reserve0 = result.Reserve0
	reserves.Reserve1 = result.Reserve1
	reserves.BlockTimestampLast = result.BlockTimestampLast

	reservesToken0Chan <- &reserves
	reservesToken1Chan <- &reserves
}

func StartArbitrage(poolPair string, amount *big.Int, routes *[]*big.Int, path []common.Address, contractAddress string) {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		fmt.Println(err)
	}

	poolAddress := common.HexToAddress(contractAddress)
	instance, err := arbitas.NewArbitas(poolAddress, client)
	if err != nil {
		fmt.Println(err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey())
	if err != nil {
		fmt.Println(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		fmt.Println("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		fmt.Println(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	if err != nil {
		fmt.Println(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0))
	auth.GasLimit = 1000000
	auth.GasPrice = big.NewInt(int64(7000000000))
	poolPairAddress := common.HexToAddress(poolPair)

	result, err := instance.StartArbitrage(auth, poolPairAddress, amount, *routes, path)

	fmt.Println(time.Now().UTC().String())
	if err != nil {
		fmt.Println("FALHOU O CONTRATO")
		fmt.Println(result)
		fmt.Println(err)
	} else {
		fmt.Println("BOMBOU O CONTRATO")
		fmt.Println(result.Hash())
		fmt.Println(result)
	}
}

func CurrentBlock() uint64 {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		fmt.Println(err)
	}

	result, err := client.BlockNumber(context.Background())
	if err != nil {
		fmt.Println(err)
	}
	if err != nil {
		fmt.Println("FALHOU O GET BLOCK")
	}
	return result
}

func SubscribeNewBlock(channel chan *types.Header) ethereum.Subscription {
	client, err := ethclient.Dial(network_websocket)
	if err != nil {
		fmt.Println(err)
	}

	sub, err := client.SubscribeNewHead(context.Background(), channel)
	if err != nil {
		fmt.Println(err)
	}

	return sub
}
