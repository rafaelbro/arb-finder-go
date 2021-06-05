import OneInchBinanceChainService from "../clients/1InchService/OneInchBinanceChainService";
import { QuoteResponse } from "../clients/1InchService/types/QuoteResponse";
import Token from "../tokens";
import { ArbitrageHumanReadableValues, ArbitrageResult, ArbitrageValues } from "../types/ArbitrageResult";

class CryptoArbitrageService {
  public async canArbitrate(
    amount: number,
    token1: Token,
    token2: Token,
    humanReadable = false
  ): Promise<ArbitrageResult> {

    const bigIntAmount = BigInt(amount) * (10n ** 18n)

    const operation1 = await OneInchBinanceChainService.quote({
      fromTokenAddress: token1,
      toTokenAddress: token2,
      amount: bigIntAmount.toString()
    });

    const operation2 = await OneInchBinanceChainService.quote({
      fromTokenAddress: token2,
      toTokenAddress: token1,
      amount: operation1.toTokenAmount
    });

    const values = await this.arbitrageHasProfit([operation1, operation2]);

    const result = {
      fromToken: token1,
      toToken: token2,
      ...values,
    } as ArbitrageResult;

    if (humanReadable){
      result['humanReadable'] = this.humanReadable(values);
    }

    return result;
  }

  public async arbitrageHasProfit(operations: QuoteResponse[]): Promise<ArbitrageValues> {
    const count = operations.length;
    const first = operations[0];
    const last = operations[count - 1];
    const fromToken = first.fromToken.address;
    const grossProfit = BigInt(last.toTokenAmount) - BigInt(first.fromTokenAmount)

    let protocols = [];
    let estimatedGas = 0;
    for(const operation of operations){
      estimatedGas += operation.estimatedGas;
      protocols.push(operation.protocols[0].map((protocol) => protocol.map(p => p.name )).flat());
    }

    const feeInSourceToken = await this.convertGasToToken(estimatedGas, fromToken);
    const netProfit = grossProfit - feeInSourceToken;

    return {
      hasProfit: netProfit > 0,
      grossProfit: grossProfit.toString(),
      fee: feeInSourceToken.toString(),
      netProfit: netProfit.toString(),
      protocols: protocols,
    }
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

  private humanReadable(values: ArbitrageValues): ArbitrageHumanReadableValues{
    return {
      grossProfit: (parseFloat(values.grossProfit) / 10 ** 18).toString(),
      fee: (parseFloat(values.fee) / 10 ** 18).toString(),
      netProfit: (parseFloat(values.netProfit) / 10 ** 18).toString()
    }
  }

}

export default new CryptoArbitrageService();
