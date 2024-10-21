import { tmCv } from "..";
import { shapeStyle } from "../interface";
import { Shape } from "../shape";
interface roundRect extends shapeStyle {
    radius: [number, number, number, number]
}
export class tmRoundRect extends Shape {
    radius: [number, number, number, number] = [0, 0, 0, 0]
    constructor(tmcv: tmCv, arg: Partial<roundRect>) {
        super(tmcv, arg)
        this.radius = arg?.radius ?? this.radius;
    }
    draw(): this {
        if (!this.canvas.ctx) return this;
        let ctx = this.canvas.ctx;
        const x = this.x;
        const y = this.y;
        const width = this.width;
        const height = this.height;
        const fillStyle = this.fillStyle;
        const strokeStyle = this.strokeStyle;
        const lineWidth = this.lineWidth
        const radius: [number, number, number, number] = this.radius
        const topLeftRadius = radius[0];
        const topRightRadius = radius[1];
        const bottomLeftRadius = radius[3];
        const bottomRightRadius = radius[2];
        ctx.beginPath();
        ctx.moveTo(x + topLeftRadius, y);
        ctx.lineTo(x + width - topRightRadius, y);
        ctx.quadraticCurveTo(x + width, y, x + width, y + topRightRadius);
        ctx.lineTo(x + width, y + height - bottomRightRadius);
        ctx.quadraticCurveTo(x + width, y + height, x + width - bottomRightRadius, y + height);
        ctx.lineTo(x + bottomLeftRadius, y + height);
        ctx.quadraticCurveTo(x, y + height, x, y + height - bottomLeftRadius);
        ctx.lineTo(x, y + topLeftRadius);
        ctx.quadraticCurveTo(x, y, x + topLeftRadius, y);
        ctx.closePath();

        if (ctx.setFillStyle) {
            ctx.setFillStyle(fillStyle)
            ctx.setLineWidth(lineWidth)
            ctx.setStrokeStyle(strokeStyle)
        } else {
            ctx.fillStyle = fillStyle;
            ctx.lineWidth = lineWidth;
            ctx.strokeStyle = strokeStyle;
        }

        ctx.fill()
        ctx.stroke()
        return this;
    }
}