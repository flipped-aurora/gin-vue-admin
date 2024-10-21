declare namespace UniNamespace {
  interface WifiError {
    /**
     * 错误信息
     *
     * | 错误码 | 错误信息 | 说明 |
     * | - | - | - |
     * | 0 | ok | 正常 |
     * | 12000 | not init | 未先调用 `startWifi` 接口 |
     * | 12001 | system not support | 当前系统不支持相关能力 |
     * | 12002 | password error Wi-Fi | 密码错误 |
     * | 12003 | connection timeout | 连接超时, 仅 Android 支持 |
     * | 12004 | duplicate request | 重复连接 Wi-Fi |
     * | 12005 | wifi not turned on | Android 特有，未打开 Wi-Fi 开关 |
     * | 12006 | gps not turned on | Android 特有，未打开 GPS 定位开关 |
     * | 12007 | user denied | 用户拒绝授权链接 Wi-Fi |
     * | 12008 | invalid SSID | 无效 SSID |
     * | 12009 | system config err | 系统运营商配置拒绝连接 Wi-Fi |
     * | 12010 | system internal error | 系统其他错误，需要在 errmsg 打印具体的错误原因 |
     * | 12011 | weapp in background | 应用在后台无法配置 Wi-Fi |
     * | 12013 | wifi config may be expired | 系统保存的 Wi-Fi 配置过期，建议忘记 Wi-Fi 后重试，仅 Android 支持 |
     * | 12014 | invalid WEP / WPA password | iOS 特有，无效的 WEP / WPA 密码 |
     */
    errMsg: string;
    /**
     * 错误码
     *
     * | 错误码 | 错误信息 | 说明 |
     * | - | - | - |
     * | 0 | ok | 正常 |
     * | 12000 | not init | 未先调用 `startWifi` 接口 |
     * | 12001 | system not support | 当前系统不支持相关能力 |
     * | 12002 | password error Wi-Fi | 密码错误 |
     * | 12003 | connection timeout | 连接超时, 仅 Android 支持 |
     * | 12004 | duplicate request | 重复连接 Wi-Fi |
     * | 12005 | wifi not turned on | Android 特有，未打开 Wi-Fi 开关 |
     * | 12006 | gps not turned on | Android 特有，未打开 GPS 定位开关 |
     * | 12007 | user denied | 用户拒绝授权链接 Wi-Fi |
     * | 12008 | invalid SSID | 无效 SSID |
     * | 12009 | system config err | 系统运营商配置拒绝连接 Wi-Fi |
     * | 12010 | system internal error | 系统其他错误，需要在 errmsg 打印具体的错误原因 |
     * | 12011 | weapp in background | 应用在后台无法配置 Wi-Fi |
     * | 12013 | wifi config may be expired | 系统保存的 Wi-Fi 配置过期，建议忘记 Wi-Fi 后重试，仅 Android 支持 |
     * | 12014 | invalid WEP / WPA password | iOS 特有，无效的 WEP / WPA 密码 |
     */
    errCode: number;
  }

  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  type StartWifiCompleteCallback = (res: WifiError) => void;
  /**
   * 接口调用失败的回调函数
   */
  type StartWifiFailCallback = (res: WifiError) => void;
  /**
   * 接口调用成功的回调函数
   */
  type StartWifiSuccessCallback = (res: WifiError) => void;

  interface StartWifiOption {
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: StartWifiCompleteCallback;
    /**
     * 接口调用失败的回调函数
     */
    fail?: StartWifiFailCallback;
    /**
     * 接口调用成功的回调函数
     */
    success?: StartWifiSuccessCallback;
  }

  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  type StopWifiCompleteCallback = (res: WifiError) => void;
  /**
   * 接口调用失败的回调函数
   */
  type StopWifiFailCallback = (res: WifiError) => void;
  /**
   * 接口调用成功的回调函数
   */
  type StopWifiSuccessCallback = (res: WifiError) => void;

  interface StopWifiOption {
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: StopWifiCompleteCallback;
    /**
     * 接口调用失败的回调函数
     */
    fail?: StopWifiFailCallback;
    /**
     * 接口调用成功的回调函数
     */
    success?: StopWifiSuccessCallback;
  }

  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  type ConnectWifiCompleteCallback = (res: WifiError) => void;
  /**
   * 接口调用失败的回调函数
   */
  type ConnectWifiFailCallback = (res: WifiError) => void;
  /**
   * 接口调用成功的回调函数
   */
  type ConnectWifiSuccessCallback = (res: WifiError) => void;

  interface ConnectWifiOption {
    /**
     * Wi-Fi 设备 SSID
     */
    SSID: string;
    /**
     * Wi-Fi 设备 BSSID
     */
    BSSID?: string;
    /**
     * Wi-Fi 设备密码
     */
    password: string;
    /**
     * 需要基础库： `2.12.0`
     *
     * 跳转到系统设置页进行连接
     */
    maunal?: boolean;
    /**
     * 需要基础库： `2.22.0`
     *
     * 是否需要返回部分 Wi-Fi 信息，仅安卓生效
     */
    partialInfo?: boolean;
    /**
     * 接口调用成功的回调函数
     */
    success?: ConnectWifiSuccessCallback;
    /**
     * 接口调用失败的回调函数
     */
    fail?: ConnectWifiFailCallback;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: ConnectWifiCompleteCallback;
  }

  interface WifiInfo {
    /**
     * Wi-Fi 的 SSID
     */
    SSID: string;
    /**
     * Wi-Fi 的 BSSID
     */
    BSSID: string;
    /**
     * Wi-Fi 是否安全
     */
    secure: boolean;
    /**
     * Wi-Fi 信号强度, 安卓取值 0 ～ 100 ，iOS 取值 0 ～ 1 ，值越大强度越大
     */
    signalStrength: number;
    /**
     * Wi-Fi 频段单位 MHz
     */
    frequency: number;
  }
  interface OnWifiConnectedListenerResult {
    /**
     * Wi-Fi 信息
     *
     * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo](https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo)
     */
    wifi: WifiInfo;
  }
  /**
   * 连接上 Wi-Fi 的事件的监听函数
   */
  type OnWifiConnectedCallback = (
    result: OnWifiConnectedListenerResult
  ) => void;
  /**
   * onWifiConnected 传入的监听函数。不传此参数则移除所有监听函数。
   */
  type OffWifiConnectedCallback = (
    result: OnWifiConnectedListenerResult
  ) => void;

  interface OnWifiConnectedWithPartialInfoListenerResult {
    /**
     *
     * 只包含 SSID 属性的 WifiInfo 对象
     *
     * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo](https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo)
     */
    wifi: WifiInfo;
  }
  /**
   * 连接上 Wi-Fi 的事件的监听函数
   */
  type OnWifiConnectedWithPartialInfoCallback = (
    result: OnWifiConnectedWithPartialInfoListenerResult
  ) => void;

  interface GetConnectedWifiSuccessCallbackResult {
    /**
     * Wi-Fi 信息
     *
     * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo](https://uniapp.dcloud.net.cn/api/system/wifi.html#WifiInfo)
     */
    wifi: WifiInfo;
    errMsg: string;
  }

  /**
   * 接口调用成功的回调函数
   */
  type GetConnectedWifiSuccessCallback = (
    result: GetConnectedWifiSuccessCallbackResult
  ) => void;
  /**
   * 接口调用失败的回调函数
   */
  type GetConnectedWifiFailCallback = (res: WifiError) => void;
  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  type GetConnectedWifiCompleteCallback = (res: WifiError) => void;
  interface GetConnectedWifiOption {
    /**
     * 是否需要返回部分 Wi-Fi 信息
     */
    partialInfo?: boolean;
    /**
     * 接口调用成功的回调函数
     */
    success?: GetConnectedWifiSuccessCallback;
    /**
     * 接口调用失败的回调函数
     */
    fail?: GetConnectedWifiFailCallback;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: GetConnectedWifiCompleteCallback;
  }

  /**
   * 接口调用成功的回调函数
   */
  type GetWifiListSuccessCallback = (res: WifiError) => void;
  /**
   * 接口调用失败的回调函数
   */
  type GetWifiListFailCallback = (res: WifiError) => void;
  /**
   * 接口调用结束的回调函数（调用成功、失败都会执行）
   */
  type GetWifiListCompleteCallback = (res: WifiError) => void;
  interface GetWifiListOption {
    /**
     * 接口调用成功的回调函数
     */
    success?: GetWifiListSuccessCallback;
    /**
     * 接口调用失败的回调函数
     */
    fail?: GetWifiListFailCallback;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: GetWifiListCompleteCallback;
  }

  interface OnGetWifiListListenerResult {
    /**
     * Wi-Fi 列表数据
     */
    wifiList: WifiInfo[];
  }
  /**
   * 获取到 Wi-Fi 列表数据事件的监听函数
   */
  type OnGetWifiListCallback = (result: OnGetWifiListListenerResult) => void;
  /**
   * onGetWifiList 传入的监听函数。不传此参数则移除所有监听函数。
   */
  type OffGetWifiListCallback = (result: OnGetWifiListListenerResult) => void;

  /**
   * onWifiConnectedWithPartialInfo 传入的监听函数。不传此参数则移除所有监听函数。
   */
  type OffWifiConnectedWithPartialInfoCallback = (
    result: OnWifiConnectedWithPartialInfoListenerResult
  ) => void;
}

