import Request, {HttpRequestConfig} from '../utils/luch-request/index.js'
const http = new Request();


//用户注册
export async function registerApi(params: any ) {
    let RequestData= {
        // data:params
        ...params
    }
    return http.post('/cliUser/rigester', RequestData)
}
//用户登录
export async function loginApi(params: any ) {
    let RequestData= {
        // data:params
        ...params
    }

    return http.post('/cliLoad/login', RequestData)
}
//购买资产
export async function buyApi(params: any ) {
    let RequestData= {
        // data:params
        ...params
    }
    let config:HttpRequestConfig = {
        header:{'x-token':uni.getStorageSync('token')}
    }


    return http.post('/climainorder/buy', RequestData, config)
}