//子组件向父组件报告的信息类型。
export interface itemParenSG {
    id: number,//子组件id，唯一。
    width: number,
    height: number,
    top: number,
    bottom: number,
    left: number,
    imgWidth: number,
    imgHeight: number,
    index: number,
}