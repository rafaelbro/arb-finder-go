import Status from "http-status-codes";
import axios from "axios";
import OneInchBaseAPIService from './OneInchBaseAPIService';
import { QuoteParams } from "./types/QuoteParams";
import { QuoteResponse } from "./types/QuoteResponse";

class OneInchBinanceChainService extends OneInchBaseAPIService {
  private CHAIN_CODE = "56";
  private ACCEPTED_EXCHANGES = []

  public async quote(params: QuoteParams) {
    const apiPath = `${this.BASE_URL}/${this.CHAIN_CODE}/quote?`
    try {
      const result = await axios.get<QuoteResponse>(apiPath, { params: { ...params, parts: 1 }})


      if(result.status !== Status.OK) {
        throw new Error();
      }

      return result.data;
    } catch(error){
      throw new error;
    }
  }
}

export default new OneInchBinanceChainService();
