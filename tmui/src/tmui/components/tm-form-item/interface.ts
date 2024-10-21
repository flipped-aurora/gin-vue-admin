export interface rulesItem {
    validator?: Function | boolean,//检验函数。可以是异步回调。
    required?: boolean,//是否必填。
    message?: string,//检验不合格时的文本
    type?: string,//校验类型.
}
export interface inputPushItem {
    value: any,
    isRequiredError: boolean,//true,错误，false正常 检验状态
    componentsName: string,//表单组件类型。
    message: string,//检验信息提示语。

}