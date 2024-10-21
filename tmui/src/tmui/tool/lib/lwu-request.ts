/**
 * 显示loading
 * @param params - 参数
 * - title - 标题
 * - mask - 是否显示透明蒙层
 */
export const loading = (params: {
    title: string;
    mask?: boolean;
}) => {
    uni.showLoading({
        title: params.title,
        mask: params.mask || true
    })
}

/**
 * 消息提示框
 * @param params - 参数
 * - title - 标题
 * - duration - 提示的延迟时间
 * @returns {void}
 */
export const msg = (params: {
    title: string;
    duration?: number;
}) => {
    uni.showToast({
        title: params.title,
        icon: 'none',
        image: '',
        mask: false,
        duration: params.duration || 2000
    })
}

/**
 * 网络请求参数配置
 */
export interface RequestOptionsConfig {
    baseUrl?: {
        /**
         * 开发环境域名
         */
        dev: string;
        /**
         * 生产环境域名
         */
        pro: string;
    };
    /**
     * 调试模式，开启后控制台会显示内部调试打印信息
     */
    debug?: boolean;
    /**
     * 请求过程是否显示loading
     */
    loading?: boolean;
    /**
     * 请求中loading弹窗的提示文本，默认为 `'请求中...'`
     */
    loadingText?: string;
    /**
     * 本次请求的队列ID，需要过滤重复请求、手动中断拦截的场景使用
     */
    task_id?: string;
    /**
     * 自定义请求前拦截
     */
    before?: Function;
    /**
     * 自定义请求后拦截
     */
    after?: Function;
    /**
     * 自定义请求头
     */
    header?: object;
    // 下面配置项参考 [uniapp 网络请求API](https://uniapp.dcloud.io/api/request/request.html)
    /**
     * 请求超时时间，单位ms
     */
    timeout?: number;
    /**
     * 请求方式，有效值：`'GET'`、`'POST'`、`'PUT'`、`'DELETE'`、`'CONNECT'`、`'HEAD'`、`'OPTIONS'`、`'TRACE'`
     */
    method?: 'GET' | 'POST' | 'PUT' | 'DELETE' | 'CONNECT' | 'HEAD' | 'OPTIONS' | 'TRACE';
    /**
     * 如果设为 json，会对返回的数据进行一次 JSON.parse，非 json 不会进行 JSON.parse
     */
    dataType?: string;
    /**
     * 设置响应的数据类型。合法值：`text`、`arraybuffer`
     */
    responseType?: string;
    /**
     * 验证 ssl 证书
     */
    sslVerify?: boolean;
    /**
     * 跨域请求时是否携带凭证（cookies）
     */
    withCredentials?: boolean;
    /**
     * DNS解析时优先使用ipv4
     */
    firstIpv4?: boolean;
    /**
     * 业务错误代码拦截处理程序，请根据业务实际情况灵活设置
     * @param code 
     * @param errMsg 
     * @returns 
     */
    errorHandleByCode?: (code: number, errMsg?: string) => void;
    /**
     * API错误拦截处理程序，请根据业务实际情况灵活设置
     * + `1.1.0` 及以上版本支持
     * 
     * @param data API返回内容
     * @param args uniapp请求API回调结果
     */
    apiErrorInterception?: (data: any, args?: UniApp.RequestSuccessCallbackResult) => void;
    /**
     * 网络异常或者断网处理程序，建议更新缓存中是否断网或者网络繁忙的标识以便前端页面展示没有网络或者断网的通用异常页面
     * @returns
     */
    networkExceptionHandle?: () => void;
    /**
     * 请求成功时接口响应描述信息字段名称，默认为 `'msg'`
     */
    requestSuccessResponseMsgName?: string;
    /**
     * 缓存中token字段名称，方便请求库从缓存获取token完成自动填充token
     */
    tokenStorageKeyName?: string;
    /**
     * 自定义获取token处理程序，通过promise返回最新token值即可
     * + `1.0.2` 及以上版本支持
     * @returns 
     */
    tokenValue?: () => Promise<unknown>;
    /**
     * 自定义构建URL参数方式，即用什么方式把请求的params对象数据转为`a=1&b=2`的格式，默认使用NodeJS内置对象 `URLSearchParams` 转化，可以自定义通过 `qs` 插件的方式转化
     * + `1.0.2` 及以上版本支持
     * 
     * @example
     * ```ts
     * // qs 插件转化示例
     * import qs from 'qs';
     * 
     * return qs.stringify(obj);
     * ```
     */
    buildQueryString?: (obj: object) => string;
    /**
     * 请求携带token的方式，有效值：`header`、`body`
     */
    takeTokenMethod?: 'header' | 'body';
    /**
     * 请求携带token的字段名称，header方式默认为 `Authorization`
     */
    takenTokenKeyName?: string;
    /**
     * 是否自动刷新token
     */
    autoRefreshToken?: boolean;
    /**
     * 自动刷新token程序，返回promise，`autoRefreshToken` 为 `true`时生效
     */
    refreshTokenHandle?: () => Promise<unknown>;
    /**
     * 自定义token失效的错误代码，便于请求库内部做自动刷新token判断
     */
    tokenExpiredCode?: number;
    /**
     * 请求失败是否自动重试
     */
    retry?: boolean;
    /**
     * 请求失败自动重试次数
     */
    retryCount?: number;
    /**
     * 请求失败重试次数是否自动计算，失败重试次数上限依然是设置的retryCount值
     */
    retryCountAutoOffRetry?: boolean;
    /**
     * 请求失败用来生成重试时间上限（指数退避算法需要），单位秒
     */
    retryMaximum?: number;
    /**
     * 请求失败执行重试时间上限（指数退避算法需要），达到上限后不再重试
     */
    retryDeadline?: number;
}

