import bigDecimal from "js-big-decimal";
import OneInchBinanceChainService from "../clients/1InchService/OneInchBinanceChainService";
import { OneInchProtocol, QuoteResponse } from "../clients/1InchService/types/QuoteResponse";
import { PAIRS } from "../pairs";
import Token from "../tokens";
import { ArbitrageResult, ArbitrageValues } from "../types/ArbitrageResult";
import Web3Connector from "../connectors/Web3/Web3Connector";
import { Exchanges, ExchangesMap } from "../exchanges";
import { ContractReserves } from "../connectors/Web3/Web3ConnectorTypes";
class CryptoArbitrageService {
  private allowedTokensACryptos = new Set([Token.BUSD.toString(),Token.USDT.toString(),Token.DAI.toString(),Token.USDC.toString(),Token.VAI.toString()]);
  private allowedTokensEllipsis = new Set([Token.BUSD.toString(),Token.USDT.toString(),Token.DAI.toString(),Token.USDC.toString()]);
  private FEE = 1.005;

  public async canArbitrate(
    token0: Token,
    amountToken0: number,
    token1: Token,
    amountToken1: number,
  ) {

    const getKey = (value: string) => {
      return Object.keys(Token).find(x => Token[x as keyof typeof Token] === value);
    }

    const bigIntAmountToken0= BigInt(amountToken0) * (10n ** 18n);
    const bigIntAmountToken1= BigInt(amountToken1) * (10n ** 18n);

    const poolPair = PAIRS[token0][token1].pairPoolAddress;
    const poolReserves = await Web3Connector.getReserves(poolPair);
    const [poolRatioFrom0, poolRatioFrom1] = await this.getPoolRatio(poolReserves);

    Web3Connector.getBlockNumber().then((block) => console.log(`[BLOCK LOG] from: ${token0} - to: ${token1} - block: ${block}`));
    OneInchBinanceChainService.quote({
      fromTokenAddress: token0,
      toTokenAddress: token1,
      amount: bigIntAmountToken0.toString(),
      protocols: Exchanges.join(',')
    }).then(async (quoteFrom0) => {
      const quoteFrom0Ratio = await this.getQuoteRatio(quoteFrom0);
      const loanPayFrom0 = bigDecimal.divide(amountToken0, poolRatioFrom0, 7);
      const grossLoanPayFrom0 = bigDecimal.multiply(loanPayFrom0, this.FEE);
      const tradeAmountFrom0 = bigDecimal.divide(amountToken0, quoteFrom0Ratio.netRatio, 7);
      const profitFrom0 = bigDecimal.subtract(tradeAmountFrom0, grossLoanPayFrom0);
      const hasProfitToken0 = bigDecimal.compareTo(profitFrom0, 0) >= 0
      const liquidityToken0 = this.hasLiquidity(bigIntAmountToken0, BigInt(poolReserves.token0Reserve))
      const hasLiquidityToken0 = bigDecimal.compareTo(liquidityToken0, '0.01') === -1;
      const runContractToken0 = hasLiquidityToken0 && hasProfitToken0;


      if(runContractToken0){
        const amount = bigIntAmountToken0;
        const [routers, addressList] = this.getRoutersAndPath(quoteFrom0);

        await Web3Connector.startArbitrage(amount, routers, addressList);
        console.log(`${token0};${token1};${loanPayFrom0};${grossLoanPayFrom0};${tradeAmountFrom0};${profitFrom0};${quoteFrom0.estimatedGas};${liquidityToken0}`);
      }
    })

    Web3Connector.getBlockNumber().then((block) => console.log(`[BLOCK LOG] from: ${token1} - to: ${token0} - block: ${block}`));
    OneInchBinanceChainService.quote({
      fromTokenAddress: token1,
      toTokenAddress: token0,
      amount: bigIntAmountToken1.toString(),
      protocols: Exchanges.join(',')
    }).then(async (quoteFrom1) => {
      const quoteFrom1Ratio = await this.getQuoteRatio(quoteFrom1);
      const loanPayFrom1 = bigDecimal.divide(amountToken1, poolRatioFrom1, 7);
      const grossLoanPayFrom1 = bigDecimal.multiply(loanPayFrom1, this.FEE);
      const tradeAmountFrom1 = bigDecimal.divide(amountToken1, quoteFrom1Ratio.netRatio, 7);
      const profitFrom1 = bigDecimal.subtract(tradeAmountFrom1, grossLoanPayFrom1);
      const hasProfitToken1 = bigDecimal.compareTo(profitFrom1, 0) >= 0
      const liquidityToken1 = this.hasLiquidity(bigIntAmountToken1, BigInt(poolReserves.token1Reserve))
      const hasLiquidityToken1 = bigDecimal.compareTo(liquidityToken1, '0.01') === -1;
      const runContractToken1 = hasLiquidityToken1 && hasProfitToken1;


      if(runContractToken1){
        const amount = bigIntAmountToken1;
        const [routers, addressList] = this.getRoutersAndPath(quoteFrom1);

        await Web3Connector.startArbitrage(amount, routers, addressList);
        console.log(`${token1};${token0};${loanPayFrom1};${grossLoanPayFrom1};${tradeAmountFrom1};${profitFrom1};${quoteFrom1.estimatedGas};${liquidityToken1}`);
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
        fromTokenAddress: Token.WBNB,
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

  public getRoutersAndPath(quote: QuoteResponse): [number[], string[]]{
    const routes: number[] = [];
    const path: string[] = []
    let first = true;
    const protocolsList = quote.protocols[0];
    for(const current of protocolsList){
      const protocol = current[0];
      if(first){
        path.push(protocol.fromTokenAddress);
        first = false;
      }

      path.push(protocol.toTokenAddress);
      if(protocol.name === 'ACRYPTOS'){
        routes.push(this.exchangeForACryptos(protocol));
      } else if(protocol.name === 'ELLIPSIS_FINANCE'){
        routes.push(this.exchangeForEllipsis(protocol));
      }else{
        routes.push(ExchangesMap[protocol.name as keyof typeof ExchangesMap]);
      }
    }
    return [routes, path];
  }

  public exchangeForACryptos(protocol: OneInchProtocol){
    if(!this.allowedTokensACryptos.has(protocol.fromTokenAddress) || !this.allowedTokensACryptos.has(protocol.toTokenAddress)){
      throw new Error('TOKEN NOT ALLOWED FOR ACRYPTOS');
    }

    if(protocol.fromTokenAddress === Token.VAI || protocol.toTokenAddress === Token.VAI){
      return ExchangesMap.ACRYPTOS_META;
    }
    return ExchangesMap.ACRYPTOS_CORE;
  }

  public exchangeForEllipsis(protocol: OneInchProtocol){
    if(!this.allowedTokensEllipsis.has(protocol.fromTokenAddress) || !this.allowedTokensEllipsis.has(protocol.toTokenAddress)){
      throw new Error('TOKEN NOT ALLOWED FOR ELLIPSIS');
    }

    if(protocol.fromTokenAddress === Token.DAI || protocol.toTokenAddress === Token.DAI){
      return ExchangesMap.ELLIPSIS_META;
    }
    return ExchangesMap.ELLIPSIS_CORE;
  }
}

export default new CryptoArbitrageService();
