import Status from "http-status-codes";
import axios from "axios";
import OneInchBaseAPIService from './OneInchBaseAPIService';
import { QuoteParams } from "./types/QuoteParams";
import { QuoteResponse } from "./types/QuoteResponse";

class OneInchBinanceChainService extends OneInchBaseAPIService {
  private CHAIN_CODE = "56";

  public async quote(params: QuoteParams) {
    const apiPath = `${this.BASE_URL}/${this.CHAIN_CODE}/quote?`

    const result = await axios.get<QuoteResponse>(apiPath, { params: params })

    if(result.status !== Status.OK) {
      throw new Error();
    }

    return result.data;
  }
}

export default new OneInchBinanceChainService();