/**
 * 这里是原来的全局配置和请求配置未分开的请求配置类型
 */
export interface RequestOptions {
    task_id?: string;
    before?: Function;
    after?: Function;
    header?: object;
    method?: "GET" | "OPTIONS" | "HEAD" | "POST" | "PUT" | "DELETE" | "TRACE" | "CONNECT";
    timeout?: number;
    dataType?: string;
    responseType?: string;
    sslVerify?: boolean;
    withCredentials?: boolean;
    firstIpv4?: boolean;
    retryCount?: number;
};

/**
 * @param {number} times 重试次数
 * @param {number} maximum_offretry 最大等待秒数
 * @returns {number}
 * 指数退避算法简介：
 * 为了解决如何设置适当的重传等待时间而存在的算法，基本流程如下：
 * - 1.客户端发起请求
 * - 2.请求失败，等待1 + random_number_milliseconds秒后重试请求。
 * - 3.请求失败，等待2 + random_number_milliseconds秒后重试请求。
 * - 4.请求失败，等待4 + random_number_milliseconds秒后重试请求。
 * - 5.以此类推，直至达到设置的等待时间上限为止结束请求，具体算法公式如下：
 *  Math.min((2 ** n + ranom_number_milliseconds), maxium_backoff)  
 * 上述的random_number_milliseconds为1到1000的随机毫秒数
 */
const makeRetryTimeout = (times: number, maximum_offretry: number): number => {
    const random_number_milliseconds = Math.floor(Math.random() * 1000);
    return Math.min(Math.pow(2, times) * 1000 + random_number_milliseconds, maximum_offretry);
}

/**
 * 对象转query string的参数字符串
 * @param obj 需要转化的对象参数
 */
const objToQueryString = (obj: object): string => {
    if (typeof obj === 'object' && obj !== null) {
        return Object.keys(obj)
            .map((key) => `${key}=${encodeURIComponent((obj as any)[key])}`)
            .join('&');
    }

    return JSON.stringify(obj);
}

/**
 * 网络请求库封装
 * @public
 */
export class Http {
    /**
     * 当前请求任务
     */
    private currentRequestTask: UniApp.RequestTask = {
        abort: () => { },
        onHeadersReceived: () => { },
        offHeadersReceived: () => { }
    };
    private requestTasksName = 'LWU-REQUEST-TASKS';
    /**
     * 请求锁
     */
    private lock: boolean = true;
    /**
     * 请求列表
     */
    private pending: Function[] = [];
    /**
     * 请求失败自动重试次数
     */
    private retryCount: number = 3;
    /**
     * 请求失败用来生成重试时间上限（指数退避算法需要），单位秒
     */
    private retryMaximum: number = 64;
    /**
     * 重试等待时间列表
     */
    private retryTimeout: (number | undefined)[] = [];
    /**
     * 重试等待时间上限
     */
    private retryDeadline: number = 10000;
    /**
     * 配置信息
     */
    private config: RequestOptionsConfig = {
        baseUrl: {
            pro: '',
            dev: ''
        },
        /**
         * 业务错误代码拦截处理程序，请根据业务实际情况灵活设置
         * @param code 
         * @param errMsg 
         * @returns 
         */
        errorHandleByCode: (code: number, errMsg?: string) => { }
    };

