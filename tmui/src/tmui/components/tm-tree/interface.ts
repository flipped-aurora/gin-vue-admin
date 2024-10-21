export interface baseNodeData {
    icon?: string,//节点图标。
    color?: string,//节点颜色主题
    disabled?: boolean,//节点是否禁用
    text: string,//节点标题
    id: number | string,//节点标识
    checked?: boolean,
    expanded?: boolean //是否父节点打开。
}