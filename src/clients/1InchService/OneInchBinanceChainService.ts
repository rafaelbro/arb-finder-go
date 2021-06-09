import Status from "http-status-codes";
import axios from "axios";
import OneInchBaseAPIService from './OneInchBaseAPIService';
import { QuoteParams } from "./types/QuoteParams";
import { QuoteResponse } from "./types/QuoteResponse";
import Exchanges from '../../exchanges';

class OneInchBinanceChainService extends OneInchBaseAPIService {
  private CHAIN_CODE = "56";

  public async quote(params: QuoteParams) {
    // console.time(`1Inch ${params.fromTokenAddress} - ${params.toTokenAddress}`);
    const apiPath = `${this.BASE_URL}/${this.CHAIN_CODE}/quote?`
    try {
      const result = await axios.get<QuoteResponse>(apiPath, { params: { ...params, parts: 1 }})

      if(result.status !== Status.OK) {
        console.log(result);
        throw new Error(JSON.stringify(result));
      }

      // console.timeEnd(`1Inch ${params.fromTokenAddress} - ${params.toTokenAddress}`);
      return result.data;
    } catch(error){
      // console.timeEnd(`1Inch ${params.fromTokenAddress} - ${params.toTokenAddress}`);
      throw error;

    }
  }
}

export default new OneInchBinanceChainService();
