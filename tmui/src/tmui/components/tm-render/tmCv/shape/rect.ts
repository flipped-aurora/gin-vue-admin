import { Shape } from "../shape";

export class tmRect extends Shape {
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
        if (ctx.setFillStyle) {
            ctx.setFillStyle(fillStyle)
            ctx.setLineWidth(lineWidth)
            ctx.setStrokeStyle(strokeStyle)
        } else {
            ctx.fillStyle = fillStyle;
            ctx.lineWidth = lineWidth;
            ctx.strokeStyle = strokeStyle;
        }

        ctx.fillRect(x, y, width, height)
        ctx.strokeRect(x, y, width, height)
        return this;
    }
}