package bscconnector

import (
	arbitas "arb-finder/src/bscconnector/arbitas_contract"
	pancakepaircontract "arb-finder/src/bscconnector/pancake_pair_contract"
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

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
	Reserve0           *big.Int
	Reserve1           *big.Int
	RationFrom0        *big.Float
	RationFrom1        *big.Float
	BlockTimestampLast uint32
}

func Reserves(poolId string, reservesChan chan *Reserve) {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		log.Fatal(err)
	}

	poolAddress := common.HexToAddress(poolId)
	instance, err := pancakepaircontract.NewPancakepaircontract(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	result, err := instance.GetReserves(nil)
	if err != nil {
		log.Fatal(err)
	}

	reserve0Float := new(big.Float).SetInt(result.Reserve0)
	reserve1Float := new(big.Float).SetInt(result.Reserve1)

	var reserves Reserve
	reserves.Reserve0 = result.Reserve0
	reserves.Reserve1 = result.Reserve1
	reserves.BlockTimestampLast = result.BlockTimestampLast
	reserves.RationFrom0 = new(big.Float).Quo(reserve0Float, reserve1Float)
	reserves.RationFrom1 = new(big.Float).Quo(reserve1Float, reserve0Float)

	reservesChan <- &reserves
	reservesChan <- &reserves
}

func StartArbitrage(amount *big.Int, routes *[]*big.Int, path []common.Address, contractAddress string) {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		log.Fatal(err)
	}

	poolAddress := common.HexToAddress(contractAddress)
	instance, err := arbitas.NewArbitas(poolAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	privateKey, err := crypto.HexToECDSA(PrivateKey())
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(56))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(int64(0))
	auth.GasPrice = big.NewInt(int64(6000000000))

	result, err := instance.StartArbitrage(auth, amount, *routes, path)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("FALHOU O CONTRATO")
	} else {
		fmt.Println("BOMBOU O CONTRATO")
		fmt.Println(result.Data())
	}
}

func CurrentBlock() uint64 {
	client, err := ethclient.Dial(network_url)
	if err != nil {
		log.Fatal(err)
	}

	result, err := client.BlockNumber(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println("FALHOU O GET BLOCK")
	}
	return result
}

func SubscribeNewBlock(channel chan *types.Header) ethereum.Subscription {
	client, err := ethclient.Dial(network_websocket)
	if err != nil {
		log.Fatal(err)
	}

	sub, err := client.SubscribeNewHead(context.Background(), channel)
	if err != nil {
		log.Fatal(err)
	}

	return sub
}
