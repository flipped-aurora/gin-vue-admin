declare namespace UniNamespace {
  interface GeneralCallbackResult {
    /**
     * 错误信息
     */
    errMsg: string;
  }

  /**
   * 小程序错误事件的监听函数
   */
  type OnAppErrorCallback = (
    /** 错误信息，包含堆栈 */
    error: string
  ) => void;

  /**
   * onError 传入的监听函数。不传此参数则移除所有监听函数。
   */
  type OffAppErrorCallback = (res: GeneralCallbackResult) => void;

  interface SourceError {
    subject?: string;
    code?: number;
    message?: string;
    cause?: SourceError | AggregateError;
  }

  // 聚合源错误信息
  interface AggregateError extends SourceError {
    errors: Array<SourceError | AggregateError>;
  }

  interface UniError {
    errSubject: string;
    errCode: number;
    errMsg: string;
    data?: object;
    cause?: SourceError | AggregateError;
  }
}

interface Uni {
  /**
   *
   * 需要基础库： `2.1.2`
   *
   * 在插件中使用：不支持
   *
   * 监听小程序错误事件。如脚本错误或 API 调用报错等。该事件与 [`App.onError`](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onerrorstring-error) 的回调时机与参数一致。
   *
   * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onError.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onError.html)
   */
  onError(listener: UniNamespace.OnAppErrorCallback): void;

  /**
   *
   * 需要基础库： `2.1.2`
   *
   * 在插件中使用：不支持
   *
   * 移除小程序错误事件的监听函数
   *
   *
   * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offError.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offError.html)
   */
  offError(listener?: UniNamespace.OffAppErrorCallback): void;
}
