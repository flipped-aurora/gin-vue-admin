export interface shapeStyle {
    fillStyle?: string,
    lineWidth?: number,
    strokeStyle?: string,
    lineDashOffset?: number,
    lineDash?: number[],
    sAngle?: number,
    eAngle?: number,
    /**一些圆弧等用到的内半径 */
    r: number,
    /**圆环时可能用到的内半径 */
    innerRadius: number,
    x?: number,
    y?: number,
    width?: number,
    height?: number,
    /** 绘制圆角矩形时可以使用，四个边角的圆角控制。 */
    radius?: [number, number, number, number],
    lineCap?: 'butt' | 'round' | 'square',
    lineJoin?: 'bevel' | 'round' | 'miter',
    lineHeight?: number,
    /** 绘制文字时最大宽度，可以不设置，自动断行。 */
    textMaxWidth?: number,
    /** 字符间距不支持断行 */
    letterSpace?: number,
    /** 要绘制的文本 */
    text?: string | [{
        text: string,
        color?: string,
        fontSize?: number
    }],
    /** 默认14号 */
    fontSize?: number,
    textAlign?: string

    [key: string]: any;
}
export interface shape extends shapeStyle {
    animate<T>(duration: number, arg: T): void
}

export type Rect = 'Rect'
