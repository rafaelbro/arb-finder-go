package oneinchservice

import (
	"encoding/json"
	"fmt"
	"math/big"
	"net/http"
	"strings"
	"time"
)

const (
	chain_code = "56"
	base_url   = "https://api.1inch.exchange/v3.0"
)

type OneInchToken struct {
	Symbol   string
	Name     string
	Address  string
	Decimals int32
	LogoURI  string
}

type OneInchProtocol struct {
	Name             string
	Part             int32
	FromTokenAddress string
	ToTokenAddress   string
}

type QuoteResponse struct {
	FromToken       OneInchToken
	ToToken         OneInchToken
	ToTokenAmount   string
	FromTokenAmount string
	Protocols       [][][]OneInchProtocol
	EstimatedGas    int64
}

func Quote(fromTokenAddress string, toTokenAddress string, amount *big.Int) (*QuoteResponse, error) {
	protocols := []string{"PANCAKESWAP", "PANCAKESWAP_V2", "APESWAP", "WAULTSWAP", "BAKERYSWAP", "MDEX", "ACRYPTOS", "ELLIPSIS_FINANCE"}
	url := fmt.Sprintf("%s/%s/quote?fromTokenAddress=%s&toTokenAddress=%s&amount=%s&parts=1&protocols=%s",
		base_url, chain_code, fromTokenAddress, toTokenAddress, amount.String(), strings.Join(protocols, ","))

	client := http.Client{
		Timeout: 5 * time.Second,
	}
	res, err := client.Get(url)
	if err != nil {
		return nil, err
	}

	if res.StatusCode == 200 {
		decoder := json.NewDecoder(res.Body)
		var quoteData QuoteResponse
		err = decoder.Decode(&quoteData)
		if err != nil {
			return nil, err
		}
		return &quoteData, nil
	}
	return nil, err
}
