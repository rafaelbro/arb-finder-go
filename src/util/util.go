package util

import "math/big"

var Tokens map[string]string = map[string]string{
	"CAKE": "0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82",
	"DAI":  "0x1af3f329e8be154074d8769d1ffa4ee058b1dbc3",
	"ETH":  "0x2170ed0880ac9a755fd29b2688956bd959f933f8",
	"ADA":  "0x3ee2200efb3400fabb9aacf31297cbdd1d435d47",
	"VAI":  "0x4bd17003473389a42daf6a0a729f6fdb328bbbd7",
	"USDT": "0x55d398326f99059ff775485246999027b3197955",
	"BTCB": "0x7130d2a12b9bcbfae4f2634d864a1ee1ce3ead9c",
	"USDC": "0x8ac76a51cc950d9822d68b83fe1ad97b32cd580d",
	"WBNB": "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c",
	"BUSD": "0xe9e7cea3dedca5984780bafc599bd69add087d56",
}

var ExchangesMap map[string]int64 = map[string]int64{
	"PANCAKESWAP":    0,
	"PANCAKESWAP_V2": 1,
	"APESWAP":        2,
	"WAULTSWAP":      3,
	"BAKERYSWAP":     4,
	"MDEX":           5,
	"ACRYPTOS_CORE":  6,
	"ACRYPTOS_META":  7,
	"ELLIPSIS_CORE":  8,
	"ELLIPSIS_META":  9,
}

var ContractAddress = "0xFA714b0381fDDA77229f0C801AdF27ff2dE6F0fb"

func ConvertToCryptoValue(value int64) *big.Int {
	landingZeros, exponent := big.NewInt(10), big.NewInt(18)
	landingZeros.Exp(landingZeros, exponent, nil)
	valueBig := big.NewInt(value)
	valueBig.Mul(valueBig, landingZeros)
	return valueBig
}

var Pairs map[string]map[string]map[string]string = map[string]map[string]map[string]string{
	Tokens["CAKE"]: {
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x0eD7e52944161450477ee417DE9Cd3a859b14fD0",
			"WAULTSWAP":      "0x7b12531Eb75F06A8C9cA4A5f27dbB952FD2A7430",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x804678fa97d91B974ec2af3c843270886528a9E6",
		},
	},
	Tokens["DAI"]: {
		Tokens["USDT"]: {
			"PANCAKESWAP_V2": "0xf6f5CE9a91Dd4FAe2d2eD92E25F2A4dc8564F174",
		},
		Tokens["USDC"]: {
			"PANCAKESWAP_V2": "0xadBba1EF326A33FDB754f14e62A96D5278b942Bd",
		},
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0xc7c3cCCE4FA25700fD5574DA7E200ae28BBd36A3",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x66FDB2eCCfB58cF098eaa419e5EfDe841368e489",
			"WAULTSWAP":      "0xcED829cB73d21B34a0AD4687C3Cd7D398172DBD8",
		},
	},
	Tokens["ETH"]: {
		Tokens["USDT"]: {
			"PANCAKESWAP_V2": "0x531FEbfeb9a61D948c384ACFBe6dCc51057AEa7e",
		},
		Tokens["BTCB"]: {
			"PANCAKESWAP_V2": "0xD171B26E4484402de70e3Ea256bE5A2630d7e88D",
			"WAULTSWAP":      "0xBb43C776D9dDDaD1395e1543545d05E138ccb4BA",
		},
		Tokens["USDC"]: {
			"PANCAKESWAP_V2": "0xEa26B78255Df2bBC31C1eBf60010D78670185bD0",
		},
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x74E4716E431f45807DCF19f284c7aA99F18a4fbc",
			"WAULTSWAP":      "0x04253aB3ff54D2E03b717BF6810a0a2Fd228365a",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x7213a321F1855CF1779f42c0CD85d3D95291D34C",
			"WAULTSWAP":      "0x40a2739d8B2CDDd5EDB8B563BA8e4c3326e23716",
		},
	},
	Tokens["VAI"]: {
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x133ee93FE93320e1182923E1a640912eDE17C90C",
		},
		Tokens["WBNB"]: {
			"WAULTSWAP": "0xB1b116D2814E92f3fAE565808695F93d0C2a2264",
		},
	},
	Tokens["USDT"]: {
		Tokens["BTCB"]: {
			"WAULTSWAP": "0x3F2e3461fd0E0eaA0fCc9Ec7A40C8B19b27da0b6",
		},
		Tokens["USDC"]: {
			"PANCAKESWAP_V2": "0xEc6557348085Aa57C72514D67070dC863C0a5A8c",
		},
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE",
			"WAULTSWAP":      "0xd6196036cB72BB921E013189CC594feC29453C2E",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x7EFaEf62fDdCCa950418312c6C91Aef321375A00",
			"WAULTSWAP":      "0x9Ce20a5169A3CD64A98C2C200aA995A2d8c8830e",
		},
	},
	Tokens["BTCB"]: {
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x61EB789d75A95CAa3fF50ed7E47b96c132fEc082",
			"WAULTSWAP":      "0xfCc62FB56c8E0630001B6EcC9eD38518D39499B2",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0xF45cd219aEF8618A92BAa7aD848364a158a24F33",
			"WAULTSWAP":      "0x61Ad21f79D1Bf96206Ad28d97B15D98a55944a2a",
		},
	},
	Tokens["USDC"]: {
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b",
			"WAULTSWAP":      "0xA34337690711CE3F265f56Ebd545Dda00d7C0405",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1",
		},
	},
	Tokens["WBNB"]: {
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16",
			"WAULTSWAP":      "0x4bbed8D9A1B27A4DDd84a3368A850e78c9580404",
		},
	},
}

var TradeQuantity map[string]int64 = map[string]int64{
	Tokens["CAKE"]: 250,
	Tokens["DAI"]:  5000,
	Tokens["ETH"]:  2,
	Tokens["ADA"]:  0,
	Tokens["VAI"]:  5000,
	Tokens["USDT"]: 5000,
	Tokens["BTCB"]: 1,
	Tokens["USDC"]: 5000,
	Tokens["WBNB"]: 12,
	Tokens["BUSD"]: 5000,
}

var RouterFeeMap map[string]int64 = map[string]int64{
	"PANCAKESWAP_V2": 25,
	"WAULTSWAP":      20,
}

var Spread int64 = 0
