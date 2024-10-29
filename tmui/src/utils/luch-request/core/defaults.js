/**
 * 默认的全局配置
 */


export default {
  // baseURL: 'https://osl.fund/api',
  baseURL: 'http://127.0.0.1:8888',
  // baseURL: 'http://sazuky.natapp1.cc',
  header: {'Content-Type': 'application/json; charset=utf-8' },
  method: 'GET',
  dataType: 'json',
  paramsSerializer: null,
  // #ifndef MP-ALIPAY
  responseType: 'text',
  // #endif
  custom: {},
  // #ifdef H5 || APP-PLUS || MP-WEIXIN || MP-ALIPAY || MP-TOUTIAO || MP-KUAISHOU
  timeout: 60000,
  // #endif
  // #ifdef APP-PLUS
  sslVerify: false,
  // #endif
  // #ifdef H5
  withCredentials: false,
  // #endif
  // #ifdef APP-PLUS
  firstIpv4: false,
  // #endif
  validateStatus: function validateStatus(status) {
    return status >= 200 && status < 300
  },
  // 是否尝试将响应数据json化
  forcedJSONParsing: true
}
