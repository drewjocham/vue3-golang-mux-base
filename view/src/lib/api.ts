import axios, {AxiosInstance} from "axios";
import { EnvironmentHelper } from "./EnvironmentHelper";
import {TestResponseRO} from "./models/TestResponse";

const url = new EnvironmentHelper()

const request: AxiosInstance = axios.create({
    baseURL: url.baseUrl,
    headers: {
        'content-type': 'application/json',
    },
    //params: {base64_encoded: 'true', fields: 'stdout'},
});

export const api = {

    async getTest() {
        try {
            const res = await request.get<TestResponseRO>("/v1/test");

            console.log(res.data)

            return res.data
        } catch (err) {
            console.error(err)
        }
    },

}


