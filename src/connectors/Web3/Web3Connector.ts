import Web3 from "web3"
import { getReservesAbi } from "../abis/getReservesAbi";
import { ContractReserves } from "./Web3ConnectorTypes"

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
}

export default new Web3Connector();
