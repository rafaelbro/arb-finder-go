import { OneInchProtocol } from "../clients/1InchService/types/QuoteResponse";
import Token from "../tokens";

export interface ArbitrageResult{
  fromToken: Token,
  toToken: Token,
  hasProfit: boolean,
  poolRatio: string,
  quoteRatio: string,
  profitPerUnit: string,
  protocols: OneInchProtocol[][],
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
