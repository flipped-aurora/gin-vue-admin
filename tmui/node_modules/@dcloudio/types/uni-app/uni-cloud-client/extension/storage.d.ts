declare namespace UniCloudNamespace {
  interface GeneralCallbackResult {
    /**
     * 状态码，操作成功则不返回
     */
    errCode: string;
    /**
     * 错误信息
     */
    errMsg: string;
  }

  interface OnUploadProgressResult {
    /**
     * 已上传大小
     */
    loaded: number;
    /**
     * 上传文件总大小
     */
    total: number;
  }

  interface UploadFileResult {
    /**
     * 文件唯一 ID，用来访问文件，建议存储起来
     */
    fileID: string;
  }

  interface UploadFileOptions {
    /**
     * 文件的绝对路径，包含文件名。例如 foo/bar.jpg、foo/bar/baz.jpg 等
     */
    cloudPath: string;
    /**
     * 要上传的文件对象
     */
    filePath: string;
    /**
     * 阿里云是否以cloudPath作为实际存储路径，默认为false（否）
     */
    cloudPathAsRealPath?: boolean;
    /**
     * 上传进度回调
     */
    onUploadProgress?: (result: OnUploadProgressResult) => void;
    /**
     * 成功返回的回调函数
     */
    success?: (result: UploadFileResult) => void;
    /**
     * 失败返回的回调函数
     */
    fail?: (result: GeneralCallbackResult) => void;
    /**
     * 结束的回调函数（调用成功、失败都会执行
     */
    complete?: (result: UploadFileResult) => void;
  }

  interface DeleteFileItem {
    /**
     * 云端fileID
     */
    fileID: string;
    /**
     * 状态码，操作成功则不返回
     */
    code: string;
  }

  interface DeleteFileResult {
    /**
     * 要删除的文件 ID 组成的数组
     */
    fileList: DeleteFileItem[];
  }

  interface DeleteFileOptions {
    /**
     * 文件ID组成的数组
     */
    fileList: any[];
    /**
     * 成功返回的回调函数
     */
    success?: (result: DeleteFileResult) => void;
    /**
     * 失败返回的回调函数
     */
    fail?: (result: GeneralCallbackResult) => void;
    /**
     * 结束的回调函数（调用成功、失败都会执行
     */
    complete?: (result: DeleteFileResult) => void;
  }

  interface GetTempFileURLItem {
    /**
     * 文件 ID
     */
    fileID: string;
    /**
     * 状态码，操作成功则为 SUCCESS
     */
    code: string;
    /**
     * 文件访问链接
     */
    tempFileURL: string;
  }

  interface GetTempFileURLResult {
    /**
     * 存储下载链接的数组
     */
    fileList: any[];
  }

  interface GetTempFileURLOptions {
    /**
     * 文件ID组成的数组
     */
    fileList: any[];
    /**
     * 成功返回的回调函数
     */
    success?: (result: GetTempFileURLResult) => void;
    /**
     * 失败返回的回调函数
     */
    fail?: (result: GeneralCallbackResult) => void;
    /**
     * 结束的回调函数（调用成功、失败都会执行
     */
    complete?: (result: GetTempFileURLResult) => void;
  }

  interface ChooseAndUploadFileSuccessCallbackResult {
    /**
     * 错误信息
     */
    errMsg: string;
    /**
     * 临时文件路径列表
     */
    tempFilePaths: any[];
    /**
     * 文件列表，每一项是一个 File 对象
     */
    tempFiles: any[];
  }

  interface ChooseAndUploadFileOnUploadProgressCallbackResult {
    /**
     * 触发当前上传进度回调的文件序号
     */
    index: number;
    /**
     * 已上传大小
     */
    loaded: number;
    /**
     * 总大小
     */
    total: number;
    /**
     * 临时文件路径
     */
    tempFilePath: string;
    /**
     * 文件对象
     */
    tempFile: any;
  }

  interface ChooseAndUploadFileOnChooseFileCallbackResult {
    /**
     * 错误信息
     */
    errMsg: string;
    /**
     * 临时文件路径列表
     */
    tempFilePaths: any[];
    /**
     * 文件列表，每一项是一个 File 对象
     */
    tempFiles: any[];
  }

  interface ChooseAndUploadFileOptions {
    /**
     * 文件类型
     * - image: 图片
     * - video: 视频
     * - all: 任意文件
     */
    type: 'image' | 'video' | 'all';
    /**
     * 文件数量
     */
    count?: number;
    /**
     * 允许的文件后缀数组
     */
    extension?: any[];
    /**
     * original 原图，compressed 压缩图，默认二者都有
     */
    sizeType?: string | string[];
    /**
     * album 从相册选图，camera 使用相机，默认二者都有
     */
    sourceType?: string | string[];
    /**
     * 摄像切换
     * - front: 前置摄像头
     * - back: 后置摄像头
     */
    camera?: 'front' | 'back';
    /**
     * 是否压缩所选的视频源文件，默认值为true，需要压缩
     */
    compressed?: boolean;
    /**
     * 拍摄视频最长拍摄时间，单位秒。最长支持 60 秒
     */
    maxDuration?: number;
    /**
     * 选择文件后的回调
     */
    onChooseFile?: (result: ChooseAndUploadFileOnChooseFileCallbackResult) => void;
    /**
     * 上传进度回调
     */
    onUploadProgress?: (result: ChooseAndUploadFileOnUploadProgressCallbackResult) => void;
    /**
     * 接口调用成功的回调函数
     */
    success?: (result: ChooseAndUploadFileSuccessCallbackResult) => void;
    /**
     * 接口调用失败的回调函数
     */
    fail?: (result: any) => void;
    /**
     * 接口调用结束的回调函数（调用成功、失败都会执行）
     */
    complete?: (result: any) => void;
  }

  interface GetFileInfoOptions {
    fileList: Array<string>;
  }

  interface GetFileInfoResponseFileItem {
    /**
     * 文件ID
     */
    fileId: string;
    /**
     * 文件上传时间（秒）
     */
    gmtCreate: number;
    /**
     * 文件最近更改时间（秒）
     */
    gmtModified: number;
    /**
     * 文件原始名称
     */
    name: string;
    /**
     * 文件大小（Byte）
     */
    size: number;
    /**
     * 文件类型
     */
    type: string;
    /**
     * 文件CDN加速访问下载链接
     */
    url: string;
  }

  interface GetFileInfoResponse {
    fileList: Array<GetFileInfoResponseFileItem>;
  }

  interface UniCloud {
    /**
     * 上传文件到云端
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/storage?id=uploadfile](https://uniapp.dcloud.io/uniCloud/storage?id=uploadfile)
     */
    uploadFile(options: UploadFileOptions): Promise<any>;
    /**
     * 选择并上传文件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/storage?id=chooseanduploadfile](https://uniapp.dcloud.io/uniCloud/storage?id=chooseanduploadfile)
     */
    chooseAndUploadFile(options: ChooseAndUploadFileOptions): Promise<any>;
    /**
     * 删除云端文件
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/storage?id=deletefile](https://uniapp.dcloud.io/uniCloud/storage?id=deletefile)
     */
    deleteFile(options: DeleteFileOptions): Promise<any>;
    /**
     * 获取文件临时链接
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/storage?id=gettempfileurl](https://uniapp.dcloud.io/uniCloud/storage?id=gettempfileurl)
     */
    getTempFileURL(options: GetTempFileURLOptions): Promise<any>;
    /**
     * 获取文件信息，阿里云专用
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/storage?id=get-file-info](https://uniapp.dcloud.io/uniCloud/storage?id=get-file-info)
     */
    getFileInfo(options: GetFileInfoOptions): Promise<GetFileInfoResponse>;
  }
}
