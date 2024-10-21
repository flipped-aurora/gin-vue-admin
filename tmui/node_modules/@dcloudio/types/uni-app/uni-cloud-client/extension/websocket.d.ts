declare namespace UniCloudNamespace {
  interface ConnectWebSocketOptions {
    /**
     * WebSocket云函数/云对象名称
     */
    name: string | string.CloudFunctionString;
    /**
     * 建立连接时需要传递的参数, 仅在 connection 事件中接收到
     */
    query?: Record<string, string>;
  }

  interface UniCloud {
    /**
     * 快速连接 WebSocket 服务
     *
     * 文档: [uniCloud.connectWebSocket](https://doc.dcloud.net.cn/uniCloud/websocket.html#unicloud-connectwebsocket)
     */
    connectWebSocket(options: ConnectWebSocketOptions): UniApp.SocketTask;
  }
}
