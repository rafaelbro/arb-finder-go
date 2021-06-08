import Status from "http-status-codes";
import axios from "axios";
import OneInchBaseAPIService from './OneInchBaseAPIService';
import { QuoteParams } from "./types/QuoteParams";
import { QuoteResponse } from "./types/QuoteResponse";
import Exchanges from '../../exchanges';

class OneInchBinanceChainService extends OneInchBaseAPIService {
  private CHAIN_CODE = "56";

  public async quote(params: QuoteParams) {
    const apiPath = `${this.BASE_URL}/${this.CHAIN_CODE}/quote?`
    try {
      const result = await axios.get<QuoteResponse>(apiPath, { params: { ...params, parts: 1 }})

      if(result.status !== Status.OK) {
        console.log(result);
        throw new Error(JSON.stringify(result));
      }

      return result.data;
    } catch(error){
      throw error;
    }
  }
}

export default new OneInchBinanceChainService();
