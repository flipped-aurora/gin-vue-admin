declare namespace UniCloudNamespace {
  interface Interceptor {
    invoke?: (result: any) => void;
    success?: (result: any) => void;
    fail?: (result: any) => void;
    complete?: (result: any) => void;
  }

  interface BaseObjectInterceptorArgs {
    objectName: string;
    methodName: string;
    params: string;
  }

  interface SuccessObjectInterceptorArgs extends BaseObjectInterceptorArgs {
    result: any;
  }

  interface FailObjectInterceptorArgs extends BaseObjectInterceptorArgs {
    error: UniError;
  }

  type CompleteObjectInterceptorArgs = SuccessObjectInterceptorArgs | FailObjectInterceptorArgs;

  interface ObjectInterceptor {
    invoke?: (result: BaseObjectInterceptorArgs) => void;
    success?: (result: SuccessObjectInterceptorArgs) => void;
    fail?: (result: FailObjectInterceptorArgs) => void;
    complete?: (result: CompleteObjectInterceptorArgs) => void;
  }

  interface UniCloud {
    /**
     * 添加拦截器
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#add-interceptor](https://uniapp.dcloud.io/uniCloud/client-sdk.html#add-interceptor)
     */
    addInterceptor(apiName: string, interceptor: Interceptor): void;
    /**
     * 移除拦截器
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#remove-interceptor](https://uniapp.dcloud.io/uniCloud/client-sdk.html#remove-interceptor)
     */
    removeInterceptor(apiName: string, interceptor?: Interceptor): void;
    /**
     * 拦截云对象请求
     *
     * 文档: [https://uniapp.dcloud.io/uniCloud/client-sdk.html#intercept-object](https://uniapp.dcloud.io/uniCloud/client-sdk.html#intercept-object)
     */
    interceptObject(interceptor: ObjectInterceptor): void;
  }
}
