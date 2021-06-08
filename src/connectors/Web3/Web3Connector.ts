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
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);

    const contract = new web3.eth.Contract(getReservesAbi as any, address);
    const result = await contract.methods.getReserves().call();
    return {
      token0Reserve: result._reserve0,
      token1Reserve: result._reserve1,
    } as ContractReserves
  }

  public async startArbitrage(amount: bigint, routers: number[], addressList: string[]){
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contractAddress = web3.utils.toChecksumAddress('0x38985C79Cb073542623339a4ab4268189f4871c8');
    web3.eth.accounts.wallet.add(account);
    web3.eth.defaultAccount = account.address;
    console.log(account.address);

    const addressListFormat = addressList.map(addr => web3.utils.toChecksumAddress(addr));
    console.log(`[CHAMOUUUUUUUUU] startArbitrage(${amount.toString()}, ${routers}, ${addressListFormat})`);

    const contractOptions = {
      from: account.address,
      gasPrice: '5000000000',
      gas: 20000000,
    }
    const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress, contractOptions);
    const result = await contract.methods.startArbitrage(amount.toString(), routers, addressListFormat).send();
    console.log(result);
  }

  public async getRouterIn(amount: bigint, routers: number[], addressList: string[]){
    const provider = new Web3.providers.HttpProvider(this.PROVIDER_URL);
    const web3 = new Web3(provider);
    const account = web3.eth.accounts.privateKeyToAccount(privateKey);
    const contractAddress = web3.utils.toChecksumAddress('0x38985C79Cb073542623339a4ab4268189f4871c8');
    web3.eth.accounts.wallet.add(account);
    web3.eth.defaultAccount = account.address;

    const addressListFormat = addressList.map(addr => web3.utils.toChecksumAddress(addr));
    console.log(`[CHAMOUUUUUUUUU] getRouterIn(0)`);

    const contract = new web3.eth.Contract(startArbitrageAbi as any, contractAddress);
    const result = await contract.methods.getRouterIn(0).call();
    console.log(result);
    await sleep(10000);
  }
}

export default new Web3Connector();