    /**
     * 获取请求配置
     * @author KviewUI <kviewui@163.com>
     * @param {RequestOptionsConfig} config - 配置信息
     * @returns 
     */
    private useConfig = (config: RequestOptionsConfig) => {
        return {
            baseUrl: config.baseUrl ?? {
                pro: '',
                dev: ''
            },
            /**
             * 调试模式，开启后控制台会显示内部调试打印信息
             */
            debug: config.debug ?? false,
            /**
             * 请求过程是否显示loading
             */
            loading: config.loading ?? true,
            /**
             * 请求中loading弹窗的提示文本，默认为 `'请求中...'`
             */
            loadingText: config.loadingText ?? '请求中...',
            /**
             * 本次请求的队列ID，需要过滤重复请求、手动中断拦截的场景使用
             */
            task_id: config.task_id ?? '',
            /**
             * 自定义请求前拦截
             */
            before: config.before,
            /**
             * 自定义请求后拦截
             */
            after: config.after,
            /**
             * 自定义请求头
             */
            header: config.header,
            // 下面配置项参考 [uniapp 网络请求API](https://uniapp.dcloud.io/api/request/request.html)
            /**
             * 请求超时时间，单位`ms`
             */
            timeout: config.timeout ?? 6000,
            /**
             * 请求方式，有效值：`'GET'`、`'POST'`、`'PUT'`、`'DELETE'`、`'CONNECT'`、`'HEAD'`、`'OPTIONS'`、`'TRACE'`
             */
            method: config.method ?? 'GET',
            /**
             * 如果设为 json，会对返回的数据进行一次 JSON.parse，非 json 不会进行 JSON.parse
             */
            dataType: config.dataType ?? 'json',
            /**
             * 设置响应的数据类型。合法值：`'text'`、`'arraybuffer'`
             */
            responseType: config.responseType ?? 'text',
            /**
             * 验证 ssl 证书
             */
            sslVerify: config.sslVerify ?? true,
            /**
             * 跨域请求时是否携带凭证（cookies）
             */
            withCredentials: config.withCredentials ?? false,
            /**
             * DNS解析时优先使用ipv4
             */
            firstIpv4: config.firstIpv4 ?? false,
            /**
             * 业务错误代码拦截处理程序，请根据业务实际情况灵活设置
             * @param code 
             * @param errMsg 
             * @returns 
             */
            errorHandleByCode: config.errorHandleByCode,
            /**
             * API错误拦截处理程序，请根据业务实际情况灵活设置
             * @param data API返回内容
             * @param args uniapp请求API回调结果
             */
            apiErrorInterception: config.apiErrorInterception,
            /**
             * 网络异常或者断网处理程序，建议更新缓存中是否断网或者网络繁忙的标识以便前端页面展示没有网络或者断网的通用异常页面
             * @returns
             */
            networkExceptionHandle: () => { },
            /**
             * 请求成功时接口响应描述信息字段名称，默认为 `'msg'`
             */
            requestSuccessResponseMsgName: config.requestSuccessResponseMsgName ?? 'msg',
            /**
             * 缓存中token字段名称，方便请求库从缓存获取token完成自动填充token
             */
            tokenStorageKeyName: config.tokenStorageKeyName ?? '',
            /**
             * 自定义获取token处理程序，通过promise返回最新token值即可
             * @returns 
             * @example
             * ```ts
             * tokenValue: () => {
             *      return new Promise((resolve, _) => {
             *          // 获取最新token演示
             *          const token = getToken();
             *          token && resolve(token);
             *      });
             * }
             * ```
             */
            tokenValue: config.tokenValue,
            /**
             * 自定义构建URL参数方式，即用什么方式把请求的params对象数据转为`a=1&b=2`的格式，默认使用NodeJS内置对象 `URLSearchParams` 转化，可以自定义通过 `qs` 插件的方式转化
             * 
             * @example
             * ```ts
             * // qs 插件转化示例
             * import qs from 'qs';
             * 
             * return qs.stringify(obj);
             * ```
             */
            buildQueryString: config.buildQueryString,
            /**
             * 请求携带token的方式，有效值：header、body
             */
            takeTokenMethod: config.takeTokenMethod ?? 'header',
            /**
             * 请求携带token的字段名称，header方式默认为 `Authorization`
             */
            takenTokenKeyName: config.takenTokenKeyName ?? 'Authorization',
            /**
             * 是否自动刷新token
             */
            autoRefreshToken: false,
            /**
             * 自动刷新token程序，返回promise，`autoRefreshToken` 为 `true`时生效
             */
            refreshTokenHandle: config.refreshTokenHandle,
            /**
             * 自定义token失效的错误代码，便于请求库内部做自动刷新token判断
             */
            tokenExpiredCode: config.tokenExpiredCode ?? 403,
            /**
             * 请求失败是否自动重试
             */
            retry: config.retry ?? false,
            /**
             * 请求失败自动重试次数
             */
            retryCount: config.retryCount ?? 3,
            /**
             * 请求失败重试次数是否自动计算，失败重试次数上限依然是设置的retryCount值
             */
            retryCountAutoOffRetry: config.retryCountAutoOffRetry ?? true,
            /**
             * 请求失败用来生成重试时间上限（指数退避算法需要），单位秒
             */
            retryMaximum: config.retryMaximum ?? 64,
            /**
             * 请求失败执行重试时间上限（指数退避算法需要），达到上限后不再重试
             */
            retryDeadline: config.retryDeadline ?? 10000
        }
    }

