import { tmCv } from ".";
import { color } from "./color";
import { shapeStyle } from "./interface";
import { division, addition, multiply, subtract, uuid } from "./util";

export class Shape {
    id: string = uuid(16)

    x: number = 0;
    y: number = 0;
    width: number = 0;
    height: number = 0;
    fillStyle: string = 'rgba(0,0,0,0)';
    strokeStyle: string = 'rgba(0,0,0,0)';
    lineWidth: number = 0;
    lineDashOffset: number = 0
    lineDash: number[] = [0]
    sAngle: number = 0
    eAngle: number = 0
    radius: [number, number, number, number] = [0, 0, 0, 0]
    /**一些圆弧等用到的内半径 */
    r: number = 0
    /**圆环时可能用到的内半径 */
    innerRadius: number = 0
    lineCap: any = 'round'
    lineJoin: any = 'round'
    textMaxWidth: number = 0
    /** 字符间距不支持断行 */
    letterSpace: number = 0
    text: string = ''
    playing: boolean = false;
    fontSize: number = 14
    textAlign: 'left' | 'right' | 'center' = 'left'

    canvas: tmCv
    constructor(tmcv: tmCv, arg: Partial<shapeStyle>) {
        for (const item in arg) {
            // @ts-expect-error
            this[item] = arg[item]
            // @ts-expect-error
            this['back_' + item] = arg[item]
        }
        this.canvas = tmcv;
        this.draw()
    }
    /**绘制 */
    draw<T>(arg?: shapeStyle) {
        return this
    }
    /**动画 */
    public animate(arg: Partial<shapeStyle>, animateConfig?: {
        /**动画时长 */
        duration?: number,
        /**重复次数 */
        repeat?: number,
        /**是否来回动画0-1，1-0这样的过渡效果 */
        yoyo?: boolean
    }) {
        const config = {
            duration: 500,
            repeat: 0,
            yoyo: false,
            onStart: () => { },
            ...(animateConfig || {})
        }
        this.playing = true;

        return this.canvas.animation(config, (progress: number) => {
            // 计算动画过程中矩形的位置和颜色
            let { x, y, width, height, radius, eAngle, sAngle, textMaxWidth, letterSpace, fontSize } = this;
            if (typeof arg.x !== 'undefined') {
                const baseX = this['back_x'] ?? 0;
                x = baseX + Number(arg.x - baseX) * progress

            }
            if (typeof arg.y !== 'undefined') {
                const baseY = this['back_y'] ?? 0;
                y = baseY + Number(arg.y - baseY) * progress
            }
            if (typeof arg.width !== 'undefined') {
                const baseWidth = this['back_width'] ?? 0;
                width = baseWidth + Math.abs((arg.width - baseWidth) * (progress))
            }
            if (typeof arg.width !== 'undefined') {
                const baseHeight = this['back_height'] ?? 0;
                height = baseHeight + Math.abs((arg.height - baseHeight) * (progress))
            }
            if (typeof arg.radius !== 'undefined' && radius) {
                this.radius = multiply(arg.radius, progress)
            }
            if (typeof arg.eAngle !== 'undefined' && eAngle) {
                this.eAngle = Number(arg.eAngle) * progress
            }
            if (typeof arg.sAngle !== 'undefined' && sAngle) {
                this.sAngle = Number(arg.sAngle) * progress
            }
            if (typeof arg?.textMaxWidth !== 'undefined') {
                this.textMaxWidth = Number(arg.textMaxWidth) * progress
            }
            if (typeof arg?.letterSpace !== 'undefined' && letterSpace) {
                this.letterSpace = Number(arg.letterSpace) * progress
            }
            if (typeof arg?.fontSize !== 'undefined') {
                const baseFontSize = this['back_fontSize'] ?? 14;
                this.fontSize = baseFontSize + (Number(arg.fontSize) - baseFontSize) * progress

            }

            if (typeof arg.fillStyle !== 'undefined') {
                const [rgbastr, rgba] = color.convertColorToRGBA(arg.fillStyle);
                const [s_rgbas, s_rgba] = color.convertColorToRGBA(this['back_fillStyle']);
                const rgbaResult = subtract(s_rgba, rgba, true)
                rgbaResult[0] = s_rgba[0] + progress * rgbaResult[0]
                rgbaResult[1] = s_rgba[1] + progress * rgbaResult[1]
                rgbaResult[2] = s_rgba[2] + progress * rgbaResult[2]
                rgbaResult[3] = s_rgba[3] + progress * rgbaResult[3]
                this.fillStyle = `rgba(${rgbaResult[0]},${rgbaResult[1]},${rgbaResult[2]},${rgbaResult[3]})`
            }
            if (typeof arg.strokeStyle !== 'undefined') {
                const [rgbastr, rgba] = color.convertColorToRGBA(arg.strokeStyle);
                const [s_rgbas, s_rgba] = color.convertColorToRGBA(this['back_strokeStyle']);
                const rgbaResult = subtract(s_rgba, rgba, true)
                rgbaResult[0] = s_rgba[0] + progress * rgbaResult[0]
                rgbaResult[1] = s_rgba[1] + progress * rgbaResult[1]
                rgbaResult[2] = s_rgba[2] + progress * rgbaResult[2]
                rgbaResult[3] = s_rgba[3] + progress * rgbaResult[3]
                this.strokeStyle = `rgba(${rgbaResult[0]},${rgbaResult[1]},${rgbaResult[2]},${rgbaResult[3]})`
            }
            this.x = x
            this.y = y
            this.width = width
            this.height = height
            this.canvas.draw()

        }).then(() => {
            this.playing = false;
            return Promise.resolve(this)
        })
    }
}