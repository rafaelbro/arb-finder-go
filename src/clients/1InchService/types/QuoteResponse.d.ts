export interface QuoteResponse {
  fromToken: OneInchToken;
  toToken: OneInchToken;
  toTokenAmount: string;
  fromTokenAmount: string;
  protocols: OneInchProtocol[][][];
  estimatedGas: number;
}

export interface OneInchToken {
  symbol: string;
  name: string;
  address: string;
  decimals: number;
  logoURI: string;
}

export interface OneInchProtocol {
  name: string;
  part: number;
  fromTokenAddress: string;
  toTokenAddress: string;
}
