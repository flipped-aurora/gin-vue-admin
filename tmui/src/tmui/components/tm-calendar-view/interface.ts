export interface monthDayItem {
    dateStr: string,//日期
    date: number | string,//当前的日。
    day: number,//星期几。
    week: number,//当前在年中的第几周。
    isNowIn: boolean,//是否是本月。
    disabled: boolean,//是否被禁用。
    extra: dateItemStyle
}
export interface dateItemStyle {
    date?: string,//日期
    text?: boolean,//浅色背景。
    color?: string,//主题色.
    extra?: string,//额外的内容，在日期下方显示的文本。
}

export interface monthYearItem {
    dateStr: string,//日期
    month: number,//月份
    isVaild: boolean,//是否在允许选中范围内。
    dateStr2?: string,
    year?: string | number

}

export interface weekItem {
    week: number,//周次
    dateStr: string,//日期
    dateStr2?: string,
    date: number | string,
    day: number,//星期几。
    isNowIn: boolean,//是否是本月。
    isVaild: boolean,//是否在允许选中范围内。
}

export interface yearItem {
    dateStr: string,//日期
    year: number,//月份
    isVaild: boolean,//是否在允许选中范围内。
}