import { AxiosResponse } from "axios";
import { QuoteParams } from "./1InchService/types/QuoteParams";
import { QuoteResponse } from "./1InchService/types/QuoteResponse";

interface IBaseClientService{
  BASE_URL: string
  quote(params: QuoteParams): Promise<QuoteResponse>;
}

export default IBaseClientService;