    /**
     * 初始化配置内容
     * @param config 
     */
    private initConfig = (config: RequestOptionsConfig) => {
        this.config = {
            ...this.useConfig(config)
        };

        if (!this.config.retry) {
            this.retryCount = 0;
        } else {
            if (this.config.retryCountAutoOffRetry) {
                this.retryMaximum = (this.config.retryMaximum as number) * 1000;
                this.retryTimeout = [];
                this.retryDeadline = this.config.retryDeadline as number;

                for (let i = 0; i < this.retryCount; i++) {
                    if (this.retryDeadline < 0) {
                        break;
                    }
                    const timeout = makeRetryTimeout(i, this.retryMaximum);
                    this.retryDeadline -= timeout;
                    this.retryTimeout.push(timeout);
                }

                this.retryCount = this.retryTimeout.length;
            }
        }
    }

    /**
     * 请求失败的错误统一处理
     * @param code - 错误码
     * @param message - 自定义错误信息 
     */
    private handleError(code: number, message: string = ''): void {
        // 调用错误状态码处理程序
        if (this.config.errorHandleByCode) {
            this.config.errorHandleByCode(code, message);
        } else {
            msg({ title: message });
        }
    }

    private interceptor(url: string, before: Function | undefined, after: Function | undefined) {
        uni.addInterceptor('request', {
            invoke: (args) => {
                // 请求前拦截处理
                if (this.config.debug) {
                    console.warn(`【LwuRequest Debug:请求拦截】${JSON.stringify(args)}`);
                }

                if (this.config.loading) {
                    loading({ title: this.config.loadingText ?? '请求中...' });
                }

                if (args?.header?.contentType) {
                    args.header['content-type'] = args.header.contentType;
                    delete args.header.contentType;
                }

                // 拼接baseURI
                let baseURI: string|undefined = '';
                if (process.env.NODE_ENV === 'development') {
                    baseURI = this.config.baseUrl && this.config.baseUrl.dev;
                    // debug = this.config.debug as boolean;
                } else {
                    baseURI = this.config.baseUrl && this.config.baseUrl.pro;
                }

                let reqUrl = `${baseURI}${url}`;
                if (args.method === 'GET') {
                    args.data = this.config.buildQueryString && this.config.buildQueryString(args.data)
                        ? this.config.buildQueryString(args.data)
                        // : new URLSearchParams(Object.entries(args.data)).toString();
                        : objToQueryString(args.data);
                    args.url = `${reqUrl}?${args.data}`;
                } else {
                    args.url = reqUrl;
                }

                // 请求前自定义拦截
                if (before) {
                    before();
                }
            },
            // 响应拦截
            success: (args: UniApp.RequestSuccessCallbackResult) => {
                this.handleError(args.statusCode, (args.data as AnyObject)[this.config.requestSuccessResponseMsgName as string]);

                this.config.apiErrorInterception && this.config.apiErrorInterception(args.data, args);

                if (this.config.debug) {
                    console.warn(`【LwuRequest Debug:响应拦截】${JSON.stringify(args)}`);
                }

                if (after) {
                    after();
                }
            },
            fail: (err: UniApp.GeneralCallbackResult) => {
                if (this.config.debug) {
                    console.warn(`【LwuRequest Debug:请求拦截失败】${JSON.stringify(err)}`);
                }
            },
            complete: (res: UniApp.GeneralCallbackResult) => {
                uni.hideLoading();
                if (this.config.debug) {
                    console.warn(`【LwuRequest Debug:请求拦截完成】${JSON.stringify(res)}`);
                }
            }
        });
    }

    /**
     * 刷新token处理
     */
    private refreshToken() {
        if (this.config.refreshTokenHandle) {
            this.config.refreshTokenHandle()
                .then(() => {
                    // 重新执行业务请求
                    uni.getStorageSync('LWU-REQUEST-CALLBACK')((callback: () => void) => {
                        callback();
                    })
                })
                .catch(() => {
                    // token失效
                    this.handleError(this.config.tokenExpiredCode as number);
                });
        }
    }

