import Token from "./tokens"

export interface Pair {
  pairPoolAddress: string;
}

export const PAIRS: any = {
  [Token.CAKE]: {
    [Token.WBNB]: {
      pairPoolAddress: '0xA527a61703D82139F8a06Bc30097cC9CAA2df5A6',
    },
  },
  [Token.ADA]: {
    [Token.WBNB]: {
      pairPoolAddress: '0xA527a61703D82139F8a06Bc30097cC9CAA2df5A6',
    },
  }
}
