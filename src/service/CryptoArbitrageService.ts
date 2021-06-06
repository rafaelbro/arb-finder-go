import bigDecimal from "js-big-decimal";
import OneInchBinanceChainService from "../clients/1InchService/OneInchBinanceChainService";
import { QuoteResponse } from "../clients/1InchService/types/QuoteResponse";
import { PAIRS } from "../pairs";
import Token from "../tokens";
import { ArbitrageHumanReadableValues, ArbitrageResult, ArbitrageValues } from "../types/ArbitrageResult";

import { abi } from "../contracts/abis/test"
import Web3Connector from "../connectors/Web3/Web3Connector";
class CryptoArbitrageService {
  public async canArbitrate(
    amount: number,
    token0: Token,
    token1: Token
  ) {
    const bigIntAmount = BigInt(amount) * (10n ** 18n);
    const quote = await OneInchBinanceChainService.quote({
      fromTokenAddress: token0,
      toTokenAddress: token1,
      amount: bigIntAmount.toString()
    });

    const poolPair = PAIRS[token0][token1].pairPoolAddress;
    const poolRatio = await this.getPoolRatio(poolPair);
    const quoteRatio = await this.getQuoteRatio(quote);

    const profitPerUnit = bigDecimal.subtract(bigDecimal.multiply(quoteRatio.netRatio, 0.997), poolRatio);

    return {
      fromToken: token0,
      toToken: token1,
      hasProfit: bigDecimal.compareTo(profitPerUnit, 0) >= 0,
      poolRatio: poolRatio,
      quoteRatio: quoteRatio.netRatio,
      profitPerUnit: profitPerUnit,
      protocols: quote.protocols[0]
    } as ArbitrageResult;
  }

  public async getQuoteRatio(quote: QuoteResponse): Promise<ArbitrageValues> {
    const toToken = quote.toToken.address;

    const grossRatio = bigDecimal.divide(quote.fromTokenAmount, quote.toTokenAmount, 5)

    let estimatedGas = quote.estimatedGas;

    const feeInToToken = await this.convertGasToToken(estimatedGas, toToken);
    const netRatio = bigDecimal.divide(bigDecimal.subtract(quote.fromTokenAmount, feeInToToken),  quote.toTokenAmount, 5)

    return {
      grosRatio: grossRatio,
      fee: feeInToToken.toString(),
      netRatio: netRatio,
    }
  }

  public async getPoolRatio(poolAddress: string){
    const result = await Web3Connector.getReserves(poolAddress);
    const poolReserveRatio = bigDecimal.divide(result.token0Reserve, result.token1Reserve, 5);
    return poolReserveRatio;
  }

  public async convertGasToToken(gasAmount: number, tokenAddress: string) {
    let amount = BigInt(Math.ceil(gasAmount * (10 ** 9)));

    if(tokenAddress !== Token.BNB){
      const result = await OneInchBinanceChainService.quote({
        fromTokenAddress: Token.BNB,
        toTokenAddress: tokenAddress,
        amount: BigInt(Math.ceil(gasAmount * (10 ** 9))).toString()
      });
      amount = BigInt(result.toTokenAmount);
    }

    return amount;
  }
}

export default new CryptoArbitrageService();
