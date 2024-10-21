declare namespace UniNamespace {
  interface OnUnhandledRejectionCallbackResult {
    /**
     * 被拒绝的 Promise 对象
     */
    promise: Promise<any>;
    /**
     * 拒绝原因，一般是一个 Error 对象
     */
    reason: string;
  }

  /**
   * 未处理的 Promise 拒绝事件的回调函数
   */
  type OnUnhandledRejectionCallback = (result: OnUnhandledRejectionCallbackResult) => void;
}

interface Uni {
  /**
   * 监听未处理的 Promise 拒绝事件。该事件与 `App.onUnhandledRejection` 的回调时机与参数一致。
   *
   * **注意**
   *
   *
   * - 安卓平台暂时不支持该事件
   * - 所有的 unhandledRejection 都可以被这一监听捕获，但只有 Error 类型的才会在小程序后台触发报警。
   *
   * 最低基础库： `2.10.0`
   *
   * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onUnhandledRejection.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onUnhandledRejection.html)
   */
  onUnhandledRejection(callback: UniNamespace.OnUnhandledRejectionCallback): void;

  /**
   * 取消监听未处理的 Promise 拒绝事件
   *
   * 最低基础库： `2.10.0`
   *
   * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offUnhandledRejection.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offUnhandledRejection.html)
   */
  offUnhandledRejection(callback: UniNamespace.OnUnhandledRejectionCallback): void;
}
