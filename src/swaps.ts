import { Pair, PAIRS, TOKEN_VALUE } from "./pairs";
import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const token0List = Object.keys(PAIRS);

token0List.forEach(async (token0: any) => {
  const token1Map: Record<Token, Pair> = PAIRS[token0];
  const token1List = Object.keys(token1Map);
  const amount0 = TOKEN_VALUE[token0 as Token];



  for(const token1 of token1List){
    const amount1 = TOKEN_VALUE[token1 as Token];
    const result = await CryptoArbitrageService.canArbitrate(token0 as Token, amount0, token1 as Token, amount1);
    if (result.fromToken0.hasProfit || result.fromToken1.hasProfit) {
      console.log([
        `"${result.timestamp}"`,
        `"${'FROM TOKEN 0'}"`,
        `"${result.fromToken0.hasProfit}"`,
        `"${result.fromToken0.poolRatio}"`,
        `"${result.fromToken0.fromToken}"`,
        `"${result.fromToken0.toToken}"`,
        `"${result.fromToken0.quoteRatio}"`,
        `"${result.fromToken0.loanAmount}"`,
        `"${result.fromToken0.tradeAmount}"`,
        `"${result.fromToken0.profit}"`,
        `"${JSON.stringify(result.fromToken0.protocols)}"`,
        `"${'FROM TOKEN 1'}"`,
        `"${result.fromToken1.hasProfit}"`,
        `"${result.fromToken1.poolRatio}"`,
        `"${result.fromToken1.fromToken}"`,
        `"${result.fromToken1.toToken}"`,
        `"${result.fromToken1.quoteRatio}"`,
        `"${result.fromToken1.loanAmount}"`,
        `"${result.fromToken1.tradeAmount}"`,
        `"${result.fromToken1.profit}"`,
        `"${JSON.stringify(result.fromToken1.protocols)}"`,
      ].join(';'))
    }
  }
});