interface Uni {
  /**
   * 初始化 Wi-Fi 模块
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#startwifi](https://uniapp.dcloud.net.cn/api/system/wifi.html#startwifi)
   */
  startWifi(option?: UniNamespace.StartWifiOption): void;

  /**
   * 关闭 Wi-Fi 模块
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#stopwifi](https://uniapp.dcloud.net.cn/api/system/wifi.html#stopwifi)
   */
  stopWifi(option?: UniNamespace.StopWifiOption): void;

  /**
   * 连接 Wi-Fi。若已知 Wi-Fi 信息，可以直接利用该接口连接。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#connectwifi](https://uniapp.dcloud.net.cn/api/system/wifi.html#connectwifi)
   */
  connectWifi(option: UniNamespace.ConnectWifiOption): void;

  /**
   * 监听连接上 Wi-Fi 的事件。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#onwificonnected](https://uniapp.dcloud.net.cn/api/system/wifi.html#onwificonnected)
   */
  onWifiConnected(listener: UniNamespace.OnWifiConnectedCallback): void;

  /**
   * 移除连接上 Wi-Fi 的事件的监听函数
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#offwificonnected](https://uniapp.dcloud.net.cn/api/system/wifi.html#offwificonnected)
   */
  offWifiConnected(listener?: UniNamespace.OffWifiConnectedCallback): void;

