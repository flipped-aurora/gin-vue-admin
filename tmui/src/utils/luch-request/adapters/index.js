import buildURL from '../helpers/buildURL'
import buildFullPath from '../core/buildFullPath'
import settle from '../core/settle'
import {isUndefined} from "../utils"

/**
 * 返回可选值存在的配置
 * @param {Array} keys - 可选值数组
 * @param {Object} config2 - 配置
 * @return {{}} - 存在的配置项
 */
const mergeKeys = (keys, config2) => {
  let config = {}
  keys.forEach(prop => {
    if (!isUndefined(config2[prop])) {
      config[prop] = config2[prop]
    }
  })
  return config
}
export default (config) => {
  return new Promise((resolve, reject) => {
    let fullPath = buildURL(buildFullPath(config.baseURL, config.url), config.params, config.paramsSerializer)
    const _config = {
      url: fullPath,
      header: config.header,
      complete: (response) => {
        config.fullPath = fullPath
        response.config = config
        response.rawData = response.data
        try {
          let jsonParseHandle = false
          const forcedJSONParsingType = typeof config.forcedJSONParsing
          if (forcedJSONParsingType === 'boolean') {
            jsonParseHandle = config.forcedJSONParsing
          } else if (forcedJSONParsingType === 'object') {
            const includesMethod = config.forcedJSONParsing.include || []
            jsonParseHandle = includesMethod.includes(config.method)
          }

          // 对可能字符串不是json 的情况容错
          if (jsonParseHandle && typeof response.data === 'string') {
            response.data = JSON.parse(response.data)
          }
          // eslint-disable-next-line no-empty
        } catch (e) {
        }
        settle(resolve, reject, response)
      }
    }
    let requestTask
    if (config.method === 'UPLOAD') {
      delete _config.header['content-type']
      delete _config.header['Content-Type']
      let otherConfig = {
        // #ifdef MP-ALIPAY
        fileType: config.fileType,
        // #endif
        filePath: config.filePath,
        name: config.name
      }
      const optionalKeys = [
        // #ifdef APP-PLUS || H5
        'files',
        // #endif
        // #ifdef H5
        'file',
        // #endif
        // #ifdef H5 || APP-PLUS || MP-WEIXIN || MP-ALIPAY || MP-TOUTIAO || MP-KUAISHOU
        'timeout',
        // #endif
        'formData'
      ]
      requestTask = uni.uploadFile({..._config, ...otherConfig, ...mergeKeys(optionalKeys, config)})
    } else if (config.method === 'DOWNLOAD') {
      const optionalKeys = [
        // #ifdef H5 || APP-PLUS || MP-WEIXIN || MP-ALIPAY || MP-TOUTIAO || MP-KUAISHOU
        'timeout',
        // #endif
        // #ifdef MP
        'filePath',
        // #endif
      ]
      requestTask = uni.downloadFile({..._config, ...mergeKeys(optionalKeys, config)})
    } else {
      const optionalKeys = [
        'data',
        'method',
        // #ifdef H5 || APP-PLUS || MP-ALIPAY || MP-WEIXIN
        'timeout',
        // #endif
        'dataType',
        // #ifndef MP-ALIPAY
        'responseType',
        // #endif
        // #ifdef APP-PLUS
        'sslVerify',
        // #endif
        // #ifdef H5
        'withCredentials',
        // #endif
        // #ifdef APP-PLUS
        'firstIpv4',
        // #endif
        // #ifdef MP-WEIXIN
        'enableHttp2',
        'enableQuic',
        // #endif
        // #ifdef MP-TOUTIAO || MP-WEIXIN
        'enableCache',
        // #endif
        // #ifdef MP-WEIXIN
        'enableHttpDNS',
        'httpDNSServiceId',
        'enableChunked',
        'forceCellularNetwork',
        // #endif
        // #ifdef MP-ALIPAY
        'enableCookie',
        // #endif
        // #ifdef MP-BAIDU
        'cloudCache',
        'defer'
        // #endif
      ]
      requestTask = uni.request({..._config, ...mergeKeys(optionalKeys, config)})
    }
    if (config.getTask) {
      config.getTask(requestTask, config)
    }
  })
}
