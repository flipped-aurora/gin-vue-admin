import { linearDeep, linearDirection } from "../../tool/lib/interface"
//按钮的位置
export enum positionType {
    tl = 'tl',
    tc = 'tc',
    tr = 'tr',
    bl = 'bl',
    bc = 'bc',
    br = 'br',
}
//子菜单弹出的位置
export enum popDir {
    l = 'left',
    t = 'top',
    b = 'bottom',
    r = 'right'
}

export interface actionsItem {
    icon: string,
    label?: string,
    iconSize?: number,
    fontSize?: number,
    color?: string,//背景
    fontColor?: string,
    linear?: string,
    linearDeep?: string
	openType?:string,
	callback?:Function,
    [key: string]: any
}