  /**
   * 获取已连接的 Wi-Fi 信息。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#getconnectedwifi](https://uniapp.dcloud.net.cn/api/system/wifi.html#getconnectedwifi)
   */
  getConnectedWifi(option: UniNamespace.GetConnectedWifiOption): void;

  /**
   * 请求获取 Wi-Fi 列表。wifiList 数据会在 onGetWifiList 注册的回调中返回。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#getwifilist](https://uniapp.dcloud.net.cn/api/system/wifi.html#getwifilist)
   */
  getWifiList(option?: UniNamespace.GetWifiListOption): void;

  /**
   * 监听获取到 Wi-Fi 列表数据事件。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#ongetwifilist](https://uniapp.dcloud.net.cn/api/system/wifi.html#ongetwifilist)
   */
  onGetWifiList(listener: UniNamespace.OnGetWifiListCallback): void;

  /**
   * 移除获取到 Wi-Fi 列表数据事件的监听函数
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#offgetwifilist](https://uniapp.dcloud.net.cn/api/system/wifi.html#offgetwifilist)
   */
   offGetWifiList(listener?: UniNamespace.OffGetWifiListCallback): void;

  /**
   * 监听连接上 Wi-Fi 的事件。
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#onwificonnectedwithpartialinfo](https://uniapp.dcloud.net.cn/api/system/wifi.html#onwificonnectedwithpartialinfo)
   */
  onWifiConnectedWithPartialInfo(listener: UniNamespace.OnWifiConnectedWithPartialInfoCallback): void;

  /**
   * 移除连接上 Wi-Fi 的事件的监听函数
   *
   * 文档: [https://uniapp.dcloud.net.cn/api/system/wifi.html#offwificonnectedwithpartialinfo](https://uniapp.dcloud.net.cn/api/system/wifi.html#offwificonnectedwithpartialinfo)
   */
  offWifiConnectedWithPartialInfo(listener?: UniNamespace.OffWifiConnectedWithPartialInfoCallback): void;
}
