import Web3 from "web3"
import { Exchanges } from "../../exchanges";
import { getReservesAbi } from "../abis/getReservesAbi";
import { startArbitrageAbi } from "../abis/startArbitrageAbi";
import { ContractReserves } from "./Web3ConnectorTypes"
import sleep from "sleep-promise";
import { privateKey } from "../../ppk";

class Web3Connector {
  private getProviderURL(): string {
    if(process.env['ENV'] === 'prod'){
      return 'https://bsc-dataseed1.binance.org:443'
    }
    return 'http://127.0.0.1:8545'
  }

  public async getReserves(address: string){
    const provider = new Web3.providers.HttpProvider(this.getProviderURL());
    const web3 = new Web3(provider);

    const contract = new web3.eth.Contract(getReservesAbi as any, address);
    const result = await contract.methods.getReserves().call();
    return {
      token0Reserve: result._reserve0,
      token1Reserve: result._reserve1,
    } as ContractReserves
  }

  public async startArbitrage(amount: bigint, routers: number[], addressList: string[]){
    const provider = new Web3.providers.HttpProvider(this.getProviderURL());
    const web3 = new Web3(provider);
    const addressListFormat = addressList.map(addr => web3.utils.toChecksumAddress(addr));

    console.log(`[WOULD CONTRACT] startArbitrage(${amount.toString()}, ${routers}, ${addressListFormat}) @ ${await this.getBlockNumber()}`);

    if(process.env['CALL'] === 'true'){
      const account = web3.eth.accounts.privateKeyToAccount(privateKey);
      const contractAddress = web3.utils.toChecksumAddress('0x3389ee611BFC16Aaf7E604E8c13a7D985468514E');
      web3.eth.accounts.wallet.add(account);
      web3.eth.defaultAccount = account.address;
      const contractOptions = {
        from: account.address,
        gasPrice: '5000000000',
        gas: 20000000,
      }

      const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress, contractOptions);
      const result = await contract.methods.startArbitrage(amount.toString(), routers, addressListFormat).send();
      console.log(JSON.stringify(result));
    }
  }

  public async getRouterIn(amount: number){
    const provider = new Web3.providers.HttpProvider(this.getProviderURL());
    const web3 = new Web3(provider);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contractAddress = web3.utils.toChecksumAddress('0x3389ee611BFC16Aaf7E604E8c13a7D985468514E');
    web3.eth.accounts.wallet.add(account);
    web3.eth.defaultAccount = account.address;

    const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress);
    const result = await contract.methods.getRouterIn(amount).call();
    console.log(`getRouterIn(${amount}) = ${result}`);
  }

  public async getBlockNumber(): Promise<number>{
    const provider = new Web3.providers.HttpProvider(this.getProviderURL());
    const web3 = new Web3(provider);
    return await web3.eth.getBlockNumber();
  }
}

export default new Web3Connector();
