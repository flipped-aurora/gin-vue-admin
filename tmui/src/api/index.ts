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
//查询设置信息
export async function getSettingApi(params: any ) {
    let RequestData= {
        // data:params
        ...params
    }
    return http.get('/cliset/getCliSetPublic', RequestData)
}

//查询结算详情记录
export async function getProfitListApi(params: any ) {
    let RequestData= {
        data:params
        // ...params
    }
    return http.get('/cliprofit/getCliProfitPublic', RequestData)
}
//查询结算总表
export async function getProfitMainApi(params: any ) {
    let RequestData= {
        data:params
        // ...params
    }
    return http.get('/climainprofit/getCliMainprofitPublic', RequestData)
}
//集合查询订单信息
export async function getOrderListApi(params: any ) {
    let RequestData= {
        data:params
        // ...params
    }
    return http.get('/cliOrder/getCliOrderPublic', RequestData)
}
//聚合查询邀请信息
export async function getInviteListApi(params: any ) {
    let RequestData= {
        data:params
        // ...params
    }
    return http.get('/cliTree/getCliTreePublic', RequestData)
}
//聚合查询提币信息
export async function getWithdrawListApi(params: any ) {
    let RequestData= {
        data:params
        // ...params
    }
    return http.get('/cliwithdraw/getCliWithdrawPublic', RequestData)
}