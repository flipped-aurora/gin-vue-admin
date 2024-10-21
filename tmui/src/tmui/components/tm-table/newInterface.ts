export const defaultCellStyle: tabaleCellStyleType = {
    type: 'text',
    color: 'white',
    fontColor: 'black',
    fontSize: 26,
    light: false,
    transparent: true,
    asyncStyleCell: false,
    sort: false,
}
export interface tabaleCellData {
    value: string | number,
    opts: tabaleCellStyleType,
    [key: string]: any;
}
export interface tabaleCellStyleType {
    type?: 'button' | 'text',
    color?: string,
    fontColor?: string,
    fontSize?: number,
    light?: boolean,
    transparent?: boolean,
    /**是否头和所在列同步同的背景色和文字色,注意该参数只在header中的opts有效 */
    asyncStyleCell?: boolean,
    /**该列是否显示 排序功能，注意该参数只在header中的opts有效 */
    sort?: boolean,

}
/** 表头数据格式 */
export interface headerType {
    /**字段变量名*/
    field: string,
    /**字段名称 */
    name: string,
    opts?: tabaleCellStyleType,
    [key: string]: any;
}
export interface tableDataType {
    /**列字段名称 */
    fields: {
        columns: string[]
    },
    /**头数据,对应fields中columns字段 */
    header: Array<headerType>,
    /** 表格数据 */
    data: Array<{
        opts?: {
            [key: string]: tabaleCellStyleType
        },
        [key: string]: any;
    }>
}
