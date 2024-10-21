import { Shape } from "../shape";
import { tmCv } from "..";
import { shapeStyle } from "../interface";

export class tmImages extends Shape {

	src: string = "";
	isLoad: boolean = false;
	isErr: boolean = false;
	imgWidth: number = 0
	imgHeight: number = 0
	constructor(tmcv: tmCv, arg: Partial<shapeStyle>, src: string) {
		super(tmcv, arg)
		this.radius = arg?.radius ?? this.radius;
		this.src = src;
		this.donwloadImg()
	}
	donwloadImg() {
		let t = this;
		function err() {
			t.isLoad = true;
			t.isErr = true;

		}
		// #ifdef H5
		const img = new Image()
		img.src = this.src;
		img.onload = function (e) {
			t.imgWidth = img.width;
			t.imgHeight = img.height;
			t.isLoad = true;
			t.isErr = false;

			if (t.canvas.graphs.findIndex(el => el.id == t.id) > -1) {
				t.draw()
				if (t.canvas.ctx.draw) {
					t.canvas.draw()
				}
			}
		}
		img.onerror = err;
		img.onabort = err;
		// #endif
		// #ifdef MP-WEIXIN
		const img = this.canvas.render?.createImage()

		img.onload = function (e) {
			t.imgWidth = e.path[0].width;
			t.imgHeight = e.path[0].height;
			t.isLoad = true;
			t.isErr = false;
			t.src = img;
			if (t.canvas.graphs.findIndex(el => el.id == t.id) > -1) {
				t.draw()
				if (t.canvas.ctx.draw) {
					t.canvas.draw()
				}
			}

		}
		img.onerror = err;
		img.onabort = err;
		img.src = t.src;
		// #endif

		// #ifndef H5 || MP-WEIXIN

		uni.downloadFile({
			url: this.src,
			fail(er) {

				err()
			},
			success(tep) {
				if (tep.statusCode !== 200) return;
				t.src = tep.tempFilePath
				uni.getImageInfo({ src: tep.tempFilePath }).then(r => {
					t.imgWidth = r.width;
					t.imgHeight = r.height;
					t.isLoad = true;
					t.isErr = false;
					if (t.canvas.graphs.findIndex(el => el.id == t.id) > -1) {
						t.draw()
						if (t.canvas.ctx.draw) {
							t.canvas.draw()
						}
					}
				}).catch(() => err)
			}
		})
		// #endif
	}
	draw(): this {
		if (!this.canvas.ctx || !this.isLoad) return this;

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

		const maxWidth = this.textMaxWidth || this.canvas.opts.width
		const text = this.text
		const letterSpace = this.letterSpace;
		const fontSize = this.fontSize;
		const lineHeight = fontSize * 1.5
		ctx.save()
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
		ctx.clip()
		if (!this.isErr) {
			ctx.drawImage(this.src, 0, 0, this.imgWidth, this.imgHeight, x, y, width, height)
		}

		ctx.fill()
		ctx.stroke()
		ctx.restore();
		return this;
	}
}