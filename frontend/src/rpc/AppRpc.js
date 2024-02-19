import GoJsRpc from "./GoJsRpc";

export default class AppRpc extends GoJsRpc {
    constructor() {
        super("http://localhost:3000/gorpc")
    }

    async sum(params) {
        const res = await this.call('sum', params)

        console.log('sum res: ', res)
        console.log('success: ', res.data.Success)
        console.log('sum res: ', res.data.Data)

        if(res.data.Success === true) {
            return res.data.Data
        } else {
            throw new Error(res.data.Data)
        }
    }
}