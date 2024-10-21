export interface columnsItem {
    children?: Array<columnsItem>,//当前列表的数据列表
    text?: string | number,
    id?: string | number,
    disabled?: boolean,//禁用后，将滑动时选中不了本项。
    // defaultIndex?:number//当前列表的默认选中的值。
    [key: string]: any;
}