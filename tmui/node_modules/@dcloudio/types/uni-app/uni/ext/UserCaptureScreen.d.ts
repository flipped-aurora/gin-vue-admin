declare namespace UniNamespace {
  type OnUserCaptureScreenCallback = (res?: { errMsg: string }) => void;
}

interface Uni {
  /**
   * 监听用户主动截屏事件，用户使用系统截屏按键截屏时触发此事件。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/capture-screen.html#onusercapturescreen](https://uniapp.dcloud.net.cn/api/system/capture-screen.html#onusercapturescreen)
   */
  onUserCaptureScreen(callback: UniNamespace.OnUserCaptureScreenCallback): void;
  /**
   * 用户主动截屏事件。取消事件监听。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/capture-screen.html#offusercapturescreen](https://uniapp.dcloud.net.cn/api/system/capture-screen.html#offusercapturescreen)
   */
  offUserCaptureScreen(callback: UniNamespace.OnUserCaptureScreenCallback): void;
}
