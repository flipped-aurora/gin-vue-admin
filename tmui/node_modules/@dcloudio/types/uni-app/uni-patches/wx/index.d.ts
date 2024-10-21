declare namespace UniNamespace {
    /**
     * 需要基础库： `2.10.4`
     *
     * 网络请求过程中一些调试信息，[查看详细说明](https://developers.weixin.qq.com/miniprogram/dev/framework/performance/network.html)
     */
    interface RequestProfile {
        /**
         * SSL建立完成的时间,如果不是安全连接,则值为 0
         */
        SSLconnectionEnd: number;
        /**
         * SSL建立连接的时间,如果不是安全连接,则值为 0
         */
        SSLconnectionStart: number;
        /**
         * HTTP（TCP） 完成建立连接的时间（完成握手），如果是持久连接，则与 fetchStart 值相等。注意如果在传输层发生了错误且重新建立连接，则这里显示的是新建立的连接完成的时间。注意这里握手结束，包括安全连接建立完成、SOCKS 授权通过
         */
        connectEnd: number;
        /**
         * HTTP（TCP） 开始建立连接的时间，如果是持久连接，则与 fetchStart 值相等。注意如果在传输层发生了错误且重新建立连接，则这里显示的是新建立的连接开始的时间
         */
        connectStart: number;
        /**
         * DNS 域名查询完成的时间，如果使用了本地缓存（即无 DNS 查询）或持久连接，则与 fetchStart 值相等
         */
        domainLookupEnd: number;
        /**
         * DNS 域名查询开始的时间，如果使用了本地缓存（即无 DNS 查询）或持久连接，则与 fetchStart 值相等
         */
        domainLookupStart: number;
        /**
         * 评估当前网络下载的kbps
         */
        downstreamThroughputKbpsEstimate: number;
        /**
         * 评估的网络状态 unknown, offline, slow 2g, 2g, 3g, 4g, last/0, 1, 2, 3, 4, 5, 6
         */
        estimate_nettype: number;
        /**
         * 组件准备好使用 HTTP 请求抓取资源的时间，这发生在检查本地缓存之前
         */
        fetchStart: number;
        /**
         * 协议层根据多个请求评估当前网络的 rtt（仅供参考）
         */
        httpRttEstimate: number;
        /**
         * 当前请求的IP
         */
        peerIP: string;
        /**
         * 当前请求的端口
         */
        port: number;
        /**
         * 使用协议类型，有效值：http1.1, h2, quic, unknown
         */
        protocol: string;
        /**
         * 收到字节数
         */
        receivedBytedCount: number;
        /**
         * 最后一个 HTTP 重定向完成时的时间。有跳转且是同域名内部的重定向才算，否则值为 0
         */
        redirectEnd: number;
        /**
         * 第一个 HTTP 重定向发生时的时间。有跳转且是同域名内的重定向才算，否则值为 0
         */
        redirectStart: number;
        /**
         * HTTP请求读取真实文档结束的时间
         */
        requestEnd: number;
        /**
         * HTTP请求读取真实文档开始的时间（完成建立连接），包括从本地读取缓存。连接错误重连时，这里显示的也是新建立连接的时间
         */
        requestStart: number;
        /**
         * HTTP 响应全部接收完成的时间（获取到最后一个字节），包括从本地读取缓存
         */
        responseEnd: number;
        /**
         * HTTP 开始接收响应的时间（获取到第一个字节），包括从本地读取缓存
         */
        responseStart: number;
        /**
         * 当次请求连接过程中实时 rtt
         */
        rtt: number;
        /**
         * 发送的字节数
         */
        sendBytesCount: number;
        /**
         * 是否复用连接
         */
        socketReused: boolean;
        /**
         * 当前网络的实际下载kbps
         */
        throughputKbps: number;
        /**
         * 传输层根据多个请求评估的当前网络的 rtt（仅供参考）
         */
        transportRttEstimate: number;
    }

    interface DownloadSuccessData {
        /** 用户文件路径 (本地路径)。传入 filePath 时会返回，跟传入的 filePath 一致 */
        filePath?: string;
        /**
         * 需要基础库： `2.10.4`
         *
         * 网络请求过程中一些调试信息，[查看详细说明](https://developers.weixin.qq.com/miniprogram/dev/framework/performance/network.html)
         */
        profile?: RequestProfile;
        errMsg?: string;
    }

    interface DownloadFileOption {
        /**
         * 需要基础库： `1.8.0`
         *
         * 指定文件下载后存储的路径 (本地路径)
         */
        filePath?: string;
    }

    interface RequestSuccessCallbackResult {
        /**
         * 需要基础库： `2.10.4`
         *
         * 网络请求过程中一些调试信息，[查看详细说明](https://developers.weixin.qq.com/miniprogram/dev/framework/performance/network.html)
         */
        profile?: RequestProfile;
        errMsg?: string;
    }

    interface RewardedVideoAdOptions {
        /**
         * 需要基础库： `2.8.0`
         *
         * 是否启用多例模式，默认为false
         */
        multiton?: boolean;
    }

    interface ConnectSocketOption {
        /**
         * 需要基础库： `2.8.0`
         *
         * 是否开启压缩扩展
         */
        perMessageDeflate?: boolean;
        /**
         * 需要基础库： `2.4.0`
         *
         * 建立 TCP 连接的时候的 TCP_NODELAY 设置
         */
        tcpNoDelay?: boolean;
        /**
         * 需要基础库： `2.10.0`
         *
         * 超时时间，单位为毫秒
         */
        timeout?: number;
    }

    interface UploadFileSuccessCallbackResult {
        errMsg?: string;
    }

    interface CanvasToTempFilePathRes {
        errMsg?: string;
    }

    interface CanvasToTempFilePathOptions {
        /** 画布标识，传入 [canvas](https://developers.weixin.qq.com/miniprogram/dev/component/canvas.html) 组件实例 （canvas type="2d" 时使用该属性）。 */
        canvas?: any;
    }

    interface ChooseAddressRes {
        /** 新选择器详细收货地址信息 */
        detailInfoNew?: string;
        /** 国标收货地址第四级地址 */
        streetName?: string;
    }

    interface ChooseImageSuccessCallbackResult {
        errMsg?: string;
    }

    interface ChooseLocationSuccess {
        errMsg?: string;
    }

    interface ChooseVideoSuccess {
        errMsg?: string;
    }

    interface CompressImageSuccessResult {
        errMsg?: string;
    }

    interface CompressImageOptions {
        /**
         * 需要基础库： `2.26.0`
         *
         * 压缩后图片的高度，单位为px，若不填写则默认以compressedWidth为准等比缩放
         */
        compressHeight?: number;
    }

    interface CompressVideoSuccessData {
        errMsg?: string;
    }

    interface GetImageInfoSuccessData {
        errMsg?: string;
    }

    interface GetLocationSuccess {
        errMsg?: string;
    }

    interface GetNetworkTypeSuccess {
        /**
         * 需要基础库： `2.22.1`
         *
         * 设备是否使用了网络代理
         */
        hasSystemProxy?: boolean;
        /** 信号强弱，单位 dbm */
        signalStrength?: number;
        errMsg?: string;
    }

    interface GetSelectedTextRangeSuccessCallbackResult {
        errMsg?: string;
    }

    interface GetSettingSuccessResult {
        /**
         * [AuthSetting](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/setting/AuthSetting.html)
         *
         * 在插件中调用时，当前宿主小程序的用户授权结果
         */
        miniprogramAuthSetting?: AuthSetting;
        errMsg?: string;
    }

    interface GetStorageSuccess {
        errMsg?: string;
    }

    interface GetStorageOptions {
        /**
         * 需要基础库： `2.21.3`
         *
         * 是否开启加密存储。只有异步的 getStorage 接口支持开启加密存储。开启后，将会对 data 使用 AES128 解密，接口回调耗时将会增加。若开启加密存储，setStorage 和 getStorage 需要同时声明 encrypt 的值为 true
         */
        encrypt?: boolean;
    }

    interface GetSystemInfoResult {
        /**
         * 需要基础库： `1.8.0`
         *
         * 设备性能等级（仅 Android）。取值为：-2 或 0（该设备无法运行小游戏），-1（性能未知），>=1（设备性能值，该值越高，设备性能越好，目前最高不到50）
         */
        benchmarkLevel?: number;
        /**
         * 需要基础库： `2.15.0`
         *
         * 是否已打开调试。可通过右上角菜单或 [wx.setEnableDebug](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.setEnableDebug.html) 打开调试。
         */
        enableDebug?: boolean;
        /** `true` 表示模糊定位，`false` 表示精确定位，仅 iOS 支持 */
        locationReducedAccuracy?: boolean;
        /**
         * 需要基础库： `2.19.3`
         *
         * 允许微信使用日历的开关
         */
        phoneCalendarAuthorized?: boolean;
    }

    interface GetUserInfoRes {
        /**
         * 需要基础库： `2.7.0`
         *
         * 敏感数据对应的云 ID，开通[云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/basis/getting-started.html)的小程序才会返回，可通过云调用直接获取开放数据，详细见[云调用直接获取开放数据](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#method-cloud)
         */
        cloudID?: string;
    }

    interface GetUserProfileRes {
        /**
         * 需要基础库： `2.10.4`
         *
         * 敏感数据对应的云 ID，开通[云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/basis/getting-started.html)的小程序才会返回，可通过云调用直接获取开放数据，详细见[云调用直接获取开放数据](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#method-cloud)
         */
        cloudID?: string;
    }

    interface GetVideoInfoSuccessData {
        errMsg?: string;
    }

    interface HideShareMenuOptions {
        /**
         * 需要基础库： `2.11.3`
         *
         * 本接口为 Beta 版本，暂只在 Android 平台支持。需要隐藏的转发按钮名称列表，默认['shareAppMessage', 'shareTimeline']。按钮名称合法值包含 "shareAppMessage"、"shareTimeline" 两种
         */
        menus?: string[];
    }

    interface LoadFontFaceOptions {
        /** 字体作用范围，可选值为 webview / native，默认 webview，设置 native 可在 Canvas 2D 下使用 */
        scopes?: any[];
    }

    interface NavigateToMiniProgramOptions {
        /**
         * 需要基础库： `2.24.0`
         *
         * 不reLaunch目标小程序，直接打开目标跳转的小程序退后台时的页面，需满足以下条件：1. 目标跳转的小程序生命周期未被销毁；2. 且目标当次启动的path、query、apiCategory与上次启动相同。
         */
        noRelaunchIfPathUnchanged?: boolean;
        /**
         * 需要基础库： `2.18.1`
         *
         * 小程序链接，当传递该参数后，可以不传 appId 和 path。链接可以通过【小程序菜单】->【复制链接】获取。
         */
        shortLink?: string;
    }

    interface StopBluetoothDevicesDiscoverySuccess {
        /**
         * 错误码
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | 0 | ok | 正常 |
         * | -1 | already connect | 已连接 |
         * | 10000 | not init | 未初始化蓝牙适配器 |
         * | 10001 | not available | 当前蓝牙适配器不可用 |
         * | 10002 | no device | 没有找到指定设备 |
         * | 10003 | connection fail | 连接失败 |
         * | 10004 | no service | 没有找到指定服务 |
         * | 10005 | no characteristic | 没有找到指定特征 |
         * | 10006 | no connection | 当前连接已断开 |
         * | 10007 | property not support | 当前特征不支持此操作 |
         * | 10008 | system error | 其余所有系统上报的异常 |
         * | 10009 | system not support | Android 系统特有，系统版本低于 4.3 不支持 BLE |
         * | 10012 | operate time out | 连接超时 |
         * | 10013 | invalid_data | 连接 deviceId 为空或者是格式不正确 |
         */
        errCode?: number;
    }

    interface NotifyBLECharacteristicValueChangeOptions {
        /**
         * 需要基础库： `2.4.0`
         *
         * 设置特征订阅类型，有效值有 `notification` 和 `indication`
         */
        type?: string;
    }

    interface OpenBluetoothAdapterOptions {
        /**
         * 需要基础库： `2.10.0`
         *
         * 蓝牙模式，可作为主/从设备，仅 iOS 需要。
         *
         * 可选值：
         * - 'central': 主机模式;
         * - 'peripheral': 从机（外围设备）模式;
         */
        mode?: "central" | "peripheral";
    }

    interface OpenDocumentOptions {
        /**
         * 需要基础库： `2.11.0`
         *
         * 是否显示右上角菜单
         */
        showMenu?: boolean;
    }

    interface OpenVideoEditorSuccessData {
        errMsg?: string;
    }

    interface OpenVideoEditorOptions {
        /**
         * 需要基础库： `2.16.1`
         *
         * 视频裁剪的最大长度
         */
        maxDuration?: string;
        /**
         * 需要基础库： `2.16.1`
         *
         * 视频裁剪的最小长度
         */
        minDuration?: string;
    }

    interface PageScrollToOptions {
        /**
         * 需要基础库： `2.23.1`
         *
         * 偏移距离，需要和 selector 参数搭配使用，可以滚动到 selector 加偏移距离的位置，单位 px
         */
        offsetTop?: number;
    }

    interface PreviewImageOptions {
        /**
         * 需要基础库： `2.13.0`
         *
         * `origin`: 发送完整的referrer; `no-referrer`: 不发送。格式固定为 `https://servicewechat.com/{appid}/{version}/page-frame.html`，其中 {appid} 为小程序的 appid，{version} 为小程序的版本号，版本号为 0 表示为开发版、体验版以及审核版本，版本号为 devtools 表示为开发者工具，其余为正式版本；
         */
        referrerPolicy?: string;
        /**
         * 需要基础库： `2.13.0`
         *
         * 是否显示长按菜单。
         */
        showmenu?: boolean;
    }

    interface SaveImageToPhotosAlbumResult {
        /** 错误信息 */
        errMsg?: string;
    }

    interface ScanCodeSuccessRes {
        /** 原始数据，base64编码 */
        rawData?: string;
        errMsg?: string;
    }

    interface SetStorageOptions {
        /**
         * 需要基础库： `2.21.3`
         *
         * 是否开启加密存储。只有异步的 setStorage 接口支持开启加密存储。开启后，将会对 data 使用 AES128 加密，接口回调耗时将会增加。若开启加密存储，setStorage 和 getStorage 需要同时声明 encrypt 的值为 true。此外，由于加密后的数据会比原始数据膨胀1.4倍，因此开启 encrypt 的情况下，单个 key 允许存储的最大数据长度为 0.7MB，所有数据存储上限为 7.1MB
         */
        encrypt?: boolean;
    }

    interface ShowActionSheetRes {
        errMsg?: string;
    }

    interface ShowModalRes {
        errMsg?: string;
    }

    interface StartBeaconDiscoveryOptions {
        /** 是否校验蓝牙开关，仅在 iOS 下有效。iOS 11 起，控制面板里关掉蓝牙，还是能继续使用 Beacon 服务。 */
        ignoreBluetoothAvailable?: boolean;
    }

    interface StartBluetoothDevicesDiscoveryOptions {
        /**
         * 扫描模式，越高扫描越快，也越耗电。仅安卓微信客户端 7.0.12 及以上支持。
         *
         * 可选值：
         * - 'low': 低;
         * - 'medium': 中;
         * - 'high': 高;
         */
        powerLevel?: "low" | "medium" | "high";
    }

    interface VibrateShortOptions {
        /**
         * 需要基础库： `2.13.0`
         *
         * 震动强度类型，有效值为：heavy、medium、light
         */
        type?: string;
    }

    interface WriteBLECharacteristicValueOptions {
        /**
         * 蓝牙特征值的写模式设置，有两种模式，iOS 优先 write，安卓优先 writeNoResponse 。（基础库 2.22.0 开始支持）
         *
         * 可选值：
         * - 'write': 强制回复写，不支持时报错;
         * - 'writeNoResponse': 强制无回复写，不支持时报错;
         */
        writeType?: "write" | "writeNoResponse";
    }

    /**
     * 小程序帐号信息
     */
    interface MiniProgram {
        /**
         * 小程序 appId
         */
        appId: string;
        /**
         * 需要基础库： `2.10.0`
         *
         * 小程序版本
         *
         * 可选值：
         * - 'develop': 开发版;
         * - 'trial': 体验版;
         * - 'release': 正式版;
         */
        envVersion: "develop" | "trial" | "release";
        /**
         * 需要基础库： `2.10.2`
         *
         * 线上小程序版本号
         */
        version: string;
    }

    /**
     * 插件帐号信息（仅在插件中调用时包含这一项）
     */
    interface Plugin {
        /**
         * 插件 appId
         */
        appId: string;
        /**
         * 插件版本号
         */
        version: string;
    }

    /**
     * 帐号信息
     */
    interface AccountInfo {
        /**
         * 小程序帐号信息
         */
        miniProgram: MiniProgram;
        /**
         * 插件帐号信息（仅在插件中调用时包含这一项）
         */
        plugin: Plugin;
    }

    interface GetBatteryInfoSyncResult {
        /**
         * 是否正在充电中
         */
        isCharging: boolean;
        /**
         * 设备电量，范围 1 - 100
         */
        level: number;
    }

    /**
     * 当前运行环境对于 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的支持情况
     */
    interface SkylineInfo {
        /**
         * 当前运行环境是否支持 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html)
         */
        isSupported: boolean;
        /**
         * 当前运行环境不支持 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的原因，仅在 `isSupported` 为 `false` 时出现
         *
         * 可选值：
         * - 'client not supported': 当前微信客户端不支持 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html)，可以尝试通过升级微信客户端解决;
         * - 'baselib not supported': 当前基础库不支持 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html)，基础库会自动更新到当前客户端所能支持的最新的版本，基础库不支持时也可以尝试通过升级微信客户端解决;
         * - 'a-b test not enabled': 命中了 _We 分析_ 平台上的 AB 实验关闭的情况。详细可以查看 [Skyline 起步 > 配置 We 分析 AB 实验](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/migration/#%E9%85%8D%E7%BD%AE-We-%E5%88%86%E6%9E%90-AB-%E5%AE%9E%E9%AA%8C) 一节;
         * - 'SwitchRender option set to webview': 本地调试的快捷切换入口被设置为了强制使用 Webview. 详情可以查看 [Skyline 起步 > 快捷切换入口](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/migration/#%E5%BF%AB%E6%8D%B7%E5%88%87%E6%8D%A2%E5%85%A5%E5%8F%A3) 一节;
         */
        reason: "client not supported" | "baselib not supported" | "a-b test not enabled" | "SwitchRender option set to webview";
        /**
         * 当前运行环境 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的版本号，形如 `0.9.7`
         */
        version: string;
    }

    /**
     * 额外的缓存处理
     */
    interface ExtraOption {
        /**
         * 需要缓存的 wx api 接口，不传则表示支持缓存的接口全都做缓存处理。返回的如果是缓存数据，开发者可通过 fromCache 标记区分
         *
         * 可选值：
         * - 'wx.login': ;
         * - 'wx.checkSession': ;
         * - 'wx.getSetting': ;
         */
        apiList?: ("wx.login" | "wx.checkSession" | "wx.getSetting")[];
    }

    interface CreateCacheManagerOption {
        /**
         * 额外的缓存处理
         */
        extra?: ExtraOption;
        /**
         * 全局缓存有效时间，单位为毫秒，默认为 7 天，最长不超过 30 天
         */
        maxAge?: number;
        /**
         * 缓存模式
         *
         * 可选值：
         * - 'weakNetwork': 弱网/离线使用缓存返回;
         * - 'always': 总是使用缓存返回;
         * - 'none': 不开启，后续可手动开启/停止使用缓存返回;
         */
        mode?: "weakNetwork" | "always" | "none";
        /**
         * 全局 origin
         */
        origin?: string;
    }

    /**
     * 需要基础库： `2.24.0`
     *
     * 缓存管理器。全局只有唯一实例，一旦被创建出来即表示接入缓存管理器。其有以下几个能力：
     *
     * 1. 在网络通畅时，符合一定规则的用户网络请求（目前只包括普通 wx.request 请求）会被缓存。
     * 2. 在网络通畅时，某些 wx api 调用会被缓存。
     * 3. 进入弱网/离线状态时，会提供事件给用户，用户可以决定是否使用缓存返回。
     * 4. 提供进入和退出弱网/离线状态的事件。
     *
     * > 1. 缓存管理器中涉及的网络请求如无特指，均指普通的 wx.request 异步请求，参数和返回值中均不考虑涉及 ArrayBuffer 或 TypedArray 的情形。
     * > 2. 缓存管理器中的缓存不会占用 storage 空间，但是有大小限制，请勿在非必要的请求上使用缓存。
     */
    interface CacheManager {
        /**
         * 全局缓存有效时间
         */
        maxAge: string;
        /**
         * 当前缓存模式
         *
         * 可选值：
         * - 'weakNetwork': 默认值，弱网/离线使用缓存返回;
         * - 'always': 总是使用缓存返回;
         * - 'none': 不开启，后续可手动开启/停止使用缓存返回;
         */
        mode: "weakNetwork" | "always" | "none";
        /**
         * 全局 origin
         */
        origin: string;
        /**
         * 当前缓存管理器状态
         *
         * 可选值：
         * - 0: 不使用缓存返回;
         * - 1: 使用缓存返回;
         * - 2: 未知;
         */
        state: 0 | 1 | 2;
    }

    interface ReadCompressedFileSyncOption {
        /**
         * 文件压缩类型，目前仅支持 'br'。
         *
         * 可选值：
         * - 'br': brotli压缩文件;
         */
        compressionAlgorithm: "br";
        /**
         * 要读取的文件的路径 (本地用户文件或代码包文件)
         */
        filePath: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AccessCompleteCallback = (res: GeneralCallbackResult) => void;

    interface AccessFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory ${path}': 文件/目录不存在;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type AccessFailCallback = (result: AccessFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AccessSuccessCallback = (res: GeneralCallbackResult) => void;

    interface AccessOption {
        /**
         * 要判断是否存在的文件/目录路径 (本地路径)
         */
        path: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AccessCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AccessFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AccessSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AppendFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface AppendFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 文件不存在;
         * - 'fail illegal operation on a directory, open "${filePath}"': 指定的 filePath 是一个已经存在的目录;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有写权限;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type AppendFileFailCallback = (result: AppendFileFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AppendFileSuccessCallback = (res: GeneralCallbackResult) => void;

    interface AppendFileOption {
        /**
         * 要追加的文本或二进制数据
         */
        data: string | ArrayBuffer;
        /**
         * 要追加内容的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AppendFileCompleteCallback;
        /**
         * 指定写入文件的字符编码
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 接口调用失败的回调函数
         */
        fail?: AppendFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AppendFileSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type FileSystemManagerCloseCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface CloseFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'bad file descriptor': 无效的文件描述符;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type FileSystemManagerCloseFailCallback = (
                result: CloseFailCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type FileSystemManagerCloseSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface FileSystemManagerCloseOption {
        /**
         * 需要被关闭的文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: FileSystemManagerCloseCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: FileSystemManagerCloseFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: FileSystemManagerCloseSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type CopyFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface CopyFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail permission denied, copyFile ${srcPath} -> ${destPath}': 指定目标文件路径没有写权限;
         * - 'fail no such file or directory, copyFile ${srcPath} -> ${destPath}': 源文件不存在，或目标文件路径的上层目录不存在;
         * - 'fail the maximum size of the file storage limit is exceeded': 存储空间不足;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type CopyFileFailCallback = (result: CopyFileFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type CopyFileSuccessCallback = (res: GeneralCallbackResult) => void;

    interface CopyFileOption {
        /**
         * 目标文件路径，支持本地路径
         */
        destPath: string;
        /**
         * 源文件路径，支持本地路径
         */
        srcPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: CopyFileCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: CopyFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: CopyFileSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type FstatCompleteCallback = (res: GeneralCallbackResult) => void;

    interface FstatFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'bad file descriptor': 无效的文件描述符;
         * - 'fail permission denied': 指定的 fd 路径没有读权限;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type FstatFailCallback = (result: FstatFailCallbackResult) => void;

    /**
     * 描述文件状态的对象
     */
    interface Stats {
        /**
         * 文件最近一次被存取或被执行的时间，UNIX 时间戳，对应 POSIX stat.st_atime
         */
        lastAccessedTime: number;
        /**
         * 文件最后一次被修改的时间，UNIX 时间戳，对应 POSIX stat.st_mtime
         */
        lastModifiedTime: number;
        /**
         * 文件的类型和存取的权限，对应 POSIX stat.st_mode
         */
        mode: string;
        /**
         * 文件大小，单位：B，对应 POSIX stat.st_size
         */
        size: number;
    }

    interface FstatSuccessCallbackResult {
        /**
         *
         * Stats 对象，包含了文件的状态信息
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/Stats.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/Stats.html)
         */
        stats: Stats;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type FstatSuccessCallback = (result: FstatSuccessCallbackResult) => void;

    interface FstatOption {
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: FstatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: FstatFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: FstatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type FtruncateCompleteCallback = (res: GeneralCallbackResult) => void;

    interface FtruncateFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'bad file descriptor': 无效的文件描述符;
         * - 'fail permission denied': 指定的 fd 没有写权限;
         * - 'fail the maximum size of the file storage limit is exceeded': 存储空间不足;
         * - 'fail sdcard not mounted': android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type FtruncateFailCallback = (result: FtruncateFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type FtruncateSuccessCallback = (res: GeneralCallbackResult) => void;

    interface FtruncateOption {
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 截断位置，默认0。如果 length 小于文件长度（单位：字节），则只有前面 length 个字节会保留在文件中，其余内容会被删除；如果 length 大于文件长度，则会对其进行扩展，并且扩展部分将填充空字节（'\0'）
         */
        length: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: FtruncateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: FtruncateFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: FtruncateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetFileInfoCompleteCallback = (res: GeneralCallbackResult) => void;

    interface GetFileInfoFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail file not exist': 指定的 filePath 找不到文件;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type GetFileInfoFailCallback = (
                result: GetFileInfoFailCallbackResult
            ) => void;

    interface GetFileInfoSuccessCallbackResult {
        /**
         * 文件大小，以字节为单位
         */
        size: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetFileInfoSuccessCallback = (
                result: GetFileInfoSuccessCallbackResult
            ) => void;

    interface GetFileInfoOption {
        /**
         * 要读取的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetFileInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetFileInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetFileInfoSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetSavedFileListCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetSavedFileListFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 文件数组
     */
    interface FileItem {
        /**
         * 文件保存时的时间戳，从1970/01/01 08:00:00 到当前时间的秒数
         */
        createTime: number;
        /**
         * 文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 本地文件大小，以字节为单位
         */
        size: number;
    }

    interface GetSavedFileListSuccessCallbackResult {
        /**
         * 文件数组
         */
        fileList: FileItem[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetSavedFileListSuccessCallback = (
                result: GetSavedFileListSuccessCallbackResult
            ) => void;

    interface GetSavedFileListOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetSavedFileListCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetSavedFileListFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetSavedFileListSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type MkdirCompleteCallback = (res: GeneralCallbackResult) => void;

    interface MkdirFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory ${dirPath}': 上级目录不存在（该错误仅在 recursive = false 时生效）;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有写权限;
         * - 'fail file already exists ${dirPath}': 有同名文件或目录（该错误仅在 recursive = false 时生效）;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type MkdirFailCallback = (result: MkdirFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type MkdirSuccessCallback = (res: GeneralCallbackResult) => void;

    interface MkdirOption {
        /**
         * 创建的目录路径 (本地路径)
         */
        dirPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: MkdirCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: MkdirFailCallback;
        /**
         * 需要基础库： `2.3.0`
         *
         * 是否在递归创建该目录的上级目录后再创建该目录。如果对应的上级目录已经存在，则不创建该上级目录。如 dirPath 为 a/b/c/d 且 recursive 为 true，将创建 a 目录，再在 a 目录下创建 b 目录，以此类推直至创建 a/b/c 目录下的 d 目录。
         */
        recursive?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: MkdirSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenCompleteCallback = (res: GeneralCallbackResult) => void;

    interface OpenFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory "${filePath}"': 上级目录不存在;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type OpenFailCallback = (result: OpenFailCallbackResult) => void;

    interface OpenSuccessCallbackResult {
        /**
         * 文件描述符
         */
        fd: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type OpenSuccessCallback = (result: OpenSuccessCallbackResult) => void;

    interface OpenOption {
        /**
         * 文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenFailCallback;
        /**
         * 文件系统标志，默认值: 'r'
         *
         * 可选值：
         * - 'a': 打开文件用于追加。 如果文件不存在，则创建该文件;
         * - 'ax': 类似于 'a'，但如果路径存在，则失败;
         * - 'a+': 打开文件用于读取和追加。 如果文件不存在，则创建该文件;
         * - 'ax+': 类似于 'a+'，但如果路径存在，则失败;
         * - 'as': 打开文件用于追加（在同步模式中）。 如果文件不存在，则创建该文件;
         * - 'as+': 打开文件用于读取和追加（在同步模式中）。 如果文件不存在，则创建该文件;
         * - 'r': 打开文件用于读取。 如果文件不存在，则会发生异常;
         * - 'r+': 打开文件用于读取和写入。 如果文件不存在，则会发生异常;
         * - 'w': 打开文件用于写入。 如果文件不存在则创建文件，如果文件存在则截断文件;
         * - 'wx': 类似于 'w'，但如果路径存在，则失败;
         * - 'w+': 打开文件用于读取和写入。 如果文件不存在则创建文件，如果文件存在则截断文件;
         * - 'wx+': 类似于 'w+'，但如果路径存在，则失败;
         */
        flag?: "a" | "ax" | "a+" | "ax+" | "as" | "as+" | "r" | "r+" | "w" | "wx" | "w+" | "wx+";
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ReadCompleteCallback = (res: GeneralCallbackResult) => void;

    interface ReadFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'bad file descriptor': 无效的文件描述符;
         * - 'fail permission denied': 指定的 fd 路径没有读权限;
         * - 'fail the value of "offset" is out of range': 传入的 offset 不合法;
         * - 'fail the value of "length" is out of range': 传入的 length 不合法;
         * - 'fail sdcard not mounted': android sdcard 挂载失败;
         * - 'bad file descriptor': 无效的文件描述符;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type ReadFailCallback = (result: ReadFailCallbackResult) => void;

    interface ReadSuccessCallbackResult {
        /**
         * 被写入的缓存区的对象，即接口入参的 arrayBuffer
         */
        arrayBuffer: ArrayBuffer;
        /**
         * 实际读取的字节数
         */
        bytesRead: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ReadSuccessCallback = (result: ReadSuccessCallbackResult) => void;

    interface ReadOption {
        /**
         * 数据写入的缓冲区，必须是 ArrayBuffer 实例
         */
        arrayBuffer: ArrayBuffer;
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ReadCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ReadFailCallback;
        /**
         * 要从文件中读取的字节数，默认0
         */
        length?: number;
        /**
         * 缓冲区中的写入偏移量，默认0
         */
        offset?: number;
        /**
         * 文件读取的起始位置，如不传或传 null，则会从当前文件指针的位置读取。如果 position 是正整数，则文件指针位置会保持不变并从 position 读取文件。
         */
        position?: number;
        /**
         * 接口调用成功的回调函数
         */
        success?: ReadSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ReadCompressedFileCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface ReadCompressedFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail decompress fail': 指定的 compressionAlgorithm 与文件实际压缩格式不符;
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 所在目录不存在;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有读权限;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type ReadCompressedFileFailCallback = (
                result: ReadCompressedFileFailCallbackResult
            ) => void;

    interface ReadCompressedFileSuccessCallbackResult {
        /**
         * 文件内容
         */
        data: ArrayBuffer;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ReadCompressedFileSuccessCallback = (
                result: ReadCompressedFileSuccessCallbackResult
            ) => void;

    interface ReadCompressedFileOption {
        /**
         * 文件压缩类型，目前仅支持 'br'。
         *
         * 可选值：
         * - 'br': brotli压缩文件;
         */
        compressionAlgorithm: "br";
        /**
         * 要读取的文件的路径 (本地用户文件或代码包文件)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ReadCompressedFileCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ReadCompressedFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ReadCompressedFileSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ReadFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface ReadFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 所在目录不存在;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有读权限;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         * - 'native buffer exceed size limit': 文件大小超出上限（100M）;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type ReadFileFailCallback = (result: ReadFileFailCallbackResult) => void;

    interface ReadFileSuccessCallbackResult {
        /**
         * 文件内容
         */
        data: string | ArrayBuffer;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ReadFileSuccessCallback = (
                result: ReadFileSuccessCallbackResult
            ) => void;

    interface ReadFileOption {
        /**
         * 要读取的文件的路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ReadFileCompleteCallback;
        /**
         * 指定读取文件的字符编码，如果不传 encoding，则以 ArrayBuffer 格式读取文件的二进制内容
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 接口调用失败的回调函数
         */
        fail?: ReadFileFailCallback;
        /**
         * 需要基础库： `2.10.0`
         *
         * 指定文件的长度，如果不指定，则读到文件末尾。有效范围：[1, fileLength]。单位：byte
         */
        length?: number;
        /**
         * 需要基础库： `2.10.0`
         *
         * 从文件指定位置开始读，如果不指定，则从文件头开始读。读取的范围应该是左闭右开区间 [position, position+length)。有效范围：[0, fileLength - 1]。单位：byte
         */
        position?: number;
        /**
         * 接口调用成功的回调函数
         */
        success?: ReadFileSuccessCallback;
    }

    /**
     * 要读取的压缩包内的文件列表（当传入"all" 时表示读取压缩包内所有文件）
     */
    interface EntryItem {
        /**
         * 压缩包内文件路径
         */
        path: string;
        /**
         * 指定读取文件的字符编码，如果不传 encoding，则以 ArrayBuffer 格式读取文件的二进制内容
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 指定文件的长度，如果不指定，则读到文件末尾。有效范围：[1, fileLength]。单位：byte
         */
        length?: number;
        /**
         * 从文件指定位置开始读，如果不指定，则从文件头开始读。读取的范围应该是左闭右开区间 [position, position+length)。有效范围：[0, fileLength - 1]。单位：byte
         */
        position?: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ReadZipEntryCompleteCallback = (res: GeneralCallbackResult) => void;

    interface ReadZipEntryFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 所在目录不存在;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有读权限;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type ReadZipEntryFailCallback = (
                result: ReadZipEntryFailCallbackResult
            ) => void;

    /**
     * 文件读取结果。res.entries 是一个对象，key是文件路径，value是一个对象 FileItem ，表示该文件的读取结果。每个 FileItem 包含 data （文件内容） 和 errMsg （错误信息） 属性。
     */
    interface EntriesResult {
    }

    interface ReadZipEntrySuccessCallbackResult {
        /**
         * 文件读取结果。res.entries 是一个对象，key是文件路径，value是一个对象 FileItem ，表示该文件的读取结果。每个 FileItem 包含 data （文件内容） 和 errMsg （错误信息） 属性。
         */
        entries: EntriesResult;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ReadZipEntrySuccessCallback = (
                result: ReadZipEntrySuccessCallbackResult
            ) => void;

    interface ReadZipEntryOption {
        /**
         * 要读取的压缩包内的文件列表（当传入"all" 时表示读取压缩包内所有文件）
         */
        entries: EntryItem[] | "all";
        /**
         * 要读取的压缩包的路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ReadZipEntryCompleteCallback;
        /**
         * 统一指定读取文件的字符编码，只在 entries 值为"all"时有效。如果 entries 值为"all"且不传 encoding，则以 ArrayBuffer 格式读取文件的二进制内容
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 接口调用失败的回调函数
         */
        fail?: ReadZipEntryFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ReadZipEntrySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ReaddirCompleteCallback = (res: GeneralCallbackResult) => void;

    interface ReaddirFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory ${dirPath}': 目录不存在;
         * - 'fail not a directory ${dirPath}': dirPath 不是目录;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有读权限;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type ReaddirFailCallback = (result: ReaddirFailCallbackResult) => void;

    interface ReaddirSuccessCallbackResult {
        /**
         * 指定目录下的文件名数组。
         */
        files: string[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ReaddirSuccessCallback = (result: ReaddirSuccessCallbackResult) => void;

    interface ReaddirOption {
        /**
         * 要读取的目录路径 (本地路径)
         */
        dirPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ReaddirCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ReaddirFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ReaddirSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RemoveSavedFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface RemoveSavedFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail file not exist': 指定的 tempFilePath 找不到文件;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type RemoveSavedFileFailCallback = (
                result: RemoveSavedFileFailCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RemoveSavedFileSuccessCallback = (res: GeneralCallbackResult) => void;

    interface RemoveSavedFileOption {
        /**
         * 需要删除的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RemoveSavedFileCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RemoveSavedFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RemoveSavedFileSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RenameCompleteCallback = (res: GeneralCallbackResult) => void;

    interface RenameFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail permission denied, rename ${oldPath} -> ${newPath}': 指定源文件或目标文件没有写权限;
         * - 'fail no such file or directory, rename ${oldPath} -> ${newPath}': 源文件不存在，或目标文件路径的上层目录不存在;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type RenameFailCallback = (result: RenameFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RenameSuccessCallback = (res: GeneralCallbackResult) => void;

    interface RenameOption {
        /**
         * 新文件路径，支持本地路径
         */
        newPath: string;
        /**
         * 源文件路径，支持本地路径
         */
        oldPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RenameCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RenameFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RenameSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RmdirCompleteCallback = (res: GeneralCallbackResult) => void;

    interface RmdirFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory ${dirPath}': 目录不存在;
         * - 'fail directory not empty': 目录不为空;
         * - 'fail permission denied, open ${dirPath}': 指定的 dirPath 路径没有写权限;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type RmdirFailCallback = (result: RmdirFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RmdirSuccessCallback = (res: GeneralCallbackResult) => void;

    interface RmdirOption {
        /**
         * 要删除的目录路径 (本地路径)
         */
        dirPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RmdirCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RmdirFailCallback;
        /**
         * 需要基础库： `2.3.0`
         *
         * 是否递归删除目录。如果为 true，则删除该目录和该目录下的所有子目录以及文件。
         */
        recursive?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: RmdirSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SaveFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface SaveFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail tempFilePath file not exist': 指定的 tempFilePath 找不到文件;
         * - 'fail permission denied, open "${filePath}"': 指定的 filePath 路径没有写权限;
         * - 'fail no such file or directory "${dirPath}"': 上级目录不存在;
         * - 'fail the maximum size of the file storage limit is exceeded': 存储空间不足;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type SaveFileFailCallback = (result: SaveFileFailCallbackResult) => void;

    interface SaveFileSuccessCallbackResult {
        /**
         * 存储后的文件路径 (本地路径)
         */
        savedFilePath: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type SaveFileSuccessCallback = (
                result: SaveFileSuccessCallbackResult
            ) => void;

    interface SaveFileOption {
        /**
         * 临时存储文件路径 (本地路径)
         */
        tempFilePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SaveFileCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SaveFileFailCallback;
        /**
         * 要存储的文件路径 (本地路径)
         */
        filePath?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: SaveFileSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StatCompleteCallback = (res: GeneralCallbackResult) => void;

    interface StatFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail permission denied, open ${path}': 指定的 path 路径没有读权限;
         * - 'fail no such file or directory ${path}': 文件不存在;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type StatFailCallback = (result: StatFailCallbackResult) => void;

    interface StatSuccessCallbackResult {
        /**
         * [Stats](https://developers.weixin.qq.com/miniprogram/dev/api/file/Stats.html)|Array.&lt;[Stats](https://developers.weixin.qq.com/miniprogram/dev/api/file/Stats.html)&gt;
         *
         * 当 recursive 为 false 时，res.stats 是一个 Stats 对象。当 recursive 为 true 且 path 是一个目录的路径时，res.stats 是一个 Array，数组的每一项是一个对象，每个对象包含 path 和 stats。
         */
        stats: Stats | Stats[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type StatSuccessCallback = (result: StatSuccessCallbackResult) => void;

    interface StatOption {
        /**
         * 文件/目录路径 (本地路径)
         */
        path: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StatFailCallback;
        /**
         * 需要基础库： `2.3.0`
         *
         * 是否递归获取目录下的每个文件的 Stats 信息
         */
        recursive?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: StatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type TruncateCompleteCallback = (res: GeneralCallbackResult) => void;

    interface TruncateFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 所在目录不存在;
         * - 'fail illegal operation on a directory, open "${filePath}"': 指定的 filePath 是一个已经存在的目录;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有写权限;
         * - 'fail the maximum size of the file storage limit is exceeded': 存储空间不足;
         * - 'fail sdcard not mounted': android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type TruncateFailCallback = (result: TruncateFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type TruncateSuccessCallback = (res: GeneralCallbackResult) => void;

    interface TruncateOption {
        /**
         * 要截断的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: TruncateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: TruncateFailCallback;
        /**
         * 截断位置，默认0。如果 length 小于文件长度（字节），则只有前面 length 个字节会保留在文件中，其余内容会被删除；如果 length 大于文件长度，则会对其进行扩展，并且扩展部分将填充空字节（'\0'）
         */
        length?: number;
        /**
         * 接口调用成功的回调函数
         */
        success?: TruncateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type UnlinkCompleteCallback = (res: GeneralCallbackResult) => void;

    interface UnlinkFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail permission denied, open ${path}': 指定的 path 路径没有读权限;
         * - 'fail no such file or directory ${path}': 文件不存在;
         * - 'fail operation not permitted, unlink ${filePath}': 传入的 filePath 是一个目录;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type UnlinkFailCallback = (result: UnlinkFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type UnlinkSuccessCallback = (res: GeneralCallbackResult) => void;

    interface UnlinkOption {
        /**
         * 要删除的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: UnlinkCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: UnlinkFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: UnlinkSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type UnzipCompleteCallback = (res: GeneralCallbackResult) => void;

    interface UnzipFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail permission denied, unzip ${zipFilePath} -> ${destPath}': 指定目标文件路径没有写权限;
         * - 'fail no such file or directory, unzip ${zipFilePath} -> "${destPath}': 源文件不存在，或目标文件路径的上层目录不存在;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type UnzipFailCallback = (result: UnzipFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type UnzipSuccessCallback = (res: GeneralCallbackResult) => void;

    interface UnzipOption {
        /**
         * 目标目录路径, 支持本地路径
         */
        targetPath: string;
        /**
         * 源文件路径，支持本地路径, 只可以是 zip 压缩文件
         */
        zipFilePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: UnzipCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: UnzipFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: UnzipSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type WriteCompleteCallback = (res: GeneralCallbackResult) => void;

    interface WriteFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'bad file descriptor': 无效的文件描述符;
         * - 'fail permission denied': 指定的 fd 路径没有写权限;
         * - 'fail sdcard not mounted': android sdcard 挂载失败;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type WriteFailCallback = (result: WriteFailCallbackResult) => void;

    interface WriteSuccessCallbackResult {
        /**
         * 实际被写入到文件中的字节数（注意，被写入的字节数不一定与被写入的字符串字符数相同）
         */
        bytesWritten: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type WriteSuccessCallback = (result: WriteSuccessCallbackResult) => void;

    interface WriteOption {
        /**
         * 写入的内容，类型为 String 或 ArrayBuffer
         */
        data: string | ArrayBuffer;
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: WriteCompleteCallback;
        /**
         * 只在 data 类型是 String 时有效，指定写入文件的字符编码，默认为 utf8
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 接口调用失败的回调函数
         */
        fail?: WriteFailCallback;
        /**
         * 只在 data 类型是 ArrayBuffer 时有效，指定要写入的字节数，默认为 arrayBuffer 从0开始偏移 offset 个字节后剩余的字节数
         */
        length?: number;
        /**
         * 只在 data 类型是 ArrayBuffer 时有效，决定 arrayBuffe 中要被写入的部位，即 arrayBuffer 中的索引，默认0
         */
        offset?: number;
        /**
         * 指定文件开头的偏移量，即数据要被写入的位置。当 position 不传或者传入非 Number 类型的值时，数据会被写入当前指针所在位置。
         */
        position?: number;
        /**
         * 接口调用成功的回调函数
         */
        success?: WriteSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type WriteFileCompleteCallback = (res: GeneralCallbackResult) => void;

    interface WriteFileFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'fail no such file or directory, open ${filePath}': 指定的 filePath 所在目录不存在;
         * - 'fail permission denied, open ${dirPath}': 指定的 filePath 路径没有写权限;
         * - 'fail the maximum size of the file storage limit is exceeded': 存储空间不足;
         * - 'fail sdcard not mounted': Android sdcard 挂载失败;
         * - 'fail base64 encode error': base64 格式错误;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type WriteFileFailCallback = (result: WriteFileFailCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type WriteFileSuccessCallback = (res: GeneralCallbackResult) => void;

    interface WriteFileOption {
        /**
         * 要写入的文本或二进制数据
         */
        data: string | ArrayBuffer;
        /**
         * 要写入的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: WriteFileCompleteCallback;
        /**
         * 指定写入文件的字符编码
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': （注意，选择 base64 编码，data 只需要传 base64 内容本身，不要传 Data URI 前缀，否则会报 fail base64 encode error 错误。例如，传 aGVsbG8= 而不是传 data:image/png;base64,aGVsbG8= ）;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 接口调用失败的回调函数
         */
        fail?: WriteFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: WriteFileSuccessCallback;
    }

    interface ReadSyncOption {
        /**
         * 数据写入的缓冲区，必须是 ArrayBuffer 实例
         */
        arrayBuffer: ArrayBuffer;
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 要从文件中读取的字节数，默认0
         */
        length?: number;
        /**
         * 缓冲区中的写入偏移量，默认0
         */
        offset?: number;
        /**
         * 文件读取的起始位置，如不传或传 null，则会从当前文件指针的位置读取。如果 position 是正整数，则文件指针位置会保持不变并从 position 读取文件。
         */
        position?: number;
    }

    /**
     * 文件读取结果。 通过 [FileSystemManager.readSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readSync.html) 接口返回
     */
    interface ReadResult {
        /**
         * 被写入的缓存区的对象，即接口入参的 arrayBuffer
         */
        arrayBuffer: ArrayBuffer;
        /**
         * 实际读取的字节数
         */
        bytesRead: number;
    }

    interface FstatSyncOption {
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
    }

    interface WriteSyncOption {
        /**
         * 写入的内容，类型为 String 或 ArrayBuffer
         */
        data: string | ArrayBuffer;
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 只在 data 类型是 String 时有效，指定写入文件的字符编码，默认为 utf8
         *
         * 可选值：
         * - 'ascii': ;
         * - 'base64': ;
         * - 'binary': ;
         * - 'hex': ;
         * - 'ucs2': 以小端序读取;
         * - 'ucs-2': 以小端序读取;
         * - 'utf16le': 以小端序读取;
         * - 'utf-16le': 以小端序读取;
         * - 'utf-8': ;
         * - 'utf8': ;
         * - 'latin1': ;
         */
        encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1";
        /**
         * 只在 data 类型是 ArrayBuffer 时有效，指定要写入的字节数，默认为 arrayBuffer 从0开始偏移 offset 个字节后剩余的字节数
         */
        length?: number;
        /**
         * 只在 data 类型是 ArrayBuffer 时有效，决定 arrayBuffe 中要被写入的部位，即 arrayBuffer 中的索引，默认0
         */
        offset?: number;
        /**
         * 指定文件开头的偏移量，即数据要被写入的位置。当 position 不传或者传入非 Number 类型的值时，数据会被写入当前指针所在位置。
         */
        position?: number;
    }

    /**
     * 文件写入结果。 通过 [FileSystemManager.writeSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeSync.html) 接口返回
     */
    interface WriteResult {
        /**
         * 实际被写入到文件中的字节数（注意，被写入的字节数不一定与被写入的字符串字符数相同）
         */
        bytesWritten: number;
    }

    interface OpenSyncOption {
        /**
         * 文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 文件系统标志，默认值: 'r'
         *
         * 可选值：
         * - 'a': 打开文件用于追加。 如果文件不存在，则创建该文件;
         * - 'ax': 类似于 'a'，但如果路径存在，则失败;
         * - 'a+': 打开文件用于读取和追加。 如果文件不存在，则创建该文件;
         * - 'ax+': 类似于 'a+'，但如果路径存在，则失败;
         * - 'as': 打开文件用于追加（在同步模式中）。 如果文件不存在，则创建该文件;
         * - 'as+': 打开文件用于读取和追加（在同步模式中）。 如果文件不存在，则创建该文件;
         * - 'r': 打开文件用于读取。 如果文件不存在，则会发生异常;
         * - 'r+': 打开文件用于读取和写入。 如果文件不存在，则会发生异常;
         * - 'w': 打开文件用于写入。 如果文件不存在则创建文件，如果文件存在则截断文件;
         * - 'wx': 类似于 'w'，但如果路径存在，则失败;
         * - 'w+': 打开文件用于读取和写入。 如果文件不存在则创建文件，如果文件存在则截断文件;
         * - 'wx+': 类似于 'w+'，但如果路径存在，则失败;
         */
        flag?: "a" | "ax" | "a+" | "ax+" | "as" | "as+" | "r" | "r+" | "w" | "wx" | "w+" | "wx+";
    }

    interface CloseSyncOption {
        /**
         * 需要被关闭的文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
    }

    interface FtruncateSyncOption {
        /**
         * 文件描述符。fd 通过 [FileSystemManager.open](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html) 或 [FileSystemManager.openSync](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html) 接口获得
         */
        fd: string;
        /**
         * 截断位置，默认0。如果 length 小于文件长度（单位：字节），则只有前面 length 个字节会保留在文件中，其余内容会被删除；如果 length 大于文件长度，则会对其进行扩展，并且扩展部分将填充空字节（'\0'）
         */
        length: number;
    }

    interface TruncateSyncOption {
        /**
         * 要截断的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 截断位置，默认0。如果 length 小于文件长度（字节），则只有前面 length 个字节会保留在文件中，其余内容会被删除；如果 length 大于文件长度，则会对其进行扩展，并且扩展部分将填充空字节（'\0'）
         */
        length?: number;
    }

    interface FileSystemManager {
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.readdir](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readdir.html) 的同步版本
         *
         * **注意事项</title>
         * - readdir接口无法访问文件系统根路径(wxfile://)。
         * <title>示例代码**
         *
         * ```js
         * const fs = wx.getFileSystemManager()
         * fs.readdir({
         * dirPath: `${wx.env.USER_DATA_PATH}/example`,
         * success(res) {
         * console.log(res.files)
         * },
         * fail(res) {
         * console.error(res)
         * }
         * })
         *
         * // 同步接口
         * try {
         * const res = fs.readdirSync(`${wx.env.USER_DATA_PATH}/example`)
         * console.log(res)
         * } catch(e) {
         * console.error(e)
         * }
         * ```
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readdirSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readdirSync.html)
         */
        readdirSync(dirPath: string): string[];
        /**
         *
         * 需要基础库： `2.21.1`
         *
         * 在插件中使用：不支持
         *
         * 同步读取指定压缩类型的本地文件内容
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readCompressedFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readCompressedFileSync.html)
         */
        readCompressedFileSync(option: ReadCompressedFileSyncOption): ArrayBuffer;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 判断文件/目录是否存在
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.access.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.access.html)
         */
        access(option: AccessOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.access](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.access.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.accessSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.accessSync.html)
         */
        accessSync(path: string): void;
        /**
         *
         * 需要基础库： `2.1.0`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 在文件结尾追加内容
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.appendFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.appendFile.html)
         */
        appendFile(option: AppendFileOption): void;
        /**
         *
         * 需要基础库： `2.1.0`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.appendFile](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.appendFile.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.appendFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.appendFileSync.html)
         */
        appendFileSync(filePath: string, data: string | ArrayBuffer, encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1"): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 关闭文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.close.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.close.html)
         */
        close(option: FileSystemManagerCloseOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 复制文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.copyFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.copyFile.html)
         */
        copyFile(option: CopyFileOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.copyFile](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.copyFile.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.copyFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.copyFileSync.html)
         */
        copyFileSync(srcPath: string, destPath: string): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 获取文件的状态信息
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.fstat.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.fstat.html)
         */
        fstat(option: FstatOption): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 对文件内容进行截断操作
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.ftruncate.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.ftruncate.html)
         */
        ftruncate(option: FtruncateOption): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 获取该小程序下的 本地临时文件 或 本地缓存文件 信息
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.getFileInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.getFileInfo.html)
         */
        getFileInfo(option: GetFileInfoOption): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 获取该小程序下已保存的本地缓存文件列表
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.getSavedFileList.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.getSavedFileList.html)
         */
        getSavedFileList(option?: GetSavedFileListOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 创建目录
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.mkdir.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.mkdir.html)
         */
        mkdir(option: MkdirOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.mkdir](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.mkdir.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.mkdirSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.mkdirSync.html)
         */
        mkdirSync(dirPath: string, recursive?: boolean): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 打开文件，返回文件描述符
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.open.html)
         */
        open(option: OpenOption): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 读文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.read.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.read.html)
         */
        read(option: ReadOption): void;
        /**
         *
         * 需要基础库： `2.21.1`
         *
         * 在插件中使用：不支持
         *
         * 读取指定压缩类型的本地文件内容
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readCompressedFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readCompressedFile.html)
         */
        readCompressedFile(option: ReadCompressedFileOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 读取本地文件内容。单个文件大小上限为100M。
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readFile.html)
         */
        readFile(option: ReadFileOption): void;
        /**
         *
         * 需要基础库： `2.17.3`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 读取压缩包内的文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readZipEntry.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readZipEntry.html)
         */
        readZipEntry(option: ReadZipEntryOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 读取目录内文件列表
         *
         * **注意事项</title>
         * - readdir接口无法访问文件系统根路径(wxfile://)。
         * <title>示例代码**
         *
         * ```js
         * const fs = wx.getFileSystemManager()
         * fs.readdir({
         * dirPath: `${wx.env.USER_DATA_PATH}/example`,
         * success(res) {
         * console.log(res.files)
         * },
         * fail(res) {
         * console.error(res)
         * }
         * })
         *
         * // 同步接口
         * try {
         * const res = fs.readdirSync(`${wx.env.USER_DATA_PATH}/example`)
         * console.log(res)
         * } catch(e) {
         * console.error(e)
         * }
         * ```
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readdir.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readdir.html)
         */
        readdir(option: ReaddirOption): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 删除该小程序下已保存的本地缓存文件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.removeSavedFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.removeSavedFile.html)
         */
        removeSavedFile(option: RemoveSavedFileOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 重命名文件。可以把文件从 oldPath 移动到 newPath
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rename.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rename.html)
         */
        rename(option: RenameOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.rename](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rename.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.renameSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.renameSync.html)
         */
        renameSync(oldPath: string, newPath: string): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 删除目录
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rmdir.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rmdir.html)
         */
        rmdir(option: RmdirOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.rmdir](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rmdir.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rmdirSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.rmdirSync.html)
         */
        rmdirSync(dirPath: string, recursive?: boolean): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 保存临时文件到本地。此接口会移动临时文件，因此调用成功后，tempFilePath 将不可用。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.saveFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.saveFile.html)
         */
        saveFile(option: SaveFileOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 获取文件 Stats 对象
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.stat.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.stat.html)
         */
        stat(option: StatOption): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 对文件内容进行截断操作
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.truncate.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.truncate.html)
         */
        truncate(option: TruncateOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 删除文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unlink.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unlink.html)
         */
        unlink(option: UnlinkOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.unlink](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unlink.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unlinkSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unlinkSync.html)
         */
        unlinkSync(filePath: string): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 解压文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unzip.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.unzip.html)
         */
        unzip(option: UnzipOption): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 写入文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.write.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.write.html)
         */
        write(option: WriteOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 写文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeFile.html)
         */
        writeFile(option: WriteFileOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.writeFile](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeFile.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeFileSync.html)
         */
        writeFileSync(filePath: string, data: string | ArrayBuffer, encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1"): void;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 读文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readSync.html)
         */
        readSync(option: ReadSyncOption): ReadResult;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 同步获取文件的状态信息
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.fstatSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.fstatSync.html)
         */
        fstatSync(option: FstatSyncOption): Stats;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.stat](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.stat.html) 的同步版本
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.statSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.statSync.html)
         */
        statSync(path: string, recursive?: boolean): Stats | Stats[];
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 同步写入文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.writeSync.html)
         */
        writeSync(option: WriteSyncOption): WriteResult;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 同步打开文件，返回文件描述符
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.openSync.html)
         */
        openSync(option: OpenSyncOption): string;
        /**
         *
         * 在插件中使用：不支持
         *
         * [FileSystemManager.saveFile](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.saveFile.html) 的同步版本
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.saveFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.saveFileSync.html)
         */
        saveFileSync(tempFilePath: string, filePath?: string): string;
        /**
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * [FileSystemManager.readFile](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readFile.html) 的同步版本
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readFileSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.readFileSync.html)
         */
        readFileSync(filePath: string, encoding?: "ascii" | "base64" | "binary" | "hex" | "ucs2" | "ucs-2" | "utf16le" | "utf-16le" | "utf-8" | "utf8" | "latin1", position?: number, length?: number): string | ArrayBuffer;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 同步关闭文件
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.closeSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.closeSync.html)
         */
        closeSync(option: CloseSyncOption): undefined;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 对文件内容进行截断操作
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.ftruncateSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.ftruncateSync.html)
         */
        ftruncateSync(option: FtruncateSyncOption): undefined;
        /**
         *
         * 需要基础库： `2.16.1`
         *
         * 在插件中使用：需要基础库 `2.19.2`
         *
         * 对文件内容进行截断操作 (truncate 的同步版本)
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.truncateSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/FileSystemManager.truncateSync.html)
         */
        truncateSync(option: TruncateSyncOption): undefined;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ExitFullScreenCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ExitFullScreenFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ExitFullScreenSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ExitFullScreenOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ExitFullScreenCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ExitFullScreenFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ExitFullScreenSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ExitPictureInPictureCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ExitPictureInPictureFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ExitPictureInPictureSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface ExitPictureInPictureOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ExitPictureInPictureCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ExitPictureInPictureFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ExitPictureInPictureSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type MuteCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type MuteFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type MuteSuccessCallback = (res: GeneralCallbackResult) => void;

    interface MuteOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: MuteCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: MuteFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: MuteSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PauseCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PauseFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PauseSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PauseOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PauseCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PauseFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PauseSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PlayCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PlayFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PlaySuccessCallback = (res: GeneralCallbackResult) => void;

    interface PlayOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PlayCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PlayFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PlaySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestFullScreenCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type RequestFullScreenFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RequestFullScreenSuccessCallback = (res: GeneralCallbackResult) => void;

    interface LivePlayerContextRequestFullScreenOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestFullScreenCompleteCallback;
        /**
         * 设置全屏时的方向
         *
         * 可选值：
         * - 0: 正常竖向;
         * - 90: 屏幕逆时针90度;
         * - -90: 屏幕顺时针90度;
         */
        direction?: 0 | 90 | -90;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestFullScreenFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestFullScreenSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestPictureInPictureCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type RequestPictureInPictureFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RequestPictureInPictureSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface RequestPictureInPictureOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestPictureInPictureCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestPictureInPictureFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestPictureInPictureSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ResumeCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ResumeFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ResumeSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ResumeOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ResumeCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ResumeFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ResumeSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SnapshotCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SnapshotFailCallback = (res: GeneralCallbackResult) => void;

    interface LivePlayerContextSnapshotSuccessCallbackResult {
        /**
         * 图片的高度
         */
        height: string;
        /**
         * 图片文件的临时路径 (本地路径)
         */
        tempImagePath: string;
        /**
         * 图片的宽度
         */
        width: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type LivePlayerContextSnapshotSuccessCallback = (
                result: LivePlayerContextSnapshotSuccessCallbackResult
            ) => void;

    interface LivePlayerContextSnapshotOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SnapshotCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SnapshotFailCallback;
        /**
         * 需要基础库： `2.10.0`
         *
         * 图片的质量
         *
         * 可选值：
         * - 'raw': 原图;
         * - 'compressed': 压缩图;
         */
        quality?: "raw" | "compressed";
        /**
         * 需要基础库： `2.25.0`
         *
         * 截取的源类型
         *
         * 可选值：
         * - 'stream': 截取视频源;
         * - 'view': 截取渲染后的画面;
         */
        sourceType?: "stream" | "view";
        /**
         * 接口调用成功的回调函数
         */
        success?: LivePlayerContextSnapshotSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopSuccessCallback = (res: GeneralCallbackResult) => void;

    interface StopOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopSuccessCallback;
    }

    interface LivePlayerContext {
        /**
         *
         * 在插件中使用：支持
         *
         * 退出全屏
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.exitFullScreen.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.exitFullScreen.html)
         */
        exitFullScreen(option?: ExitFullScreenOption): void;
        /**
         *
         * 在插件中使用：支持
         *
         * 退出小窗，该方法可在任意页面调用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.exitPictureInPicture.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.exitPictureInPicture.html)
         */
        exitPictureInPicture(option?: ExitPictureInPictureOption): void;
        /**
         *
         * 在插件中使用：支持
         *
         * 静音
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.mute.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.mute.html)
         */
        mute(option?: MuteOption): void;
        /**
         *
         * 需要基础库： `1.9.90`
         *
         * 在插件中使用：支持
         *
         * 暂停
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.pause.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.pause.html)
         */
        pause(option?: PauseOption): void;
        /**
         *
         * 在插件中使用：支持
         *
         * 播放
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.play.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.play.html)
         */
        play(option?: PlayOption): void;
        /**
         *
         * 在插件中使用：支持
         *
         * 进入全屏
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.requestFullScreen.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.requestFullScreen.html)
         */
        requestFullScreen(option: LivePlayerContextRequestFullScreenOption): void;
        /**
         *
         * 需要基础库： `2.15.0`
         *
         * 在插件中使用：支持
         *
         * 进入小窗
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.requestPictureInPicture.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.requestPictureInPicture.html)
         */
        requestPictureInPicture(option?: RequestPictureInPictureOption): void;
        /**
         *
         * 需要基础库： `1.9.90`
         *
         * 在插件中使用：支持
         *
         * 恢复
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.resume.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.resume.html)
         */
        resume(option?: ResumeOption): void;
        /**
         *
         * 需要基础库： `2.7.1`
         *
         * 在插件中使用：支持
         *
         * 截图
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.snapshot.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.snapshot.html)
         */
        snapshot(option: LivePlayerContextSnapshotOption): void;
        /**
         *
         * 在插件中使用：支持
         *
         * 停止
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.stop.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.stop.html)
         */
        stop(option?: StopOption): void;
    }

    interface GetLogManagerOption {
        /**
         * 需要基础库： `2.3.2`
         *
         * 取值为0/1，取值为0表示会把 `App`、`Page` 的生命周期函数和 uni 命名空间下的函数调用写入日志，取值为1则不会。默认值是 0
         */
        level?: number;
    }

    interface LogManager {
        /**
         *
         * 在插件中使用：不支持
         *
         * 写 debug 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.debug.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.debug.html)
         */
        debug(args: any[]): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 写 info 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.info.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.info.html)
         */
        info(args: any[]): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 写 log 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.log.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.log.html)
         */
        log(args: any[]): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 写 warn 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.warn.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/LogManager.warn.html)
         */
        warn(args: any[]): void;
    }

    /**
     * MediaAudioPlayer 实例，可通过 [uni.createMediaAudioPlayer](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createMediaAudioPlayer.html) 接口获取实例。
     */
    interface MediaAudioPlayer {
        /**
         * 音量。范围 0~1。默认为 1
         */
        volume: number;
    }

    /**
     * 需要基础库： `2.9.0`
     *
     * 可通过 [MediaContainer.extractDataSource](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.extractDataSource.html) 返回。
     *
     * [MediaTrack](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaTrack.html) 音频或视频轨道，可以对轨道进行一些操作
     */
    interface MediaTrack {
        /**
         * 轨道长度，只读
         */
        duration: number;
        /**
         * 轨道类型，只读
         *
         * 可选值：
         * - 'audio': 音频轨道;
         * - 'video': 视频轨道;
         */
        kind: "audio" | "video";
        /**
         * 音量，音频轨道下有效，可写
         */
        volume: number;
    }

    interface ExtractDataSourceOption {
        /**
         * 视频源地址，只支持本地文件
         */
        source: string;
    }

    interface MediaContainer {
        /**
         *
         * 需要基础库： `2.9.0`
         *
         * 在插件中使用：支持
         *
         * 将音频或视频轨道添加到容器
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.addTrack.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.addTrack.html)
         */
        addTrack(track: MediaTrack): void;
        /**
         *
         * 需要基础库： `2.9.0`
         *
         * 在插件中使用：支持
         *
         * 将容器销毁，释放资源
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.destroy.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.destroy.html)
         */
        destroy(): void;
        /**
         *
         * 需要基础库： `2.9.0`
         *
         * 在插件中使用：支持
         *
         * 将容器内的轨道合并并导出视频文件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.export.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.export.html)
         */
        export(): void;
        /**
         *
         * 需要基础库： `2.9.0`
         *
         * 在插件中使用：支持
         *
         * 将传入的视频源分离轨道。不会自动将轨道添加到待合成的容器里。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.extractDataSource.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.extractDataSource.html)
         */
        extractDataSource(option: ExtractDataSourceOption): void;
        /**
         *
         * 需要基础库： `2.9.0`
         *
         * 在插件中使用：支持
         *
         * 将音频或视频轨道从容器中移除
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.removeTrack.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/MediaContainer.removeTrack.html)
         */
        removeTrack(track: MediaTrack): void;
    }

    interface CreateMediaRecorderOption {
        /**
         * 指定录制的时长（s)，到达自动停止。最大 7200，最小 5
         */
        duration?: number;
        /**
         * 视频 fps
         */
        fps?: number;
        /**
         * 视频关键帧间隔
         */
        gop?: number;
        /**
         * 视频比特率（kbps），最小值 600，最大值 3000
         */
        videoBitsPerSecond?: number;
    }

    interface MediaRecorder {
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 取消监听录制事件。当对应事件触发时，该回调函数不再执行。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.off.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.off.html)
         */
        off(eventName: string, callback: (...args: any[]) => any): void;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 注册监听录制事件的回调函数。当对应事件触发时，回调函数会被执行。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.on.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.on.html)
         */
        on(eventName: "start" | "stop" | "pause" | "resume" | "timeupdate", callback: (...args: any[]) => any): void;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 销毁录制器
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.destroy.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.destroy.html)
         */
        destroy(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 暂停录制
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.pause.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.pause.html)
         */
        pause(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 请求下一帧录制，在 callback 里完成一帧渲染后开始录制当前帧
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.requestFrame.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.requestFrame.html)
         */
        requestFrame(callback: (...args: any[]) => any): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 恢复录制
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.resume.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.resume.html)
         */
        resume(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 开始录制
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.start.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.start.html)
         */
        start(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 结束录制
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.stop.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/MediaRecorder.stop.html)
         */
        stop(): Promise<any>;
    }

    /**
     * 标签类型枚举
     */
    interface TechType {
        /**
         * 对应IsoDep实例，实例支持ISO-DEP (ISO 14443-4)标准的读写
         */
        isoDep: string;
        /**
         * 对应MifareClassic实例，实例支持MIFARE Classic标签的读写
         */
        mifareClassic: string;
        /**
         * 对应MifareUltralight实例，实例支持MIFARE Ultralight标签的读写
         */
        mifareUltralight: string;
        /**
         * 对应Ndef实例，实例支持对NDEF格式的NFC标签上的NDEF数据的读写
         */
        ndef: string;
        /**
         * 对应NfcA实例，实例支持NFC-A (ISO 14443-3A)标准的读写
         */
        nfcA: string;
        /**
         * 对应NfcB实例，实例支持NFC-B (ISO 14443-3B)标准的读写
         */
        nfcB: string;
        /**
         * 对应NfcF实例，实例支持NFC-F (JIS 6319-4)标准的读写
         */
        nfcF: string;
        /**
         * 对应NfcV实例，实例支持NFC-V (ISO 15693)标准的读写
         */
        nfcV: string;
    }

    /**
     * 需要基础库： `2.11.2`
     */
    interface NFCAdapter {
        /**
         * 标签类型枚举
         */
        tech: TechType;
    }

    /**
     * 需要基础库： `2.7.0`
     *
     * 离屏 canvas 实例，可通过 [uni.createOffscreenCanvas](https://developers.weixin.qq.com/miniprogram/dev/api/canvas/wx.createOffscreenCanvas.html) 创建。
     */
    interface OffscreenCanvas {
        /**
         * 画布高度
         */
        height: number;
        /**
         * 画布宽度
         */
        width: number;
    }

    interface CreateOffscreenCanvasOption {
        /**
         * 在自定义组件下，当前组件实例的 this
         */
        compInst?: any;
        /**
         * 画布高度
         */
        height?: number;
        /**
         * 创建的离屏 canvas 类型
         *
         * 可选值：
         * - 'webgl': webgl类型上下文;
         * - '2d': 2d类型上下文;
         */
        type?: "webgl" | "2d";
        /**
         * 画布宽度
         */
        width?: number;
    }

    /**
     * 单条性能数据。具体数据口径请参考[性能数据文档]((performance/perf_data##_1-4-%E9%80%9A%E8%BF%87-wx-getPerformance-%E5%9C%A8%E5%B0%8F%E7%A8%8B%E5%BA%8F%E5%86%85%E8%8E%B7%E5%8F%96))
     */
    interface PerformanceEntry {
        /**
         * 需要基础库： `2.24.0`
         *
         * 解析域名结束时间。仅 resourceTiming 指标有效。
         */
        domainLookupEnd: number;
        /**
         * 需要基础库： `2.24.0`
         *
         * 解析域名开始时间。仅 resourceTiming 指标有效。
         */
        domainLookupStart: number;
        /**
         * 耗时 ms。仅对于表示阶段的指标有效。
         */
        duration: number;
        /**
         * 指标类型
         *
         * 可选值：
         * - 'navigation': 路由;
         * - 'render': 渲染;
         * - 'script': 脚本;
         */
        entryType: "navigation" | "render" | "script";
        /**
         * 需要基础库： `2.21.2`
         *
         * 注入文件列表。仅 evaluateScript 指标有效。
         */
        fileList: string[];
        /**
         * 需要基础库： `2.21.2`
         *
         * 首次渲染参数在渲染层收到的时间。仅 firstRender 指标有效。
         */
        initDataRecvTime: number;
        /**
         * 需要基础库： `2.21.2`
         *
         * 首次渲染参数从逻辑层发出的时间。仅 firstRender 指标有效。
         */
        initDataSendTime: number;
        /**
         * 需要基础库： `2.24.0`
         *
         * 初始化性能条目的资源类型。仅 resourceTiming 指标有效。
         *
         * 可选值：
         * - 'audio': 音频;
         * - 'cover-image': cover-image 组件的图片;
         * - 'image': 组件的图片;
         * - 'open-data': 组件的图片;
         */
        initiatorType: "audio" | "cover-image" | "image" | "open-data";
        /**
         * 分包名，主包表示为 __APP__ (2.21.2 开始)。仅 evaluateScript 指标有效。
         */
        moduleName: string;
        /**
         * 指标名称
         *
         * 可选值：
         * - 'appLaunch': 小程序启动耗时。(entryType: navigation);
         * - 'route': 路由处理耗时。(entryType: navigation);
         * - 'firstRender': 页面首次渲染耗时。(entryType: render);
         * - 'firstPaint': 页面首次绘制(FP)时间点，无 duration。（iOS 不支持）(entryType: render);
         * - 'firstContentfulPaint': 页面首次内容绘制(FCP)时间点，无 duration。（iOS 14.5 以下版本不支持）(entryType: render);
         * - 'largestContentfulPaint': 页面最大内容绘制(LCP)时间点，无 duration。（iOS 不支持）(entryType: render);
         * - 'evaluateScript': 逻辑层 JS 代码注入耗时。(entryType: script);
         * - 'downloadPackage': 代码包下载耗时。(entryType: loadPackage);
         * - 'resourceTiming': 视图层资源加载耗时。(entryType: resource);
         */
        name: "appLaunch" | "route" | "firstRender" | "firstPaint" | "firstContentfulPaint" | "largestContentfulPaint" | "evaluateScript" | "downloadPackage" | "resourceTiming";
        /**
         * 路由真正响应开始时间。仅 navigation 类型指标有效。
         */
        navigationStart: number;
        /**
         * 路由详细类型，与小程序路由方法对应。仅 navigation 类型指标有效。
         */
        navigationType: string;
        /**
         * 需要基础库： `2.24.0`
         *
         * 代码包名称。仅 downloadPackage 指标有效。
         */
        packageName: string;
        /**
         * 需要基础库： `2.24.0`
         *
         * 代码包大小。仅 downloadPackage 指标有效。
         */
        packageSize: number;
        /**
         * 需要基础库： `2.23.1`
         *
         * path 对应页面实例 Id（随机生成，不保证递增）。仅 render/navigation 指标有效。
         */
        pageId: number;
        /**
         * 页面路径。仅 render 和 navigation 类型指标有效。
         */
        path: string;
        /**
         * 需要基础库： `2.23.1`
         *
         * referrerPath对应页面实例 Id（随机生成，不保证递增）。仅 route 指标有效。
         */
        referrerPageId: number;
        /**
         * 需要基础库： `2.23.1`
         *
         * 页面跳转来源页面路径。仅 route 指标有效。
         */
        referrerPath: number;
        /**
         * 开始时间，不同指标的具体含义会有差异。
         */
        startTime: number;
        /**
         * 需要基础库： `2.24.0`
         *
         * 表示获取资源的大小（以八位字节为单位）的数字。仅 resourceTiming 指标有效。(iOS 不支持）
         */
        transferSize: number;
        /**
         * 需要基础库： `2.24.0`
         *
         * 资源路径。仅 resourceTiming 指标有效。
         */
        uri: string;
        /**
         * 需要基础库： `2.21.2`
         *
         * 渲染层代码注入完成时间。仅 firstRender 指标有效。
         */
        viewLayerReadyTime: number;
        /**
         * 需要基础库： `2.21.2`
         *
         * 渲染层执行渲染结束时间。仅 firstRender 指标有效。
         */
        viewLayerRenderEndTime: number;
        /**
         * 需要基础库： `2.21.2`
         *
         * 渲染层执行渲染开始时间。仅 firstRender 指标有效。
         */
        viewLayerRenderStartTime: number;
    }

    /**
     * 需要基础库： `2.11.0`
     *
     * PerformanceObserver 对象，用于监听性能相关事件
     */
    interface PerformanceObserver {
        /**
         * 获取当前支持的所有性能指标类型
         */
        supportedEntryTypes: any[];
    }

    interface Performance {
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：不支持
         *
         * 该方法返回当前缓冲区中的所有性能数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntries.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntries.html)
         */
        getEntries(): PerformanceEntry[];
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：不支持
         *
         * 获取当前缓冲区中所有名称为 [name] 且类型为 [entryType] 的性能数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntriesByName.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntriesByName.html)
         */
        getEntriesByName(name: string, entryType?: string): PerformanceEntry[];
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：不支持
         *
         * 获取当前缓冲区中所有类型为 [entryType] 的性能数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntriesByType.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.getEntriesByType.html)
         */
        getEntriesByType(entryType: string): PerformanceEntry[];
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：不支持
         *
         * 设置缓冲区大小，默认缓冲 30 条性能数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.setBufferSize.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.setBufferSize.html)
         */
        setBufferSize(size: number): void;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：不支持
         *
         * 创建全局性能事件监听器
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.createObserver.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/Performance.createObserver.html)
         */
        createObserver(callback: (...args: any[]) => any): PerformanceObserver;
    }

    interface PreDownloadSubpackageOption {
        /**
         * 分包加载结束回调事件(加载成功、失败都会执行）
         */
        complete: (...args: any[]) => any;
        /**
         * 分包加载失败回调事件
         */
        fail: (...args: any[]) => any;
        /**
         * 分包的类型。目前仅支持填 "workers"，表示 workers 分包。
         */
        packageType: string;
        /**
         * 分包加载成功回调事件
         */
        success: (...args: any[]) => any;
    }

    interface PreDownloadSubpackageTaskOnProgressUpdateListenerResult {
        /**
         * 分包下载进度百分比
         */
        progress: number;
        /**
         * 预期需要下载的数据总长度，单位 Bytes
         */
        totalBytesExpectedToWrite: number;
        /**
         * 已经下载的数据长度，单位 Bytes
         */
        totalBytesWritten: number;
    }

    /**
     * 分包加载进度变化事件的监听函数
     */
    type PreDownloadSubpackageTaskOnProgressUpdateCallback = (
                result: PreDownloadSubpackageTaskOnProgressUpdateListenerResult
            ) => void;

    interface PreDownloadSubpackageTask {
        /**
         *
         * 需要基础库： `2.27.3`
         *
         * 在插件中使用：不支持
         *
         * 监听分包加载进度变化事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/subpackage/PreDownloadSubpackageTask.onProgressUpdate.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/subpackage/PreDownloadSubpackageTask.onProgressUpdate.html)
         */
        onProgressUpdate(listener: PreDownloadSubpackageTaskOnProgressUpdateCallback): void;
    }

    interface CurrentState {
        /**
         * 当前缓存中的日志条数
         */
        logCount: number;
        /**
         * 当前缓存中最大可存日志条数
         */
        maxLogCount: number;
        /**
         * 当前缓存最大可用空间，以字节为单位
         */
        maxSize: number;
        /**
         * 当前缓存中已使用空间，以字节为单位
         */
        size: number;
    }

    interface RealtimeTagLogManager {
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 添加过滤关键字
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.addFilterMsg.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.addFilterMsg.html)
         */
        addFilterMsg(msg: string): void;
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 写 error 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.error.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.error.html)
         */
        error(key: string, value: any): void;
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 写 info 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.info.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.info.html)
         */
        info(key: string, value: any): void;
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 设置过滤关键字
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.setFilterMsg.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.setFilterMsg.html)
         */
        setFilterMsg(msg: string): void;
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 写 warn 日志
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.warn.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeTagLogManager.warn.html)
         */
        warn(key: string, value: any): void;
    }

    interface RealtimeLogManager {
        /**
         *
         * 需要基础库： `2.19.4`
         *
         * 在插件中使用：不支持
         *
         * 实时日志会将一定时间间隔内缓存的日志聚合上报，如果该时间内缓存的内容超出限制，则会被丢弃。此方法可以获取当前缓存剩余空间。
         *
         * > 注意：基础库内部在对日志进行上报时会补充一些结构化数据，如果遇到上报溢出的情况也会补充警告日志，所以此方法获取到的当前占用信息会比预期的大一些。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.getCurrentState.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.getCurrentState.html)
         */
        getCurrentState(): CurrentState;
        /**
         *
         * 需要基础库： `2.8.1`
         *
         * 在插件中使用：不支持
         *
         * 添加过滤关键字，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.addFilterMsg.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.addFilterMsg.html)
         */
        addFilterMsg(msg: string): void;
        /**
         *
         * 需要基础库： `2.7.1`
         *
         * 在插件中使用：不支持
         *
         * 写 error 日志，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.error.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.error.html)
         */
        error(args: any[]): void;
        /**
         *
         * 需要基础库： `2.9.1`
         *
         * 在插件中使用：不支持
         *
         * 设置实时日志page参数所在的页面，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.in.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.in.html)
         */
        in(pageInstance: any): void;
        /**
         *
         * 需要基础库： `2.7.1`
         *
         * 在插件中使用：不支持
         *
         * 写 info 日志，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.info.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.info.html)
         */
        info(args: any[]): void;
        /**
         *
         * 需要基础库： `2.7.3`
         *
         * 在插件中使用：不支持
         *
         * 设置过滤关键字，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.setFilterMsg.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.setFilterMsg.html)
         */
        setFilterMsg(msg: string): void;
        /**
         *
         * 需要基础库： `2.7.1`
         *
         * 在插件中使用：不支持
         *
         * 写 warn 日志，暂不支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.warn.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.warn.html)
         */
        warn(args: any[]): void;
        /**
         *
         * 需要基础库： `2.16.0`
         *
         * 在插件中使用：需要基础库 `2.16.0`
         *
         * 获取给定标签的日志管理器实例，目前只支持在插件使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.tag.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/RealtimeLogManager.tag.html)
         */
        tag(tagName: string): RealtimeTagLogManager;
    }

    interface BindWifiOption {
        /**
         * 当前 wifi 网络的 BSSID ，可通过 wx.getConnectedWifi 获取
         */
        BSSID: string;
    }

    interface TCPSocketConnectOption {
        /**
         * 套接字要连接的地址
         */
        address: string;
        /**
         * 套接字要连接的端口
         */
        port: number;
        /**
         * 套接字要连接的超时时间，默认为 2s
         */
        timeout?: number;
    }

    /**
     * onBindWifi 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBindWifiCallback = (res: GeneralCallbackResult) => void;
    /**
     * onClose 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type UDPSocketOffCloseCallback = (res: GeneralCallbackResult) => void;
    /**
     * onConnect 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffConnectCallback = (res: GeneralCallbackResult) => void;
    /**
     * onError 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type UDPSocketOffErrorCallback = (result: GeneralCallbackResult) => void;

    interface LocalInfo {
        /**
         * 接收消息的 socket 的地址
         */
        address: string;
        /**
         * 使用的协议族，为 IPv4 或者 IPv6
         */
        family: string;
        /**
         * 端口号
         */
        port: number;
    }

    /**
     * 发送端地址信息
     */
    interface RemoteInfo {
        /**
         * 发送消息的 socket 的地址
         */
        address: string;
        /**
         * 使用的协议族，为 IPv4 或者 IPv6
         */
        family: string;
        /**
         * 端口号
         */
        port: number;
        /**
         * message 的大小，单位：字节
         */
        size: number;
    }

    interface TCPSocketOnMessageListenerResult {
        /**
         * 接收端地址信息
         */
        localInfo: LocalInfo;
        /**
         * 收到的消息
         */
        message: ArrayBuffer;
        /**
         * 发送端地址信息
         */
        remoteInfo: RemoteInfo;
    }

    /**
     * onMessage 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type TCPSocketOffMessageCallback = (
                result: TCPSocketOnMessageListenerResult
            ) => void;
    /**
     * 当一个 socket 绑定当前 wifi 网络成功时触发该事件的监听函数
     */
    type OnBindWifiCallback = (res: GeneralCallbackResult) => void;
    type UDPSocketOnCloseCallback = (res: GeneralCallbackResult) => void;
    /**
     * 当一个 socket 连接成功建立的时候触发该事件的监听函数
     */
    type OnConnectCallback = (res: GeneralCallbackResult) => void;
    type UDPSocketOnErrorCallback = (result: GeneralCallbackResult) => void;
    /**
     * 当接收到数据的时触发该事件的监听函数
     */
    type TCPSocketOnMessageCallback = (
                result: TCPSocketOnMessageListenerResult
            ) => void;

    interface TCPSocket {
        /**
         *
         * 需要基础库： `2.25.0`
         *
         * 在插件中使用：不支持
         *
         * 将 TCP Socket 绑定到当前 wifi 网络，成功后会触发 onBindWifi 事件（仅安卓支持）
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.bindWifi.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.bindWifi.html)
         */
        bindWifi(options: BindWifiOption): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 关闭连接
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.close.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.close.html)
         */
        close(): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 在给定的套接字上启动连接
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.connect.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.connect.html)
         */
        connect(options: TCPSocketConnectOption): void;
        /**
         *
         * 需要基础库： `2.25.0`
         *
         * 在插件中使用：不支持
         *
         * 移除当一个 socket 绑定当前 wifi 网络成功时触发该事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offBindWifi.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offBindWifi.html)
         */
        offBindWifi(listener?: OffBindWifiCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 移除一旦 socket 完全关闭就发出该事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offClose.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offClose.html)
         */
        offClose(listener?: UDPSocketOffCloseCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 移除当一个 socket 连接成功建立的时候触发该事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offConnect.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offConnect.html)
         */
        offConnect(listener?: OffConnectCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 移除当错误发生时触发的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offError.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offError.html)
         */
        offError(listener?: UDPSocketOffErrorCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 移除当接收到数据的时触发该事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.offMessage.html)
         */
        offMessage(listener?: TCPSocketOffMessageCallback): void;
        /**
         *
         * 需要基础库： `2.25.0`
         *
         * 在插件中使用：不支持
         *
         * 监听当一个 socket 绑定当前 wifi 网络成功时触发该事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onBindWifi.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onBindWifi.html)
         */
        onBindWifi(listener: OnBindWifiCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 监听一旦 socket 完全关闭就发出该事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onClose.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onClose.html)
         */
        onClose(listener: UDPSocketOnCloseCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 监听当一个 socket 连接成功建立的时候触发该事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onConnect.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onConnect.html)
         */
        onConnect(listener: OnConnectCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 监听当错误发生时触发
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onError.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onError.html)
         */
        onError(listener: UDPSocketOnErrorCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 监听当接收到数据的时触发该事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.onMessage.html)
         */
        onMessage(listener: TCPSocketOnMessageCallback): void;
        /**
         *
         * 在插件中使用：不支持
         *
         * 在 socket 上发送数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.write.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/TCPSocket.write.html)
         */
        write(data: string | ArrayBuffer): void;
    }

    interface UDPSocketConnectOption {
        /**
         * 要发消息的地址
         */
        address: string;
        /**
         * 要发送消息的端口号
         */
        port: number;
    }

    /**
     * onListening 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffListeningCallback = (res: GeneralCallbackResult) => void;

    interface UDPSocketOnMessageListenerResult {
        /**
         * 接收端地址信息，2.18.0 起支持
         */
        localInfo: LocalInfo;
        /**
         * 收到的消息。消息长度需要小于4096。
         */
        message: ArrayBuffer;
        /**
         * 发送端地址信息
         */
        remoteInfo: RemoteInfo;
    }

    /**
     * onMessage 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type UDPSocketOffMessageCallback = (
                result: UDPSocketOnMessageListenerResult
            ) => void;
    /**
     * 开始监听数据包消息的事件的监听函数
     */
    type OnListeningCallback = (res: GeneralCallbackResult) => void;
    /**
     * 收到消息的事件的监听函数
     */
    type UDPSocketOnMessageCallback = (
                result: UDPSocketOnMessageListenerResult
            ) => void;

    interface UDPSocketSendOption {
        /**
         * 要发消息的地址。在基础库 <= 2.9.3 版本必须是和本机同网段的 IP 地址，或安全域名列表内的域名地址；之后版本可以是任意 IP 和域名
         */
        address: string;
        /**
         * 要发送的数据
         */
        message: string | ArrayBuffer;
        /**
         * 要发送消息的端口号
         */
        port: number;
        /**
         * 发送数据的长度，仅当 message 为 ArrayBuffer 类型时有效
         */
        length?: number;
        /**
         * 发送数据的偏移量，仅当 message 为 ArrayBuffer 类型时有效
         */
        offset?: number;
        /**
         * 向指定地址发消息时，是否要开启广播，基础库 2.24.0 开始支持
         */
        setBroadcast?: boolean;
    }

    interface UDPSocket {
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 关闭 UDP Socket 实例，相当于销毁。 在关闭之后，UDP Socket 实例不能再发送消息，每次调用 `UDPSocket.send` 将会触发错误事件，并且 message 事件回调函数也不会再也执行。在 `UDPSocket` 实例被创建后将被 Native 强引用，保证其不被 GC。在 `UDPSocket.close` 后将解除对其的强引用，让 UDPSocket 实例遵从 GC。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.close.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.close.html)
         */
        close(): void;
        /**
         *
         * 需要基础库： `2.15.0`
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 预先连接到指定的 IP 和 port，需要配合 write 方法一起使用
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.connect.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.connect.html)
         */
        connect(option: UDPSocketConnectOption): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 移除关闭事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offClose.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offClose.html)
         */
        offClose(listener?: UDPSocketOffCloseCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 移除错误事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offError.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offError.html)
         */
        offError(listener?: UDPSocketOffErrorCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 移除开始监听数据包消息的事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offListening.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offListening.html)
         */
        offListening(listener?: OffListeningCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 移除收到消息的事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.offMessage.html)
         */
        offMessage(listener?: UDPSocketOffMessageCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 监听关闭事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onClose.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onClose.html)
         */
        onClose(listener: UDPSocketOnCloseCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 监听错误事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onError.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onError.html)
         */
        onError(listener: UDPSocketOnErrorCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 监听开始监听数据包消息的事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onListening.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onListening.html)
         */
        onListening(listener: OnListeningCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 监听收到消息的事件
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.onMessage.html)
         */
        onMessage(listener: UDPSocketOnMessageCallback): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 向指定的 IP 和 port 发送消息。基础库 2.9.0 起支持广播 (指定地址为 255.255.255.255)。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.send.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.send.html)
         */
        send(option: UDPSocketSendOption): void;
        /**
         *
         * 需要基础库： `2.18.0`
         *
         * 在插件中使用：支持
         *
         * 设置 IP_TTL 套接字选项，用于设置一个 IP 数据包传输时允许的最大跳步数
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.setTTL.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.setTTL.html)
         */
        setTTL(ttl: number): void;
        /**
         *
         * 需要基础库： `2.15.0`
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 用法与 send 方法相同，如果没有预先调用 connect 则与 send 无差异（注意即使调用了 connect 也需要在本接口填入地址和端口参数）
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.write.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.write.html)
         */
        write(): void;
        /**
         *
         * 在插件中使用：需要基础库 `2.11.1`
         *
         * 绑定一个系统随机分配的可用端口，或绑定一个指定的端口号
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.bind.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/UDPSocket.bind.html)
         */
        bind(port?: number): number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetLatestUserKeyCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetLatestUserKeyFailCallback = (res: GeneralCallbackResult) => void;

    interface GetLatestUserKeySuccessCallbackResult {
        /**
         * 用户加密密钥
         */
        encryptKey: string;
        /**
         * 密钥过期时间
         */
        expireTime: number;
        /**
         * 密钥初始向量
         */
        iv: string;
        /**
         * 密钥版本
         */
        version: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetLatestUserKeySuccessCallback = (
                result: GetLatestUserKeySuccessCallbackResult
            ) => void;

    interface GetLatestUserKeyOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetLatestUserKeyCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetLatestUserKeyFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetLatestUserKeySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetRandomValuesCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetRandomValuesFailCallback = (res: GeneralCallbackResult) => void;

    interface GetRandomValuesSuccessCallbackResult {
        /**
         * 随机数内容，长度为传入的字节数
         */
        randomValues: ArrayBuffer;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetRandomValuesSuccessCallback = (
                result: GetRandomValuesSuccessCallbackResult
            ) => void;

    interface GetRandomValuesOption {
        /**
         * 整数，生成随机数的字节数，最大 1048576
         */
        length: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetRandomValuesCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetRandomValuesFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetRandomValuesSuccessCallback;
    }

    interface UserCryptoManager {
        /**
         *
         * 需要基础库： `2.17.3`
         *
         * 在插件中使用：不支持
         *
         * 获取最新的用户加密密钥
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/UserCryptoManager.getLatestUserKey.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/UserCryptoManager.getLatestUserKey.html)
         */
        getLatestUserKey(option?: GetLatestUserKeyOption): void;
        /**
         *
         * 需要基础库： `2.17.3`
         *
         * 在插件中使用：不支持
         *
         * 获取密码学安全随机数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/UserCryptoManager.getRandomValues.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/UserCryptoManager.getRandomValues.html)
         */
        getRandomValues(option: GetRandomValuesOption): void;
    }

    /**
     * 平面跟踪配置
     */
    interface PlaneTrack {
        /**
         * 平面跟踪配置模式
         *
         * 可选值：
         * - 1: 检测横向平面;
         * - 2: 检测纵向平面，只有 v2 版本支持;
         * - 3: 检测横向和纵向平面，只有 v2 版本支持;
         */
        mode: 1 | 2 | 3;
    }

    /**
     * 需要基础库： `2.27.0`
     *
     * OCR检测配置。用法详情[指南文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/visionkit/ocr.html)。
     */
    interface OCRTrack {
        /**
         * 需要基础库： `2.27.0`
         *
         * OCR检测模式
         *
         * 可选值：
         * - 1: 通过摄像头实时检测;
         * - 2: 静态图片检测;
         */
        mode: 1 | 2;
    }

    /**
     * 需要基础库： `2.25.0`
     *
     * 人脸检测配置。用法详情[指南文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/visionkit/face.html)。安卓微信8.0.25开始支持，iOS微信8.0.24开始支持。
     */
    interface FaceTrack {
        /**
         * 需要基础库： `2.25.0`
         *
         * 人脸检测模式
         *
         * 可选值：
         * - 1: 通过摄像头实时检测;
         * - 2: 静态图片检测;
         */
        mode: 1 | 2;
    }

    /**
     * 跟踪能力配置，目前不同的跟踪能力之间是互斥的，默认使用平面跟踪能力。需要注意目前 track 中不同的跟踪配置存在互斥关系（比如 marker 跟踪配置和 OSD 跟踪配置不能同时存在），请按需配置。
     */
    interface Track {
        /**
         * 平面跟踪配置
         */
        plane: PlaneTrack;
        /**
         * 需要基础库： `2.27.0`
         *
         * OCR检测配置。用法详情[指南文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/visionkit/ocr.html)。
         */
        OCR?: OCRTrack;
        /**
         * 需要基础库： `2.24.5`
         *
         * OSD 跟踪配置
         */
        OSD?: boolean;
        /**
         * 需要基础库： `2.25.0`
         *
         * 人脸检测配置。用法详情[指南文档](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/visionkit/face.html)。安卓微信8.0.25开始支持，iOS微信8.0.24开始支持。
         */
        face?: FaceTrack;
        /**
         * 需要基础库： `2.24.5`
         *
         * marker 跟踪配置
         */
        marker?: boolean;
    }

    interface VKConfig {
        /**
         * 跟踪能力配置，目前不同的跟踪能力之间是互斥的，默认使用平面跟踪能力。需要注意目前 track 中不同的跟踪配置存在互斥关系（比如 marker 跟踪配置和 OSD 跟踪配置不能同时存在），请按需配置。
         */
        track: Track;
        /**
         * 需要基础库： `2.23.0`
         *
         * 绑定的 WebGLRenderingContext 对象
         */
        gl?: any;
        /**
         * 需要基础库： `2.22.0`
         *
         * vision kit 版本。
         *
         * 可选值：
         * - 'v1': v1适用于用户在平面场景下，例如桌面，地面，泛平面场景，放置虚拟物体，不提供真实世界距离。用户放置物体时，手机相机倾斜向下对着目标平面点击即可，具有广泛的机型支持;
         * - 'v2': v2提供真实物理距离的 ar 定位功能，提供平面识别功能，用户在平面范围点击放置虚拟物体的功能，具有[有限的机型支持](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/visionkit/plane.html#%E9%99%84%E5%BD%95)。iOS 设备在基础库 2.22.0 开始支持v2。安卓设备在基础库 2.25.1 开始支持v2，另外，安卓v2不支持竖直平面。;
         */
        version?: "v1" | "v2";
    }

    interface VKSize {
        /**
         * 高度
         */
        height: number;
        /**
         * 宽度
         */
        width: number;
    }

    /**
     * 需要基础库： `2.20.0`
     *
     * vision kit 会话对象。
     */
    interface VKSession {
        /**
         * 相机尺寸
         */
        cameraSize: VKSize;
        /**
         * 会话配置
         */
        config: VKConfig;
        /**
         * 会话状态
         *
         * 可选值：
         * - 0: 不可用;
         * - 1: 运行中;
         * - 2: 暂停中;
         */
        state: 0 | 1 | 2;
    }

    /**
     * 视频帧数据，若取不到则返回 null。当缓冲区为空的时候可能暂停取不到数据。
     */
    interface FrameDataOptions {
        /**
         * 帧数据
         */
        data: ArrayBuffer;
        /**
         * 帧数据高度
         */
        height: number;
        /**
         * 帧原始 dts
         */
        pkDts: number;
        /**
         * 帧原始 pts
         */
        pkPts: number;
        /**
         * 帧数据宽度
         */
        width: number;
    }

    interface VideoDecoderStartOption {
        /**
         * 需要解码的视频源文件。基础库 2.13.0 以下的版本只支持本地路径。 2.13.0 开始支持 http:// 和 https:// 协议的远程路径。
         */
        source: string;
        /**
         * 需要基础库： `2.15.0`
         *
         * 是否不需要音频轨道
         */
        abortAudio?: boolean;
        /**
         * 需要基础库： `2.15.0`
         *
         * 是否不需要视频轨道
         */
        abortVideo?: boolean;
        /**
         * 解码模式。0：按 pts 解码；1：以最快速度解码
         */
        mode?: number;
    }

    interface VideoDecoder {
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 获取下一帧的解码数据
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.getFrameData.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.getFrameData.html)
         */
        getFrameData(): FrameDataOptions;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 移除解码器
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.remove.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.remove.html)
         */
        remove(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 跳到某个时间点解码
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.seek.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.seek.html)
         */
        seek(position: number): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 开始解码
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.start.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.start.html)
         */
        start(option: VideoDecoderStartOption): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 停止解码
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.stop.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.stop.html)
         */
        stop(): Promise<any>;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 取消监听录制事件。当对应事件触发时，该回调函数不再执行
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.off.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.off.html)
         */
        off(eventName: string, callback: (...args: any[]) => any): void;
        /**
         *
         * 需要基础库： `2.11.0`
         *
         * 在插件中使用：支持
         *
         * 注册监听录制事件的回调函数。当对应事件触发时，回调函数会被执行
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.on.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.on.html)
         */
        on(eventName: "start" | "stop" | "seek" | "bufferchange" | "ended", callback: (...args: any[]) => any): void;
    }

    /**
     * 需要基础库： `2.19.0`
     *
     * 一类音频处理模块，不同的Node具备不同的功能，如GainNode(音量调整)等。一个WebAudioContextNode可以通过上下文来创建。
     * 目前已经支持以下Node：
     * IIRFilterNode
     * WaveShaperNode
     * ConstantSourceNode
     * ChannelMergerNode
     * OscillatorNode
     * GainNode
     * BiquadFilterNode
     * PeriodicWaveNode
     * BufferSourceNode
     * ChannelSplitterNode
     * ChannelMergerNode
     * DelayNode
     * DynamicsCompressorNode
     * ScriptProcessorNode
     * PannerNode
     * AnalyserNode
     */
    interface WebAudioContextNode {
        /**
         * 表示监听器的前向系统在同一笛卡尔坐标系中的水平位置，作为位置（位置x，位置和位置和位置）值。
         */
        forwardX: number;
        /**
         * 表示听众的前向方向在同一笛卡尔坐标系中作为位置（位置x，位置和位置和位置）值的垂直位置。
         */
        forwardY: number;
        /**
         * 表示与position (positionX、positionY和positionZ)值在同一笛卡尔坐标系下的听者前进方向的纵向(前后)位置。
         */
        forwardZ: number;
        /**
         * 右手笛卡尔坐标系中X轴的位置。
         */
        positionX: number;
        /**
         * 右手笛卡尔坐标系中Y轴的位置。
         */
        positionY: number;
        /**
         * 右手笛卡尔坐标系中Z轴的位置。
         */
        positionZ: number;
        /**
         * 设置监听器的方向
         */
        setOrientation: (...args: any[]) => any;
        /**
         * 设置监听器的位置
         *
         * /**
         */
        setPosition: (...args: any[]) => any;
        /**
         * 表示在与position (positionX、positionY和positionZ)值相同的笛卡尔坐标系中侦听器向前方向的水平位置。
         */
        upX: number;
        /**
         * 表示在与position (positionX、positionY和positionZ)值相同的笛卡尔坐标系中侦听器向上方向的水平位置。
         */
        upY: number;
        /**
         * 表示在与position (positionX、positionY和positionZ)值相同的笛卡尔坐标系中侦听器向后方向的水平位置。
         */
        upZ: number;
    }

    /**
     * 需要基础库： `2.19.0`
     *
     * WebAudioContext 实例，通过[uni.createWebAudioContext](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createWebAudioContext.html) 接口获取该实例。
     *
     */
    interface WebAudioContext {
        /**
         * 获取当前上下文的时间戳。
         */
        currentTime: number;
        /**
         *
         * 当前上下文的最终目标节点，一般是音频渲染设备。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/WebAudioContextNode.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/WebAudioContextNode.html)
         */
        destination: WebAudioContextNode;
        /**
         * 空间音频监听器。
         */
        listener: any;
        /**
         * 可写属性，开发者可以对该属性设置一个监听函数，当WebAudio状态改变的时候，会触发开发者设置的监听函数。
         */
        onstatechange: (...args: any[]) => any;
        /**
         * 采样率，通常在8000-96000之间，通常44100hz的采样率最为常见。
         */
        sampleRate: number;
        /**
         * 当前WebAudio上下文的状态。可能的值如下：suspended（暂停）、running（正在运行）、closed（已关闭）。需要注意的是，不要在 audioContext close后再访问state属性
         */
        state: string;
    }

    /**
     * 可选参数
     */
    interface CreateWorkerOption {
        /**
         * 需要基础库： `2.13.0`
         *
         * 是否使用实验worker。在iOS下，实验worker的JS运行效率比非实验worker提升数倍，如需在worker内进行重度计算的建议开启此选项。同时，实验worker存在极小概率会在系统资源紧张时被系统回收，因此建议配合 worker.onProcessKilled 事件使用，在worker被回收后可重新创建一个。
         */
        useExperimentalWorker?: boolean;
    }

    /**
     * 需要添加的卡券列表
     */
    interface AddCardRequestInfo {
        /**
         * 卡券的扩展参数。需将 CardExt 对象 JSON 序列化为**字符串**传入
         */
        cardExt: string;
        /**
         * 卡券 ID
         */
        cardId: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddCardCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddCardFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 卡券添加结果列表
     */
    interface AddCardResponseInfo {
        /**
         * 卡券的扩展参数，结构请参考下文
         */
        cardExt: string;
        /**
         * 用户领取到卡券的 ID
         */
        cardId: string;
        /**
         * 加密 code，为用户领取到卡券的code加密后的字符串，解密请参照：[code 解码接口](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1499332673_Unm7V)
         */
        code: string;
        /**
         * 是否成功
         */
        isSuccess: boolean;
    }

    interface AddCardSuccessCallbackResult {
        /**
         * 卡券添加结果列表
         */
        cardList: AddCardResponseInfo[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type AddCardSuccessCallback = (result: AddCardSuccessCallbackResult) => void;

    interface AddCardOption {
        /**
         * 需要添加的卡券列表
         */
        cardList: AddCardRequestInfo[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddCardCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddCardFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AddCardSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddFileToFavoritesCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddFileToFavoritesFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AddFileToFavoritesSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface AddFileToFavoritesOption {
        /**
         * 要收藏的文件地址，必须为本地路径或临时路径
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddFileToFavoritesCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddFileToFavoritesFailCallback;
        /**
         * 自定义文件名，若留空则使用filePath中的文件名
         */
        fileName?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: AddFileToFavoritesSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddPhoneCalendarCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddPhoneCalendarFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AddPhoneCalendarSuccessCallback = (res: GeneralCallbackResult) => void;

    interface AddPhoneCalendarOption {
        /**
         * 开始时间的 unix 时间戳
         */
        startTime: number;
        /**
         * 日历事件标题
         */
        title: string;
        /**
         * 是否提醒，默认 true
         */
        alarm?: boolean;
        /**
         * 提醒提前量，单位秒，默认 0 表示开始时提醒
         */
        alarmOffset?: number;
        /**
         * 是否全天事件，默认 false
         */
        allDay?: boolean;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddPhoneCalendarCompleteCallback;
        /**
         * 事件说明
         */
        description?: string;
        /**
         * 结束时间的 unix 时间戳，默认与开始时间相同
         */
        endTime?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddPhoneCalendarFailCallback;
        /**
         * 事件位置
         */
        location?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: AddPhoneCalendarSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddPhoneRepeatCalendarCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddPhoneRepeatCalendarFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AddPhoneRepeatCalendarSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface AddPhoneRepeatCalendarOption {
        /**
         * 开始时间的 unix 时间戳 (1970年1月1日开始所经过的秒数)
         */
        startTime: number;
        /**
         * 日历事件标题
         */
        title: string;
        /**
         * 是否提醒，默认 true
         */
        alarm?: boolean;
        /**
         * 提醒提前量，单位秒，默认 0 表示开始时提醒
         */
        alarmOffset?: number;
        /**
         * 是否全天事件，默认 false
         */
        allDay?: boolean;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddPhoneRepeatCalendarCompleteCallback;
        /**
         * 事件说明
         */
        description?: string;
        /**
         * 结束时间的 unix 时间戳，默认与开始时间相同
         */
        endTime?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddPhoneRepeatCalendarFailCallback;
        /**
         * 事件位置
         */
        location?: string;
        /**
         * 重复周期结束时间的 unix 时间戳，不填表示一直重复
         */
        repeatEndTime?: number;
        /**
         * 重复周期，默认 month 每月重复
         *
         * 可选值：
         * - 'day': 每天重复;
         * - 'week': 每周重复;
         * - 'month': 每月重复。该模式日期不能大于 28 日;
         * - 'year': 每年重复;
         */
        repeatInterval?: "day" | "week" | "month" | "year";
        /**
         * 接口调用成功的回调函数
         */
        success?: AddPhoneRepeatCalendarSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddVideoToFavoritesCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddVideoToFavoritesFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AddVideoToFavoritesSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface AddVideoToFavoritesOption {
        /**
         * 要收藏的视频地址，必须为本地路径或临时路径
         */
        videoPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddVideoToFavoritesCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddVideoToFavoritesFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AddVideoToFavoritesSuccessCallback;
        /**
         * 缩略图路径，若留空则使用视频首帧
         */
        thumbPath?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AuthPrivateMessageCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AuthPrivateMessageFailCallback = (res: GeneralCallbackResult) => void;

    interface AuthPrivateMessageSuccessCallbackResult {
        /**
         * 经过加密的activityId，解密后可得到原始的activityId。若解密后得到的activityId可以与开发者后台的活动id对应上则验证通过，否则表明valid字段不可靠（被篡改） 详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
         */
        encryptedData: string;
        /**
         * 错误信息
         */
        errMsg: string;
        /**
         * 加密算法的初始向量，详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
         */
        iv: string;
        /**
         * 验证是否通过
         */
        valid: boolean;
    }

    /**
     * 接口调用成功的回调函数
     */
    type AuthPrivateMessageSuccessCallback = (
                result: AuthPrivateMessageSuccessCallbackResult
            ) => void;

    interface AuthPrivateMessageOption {
        /**
         * shareTicket。可以从 wx.onShow 中获取。详情 [shareTicket](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share.html)
         */
        shareTicket: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AuthPrivateMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AuthPrivateMessageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AuthPrivateMessageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AuthorizeForMiniProgramCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AuthorizeForMiniProgramFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AuthorizeForMiniProgramSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface AuthorizeForMiniProgramOption {
        /**
         * 需要获取权限的 scope，详见 [scope 列表](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/authorize.html#scope-列表)
         *
         * 可选值：
         * - 'scope.record': ;
         * - 'scope.writePhotosAlbum': ;
         * - 'scope.camera': ;
         */
        scope: "scope.record" | "scope.writePhotosAlbum" | "scope.camera";
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AuthorizeForMiniProgramCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AuthorizeForMiniProgramFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AuthorizeForMiniProgramSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type BatchGetStorageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type BatchGetStorageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type BatchGetStorageSuccessCallback = (res: GeneralCallbackResult) => void;

    interface BatchGetStorageOption {
        /**
         * 本地缓存中指定的 keyList
         */
        keyList: string[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: BatchGetStorageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: BatchGetStorageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: BatchGetStorageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type BatchSetStorageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type BatchSetStorageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type BatchSetStorageSuccessCallback = (res: GeneralCallbackResult) => void;

    interface BatchSetStorageOption {
        /**
         * [{ key, value }]
         */
        kvList: any[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: BatchSetStorageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: BatchSetStorageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: BatchSetStorageSuccessCallback;
    }

    interface KvList {
        /**
         * key 本地缓存中指定的 key
         */
        key: string;
        /**
         * data 需要存储的内容。只支持原生类型、Date、及能够通过`JSON.stringify`序列化的对象。
         */
        value: any;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type CheckIsOpenAccessibilityCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type CheckIsOpenAccessibilityFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface CheckIsOpenAccessibilitySuccessCallbackOption {
        /**
         * iOS 上开启辅助功能旁白，安卓开启 talkback 时返回 true
         */
        open: boolean;
    }

    /**
     * 接口调用成功的回调函数
     */
    type CheckIsOpenAccessibilitySuccessCallback = (
                option: CheckIsOpenAccessibilitySuccessCallbackOption
            ) => void;

    interface CheckIsOpenAccessibilityOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: CheckIsOpenAccessibilityCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: CheckIsOpenAccessibilityFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: CheckIsOpenAccessibilitySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChooseContactCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChooseContactFailCallback = (res: GeneralCallbackResult) => void;

    interface ChooseContactSuccessCallbackOption {
        /**
         * 联系人姓名
         */
        displayName: string;
        /**
         * 手机号
         */
        phoneNumber: string;
        /**
         * 选定联系人的所有手机号（部分 Android 系统只能选联系人而不能选特定手机号）
         */
        phoneNumberList: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChooseContactSuccessCallback = (
                option: ChooseContactSuccessCallbackOption
            ) => void;

    interface ChooseContactOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChooseContactCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChooseContactFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ChooseContactSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChooseInvoiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChooseInvoiceFailCallback = (res: GeneralCallbackResult) => void;

    interface ChooseInvoiceSuccessCallbackResult {
        /**
         * 用户选中的发票信息，格式为一个 JSON 字符串，包含三个字段： card_id：所选发票卡券的 cardId，encrypt_code：所选发票卡券的加密 code，报销方可以通过 cardId 和 encryptCode 获得报销发票的信息，app_id： 发票方的 appId。
         */
        invoiceInfo: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChooseInvoiceSuccessCallback = (
                result: ChooseInvoiceSuccessCallbackResult
            ) => void;

    interface ChooseInvoiceOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChooseInvoiceCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChooseInvoiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ChooseInvoiceSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChooseLicensePlateCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChooseLicensePlateFailCallback = (res: GeneralCallbackResult) => void;

    interface ChooseLicensePlateSuccessCallbackResult {
        /**
         * 用户选择的车牌号
         */
        plateNumber: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChooseLicensePlateSuccessCallback = (
                result: ChooseLicensePlateSuccessCallbackResult
            ) => void;

    interface ChooseLicensePlateOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChooseLicensePlateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChooseLicensePlateFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ChooseLicensePlateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChooseMediaCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChooseMediaFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 本地临时文件列表
     */
    interface MediaFile {
        /**
         * 视频的时间长度
         */
        duration: number;
        /**
         * 文件类型
         *
         * 可选值：
         * - 'image': 图片;
         * - 'video': 视频;
         */
        fileType: "video" | "image";
        /**
         * 视频的高度
         */
        height: number;
        /**
         * 本地临时文件大小，单位 B
         */
        size: number;
        /**
         * 本地临时文件路径 (本地路径)
         */
        tempFilePath: string;
        /**
         * 视频缩略图临时文件路径
         */
        thumbTempFilePath: string;
        /**
         * 视频的宽度
         */
        width: number;
    }

    interface ChooseMediaSuccessCallbackResult {
        /**
         * 本地临时文件列表
         */
        tempFiles: MediaFile[];
        /**
         * 文件类型，有效值有 image 、video、mix
         */
        type: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChooseMediaSuccessCallback = (
                result: ChooseMediaSuccessCallbackResult
            ) => void;

    interface ChooseMediaOption {
        /**
         * 仅在 sourceType 为 camera 时生效，使用前置或后置摄像头
         *
         * 可选值：
         * - 'back': 使用后置摄像头;
         * - 'front': 使用前置摄像头;
         */
        camera?: "back" | "front";
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChooseMediaCompleteCallback;
        /**
         * 最多可以选择的文件个数，基础库2.25.0前，最多可支持9个文件，2.25.0及以后最多可支持20个文件
         */
        count?: number;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChooseMediaFailCallback;
        /**
         * 拍摄视频最长拍摄时间，单位秒。时间范围为 3s 至 60s 之间。不限制相册。
         */
        maxDuration?: number;
        /**
         * 文件类型
         *
         * 可选值：
         * - 'image': 只能拍摄图片或从相册选择图片;
         * - 'video': 只能拍摄视频或从相册选择视频;
         * - 'mix': 可同时选择图片和视频;
         */
        mediaType?: ("video" | "image" | "mix")[];
        /**
         * 是否压缩所选文件，基础库2.25.0前仅对 mediaType 为 image 时有效，2.25.0及以后对全量 mediaType 有效
         */
        sizeType?: string[];
        /**
         * 图片和视频选择的来源
         *
         * 可选值：
         * - 'album': 从相册选择;
         * - 'camera': 使用相机拍摄;
         */
        sourceType?: ("album" | "camera")[];
        /**
         * 接口调用成功的回调函数
         */
        success?: ChooseMediaSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChooseMessageFileCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChooseMessageFileFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 返回选择的文件的本地临时文件对象数组
     */
    interface ChooseFile {
        /**
         * 选择的文件名称
         */
        name: string;
        /**
         * 本地临时文件路径 (本地路径)
         */
        path: string;
        /**
         * 本地临时文件大小，单位 B
         */
        size: number;
        /**
         * 选择的文件的会话发送时间，Unix时间戳，工具暂不支持此属性
         */
        time: number;
        /**
         * 选择的文件类型
         *
         * 可选值：
         * - 'video': 选择了视频文件;
         * - 'image': 选择了图片文件;
         * - 'file': 选择了除图片和视频的文件;
         */
        type: "video" | "image" | "file";
    }

    interface ChooseMessageFileSuccessCallbackResult {
        /**
         * 返回选择的文件的本地临时文件对象数组
         */
        tempFiles: ChooseFile[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChooseMessageFileSuccessCallback = (
                result: ChooseMessageFileSuccessCallbackResult
            ) => void;

    interface ChooseMessageFileOption {
        /**
         * 最多可以选择的文件个数，可以 0～100
         */
        count: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChooseMessageFileCompleteCallback;
        /**
         * 需要基础库： `2.6.0`
         *
         * 根据文件拓展名过滤，仅 type==file 时有效。每一项都不能是空字符串。默认不过滤。
         */
        extension?: string[];
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChooseMessageFileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ChooseMessageFileSuccessCallback;
        /**
         * 所选的文件的类型
         *
         * 可选值：
         * - 'all': 从所有文件选择;
         * - 'video': 只能选择视频文件;
         * - 'image': 只能选择图片文件;
         * - 'file': 可以选择除了图片和视频之外的其它的文件;
         */
        type?: "all" | "video" | "image" | "file";
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ChoosePoiCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ChoosePoiFailCallback = (res: GeneralCallbackResult) => void;

    interface ChoosePoiSuccessCallbackResult {
        /**
         * 详细地址
         */
        address: string;
        /**
         * 城市名称
         */
        city: number;
        /**
         * 纬度，浮点数，范围为-90~90，负数表示南纬。使用 gcj02 国测局坐标系（即将废弃）
         */
        latitude: number;
        /**
         * 经度，浮点数，范围为-180~180，负数表示西经。使用 gcj02 国测局坐标系（即将废弃）
         */
        longitude: number;
        /**
         * 位置名称
         */
        name: string;
        /**
         * 选择城市时，值为 1，选择精确位置时，值为 2
         */
        type: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type ChoosePoiSuccessCallback = (
                result: ChoosePoiSuccessCallbackResult
            ) => void;

    interface ChoosePoiOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ChoosePoiCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ChoosePoiFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ChoosePoiSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type CreateBLEPeripheralServerCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type CreateBLEPeripheralServerFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    /**
     * 描述符的权限
     */
    interface DescriptorPermission {
        /**
         * 读
         */
        read?: boolean;
        /**
         * 写
         */
        write?: boolean;
    }

    /**
     * 描述符数据
     */
    interface CharacteristicDescriptor {
        /**
         * Descriptor 的 UUID
         */
        uuid: string;
        /**
         * 描述符的权限
         */
        permission?: DescriptorPermission;
        /**
         * 描述符数据
         */
        value?: ArrayBuffer;
    }

    /**
     * 特征权限
     */
    interface CharacteristicPermission {
        /**
         * 加密读请求
         */
        readEncryptionRequired?: boolean;
        /**
         * 可读
         */
        readable?: boolean;
        /**
         * 加密写请求
         */
        writeEncryptionRequired?: boolean;
        /**
         * 可写
         */
        writeable?: boolean;
    }

    /**
     * 特征支持的操作
     */
    interface CharacteristicProperties {
        /**
         * 回包
         */
        indicate?: boolean;
        /**
         * 订阅
         */
        notify?: boolean;
        /**
         * 读
         */
        read?: boolean;
        /**
         * 写
         */
        write?: boolean;
        /**
         * 无回复写
         */
        writeNoResponse?: boolean;
    }

    /**
     * characteristics列表
     */
    interface Characteristic {
        /**
         * characteristic 的 UUID
         */
        uuid: string;
        /**
         * 描述符数据
         */
        descriptors?: CharacteristicDescriptor[];
        /**
         * 特征权限
         */
        permission?: CharacteristicPermission;
        /**
         * 特征支持的操作
         */
        properties?: CharacteristicProperties;
        /**
         * 特征对应的二进制值
         */
        value?: ArrayBuffer;
    }

    /**
     * 描述service的Object
     */
    interface BLEPeripheralService {
        /**
         * characteristics列表
         */
        characteristics: Characteristic[];
        /**
         * 蓝牙服务的 UUID
         */
        uuid: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type AddServiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type AddServiceFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type AddServiceSuccessCallback = (res: GeneralCallbackResult) => void;

    interface AddServiceOption {
        /**
         * 描述service的Object
         */
        service: BLEPeripheralService;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: AddServiceCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: AddServiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: AddServiceSuccessCallback;
    }

    /**
     * 接口调用失败的回调函数
     */
    type SocketTaskCloseFailCallback = (res: GeneralCallbackResult) => void;

    interface BLEPeripheralServerCloseOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: FileSystemManagerCloseCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SocketTaskCloseFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: FileSystemManagerCloseSuccessCallback;
    }

    interface OnCharacteristicReadRequestListenerResult {
        /**
         * 唯一标识码，调用 [writeCharacteristicValue](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html) 时使用
         */
        callbackId: number;
        /**
         * 蓝牙特征的 UUID
         */
        characteristicId: string;
        /**
         * 蓝牙特征对应服务的 UUID
         */
        serviceId: string;
    }

    /**
     * onCharacteristicReadRequest 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffCharacteristicReadRequestCallback = (
                result: OnCharacteristicReadRequestListenerResult
            ) => void;

    interface OnCharacteristicSubscribedListenerResult {
        /**
         * 蓝牙特征的 UUID
         */
        characteristicId: string;
        /**
         * 蓝牙特征对应服务的 UUID
         */
        serviceId: string;
    }

    /**
     * onCharacteristicSubscribed 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffCharacteristicSubscribedCallback = (
                result: OnCharacteristicSubscribedListenerResult
            ) => void;
    /**
     * onCharacteristicUnsubscribed 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffCharacteristicUnsubscribedCallback = (
                result: OnCharacteristicSubscribedListenerResult
            ) => void;

    interface OnCharacteristicWriteRequestListenerResult {
        /**
         * 唯一标识码，调用 [writeCharacteristicValue](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html) 时使用
         */
        callbackId: number;
        /**
         * 蓝牙特征的 UUID
         */
        characteristicId: string;
        /**
         * 蓝牙特征对应服务的 UUID
         */
        serviceId: string;
        /**
         * 请求写入特征的二进制数据值
         */
        value: ArrayBuffer;
    }

    /**
     * onCharacteristicWriteRequest 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffCharacteristicWriteRequestCallback = (
                result: OnCharacteristicWriteRequestListenerResult
            ) => void;
    /**
     * 已连接的设备请求读当前外围设备的特征值事件的监听函数
     */
    type OnCharacteristicReadRequestCallback = (
                result: OnCharacteristicReadRequestListenerResult
            ) => void;
    /**
     * 特征订阅事件的监听函数
     */
    type OnCharacteristicSubscribedCallback = (
                result: OnCharacteristicSubscribedListenerResult
            ) => void;
    /**
     * 取消特征订阅事件的监听函数
     */
    type OnCharacteristicUnsubscribedCallback = (
                result: OnCharacteristicSubscribedListenerResult
            ) => void;
    /**
     * 已连接的设备请求写当前外围设备的特征值事件的监听函数
     */
    type OnCharacteristicWriteRequestCallback = (
                result: OnCharacteristicWriteRequestListenerResult
            ) => void;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RemoveServiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type RemoveServiceFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RemoveServiceSuccessCallback = (res: GeneralCallbackResult) => void;

    interface RemoveServiceOption {
        /**
         * service 的 UUID
         */
        serviceId: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RemoveServiceCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RemoveServiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RemoveServiceSuccessCallback;
    }

    /**
     * 需要基础库： `2.20.1`
     *
     * 以 beacon 设备形式广播的参数。
     */
    interface BeaconInfoObj {
        /**
         * Beacon 设备的主 ID
         */
        major: number;
        /**
         * Beacon 设备的次 ID
         */
        minor: number;
        /**
         * Beacon 设备广播的 UUID
         */
        uuid: string;
        /**
         * 用于判断距离设备 1 米时 RSSI 大小的参考值
         */
        measuredPower?: number;
    }

    /**
     * 广播的制造商信息。仅安卓支持，iOS 因系统限制无法定制。
     */
    interface ManufacturerData {
        /**
         * 制造商ID，0x 开头的十六进制
         */
        manufacturerId: string;
        /**
         * 制造商信息
         */
        manufacturerSpecificData?: ArrayBuffer;
    }

    /**
     * 广播自定义参数
     */
    interface AdvertiseReqObj {
        /**
         * 需要基础库： `2.20.1`
         *
         * 以 beacon 设备形式广播的参数。
         */
        beacon?: BeaconInfoObj;
        /**
         * 当前设备是否可连接
         */
        connectable?: boolean;
        /**
         * 广播中 deviceName 字段，默认为空
         */
        deviceName?: string;
        /**
         * 广播的制造商信息。仅安卓支持，iOS 因系统限制无法定制。
         */
        manufacturerData?: ManufacturerData[];
        /**
         * 要广播的服务 UUID 列表。使用 16/32 位 UUID 时请参考注意事项。
         */
        serviceUuids?: string[];
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartAdvertisingCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StartAdvertisingFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StartAdvertisingSuccessCallback = (res: GeneralCallbackResult) => void;

    interface StartAdvertisingObject {
        /**
         * 广播自定义参数
         */
        advertiseRequest: AdvertiseReqObj;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartAdvertisingCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartAdvertisingFailCallback;
        /**
         * 广播功率
         *
         * 可选值：
         * - 'low': 功率低;
         * - 'medium': 功率适中;
         * - 'high': 功率高;
         */
        powerLevel?: "low" | "medium" | "high";
        /**
         * 接口调用成功的回调函数
         */
        success?: StartAdvertisingSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopAdvertisingCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopAdvertisingFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopAdvertisingSuccessCallback = (res: GeneralCallbackResult) => void;

    interface StopAdvertisingOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopAdvertisingCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopAdvertisingFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopAdvertisingSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type WriteCharacteristicValueCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type WriteCharacteristicValueFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type WriteCharacteristicValueSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface WriteCharacteristicValueObject {
        /**
         * 蓝牙特征的 UUID
         */
        characteristicId: string;
        /**
         * 是否需要通知主机 value 已更新
         */
        needNotify: boolean;
        /**
         * 蓝牙特征对应服务的 UUID
         */
        serviceId: string;
        /**
         * characteristic 对应的二进制值
         */
        value: ArrayBuffer;
        /**
         * 可选，处理回包时使用
         */
        callbackId?: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: WriteCharacteristicValueCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: WriteCharacteristicValueFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: WriteCharacteristicValueSuccessCallback;
    }

    interface BLEPeripheralServer {
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 添加服务。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.addService.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.addService.html)
         */
        addService(option: AddServiceOption): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 关闭当前服务端。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.close.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.close.html)
         */
        close(option?: BLEPeripheralServerCloseOption): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 移除已连接的设备请求读当前外围设备的特征值事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicReadRequest.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicReadRequest.html)
         */
        offCharacteristicReadRequest(listener?: OffCharacteristicReadRequestCallback): void;
        /**
         *
         * 需要基础库： `2.13.0`
         *
         * 在插件中使用：不支持
         *
         * 移除特征订阅事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicSubscribed.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicSubscribed.html)
         */
        offCharacteristicSubscribed(listener?: OffCharacteristicSubscribedCallback): void;
        /**
         *
         * 需要基础库： `2.13.0`
         *
         * 在插件中使用：不支持
         *
         * 移除取消特征订阅事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicUnsubscribed.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicUnsubscribed.html)
         */
        offCharacteristicUnsubscribed(listener?: OffCharacteristicUnsubscribedCallback): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 移除已连接的设备请求写当前外围设备的特征值事件的监听函数
         *
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicWriteRequest.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.offCharacteristicWriteRequest.html)
         */
        offCharacteristicWriteRequest(listener?: OffCharacteristicWriteRequestCallback): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 监听已连接的设备请求读当前外围设备的特征值事件。收到该消息后需要立刻调用 [writeCharacteristicValue](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html) 写回数据，否则主机不会收到响应。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicReadRequest.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicReadRequest.html)
         */
        onCharacteristicReadRequest(listener: OnCharacteristicReadRequestCallback): void;
        /**
         *
         * 需要基础库： `2.13.0`
         *
         * 在插件中使用：不支持
         *
         * 监听特征订阅事件，仅 iOS 支持。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicSubscribed.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicSubscribed.html)
         */
        onCharacteristicSubscribed(listener: OnCharacteristicSubscribedCallback): void;
        /**
         *
         * 需要基础库： `2.13.0`
         *
         * 在插件中使用：不支持
         *
         * 监听取消特征订阅事件，仅 iOS 支持。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicUnsubscribed.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicUnsubscribed.html)
         */
        onCharacteristicUnsubscribed(listener: OnCharacteristicUnsubscribedCallback): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 监听已连接的设备请求写当前外围设备的特征值事件。收到该消息后需要立刻调用 [writeCharacteristicValue](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html) 写回数据，否则主机不会收到响应。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicWriteRequest.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.onCharacteristicWriteRequest.html)
         */
        onCharacteristicWriteRequest(listener: OnCharacteristicWriteRequestCallback): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 移除服务。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.removeService.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.removeService.html)
         */
        removeService(option: RemoveServiceOption): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 开始广播本地创建的外围设备。
         *
         * **注意**
         *
         * - Android 8.0.9 开始，支持直接使用 16/32/128 位 UUID；
         * - Android 8.0.9 以下版本只支持 128 位 UUID，使用 16/32 位的 UUID 时需要进行补位（系统会自动识别是否属于预分配区间），可以参考[蓝牙指南](https://developers.weixin.qq.com/miniprogram/dev/framework/device/ble.html)；
         * - iOS 必须直接使用 16 位的 UUID，不能补位到 128 位，否则系统组包时仍会按照 128 位传输。iOS 暂不支持 32 位 UUID。
         * - iOS 同时只能发起一个广播，安卓支持同时发起多个广播。
         * - 传 beacon 参数时，不能同时传入 deviceName，serviceUuids，manufacturerData 参数。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.startAdvertising.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.startAdvertising.html)
         */
        startAdvertising(Object: StartAdvertisingObject): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 停止广播。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.stopAdvertising.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.stopAdvertising.html)
         */
        stopAdvertising(option?: StopAdvertisingOption): void;
        /**
         *
         * 需要基础库： `2.10.3`
         *
         * 在插件中使用：不支持
         *
         * 往指定特征写入二进制数据值，并通知已连接的主机，从机的特征值已发生变化，该接口会处理是走回包还是走订阅。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.writeCharacteristicValue.html)
         */
        writeCharacteristicValue(Object: WriteCharacteristicValueObject): void;
    }

    interface CreateBLEPeripheralServerSuccessCallbackResult {
        /**
         *
         * 外围设备的服务端。
         *
         * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/BLEPeripheralServer.html)
         */
        server: BLEPeripheralServer;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type CreateBLEPeripheralServerSuccessCallback = (
                result: CreateBLEPeripheralServerSuccessCallbackResult
            ) => void;

    interface CreateBLEPeripheralServerOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: CreateBLEPeripheralServerCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: CreateBLEPeripheralServerFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: CreateBLEPeripheralServerSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type CropImageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type CropImageFailCallback = (res: GeneralCallbackResult) => void;

    interface EditImageSuccessCallbackResult {
        /**
         * 编辑后图片的临时文件路径 (本地路径)
         */
        tempFilePath: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type CropImageSuccessCallback = (
                result: EditImageSuccessCallbackResult
            ) => void;

    interface CropImageOption {
        /**
         * 裁剪比例
         *
         * 可选值：
         * - '16:9': 宽高比为16比9;
         * - '9:16': 宽高比为9比16;
         * - '4:3': 宽高比为4比3;
         * - '3:4': 宽高比为3比4;
         * - '5:4': 宽高比为5比4;
         * - '4:5': 宽高比为4比5;
         * - '1:1': 宽高比为1比1;
         */
        cropScale: "16:9" | "9:16" | "4:3" | "3:4" | "5:4" | "4:5" | "1:1";
        /**
         * 图片路径，图片的路径，支持本地路径、代码包路径
         */
        src: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: CropImageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: CropImageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: CropImageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type DisableAlertBeforeUnloadCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type DisableAlertBeforeUnloadFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type DisableAlertBeforeUnloadSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface DisableAlertBeforeUnloadOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: DisableAlertBeforeUnloadCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: DisableAlertBeforeUnloadFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: DisableAlertBeforeUnloadSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type EditImageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type EditImageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type EditImageSuccessCallback = (
                result: EditImageSuccessCallbackResult
            ) => void;

    interface EditImageOption {
        /**
         * 图片路径，图片的路径，支持本地路径、代码包路径
         */
        src: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: EditImageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: EditImageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: EditImageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type EnableAlertBeforeUnloadCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type EnableAlertBeforeUnloadFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type EnableAlertBeforeUnloadSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface EnableAlertBeforeUnloadOption {
        /**
         * 询问对话框内容
         */
        message: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: EnableAlertBeforeUnloadCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: EnableAlertBeforeUnloadFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: EnableAlertBeforeUnloadSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ExitMiniProgramCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ExitMiniProgramFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ExitMiniProgramSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ExitMiniProgramOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ExitMiniProgramCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ExitMiniProgramFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ExitMiniProgramSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ExitVoIPChatCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ExitVoIPChatFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ExitVoIPChatSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ExitVoIPChatOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ExitVoIPChatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ExitVoIPChatFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ExitVoIPChatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type FaceDetectCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type FaceDetectFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 人脸角度信息，取值范围 [-1, 1]，数值越接近 0 表示越正对摄像头
     */
    interface FaceAngel {
        /**
         * 仰俯角（点头）
         */
        pitch: number;
        /**
         * 翻滚角（左右倾）
         */
        roll: number;
        /**
         * 偏航角（摇头）
         */
        yaw: number;
    }

    /**
     * 人脸置信度，取值范围 [0, 1]，数值越大置信度越高（遮挡越少）
     */
    interface FaceConf {
        /**
         * 整体可信度
         */
        global: number;
        /**
         * 左眼可信度
         */
        leftEye: number;
        /**
         * 嘴巴可信度
         */
        mouth: number;
        /**
         * 鼻子可信度
         */
        nose: number;
        /**
         * 右眼可信度
         */
        rightEye: number;
    }

    interface FaceDetectSuccessCallbackResult {
        /**
         * 人脸角度信息，取值范围 [-1, 1]，数值越接近 0 表示越正对摄像头
         */
        angleArray: FaceAngel;
        /**
         * 人脸置信度，取值范围 [0, 1]，数值越大置信度越高（遮挡越少）
         */
        confArray: FaceConf;
        /**
         * 脸部方框数值，对象包含 height, weight, originX, originY 四个属性 (origin 为方框左上角坐标)
         */
        detectRect: any;
        /**
         * 多人模式（enableMultiFace）下的人脸信息，每个对象包含上述其它属性
         */
        faceInfo: any;
        /**
         * 标记人脸轮廓的 106 个点位置数组，数组每个对象包含 x 和 y
         */
        pointArray: any;
        /**
         * 脸部中心点横坐标，检测不到人脸则为 -1
         */
        x: number;
        /**
         * 脸部中心点纵坐标，检测不到人脸则为 -1
         */
        y: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type FaceDetectSuccessCallback = (
                result: FaceDetectSuccessCallbackResult
            ) => void;

    interface FaceDetectOption {
        /**
         * 图像像素点数据，每四项表示一个像素点的 RGBA
         */
        frameBuffer: ArrayBuffer;
        /**
         * 图像高度
         */
        height: number;
        /**
         * 图像宽度
         */
        width: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: FaceDetectCompleteCallback;
        /**
         * 是否返回当前图像的人脸角度信息
         */
        enableAngle?: boolean;
        /**
         * 是否返回当前图像的人脸的置信度（可表示器官遮挡情况）
         */
        enableConf?: boolean;
        /**
         * 是否返回多张人脸的信息
         */
        enableMultiFace?: boolean;
        /**
         * 是否返回当前图像的人脸（106 个点）
         */
        enablePoint?: boolean;
        /**
         * 接口调用失败的回调函数
         */
        fail?: FaceDetectFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: FaceDetectSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetAvailableAudioSourcesCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetAvailableAudioSourcesFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetAvailableAudioSourcesSuccessCallbackResult {
        /**
         * 支持的音频输入源列表，可在 [RecorderManager.start()](https://developers.weixin.qq.com/miniprogram/dev/api/media/recorder/RecorderManager.start.html) 接口中使用。返回值定义参考 https://developer.android.com/reference/kotlin/android/media/MediaRecorder.AudioSource
         *
         * 可选值：
         * - 'auto': 自动设置，默认使用手机麦克风，插上耳麦后自动切换使用耳机麦克风，所有平台适用;
         * - 'buildInMic': 手机麦克风，仅限 iOS;
         * - 'headsetMic': 耳机麦克风，仅限 iOS;
         * - 'mic': 麦克风（没插耳麦时是手机麦克风，插耳麦时是耳机麦克风），仅限 Android;
         * - 'camcorder': 同 mic，适用于录制音视频内容，仅限 Android;
         * - 'voice_communication': 同 mic，适用于实时沟通，仅限 Android;
         * - 'voice_recognition': 同 mic，适用于语音识别，仅限 Android;
         */
        audioSources: ("auto" | "buildInMic" | "headsetMic" | "mic" | "camcorder" | "voice_communication" | "voice_recognition")[];
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetAvailableAudioSourcesSuccessCallback = (
                result: GetAvailableAudioSourcesSuccessCallbackResult
            ) => void;

    interface GetAvailableAudioSourcesOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetAvailableAudioSourcesCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetAvailableAudioSourcesFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetAvailableAudioSourcesSuccessCallback;
    }

    interface BluetoothError {
        /**
         * 错误信息
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | 0 | ok | 正常 |
         * | -1 | already connect | 已连接 |
         * | 10000 | not init | 未初始化蓝牙适配器 |
         * | 10001 | not available | 当前蓝牙适配器不可用 |
         * | 10002 | no device | 没有找到指定设备 |
         * | 10003 | connection fail | 连接失败 |
         * | 10004 | no service | 没有找到指定服务 |
         * | 10005 | no characteristic | 没有找到指定特征 |
         * | 10006 | no connection | 当前连接已断开 |
         * | 10007 | property not support | 当前特征不支持此操作 |
         * | 10008 | system error | 其余所有系统上报的异常 |
         * | 10009 | system not support | Android 系统特有，系统版本低于 4.3 不支持 BLE |
         * | 10012 | operate time out | 连接超时 |
         * | 10013 | invalid_data | 连接 deviceId 为空或者是格式不正确 |
         */
        errMsg: string;
        /**
         * 错误码
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | 0 | ok | 正常 |
         * | -1 | already connect | 已连接 |
         * | 10000 | not init | 未初始化蓝牙适配器 |
         * | 10001 | not available | 当前蓝牙适配器不可用 |
         * | 10002 | no device | 没有找到指定设备 |
         * | 10003 | connection fail | 连接失败 |
         * | 10004 | no service | 没有找到指定服务 |
         * | 10005 | no characteristic | 没有找到指定特征 |
         * | 10006 | no connection | 当前连接已断开 |
         * | 10007 | property not support | 当前特征不支持此操作 |
         * | 10008 | system error | 其余所有系统上报的异常 |
         * | 10009 | system not support | Android 系统特有，系统版本低于 4.3 不支持 BLE |
         * | 10012 | operate time out | 连接超时 |
         * | 10013 | invalid_data | 连接 deviceId 为空或者是格式不正确 |
         */
        errCode: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetBLEMTUCompleteCallback = (res: BluetoothError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetBLEMTUFailCallback = (res: BluetoothError) => void;

    interface GetBLEMTUSuccessCallbackResult {
        /**
         * 最大传输单元
         */
        mtu: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetBLEMTUSuccessCallback = (
                result: GetBLEMTUSuccessCallbackResult
            ) => void;

    interface GetBLEMTUOption {
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetBLEMTUCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetBLEMTUFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetBLEMTUSuccessCallback;
        /**
         * 写模式 （iOS 特有参数）
         *
         * 可选值：
         * - 'write': 有回复写;
         * - 'writeNoResponse': 无回复写;
         */
        writeType?: "write" | "writeNoResponse";
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetBackgroundAudioPlayerStateCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetBackgroundAudioPlayerStateFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetBackgroundAudioPlayerStateSuccessCallbackResult {
        /**
         * 选定音频的播放位置（单位：s），只有在音乐播放中时返回
         */
        currentPosition: number;
        /**
         * 歌曲数据链接，只有在音乐播放中时返回
         */
        dataUrl: string;
        /**
         * 音频的下载进度百分比，只有在音乐播放中时返回
         */
        downloadPercent: number;
        /**
         * 选定音频的长度（单位：s），只有在音乐播放中时返回
         */
        duration: number;
        /**
         * 播放状态
         *
         * 可选值：
         * - 0: 暂停中;
         * - 1: 播放中;
         * - 2: 没有音乐播放;
         */
        status: 0 | 1 | 2;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetBackgroundAudioPlayerStateSuccessCallback = (
                result: GetBackgroundAudioPlayerStateSuccessCallbackResult
            ) => void;

    interface GetBackgroundAudioPlayerStateOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetBackgroundAudioPlayerStateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetBackgroundAudioPlayerStateFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetBackgroundAudioPlayerStateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetBackgroundFetchDataCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetBackgroundFetchDataFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetBackgroundFetchDataSuccessCallbackResult {
        /**
         * 缓存数据
         */
        fetchedData: string;
        /**
         * 小程序页面路径
         */
        path: string;
        /**
         * 传给页面的 query 参数
         */
        query: string;
        /**
         * 进入小程序的场景值
         */
        scene: number;
        /**
         * 客户端拿到缓存数据的时间戳 ms。(iOS 时间戳存在异常，8.0.27 修复)
         */
        timeStamp: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetBackgroundFetchDataSuccessCallback = (
                result: GetBackgroundFetchDataSuccessCallbackResult
            ) => void;

    interface GetBackgroundFetchDataOption {
        /**
         * 缓存数据类别，取值为 periodic 或 pre
         */
        fetchType: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetBackgroundFetchDataCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetBackgroundFetchDataFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetBackgroundFetchDataSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetBackgroundFetchTokenCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetBackgroundFetchTokenFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetBackgroundFetchTokenSuccessCallbackResult {
        /**
         * 接口调用结果
         */
        errMsg: string;
        /**
         * 自定义的登录态
         */
        token: number;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetBackgroundFetchTokenSuccessCallback = (
                result: GetBackgroundFetchTokenSuccessCallbackResult
            ) => void;

    interface GetBackgroundFetchTokenOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetBackgroundFetchTokenCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetBackgroundFetchTokenFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetBackgroundFetchTokenSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetChannelsLiveInfoCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetChannelsLiveInfoFailCallback = (res: GeneralCallbackResult) => void;

    interface GetChannelsLiveInfoSuccessCallbackResult {
        /**
         * 直播主题
         */
        description: string;
        /**
         * 直播 feedId
         */
        feedId: string;
        /**
         * 视频号头像
         */
        headUrl: string;
        /**
         * 视频号昵称
         */
        nickname: string;
        /**
         * 直播 nonceId
         */
        nonceId: string;
        /**
         * 直播状态，2直播中，3直播结束
         */
        status: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetChannelsLiveInfoSuccessCallback = (
                result: GetChannelsLiveInfoSuccessCallbackResult
            ) => void;

    interface GetChannelsLiveInfoOption {
        /**
         * 视频号 id，以“sph”开头的id，可在视频号助手获取
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetChannelsLiveInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetChannelsLiveInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetChannelsLiveInfoSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetChannelsLiveNoticeInfoCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetChannelsLiveNoticeInfoFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetChannelsLiveNoticeInfoSuccessCallbackResult {
        /**
         * 直播封面
         */
        headUrl: string;
        /**
         * 视频号昵称
         */
        nickname: string;
        /**
         * 预告 id
         */
        noticeId: string;
        /**
         * 需要基础库： `2.24.6`
         *
         * 除最近的一条预告信息外，其他的预告信息列表（注意：每次最多返回按时间戳增序排列的15个预告信息，其中时间最近的那个预告信息会在接口其他的返回参数中展示，其余的预告信息会在该字段中展示）。
         */
        otherInfos: any[];
        /**
         * 是否可预约
         */
        reservable: boolean;
        /**
         * 开始时间
         */
        startTime: string;
        /**
         * 预告状态：0可用 1取消 2已用
         */
        status: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetChannelsLiveNoticeInfoSuccessCallback = (
                result: GetChannelsLiveNoticeInfoSuccessCallbackResult
            ) => void;

    interface GetChannelsLiveNoticeInfoOption {
        /**
         * 视频号 id，以“sph”开头的id，可在视频号助手获取
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetChannelsLiveNoticeInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetChannelsLiveNoticeInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetChannelsLiveNoticeInfoSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetChannelsShareKeyCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetChannelsShareKeyFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 推广员
     */
    interface PromoterResult {
        /**
         * 推广员昵称
         */
        finderNickname: string;
        /**
         * 推广员 id
         */
        promoterId: string;
        /**
         * 推广员 openid
         */
        promoterOpenId: string;
    }

    interface GetChannelsShareKeySuccessCallbackResult {
        /**
         * 推广员
         */
        promoter: PromoterResult;
        /**
         * 分享者 openid
         */
        sharerOpenId: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetChannelsShareKeySuccessCallback = (
                result: GetChannelsShareKeySuccessCallbackResult
            ) => void;

    interface GetChannelsShareKeyOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetChannelsShareKeyCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetChannelsShareKeyFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetChannelsShareKeySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetFuzzyLocationCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetFuzzyLocationFailCallback = (res: GeneralCallbackResult) => void;

    interface GetFuzzyLocationSuccessCallbackResult {
        /**
         * 纬度，范围为 -90~90，负数表示南纬
         */
        latitude: number;
        /**
         * 经度，范围为 -180~180，负数表示西经
         */
        longitude: number;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetFuzzyLocationSuccessCallback = (
                result: GetFuzzyLocationSuccessCallbackResult
            ) => void;

    interface GetFuzzyLocationOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetFuzzyLocationCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetFuzzyLocationFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetFuzzyLocationSuccessCallback;
        /**
         * wgs84 返回 gps 坐标，gcj02 返回可用于 wx.openLocation 的坐标
         */
        type?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetGroupEnterInfoCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetGroupEnterInfoFailCallback = (res: GeneralCallbackResult) => void;

    interface GetGroupEnterInfoSuccessCallbackResult {
        /**
         * 需要基础库： `2.7.0`
         *
         * 敏感数据对应的云 ID，开通[云开发](https://developers.weixin.qq.com/miniprogram/dev/wxcloud/basis/getting-started.html)的小程序才会返回，可通过云调用直接获取开放数据，详细见[云调用直接获取开放数据](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#method-cloud)
         */
        cloudID: string;
        /**
         * 包括敏感数据在内的完整转发信息的加密数据，详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
         */
        encryptedData: string;
        /**
         * 错误信息
         */
        errMsg: string;
        /**
         * 加密算法的初始向量，详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
         */
        iv: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetGroupEnterInfoSuccessCallback = (
                result: GetGroupEnterInfoSuccessCallbackResult
            ) => void;

    interface GetGroupEnterInfoOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetGroupEnterInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetGroupEnterInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetGroupEnterInfoSuccessCallback;
    }

    interface NFCError {
        /**
         * 错误信息
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | 0 | ok | 正常 |
         * | 13000 |  | 当前设备不支持NFC |
         * | 13001 |  | 当前设备支持NFC，但系统NFC开关未开启 |
         * | 13002 |  | 当前设备支持NFC，但不支持HCE |
         * | 13003 |  | AID列表参数格式错误 |
         * | 13004 |  | 未设置微信为默认NFC支付应用 |
         * | 13005 |  | 返回的指令不合法 |
         * | 13006 |  | 注册AID失败 |
         */
        errMsg: string;
        /**
         * 错误码
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | 0 | ok | 正常 |
         * | 13000 |  | 当前设备不支持NFC |
         * | 13001 |  | 当前设备支持NFC，但系统NFC开关未开启 |
         * | 13002 |  | 当前设备支持NFC，但不支持HCE |
         * | 13003 |  | AID列表参数格式错误 |
         * | 13004 |  | 未设置微信为默认NFC支付应用 |
         * | 13005 |  | 返回的指令不合法 |
         * | 13006 |  | 注册AID失败 |
         */
        errCode: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetHCEStateCompleteCallback = (res: NFCError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetHCEStateFailCallback = (res: NFCError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type GetHCEStateSuccessCallback = (res: NFCError) => void;

    interface GetHCEStateOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetHCEStateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetHCEStateFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetHCEStateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetLocalIPAddressCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetLocalIPAddressFailCallback = (res: GeneralCallbackResult) => void;

    interface GetLocalIPAddressSuccessCallbackResult {
        /**
         * 错误信息
         */
        errMsg: string;
        /**
         * 本机局域网IP地址
         */
        localip: string;
        /**
         * 本机局域网子网掩码，基础库 2.24.0 开始支持
         */
        netmask: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetLocalIPAddressSuccessCallback = (
                result: GetLocalIPAddressSuccessCallbackResult
            ) => void;

    interface GetLocalIPAddressOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetLocalIPAddressCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetLocalIPAddressFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetLocalIPAddressSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetScreenRecordingStateCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetScreenRecordingStateFailCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface GetScreenRecordingStateSuccessCallbackResult {
        /**
         * 录屏状态
         *
         * 可选值：
         * - 'on': 开启;
         * - 'off': 关闭;
         */
        state: "on" | "off";
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetScreenRecordingStateSuccessCallback = (
                result: GetScreenRecordingStateSuccessCallbackResult
            ) => void;

    interface GetScreenRecordingStateOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetScreenRecordingStateCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetScreenRecordingStateFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetScreenRecordingStateSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetShareInfoCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetShareInfoFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type GetShareInfoSuccessCallback = (
                result: GetGroupEnterInfoSuccessCallbackResult
            ) => void;

    interface GetShareInfoOption {
        /**
         * shareTicket
         */
        shareTicket: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetShareInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetShareInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetShareInfoSuccessCallback;
        /**
         * 需要基础库： `1.9.90`
         *
         * 超时时间，单位 ms
         */
        timeout?: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetSkylineInfoCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetSkylineInfoFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type GetSkylineInfoSuccessCallback = (
                /** 当前运行环境对于 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的支持情况 */
                result: SkylineInfo
            ) => void;

    interface GetSkylineInfoOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetSkylineInfoCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetSkylineInfoFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetSkylineInfoSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetSystemInfoAsyncCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetSystemInfoAsyncFailCallback = (res: GeneralCallbackResult) => void;

    /**
     * 需要基础库： `2.12.3`
     *
     * 当前小程序运行的宿主环境
     */
    interface SystemInfoHost {
        /**
         * 宿主 app 对应的 appId
         */
        appId: string;
    }

    interface SystemInfo {
        /**
         * 需要基础库： `1.1.0`
         *
         * 客户端基础库版本
         */
        SDKVersion: string;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信使用相册的开关（仅 iOS 有效）
         */
        albumAuthorized: boolean;
        /**
         * 需要基础库： `1.8.0`
         *
         * 设备性能等级（仅 Android）。取值为：-2 或 0（该设备无法运行小游戏），-1（性能未知），>=1（设备性能值，该值越高，设备性能越好，目前最高不到50）
         */
        benchmarkLevel: number;
        /**
         * 需要基础库： `2.6.0`
         *
         * 蓝牙的系统开关
         */
        bluetoothEnabled: boolean;
        /**
         * 需要基础库： `1.5.0`
         *
         * 设备品牌
         */
        brand: string;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信使用摄像头的开关
         */
        cameraAuthorized: boolean;
        /**
         * 设备方向
         *
         * 可选值：
         * - 'portrait': 竖屏;
         * - 'landscape': 横屏;
         */
        deviceOrientation: "portrait" | "landscape";
        /**
         * 需要基础库： `2.15.0`
         *
         * 是否已打开调试。可通过右上角菜单或 [uni.setEnableDebug](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.setEnableDebug.html) 打开调试。
         */
        enableDebug: boolean;
        /**
         * 需要基础库： `1.5.0`
         *
         * 用户字体大小（单位px）。以微信客户端「我-设置-通用-字体大小」中的设置为准
         */
        fontSizeSetting: number;
        /**
         * 需要基础库： `2.12.3`
         *
         * 当前小程序运行的宿主环境
         */
        host: SystemInfoHost;
        /**
         * 微信设置的语言
         */
        language: string;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信使用定位的开关
         */
        locationAuthorized: boolean;
        /**
         * 需要基础库： `2.6.0`
         *
         * 地理位置的系统开关
         */
        locationEnabled: boolean;
        /**
         * `true` 表示模糊定位，`false` 表示精确定位，仅 iOS 支持
         */
        locationReducedAccuracy: boolean;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信使用麦克风的开关
         */
        microphoneAuthorized: boolean;
        /**
         * 设备型号。新机型刚推出一段时间会显示unknown，微信会尽快进行适配。
         */
        model: string;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信通知带有提醒的开关（仅 iOS 有效）
         */
        notificationAlertAuthorized: boolean;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信通知的开关
         */
        notificationAuthorized: boolean;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信通知带有标记的开关（仅 iOS 有效）
         */
        notificationBadgeAuthorized: boolean;
        /**
         * 需要基础库： `2.6.0`
         *
         * 允许微信通知带有声音的开关（仅 iOS 有效）
         */
        notificationSoundAuthorized: boolean;
        /**
         * 需要基础库： `2.19.3`
         *
         * 允许微信使用日历的开关
         */
        phoneCalendarAuthorized: boolean;
        /**
         * 设备像素比
         */
        pixelRatio: number;
        /**
         * 客户端平台
         *
         * 可选值：
         * - 'ios': iOS微信（包含 iPhone、iPad）;
         * - 'android': Android微信;
         * - 'windows': Windows微信;
         * - 'mac': macOS微信;
         * - 'devtools': 微信开发者工具;
         */
        platform: "ios" | "android" | "windows" | "mac" | "devtools";
        /**
         * 需要基础库： `2.7.0`
         *
         * 在竖屏正方向下的安全区域
         */
        safeArea: SafeArea;
        /**
         * 需要基础库： `1.1.0`
         *
         * 屏幕高度，单位px
         */
        screenHeight: number;
        /**
         * 需要基础库： `1.1.0`
         *
         * 屏幕宽度，单位px
         */
        screenWidth: number;
        /**
         * 需要基础库： `1.9.0`
         *
         * 状态栏的高度，单位px
         */
        statusBarHeight: number;
        /**
         * 操作系统及版本
         */
        system: string;
        /**
         * 微信版本号
         */
        version: string;
        /**
         * 需要基础库： `2.6.0`
         *
         * Wi-Fi 的系统开关
         */
        wifiEnabled: boolean;
        /**
         * 可使用窗口高度，单位px
         */
        windowHeight: number;
        /**
         * 可使用窗口宽度，单位px
         */
        windowWidth: number;
        /**
         * 需要基础库： `2.11.0`
         *
         * 系统当前主题，取值为`light`或`dark`，全局配置`"darkmode":true`时才能获取，否则为 undefined （不支持小游戏）
         *
         * 可选值：
         * - 'dark': 深色主题;
         * - 'light': 浅色主题;
         */
        theme?: "dark" | "light";
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetSystemInfoAsyncSuccessCallback = (result: SystemInfo) => void;

    interface GetSystemInfoAsyncOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetSystemInfoAsyncCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetSystemInfoAsyncFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetSystemInfoAsyncSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type GetWeRunDataCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type GetWeRunDataFailCallback = (res: GeneralCallbackResult) => void;

    interface GetWeRunDataSuccessCallbackResult {
        /**
         * 需要基础库： `2.7.0`
         *
         * 敏感数据对应的云 ID，开通云开发的小程序才会返回，可通过云调用直接获取开放数据，详细见[云调用直接获取开放数据](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html#method-cloud)
         */
        cloudID: string;
        /**
         * 包括敏感数据在内的完整用户信息的加密数据，详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)。解密后得到的数据结构见后文
         */
        encryptedData: string;
        /**
         * 加密算法的初始向量，详细见[加密数据解密算法](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/signature.html)
         */
        iv: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type GetWeRunDataSuccessCallback = (
                result: GetWeRunDataSuccessCallbackResult
            ) => void;

    interface GetWeRunDataOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: GetWeRunDataCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: GetWeRunDataFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: GetWeRunDataSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type InitFaceDetectCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type InitFaceDetectFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type InitFaceDetectSuccessCallback = (res: GeneralCallbackResult) => void;

    interface InitFaceDetectOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: InitFaceDetectCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: InitFaceDetectFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: InitFaceDetectSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type IsBluetoothDevicePairedCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type IsBluetoothDevicePairedFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type IsBluetoothDevicePairedSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface IsBluetoothDevicePairedOption {
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: IsBluetoothDevicePairedCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: IsBluetoothDevicePairedFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: IsBluetoothDevicePairedSuccessCallback;
    }

    interface VoIP1v1ChatUser {
        /**
         * 昵称
         */
        nickname: string;
        /**
         * 小程序内 openid
         */
        openid: string;
        /**
         * 头像
         */
        headImage?: string;
    }

    interface Join1v1ChatError {
        /**
         * 错误信息
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | -20000 | not open 1v1 Chat | 未开通双人通话 |
         * | -20001 | device not support | 当前设备不支持 |
         * | -20002 | on call | 正在通话中 |
         * | -20003 | occupied by other miniprogram | 其它小程序正在通话中 |
         * | -30000 | system error | 内部系统错误 |
         * | -30001 | wechat has no camera authorization | 微信缺失相机权限 |
         * | -30002 | wechat has no record authorization | 微信缺失录音权限 |
         * | -30003 | miniprogram has no record authorization | 小程序缺失录音权限 |
         * | -30004 | miniprogram has no camera authorization | 小程序缺失相机权限 |
         * | -1 |  | 当前已在房间内 |
         * | -2 |  | 录音设备被占用，可能是当前正在使用微信内语音通话或系统通话 |
         * | -3 |  | 加入会话期间退出（可能是用户主动退出，或者退后台、来电等原因），因此加入失败 |
         * | -1000 |  | 系统错误 |
         */
        errMsg: string;
        /**
         * 错误码
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | -20000 | not open 1v1 Chat | 未开通双人通话 |
         * | -20001 | device not support | 当前设备不支持 |
         * | -20002 | on call | 正在通话中 |
         * | -20003 | occupied by other miniprogram | 其它小程序正在通话中 |
         * | -30000 | system error | 内部系统错误 |
         * | -30001 | wechat has no camera authorization | 微信缺失相机权限 |
         * | -30002 | wechat has no record authorization | 微信缺失录音权限 |
         * | -30003 | miniprogram has no record authorization | 小程序缺失录音权限 |
         * | -30004 | miniprogram has no camera authorization | 小程序缺失相机权限 |
         * | -1 |  | 当前已在房间内 |
         * | -2 |  | 录音设备被占用，可能是当前正在使用微信内语音通话或系统通话 |
         * | -3 |  | 加入会话期间退出（可能是用户主动退出，或者退后台、来电等原因），因此加入失败 |
         * | -1000 |  | 系统错误 |
         */
        errCode: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type Join1v1ChatCompleteCallback = (res: Join1v1ChatError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type Join1v1ChatFailCallback = (res: Join1v1ChatError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type Join1v1ChatSuccessCallback = (res: Join1v1ChatError) => void;

    interface Join1v1ChatOption {
        /**
         * 呼叫方信息
         */
        caller: VoIP1v1ChatUser;
        /**
         * 接听方信息
         */
        listener: VoIP1v1ChatUser;
        /**
         * 窗口背景色(音频通话背景以及小窗模式背景)
         *
         * 可选值：
         * - 0: #262930;
         * - 1: #FA5151;
         * - 2: #FA9D3B;
         * - 3: #3D7257;
         * - 4: #1485EE;
         * - 5: #6467F0;
         */
        backgroundType?: 0 | 1 | 2 | 3 | 4 | 5;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: Join1v1ChatCompleteCallback;
        /**
         * 不允许切换到语音通话
         */
        disableSwitchVoice?: boolean;
        /**
         * 接口调用失败的回调函数
         */
        fail?: Join1v1ChatFailCallback;
        /**
         * 小窗样式
         */
        minWindowType?: number;
        /**
         * 通话类型
         *
         * 可选值：
         * - 'voice': 语音通话;
         * - 'video': 视频通话;
         */
        roomType?: "video" | "voice";
        /**
         * 接口调用成功的回调函数
         */
        success?: Join1v1ChatSuccessCallback;
    }

    interface JoinVoIPChatError {
        /**
         * 错误信息
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | -1 | 当前已在房间内 |  |
         * | -2 | 录音设备被占用，可能是当前正在使用微信内语音通话或系统通话 |  |
         * | -3 | 加入会话期间退出（可能是用户主动退出，或者退后台、来电等原因），因此加入失败 |  |
         * | -1000 | 系统错误 |  |
         */
        errMsg: string;
        /**
         * 错误码
         *
         * | 错误码 | 错误信息 | 说明 |
         * | - | - | - |
         * | -1 | 当前已在房间内 |  |
         * | -2 | 录音设备被占用，可能是当前正在使用微信内语音通话或系统通话 |  |
         * | -3 | 加入会话期间退出（可能是用户主动退出，或者退后台、来电等原因），因此加入失败 |  |
         * | -1000 | 系统错误 |  |
         */
        errCode: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type JoinVoIPChatCompleteCallback = (res: JoinVoIPChatError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type JoinVoIPChatFailCallback = (res: JoinVoIPChatError) => void;

    /**
     * 静音设置
     */
    interface MuteConfig {
        /**
         * 是否静音耳机
         */
        muteEarphone?: boolean;
        /**
         * 是否静音麦克风
         */
        muteMicrophone?: boolean;
    }

    interface JoinVoIPChatSuccessCallbackResult {
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果
         */
        errMsg: string;
        /**
         * 在此通话中的成员 openId 名单
         */
        openIdList: string[];
    }

    /**
     * 接口调用成功的回调函数
     */
    type JoinVoIPChatSuccessCallback = (
                result: JoinVoIPChatSuccessCallbackResult
            ) => void;

    interface JoinVoIPChatOption {
        /**
         * 小游戏内此房间/群聊的 ID。同一时刻传入相同 groupId 的用户会进入到同个实时语音房间。
         */
        groupId: string;
        /**
         * 验证所需的随机字符串
         */
        nonceStr: string;
        /**
         * 签名，用于验证小游戏的身份
         */
        signature: string;
        /**
         * 验证所需的时间戳
         */
        timeStamp: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: JoinVoIPChatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: JoinVoIPChatFailCallback;
        /**
         * 静音设置
         */
        muteConfig?: MuteConfig;
        /**
         * 房间类型
         *
         * 可选值：
         * - 'voice': 音频房间，用于语音通话;
         * - 'video': 视频房间，结合 [voip-room](https://developers.weixin.qq.com/miniprogram/dev/component/voip-room.html) 组件可显示成员画面;
         */
        roomType?: "video" | "voice";
        /**
         * 接口调用成功的回调函数
         */
        success?: JoinVoIPChatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type MakeBluetoothPairCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type MakeBluetoothPairFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type MakeBluetoothPairSuccessCallback = (res: GeneralCallbackResult) => void;

    interface MakeBluetoothPairOption {
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
        /**
         * pin 码，Base64 格式。
         */
        pin: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: MakeBluetoothPairCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: MakeBluetoothPairFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: MakeBluetoothPairSuccessCallback;
        /**
         * 超时时间，单位 ms
         */
        timeout?: number;
    }

    /**
     * onAppHide 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffAppHideCallback = (res: GeneralCallbackResult) => void;
    /**
     * onAppShow 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffAppShowCallback = (res: GeneralCallbackResult) => void;
    /**
     * onAudioInterruptionBegin 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffAudioInterruptionBeginCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * onAudioInterruptionEnd 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffAudioInterruptionEndCallback = (res: GeneralCallbackResult) => void;

    interface OnBLECharacteristicValueChangeListenerResult {
        /**
         * 蓝牙特征的 UUID
         */
        characteristicId: string;
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
        /**
         * 蓝牙特征对应服务的 UUID
         */
        serviceId: string;
        /**
         * 特征最新的值
         */
        value: ArrayBuffer;
    }

    /**
     * onBLECharacteristicValueChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBLECharacteristicValueChangeCallback = (
                result: OnBLECharacteristicValueChangeListenerResult
            ) => void;

    interface OnBLEConnectionStateChangeListenerResult {
        /**
         * 是否处于已连接状态
         */
        connected: boolean;
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
    }

    /**
     * onBLEConnectionStateChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBLEConnectionStateChangeCallback = (
                result: OnBLEConnectionStateChangeListenerResult
            ) => void;

    interface OnBLEMTUChangeListenerResult {
        /**
         * 蓝牙设备 id
         */
        deviceId: string;
        /**
         * 最大传输单元
         */
        mtu: number;
    }

    /**
     * onBLEMTUChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBLEMTUChangeCallback = (
                result: OnBLEMTUChangeListenerResult
            ) => void;

    interface OnBLEPeripheralConnectionStateChangedListenerResult {
        /**
         * 连接目前状态
         */
        connected: boolean;
        /**
         * 连接状态变化的设备 id
         */
        deviceId: string;
        /**
         * server 的 UUID
         */
        serverId: string;
    }

    /**
     * onBLEPeripheralConnectionStateChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBLEPeripheralConnectionStateChangedCallback = (
                result: OnBLEPeripheralConnectionStateChangedListenerResult
            ) => void;
    /**
     * onBeaconServiceChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBeaconServiceChangeCallback = (res: GeneralCallbackResult) => void;
    /**
     * onBeaconUpdate 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBeaconUpdateCallback = (res: GeneralCallbackResult) => void;

    interface OnBluetoothAdapterStateChangeListenerResult {
        /**
         * 蓝牙适配器是否可用
         */
        available: boolean;
        /**
         * 蓝牙适配器是否处于搜索状态
         */
        discovering: boolean;
    }

    /**
     * onBluetoothAdapterStateChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBluetoothAdapterStateChangeCallback = (
                result: OnBluetoothAdapterStateChangeListenerResult
            ) => void;
    /**
     * onBluetoothDeviceFound 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffBluetoothDeviceFoundCallback = (res: GeneralCallbackResult) => void;

    interface OnCopyUrlListenerResult {
        /**
         * 用短链打开小程序时当前页面携带的查询字符串。小程序中使用时，应在进入页面时调用 `wx.onCopyUrl` 自定义 `query`，退出页面时调用 `wx.offCopyUrl`，防止影响其它页面。
         */
        query: string;
    }

    /**
     * onCopyUrl 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffCopyUrlCallback = (result: OnCopyUrlListenerResult) => void;
    /**
     * onDeviceMotionChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffDeviceMotionChangeCallback = (res: GeneralCallbackResult) => void;

    interface OnHCEMessageListenerResult {
        /**
         * `messageType=1` 时 ,客户端接收到 NFC 设备的指令
         */
        data: ArrayBuffer;
        /**
         * 消息类型
         *
         * 可选值：
         * - 1: HCE APDU Command类型，小程序需对此指令进行处理，并调用 sendHCEMessage 接口返回处理指令;
         * - 2: 设备离场事件类型;
         */
        messageType: 1 | 2;
        /**
         * `messageType=2` 时，原因
         */
        reason: number;
    }

    /**
     * onHCEMessage 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffHCEMessageCallback = (result: OnHCEMessageListenerResult) => void;

    interface OnLazyLoadErrorListenerResult {
        /**
         * 详细信息
         */
        errMsg: string;
        /**
         * 异步组件所属的分包
         */
        subpackage: any[];
        /**
         * 'subpackage' 失败类型
         */
        type: string;
    }

    /**
     * onLazyLoadError 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffLazyLoadErrorCallback = (
                result: OnLazyLoadErrorListenerResult
            ) => void;
    /**
     * onLocalServiceDiscoveryStop 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffLocalServiceDiscoveryStopCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OnLocalServiceFoundListenerResult {
        /**
         * 服务的 ip 地址
         */
        ip: string;
        /**
         * 服务的端口
         */
        port: number;
        /**
         * 服务的名称
         */
        serviceName: string;
        /**
         * 服务的类型
         */
        serviceType: string;
    }

    /**
     * onLocalServiceFound 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffLocalServiceFoundCallback = (
                result: OnLocalServiceFoundListenerResult
            ) => void;

    interface OnLocalServiceLostListenerResult {
        /**
         * 服务的名称
         */
        serviceName: string;
        /**
         * 服务的类型
         */
        serviceType: string;
    }

    /**
     * onLocalServiceLost 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffLocalServiceLostCallback = (
                result: OnLocalServiceLostListenerResult
            ) => void;
    /**
     * onLocalServiceResolveFail 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffLocalServiceResolveFailCallback = (
                result: OnLocalServiceLostListenerResult
            ) => void;

    interface OnNetworkWeakChangeListenerResult {
        /**
         * 当前网络类型
         */
        networkType: string;
        /**
         * 当前是否处于弱网状态
         */
        weakNet: boolean;
    }

    /**
     * onNetworkWeakChange 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffNetworkWeakChangeCallback = (
                result: OnNetworkWeakChangeListenerResult
            ) => void;

    interface OnPageNotFoundListenerResult {
        /**
         * 是否本次启动的首个页面（例如从分享等入口进来，首个页面是开发者配置的分享页面）
         */
        isEntryPage: boolean;
        /**
         * 不存在页面的路径 (代码包路径)
         */
        path: string;
        /**
         * 打开不存在页面的 query 参数
         */
        query: any;
    }

    /**
     * onPageNotFound 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffPageNotFoundCallback = (
                result: OnPageNotFoundListenerResult
            ) => void;

    interface OnScreenRecordingStateChangedListenerResult {
        /**
         * 录屏状态
         *
         * 可选值：
         * - 'start': 开始录屏;
         * - 'stop': 结束录屏;
         */
        state: "start" | "stop";
    }

    /**
     * onScreenRecordingStateChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffScreenRecordingStateChangedCallback = (
                result: OnScreenRecordingStateChangedListenerResult
            ) => void;

    interface OnVoIPChatInterruptedListenerResult {
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果（错误原因）
         */
        errMsg: string;
    }

    /**
     * onVoIPChatInterrupted 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffVoIPChatInterruptedCallback = (
                result: OnVoIPChatInterruptedListenerResult
            ) => void;

    interface OnVoIPChatMembersChangedListenerResult {
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果
         */
        errMsg: string;
        /**
         * 还在实时语音通话中的成员 openId 名单
         */
        openIdList: string[];
    }

    /**
     * onVoIPChatMembersChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffVoIPChatMembersChangedCallback = (
                result: OnVoIPChatMembersChangedListenerResult
            ) => void;

    interface OnVoIPChatSpeakersChangedListenerResult {
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果（错误原因）
         */
        errMsg: string;
        /**
         * 还在实时语音通话中的成员 openId 名单
         */
        openIdList: string[];
    }

    /**
     * onVoIPChatSpeakersChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffVoIPChatSpeakersChangedCallback = (
                result: OnVoIPChatSpeakersChangedListenerResult
            ) => void;

    interface OnVoIPChatStateChangedListenerResult {
        /**
         * 事件码
         */
        code: number;
        /**
         * 附加信息
         */
        data: any;
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果
         */
        errMsg: string;
    }

    /**
     * onVoIPChatStateChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffVoIPChatStateChangedCallback = (
                result: OnVoIPChatStateChangedListenerResult
            ) => void;

    interface OnVoIPVideoMembersChangedListenerResult {
        /**
         * 错误码
         */
        errCode: number;
        /**
         * 调用结果
         */
        errMsg: string;
        /**
         * 开启视频的成员名单
         */
        openIdList: string[];
    }

    /**
     * onVoIPVideoMembersChanged 传入的监听函数。不传此参数则移除所有监听函数。
     */
    type OffVoIPVideoMembersChangedCallback = (
                result: OnVoIPVideoMembersChangedListenerResult
            ) => void;
    /**
     * 小程序切后台事件的监听函数
     */
    type OnAppHideCallback = (res: GeneralCallbackResult) => void;

    /**
     * 打开的文件信息数组，只有从聊天素材场景打开（scene为1173）才会携带该参数
     */
    interface ForwardMaterials {
        /**
         * 文件名
         */
        name: string;
        /**
         * 文件路径（如果是webview则是url）
         */
        path: string;
        /**
         * 文件大小
         */
        size: number;
        /**
         * 文件的mimetype类型
         */
        type: string;
    }

    /**
     * 启动参数
     */
    interface LaunchOptionsApp {
        /**
         * 需要基础库： `2.20.0`
         *
         * API 类别
         *
         * 可选值：
         * - 'default': 默认类别;
         * - 'nativeFunctionalized': 原生功能化，视频号直播商品、商品橱窗等场景打开的小程序;
         * - 'browseOnly': 仅浏览，朋友圈快照页等场景打开的小程序;
         * - 'embedded': 内嵌，通过打开半屏小程序能力打开的小程序;
         */
        apiCategory: "default" | "nativeFunctionalized" | "browseOnly" | "embedded";
        /**
         * 打开的文件信息数组，只有从聊天素材场景打开（scene为1173）才会携带该参数
         */
        forwardMaterials: ForwardMaterials[];
        /**
         * 启动小程序的路径 (代码包路径)
         */
        path: string;
        /**
         * 启动小程序的 query 参数
         */
        query: any;
        /**
         * 来源信息。从另一个小程序、公众号或 App 进入小程序时返回。否则返回 `{}`。(参见后文注意)
         */
        referrerInfo: ReferrerInfo;
        /**
         * 启动小程序的[场景值](https://developers.weixin.qq.com/miniprogram/dev/component/xr-frame/core/scene.html)
         */
        scene: number;
        /**
         * 从微信群聊/单聊打开小程序时，chatType 表示具体微信群聊/单聊类型
         *
         * 可选值：
         * - 1: 微信联系人单聊;
         * - 2: 企业微信联系人单聊;
         * - 3: 普通微信群聊;
         * - 4: 企业微信互通群聊;
         */
        chatType?: 1 | 2 | 3 | 4;
        /**
         * shareTicket，详见[获取更多转发信息](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share.html)
         */
        shareTicket?: string;
    }

    /**
     * 小程序切前台事件的监听函数
     */
    type OnAppShowCallback = (
                /** 启动参数 */
                options: LaunchOptionsApp
            ) => void;
    /**
     * 音频因为受到系统占用而被中断开始事件的监听函数
     */
    type OnAudioInterruptionBeginCallback = (res: GeneralCallbackResult) => void;
    /**
     * 音频中断结束事件的监听函数
     */
    type OnAudioInterruptionEndCallback = (res: GeneralCallbackResult) => void;
    /**
     * 蓝牙低功耗的最大传输单元变化事件的监听函数
     */
    type OnBLEMTUChangeCallback = (result: OnBLEMTUChangeListenerResult) => void;
    /**
     * 当前外围设备被连接或断开连接事件的监听函数
     */
    type OnBLEPeripheralConnectionStateChangedCallback = (
                result: OnBLEPeripheralConnectionStateChangedListenerResult
            ) => void;
    /**
     * 音乐暂停事件的监听函数
     */
    type OnBackgroundAudioPauseCallback = (res: GeneralCallbackResult) => void;
    /**
     * 音乐播放事件的监听函数
     */
    type OnBackgroundAudioPlayCallback = (res: GeneralCallbackResult) => void;
    /**
     * 音乐停止事件的监听函数
     */
    type OnBackgroundAudioStopCallback = (res: GeneralCallbackResult) => void;

    interface OnBackgroundFetchDataListenerResult {
        /**
         * 缓存数据类别，取值为 periodic 或 pre
         */
        fetchType: string;
        /**
         * 缓存数据
         */
        fetchedData: string;
        /**
         * 小程序页面路径
         */
        path: string;
        /**
         * 传给页面的 query 参数
         */
        query: string;
        /**
         * 进入小程序的场景值
         */
        scene: number;
        /**
         * 客户端拿到缓存数据的时间戳
         */
        timeStamp: number;
    }

    /**
     * 收到 backgroundFetch 数据事件的监听函数
     */
    type OnBackgroundFetchDataCallback = (
                result: OnBackgroundFetchDataListenerResult
            ) => void;
    /**
     * 用户点击右上角菜单的「复制链接」按钮时触发的事件的监听函数
     */
    type OnCopyUrlCallback = (result: OnCopyUrlListenerResult) => void;

    interface OnDeviceMotionChangeListenerResult {
        /**
         * 当 手机坐标 X/Y 和 地球 X/Y 重合时，绕着 Z 轴转动的夹角为 alpha，范围值为 [0, 2*PI)。逆时针转动为正。
         */
        alpha: number;
        /**
         * 当手机坐标 Y/Z 和地球 Y/Z 重合时，绕着 X 轴转动的夹角为 beta。范围值为 [-1*PI, PI) 。顶部朝着地球表面转动为正。也有可能朝着用户为正。
         */
        beta: number;
        /**
         * 当手机 X/Z 和地球 X/Z 重合时，绕着 Y 轴转动的夹角为 gamma。范围值为 [-1*PI/2, PI/2)。右边朝着地球表面转动为正。
         */
        gamma: number;
    }

    /**
     * 设备方向变化事件的监听函数
     */
    type OnDeviceMotionChangeCallback = (
                result: OnDeviceMotionChangeListenerResult
            ) => void;
    /**
     * 接收 NFC 设备消息事件的监听函数
     */
    type OnHCEMessageCallback = (result: OnHCEMessageListenerResult) => void;
    /**
     * 小程序异步组件加载失败事件的监听函数
     */
    type OnLazyLoadErrorCallback = (
                result: OnLazyLoadErrorListenerResult
            ) => void;
    /**
     * mDNS 服务停止搜索的事件的监听函数
     */
    type OnLocalServiceDiscoveryStopCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * mDNS 服务发现的事件的监听函数
     */
    type OnLocalServiceFoundCallback = (
                result: OnLocalServiceFoundListenerResult
            ) => void;
    /**
     * mDNS 服务离开的事件的监听函数
     */
    type OnLocalServiceLostCallback = (
                result: OnLocalServiceLostListenerResult
            ) => void;
    /**
     * mDNS 服务解析失败的事件的监听函数
     */
    type OnLocalServiceResolveFailCallback = (
                result: OnLocalServiceLostListenerResult
            ) => void;
    /**
     * 弱网状态变化事件的监听函数
     */
    type OnNetworkWeakChangeCallback = (
                result: OnNetworkWeakChangeListenerResult
            ) => void;
    /**
     * 小程序要打开的页面不存在事件的监听函数
     */
    type OnPageNotFoundCallback = (result: OnPageNotFoundListenerResult) => void;
    /**
     * 用户录屏事件的监听函数
     */
    type OnScreenRecordingStateChangedCallback = (
                result: OnScreenRecordingStateChangedListenerResult
            ) => void;
    /**
     * 被动断开实时语音通话事件的监听函数
     */
    type OnVoIPChatInterruptedCallback = (
                result: OnVoIPChatInterruptedListenerResult
            ) => void;
    /**
     * 实时语音通话成员在线状态变化事件的监听函数
     */
    type OnVoIPChatMembersChangedCallback = (
                result: OnVoIPChatMembersChangedListenerResult
            ) => void;
    /**
     * 实时语音通话成员通话状态变化事件的监听函数
     */
    type OnVoIPChatSpeakersChangedCallback = (
                result: OnVoIPChatSpeakersChangedListenerResult
            ) => void;
    /**
     * 房间状态变化事件的监听函数
     */
    type OnVoIPChatStateChangedCallback = (
                result: OnVoIPChatStateChangedListenerResult
            ) => void;
    /**
     * 实时语音通话成员视频状态变化事件的监听函数
     */
    type OnVoIPVideoMembersChangedCallback = (
                result: OnVoIPVideoMembersChangedListenerResult
            ) => void;

    /**
     * 需要打开的卡券列表
     */
    interface OpenCardRequestInfo {
        /**
         * 卡券 ID
         */
        cardId: string;
        /**
         * 由 [uni.addCard](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/card/wx.addCard.html) 的返回对象中的加密 code 通过解密后得到，解密请参照：[code 解码接口](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=mp1499332673_Unm7V)
         */
        code: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenCardCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenCardFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenCardSuccessCallback = (res: GeneralCallbackResult) => void;

    interface OpenCardOption {
        /**
         * 需要打开的卡券列表
         */
        cardList: OpenCardRequestInfo[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenCardCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenCardFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenCardSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenChannelsActivityCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenChannelsActivityFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenChannelsActivitySuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OpenChannelsActivityOption {
        /**
         * 视频 feedId
         */
        feedId: string;
        /**
         * 视频号 id，以“sph”开头的id，可在视频号助手获取
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenChannelsActivityCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenChannelsActivityFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenChannelsActivitySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenChannelsEventCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenChannelsEventFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenChannelsEventSuccessCallback = (res: GeneralCallbackResult) => void;

    interface OpenChannelsEventOption {
        /**
         * 活动 id
         */
        eventId: string;
        /**
         * 视频号 id，以“sph”开头的id，可在视频号助手获取
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenChannelsEventCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenChannelsEventFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenChannelsEventSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenChannelsLiveCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenChannelsLiveFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenChannelsLiveSuccessCallback = (res: GeneralCallbackResult) => void;

    interface OpenChannelsLiveOption {
        /**
         * 视频号 id，以“sph”开头的id，可在视频号助手获取
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenChannelsLiveCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenChannelsLiveFailCallback;
        /**
         * 直播 feedId，通过 getChannelsLiveInfo 接口获取（基础库 v2.19.2 之前的版本需要填写）
         */
        feedId?: string;
        /**
         * 直播 nonceId，通过 getChannelsLiveInfo 接口获取（基础库 v2.19.2 之前的版本需要填写）
         */
        nonceId?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenChannelsLiveSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenChannelsUserProfileCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenChannelsUserProfileFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenChannelsUserProfileSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OpenChannelsUserProfileOption {
        /**
         * 视频号 id
         */
        finderUserName: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenChannelsUserProfileCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenChannelsUserProfileFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenChannelsUserProfileSuccessCallback;
    }

    /**
     * 客服信息
     */
    interface ExtInfoOption {
        /**
         * 客服链接
         */
        url: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenCustomerServiceChatCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenCustomerServiceChatFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenCustomerServiceChatSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OpenCustomerServiceChatOption {
        /**
         * 企业ID
         */
        corpId: string;
        /**
         * 客服信息
         */
        extInfo: ExtInfoOption;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenCustomerServiceChatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenCustomerServiceChatFailCallback;
        /**
         * 气泡消息图片
         */
        sendMessageImg?: string;
        /**
         * 气泡消息小程序路径
         */
        sendMessagePath?: string;
        /**
         * 气泡消息标题
         */
        sendMessageTitle?: string;
        /**
         * 是否发送小程序气泡消息
         */
        showMessageCard?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenCustomerServiceChatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenEmbeddedMiniProgramCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenEmbeddedMiniProgramFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenEmbeddedMiniProgramSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OpenEmbeddedMiniProgramOption {
        /**
         * 要打开的小程序 appId
         */
        appId: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenEmbeddedMiniProgramCompleteCallback;
        /**
         * 要打开的小程序版本。仅在当前小程序为开发版或体验版时此参数有效。如果当前小程序是正式版，则打开的小程序必定是正式版。
         *
         * 可选值：
         * - 'develop': 开发版;
         * - 'trial': 体验版;
         * - 'release': 正式版;
         */
        envVersion?: "develop" | "trial" | "release";
        /**
         * 需要传递给目标小程序的数据，目标小程序可在 `App.onLaunch`，`App.onShow` 中获取到这份数据。如果跳转的是小游戏，可以在 [uni.onShow](#)、[uni.getLaunchOptionsSync](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/life-cycle/wx.getLaunchOptionsSync.html) 中可以获取到这份数据数据。
         */
        extraData?: any;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenEmbeddedMiniProgramFailCallback;
        /**
         * 需要基础库： `2.24.0`
         *
         * 不reLaunch目标小程序，直接打开目标跳转的小程序退后台时的页面，需满足以下条件：1. 目标跳转的小程序生命周期未被销毁；2. 且目标当次启动的path、query、apiCategory与上次启动相同。
         */
        noRelaunchIfPathUnchanged?: boolean;
        /**
         * 打开的页面路径，如果为空则打开首页。path 中 ? 后面的部分会成为 query，在小程序的 `App.onLaunch`、`App.onShow` 和 `Page.onLoad` 的回调函数或小游戏的 [uni.onShow](#) 回调函数、[uni.getLaunchOptionsSync](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/life-cycle/wx.getLaunchOptionsSync.html) 中可以获取到 query 数据。对于小游戏，可以只传入 query 部分，来实现传参效果，如：传入 "?foo=bar"。
         */
        path?: string;
        /**
         * 小程序链接，当传递该参数后，可以不传 appId 和 path。链接可以通过【小程序菜单】->【复制链接】获取。仅 verify=binding 支持。
         */
        shortLink?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenEmbeddedMiniProgramSuccessCallback;
        /**
         * 需要基础库： `2.24.3`
         *
         * 校验方式。
         *
         * 可选值：
         * - 'binding': 校验小程序管理后台的绑定关系。;
         * - 'unionProduct': 校验目标打开链接是否为[小程序联盟](https://developers.weixin.qq.com/doc/ministore/union/brief-introduction.html)商品。;
         */
        verify?: "binding" | "unionProduct";
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type OpenSystemBluetoothSettingCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type OpenSystemBluetoothSettingFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type OpenSystemBluetoothSettingSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface OpenSystemBluetoothSettingOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: OpenSystemBluetoothSettingCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: OpenSystemBluetoothSettingFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: OpenSystemBluetoothSettingSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PauseBackgroundAudioCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PauseBackgroundAudioFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PauseBackgroundAudioSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface PauseBackgroundAudioOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PauseBackgroundAudioCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PauseBackgroundAudioFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PauseBackgroundAudioSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PauseVoiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PauseVoiceFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PauseVoiceSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PauseVoiceOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PauseVoiceCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PauseVoiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PauseVoiceSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PlayBackgroundAudioCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PlayBackgroundAudioFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PlayBackgroundAudioSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface PlayBackgroundAudioOption {
        /**
         * 音乐链接，目前支持的格式有 m4a, aac, mp3, wav
         */
        dataUrl: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PlayBackgroundAudioCompleteCallback;
        /**
         * 封面URL
         */
        coverImgUrl?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PlayBackgroundAudioFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PlayBackgroundAudioSuccessCallback;
        /**
         * 音乐标题
         */
        title?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PlayVoiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PlayVoiceFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PlayVoiceSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PlayVoiceOption {
        /**
         * 需要播放的语音文件的文件路径 (本地路径)
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PlayVoiceCompleteCallback;
        /**
         * 需要基础库： `1.6.0`
         *
         * 指定播放时长，到达指定的播放时长后会自动停止播放，单位：秒
         */
        duration?: number;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PlayVoiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PlayVoiceSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PluginLoginCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PluginLoginFailCallback = (res: GeneralCallbackResult) => void;

    interface PluginLoginSuccessCallbackResult {
        /**
         * 用于换取 openpid 的凭证（有效期五分钟）。插件开发者可以用此 code 在开发者服务器后台调用 [auth.getPluginOpenPId](#) 换取 openpid。
         */
        code: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type PluginLoginSuccessCallback = (
                result: PluginLoginSuccessCallbackResult
            ) => void;

    interface PluginLoginOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PluginLoginCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PluginLoginFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PluginLoginSuccessCallback;
    }

    interface Asset {
        src: string;
        /**
         * 可选值：
         * - 'font': 字体;
         * - 'image': 图片;
         */
        type: "image" | "font";
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PreloadAssetsCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PreloadAssetsFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PreloadAssetsSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PreloadAssetsOption {
        data: Asset[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PreloadAssetsCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PreloadAssetsFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PreloadAssetsSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PreloadSkylineViewCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PreloadSkylineViewFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PreloadSkylineViewSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface PreloadSkylineViewOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PreloadSkylineViewCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PreloadSkylineViewFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PreloadSkylineViewSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PreloadWebviewCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PreloadWebviewFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PreloadWebviewSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PreloadWebviewOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PreloadWebviewCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PreloadWebviewFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: PreloadWebviewSuccessCallback;
    }

    /**
     * 需要预览的资源列表
     */
    interface MediaSource {
        /**
         * 图片或视频的地址
         */
        url: string;
        /**
         * 视频的封面图片
         */
        poster?: string;
        /**
         * 资源的类型，默认为图片
         *
         * 可选值：
         * - 'image': 图片;
         * - 'video': 视频;
         */
        type?: "video" | "image";
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type PreviewMediaCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type PreviewMediaFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type PreviewMediaSuccessCallback = (res: GeneralCallbackResult) => void;

    interface PreviewMediaOption {
        /**
         * 需要预览的资源列表
         */
        sources: MediaSource[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: PreviewMediaCompleteCallback;
        /**
         * 当前显示的资源序号
         */
        current?: number;
        /**
         * 接口调用失败的回调函数
         */
        fail?: PreviewMediaFailCallback;
        /**
         * 需要基础库： `2.13.0`
         *
         * `origin`: 发送完整的referrer; `no-referrer`: 不发送。格式固定为 `https://servicewechat.com/{appid}/{version}/page-frame.html`，其中 {appid} 为小程序的 appid，{version} 为小程序的版本号，版本号为 0 表示为开发版、体验版以及审核版本，版本号为 devtools 表示为开发者工具，其余为正式版本；
         */
        referrerPolicy?: string;
        /**
         * 需要基础库： `2.13.0`
         *
         * 是否显示长按菜单。
         */
        showmenu?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: PreviewMediaSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestOrderPaymentCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type RequestOrderPaymentFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RequestOrderPaymentSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface RequestOrderPaymentOption {
        /**
         * 随机字符串，长度为32个字符以下
         */
        nonceStr: string;
        /**
         * 统一下单接口返回的 prepay_id 参数值，提交格式如：prepay_id=***
         */
        package: string;
        /**
         * 签名，具体见微信支付文档
         */
        paySign: string;
        /**
         * 时间戳，从 1970 年 1 月 1 日 00:00:00 至今的秒数，即当前的时间
         */
        timeStamp: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestOrderPaymentCompleteCallback;
        /**
         * 外部APP用户ID
         */
        extUserUin?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestOrderPaymentFailCallback;
        /**
         * 订单信息，仅在需要校验的场景下需要传递，具体见[接口说明](https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/order/requestOrderPayment)
         */
        orderInfo?: any;
        /**
         * 签名算法，应与后台下单时的值一致
         *
         * 可选值：
         * - 'MD5': 仅在 v2 版本接口适用;
         * - 'HMAC-SHA256': 仅在 v2 版本接口适用;
         * - 'RSA': 仅在 v3 版本接口适用;
         */
        signType?: "MD5" | "HMAC-SHA256" | "RSA";
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestOrderPaymentSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestPluginPaymentCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type RequestPluginPaymentFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type RequestPluginPaymentSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface RequestPluginPaymentOption {
        /**
         * 需要显示在页面中的金额，单位为分
         */
        fee: number;
        /**
         * 任意数据，传递给功能页中的响应函数
         */
        paymentArgs: any;
        /**
         * 插件版本
         *
         * 可选值：
         * - 'develop': 开发版;
         * - 'trial': 体验版;
         * - 'release': 正式版;
         */
        version: "develop" | "trial" | "release";
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestPluginPaymentCompleteCallback;
        /**
         * 需要显示在页面中的货币符号的代码
         */
        currencyType?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestPluginPaymentFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestPluginPaymentSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestSubscribeDeviceMessageCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface RequestSubscribeDeviceMessageFailCallbackResult {
        /**
         * 接口调用失败错误码，有可能为空
         */
        errCode: number;
        /**
         * 接口调用失败错误信息
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type RequestSubscribeDeviceMessageFailCallback = (
                result: RequestSubscribeDeviceMessageFailCallbackResult
            ) => void;

    interface RequestSubscribeDeviceMessageSuccessCallbackResult {
        /**
         * 接口调用成功时errMsg值为'requestSubscribeDeviceMessage:ok'
         */
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type RequestSubscribeDeviceMessageSuccessCallback = (
                result: RequestSubscribeDeviceMessageSuccessCallbackResult
            ) => void;

    interface RequestSubscribeDeviceMessageOption {
        /**
         * 设备型号 id 。通过微信公众平台注册设备获得。
         */
        modelId: string;
        /**
         * 设备唯一序列号。由厂商分配，长度不能超过128字节。字符只接受数字，大小写字母，下划线（_）和连字符（-）。
         */
        sn: string;
        /**
         * 设备票据，5分钟内有效。
         */
        snTicket: string;
        /**
         * 需要订阅的消息模板的 id 的集合，一次调用最多可订阅3条消息
         */
        tmplIds: any[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestSubscribeDeviceMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestSubscribeDeviceMessageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestSubscribeDeviceMessageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type RequestSubscribeMessageCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface RequestSubscribeMessageFailCallbackResult {
        /**
         * 接口调用失败错误码
         */
        errCode: number;
        /**
         * 接口调用失败错误信息
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type RequestSubscribeMessageFailCallback = (
                result: RequestSubscribeMessageFailCallbackResult
            ) => void;

    interface RequestSubscribeMessageSuccessCallbackResult {
        /**
         * 接口调用成功时errMsg值为'requestSubscribeMessage:ok'
         */
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type RequestSubscribeMessageSuccessCallback = (
                result: RequestSubscribeMessageSuccessCallbackResult
            ) => void;

    interface RequestSubscribeMessageOption {
        /**
         * 需要订阅的消息模板的id的集合，一次调用最多可订阅3条消息（注意：iOS客户端7.0.6版本、Android客户端7.0.7版本之后的一次性订阅/长期订阅才支持多个模板消息，iOS客户端7.0.5版本、Android客户端7.0.6版本之前的一次订阅只支持一个模板消息）消息模板id在[微信公众平台(mp.weixin.qq.com)-功能-订阅消息]中配置。每个tmplId对应的模板标题需要不相同，否则会被过滤。
         */
        tmplIds: any[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: RequestSubscribeMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: RequestSubscribeMessageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: RequestSubscribeMessageSuccessCallback;
    }

    interface ReserveChannelsLiveOption {
        /**
         * 预告 id，通过 getChannelsLiveNoticeInfo 接口获取
         */
        noticeId: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SaveFileToDiskCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SaveFileToDiskFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SaveFileToDiskSuccessCallback = (res: GeneralCallbackResult) => void;

    interface SaveFileToDiskOption {
        /**
         * 待保存文件路径
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SaveFileToDiskCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SaveFileToDiskFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SaveFileToDiskSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SeekBackgroundAudioCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SeekBackgroundAudioFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SeekBackgroundAudioSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface SeekBackgroundAudioOption {
        /**
         * 音乐位置，单位：秒
         */
        position: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SeekBackgroundAudioCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SeekBackgroundAudioFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SeekBackgroundAudioSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SendHCEMessageCompleteCallback = (res: NFCError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SendHCEMessageFailCallback = (res: NFCError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SendHCEMessageSuccessCallback = (res: NFCError) => void;

    interface SendHCEMessageOption {
        /**
         * 二进制数据
         */
        data: ArrayBuffer;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SendHCEMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SendHCEMessageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SendHCEMessageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SendSmsCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SendSmsFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SendSmsSuccessCallback = (res: GeneralCallbackResult) => void;

    interface SendSmsOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SendSmsCompleteCallback;
        /**
         * 预填到发送短信面板的内容
         */
        content?: string;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SendSmsFailCallback;
        /**
         * 预填到发送短信面板的手机号
         */
        phoneNumber?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: SendSmsSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetBackgroundFetchTokenCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetBackgroundFetchTokenFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetBackgroundFetchTokenSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface SetBackgroundFetchTokenOption {
        /**
         * 自定义的登录态
         */
        token: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetBackgroundFetchTokenCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetBackgroundFetchTokenFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetBackgroundFetchTokenSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetEnable1v1ChatCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetEnable1v1ChatFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetEnable1v1ChatSuccessCallback = (res: GeneralCallbackResult) => void;

    interface SetEnable1v1ChatOption {
        /**
         * 是否开启
         */
        enable: boolean;
        /**
         * 窗口背景色(音频通话背景以及小窗模式背景)
         *
         * 可选值：
         * - 0: #262930;
         * - 1: #FA5151;
         * - 2: #FA9D3B;
         * - 3: #3D7257;
         * - 4: #1485EE;
         * - 5: #6467F0;
         */
        backgroundType?: 0 | 1 | 2 | 3 | 4 | 5;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetEnable1v1ChatCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetEnable1v1ChatFailCallback;
        /**
         * 小窗样式
         */
        minWindowType?: number;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetEnable1v1ChatSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetInnerAudioOptionCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetInnerAudioOptionFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetInnerAudioOptionSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface SetInnerAudioOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetInnerAudioOptionCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetInnerAudioOptionFailCallback;
        /**
         * 是否与其他音频混播，设置为 true 之后，不会终止其他应用或微信内的音乐
         */
        mixWithOther?: boolean;
        /**
         * （仅在 iOS 生效）是否遵循静音开关，设置为 false 之后，即使是在静音模式下，也能播放声音
         */
        obeyMuteSwitch?: boolean;
        /**
         * true 代表用扬声器播放，false 代表听筒播放，默认值为 true。
         */
        speakerOn?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetInnerAudioOptionSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetTopBarTextCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetTopBarTextFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetTopBarTextSuccessCallback = (res: GeneralCallbackResult) => void;

    interface SetTopBarTextOption {
        /**
         * 置顶栏文字
         */
        text: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetTopBarTextCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetTopBarTextFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetTopBarTextSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetVisualEffectOnCaptureCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetVisualEffectOnCaptureFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetVisualEffectOnCaptureSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface SetVisualEffectOnCaptureOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetVisualEffectOnCaptureCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetVisualEffectOnCaptureFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetVisualEffectOnCaptureSuccessCallback;
        /**
         * 截屏/录屏时的表现，仅支持 none / hidden，传入 hidden 则表示在截屏/录屏时隐藏屏幕
         */
        visualEffect?: string;
    }

    /**
     * 提供预设的 Wi-Fi 信息列表
     */
    interface WifiData {
        /**
         * Wi-Fi 的 BSSID
         */
        BSSID?: string;
        /**
         * Wi-Fi 的 SSID
         */
        SSID?: string;
        /**
         * Wi-Fi 设备密码
         */
        password?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetWifiListCompleteCallback = (res: WifiError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetWifiListFailCallback = (res: WifiError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetWifiListSuccessCallback = (res: WifiError) => void;

    interface SetWifiListOption {
        /**
         * 提供预设的 Wi-Fi 信息列表
         */
        wifiList: WifiData[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetWifiListCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetWifiListFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetWifiListSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SetWindowSizeCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SetWindowSizeFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SetWindowSizeSuccessCallback = (res: GeneralCallbackResult) => void;

    interface SetWindowSizeOption {
        /**
         * 窗口高度，以像素为单位
         */
        height: number;
        /**
         * 窗口宽度，以像素为单位
         */
        width: number;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SetWindowSizeCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SetWindowSizeFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SetWindowSizeSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ShareFileMessageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ShareFileMessageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ShareFileMessageSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ShareFileMessageOption {
        /**
         * 要分享的文件地址，必须为本地路径或临时路径
         */
        filePath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ShareFileMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ShareFileMessageFailCallback;
        /**
         * 自定义文件名，若留空则使用filePath中的文件名
         */
        fileName?: string;
        /**
         * 接口调用成功的回调函数
         */
        success?: ShareFileMessageSuccessCallback;
    }

    /**
     * 运动数据列表
     */
    interface WxaSportRecord {
        /**
         * 消耗卡路里
         */
        calorie: number;
        /**
         * 运动距离
         */
        distance: number;
        /**
         * 运动时长
         */
        time: number;
        /**
         * 运动项目id
         */
        typeId: number;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ShareToWeRunCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ShareToWeRunFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ShareToWeRunSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ShareToWeRunOption {
        /**
         * 运动数据列表
         */
        recordList: WxaSportRecord[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ShareToWeRunCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ShareToWeRunFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ShareToWeRunSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ShareVideoMessageCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ShareVideoMessageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ShareVideoMessageSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ShareVideoMessageOption {
        /**
         * 要分享的视频地址，必须为本地路径或临时路径
         */
        videoPath: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ShareVideoMessageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ShareVideoMessageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ShareVideoMessageSuccessCallback;
        /**
         * 缩略图路径，若留空则使用视频首帧
         */
        thumbPath?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ShowRedPackageCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ShowRedPackageFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ShowRedPackageSuccessCallback = (res: GeneralCallbackResult) => void;

    interface ShowRedPackageOption {
        /**
         * 封面地址
         */
        url: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ShowRedPackageCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ShowRedPackageFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ShowRedPackageSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type ShowShareImageMenuCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type ShowShareImageMenuFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type ShowShareImageMenuSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface ShowShareImageMenuOption {
        /**
         * 要分享的图片地址，必须为本地路径或临时路径
         */
        path: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: ShowShareImageMenuCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: ShowShareImageMenuFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: ShowShareImageMenuSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartDeviceMotionListeningCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StartDeviceMotionListeningFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StartDeviceMotionListeningSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StartDeviceMotionListeningOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartDeviceMotionListeningCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartDeviceMotionListeningFailCallback;
        /**
         * 监听设备方向的变化回调函数的执行频率
         *
         * 可选值：
         * - 'game': 适用于更新游戏的回调频率，在 20ms/次 左右;
         * - 'ui': 适用于更新 UI 的回调频率，在 60ms/次 左右;
         * - 'normal': 普通的回调频率，在 200ms/次 左右;
         */
        interval?: "game" | "ui" | "normal";
        /**
         * 接口调用成功的回调函数
         */
        success?: StartDeviceMotionListeningSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartHCECompleteCallback = (res: NFCError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StartHCEFailCallback = (res: NFCError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StartHCESuccessCallback = (res: NFCError) => void;

    interface StartHCEOption {
        /**
         * 需要注册到系统的 AID 列表
         */
        aid_list: string[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartHCECompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartHCEFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StartHCESuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartLocalServiceDiscoveryCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StartLocalServiceDiscoveryFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'invalid param': serviceType 为空;
         * - 'scan task already exist': 在当前 startLocalServiceDiscovery 发起的搜索未停止的情况下，再次调用 startLocalServiceDiscovery;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type StartLocalServiceDiscoveryFailCallback = (
                result: StartLocalServiceDiscoveryFailCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StartLocalServiceDiscoverySuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StartLocalServiceDiscoveryOption {
        /**
         * 要搜索的服务类型
         */
        serviceType: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartLocalServiceDiscoveryCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartLocalServiceDiscoveryFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StartLocalServiceDiscoverySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartLocationUpdateBackgroundCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StartLocationUpdateBackgroundFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StartLocationUpdateBackgroundSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StartLocationUpdateBackgroundOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartLocationUpdateBackgroundCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartLocationUpdateBackgroundFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StartLocationUpdateBackgroundSuccessCallback;
        /**
         * wgs84 返回 gps 坐标，gcj02 返回可用于 wx.openLocation 的坐标
         */
        type?: string;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StartRecordCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StartRecordFailCallback = (res: GeneralCallbackResult) => void;

    interface StartRecordSuccessCallbackResult {
        /**
         * 录音文件的临时路径 (本地路径)
         */
        tempFilePath: string;
        errMsg: string;
    }

    /**
     * 接口调用成功的回调函数
     */
    type WxStartRecordSuccessCallback = (
                result: StartRecordSuccessCallbackResult
            ) => void;

    interface WxStartRecordOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StartRecordCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StartRecordFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: WxStartRecordSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopBackgroundAudioCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopBackgroundAudioFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopBackgroundAudioSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StopBackgroundAudioOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopBackgroundAudioCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopBackgroundAudioFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopBackgroundAudioSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopDeviceMotionListeningCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopDeviceMotionListeningFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopDeviceMotionListeningSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StopDeviceMotionListeningOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopDeviceMotionListeningCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopDeviceMotionListeningFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopDeviceMotionListeningSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopFaceDetectCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopFaceDetectFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopFaceDetectSuccessCallback = (res: GeneralCallbackResult) => void;

    interface StopFaceDetectOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopFaceDetectCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopFaceDetectFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopFaceDetectSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopHCECompleteCallback = (res: NFCError) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopHCEFailCallback = (res: NFCError) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopHCESuccessCallback = (res: NFCError) => void;

    interface StopHCEOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopHCECompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopHCEFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopHCESuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopLocalServiceDiscoveryCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StopLocalServiceDiscoveryFailCallbackResult {
        /**
         * 错误信息
         *
         * 可选值：
         * - 'task not found': 在当前没有处在搜索服务中的情况下调用 stopLocalServiceDiscovery;
         */
        errMsg: string;
    }

    /**
     * 接口调用失败的回调函数
     */
    type StopLocalServiceDiscoveryFailCallback = (
                result: StopLocalServiceDiscoveryFailCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopLocalServiceDiscoverySuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface StopLocalServiceDiscoveryOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopLocalServiceDiscoveryCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopLocalServiceDiscoveryFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopLocalServiceDiscoverySuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopRecordCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopRecordFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type WxStopRecordSuccessCallback = (res: GeneralCallbackResult) => void;

    interface WxStopRecordOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopRecordCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopRecordFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: WxStopRecordSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type StopVoiceCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type StopVoiceFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type StopVoiceSuccessCallback = (res: GeneralCallbackResult) => void;

    interface StopVoiceOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: StopVoiceCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: StopVoiceFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: StopVoiceSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type SubscribeVoIPVideoMembersCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type SubscribeVoIPVideoMembersFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type SubscribeVoIPVideoMembersSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface SubscribeVoIPVideoMembersOption {
        /**
         * 订阅的成员列表
         */
        openIdList: string[];
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: SubscribeVoIPVideoMembersCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: SubscribeVoIPVideoMembersFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: SubscribeVoIPVideoMembersSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type UpdateShareMenuCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type UpdateShareMenuFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type UpdateShareMenuSuccessCallback = (res: GeneralCallbackResult) => void;

    /**
     * 参数列表
     */
    interface UpdatableMessageFrontEndParameter {
        /**
         * 参数名
         */
        name: string;
        /**
         * 参数值
         */
        value: string;
    }

    /**
     * 需要基础库： `2.4.0`
     *
     * 动态消息的模板信息
     */
    interface UpdatableMessageFrontEndTemplateInfo {
        /**
         * 参数列表
         */
        parameterList: UpdatableMessageFrontEndParameter[];
    }

    interface UpdateShareMenuOption {
        /**
         * 需要基础库： `2.4.0`
         *
         * 动态消息的 activityId。通过 [updatableMessage.createActivityId](#) 接口获取
         */
        activityId?: string;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: UpdateShareMenuCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: UpdateShareMenuFailCallback;
        /**
         * 需要基础库： `2.13.0`
         *
         * 是否是私密消息。详见 [小程序私密消息](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share/private-message.html)
         */
        isPrivateMessage?: boolean;
        /**
         * 需要基础库： `2.4.0`
         *
         * 是否是动态消息，详见[动态消息](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share/updatable-message.html)
         */
        isUpdatableMessage?: boolean;
        /**
         * 接口调用成功的回调函数
         */
        success?: UpdateShareMenuSuccessCallback;
        /**
         * 需要基础库： `2.4.0`
         *
         * 动态消息的模板信息
         */
        templateInfo?: UpdatableMessageFrontEndTemplateInfo;
        /**
         * 需要基础库： `2.11.0`
         *
         * 群待办消息的id，通过toDoActivityId可以把多个群待办消息聚合为同一个。通过 [updatableMessage.createActivityId](#) 接口获取。详见[群待办消息](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share.html)
         */
        toDoActivityId?: string;
        /**
         * 是否使用带 shareTicket 的转发[详情](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share.html)
         */
        withShareTicket?: boolean;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type UpdateVoIPChatMuteConfigCompleteCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用失败的回调函数
     */
    type UpdateVoIPChatMuteConfigFailCallback = (
                res: GeneralCallbackResult
            ) => void;
    /**
     * 接口调用成功的回调函数
     */
    type UpdateVoIPChatMuteConfigSuccessCallback = (
                res: GeneralCallbackResult
            ) => void;

    interface UpdateVoIPChatMuteConfigOption {
        /**
         * 静音设置
         */
        muteConfig: MuteConfig;
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: UpdateVoIPChatMuteConfigCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: UpdateVoIPChatMuteConfigFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: UpdateVoIPChatMuteConfigSuccessCallback;
    }

    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    type UpdateWeChatAppCompleteCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    type UpdateWeChatAppFailCallback = (res: GeneralCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    type UpdateWeChatAppSuccessCallback = (res: GeneralCallbackResult) => void;

    interface UpdateWeChatAppOption {
        /**
         * 接口调用结束的回调函数（调用成功、失败都会执行）
         */
        complete?: UpdateWeChatAppCompleteCallback;
        /**
         * 接口调用失败的回调函数
         */
        fail?: UpdateWeChatAppFailCallback;
        /**
         * 接口调用成功的回调函数
         */
        success?: UpdateWeChatAppSuccessCallback;
    }
}

interface Uni {
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：不支持
     *
     * 从本地缓存中同步批量获取指定 key 的内容。
     *
     *
     * ****
     *
     * 对于多个key的读取, 批量读取在性能上优于多次getStorageSync读取
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchGetStorageSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchGetStorageSync.html)
     */
    batchGetStorageSync(keyList: string[]): any[];
    /**
     *
     * 需要基础库： `2.2.2`
     *
     * 在插件中使用：需要基础库 `2.2.2`
     *
     * 获取当前帐号信息。线上小程序版本号仅支持在正式版小程序中获取，开发版和体验版中无法获取。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/account-info/wx.getAccountInfoSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/account-info/wx.getAccountInfoSync.html)
     */
    getAccountInfoSync(): UniNamespace.AccountInfo;
    /**
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * [uni.getBatteryInfo](https://developers.weixin.qq.com/miniprogram/dev/api/device/battery/wx.getBatteryInfo.html) 的同步版本
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/battery/wx.getBatteryInfoSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/battery/wx.getBatteryInfoSync.html)
     */
    getBatteryInfoSync(): UniNamespace.GetBatteryInfoSyncResult;
    /**
     *
     * 需要基础库： `2.17.0`
     *
     * 在插件中使用：不支持
     *
     * 给定实验参数数组，获取对应的实验参数值
     *
     * **提示**
     *
     * 假设实验参数有 `color`, `size`
     * 调用 wx.getExptInfoSync() 会返回 `{color:'#fff',size:20}` 类似的结果
     * 而 wx.getExptInfoSync(['color']) 则只会返回 `{color:'#fff'}`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.getExptInfoSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.getExptInfoSync.html)
     */
    getExptInfoSync(keys?: string[]): any;
    /**
     *
     * 需要基础库： `2.26.2`
     *
     * 在插件中使用：需要基础库 `2.26.2`
     *
     * 获取当前运行环境对于 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的支持情况
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSkylineInfoSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSkylineInfoSync.html)
     */
    getSkylineInfoSync(): UniNamespace.SkylineInfo;
    /**
     *
     * 需要基础库： `2.26.3`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 获取 Webview 小程序的 UserAgent
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getRendererUserAgent.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getRendererUserAgent.html)
     */
    getRendererUserAgent(): Promise<string>;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createAudioContext.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createAudioContext.html)
     */
    createAudioContext(id: string, component?: any): UniNamespace.AudioContext;
    /**
     *
     * 需要基础库： `2.24.0`
     *
     * 在插件中使用：不支持
     *
     * 创建缓存管理器
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/cachemanager/wx.createCacheManager.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/cachemanager/wx.createCacheManager.html)
     */
    createCacheManager(option: UniNamespace.CreateCacheManagerOption): UniNamespace.CacheManager;
    /**
     *
     * 需要基础库： `1.9.9`
     *
     * 在插件中使用：需要基础库 `2.19.2`
     *
     * 获取全局唯一的文件管理器
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/wx.getFileSystemManager.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/wx.getFileSystemManager.html)
     */
    getFileSystemManager(): UniNamespace.FileSystemManager;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 创建 [live-player](https://developers.weixin.qq.com/miniprogram/dev/component/live-player.html) 上下文 [LivePlayerContext](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/LivePlayerContext.html) 对象。建议使用 [uni.createSelectorQuery](https://developers.weixin.qq.com/miniprogram/dev/api/wxml/wx.createSelectorQuery.html) 获取 context 对象。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/live/wx.createLivePlayerContext.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/live/wx.createLivePlayerContext.html)
     */
    createLivePlayerContext(id: string, component?: any): UniNamespace.LivePlayerContext;
    /**
     *
     * 需要基础库： `2.1.0`
     *
     * 在插件中使用：不支持
     *
     * 获取日志管理器对象。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.getLogManager.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.getLogManager.html)
     */
    getLogManager(option: UniNamespace.GetLogManagerOption): UniNamespace.LogManager;
    /**
     *
     * 需要基础库： `2.13.0`
     *
     * 在插件中使用：支持
     *
     * 创建媒体音频播放器对象 [MediaAudioPlayer](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/MediaAudioPlayer.html) 对象，可用于播放视频解码器 [VideoDecoder](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/VideoDecoder.html) 输出的音频。
     *
     *
     * **完整demo（小游戏）**
     *
     * - https://developers.weixin.qq.com/s/SF2duHmb7MjI
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createMediaAudioPlayer.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createMediaAudioPlayer.html)
     */
    createMediaAudioPlayer(): UniNamespace.MediaAudioPlayer;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.10.0`
     *
     * 创建音视频处理容器，最终可将容器中的轨道合成一个视频
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/wx.createMediaContainer.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-processing/wx.createMediaContainer.html)
     */
    createMediaContainer(): UniNamespace.MediaContainer;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：需要基础库 `2.11.0`
     *
     * 创建 WebGL 画面录制器，可逐帧录制在 WebGL 上渲染的画面并导出视频文件
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/wx.createMediaRecorder.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/wx.createMediaRecorder.html)
     */
    createMediaRecorder(canvas: any, options: UniNamespace.CreateMediaRecorderOption): UniNamespace.MediaRecorder;
    /**
     *
     * 需要基础库： `2.11.2`
     *
     * 在插件中使用：需要基础库 `2.11.2`
     *
     * 获取 NFC 实例
     *
     * **示例代码**
     *
     * [在微信开发者工具中查看示例](https://developers.weixin.qq.com/s/1WsbDwmb75ig)
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc/wx.getNFCAdapter.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc/wx.getNFCAdapter.html)
     */
    getNFCAdapter(): UniNamespace.NFCAdapter;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：需要基础库 `2.16.1`
     *
     * 创建离屏 canvas 实例
     *
     * **离屏 Canvas 类型不可混用**
     *
     * 由于 webgl canvas 和 2d canvas 的底层实现方式不同，因此必须要在调用 `wx.createOffscreenCanvas` 时提前指定类型。
     *
     * 指定类型后，离屏 canvas `getContext(type)` 调用不允许混用，如不能对 webgl canvas 调用 `getContext('2d')`。
     *
     * 同样的，不同类型 canvas 调用 `createImage` 创建的图片对象也不支持混用，使用时请注意尽量使用 canvas 自身的 `createImage` 创建图片对象。
     *
     * **与 MediaRecorder 结合**
     *
     * 离屏 webgl canvas 支持作为参数传递给 [`wx.createMediaRecorder`](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/wx.createMediaRecorder.html), 离屏 2d canvas 暂不支持。
     *
     * **旧版 createOffscreenCanvas**
     *
     * 旧版函数签名为 `wx.createOffscreenCanvas(width: number, height: number, this: object): OffscreenCanvas`，从基础库 2.7.0 开始支持
     *
     * 从基础库 2.16.1 开始改为 `wx.createOffscreenCanvas(options: object): OffscreenCanvas`，向下兼容旧版入参。
     * 但需注意旧版入参只能创建 webgl 类型，如需创建 2d 类型则必须使用新版。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/canvas/wx.createOffscreenCanvas.html](https://developers.weixin.qq.com/miniprogram/dev/api/canvas/wx.createOffscreenCanvas.html)
     */
    createOffscreenCanvas(width: number, height: number, component?: any): UniNamespace.OffscreenCanvas;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：需要基础库 `2.16.1`
     *
     * 创建离屏 canvas 实例
     *
     * **离屏 Canvas 类型不可混用**
     *
     * 由于 webgl canvas 和 2d canvas 的底层实现方式不同，因此必须要在调用 `wx.createOffscreenCanvas` 时提前指定类型。
     *
     * 指定类型后，离屏 canvas `getContext(type)` 调用不允许混用，如不能对 webgl canvas 调用 `getContext('2d')`。
     *
     * 同样的，不同类型 canvas 调用 `createImage` 创建的图片对象也不支持混用，使用时请注意尽量使用 canvas 自身的 `createImage` 创建图片对象。
     *
     * **与 MediaRecorder 结合**
     *
     * 离屏 webgl canvas 支持作为参数传递给 [`wx.createMediaRecorder`](https://developers.weixin.qq.com/miniprogram/dev/api/media/media-recorder/wx.createMediaRecorder.html), 离屏 2d canvas 暂不支持。
     *
     * **旧版 createOffscreenCanvas**
     *
     * 旧版函数签名为 `wx.createOffscreenCanvas(width: number, height: number, this: object): OffscreenCanvas`，从基础库 2.7.0 开始支持
     *
     * 从基础库 2.16.1 开始改为 `wx.createOffscreenCanvas(options: object): OffscreenCanvas`，向下兼容旧版入参。
     * 但需注意旧版入参只能创建 webgl 类型，如需创建 2d 类型则必须使用新版。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/canvas/wx.createOffscreenCanvas.html](https://developers.weixin.qq.com/miniprogram/dev/api/canvas/wx.createOffscreenCanvas.html)
     */
    createOffscreenCanvas(option: UniNamespace.CreateOffscreenCanvasOption): UniNamespace.OffscreenCanvas;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：支持
     *
     * 获取当前小程序性能相关的信息。关于小程序启动性能优化的更多内容，请参考[启动性能指南](https://developers.weixin.qq.com/miniprogram/dev/framework/performance/tips/start.html)。
     *
     * ****
     *
     * 目前支持获取以下几类性能指标，具体内容请参考 [PerformanceEntry](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/PerformanceEntry.html)：
     *
     * | 指标类型（entryType） | 指标名称          | 最低版本 ｜
     * | ------------------- | ---------------- | ------ |
     * | 路由（navigation）   | route: 路由性能 | |
     * | 路由（navigation）   | appLaunch: 小程序启动耗时 | |
     * | 渲染（render）       | firstRender: 页面首次渲染耗时 | |
     * | 渲染（render）       | firstPaint: 页面首次绘制 | <2.21.2> |
     * | 渲染（render）       | firstContentfulPaint: 页面首次内容绘制 |  <2.21.2> |
     * | 渲染（render）       | largestContentfulPaint: 页面最大内容绘制 | <2.23.1> |
     * | 脚本（script）       | evaluateScript: 注入脚本耗时  | |
     * | 包加载（loadPackage）| downloadPackage: 代码包下载耗时  | <2.24.0> |
     * | 资源（resource）     | resourceTiming: 视图层资源加载耗时  | <2.24.0> |
     *
     *
     * **注意**
     *
     * - 目前，当开启代码 [按需注入](https://developers.weixin.qq.com/miniprogram/dev/framework/ability/lazyload.html) 时，`evaluateScript` 将仅包含公有部分代码（2.21.2 开始会区分公共部分/页面和组件的部分），页面和组件的代码注入的时间会包含在 `firstRender` 中（因为页面和组件的代码注入过程成为了首次渲染过程的一部分）。因此开启按需注入后，脚本耗时降低，渲染时间提高属于正常现象，优化效果可以关注整体启动耗时（`appLaunch`）来评估。
     * - firstPaint 和 firstContentfulPaint 指标在开启 vConsole 的情况下，由于绘制 vConsole 面板，会导致数据提前。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.getPerformance.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.getPerformance.html)
     */
    getPerformance(): UniNamespace.Performance;
    /**
     *
     * 需要基础库： `2.27.3`
     *
     * 在插件中使用：不支持
     *
     * 触发分包预下载。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/subpackage/wx.preDownloadSubpackage.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/subpackage/wx.preDownloadSubpackage.html)
     */
    preDownloadSubpackage(option: UniNamespace.PreDownloadSubpackageOption): UniNamespace.PreDownloadSubpackageTask;
    /**
     *
     * 需要基础库： `2.7.1`
     *
     * 在插件中使用：需要基础库 `2.16.0`
     *
     * 获取实时日志管理器对象。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.getRealtimeLogManager.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/debug/wx.getRealtimeLogManager.html)
     */
    getRealtimeLogManager(): UniNamespace.RealtimeLogManager;
    /**
     *
     * 需要基础库： `2.18.0`
     *
     * 在插件中使用：支持
     *
     * 创建一个 TCP Socket 实例。使用前请注意阅读[相关说明](https://developers.weixin.qq.com/miniprogram/dev/framework/ability/network.html)。
     *
     * **连接限制**
     *
     * - 允许与局域网内的非本机 IP 通信
     * - 允许与配置过的服务器域名通信，详见[相关说明](https://developers.weixin.qq.com/miniprogram/dev/framework/ability/network.html)
     * - 禁止与以下端口号连接：`1024 以下` `1099` `1433` `1521` `1719` `1720` `1723` `2049` `2375` `3128` `3306` `3389` `3659` `4045` `5060` `5061` `5432` `5984` `6379` `6000` `6566` `7001` `7002` `8000-8100` `8443` `8888` `9200` `9300` `10051` `10080` `11211` `27017` `27018` `27019`
     * - 每 5 分钟内最多创建 20 个 TCPSocket
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/wx.createTCPSocket.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/tcp/wx.createTCPSocket.html)
     */
    createTCPSocket(): UniNamespace.TCPSocket;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.11.1`
     *
     * 创建一个 UDP Socket 实例。使用前请注意阅读[相关说明](https://developers.weixin.qq.com/miniprogram/dev/framework/ability/network.html)。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/wx.createUDPSocket.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/udp/wx.createUDPSocket.html)
     */
    createUDPSocket(): UniNamespace.UDPSocket;
    /**
     *
     * 需要基础库： `2.17.3`
     *
     * 在插件中使用：不支持
     *
     * 获取用户加密模块
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/wx.getUserCryptoManager.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/crypto/wx.getUserCryptoManager.html)
     */
    getUserCryptoManager(): UniNamespace.UserCryptoManager;
    /**
     *
     * 需要基础库： `2.20.0`
     *
     * 在插件中使用：需要基础库 `2.20.0`
     *
     * 创建 vision kit 会话对象。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ai/visionkit/wx.createVKSession.html](https://developers.weixin.qq.com/miniprogram/dev/api/ai/visionkit/wx.createVKSession.html)
     */
    createVKSession(option: UniNamespace.VKConfig): UniNamespace.VKSession;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：需要基础库 `2.11.0`
     *
     * 创建视频解码器，可逐帧获取解码后的数据
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/wx.createVideoDecoder.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video-decoder/wx.createVideoDecoder.html)
     */
    createVideoDecoder(): UniNamespace.VideoDecoder;
    /**
     *
     * 需要基础库： `2.19.0`
     *
     * 在插件中使用：不支持
     *
     * 创建 WebAudio 上下文。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createWebAudioContext.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.createWebAudioContext.html)
     */
    createWebAudioContext(): UniNamespace.WebAudioContext;
    /**
     *
     * 需要基础库： `1.9.90`
     *
     * 在插件中使用：需要基础库 `2.18.1`
     *
     * 创建一个 Worker 线程
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/worker/wx.createWorker.html](https://developers.weixin.qq.com/miniprogram/dev/api/worker/wx.createWorker.html)
     */
    createWorker(scriptPath: string, options?: UniNamespace.CreateWorkerOption): UniNamespace.Worker;
    /**
     *
     * 需要基础库： `2.22.0`
     *
     * 在插件中使用：需要基础库 `2.22.0`
     *
     * 判断支持版本
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ai/visionkit/wx.isVKSupport.html](https://developers.weixin.qq.com/miniprogram/dev/api/ai/visionkit/wx.isVKSupport.html)
     */
    isVKSupport(version: "v1" | "v2"): boolean;
    /**
     *
     * 需要基础库： `2.14.0`
     *
     * 在插件中使用：不支持
     *
     * 根据传入的 buffer 创建一个唯一的 URL 存在内存中
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.createBufferURL.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.createBufferURL.html)
     */
    createBufferURL(buffer: ArrayBuffer | Int8Array | Uint8Array | Uint8ClampedArray | Int16Array | Uint16Array | Int32Array | Uint32Array | Float32Array | Float64Array): string;
    /**
     *
     * 需要基础库： `1.1.0`
     *
     * 在插件中使用：不支持
     *
     * 批量添加卡券。只有通过 [认证](https://developers.weixin.qq.com/miniprogram/product/renzheng.html) 的小程序或文化互动类目的小游戏才能使用。更多文档请参考 [微信卡券接口文档](https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=1490190158&version=1&lang=zh_CN&platform=2)。
     *
     * **cardExt 说明**
     *
     * cardExt 是卡券的扩展参数，其值是一个 JSON 字符串。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/card/wx.addCard.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/card/wx.addCard.html)
     */
    addCard(option: UniNamespace.AddCardOption): void;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：不支持
     *
     * 收藏文件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/favorites/wx.addFileToFavorites.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/favorites/wx.addFileToFavorites.html)
     */
    addFileToFavorites(option: UniNamespace.AddFileToFavoritesOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：不支持
     *
     * 向系统日历添加事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/calendar/wx.addPhoneCalendar.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/calendar/wx.addPhoneCalendar.html)
     */
    addPhoneCalendar(option: UniNamespace.AddPhoneCalendarOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：不支持
     *
     * 向系统日历添加重复事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/calendar/wx.addPhoneRepeatCalendar.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/calendar/wx.addPhoneRepeatCalendar.html)
     */
    addPhoneRepeatCalendar(option: UniNamespace.AddPhoneRepeatCalendarOption): void;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：不支持
     *
     * 收藏视频
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/favorites/wx.addVideoToFavorites.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/favorites/wx.addVideoToFavorites.html)
     */
    addVideoToFavorites(option: UniNamespace.AddVideoToFavoritesOption): void;
    /**
     *
     * 需要基础库： `2.13.0`
     *
     * 在插件中使用：不支持
     *
     * 验证私密消息。用法详情见 [小程序私密消息使用指南](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/share/private-message.html)
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.authPrivateMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.authPrivateMessage.html)
     */
    authPrivateMessage(option?: UniNamespace.AuthPrivateMessageOption): void;
    /**
     *
     * 需要基础库： `2.14.4`
     *
     * 在插件中使用：需要基础库 `2.14.4`
     *
     * **仅小程序插件中能调用该接口**，用法同 [uni.authorize](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/authorize/wx.authorize.html)。目前仅支持三种 scope（见下）
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/authorize/wx.authorizeForMiniProgram.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/authorize/wx.authorizeForMiniProgram.html)
     */
    authorizeForMiniProgram(option: UniNamespace.AuthorizeForMiniProgramOption): void;
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：不支持
     *
     * 从本地缓存中异步批量获取指定 key 的内容。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchGetStorage.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchGetStorage.html)
     */
    batchGetStorage(option: UniNamespace.BatchGetStorageOption): void;
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 将数据批量存储在本地缓存中指定的 key 中。会覆盖掉原来该 key 对应的内容。除非用户主动删除或因存储空间原因被系统清理，否则数据都一直可用。单个 key 允许存储的最大数据长度为 1MB，所有数据存储上限为 10MB。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchSetStorage.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchSetStorage.html)
     */
    batchSetStorage(option: UniNamespace.BatchSetStorageOption): void;
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 将数据批量存储在本地缓存中指定的 key 中。会覆盖掉原来该 key 对应的内容。除非用户主动删除或因存储空间原因被系统清理，否则数据都一直可用。单个 key 允许存储的最大数据长度为 1MB，所有数据存储上限为 10MB。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchSetStorageSync.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.batchSetStorageSync.html)
     */
    batchSetStorageSync(kvList: UniNamespace.KvList[]): void;
    /**
     *
     * 需要基础库： `2.13.0`
     *
     * 在插件中使用：不支持
     *
     * 检测是否开启视觉无障碍功能。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/accessibility/wx.checkIsOpenAccessibility.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/accessibility/wx.checkIsOpenAccessibility.html)
     */
    checkIsOpenAccessibility(option?: UniNamespace.CheckIsOpenAccessibilityOption): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 拉起手机通讯录，选择联系人。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/contact/wx.chooseContact.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/contact/wx.chooseContact.html)
     */
    chooseContact(option?: UniNamespace.ChooseContactOption): void;
    /**
     *
     * 需要基础库： `2.3.0`
     *
     * 在插件中使用：需要基础库 `2.16.1`
     *
     * 选择用户已有的发票。
     *
     * **通过 cardId 和 encryptCode 获得报销发票的信息**
     *
     * 请参考[微信电子发票文档](https://mp.weixin.qq.com/wiki?t=resource/res_main&id=21517918939oae3U)中，「查询报销发票信息」部分。
     * 其中 `access_token` 的获取请参考[auth.getAccessToken](#)文档
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/invoice/wx.chooseInvoice.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/invoice/wx.chooseInvoice.html)
     */
    chooseInvoice(option?: UniNamespace.ChooseInvoiceOption): void;
    /**
     *
     * 需要基础库： `2.19.0`
     *
     * 在插件中使用：不支持
     *
     * 选择车牌号
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/license-plate/wx.chooseLicensePlate.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/license-plate/wx.chooseLicensePlate.html)
     */
    chooseLicensePlate(option?: UniNamespace.ChooseLicensePlateOption): void;
    /**
     *
     * 需要基础库： `2.10.0`
     *
     * 在插件中使用：需要基础库 `2.11.1`
     *
     * 拍摄或从手机相册中选择图片或视频。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/video/wx.chooseMedia.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/video/wx.chooseMedia.html)
     */
    chooseMedia(option: UniNamespace.ChooseMediaOption): void;
    /**
     *
     * 需要基础库： `2.5.0`
     *
     * 在插件中使用：不支持
     *
     * 从客户端会话选择文件。
     *
     * ****
     *
     * ```js
     * wx.chooseMessageFile({
     * count: 10,
     * type: 'image',
     * success (res) {
     * // tempFilePath可以作为img标签的src属性显示图片
     * const tempFilePaths = res.tempFiles
     * }
     * })
     * ```
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.chooseMessageFile.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.chooseMessageFile.html)
     */
    chooseMessageFile(option: UniNamespace.ChooseMessageFileOption): void;
    /**
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.choosePoi.html](https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.choosePoi.html)
     */
    choosePoi(option: UniNamespace.ChoosePoiOption): void;
    /**
     *
     * 需要基础库： `2.10.3`
     *
     * 在插件中使用：需要基础库 `2.22.1`
     *
     * 建立本地作为蓝牙低功耗外围设备的服务端，可创建多个。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.createBLEPeripheralServer.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.createBLEPeripheralServer.html)
     */
    createBLEPeripheralServer(option?: UniNamespace.CreateBLEPeripheralServerOption): void;
    /**
     *
     * 需要基础库： `2.26.0`
     *
     * 在插件中使用：不支持
     *
     * 裁剪图片接口
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.cropImage.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.cropImage.html)
     */
    cropImage(option: UniNamespace.CropImageOption): void;
    /**
     *
     * 需要基础库： `2.12.0`
     *
     * 在插件中使用：不支持
     *
     * 关闭小程序页面返回询问对话框。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ui/interaction/wx.disableAlertBeforeUnload.html](https://developers.weixin.qq.com/miniprogram/dev/api/ui/interaction/wx.disableAlertBeforeUnload.html)
     */
    disableAlertBeforeUnload(option?: UniNamespace.DisableAlertBeforeUnloadOption): void;
    /**
     *
     * 需要基础库： `2.22.0`
     *
     * 在插件中使用：不支持
     *
     * 编辑图片接口
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.editImage.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.editImage.html)
     */
    editImage(option: UniNamespace.EditImageOption): void;
    /**
     *
     * 需要基础库： `2.12.0`
     *
     * 在插件中使用：不支持
     *
     * 开启小程序页面返回询问对话框。
     *
     * ## 弹窗条件
     * * 当用户在小程序内非首页页面/最底层页
     * * 官方导航栏上的的返回
     * * 全屏模式下自绘返回键
     * * android 系统 back 键时
     *
     * ## 注意事项
     * * 手势滑动返回时不做拦截
     * * 在任何场景下，此功能都不应拦住用户退出小程序的行为
     *
     * **示例代码**
     *
     * [在微信开发者工具中查看示例](https://developers.weixin.qq.com/s/MTPm9Cmh7VfT)
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ui/interaction/wx.enableAlertBeforeUnload.html](https://developers.weixin.qq.com/miniprogram/dev/api/ui/interaction/wx.enableAlertBeforeUnload.html)
     */
    enableAlertBeforeUnload(option: UniNamespace.EnableAlertBeforeUnloadOption): void;
    /**
     *
     * 需要基础库： `2.17.3`
     *
     * 在插件中使用：不支持
     *
     * 退出当前小程序。必须有点击行为才能调用成功。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.exitMiniProgram.html](https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.exitMiniProgram.html)
     */
    exitMiniProgram(option?: UniNamespace.ExitMiniProgramOption): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.0`
     *
     * 退出（销毁）实时语音通话
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.exitVoIPChat.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.exitVoIPChat.html)
     */
    exitVoIPChat(option?: UniNamespace.ExitVoIPChatOption): void;
    /**
     *
     * 需要基础库： `2.18.0`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.faceDetect.html](https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.faceDetect.html)
     */
    faceDetect(option: UniNamespace.FaceDetectOption): void;
    /**
     *
     * 需要基础库： `2.1.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 获取当前支持的音频输入源
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.getAvailableAudioSources.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.getAvailableAudioSources.html)
     */
    getAvailableAudioSources(option?: UniNamespace.GetAvailableAudioSourcesOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.20.1`
     *
     * 获取蓝牙低功耗的最大传输单元。需在 [uni.createBLEConnection](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.createBLEConnection.html) 调用成功后调用。
     *
     * **注意**
     *
     * - 小程序中 MTU 为 ATT_MTU，包含 Op-Code 和 Attribute Handle 的长度，实际可以传输的数据长度为 `ATT_MTU - 3`
     * - iOS 系统中 MTU 为固定值；安卓系统中，MTU 会在系统协商成功之后发生改变，建议使用 [uni.onBLEMTUChange](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.onBLEMTUChange.html) 监听。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.getBLEMTU.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.getBLEMTU.html)
     */
    getBLEMTU(option: UniNamespace.GetBLEMTUOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.getBackgroundAudioPlayerState.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.getBackgroundAudioPlayerState.html)
     */
    getBackgroundAudioPlayerState(option?: UniNamespace.GetBackgroundAudioPlayerStateOption): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 拉取 backgroundFetch 客户端缓存数据。
     * 当调用接口时，若当次请求未结束，会先返回本地的旧数据（之前打开小程序时请求的），如果本地没有旧数据会返回失败，而不会等待请求完成。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.getBackgroundFetchData.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.getBackgroundFetchData.html)
     */
    getBackgroundFetchData(option: UniNamespace.GetBackgroundFetchDataOption): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 获取设置过的自定义登录态。若无，则返回 fail。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.getBackgroundFetchToken.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.getBackgroundFetchToken.html)
     */
    getBackgroundFetchToken(option?: UniNamespace.GetBackgroundFetchTokenOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：不支持
     *
     * 获取视频号直播信息
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsLiveInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsLiveInfo.html)
     */
    getChannelsLiveInfo(option: UniNamespace.GetChannelsLiveInfoOption): void;
    /**
     *
     * 需要基础库： `2.19.0`
     *
     * 在插件中使用：不支持
     *
     * 获取视频号直播预告信息
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsLiveNoticeInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsLiveNoticeInfo.html)
     */
    getChannelsLiveNoticeInfo(option: UniNamespace.GetChannelsLiveNoticeInfoOption): void;
    /**
     *
     * 需要基础库： `2.22.1`
     *
     * 在插件中使用：不支持
     *
     * 获取视频号直播卡片/视频卡片的分享来源，仅当卡片携带了分享信息、同时用户已授权该小程序获取视频号分享信息且启动场景值为 1177、1184、1195、1208 时可用。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsShareKey.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.getChannelsShareKey.html)
     */
    getChannelsShareKey(option?: UniNamespace.GetChannelsShareKeyOption): void;
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：支持
     *
     * 获取当前的模糊地理位置。
     * ## 使用方法
     * 自 2022 年 7 月 14 日后发布的小程序，若使用该接口，需要在 app.json 中进行声明，否则将无法正常使用该接口，2022年7月14日前发布的小程序不受影响。[具体规则见公告](https://developers.weixin.qq.com/community/develop/doc/000a02f2c5026891650e7f40351c01)
     * ## 申请开通
     * 暂只针对具备与地理位置强相关的使用场景的小程序开放，在小程序管理后台，「开发」-「开发管理」-「接口设置」中自助开通该接口权限。 从2022年7月14日开始在代码审核环节将检测该接口是否已完成开通，如未开通，将在代码提审环节进行拦截。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.getFuzzyLocation.html](https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.getFuzzyLocation.html)
     */
    getFuzzyLocation(option: UniNamespace.GetFuzzyLocationOption): void;
    /**
     *
     * 需要基础库： `2.10.4`
     *
     * 在插件中使用：不支持
     *
     * 获取微信群聊场景下的小程序启动信息。群聊场景包括群聊小程序消息卡片、群待办、群工具。可用于获取当前群的 opengid。
     * ## 注意事项
     * - 基础库 v2.10.4 开始支持获取群工具小程序启动信息
     * - 基础库 v2.17.3 开始支持获取群聊小程序消息卡片、群待办小程序启动信息
     *
     *
     * **Tips**
     *
     * - 如需要展示群名称，小程序可以使用[开放数据组件](https://developers.weixin.qq.com/miniprogram/dev/component/open-data.html)
     * - 小游戏可以通过 `wx.getGroupInfo` 接口获取群名称
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/group/wx.getGroupEnterInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/group/wx.getGroupEnterInfo.html)
     */
    getGroupEnterInfo(option: UniNamespace.GetGroupEnterInfoOption): void;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 判断当前设备是否支持 HCE 能力。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.getHCEState.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.getHCEState.html)
     */
    getHCEState(option?: UniNamespace.GetHCEStateOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 获取局域网IP地址
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.getLocalIPAddress.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.getLocalIPAddress.html)
     */
    getLocalIPAddress(option: UniNamespace.GetLocalIPAddressOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：不支持
     *
     * 获取密码学安全随机数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/crypto/wx.getRandomValues.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/crypto/wx.getRandomValues.html)
     */
    getRandomValues(option: UniNamespace.GetRandomValuesOption): void;
    /**
     *
     * 需要基础库： `2.24.0`
     *
     * 在插件中使用：不支持
     *
     * 查询用户是否在录屏。
     *
     * **示例代码**
     *
     * ```js
     * wx.getScreenRecordingState({
     *  success: function (res) {
     *    console.log(res.state)
     *  },
     * })
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.getScreenRecordingState.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.getScreenRecordingState.html)
     */
    getScreenRecordingState(option?: UniNamespace.GetScreenRecordingStateOption): void;
    /**
     *
     * 需要基础库： `1.1.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 在插件中使用时，只能在当前插件的页面中调用
     *
     * 获取转发详细信息
     *
     *
     * **Tips**
     *
     * - 如需要展示群名称，小程序可以使用 [开放数据组件](https://developers.weixin.qq.com/miniprogram/dev/component/open-data.html)
     * - 小游戏可以通过 [`wx.getGroupInfo`](#) 接口获取群名称
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.getShareInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.getShareInfo.html)
     */
    getShareInfo(option: UniNamespace.GetShareInfoOption): void;
    /**
     *
     * 需要基础库： `2.26.2`
     *
     * 在插件中使用：需要基础库 `2.26.2`
     *
     * 获取当前运行环境对于 [Skyline 渲染引擎](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 的支持情况
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSkylineInfo.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSkylineInfo.html)
     */
    getSkylineInfo(option?: UniNamespace.GetSkylineInfoOption): void;
    /**
     *
     * 需要基础库： `2.14.1`
     *
     * 在插件中使用：不支持
     *
     * 异步获取系统信息。需要一定的微信客户端版本支持，在不支持的客户端上，会使用同步实现来返回。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSystemInfoAsync.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.getSystemInfoAsync.html)
     */
    getSystemInfoAsync(option?: UniNamespace.GetSystemInfoAsyncOption): void;
    /**
     *
     * 需要基础库： `1.2.0`
     *
     * 在插件中使用：不支持
     *
     * 获取用户过去三十天微信运动步数。需要先调用 [uni.login](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.login.html) 接口。步数信息会在用户主动进入小程序时更新。
     *
     *
     * stepInfoList 中，每一项结构如下：
     *
     * | 属性 | 类型 | 说明 |
     * | --- | ---- | --- |
     * | timestamp | number | 时间戳，表示数据对应的时间 |
     * | step | number | 微信运动步数 |
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/werun/wx.getWeRunData.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/werun/wx.getWeRunData.html)
     */
    getWeRunData(option?: UniNamespace.GetWeRunDataOption): void;
    /**
     *
     * 需要基础库： `2.18.0`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.initFaceDetect.html](https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.initFaceDetect.html)
     */
    initFaceDetect(option?: UniNamespace.InitFaceDetectOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.19.1`
     *
     * 查询蓝牙设备是否配对，仅安卓支持。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.isBluetoothDevicePaired.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.isBluetoothDevicePaired.html)
     */
    isBluetoothDevicePaired(option: UniNamespace.IsBluetoothDevicePairedOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：不支持
     *
     * 加入（创建）双人通话。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.join1v1Chat.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.join1v1Chat.html)
     */
    join1v1Chat(option: UniNamespace.Join1v1ChatOption): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.0`
     *
     * 加入 (创建) 实时语音通话，更多信息可见 [实时语音指南](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/voip-chat.html)。调用前需要用户授权 `scope.record`，若房间类型为视频房间需要用户授权 `scope.camera`。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.joinVoIPChat.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.joinVoIPChat.html)
     */
    joinVoIPChat(option: UniNamespace.JoinVoIPChatOption): void;
    /**
     *
     * 需要基础库： `2.12.0`
     *
     * 在插件中使用：需要基础库 `2.12.0`
     *
     * 蓝牙配对接口，仅安卓支持。
     *
     * 通常情况下（需要指定 `pin` 码或者密码时）系统会接管配对流程，直接调用 [uni.createBLEConnection](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.createBLEConnection.html) 即可。该接口只应当在开发者不想让用户手动输入 `pin` 码且真机验证确认可以正常生效情况下用。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.makeBluetoothPair.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.makeBluetoothPair.html)
     */
    makeBluetoothPair(option: UniNamespace.MakeBluetoothPairOption): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 移除小程序切后台事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAppHide.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAppHide.html)
     */
    offAppHide(listener?: UniNamespace.OffAppHideCallback): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 移除小程序切前台事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAppShow.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAppShow.html)
     */
    offAppShow(listener?: UniNamespace.OffAppShowCallback): void;
    /**
     *
     * 需要基础库： `2.6.2`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除音频因为受到系统占用而被中断开始事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAudioInterruptionBegin.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAudioInterruptionBegin.html)
     */
    offAudioInterruptionBegin(listener?: UniNamespace.OffAudioInterruptionBeginCallback): void;
    /**
     *
     * 需要基础库： `2.6.2`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除音频中断结束事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAudioInterruptionEnd.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offAudioInterruptionEnd.html)
     */
    offAudioInterruptionEnd(listener?: UniNamespace.OffAudioInterruptionEndCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除蓝牙低功耗设备的特征值变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLECharacteristicValueChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLECharacteristicValueChange.html)
     */
    offBLECharacteristicValueChange(listener?: UniNamespace.OffBLECharacteristicValueChangeCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除蓝牙低功耗连接状态改变事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLEConnectionStateChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLEConnectionStateChange.html)
     */
    offBLEConnectionStateChange(listener?: UniNamespace.OffBLEConnectionStateChangeCallback): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.20.1`
     *
     * 移除蓝牙低功耗的最大传输单元变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLEMTUChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.offBLEMTUChange.html)
     */
    offBLEMTUChange(listener?: UniNamespace.OffBLEMTUChangeCallback): void;
    /**
     *
     * 需要基础库： `2.10.3`
     *
     * 在插件中使用：需要基础库 `2.22.1`
     *
     * 移除当前外围设备被连接或断开连接事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.offBLEPeripheralConnectionStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.offBLEPeripheralConnectionStateChanged.html)
     */
    offBLEPeripheralConnectionStateChanged(listener?: UniNamespace.OffBLEPeripheralConnectionStateChangedCallback): void;
    /**
     *
     * 需要基础库： `2.8.1`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除 Beacon 服务状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/ibeacon/wx.offBeaconServiceChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/ibeacon/wx.offBeaconServiceChange.html)
     */
    offBeaconServiceChange(listener?: UniNamespace.OffBeaconServiceChangeCallback): void;
    /**
     *
     * 需要基础库： `2.8.1`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除 Beacon 设备更新事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/ibeacon/wx.offBeaconUpdate.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/ibeacon/wx.offBeaconUpdate.html)
     */
    offBeaconUpdate(listener?: UniNamespace.OffBeaconUpdateCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除蓝牙适配器状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.offBluetoothAdapterStateChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.offBluetoothAdapterStateChange.html)
     */
    offBluetoothAdapterStateChange(listener?: UniNamespace.OffBluetoothAdapterStateChangeCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除搜索到新设备的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.offBluetoothDeviceFound.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth/wx.offBluetoothDeviceFound.html)
     */
    offBluetoothDeviceFound(listener?: UniNamespace.OffBluetoothDeviceFoundCallback): void;
    /**
     *
     * 需要基础库： `2.14.3`
     *
     * 在插件中使用：不支持
     *
     * 移除用户点击右上角菜单的「复制链接」按钮时触发的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.offCopyUrl.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.offCopyUrl.html)
     */
    offCopyUrl(listener?: UniNamespace.OffCopyUrlCallback): void;
    /**
     *
     * 需要基础库： `2.9.3`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除设备方向变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.offDeviceMotionChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.offDeviceMotionChange.html)
     */
    offDeviceMotionChange(listener?: UniNamespace.OffDeviceMotionChangeCallback): void;
    /**
     *
     * 需要基础库： `2.8.1`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除接收 NFC 设备消息事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.offHCEMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.offHCEMessage.html)
     */
    offHCEMessage(listener?: UniNamespace.OffHCEMessageCallback): void;
    /**
     *
     * 需要基础库： `2.24.3`
     *
     * 在插件中使用：不支持
     *
     * 移除小程序异步组件加载失败事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offLazyLoadError.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offLazyLoadError.html)
     */
    offLazyLoadError(listener?: UniNamespace.OffLazyLoadErrorCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除 mDNS 服务停止搜索的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceDiscoveryStop.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceDiscoveryStop.html)
     */
    offLocalServiceDiscoveryStop(listener?: UniNamespace.OffLocalServiceDiscoveryStopCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除 mDNS 服务发现的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceFound.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceFound.html)
     */
    offLocalServiceFound(listener?: UniNamespace.OffLocalServiceFoundCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除 mDNS 服务离开的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceLost.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceLost.html)
     */
    offLocalServiceLost(listener?: UniNamespace.OffLocalServiceLostCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 移除 mDNS 服务解析失败的事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceResolveFail.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.offLocalServiceResolveFail.html)
     */
    offLocalServiceResolveFail(listener?: UniNamespace.OffLocalServiceResolveFailCallback): void;
    /**
     *
     * 需要基础库： `2.21.0`
     *
     * 在插件中使用：不支持
     *
     * 移除弱网状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.offNetworkWeakChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.offNetworkWeakChange.html)
     */
    offNetworkWeakChange(listener?: UniNamespace.OffNetworkWeakChangeCallback): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 移除小程序要打开的页面不存在事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offPageNotFound.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.offPageNotFound.html)
     */
    offPageNotFound(listener?: UniNamespace.OffPageNotFoundCallback): void;
    /**
     *
     * 需要基础库： `2.24.0`
     *
     * 在插件中使用：不支持
     *
     * 移除用户录屏事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.offScreenRecordingStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.offScreenRecordingStateChanged.html)
     */
    offScreenRecordingStateChanged(listener?: UniNamespace.OffScreenRecordingStateChangedCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除被动断开实时语音通话事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatInterrupted.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatInterrupted.html)
     */
    offVoIPChatInterrupted(listener?: UniNamespace.OffVoIPChatInterruptedCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除实时语音通话成员在线状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatMembersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatMembersChanged.html)
     */
    offVoIPChatMembersChanged(listener?: UniNamespace.OffVoIPChatMembersChangedCallback): void;
    /**
     *
     * 需要基础库： `2.9.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 移除实时语音通话成员通话状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatSpeakersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatSpeakersChanged.html)
     */
    offVoIPChatSpeakersChanged(listener?: UniNamespace.OffVoIPChatSpeakersChangedCallback): void;
    /**
     *
     * 需要基础库： `2.16.0`
     *
     * 在插件中使用：不支持
     *
     * 移除房间状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPChatStateChanged.html)
     */
    offVoIPChatStateChanged(listener?: UniNamespace.OffVoIPChatStateChangedCallback): void;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：不支持
     *
     * 移除实时语音通话成员视频状态变化事件的监听函数
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPVideoMembersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.offVoIPVideoMembersChanged.html)
     */
    offVoIPVideoMembersChanged(listener?: UniNamespace.OffVoIPVideoMembersChangedCallback): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 监听小程序切后台事件。该事件与 [`App.onHide`](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onhide) 的回调时机一致。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAppHide.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAppHide.html)
     */
    onAppHide(listener: UniNamespace.OnAppHideCallback): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 监听小程序切前台事件。该事件与 [`App.onShow`](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onshowobject-object) 的回调参数一致。
     *
     * **返回有效 referrerInfo 的场景**
     *
     * | 场景值 | 场景                            | appId含义  |
     * | ------ | ------------------------------- | ---------- |
     * | 1020   | 公众号 profile 页相关小程序列表 | 来源公众号 |
     * | 1035   | 公众号自定义菜单                | 来源公众号 |
     * | 1036   | App 分享消息卡片                | 来源App    |
     * | 1037   | 小程序打开小程序                | 来源小程序 |
     * | 1038   | 从另一个小程序返回              | 来源小程序 |
     * | 1043   | 公众号模板消息                  | 来源公众号 |
     *
     * **不同 apiCategory 场景下的 API 限制**
     *
     * `X` 表示 API 被限制无法使用；不在表格中的 API 不限制。
     *
     * |                                       | default | nativeFunctionalized | browseOnly | embedded |
     * |-|-|-|-|-|
     * |navigateToMiniProgram                  |         | `X`                  | `X`        |          |
     * |openSetting                            |         |                      | `X`        |          |
     * |&lt;button open-type="share"&gt;       |         | `X`                  | `X`        | `X`      |
     * |&lt;button open-type="feedback"&gt;    |         |                      | `X`        |          |
     * |&lt;button open-type="open-setting"&gt;|         |                      | `X`        |          |
     * |openEmbeddedMiniProgram                |         | `X`                  | `X`        | `X`      |
     *
     * **注意**
     *
     * 部分版本在无`referrerInfo`的时候会返回 `undefined`，建议使用 `options.referrerInfo && options.referrerInfo.appId` 进行判断。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAppShow.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAppShow.html)
     */
    onAppShow(listener: UniNamespace.OnAppShowCallback): void;
    /**
     *
     * 需要基础库： `2.6.2`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听音频因为受到系统占用而被中断开始事件。以下场景会触发此事件：闹钟、电话、FaceTime 通话、微信语音聊天、微信视频聊天、有声广告开始播放、实名认证页面弹出等。此事件触发后，小程序内所有音频会暂停。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAudioInterruptionBegin.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAudioInterruptionBegin.html)
     */
    onAudioInterruptionBegin(listener: UniNamespace.OnAudioInterruptionBeginCallback): void;
    /**
     *
     * 需要基础库： `2.6.2`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听音频中断结束事件。在收到 onAudioInterruptionBegin 事件之后，小程序内所有音频会暂停，收到此事件之后才可再次播放成功
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAudioInterruptionEnd.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onAudioInterruptionEnd.html)
     */
    onAudioInterruptionEnd(listener: UniNamespace.OnAudioInterruptionEndCallback): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.20.1`
     *
     * 监听蓝牙低功耗的最大传输单元变化事件（仅安卓触发）。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.onBLEMTUChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-ble/wx.onBLEMTUChange.html)
     */
    onBLEMTUChange(listener: UniNamespace.OnBLEMTUChangeCallback): void;
    /**
     *
     * 需要基础库： `2.10.3`
     *
     * 在插件中使用：需要基础库 `2.22.1`
     *
     * 监听当前外围设备被连接或断开连接事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.onBLEPeripheralConnectionStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/bluetooth-peripheral/wx.onBLEPeripheralConnectionStateChanged.html)
     */
    onBLEPeripheralConnectionStateChanged(listener: UniNamespace.OnBLEPeripheralConnectionStateChangedCallback): void;
    /**
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioPause.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioPause.html)
     */
    onBackgroundAudioPause(listener: UniNamespace.OnBackgroundAudioPauseCallback): void;
    /**
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioPlay.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioPlay.html)
     */
    onBackgroundAudioPlay(listener: UniNamespace.OnBackgroundAudioPlayCallback): void;
    /**
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioStop.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.onBackgroundAudioStop.html)
     */
    onBackgroundAudioStop(listener: UniNamespace.OnBackgroundAudioStopCallback): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 监听收到 backgroundFetch 数据事件。如果监听时请求已经完成，则事件不会触发。建议和 [uni.getBackgroundFetchData](https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.getBackgroundFetchData.html) 配合使用
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.onBackgroundFetchData.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.onBackgroundFetchData.html)
     */
    onBackgroundFetchData(listener: UniNamespace.OnBackgroundFetchDataCallback): void;
    /**
     *
     * 需要基础库： `2.14.3`
     *
     * 在插件中使用：不支持
     *
     * 监听用户点击右上角菜单的「复制链接」按钮时触发的事件。本接口为 Beta 版本，暂只在 Android 平台支持。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.onCopyUrl.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.onCopyUrl.html)
     */
    onCopyUrl(listener: UniNamespace.OnCopyUrlCallback): void;
    /**
     *
     * 需要基础库： `2.3.0`
     *
     * 在插件中使用：不支持
     *
     * 监听设备方向变化事件。频率根据 [uni.startDeviceMotionListening()](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.startDeviceMotionListening.html) 的 interval 参数。可以使用 [uni.stopDeviceMotionListening()](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.stopDeviceMotionListening.html) 停止监听。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.onDeviceMotionChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.onDeviceMotionChange.html)
     */
    onDeviceMotionChange(listener: UniNamespace.OnDeviceMotionChangeCallback): void;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 监听接收 NFC 设备消息事件。仅能注册一个监听
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.onHCEMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.onHCEMessage.html)
     */
    onHCEMessage(listener: UniNamespace.OnHCEMessageCallback): void;
    /**
     *
     * 需要基础库： `2.24.3`
     *
     * 在插件中使用：不支持
     *
     * 监听小程序异步组件加载失败事件。
     *
     * **注意**
     *
     * - 加载异步组件通常需要下载分包，若分包下载超时，则会触发 errMsg 为 "loadSubpackage: timeout" 的回调，默认超时等待时间为 5 秒。
     * - 可以通过第二个参数指定超时时间（单位：ms），该设置全局有效，多次指定超时时间则覆盖前面。
     * - 分包确认下载失败时，会再次触发 errMsg 为 "loadSubpackage: fail" 的回调。
     * - 若在页面中使用该接口进行监听，请确保在必要时手动调用 offLazyLoadError 取消监听，以避免非预期的内存泄漏。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onLazyLoadError.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onLazyLoadError.html)
     */
    onLazyLoadError(listener: UniNamespace.OnLazyLoadErrorCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听 mDNS 服务停止搜索的事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceDiscoveryStop.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceDiscoveryStop.html)
     */
    onLocalServiceDiscoveryStop(listener: UniNamespace.OnLocalServiceDiscoveryStopCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听 mDNS 服务发现的事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceFound.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceFound.html)
     */
    onLocalServiceFound(listener: UniNamespace.OnLocalServiceFoundCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听 mDNS 服务离开的事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceLost.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceLost.html)
     */
    onLocalServiceLost(listener: UniNamespace.OnLocalServiceLostCallback): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 监听 mDNS 服务解析失败的事件
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceResolveFail.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.onLocalServiceResolveFail.html)
     */
    onLocalServiceResolveFail(listener: UniNamespace.OnLocalServiceResolveFailCallback): void;
    /**
     *
     * 需要基础库： `2.21.0`
     *
     * 在插件中使用：不支持
     *
     * 监听弱网状态变化事件
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.onNetworkWeakChange.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/network/wx.onNetworkWeakChange.html)
     */
    onNetworkWeakChange(listener: UniNamespace.OnNetworkWeakChangeCallback): void;
    /**
     *
     * 需要基础库： `2.1.2`
     *
     * 在插件中使用：不支持
     *
     * 监听小程序要打开的页面不存在事件。该事件与 [`App.onPageNotFound`](https://developers.weixin.qq.com/miniprogram/dev/reference/api/App.html#onpagenotfoundobject-object) 的回调时机一致。
     *
     * **注意**
     *
     * - 开发者可以在回调中进行页面重定向，但必须在回调中**同步**处理，异步处理（例如 `setTimeout` 异步执行）无效。
     * - 若开发者没有调用 [uni.onPageNotFound](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onPageNotFound.html) 绑定监听，也没有声明 `App.onPageNotFound`，当跳转页面不存在时，将推入微信客户端原生的页面不存在提示页面。
     * - 如果回调中又重定向到另一个不存在的页面，将推入微信客户端原生的页面不存在提示页面，并且不再第二次回调。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onPageNotFound.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/app/app-event/wx.onPageNotFound.html)
     */
    onPageNotFound(listener: UniNamespace.OnPageNotFoundCallback): void;
    /**
     *
     * 需要基础库： `2.24.0`
     *
     * 在插件中使用：不支持
     *
     * 监听用户录屏事件。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.onScreenRecordingStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.onScreenRecordingStateChanged.html)
     */
    onScreenRecordingStateChanged(listener: UniNamespace.OnScreenRecordingStateChangedCallback): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 监听被动断开实时语音通话事件。包括小游戏切入后端时断开
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatInterrupted.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatInterrupted.html)
     */
    onVoIPChatInterrupted(listener: UniNamespace.OnVoIPChatInterruptedCallback): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 监听实时语音通话成员在线状态变化事件。有成员加入/退出通话时触发回调
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatMembersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatMembersChanged.html)
     */
    onVoIPChatMembersChanged(listener: UniNamespace.OnVoIPChatMembersChangedCallback): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 监听实时语音通话成员通话状态变化事件。有成员开始/停止说话时触发回调
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatSpeakersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatSpeakersChanged.html)
     */
    onVoIPChatSpeakersChanged(listener: UniNamespace.OnVoIPChatSpeakersChangedCallback): void;
    /**
     *
     * 需要基础库： `2.16.0`
     *
     * 在插件中使用：不支持
     *
     * 监听房间状态变化事件。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatStateChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPChatStateChanged.html)
     */
    onVoIPChatStateChanged(listener: UniNamespace.OnVoIPChatStateChangedCallback): void;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：不支持
     *
     * 监听实时语音通话成员视频状态变化事件。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPVideoMembersChanged.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.onVoIPVideoMembersChanged.html)
     */
    onVoIPVideoMembersChanged(listener: UniNamespace.OnVoIPVideoMembersChangedCallback): void;
    /**
     *
     * 需要基础库： `1.1.0`
     *
     * 在插件中使用：不支持
     *
     * 查看微信卡包中的卡券。只有通过 [认证](https://developers.weixin.qq.com/miniprogram/product/renzheng.html) 的小程序或文化互动类目的小游戏才能使用。更多文档请参考 [微信卡券接口文档](https://mp.weixin.qq.com/cgi-bin/announce?action=getannouncement&key=1490190158&version=1&lang=zh_CN&platform=2)。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/card/wx.openCard.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/card/wx.openCard.html)
     */
    openCard(option: UniNamespace.OpenCardOption): void;
    /**
     *
     * 需要基础库： `2.19.2`
     *
     * 在插件中使用：不支持
     *
     * 打开视频号视频
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsActivity.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsActivity.html)
     */
    openChannelsActivity(option: UniNamespace.OpenChannelsActivityOption): void;
    /**
     *
     * 需要基础库： `2.21.0`
     *
     * 在插件中使用：不支持
     *
     * 打开视频号活动页
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsEvent.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsEvent.html)
     */
    openChannelsEvent(option: UniNamespace.OpenChannelsEventOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：不支持
     *
     * 打开视频号直播
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsLive.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsLive.html)
     */
    openChannelsLive(option: UniNamespace.OpenChannelsLiveOption): void;
    /**
     *
     * 需要基础库： `2.21.2`
     *
     * 在插件中使用：不支持
     *
     * 打开视频号主页
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsUserProfile.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.openChannelsUserProfile.html)
     */
    openChannelsUserProfile(option: UniNamespace.OpenChannelsUserProfileOption): void;
    /**
     *
     * 需要基础库： `2.19.0`
     *
     * 在插件中使用：不支持
     *
     * 打开微信客服，页面产生点击事件（例如 button 上 bindtap 的回调中）后才可调用。了解更多信息，可以参考[微信客服介绍](https://work.weixin.qq.com/kf/)。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/service-chat/wx.openCustomerServiceChat.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/service-chat/wx.openCustomerServiceChat.html)
     */
    openCustomerServiceChat(option: UniNamespace.OpenCustomerServiceChatOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.26.2`
     *
     * 打开半屏小程序。接入指引请参考 [半屏小程序能力](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/openEmbeddedMiniProgram.html)。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.openEmbeddedMiniProgram.html](https://developers.weixin.qq.com/miniprogram/dev/api/navigate/wx.openEmbeddedMiniProgram.html)
     */
    openEmbeddedMiniProgram(option: UniNamespace.OpenEmbeddedMiniProgramOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 跳转系统蓝牙设置页。仅支持安卓。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.openSystemBluetoothSetting.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/system/wx.openSystemBluetoothSetting.html)
     */
    openSystemBluetoothSetting(option?: UniNamespace.OpenSystemBluetoothSettingOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.pauseBackgroundAudio.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.pauseBackgroundAudio.html)
     */
    pauseBackgroundAudio(option?: UniNamespace.PauseBackgroundAudioOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.pauseVoice.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.pauseVoice.html)
     */
    pauseVoice(option?: UniNamespace.PauseVoiceOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.playBackgroundAudio.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.playBackgroundAudio.html)
     */
    playBackgroundAudio(option: UniNamespace.PlayBackgroundAudioOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.playVoice.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.playVoice.html)
     */
    playVoice(option: UniNamespace.PlayVoiceOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.20.1`
     *
     * __该接口仅在小程序插件中可调用__，调用接口获得插件用户标志凭证（code）。插件可以此凭证换取用于识别用户的标识 openpid。用户不同、宿主小程序不同或插件不同的情况下，该标识均不相同，即当且仅当同一个用户在同一个宿主小程序中使用同一个插件时，openpid 才会相同。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.pluginLogin.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/login/wx.pluginLogin.html)
     */
    pluginLogin(args?: UniNamespace.PluginLoginOption): void;
    /**
     *
     * 需要基础库： `2.22.1`
     *
     * 在插件中使用：不支持
     *
     * 为视图层预加载媒体资源文件, 目前支持：font，image
     *
     *
     * ****
     *
     * - 开发过程中，可在开发者工具network面板查看预加载情况。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadAssets.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadAssets.html)
     */
    preloadAssets(option: UniNamespace.PreloadAssetsOption): void;
    /**
     *
     * 需要基础库： `2.24.7`
     *
     * 在插件中使用：需要基础库 `2.24.7`
     *
     * 预加载下个页面所需要的 [Skyline](https://developers.weixin.qq.com/miniprogram/dev/framework/runtime/skyline/introduction.html) 运行环境。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadSkylineView.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadSkylineView.html)
     */
    preloadSkylineView(option?: UniNamespace.PreloadSkylineViewOption): void;
    /**
     *
     * 需要基础库： `2.15.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 预加载下个页面的 WebView。参见[预加载下个页面的时机](https://developers.weixin.qq.com/miniprogram/dev/framework/performance/tips/runtime_nav.html#_2-4-控制预加载下个页面的时机)
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadWebview.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.preloadWebview.html)
     */
    preloadWebview(option?: UniNamespace.PreloadWebviewOption): void;
    /**
     *
     * 需要基础库： `2.12.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 预览图片和视频。
     *
     * **支持长按识别的码**
     *
     * | 类型 | 说明 | 最低版本 |
     * |------|------| -------|
     * | 小程序码 |    |
     * | 微信个人码 | 不支持小游戏   | [2.18.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) |
     * | 企业微信个人码 | 不支持小游戏   | [2.18.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) |
     * | 普通群码 | 指仅包含微信用户的群，不支持小游戏   | [2.18.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) |
     * | 互通群码 |  指既有微信用户也有企业微信用户的群，不支持小游戏  | [2.18.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) |
     * | 公众号二维码 | 不支持小游戏  | [2.18.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) |
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.previewMedia.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/image/wx.previewMedia.html)
     */
    previewMedia(option: UniNamespace.PreviewMediaOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 在插件中使用时，可以被正常调用，但目前不会进行统计展示
     *
     * 自定义分析数据上报接口。使用前，需要在小程序管理后台自定义分析中新建事件，配置好事件名与字段。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportAnalytics.html](https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportAnalytics.html)
     */
    reportAnalytics(eventName: string, data: any): void;
    /**
     *
     * 需要基础库： `2.14.4`
     *
     * 在插件中使用：不支持
     *
     * 事件上报
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportEvent.html](https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportEvent.html)
     */
    reportEvent(eventId: string, data?: any): void;
    /**
     *
     * 需要基础库： `2.0.1`
     *
     * 在插件中使用：不支持
     *
     * 自定义业务数据监控上报接口。
     *
     * **使用说明**
     *
     * 使用前，需要在「小程序管理后台-运维中心-性能监控-业务数据监控」中新建监控事件，配置监控描述与告警类型。每一个监控事件对应唯一的监控ID，开发者最多可以创建128个监控事件。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportMonitor.html](https://developers.weixin.qq.com/miniprogram/dev/api/data-analysis/wx.reportMonitor.html)
     */
    reportMonitor(name: string, value: number): void;
    /**
     *
     * 需要基础库： `2.9.2`
     *
     * 在插件中使用：需要基础库 `2.9.3`
     *
     * 小程序测速上报。使用前，需要在小程序管理后台配置。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.reportPerformance.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/performance/wx.reportPerformance.html)
     */
    reportPerformance(id: number, value: number, dimensions?: string | any[]): void;
    /**
     *
     * 需要基础库： `2.16.0`
     *
     * 在插件中使用：不支持
     *
     * 仅接入了[自定义版交易组件](https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/Introduction2)的小程序需要使用，普通小程序可直接使用 [`wx.requestPayment`](https://developers.weixin.qq.com/miniprogram/dev/api/payment/wx.requestPayment.html)。
     *
     * **前置检查**
     *
     * 接入自定义版交易组件之后，若要发起微信支付，请先查询[需要校验的场景](https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/order/check_scene)。
     * 在需要校验的场景中，发起微信支付时，必须使用该接口，需要按照要求传入相关的[订单信息](https://developers.weixin.qq.com/miniprogram/dev/framework/ministore/minishopopencomponent2/API/order/add_order)进行校验，校验通过后用户才可以完成当前订单的支付，非需要校验的场景则可以按照商家要求自行选择传入订单信息或不传入。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/payment/wx.requestOrderPayment.html](https://developers.weixin.qq.com/miniprogram/dev/api/payment/wx.requestOrderPayment.html)
     */
    requestOrderPayment(args: UniNamespace.RequestOrderPaymentOption): void;
    /**
     *
     * 需要基础库： `2.22.1`
     *
     * 在插件中使用：需要基础库 `2.22.1`
     *
     * 插件中发起支付。
     *
     * **Tip**
     *
     * 1. `tip`: 小程序与插件绑定在同一个open平台账号上且小程序与插件均为open账号的同主体/关联主体时，调用此接口将直接拉起支付收银台。
     * 1. `tip`: 这个接口本身可以在开发者工具中使用，但功能页的跳转目前不支持在开发者工具中调试，请在真机上测试。
     * 1. `tip`: 跳转支付功能页需要在 `app.json` 中配置 `"functionalPages": true`
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/payment/wx.requestPluginPayment.html](https://developers.weixin.qq.com/miniprogram/dev/api/payment/wx.requestPluginPayment.html)
     */
    requestPluginPayment(option: UniNamespace.RequestPluginPaymentOption): void;
    /**
     *
     * 需要基础库： `2.20.0`
     *
     * 在插件中使用：不支持
     *
     * 订阅设备消息接口，调用后弹出授权框，用户同意后会允许开发者给用户发送订阅模版消息。当用户点击“允许”按钮时，模板消息会被添加到用户的小程序设置页，通过 wx.getSetting 接口可获取用户对相关模板消息的订阅状态。
     *
     * **错误码**
     *
     * | errCode | errMsg                                                 | 说明                                                           |
     * | ------- | ------------------------------------------------------ | -------------------------------------------------------------- |
     * | 10001   | TmplIds can't be empty                                 | tmplIds 为空                                                  |
     * | 10004   | Invalid template id                                    | tmplId 参数类型错误                                            |
     * | 20001   | No template data return, verify the template id exist  | tmplId 为空                                                  |
     * | 20003   | Templates count out of max bounds                      | tmplId 数量超过上限                                           |
     * | 19720726   | check sn_ticket fail                                | snTicket 不合法                                              |
     * | 19720727   | sn_ticket expire                                    | snTicket 过期                                               |
     * | 19720728   | err_not_found_tid                                    | tmplId 不存在                                              |
     * | 19720736   | template_id do not match model_id                   | modelId 类型与 tmplId 类型不符                               |
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeDeviceMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeDeviceMessage.html)
     */
    requestSubscribeDeviceMessage(option: UniNamespace.RequestSubscribeDeviceMessageOption): void;
    /**
     *
     * 需要基础库： `2.4.4`
     *
     * 在插件中使用：不支持
     *
     * 调起客户端小程序订阅消息界面，返回用户订阅消息的操作结果。当用户勾选了订阅面板中的“总是保持以上选择，不再询问”时，模板消息会被添加到用户的小程序设置页，通过 [uni.getSetting](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/setting/wx.getSetting.html) 接口可获取用户对相关模板消息的订阅状态。
     *
     * ## 注意事项
     * - 一次性模板 id 和永久模板 id 不可同时使用。
     * - 低版本基础库2.4.4~2.8.3 已支持订阅消息接口调用，仅支持传入一个一次性 tmplId / 永久 tmplId。
     * - [2.8.2](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) 版本开始，用户发生点击行为或者发起支付回调后，才可以调起订阅消息界面。
     * - [2.10.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) 版本开始，开发版和体验版小程序将禁止使用模板消息 formId。
     * - 一次授权调用里，每个tmplId对应的模板标题不能存在相同的，若出现相同的，只保留一个。
     * - [2.10.0](https://developers.weixin.qq.com/miniprogram/dev/framework/compatibility.html) 版本开始，支持订阅语音消息提醒，[详情](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/subscribe-message.html)
     *
     * **错误码**
     *
     * | errCode | errMsg                                                 | 说明                                                           |
     * | ------- | ------------------------------------------------------ | -------------------------------------------------------------- |
     * | 10001   | TmplIds can't be empty                                 | 参数传空了                                                     |
     * | 10002   | Request list fail                                       | 网络问题，请求消息列表失败                                     |
     * | 10003   | Request subscribe fail                                 | 网络问题，订阅请求发送失败                                     |
     * | 10004   | Invalid template id                                    | 参数类型错误                                                   |
     * | 10005   | Cannot show subscribe message UI                       | 无法展示 UI，一般是小程序这个时候退后台了导致的                |
     * | 20001   | No template data return, verify the template id exist  | 没有模板数据，一般是模板 ID 不存在 或者和模板类型不对应 导致的 |
     * | 20002   | Templates type must be same                            | 模板消息类型 既有一次性的又有永久的                            |
     * | 20003   | Templates count out of max bounds                      | 模板消息数量超过上限                                           |
     * | 20004   | The main switch is switched off                        | 用户关闭了主开关，无法进行订阅                                 |
     * | 20005   | This mini program was banned from subscribing messages | 小程序被禁封                                                   |
     * | 20013   | Reject DeviceMsg Template                              | 不允许通过该接口订阅设备消息                                      |
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/subscribe-message/wx.requestSubscribeMessage.html)
     */
    requestSubscribeMessage(option: UniNamespace.RequestSubscribeMessageOption): void;
    /**
     *
     * 需要基础库： `2.19.0`
     *
     * 在插件中使用：不支持
     *
     * 预约视频号直播
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.reserveChannelsLive.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/channels/wx.reserveChannelsLive.html)
     */
    reserveChannelsLive(option: UniNamespace.ReserveChannelsLiveOption): void;
    /**
     *
     * 需要基础库： `2.14.0`
     *
     * 在插件中使用：不支持
     *
     * 根据 URL 销毁存在内存中的数据
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.revokeBufferURL.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/wx.revokeBufferURL.html)
     */
    revokeBufferURL(url: string): void;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 保存文件系统的文件到用户磁盘，仅在 PC 端支持
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/file/wx.saveFileToDisk.html](https://developers.weixin.qq.com/miniprogram/dev/api/file/wx.saveFileToDisk.html)
     */
    saveFileToDisk(option: UniNamespace.SaveFileToDiskOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.seekBackgroundAudio.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.seekBackgroundAudio.html)
     */
    seekBackgroundAudio(option: UniNamespace.SeekBackgroundAudioOption): void;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 发送 NFC 消息。仅在安卓系统下有效。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.sendHCEMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.sendHCEMessage.html)
     */
    sendHCEMessage(option: UniNamespace.SendHCEMessageOption): void;
    /**
     *
     * 需要基础库： `2.25.0`
     *
     * 在插件中使用：不支持
     *
     * 拉起手机发送短信界面。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/sms/wx.sendSms.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/sms/wx.sendSms.html)
     */
    sendSms(option: UniNamespace.SendSmsOption): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 设置自定义登录态，在周期性拉取数据时带上，便于第三方服务器验证请求合法性
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.setBackgroundFetchToken.html](https://developers.weixin.qq.com/miniprogram/dev/api/storage/background-fetch/wx.setBackgroundFetchToken.html)
     */
    setBackgroundFetchToken(option: UniNamespace.SetBackgroundFetchTokenOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：不支持
     *
     * 开启双人通话。设置 `enable` 为 `false` 时，无法接听呼叫。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.setEnable1v1Chat.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.setEnable1v1Chat.html)
     */
    setEnable1v1Chat(option: UniNamespace.SetEnable1v1ChatOption): void;
    /**
     *
     * 需要基础库： `2.3.0`
     *
     * 在插件中使用：需要基础库 `2.10.0`
     *
     * 设置 [InnerAudioContext](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/InnerAudioContext.html) 的播放选项。设置之后对当前小程序全局生效。
     *
     * ****
     *
     * ## 注意事项
     * - 为保证微信整体体验，speakerOn 为 true 时，客户端会忽略 mixWithOthers 参数的内容，强制与其它音频互斥
     * - 不支持在播放音频的过程中切换为扬声器播放，开发者如需切换可以先暂停当前播放的音频并记录下当前暂停的时间点，然后切换后重新从原来暂停的时间点开始播放音频
     * - 目前 wx.setInnerAudioOption 接口不兼容 wx.createWebAudioContext 接口，也不兼容 wx.createInnerAudioContext 开启 useWebAudioImplement 的情况，将在后续版本中支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.setInnerAudioOption.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.setInnerAudioOption.html)
     */
    setInnerAudioOption(option: UniNamespace.SetInnerAudioOption): void;
    /**
     *
     * 需要基础库： `1.4.3`
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ui/sticky/wx.setTopBarText.html](https://developers.weixin.qq.com/miniprogram/dev/api/ui/sticky/wx.setTopBarText.html)
     */
    setTopBarText(option: UniNamespace.SetTopBarTextOption): void;
    /**
     *
     * 需要基础库： `2.20.1`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 设置截屏/录屏时屏幕表现，仅支持在 Android 端调用
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.setVisualEffectOnCapture.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/screen/wx.setVisualEffectOnCapture.html)
     */
    setVisualEffectOnCapture(option: UniNamespace.SetVisualEffectOnCaptureOption): void;
    /**
     *
     * 需要基础库： `1.6.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 设置 `wifiList` 中 AP 的相关信息。在 `onGetWifiList` 回调后调用，**iOS特有接口**。
     *
     * **注意**
     *
     * - 该接口只能在 `onGetWifiList` 回调之后才能调用。
     * - 此时客户端会挂起，等待小程序设置 Wi-Fi 信息，请务必尽快调用该接口，若无数据请传入一个空数组。
     * - 有可能随着周边 Wi-Fi 列表的刷新，单个流程内收到多次带有存在重复的 Wi-Fi 列表的回调。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/wifi/wx.setWifiList.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/wifi/wx.setWifiList.html)
     */
    setWifiList(option: UniNamespace.SetWifiListOption): void;
    /**
     *
     * 需要基础库： `2.10.1`
     *
     * 在插件中使用：不支持
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ui/window/wx.setWindowSize.html](https://developers.weixin.qq.com/miniprogram/dev/api/ui/window/wx.setWindowSize.html)
     */
    setWindowSize(option: UniNamespace.SetWindowSizeOption): void;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：不支持
     *
     * 转发文件到聊天
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.shareFileMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.shareFileMessage.html)
     */
    shareFileMessage(option: UniNamespace.ShareFileMessageOption): void;
    /**
     *
     * 在插件中使用：不支持
     *
     * 分享数据到微信运动。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/werun/wx.shareToWeRun.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/werun/wx.shareToWeRun.html)
     */
    shareToWeRun(option: UniNamespace.ShareToWeRunOption): void;
    /**
     *
     * 需要基础库： `2.16.1`
     *
     * 在插件中使用：不支持
     *
     * 转发视频到聊天
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.shareVideoMessage.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.shareVideoMessage.html)
     */
    shareVideoMessage(option: UniNamespace.ShareVideoMessageOption): void;
    /**
     *
     * 需要基础库： `2.10.0`
     *
     * 在插件中使用：不支持
     *
     * 拉取h5领取红包封面页。获取参考红包封面地址参考 [微信红包封面开发平台](https://cover.weixin.qq.com/cgi-bin/mmcover-bin/readtemplate?t=page/index#/doc?page=introduce)。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/open-api/redpackage/wx.showRedPackage.html](https://developers.weixin.qq.com/miniprogram/dev/api/open-api/redpackage/wx.showRedPackage.html)
     */
    showRedPackage(option: UniNamespace.ShowRedPackageOption): void;
    /**
     *
     * 需要基础库： `2.14.3`
     *
     * 在插件中使用：需要基础库 `2.16.0`
     *
     * 打开分享图片弹窗，可以将图片发送给朋友、收藏或下载
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.showShareImageMenu.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.showShareImageMenu.html)
     */
    showShareImageMenu(option: UniNamespace.ShowShareImageMenuOption): void;
    /**
     *
     * 需要基础库： `2.3.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 开始监听设备方向的变化。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.startDeviceMotionListening.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.startDeviceMotionListening.html)
     */
    startDeviceMotionListening(option?: UniNamespace.StartDeviceMotionListeningOption): void;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 初始化 NFC 模块。（HCE 模式仅安卓支持）
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.startHCE.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.startHCE.html)
     */
    startHCE(option: UniNamespace.StartHCEOption): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 开始搜索局域网下的 mDNS 服务。搜索的结果会通过 wx.onLocalService* 事件返回。
     *
     * **注意**
     *
     * 1. 由于操作系统相关能力变更，iOS 微信客户端 7.0.18 及以上版本无法使用 mDNS 相关接口，安卓版本不受影响
     * 2. wx.startLocalServiceDiscovery 是一个消耗性能的行为，开始 30 秒后会自动 stop 并执行 wx.onLocalServiceDiscoveryStop 注册的回调函数。
     * 3. 在调用 wx.startLocalServiceDiscovery 后，在这次搜索行为停止后才能发起下次 wx.startLocalServiceDiscovery。停止本次搜索行为的操作包括调用 wx.stopLocalServiceDiscovery 和 30 秒后系统自动 stop 本次搜索。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.startLocalServiceDiscovery.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.startLocalServiceDiscovery.html)
     */
    startLocalServiceDiscovery(option: UniNamespace.StartLocalServiceDiscoveryOption): void;
    /**
     *
     * 需要基础库： `2.8.0`
     *
     * 在插件中使用：不支持
     *
     * 开启小程序进入前后台时均接收位置消息，需引导用户开启[授权](https://developers.weixin.qq.com/miniprogram/dev/framework/open-ability/authorize.html#后台定位)。授权以后，小程序在运行中或进入后台均可接受位置消息变化。
     *  ## 使用方法
     * 自 2022 年 7 月 14 日后发布的小程序，若使用该接口，需要在 app.json 中进行声明，否则将无法正常使用该接口，2022年7月14日前发布的小程序不受影响。[具体规则见公告](https://developers.weixin.qq.com/community/develop/doc/000a02f2c5026891650e7f40351c01)
     *
     *  ## 申请开通
     *  暂只针对如下类目的小程序开放，需要先通过类目审核，再在小程序管理后台，「开发」-「开发管理」-「接口设置」中自助开通该接口权限。从2022年7月14日开始，在代码审核环节将检测该接口是否已完成开通，如未开通，将在代码提审环节进行拦截。
     *
     *  ### 国内主体开放类目
     *
     *  | 一级类目/主体类型	| 二级类目	| 应用场景 |
     *  | -------------- | -------| -------- |
     *  | 电商平台 |	/	| 在小程序内提供线下商超导览、导航服务 |
     *  | 商家自营 |	/	| 在小程序内提供线下商超导览、导航服务 |
     *  | 交通服务 |	/	| 代驾服务、打车出行、城市共享交通、实时导航服务等 |
     *  | 生活服务 |	跑腿、共享服务	| 含有B端小程序配送服务，基于地理位置共享工具类服务  |
     *  | 物流服务 |	收件/派件、查件、邮政、装卸搬运、快递柜、货物运输	 |	提供B端小程序快递/货物收发服务 |
     *  | 餐饮服务 | 	点餐平台、外卖平台	| 提供B端小程序餐饮配送服务、线下门店实时导航 |
     *  | 工具	| 健康管理 |	基于实时地理位置提供身体管理记录等服务 |
     *  | 旅游 | 景区服务、住宿服务 | 在小程序内提供景区导航、导览服务、酒店导航服务 |
     *  | 政务民生 |	/	 | 提供政务单位相关业务 |
     *  | 政府主体帐号 |	/	| 提供政务单位相关业务 |
     *
     *  ### 海外主体开放类目
     *  | 一级类目/主体类型	| 二级类目	| 应用场景 |
     *  | -------------- | -------| -------- |
     *  | 交通服务	 | /	 | 代驾服务、打车出行、城市共享交通、实时导航服务等 |
     *  | 生活服务	 | 家政、外送	 | 含有B端小程序配送服务，基于地理位置导航上门服务 |
     *  | 快递业与邮政	 | /	 | 提供B端小程序快递/货物收发服务 |
     *  | 餐饮服务	 | 外卖点餐	 | 提供B端小程序餐饮配送服务、线下门店实时导航 |
     *  | 电商平台	 | /	 | 在小程序内提供线下商超导览、导航服务 |
     *  | 跨境电商	 | /	 | 在小程序内提供线下商超导览、导航服务 |
     *  | 本地服务	 | 服装/鞋/箱包、玩具、家电/数码/手机、美妆/洗护、珠宝/饰品/眼镜/钟表、运动/户外/乐器、鲜花/园艺/工艺品、家居/家饰/家纺、办公/文具、机械/电子器件、酒、食品、百货/超市/便利店、宠物食品/用品	 | 在小程序内提供线下商超导览、导航服务 |
     *
     * **注意**
     *
     * - 安卓微信7.0.6版本，iOS 7.0.5版本起支持该接口
     * - 需在app.json中配置requiredBackgroundModes: ['location']后使用
     * - 获取位置信息需配置[地理位置用途说明](https://developers.weixin.qq.com/miniprogram/dev/reference/configuration/app.html#permission)。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.startLocationUpdateBackground.html](https://developers.weixin.qq.com/miniprogram/dev/api/location/wx.startLocationUpdateBackground.html)
     */
    startLocationUpdateBackground(option: UniNamespace.StartLocationUpdateBackgroundOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/recorder/wx.startRecord.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/recorder/wx.startRecord.html)
     */
    startRecord(option?: UniNamespace.WxStartRecordOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.stopBackgroundAudio.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/background-audio/wx.stopBackgroundAudio.html)
     */
    stopBackgroundAudio(option?: UniNamespace.StopBackgroundAudioOption): void;
    /**
     *
     * 需要基础库： `2.3.0`
     *
     * 在插件中使用：需要基础库 `2.9.1`
     *
     * 停止监听设备方向的变化。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.stopDeviceMotionListening.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/motion/wx.stopDeviceMotionListening.html)
     */
    stopDeviceMotionListening(option?: UniNamespace.StopDeviceMotionListeningOption): void;
    /**
     *
     * 需要基础库： `2.18.0`
     *
     * 在插件中使用：需要基础库 `2.21.3`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.stopFaceDetect.html](https://developers.weixin.qq.com/miniprogram/dev/api/ai/face/wx.stopFaceDetect.html)
     */
    stopFaceDetect(option?: UniNamespace.StopFaceDetectOption): void;
    /**
     *
     * 需要基础库： `1.7.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 关闭 NFC 模块。仅在安卓系统下有效。
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.stopHCE.html](https://developers.weixin.qq.com/miniprogram/dev/api/device/nfc-hce/wx.stopHCE.html)
     */
    stopHCE(option?: UniNamespace.StopHCEOption): void;
    /**
     *
     * 需要基础库： `2.4.0`
     *
     * 在插件中使用：需要基础库 `2.15.0`
     *
     * 停止搜索 mDNS 服务
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.stopLocalServiceDiscovery.html](https://developers.weixin.qq.com/miniprogram/dev/api/network/mdns/wx.stopLocalServiceDiscovery.html)
     */
    stopLocalServiceDiscovery(option?: UniNamespace.StopLocalServiceDiscoveryOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/recorder/wx.stopRecord.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/recorder/wx.stopRecord.html)
     */
    stopRecord(option?: UniNamespace.WxStopRecordOption): void;
    /**
     *
     * 在插件中使用：需要基础库 `1.9.6`
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.stopVoice.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/audio/wx.stopVoice.html)
     */
    stopVoice(option?: UniNamespace.StopVoiceOption): void;
    /**
     *
     * 需要基础库： `2.11.0`
     *
     * 在插件中使用：需要基础库 `2.11.0`
     *
     * 订阅视频画面成员。对于视频房间，当成员超过两人时需进行订阅，否则只能看到最先加入房间的两人画面。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.subscribeVoIPVideoMembers.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.subscribeVoIPVideoMembers.html)
     */
    subscribeVoIPVideoMembers(option: UniNamespace.SubscribeVoIPVideoMembersOption): void;
    /**
     *
     * 需要基础库： `1.2.0`
     *
     * 在插件中使用：需要基础库 `2.1.0`
     *
     * 在插件中使用时，只能在当前插件的页面中调用
     *
     * 更新转发属性
     *
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.updateShareMenu.html](https://developers.weixin.qq.com/miniprogram/dev/api/share/wx.updateShareMenu.html)
     */
    updateShareMenu(option: UniNamespace.UpdateShareMenuOption): void;
    /**
     *
     * 需要基础库： `2.7.0`
     *
     * 在插件中使用：需要基础库 `2.9.0`
     *
     * 更新实时语音静音设置
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.updateVoIPChatMuteConfig.html](https://developers.weixin.qq.com/miniprogram/dev/api/media/voip/wx.updateVoIPChatMuteConfig.html)
     */
    updateVoIPChatMuteConfig(option: UniNamespace.UpdateVoIPChatMuteConfigOption): void;
    /**
     *
     * 需要基础库： `2.12.0`
     *
     * 在插件中使用：需要基础库 `2.12.0`
     *
     * 更新客户端版本。当判断用户小程序所在客户端版本过低时，可使用该接口跳转到更新微信页面。
     *
     * 文档: [https://developers.weixin.qq.com/miniprogram/dev/api/base/update/wx.updateWeChatApp.html](https://developers.weixin.qq.com/miniprogram/dev/api/base/update/wx.updateWeChatApp.html)
     */
    updateWeChatApp(option?: UniNamespace.UpdateWeChatAppOption): void;
}
