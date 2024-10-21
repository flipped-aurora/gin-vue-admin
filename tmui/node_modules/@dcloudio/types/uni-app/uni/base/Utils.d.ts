export interface AsyncOptions {
  /**
   * 接口调用成功的回调函数
   */
  success?: (result: any) => void;
  /**
   * 接口调用失败的回调函数
   */
  fail?: (result: any) => void;
  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  complete?: (result: any) => void;
}
