declare namespace UniCloudNamespace {
  interface CurrentUserInfo {
    /**
     * 当前用户uid
     */
    uid: string;
    /**
     * 当前用户角色列表
     */
    role: any[];
    /**
     * 当前用户权限列表
     */
    permission: any[];
  }

  interface UniCloudOptions {
    /**
     * 选择服务供应商
     * - tencent: 选择腾讯云作为服务商
     * - aliyun: 选择阿里云作为服务商
     */
    provider: 'tencent' | 'aliyun';
    /**
     * 服务空间ID
     */
    spaceId: string;
    /**
     * 服务空间对应的clientSecret
     */
    clientSecret?: string;
    /**
     * 服务空间地址
     */
    endpoint?: string;
  }

  interface UniError {
    /**
     * 错误模块名
     */
    errSubject?: string;
    /**
     * 错误码
     */
    errCode: number | string;
    /**
     * 错误信息
     */
    errMsg: string;
    /**
     * 请求id方便排错
     */
    requestId?: string;
    /**
     * 错误详情
     */
    detail?: any;
    /**
     * 上级错误
     */
    cause?: UniError;
  }

  interface CallFunctionResult {
    /**
     * 云函数执行结果
     */
    result: any;
    /**
     * 请求序列号，用于错误排查
     */
    requestId?: string;
  }

  enum SECRET_TYPE {
    /**
     * 不加密
     */
    none = 'none',
    /**
     * 仅请求参数加密
     */
    request = 'request',
    /**
     * 仅响应结果加密
     */
    response = 'response',
    /**
     * 请求响应均加密
     */
    both = 'both'
  }
  interface CallFunctionOptions {
    /**
     * 云函数名
     */
    name: string | string.CloudFunctionString;
    /**
     * 传递给云函数的参数
     */
    data?: any;
    /**
     * 安全网络类型
     * - none：不加密
     * - request：仅请求参数加密
     * - response：仅响应结果加密
     * - both：请求响应均加密
     */
    secretType?: keyof typeof SECRET_TYPE;
    /**
     * 成功返回的回调函数
     */
    success?: (result: CallFunctionResult) => void;
    /**
     * 失败返回的回调函数
     */
    fail?: (result: any) => void;
    /**
     * 结束的回调函数（调用成功、失败都会执行
     */
    complete?: (result: CallFunctionResult) => void;
  }

  interface ImportObjectLoadingOptions {
    /**
     * loading界面文字
     */
    text?: string;
    /**
     * loading是否显示透明遮罩
     */
    mask?: boolean;
  }

  interface ImportObjectErrorOptions {
    /**
     * 错误提示类型，modal | toast
     */
    type?: 'modal' | 'toast';
    /**
     * 是否显示重试按钮，type为modal时生效
     */
    retry?: boolean;
  }

  interface ParseSystemErrorOptions {
    /**
     * 云对象名
     */
    objectName: string;
    /**
     * 调用的方法名
     */
    methodName: string;
    /**
     * 参数
     */
    params: Array<any>;
    /**
     * 错误码
     */
    errCode: string | number;
    /**
     * 错误信息
     */
    errMsg: string;
  }

  interface ParsedSystemErrorResult {
    errMsg: string;
  }

  interface ImportObjectOptions {
    /**
     * 是否移除自动展示的ui
     */
    customUI?: boolean;
    /**
     * loading界面配置
     */
    loadingOptions?: ImportObjectLoadingOptions;
    /**
     * 错误提示配置
     */
    errorOptions?: ImportObjectErrorOptions;
    /**
     * 使用安全网络的方法及安全网络类型
     */
    secretMethods?: Record<string, keyof typeof SECRET_TYPE>;
    /**
     * 转化云对象内未捕获的错误或客户端网络错误
     */
    parseSystemError?: (options: ParseSystemErrorOptions) => Promise<ParsedSystemErrorResult> | ParsedSystemErrorResult;
  }

  interface InitSecureNetworkByWeixinOptions {
    /**
     * 是否握手并自动调用uni-id-co的微信登录，默认仅调用uni-id-co的secureNetworkHandshakeByWeixin不调用微信登录
     */
    callLoginByWeixin?: boolean;
    /**
     * 用户openid，传此参数时不会调用uni-id-co的任何方法
     */
    openid?: string;
  }

  interface InitSecureNetworkByWeixinResponse {
    code?: string;
  }

  interface UniCloud {
    /** 用于快速开发datacom规范的组件 */
    mixinDatacom: any;
    /**
     * 服务空间初始化，返回uniCloud实例
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/init](https://uniapp.dcloud.io/uniCloud/init)
     */
    init(options: UniCloudOptions): UniCloud;
    /**
     * 设置自定义clientInfo信息
     *
     * 文档: [https://doc.dcloud.net.cn/uniCloud/client-sdk.html#set-custom-client-info](https://doc.dcloud.net.cn/uniCloud/client-sdk.html#set-custom-client-info)
     */
    setCustomClientInfo(options: object): void;
    /**
     * 调用云函数
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cf-functions?id=clientcallfunction](https://uniapp.dcloud.io/uniCloud/cf-functions?id=clientcallfunction)
     */
    callFunction(options: CallFunctionOptions): Promise<any>;
    /**
     * 引用云对象
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/cloud-obj](https://uniapp.dcloud.io/uniCloud/cloud-obj)
     */
    importObject(objectName: string | string.CloudObjectString, importObjectOptions?: ImportObjectOptions): any;
    /**
     * 获取当前用户缓存在token内的信息
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#client-getcurrentuserinfo](https://uniapp.dcloud.io/uniCloud/client-sdk.html#client-getcurrentuserinfo)
     */
    getCurrentUserInfo(): CurrentUserInfo;
    /**
     * 微信小程序安全网络初始化
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#client-getcurrentuserinfo](https://uniapp.dcloud.io/uniCloud/client-sdk.html#init-secure-network-by-weixin)
     */
    initSecureNetworkByWeixin(options?: InitSecureNetworkByWeixinOptions): Promise<InitSecureNetworkByWeixinResponse>;
  }
}
