import { OneInchProtocol } from "../clients/1InchService/types/QuoteResponse";
import Token from "../tokens";

export interface ArbitrageResult{
  timestamp: string,
  fromToken0: {
    hasProfit: boolean,
    poolRatio: string,
    fromToken: Token,
    toToken: Token,
    quoteRatio: string,
    loanAmount: string,
    tradeAmount: string,
    profit: string,
    protocols: string,
  },
  fromToken1: {
    hasProfit: boolean
    poolRatio: string,
    fromToken: Token,
    toToken: Token,
    quoteRatio: string,
    loanAmount: string,
    tradeAmount: string,
    profit: string,
    protocols: string,
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
