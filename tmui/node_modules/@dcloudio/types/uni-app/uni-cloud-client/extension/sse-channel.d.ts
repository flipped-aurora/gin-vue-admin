declare namespace UniCloudNamespace {
  class SSEChannel {
    constructor();
    /**
     * 开启消息通道
     *
     * 文档：[https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#channel-open](https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#channel-open)
     */
    open: () => Promise<void>;
    /**
     * 关闭消息通道
     *
     * 文档：[https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#channel-close](https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#channel-close)
     */
    close: () => void;
    /**
     * 添加通道开启事件监听器
     */
    on(event: 'open', callback: () => any): void;
    /**
     * 添加消息接收事件监听器
     */
    on(event: 'message', callback: (message?: any) => any): void;
    /**
     * 添加通道消息完毕事件监听器
     */
    on(event: 'end', callback: (message?: any) => any): void;
    /**
     * 添加通道关闭事件监听器
     */
    on(event: 'close', callback: () => any): void;
    /**
     * 添加通道错误事件监听器
     */
    on(event: 'error', callback: (err?: any) => any): void;
    /**
     * 添加通道开启事件监听器
     */
    addListener(event: 'open', callback: () => any): void;
    /**
     * 添加消息接收事件监听器
     */
    addListener(event: 'message', callback: (message?: any) => any): void;
    /**
     * 添加通道消息完毕事件监听器
     */
    addListener(event: 'end', callback: (message?: any) => any): void;
    /**
     * 添加通道关闭事件监听器
     */
    addListener(event: 'close', callback: () => any): void;
    /**
     * 添加通道错误事件监听器
     */
    addListener(event: 'error', callback: (err?: any) => any): void;
    /**
     * 移除通道开启事件监听器
     */
    off(event: 'open', callback: () => any): void;
    /**
     * 移除消息接收事件监听器
     */
    off(event: 'message', callback: (message?: any) => any): void;
    /**
     * 移除通道消息完毕事件监听器
     */
    off(event: 'end', callback: (message?: any) => any): void;
    /**
     * 移除通道关闭事件监听器
     */
    off(event: 'close', callback: () => any): void;
    /**
     * 移除通道错误事件监听器
     */
    off(event: 'error', callback: (err?: any) => any): void;
    /**
     * 移除通道开启事件监听器
     */
    removeListener(event: 'open', callback: () => any): void;
    /**
     * 移除消息接收事件监听器
     */
    removeListener(event: 'message', callback: (message?: any) => any): void;
    /**
     * 移除通道消息完毕事件监听器
     */
    removeListener(event: 'end', callback: (message?: any) => any): void;
    /**
     * 移除通道关闭事件监听器
     */
    removeListener(event: 'close', callback: () => any): void;
    /**
     * 移除通道错误事件监听器
     */
    removeListener(event: 'error', callback: (err?: any) => any): void;
    /**
     * 移除指定事件的所有监听器
     */
    removeAllListener(event: 'open' | 'message' | 'end' | 'close' | 'error'): void;
  }

  interface UniCloud {
    /**
     * 云函数请求中的中间状态通知通道类
     *
     * 文档：[https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#create-sse-channel](https://uniapp.dcloud.net.cn/uniCloud/sse-channel.html#create-sse-channel)
     */
    SSEChannel: SSEChannel;
  }
}
