export type HttpTask = UniApp.RequestTask | UniApp.UploadTask | UniApp.DownloadTask;

export type HttpRequestTask = UniApp.RequestTask;

export type HttpUploadTask = UniApp.UploadTask;

export type HttpDownloadTask = UniApp.DownloadTask;

export type HttpMethod =
    "GET"
    | "POST"
    | "PUT"
    | "DELETE"
    | "CONNECT"
    | "HEAD"
    | "OPTIONS"
    | "TRACE"
    | "UPLOAD"
    | "DOWNLOAD";

export type HttpRequestHeader = Record<string, string>;

export type HttpParams = Record<string, any>;

export type HttpData = Record<string, any>;

export type HttpResponseType = 'arraybuffer' | 'text';

export type HttpCustom = Record<string, any>;

export type HttpFileType = 'image' | 'video' | 'audio';

export type HttpFormData = Record<string, any>;

export type HttpResponseHeader = Record<string, string> & {
    "set-cookie"?: string[]
};

export interface HttpRequestConfig<T = HttpTask> {
    /** @desc 请求服务器接口地址 */
    url?: string;
    /** @desc 请求方式，默认为 GET */
    method?: HttpMethod;
    /** @desc 请求基地址 */
    baseURL?: string;
    /** @desc 请求头信息，不能设置 Referer，App、H5 端会自动带上 cookie，且 H5 端不可手动修改 */
    header?: HttpRequestHeader;
    /** @desc 请求查询参数，自动拼接为查询字符串 */
    params?: HttpParams;
    /** @desc 请求体参数 */
    data?: HttpData;
    /** @desc 超时时间，单位 ms，默认为 60000，仅 H5 (HBuilderX 2.9.9+)、APP (HBuilderX 2.9.9+)、微信小程序 (2.10.0)、支付宝小程序支持 */
    timeout?: number;
    /** @desc 跨域请求时是否携带凭证 (cookies)，默认为 false，仅 H5 (HBuilderX 2.6.15+) 支持 */
    withCredentials?: boolean;
    /** @desc 设置响应的数据类型，支付宝小程序不支持 */
    responseType?: HttpResponseType;
    /** @desc 全局自定义验证器 */
    validateStatus?: ((statusCode: number) => boolean) | null;


    /** params 参数自定义处理 */
    paramsSerializer?: (params: AnyObject) => string | void;

    /** @desc 默认为 json，如果设为 json，会尝试对返回的数据做一次 JSON.parse */
    dataType?: string;
    /** @desc DNS 解析时是否优先使用 ipv4，默认为 false，仅 App-Android (HBuilderX 2.8.0+) 支持 */
    firstIpv4?: boolean;
    /** @desc 是否验证 SSL 证书，默认为 true，仅 App-Android (HBuilderX 2.3.3+) 支持 */
    sslVerify?: boolean;

    /** @desc 开启 http2;微信小程序 */
    enableHttp2?: boolean;

    /** @desc 开启 quic；微信小程序 */
    enableQuic?: boolean;
    /** @desc 开启 cache;微信小程序、字节跳动小程序 2.31.0+ */
    enableCache?: boolean;
    /** @desc 开启 httpDNS;微信小程序 */
    enableHttpDNS?: boolean;
    /** @desc httpDNS 服务商；微信小程序 */
    httpDNSServiceId?: string;
    /** @desc 开启 transfer-encoding chunked;微信小程序 */
    enableChunked?: boolean;
    /** @desc wifi下使用移动网络发送请求;微信小程序 */
    forceCellularNetwork?: boolean;
    /** @desc 开启后可在headers中编辑cookie;支付宝小程序 10.2.33+ */
    enableCookie?: boolean;
    /** @desc 是否开启云加速;百度小程序 3.310.11+ */
    cloudCache?: boolean | object;
    /** @desc 控制当前请求是否延时至首屏内容渲染后发送；百度小程序 3.310.11+ */
    defer?: boolean;

    /** @desc 自定义参数 */
    custom?: HttpCustom;

    /** @desc 返回当前请求的 task 和 options，不要在这里修改 options */
    getTask?: (task: T, options: HttpRequestConfig<T>) => void;

    /** @desc 需要上传的文件列表，使用 files 时，filePath 和 name 不生效，仅支持 App、H5 (2.6.15+) */
    files?: { name?: string; file?: File; uri: string; }[];
    /** @desc 文件类型，仅支付宝小程序支持且为必填项 */
    fileType?: HttpFileType;
    /** @desc 要上传的文件对象，仅 H5 (2.6.15+) 支持 */
    file?: File;
    /** @desc 要上传文件资源的路径，使用 files 时，filePath 和 name 不生效 */
    filePath?: string;
    /** @desc 文件对应的 key，开发者在服务器端通过这个 key 可以获取到文件二进制内容，使用 files 时，filePath 和 name 不生效 */
    name?: string;
    /** @desc 请求中其他额外的 form data */
    formData?: HttpFormData;
}

export interface HttpResponse<T = any, D = HttpTask> {
    data: T;
    statusCode: number;
    header: HttpResponseHeader;
    config: HttpRequestConfig<D>;
    cookies: string[];
    errMsg: string;
    rawData: any;
}

export interface HttpUploadResponse<T = any, D = HttpTask> {
    data: T;
    statusCode: number;
    config: HttpRequestConfig<D>;
    errMsg: string;
    rawData: any;
}

export interface HttpDownloadResponse extends HttpResponse {
    tempFilePath: string;
    apFilePath?: string;
    filePath?: string;
    fileContent?: string;
}

export interface HttpError<T = any, D = HttpTask> {
    data?: T;
    statusCode?: number;
    header?: HttpResponseHeader;
    config: HttpRequestConfig<D>;
    cookies?: string[];
    errMsg: string;
}

export interface HttpPromise<T = any> extends Promise<HttpResponse<T>> {
}

export interface HttpInterceptorManager<V, E = V> {
    use(onFulfilled?: (value: V) => V | Promise<V>, onRejected?: (error: E) => T | Promise<E>): void;

    eject(id: number): void;
}

export abstract class HttpRequestAbstract {
    constructor(config?: HttpRequestConfig);

    interceptors: {
        request: HttpInterceptorManager<HttpRequestConfig>;
        response: HttpInterceptorManager<HttpResponse, HttpError>;
    }

    request<T = any, R = HttpResponse<T>, D = HttpRequestTask>(config: HttpRequestConfig<D>): Promise<R>;

    get<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, config?: HttpRequestConfig<D>): Promise<R>;

    delete<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    head<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    options<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    post<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    put<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    config: HttpRequestConfig;

    setConfig<D = HttpTask>(onSend: (config: HttpRequestConfig<D>) => HttpRequestConfig<D>): void;

    connect<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    trace<T = any, R = HttpResponse<T>, D = HttpRequestTask>(url: string, data?: HttpData, config?: HttpRequestConfig<D>): Promise<R>;

    upload<T = any, R = HttpUploadResponse<T>, D = HttpUploadTask>(url: string, config?: HttpRequestConfig<D>): Promise<R>;

    download<T = any, R = HttpDownloadResponse<T>, D = HttpDownloadTask>(url: string, config?: HttpRequestConfig<D>): Promise<R>;

    middleware<T = any, R = HttpResponse<T>, D = HttpTask>(config: HttpRequestConfig<D>): Promise<R>;
}

declare class HttpRequest extends HttpRequestAbstract {
}

export default HttpRequest;