    public request(url: string, data: any = {}, options: RequestOptionsConfig = {
        header: {},
        method: this.config.method,
        timeout: this.config.timeout,
        dataType: this.config.dataType,
        responseType: this.config.responseType,
        sslVerify: this.config.sslVerify,
        withCredentials: this.config.withCredentials,
        firstIpv4: this.config.firstIpv4,
        retryCount: this.config.retryCount
    }) {
        // 初始化配置
        this.initConfig(options);
        // 判断该请求队列是否存在，如果存在则中断请求
        const requestTasks = uni.getStorageSync(this.requestTasksName);

        if (options?.task_id && requestTasks[options?.task_id]) {
            if (this.config.debug) {
                console.warn(`【LwuRequest Debug】请求ID${options.task_id}有重复项已自动过滤`);
            }

            requestTasks[options?.task_id]?.abort();
        }

        return new Promise((resolve, reject) => {
            // 拦截器
            this.interceptor(url, options.before, options.after);
            // let header: any = {};

            // 判断是否存在token，如果存在则在请求头统一添加token，token获取从config配置获取
            let token = uni.getStorageSync(this.config.tokenStorageKeyName as string);

            const setToken = () => {
                return new Promise((resolve, _) => {
                    token && resolve(token);

                    if (this.config.tokenValue) {
                        this.config.tokenValue().then(res => {
                            res && resolve(res);
                            resolve(false);
                        })
                    } else {
                        resolve('');
                    }
                });
            }

            setToken().then(getToken => {
                if (getToken) {
                    if (this.config.takeTokenMethod === 'header') {
                        (options.header as any)[this.config.takenTokenKeyName as string] = getToken;
                    }

                    if (this.config.takeTokenMethod === 'body') {
                        data[this.config.takenTokenKeyName as string] = getToken;
                    }
                }

                // 发起请求
                this.currentRequestTask = uni.request({
                    url: url,
                    data: data,
                    // header: reqHeader.header,
                    header: {
                        ...options.header
                    },
                    method: options.method,
                    timeout: options.timeout,
                    dataType: options.dataType,
                    responseType: options.responseType,
                    sslVerify: options.sslVerify,
                    withCredentials: options.withCredentials,
                    firstIpv4: options.firstIpv4,
                    success: (res: UniApp.RequestSuccessCallbackResult) => {
                        if (res.statusCode !== this.config.tokenExpiredCode) {
                            resolve(res.data);
                        } else {
                            // 刷新token
                            this.refreshToken();
                            uni.setStorageSync('LWU-REQUEST-CALLBACK', () => {
                                resolve(this.request(url, data, options));
                            });
                        }
                    },
                    fail: (err: UniApp.GeneralCallbackResult) => {
                        this.retryCount = options.retryCount ?? 3;

                        if (this.retryCount === 0) {
                            reject(err);
                        } else {
                            if (this.config.debug) {
                                console.warn(`【LwuRequest Debug】自动重试次数：${this.retryCount}`);
                            }
                            this.retryCount--;
                            setTimeout(this.request, this.retryTimeout.shift());
                            // 网络异常或者断网处理
                            this.config.networkExceptionHandle && this.config.networkExceptionHandle();
                        }
                    }
                });

                // 判断是否设置请求队列ID
                if (options?.task_id) {
                    // 当前请求存入缓存
                    let tasks: UniApp.RequestTask[] = [];
                    tasks[options?.task_id as any] = this.currentRequestTask;
                    uni.setStorageSync(this.requestTasksName, tasks);
                }
            });
        });
    }

    public get(url: string, data: object = {}, options: RequestOptions = {}) {
        return this.request(url, data, {
            method: 'GET',
            ...options
        });
    }

    public post(url: string, data: object = {}, options: RequestOptions = {}) {
        return this.request(url, data, {
            method: 'POST',
            ...options
        });
    }

    public put(url: string, data: object = {}, options: RequestOptions = {}) {
        return this.request(url, data, {
            method: 'POST',
            ...options
        });
    }

    public delete(url: string, data: object = {}, options: RequestOptions = {}) {
        return this.request(url, data, {
            method: 'DELETE',
            ...options
        });
    }

    /**
     * 中断请求，不传 `task_id` 时默认中断当前任务
     * @param task_id 
     */
    public abort(task_id: string = '') {
        const requestTask = uni.getStorageSync(this.requestTasksName);

        if (requestTask[task_id]) {
            requestTask[task_id].abort();
        } else {
            this.currentRequestTask.abort();
        }
    }
}

