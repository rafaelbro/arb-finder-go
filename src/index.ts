import Web3Connector from "./connectors/Web3/Web3Connector";
import Exchanges from "./exchanges";
import { TOKEN_VALUE } from "./pairs";
import CryptoArbitrageService from "./service/CryptoArbitrageService";
import Token from "./tokens";

const token0 = Token[process.argv[2] as keyof typeof Token];
const token1 = Token[process.argv[3] as keyof typeof Token];

const amount0 = TOKEN_VALUE[token0 as Token];
const amount1 = TOKEN_VALUE[token1 as Token];

CryptoArbitrageService.canArbitrate(token0, amount0, token1, amount1);
  // .then(async (result) => {
  //   const fromToken0 = result.fromToken0
  //   const fromToken1 = result.fromToken1;

  //   let amount: bigint;
  //   let routers: number[] = [];
  //   let addressList: string[] = [];

  //   if(fromToken0.runContract){
  //     amount = BigInt(amount0) * (10n ** 18n);
  //     routers = fromToken0.protocols[0].map((p) => Exchanges[p[0].name as keyof typeof Exchanges]);
  //     addressList = fromToken0.protocols[0].map((p) => p[0].toTokenAddress);
  //     addressList.unshift(token0);

  //     await Web3Connector.startArbitrage(amount, routers, addressList);
  //   } else if(fromToken1 && fromToken1.runContract) {
  //     amount = BigInt(amount1) * (10n ** 18n);
  //     routers = fromToken1.protocols[0].map((p) => Exchanges[p[0].name as keyof typeof Exchanges]);
  //     addressList = fromToken1.protocols[0].map((p) => p[0].toTokenAddress);
  //     addressList.unshift(token1);

  //     await Web3Connector.startArbitrage(amount, routers, addressList);
  //   }

  //   if (result.fromToken0.runContract || (result.fromToken1 && result.fromToken1.runContract)) {
  //     console.log(JSON.stringify(result));
  //     console.log("");
  //     if(result.fromToken0.runContract){
  //       console.log(JSON.stringify(result.fromToken0.protocols));
  //     }else if(result.fromToken1) {
  //       console.log(JSON.stringify(result.fromToken1.protocols));
  //     }
  //   }
  // });
