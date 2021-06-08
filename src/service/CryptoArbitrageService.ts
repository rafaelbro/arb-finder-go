import bigDecimal from "js-big-decimal";
import OneInchBinanceChainService from "../clients/1InchService/OneInchBinanceChainService";
import { QuoteResponse } from "../clients/1InchService/types/QuoteResponse";
import { PAIRS } from "../pairs";
import Token from "../tokens";
import { ArbitrageResult, ArbitrageValues } from "../types/ArbitrageResult";
import Web3Connector from "../connectors/Web3/Web3Connector";
import Exchanges from "../exchanges";
import { ContractReserves } from "../connectors/Web3/Web3ConnectorTypes";
class CryptoArbitrageService {
  public async canArbitrate(
    token0: Token,
    amountToken0: number,
    token1: Token,
    amountToken1: number,
  ) {

    const bigIntAmountToken0= BigInt(amountToken0) * (10n ** 18n);
    const bigIntAmountToken1= BigInt(amountToken1) * (10n ** 18n);

    const quoteFrom0 = await OneInchBinanceChainService.quote({
      fromTokenAddress: token0,
      toTokenAddress: token1,
      amount: bigIntAmountToken0.toString(),
      protocols: Object.keys(Exchanges).join(',')
    });

    const quoteFrom1 = await OneInchBinanceChainService.quote({
      fromTokenAddress: token1,
      toTokenAddress: token0,
      amount: bigIntAmountToken1.toString(),
      protocols: Object.keys(Exchanges).join(',')
    });

    const poolPair = PAIRS[token0][token1].pairPoolAddress;
    const poolReserves = await Web3Connector.getReserves(poolPair);
    const [poolRatioFrom0, poolRatioFrom1] = await this.getPoolRatio(poolReserves);

    const quoteFrom0Ratio = await this.getQuoteRatio(quoteFrom0);
    const grossLoanPayFrom0 = bigDecimal.add(bigDecimal.divide(amountToken0, poolRatioFrom0, 5), bigDecimal.multiply(amountToken0, 0.004));
    const tradeAmountFrom0 = bigDecimal.divide(amountToken0, quoteFrom0Ratio.netRatio, 5);
    const profitFrom0 = bigDecimal.subtract(tradeAmountFrom0, grossLoanPayFrom0);


    const quoteFrom1Ratio = await this.getQuoteRatio(quoteFrom1);
    const grossLoanPayFrom1 = bigDecimal.add(bigDecimal.divide(amountToken1, poolRatioFrom1, 5), bigDecimal.multiply(amountToken1, 0.004));
    const tradeAmountFrom1 = bigDecimal.divide(amountToken1, quoteFrom1Ratio.netRatio, 5);
    const profitFrom1 = bigDecimal.subtract(tradeAmountFrom1, grossLoanPayFrom1);

    return {
      timestamp: new Date().toUTCString(),
      fromToken0: {
        fromToken: Object.keys(Token).find((x: any) => Token[x as keyof typeof Token] == token0),
        toToken: Object.keys(Token).find((x: any) => Token[x as keyof typeof Token] == token1),
        poolRatio: poolRatioFrom0,
        quoteRatio: quoteFrom0Ratio.netRatio,
        loanAmount: grossLoanPayFrom0,
        tradeAmount: tradeAmountFrom0,
        profit: profitFrom0,
        protocols: quoteFrom0.protocols,
        hasProfit: bigDecimal.compareTo(profitFrom0, 0) >= 0,
        hasLiquidity: this.hasLiquidity(bigIntAmountToken0, BigInt(poolReserves.token0Reserve)),
        runContract: bigDecimal.compareTo(profitFrom0, 0) >= 0 && this.hasLiquidity(bigIntAmountToken0, BigInt(poolReserves.token0Reserve)),
      },
      fromToken1: {
        fromToken: Object.keys(Token).find((x: any) => Token[x as keyof typeof Token] == token1),
        toToken: Object.keys(Token).find((x: any) => Token[x as keyof typeof Token] == token0),
        poolRatio: poolRatioFrom1,
        quoteRatio: quoteFrom1Ratio.netRatio,
        loanAmount: grossLoanPayFrom1,
        tradeAmount: tradeAmountFrom1,
        profit: profitFrom1,
        protocols: quoteFrom1.protocols,
        hasProfit: bigDecimal.compareTo(profitFrom1, 0) >= 0,
        hasLiquidity: this.hasLiquidity(bigIntAmountToken1, BigInt(poolReserves.token1Reserve)),
        runContract: bigDecimal.compareTo(profitFrom1, 0) >= 0 && this.hasLiquidity(bigIntAmountToken1, BigInt(poolReserves.token1Reserve)),
      },
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

  public async getPoolRatio(poolReserves: ContractReserves){
    const poolReserveRatioFrom0 = bigDecimal.divide(poolReserves.token0Reserve, poolReserves.token1Reserve, 5);
    const poolReserveRatioFrom1 = bigDecimal.divide(poolReserves.token1Reserve, poolReserves.token0Reserve, 5);
    return [poolReserveRatioFrom0, poolReserveRatioFrom1];
  }

  public async convertGasToToken(gasAmount: number, tokenAddress: string) {
    let amount = BigInt(Math.ceil(gasAmount * (10 ** 9)));

    if(tokenAddress !== Token.WBNB){
      const result = await OneInchBinanceChainService.quote({
        fromTokenAddress: Token.BNB,
        toTokenAddress: tokenAddress,
        amount: BigInt(Math.ceil(gasAmount * (10 ** 9))).toString()
      });
      amount = BigInt(result.toTokenAmount);
    }

    return amount;
  }

  public hasLiquidity(tokenAmount: bigint, reserve: bigint): boolean{
    return bigDecimal.compareTo(bigDecimal.divide(tokenAmount.toString(), reserve.toString(), 5), '0.01') === -1;
  }
}

export default new CryptoArbitrageService();
