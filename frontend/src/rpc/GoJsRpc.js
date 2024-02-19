import axios from 'axios'

export default class GoJsRpc {
    
    constructor(endpoint) {
        this.endpoint = endpoint
    }

    async call(method, params) {
        const data = {
            method,
            params: this.prepareParams(params)
        }

        return await axios.post(this.endpoint, data)
    }

    prepareParams(params) {
        const p = {}

        for (const key in params) {
            p[key] = {
                value: params[key]
            }
        }

        return p
    }

}