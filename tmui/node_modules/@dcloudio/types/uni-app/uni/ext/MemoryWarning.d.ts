declare namespace UniNamespace {
  interface OnMemoryWarningCallbackResult {
    /**
     * 仅 Android 有该字段，对应系统内存告警等级宏定义
     */
    level: number;
  }
  type OnMemoryWarningCallback = (res: OnMemoryWarningCallbackResult) => void;
}
interface Uni {
  /**
   * 监听内存不足告警事件。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/memory.html#onmemorywarning](https://uniapp.dcloud.net.cn/api/system/memory.html#onmemorywarning)
   */
  onMemoryWarning(callback: UniNamespace.OnMemoryWarningCallback): void;
  /**
   * 取消监听内存不足告警事件。不传入 callback 则取消所有监听。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/memory.html#offmemorywarning](https://uniapp.dcloud.net.cn/api/system/memory.html#offmemorywarning)
   */
  offMemoryWarning(callback?: UniNamespace.OnMemoryWarningCallback): void;
}
