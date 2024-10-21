declare namespace UniCloudNamespace {
  interface UniCloudResponseEvent {
    type: 'clientdb' | 'cloudobject' | 'cloudfunction';
    content: any;
  }

  interface UniCloudNeedLoginEvent {
    errCode: string | number;
    errMsg: string;
    uniIdRedirectUrl: string;
  }

  interface UniCloudRefreshTokenEvent {
    token: string;
    tokenExpired: number;
  }

  interface UniCloud {
    /** 事件回调 */
    on(eventName: string, callback: (result: any) => void): void;
    /**
     * 监听云函数、云对象、clientDB的响应
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-response](https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-response)
     */
    onResponse(callback: (result?: UniCloudResponseEvent) => void): void;
    /**
     * 移除监听云函数、云对象、clientDB的响应
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-response](https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-response)
     */
    offResponse(callback: (result?: UniCloudResponseEvent) => void): void;
    /**
     * 监听需要登录事件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-need-login](https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-need-login)
     */
    onNeedLogin(callback: (result?: UniCloudNeedLoginEvent) => void): void;
    /**
     * 移除监听需要登录事件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-need-login](https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-need-login)
     */
    offNeedLogin(callback: (result?: UniCloudNeedLoginEvent) => void): void;
    /**
     * 监听token刷新事件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-refresh-token](https://uniapp.dcloud.io/uniCloud/client-sdk.html#on-refresh-token)
     */
    onRefreshToken(callback: (result?: UniCloudRefreshTokenEvent) => void): void;
    /**
     * 移除监听token刷新事件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-refresh-token](https://uniapp.dcloud.io/uniCloud/client-sdk.html#off-refresh-token)
     */
    offRefreshToken(callback: (result: UniCloudRefreshTokenEvent) => void): void;
  }
}
