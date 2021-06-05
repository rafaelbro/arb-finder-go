import Token from "../tokens";

export interface ArbitrageResult extends ArbitrageValues{
  fromToken: Token,
  toToken: Token,
  humanReadable?: ArbitrageHumanReadableValues
}

export interface ArbitrageValues {
  hasProfit: boolean;
  grossProfit: string;
  fee: string;
  netProfit: string;
  protocols: string[][],
}

export interface ArbitrageHumanReadableValues {
  grossProfit: string;
  fee: string;
  netProfit: string;
}

10000.000000000000000000
