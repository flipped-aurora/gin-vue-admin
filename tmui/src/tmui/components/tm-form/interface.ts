
interface rulesItem {
    validator?: Function | boolean,//检验函数。可以是异步回调。
    required?: boolean,//是否必填。
    message?: string,//检验不合格时的文本
    type?: string,//校验类型.
}
export interface formItem {
    label: string,//标签名称。
    field: string,//字段名称key.
    value: any,
    isRequiredError: boolean,//true,错误，false正常 检验状态
    message: string,//检验信息提示语。
    id: number | string,//表单唯一标识id
    componentsName: string,//表单组件类型。
    rules: rulesItem[]
}

export interface validateResultListType {
    message: string,
    validator: boolean, //通过为true,失败为false
    field: string,//字段名称key.
    [key: string]: any
}