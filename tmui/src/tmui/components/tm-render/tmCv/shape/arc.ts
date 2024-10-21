import { Shape } from "../shape";

export class tmArc extends Shape {
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
        const r = this.r
        const sAngle = this.sAngle * (Math.PI / 180) - Math.PI / 2
        const eAngle = this.eAngle * (Math.PI / 180) - Math.PI / 2

        ctx.beginPath();
        if (ctx.setFillStyle) {
            ctx.setFillStyle(fillStyle)
            ctx.setLineWidth(lineWidth)
            ctx.setStrokeStyle(strokeStyle)
            ctx.setLineJoin(this.lineJoin)
            ctx.setLineDash(this.lineDash, this.lineDashOffset)
        } else {
            ctx.fillStyle = fillStyle;
            ctx.lineWidth = lineWidth;
            ctx.strokeStyle = strokeStyle;
            ctx.lineJoin = this.lineJoin
            ctx.setLineDash(this.lineDash, this.lineDashOffset)
        }

        ctx.arc(x, y, r, sAngle, eAngle)

        ctx.fill()
        ctx.stroke()

        return this;
    }
}