import { Pair, PAIRS } from "./pairs";
import Token from "./tokens";
import { execSync } from "child_process";

const token0List = Object.keys(PAIRS);

const getKey = (value: string) => {
  return Object.keys(Token).find(x => Token[x as keyof typeof Token] === value);
}

for(const token0 of token0List) {
  const token1Map: Record<Token, Pair> = PAIRS[token0];
  const token1List = Object.keys(token1Map);

  for(const token1 of token1List){
    console.log(`Tokens: ${getKey(token0)} --- ${getKey(token1)}`);
    console.time("start");
    execSync(`PROD=${process.env['PROD']} node dist/index.js ${getKey(token0)} ${getKey(token1)} >> out`);
    console.timeEnd("start");
  }
};
