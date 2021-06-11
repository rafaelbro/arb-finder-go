import Token from "./tokens"

export interface Pair {
  pairPoolAddress: string;
}

16370061067567019379

export const TOKEN_VALUE: Record<Token, number> = {
  [Token.CAKE]: 150,
  [Token.DAI]: 5000,
  [Token.ETH]: 2,
  [Token.ADA]: 0,
  [Token.VAI]: 5000,
  [Token.USDT]: 5000,
  [Token.BTCB]: 1,
  [Token.USDC]: 5000,
  [Token.WBNB]: 12,
  [Token.BUSD]: 5000,
}

export const PAIRS: any = {
  [Token.CAKE]: {
    // [Token.DAI]: {
    //   pairPoolAddress: '0x29084c8D46cCe9B50629F01eBBbE11523A7A3EC2'
    // },
    // [Token.ETH]: {
    //   pairPoolAddress: '0x3f23B4F1a35794306ba4f3176934012dC73312D1',
    // },
    // [Token.VAI]: {
    //   pairPoolAddress: '0x090a10634b8a6850DB6d92cbe1aA747861b11552',
    // },
    // [Token.USDT]: {
    //   pairPoolAddress: '0xA39Af17CE4a8eb807E076805Da1e2B8EA7D0755b',
    // },
    // [Token.BTCB]: {
    //   pairPoolAddress: '0x3b2B70Dd9684deEE3f942653a418bE34fc0D525b',
    // },
    // [Token.USDC]: {
    //   pairPoolAddress: '0x177d955dCA80443A09E7a7F5946cA16b8E0dcA1C'
    // },
    [Token.WBNB]: {
      pairPoolAddress: '0x0eD7e52944161450477ee417DE9Cd3a859b14fD0',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x804678fa97d91B974ec2af3c843270886528a9E6',
    },
  },
  [Token.DAI]: {
    // [Token.ETH]: {
    //   pairPoolAddress: '0x9Ca4eBCf0a8ebB90069f06c0A814433201B99475',
    // },
    // [Token.VAI]: {
    //   pairPoolAddress: '0xfE38f05Fa1BA8f6B7A63f80b7808b5272941d47f',
    // },
    [Token.USDT]: {
      pairPoolAddress: '0xf6f5CE9a91Dd4FAe2d2eD92E25F2A4dc8564F174',
    },
    // [Token.BTCB]: {
    //   pairPoolAddress: '0x15e458Ab8B1370c3CB925093Ece0693aAae631fb',
    // },
    [Token.USDC]: {
      pairPoolAddress: '0xadBba1EF326A33FDB754f14e62A96D5278b942Bd'
    },
    [Token.WBNB]: {
      pairPoolAddress: '0xc7c3cCCE4FA25700fD5574DA7E200ae28BBd36A3',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x66FDB2eCCfB58cF098eaa419e5EfDe841368e489',
    },
  },
  [Token.ETH]: {
    // [Token.VAI]: {
    //   pairPoolAddress: '0x00A994e2EaA83340FC40Fe15e2Def6BebD04D4CA',
    // },
    [Token.USDT]: {
      pairPoolAddress: '0x531FEbfeb9a61D948c384ACFBe6dCc51057AEa7e',
    },
    [Token.BTCB]: {
      pairPoolAddress: '0xD171B26E4484402de70e3Ea256bE5A2630d7e88D',
    },
    [Token.USDC]: {
      pairPoolAddress: '0xEa26B78255Df2bBC31C1eBf60010D78670185bD0'
    },
    [Token.WBNB]: {
      pairPoolAddress: '0x74E4716E431f45807DCF19f284c7aA99F18a4fbc',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x7213a321F1855CF1779f42c0CD85d3D95291D34C',
    },
  },
  [Token.VAI]: {
    // [Token.USDT]: {
    //   pairPoolAddress: '0xD94FeFc80a7d10d4708b140c7210569061a7eddb',
    // },
    // [Token.BTCB]: {
    //   pairPoolAddress: '0x87da5D669F1F92E0E6f4BbB25630558Af7899561',
    // },
    // [Token.WBNB]: {
    //   pairPoolAddress: '0x3955d04E88cAa2482ab4815431e703E4d65Ec93C',
    // },
    [Token.BUSD]: {
      pairPoolAddress: '0x133ee93FE93320e1182923E1a640912eDE17C90C',
    },
  },
  [Token.USDT]: {
    // [Token.BTCB]: {
    //   pairPoolAddress: '0x3F803EC2b816Ea7F06EC76aA2B6f2532F9892d62',
    // },
    [Token.USDC]: {
      pairPoolAddress: '0xEc6557348085Aa57C72514D67070dC863C0a5A8c'
    },
    [Token.WBNB]: {
      pairPoolAddress: '0x16b9a82891338f9bA80E2D6970FddA79D1eb0daE',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x7EFaEf62fDdCCa950418312c6C91Aef321375A00',
    }
  },
  [Token.BTCB]: {
    // [Token.USDC]: {
    //   pairPoolAddress: '0x2dF244535624761f6fCc381CaE3e9b903429d9Ff'
    // },
    [Token.WBNB]: {
      pairPoolAddress: '0x61EB789d75A95CAa3fF50ed7E47b96c132fEc082',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0xF45cd219aEF8618A92BAa7aD848364a158a24F33',
    }
  },
  [Token.USDC]: {
    [Token.WBNB]: {
      pairPoolAddress: '0xd99c7F6C65857AC913a8f880A4cb84032AB2FC5b',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x2354ef4DF11afacb85a5C7f98B624072ECcddbB1',
    }
  },
  [Token.WBNB]: {
    [Token.BUSD]: {
      pairPoolAddress: '0x58F876857a02D6762E0101bb5C46A8c1ED44Dc16',
    },
  },
}

export const PAIRS_V1: any = {
  [Token.CAKE]: {
    [Token.WBNB]: {
      pairPoolAddress: '0xA527a61703D82139F8a06Bc30097cC9CAA2df5A6',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0x0Ed8E0A2D99643e1e65CCA22Ed4424090B8B7458',
    },
    [Token.VAI]: {
      pairPoolAddress: '0xF1cc1FB4aC01ee74186e5e999c7027371218B232',
    },
    [Token.USDT]: {
      pairPoolAddress: '0x3f3d4CE222A7C919EA7f0231471c77478E36Fc0d',
    },
  },
  [Token.VAI]: {
    [Token.WBNB]: {
      pairPoolAddress: '0xe62C5A3355068Cc383D89ab831E000473C043Cae',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0xfF17ff314925Dff772b71AbdFF2782bC913B3575',
    },
    [Token.USDT]: {
      pairPoolAddress: '0xFac8e890218150E8c010A42ee745A8aF2063DB80',
    }
  },
  [Token.USDT]: {
    [Token.WBNB]: {
      pairPoolAddress: '0x20bCC3b8a0091dDac2d0BC30F68E6CBb97de59Cd',
    },
    [Token.BUSD]: {
      pairPoolAddress: '0xc15fa3E22c912A276550F3E5FE3b0Deb87B55aCd',
    }
  },
  [Token.WBNB]: {
    [Token.BUSD]: {
      pairPoolAddress: '0x1B96B92314C44b159149f7E0303511fB2Fc4774f',
    },
  },
}
