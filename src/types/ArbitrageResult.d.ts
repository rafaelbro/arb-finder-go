import { OneInchProtocol } from "../clients/1InchService/types/QuoteResponse";
import Token from "../tokens";

export interface ArbitrageResult{
  timestamp: string,
  fromToken0: {
    runContract: boolean,
    hasProfit: boolean,
    poolRatio: string,
    fromToken: Token,
    toToken: Token,
    quoteRatio: string,
    loanAmount: string,
    tradeAmount: string,
    profit: string,
    liquidityRatio: string,
    hasLiquidity: boolean,
    protocols: OneInchProtocol[][][],
  },
  fromToken1?: {
    runContract: boolean,
    hasProfit: boolean
    poolRatio: string,
    fromToken: Token,
    toToken: Token,
    quoteRatio: string,
    loanAmount: string,
    tradeAmount: string,
    profit: string,
    liquidityRatio: string,
    hasLiquidity: boolean,
    protocols: OneInchProtocol[][][],
  }
}

export interface ArbitrageValues {
  grosRatio: string;
  fee: string;
  netRatio: string;
}

export interface ArbitrageHumanReadableValues {
  grossProfit: string;
  fee: string;
  netProfit: string;
}
