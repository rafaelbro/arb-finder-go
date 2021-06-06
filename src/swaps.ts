import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const swaps = {
  [Token.BNB]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT, Token.ETH, Token.ADA, Token.WBNB],
  [Token.ETH]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT, Token.BNB, Token.ADA, Token.WBNB],
  [Token.ADA]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT, Token.BNB, Token.ETH, Token.WBNB],
  [Token.WBNB]: [ Token.DAI, Token.BUSD, Token.USDC, Token.USDT, Token.BNB, Token.ETH, Token.ADA],
} as any

const keys = Object.keys(swaps);

keys.forEach(async (key) => {
  const toTokens: Token[] = swaps[key];

  for(const token of toTokens){
    // const result = await CryptoArbitrageService.canArbitrate(10, key as Token, token as Token);
    // console.log([result.fromToken, result.toToken, result.hasProfit, result.grossProfit, result.fee, result.netProfit].join(','))
  }
});
