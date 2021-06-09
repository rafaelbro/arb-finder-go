import Web3 from "web3"
import Exchanges from "../../exchanges";
import { getReservesAbi } from "../abis/getReservesAbi";
import { startArbitrageAbi } from "../abis/startArbitrageAbi";
import { ContractReserves } from "./Web3ConnectorTypes"
import sleep from "sleep-promise";
import { privateKey } from "../../ppk";

class Web3Connector {
  private PROVIDER_URL = 'https://bsc-dataseed1.binance.org:443';

  public async getReserves(address: string){
    //console.time("Web3 GetReserves");
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);

    const contract = new web3.eth.Contract(getReservesAbi as any, address);
    const result = await contract.methods.getReserves().call();
    //console.timeEnd("Web3 GetReserves");
    return {
      token0Reserve: result._reserve0,
      token1Reserve: result._reserve1,
    } as ContractReserves
  }

  public async startArbitrage(amount: bigint, routers: number[], addressList: string[]){
    // console.time("Web3 startArbitrage");
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contractAddress = web3.utils.toChecksumAddress('0x3389ee611BFC16Aaf7E604E8c13a7D985468514E');
    web3.eth.accounts.wallet.add(account);
    web3.eth.defaultAccount = account.address;
    console.log(account.address);

    const addressListFormat = addressList.map(addr => web3.utils.toChecksumAddress(addr));

    const contractOptions = {
      from: account.address,
      gasPrice: '5000000000',
      gas: 20000000,
    }
    const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress, contractOptions);
    if(process.env['PROD'] === 'true'){
      const result = await contract.methods.startArbitrage(amount.toString(), routers, addressListFormat).send();
      console.log(`[CHAMOUUUUUUUUU] startArbitrage(${amount.toString()}, ${routers}, ${addressListFormat})`);
      console.log(JSON.stringify(result));
    }
    // console.timeEnd("Web3 startArbitrage");
  }

  public async getRouterIn(amount: number){
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contractAddress = web3.utils.toChecksumAddress('0x3389ee611BFC16Aaf7E604E8c13a7D985468514E');
    web3.eth.accounts.wallet.add(account);
    web3.eth.defaultAccount = account.address;


    const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress);
    const result = await contract.methods.getRouterIn(amount).call();
    console.log(`getRouterIn(${amount}) = ${result}`);
  }
}

export default new Web3Connector();
