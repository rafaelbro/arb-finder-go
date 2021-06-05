import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const swaps = {
  [Token.BNB]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT],
  [Token.ETH]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT],
  [Token.ADA]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT],
  [Token.WBNB]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT],
} as any

const keys = Object.keys(swaps);

keys.forEach(async (key) => {
  const toTokens: Token[] = swaps[key];

  for(const token of toTokens){
    const result = await CryptoArbitrageService.canArbitrate(10, key as Token, token as Token);
    console.log([result.fromToken, result.toToken, result.hasProfit, result.grossProfit, result.fee, result.netProfit].join(','))
  }
});
