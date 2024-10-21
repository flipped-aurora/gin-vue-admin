import { Shape } from "../shape";

export class tmRing extends Shape {
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
        const innerRadius = this.innerRadius
        if (ctx.setFillStyle) {
            // ctx.setFillStyle(fillStyle)
            ctx.setLineWidth(r - innerRadius)
            ctx.setStrokeStyle(strokeStyle)
            ctx.setLineJoin('round')
            ctx.setLineCap('round')
            ctx.setLineDash(this.lineDash, this.lineDashOffset)
        } else {
            // ctx.fillStyle = fillStyle;
            ctx.lineWidth = r - innerRadius;
            ctx.strokeStyle = strokeStyle;
            ctx.lineJoin = 'round'
            ctx.lineCap = 'round'
            ctx.setLineDash(this.lineDash, this.lineDashOffset)
        }
        ctx.beginPath()
        ctx.arc(x, y, r - (r - innerRadius), sAngle, eAngle)
        // ctx.fill()
        ctx.stroke()
        return this;
    }
}