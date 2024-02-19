import axios from 'axios'

export default class GoJsRpc {
    
    constructor(endpoint) {
        this.endpoint = endpoint

        return new Proxy(this, {
            get: function(gojsrpc, field) {
                if (field in gojsrpc){ 
                    return gojsrpc[field]; // normal case
                }

                console.log('__call: ', field)

                return function (params) {
                    // This function will be executed when property is accessed as a function
                    return this.call(field, params)
                }
            }
        });
    }

    async call(method, params) {
        return new Promise(async (resolve, reject) => {
            console.log('calling: ', method, params)

            const data = {
                method,
                params: this.prepareParams(params)
            }

            const res = await axios.post(this.endpoint, data)

            if(res.data.Success === true) {
                resolve(res.data.Data)
            } else {
                reject(new Error(res.data.Data))
            }
        })
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