declare namespace UniNamespace {
  interface StartFacialRecognitionVerifyCallbackResult {
    /**
     * 错误码，成功时为0
     */
    errCode: number;
    /**
     * 错误信息
     */
    errMsg: string;
    /**
     * 抛出错误的模块/主题名
     */
    errSubject?: string;
    /**
     * 引起此错误的下层错误
     */
    cause?: any;
  }
  interface StartFacialRecognitionVerifyOption {
    /**
     * 认证流水号，由服务端根据接入的业务模式调用对应的初始化接口获取
     */
    certifyId: string;
    /**
     * 刷脸圈的颜色
     */
    progressBarColor?: string;
    /**
     * 认证界面UI朝向。port 为竖屏，land 为横屏，默认为 port
     */
    screenOrientation?: string;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: (result: StartFacialRecognitionVerifyCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    fail?: (result: StartFacialRecognitionVerifyCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    success?: (result: StartFacialRecognitionVerifyCallbackResult) => void;
  }
}

interface Uni {
  /**
   * 获取设备信息MetaInfo
   *
   * 文档: [https://uniapp.dcloud.net.cn/uniCloud/frv/dev.html#get-meta-info](https://uniapp.dcloud.net.cn/uniCloud/frv/dev.html#get-meta-info)
   */
  getFacialRecognitionMetaInfo(): string;
  /**
   * 调起实人认证界面
   *
   * 文档: [https://uniapp.dcloud.net.cn/uniCloud/frv/dev.html#start-frv](https://uniapp.dcloud.net.cn/uniCloud/frv/dev.html#start-frv)
   */
  startFacialRecognitionVerify(options: UniNamespace.StartFacialRecognitionVerifyOption): void;
}
