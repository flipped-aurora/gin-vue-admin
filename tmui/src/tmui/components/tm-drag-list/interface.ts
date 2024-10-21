export interface itemList {
    _id?: number | string,//内部标识
    top?: number | string,
    i?: number | string,
    text?: number | string,//项目文字
    icon?: string,//项目图标，左边
    color?: string,//文字颜色
    [prop: string]: any
}