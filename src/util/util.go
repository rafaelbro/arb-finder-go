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

var ContractAddress = "0x5735f3d06c3F743BCE381a90C7fAcD1400e5A4Fa"

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
			"MDEX":           "0xA13aFe2DF0fA0bb11F2aeAAAF98aC1D591E108d1",
			"APESWAP":        "0x60593Abea55e9Ea9d31c1b6473191cD2475a720D",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x804678fa97d91B974ec2af3c843270886528a9E6",
		},
	},
	Tokens["DAI"]: {
		Tokens["USDT"]: {
			"PANCAKESWAP_V2": "0xf6f5CE9a91Dd4FAe2d2eD92E25F2A4dc8564F174",
			"MDEX":           "0x59B76b5D39370ba2Aa7e723c639861266e85BFEc",
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
			"APESWAP":        "0x8b6EcEA3e9bd6290c2150A89AF6c69887AaF1870",
		},
	},
	Tokens["ETH"]: {
		Tokens["USDT"]: {
			"PANCAKESWAP_V2": "0x531FEbfeb9a61D948c384ACFBe6dCc51057AEa7e",
			"MDEX":           "0x0FB881c078434b1C0E4d0B64d8c64d12078b7Ce2",
		},
		Tokens["BTCB"]: {
			"PANCAKESWAP_V2": "0xD171B26E4484402de70e3Ea256bE5A2630d7e88D",
			"WAULTSWAP":      "0xBb43C776D9dDDaD1395e1543545d05E138ccb4BA",
			"MDEX":           "0x577d005912C49B1679B4c21E334FdB650E92C077",
		},
		Tokens["USDC"]: {
			"PANCAKESWAP_V2": "0xEa26B78255Df2bBC31C1eBf60010D78670185bD0",
		},
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x74E4716E431f45807DCF19f284c7aA99F18a4fbc",
			"WAULTSWAP":      "0x04253aB3ff54D2E03b717BF6810a0a2Fd228365a",
			"MDEX":           "0x82E8F9e7624fA038DfF4a39960F5197A43fa76aa",
			"APESWAP":        "0xA0C3Ef24414ED9C9B456740128d8E63D016A9e11",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x7213a321F1855CF1779f42c0CD85d3D95291D34C",
			"WAULTSWAP":      "0x40a2739d8B2CDDd5EDB8B563BA8e4c3326e23716",
			"MDEX":           "0xc0BA2569e473974e9004CEEEae76Aeaea521525c",
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
			"MDEX":      "0xda28Eb7ABa389C1Ea226A420bCE04Cb565Aafb85",
		},
		Tokens["USDC"]: {
			"PANCAKESWAP_V2": "0xEc6557348085Aa57C72514D67070dC863C0a5A8c",
			"MDEX":           "0x9f4Da89774570E27170873BefD139a79CB1A3da2",
		},
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE",
			"WAULTSWAP":      "0xd6196036cB72BB921E013189CC594feC29453C2E",
			"MDEX":           "0x09CB618bf5eF305FadfD2C8fc0C26EeCf8c6D5fd",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x7EFaEf62fDdCCa950418312c6C91Aef321375A00",
			"WAULTSWAP":      "0x9Ce20a5169A3CD64A98C2C200aA995A2d8c8830e",
			"MDEX":           "0x62c1dEC1fF328DCdC157Ae0068Bb21aF3967aCd9",
			"APESWAP":        "0x2e707261d086687470B515B320478Eb1C88D49bb",
		},
	},
	Tokens["BTCB"]: {
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0x61EB789d75A95CAa3fF50ed7E47b96c132fEc082",
			"WAULTSWAP":      "0xfCc62FB56c8E0630001B6EcC9eD38518D39499B2",
			"MDEX":           "0x969f2556F786a576F32AeF6c1D6618f0221Ec70e",
			"APESWAP":        "0x1E1aFE9D9c5f290d8F6996dDB190bd111908A43D",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0xF45cd219aEF8618A92BAa7aD848364a158a24F33",
			"WAULTSWAP":      "0x61Ad21f79D1Bf96206Ad28d97B15D98a55944a2a",
			"MDEX":           "0x4fb8253432FB3e92109c91E3Ff2b85FfA0f6A1F4",
		},
	},
	Tokens["USDC"]: {
		Tokens["WBNB"]: {
			"PANCAKESWAP_V2": "0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b",
			"WAULTSWAP":      "0xA34337690711CE3F265f56Ebd545Dda00d7C0405",
		},
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1",
			"APESWAP":        "0xC087C78AbaC4A0E900a327444193dBF9BA69058E",
		},
	},
	Tokens["WBNB"]: {
		Tokens["BUSD"]: {
			"PANCAKESWAP_V2": "0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16",
			"WAULTSWAP":      "0x4bbed8D9A1B27A4DDd84a3368A850e78c9580404",
			"MDEX":           "0x340192D37d95fB609874B1db6145ED26d1e47744",
			"APESWAP":        "0x51e6D27FA57373d8d4C256231241053a70Cb1d93",
		},
	},
}

var TradeQuantity map[string]int64 = map[string]int64{
	Tokens["CAKE"]: 500,
	Tokens["DAI"]:  10000,
	Tokens["ETH"]:  4,
	Tokens["ADA"]:  0,
	Tokens["VAI"]:  10000,
	Tokens["USDT"]: 10000,
	Tokens["BTCB"]: 1,
	Tokens["USDC"]: 10000,
	Tokens["WBNB"]: 25,
	Tokens["BUSD"]: 10000,
}

var RouterFeeMap map[string]int64 = map[string]int64{
	"PANCAKESWAP_V2": 25,
	"WAULTSWAP":      20,
	"MDEX":           30,
	"APESWAP":        25,
}

var Spread int64 = 3
