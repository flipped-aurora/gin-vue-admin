
export interface DataTableColumn {
    title: string,//列表的标题。默认为：""
    key: string,//这个key需要和tabdata中的key相同，表示同一列。
    width: number,//列表宽
    align: string,//对齐方向，start左,center中,end右,默认center
    sort?: boolean,//是否显示排序,默认false
    bgColor: string,//当前头的背景色。默认grey
    cellColor: string,//当前列的背景色。,如果为"",则使用行数据的背景，如果行背景也没有提供，使用white.
    light: boolean,//背景色是否是浅色
    sortType: string,//desc降序，升序asce,none,无排序
    unit: string,//单位
    headerFontSize: number, // 头部字体大小
    fontSize: number, // 内容字体大小
    headerProps: tSheetProps, // 头部tSheet props
    cellProps: tSheetProps, // 内容 tSheet props
    _headerProps?: _tSheetProps,
    _cellProps?: _tSheetProps,
    minWidth: number,
    ellipsis: boolean, //文本超出隐藏
    slot?: string
}
export interface tSheetProps {
    parenClass: string,
    contStyle: string,
    height: number,
    width: number,
    color: string,
    transprent: boolean | string,
    border: number | string,
    margin: any,
    padding: any,
    unit: string,
    hoverClass: string,
    darkBgColor: string,
    noLevel: boolean,
    blur: boolean,
    _style: object | string,
    _class: object | string,
    followTheme: boolean | string,
    dark: boolean | string,
    followDark: boolean | string,
    round: number,
    shadow: number,
    outlined: boolean,
    borderStyle: string,
    borderDirection: string,
    text: boolean | string,
    linear: string,
    linearDeep: string,
    isDisabledRoundAndriod: boolean | string,
}

export interface _tSheetProps {
    width: number,
    height: number,
    fontSize: number
}