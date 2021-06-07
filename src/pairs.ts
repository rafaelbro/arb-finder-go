import Token from "./tokens"

export interface Pair {
  pairPoolAddress: string;
}

export const TOKEN_VALUE: Record<Token, number> = {
  [Token.CAKE]: 250,
  [Token.DAI]: 5000,
  [Token.ETH]: 0,
  [Token.ADA]: 0,
  [Token.VAI]: 5000,
  [Token.WBNB]: 12,
  [Token.BNB]: 12,
  [Token.USDT]: 5000,
  [Token.USDC]: 0,
  [Token.BUSD]: 5000,
}

export const PAIRS: any = {
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
