import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const fromToken = Token[process.argv[2] as keyof typeof Token];
const amountToken0 = parseInt(process.argv[3]);
const toToken = Token[process.argv[4] as keyof typeof Token];
const amountToken1 = parseInt(process.argv[5]);

CryptoArbitrageService.canArbitrate(fromToken, amountToken0, toToken, amountToken1)
  .then(result => console.log(JSON.stringify(result)));
