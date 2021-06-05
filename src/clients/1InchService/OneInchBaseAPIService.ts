import { AxiosResponse } from "axios";
import IBaseClientService from "../IBaseClientService";
import { QuoteParams } from "./types/QuoteParams";
import { QuoteResponse } from "./types/QuoteResponse";

export default abstract class OneInchBaseAPIService implements IBaseClientService{
   BASE_URL = "https://api.1inch.exchange/v3.0";
   abstract quote(params: QuoteParams): Promise<QuoteResponse>;
}

