export interface QuoteParams {
  fromTokenAddress: string;
  toTokenAddress: string;
  amount: string;
  fee?: number;
  protocols?: string;
  gasPrice?: string;
  complexityLevel?: string
  connectorTokens?: string;
  gasLimit?: number;
  parts?: number;
  mainRouteParts?: number;
}
