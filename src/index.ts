import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const amount = parseInt(process.argv[2]);
const fromToken = Token[process.argv[3] as keyof typeof Token];
const toToken = Token[process.argv[4] as keyof typeof Token];

CryptoArbitrageService.canArbitrate(amount, fromToken, toToken)
  .then(result => console.log(result));
