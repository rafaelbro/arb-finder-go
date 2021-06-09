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

    const poolPair = PAIRS[token0][token1].pairPoolAddress;
    const poolReserves = await Web3Connector.getReserves(poolPair);
    const [poolRatioFrom0, poolRatioFrom1] = await this.getPoolRatio(poolReserves);

    OneInchBinanceChainService.quote({
      fromTokenAddress: token0,
      toTokenAddress: token1,
      amount: bigIntAmountToken0.toString(),
      protocols: Object.keys(Exchanges).join(',')
    }).then(async (quoteFrom0) => {
      const quoteFrom0Ratio = await this.getQuoteRatio(quoteFrom0);
      const grossLoanPayFrom0 = bigDecimal.multiply(bigDecimal.divide(amountToken0, poolRatioFrom0, 7), 1.01);
      const tradeAmountFrom0 = bigDecimal.divide(amountToken0, quoteFrom0Ratio.netRatio, 7);
      const profitFrom0 = bigDecimal.subtract(tradeAmountFrom0, grossLoanPayFrom0);
      const hasProfitToken0 = bigDecimal.compareTo(profitFrom0, 0) >= 0
      const liquidityToken0 = this.hasLiquidity(bigIntAmountToken0, BigInt(poolReserves.token0Reserve))
      const hasLiquidityToken0 = bigDecimal.compareTo(liquidityToken0, '0.01') === -1;
      const runContractToken0 = hasLiquidityToken0 && hasProfitToken0;

      if(runContractToken0){
        const amount = bigIntAmountToken0;
        const routers = quoteFrom0.protocols[0].map((p: any) => Exchanges[p[0].name as keyof typeof Exchanges]);
        const addressList = quoteFrom0.protocols[0].map((p: any) => p[0].toTokenAddress);
        addressList.unshift(token0);

        await Web3Connector.startArbitrage(amount, routers, addressList);
      }
    })


    OneInchBinanceChainService.quote({
      fromTokenAddress: token1,
      toTokenAddress: token0,
      amount: bigIntAmountToken1.toString(),
      protocols: Object.keys(Exchanges).join(',')
    }).then(async (quoteFrom1) => {
      const quoteFrom1Ratio = await this.getQuoteRatio(quoteFrom1);
      const grossLoanPayFrom1 = bigDecimal.multiply(bigDecimal.divide(amountToken1, poolRatioFrom1, 7), 1.01);
      const tradeAmountFrom1 = bigDecimal.divide(amountToken1, quoteFrom1Ratio.netRatio, 7);
      const profitFrom1 = bigDecimal.subtract(tradeAmountFrom1, grossLoanPayFrom1);
      const hasProfitToken1 = bigDecimal.compareTo(profitFrom1, 0) >= 0
      const liquidityToken1 = this.hasLiquidity(bigIntAmountToken1, BigInt(poolReserves.token0Reserve))
      const hasLiquidityToken1 = bigDecimal.compareTo(liquidityToken1, '0.01') === -1;
      const runContractToken1 = hasLiquidityToken1 && hasProfitToken1;

      if(runContractToken1){
        const amount = bigIntAmountToken1;
        const routers = quoteFrom1.protocols[0].map((p: any) => Exchanges[p[0].name as keyof typeof Exchanges]);
        const addressList = quoteFrom1.protocols[0].map((p: any) => p[0].toTokenAddress);
        addressList.unshift(token0);

        await Web3Connector.startArbitrage(amount, routers, addressList);
      }
    });
  }

  public async getQuoteRatio(quote: QuoteResponse): Promise<ArbitrageValues> {
    const toToken = quote.toToken.address;

    const grossRatio = bigDecimal.divide(quote.fromTokenAmount, quote.toTokenAmount, 7)

    let estimatedGas = quote.estimatedGas;

    const feeInToToken = 0 ;//await this.convertGasToToken(estimatedGas, toToken);
    const netRatio = bigDecimal.divide(quote.fromTokenAmount, quote.toTokenAmount, 7)

    return {
      grosRatio: grossRatio,
      fee: feeInToToken.toString(),
      netRatio: netRatio,
    }
  }

  public async getPoolRatio(poolReserves: ContractReserves){
    const poolReserveRatioFrom0 = bigDecimal.divide(poolReserves.token0Reserve, poolReserves.token1Reserve, 7);
    const poolReserveRatioFrom1 = bigDecimal.divide(poolReserves.token1Reserve, poolReserves.token0Reserve, 7);
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

  public hasLiquidity(tokenAmount: bigint, reserve: bigint): string{
    return bigDecimal.divide(tokenAmount.toString(), reserve.toString(), 7);
  }
}

export default new CryptoArbitrageService();
