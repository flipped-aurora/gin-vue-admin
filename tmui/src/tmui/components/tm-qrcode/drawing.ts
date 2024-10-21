/**
 * canvas-qr
 * create qrcode steam based on node-canvas
 * core code is from https://github.com/neocotic/qr.js
 */

//globs
import { generateFrame } from "./qrcode"
import { qrOpts, qrOptsDefault } from "./interface";
var Canvas = null
	, Image = null;


/**
 * @param size {Number} canvassize
 * @param lw {Number} logo width
 * @param lh {Number} logo height
 * return {Object}
 */
function computeLogoPos(size, lw, lh) {
	return {
		x: (size - lw) / 2
		, y: (size - lh) / 2
		, w: lw
		, h: lh
	}
}

/**
 * @param option {Object}
 * return {WriteSteam}
 */
export async function qr(ctx, option, canvas2d) {

	if (!ctx) return;


	var defaults = Object.assign({
		...qrOptsDefault
	}, option)
		, size = defaults.size
		, borderWidth = size * defaults.border
		, qrSize = size - borderWidth * 2
		, backgroundColor = defaults.backgroundColor
		, backgroundImage = defaults.backgroundImage
		, logoImage = defaults.logoImage
		, i
		, j
		, points
		, width
		, fo
		, px
		, logoPos
		, linearDir = defaults.linearDir

	// var cvs = ctx
	var c2d = ctx
	if (!c2d?.width) {
		c2d.width = size
	}
	if (!c2d?.height) {
		c2d.height = size
	}
	//base background
	fillStyle(c2d, defaults.baseColor, linearDir)
	c2d.fillRect(0, 0, size, size)
	c2d.save()
	//draw background color
	if (backgroundColor) {
		fillStyle(c2d, backgroundColor, linearDir)
		c2d.fillRect(0, 0, size, size)
	}
	c2d.restore()
	//draw backgroundImage
	if (backgroundImage) {

		if (canvas2d) {
			await drawImage(canvas2d, c2d, { width: size, height: size, src: backgroundImage, x: 0, y: 0 })
		} else {
			c2d.drawImage(backgroundImage, 0, 0, size, size)
		}

	}

	//now draw qrcode
	fo = generateFrame(defaults.str, defaults.ecc, qrSize)
	points = fo.frameBuffer
	width = fo.width
	px = qrSize / width

	fillStyle(c2d, defaults.forgroundColor, linearDir, linearDir)
	for (i = 0; i < width; i++) {
		for (j = 0; j < width; j++) {
			if (points[j * width + i]) {
				c2d.fillRect(borderWidth + px * i, borderWidth + px * j, px, px)
			}
		}
	}

	//draw logo image
	if (logoImage) {
		logoPos = computeLogoPos(size, defaults.logoWidth, defaults.logoHeight)

		if (canvas2d) {
			await drawImage(canvas2d, c2d, { width: defaults.logoWidth, height: defaults.logoHeight, src: logoImage, x: logoPos.x, y: logoPos.y })
		} else {
			c2d.drawImage(logoImage, logoPos.x, logoPos.y, defaults.logoWidth, defaults.logoHeight)
		}

	}


	//非2d类需要draw。2d无需draw
	if (!canvas2d) {
		c2d.draw()
	}
	uni.hideLoading()
	return ctx

}

function fillStyle(ctx, value, linearDir = "left") {
	//绘制渐变
	if (typeof value == 'object' && Array.isArray(value)) {
		let w2w = parseInt(String(ctx.width / 2))
		
		var gradient = ctx.createLinearGradient(w2w, 0, w2w, ctx.width);
		if (linearDir == "left") {
			gradient = ctx.createLinearGradient(ctx.width, w2w, 0, w2w);
		} else if (linearDir == "bottom") {
			gradient = ctx.createLinearGradient(w2w, 0, w2w, ctx.width);
		} else if (linearDir == "top") {
			gradient = ctx.createLinearGradient(w2w, ctx.width, w2w, 0);
		} else if (linearDir == "right") {
			gradient = ctx.createLinearGradient(0, w2w, ctx.width, w2w);
		} else if (linearDir == "tlbr") {
			gradient = ctx.createLinearGradient(0, 0, ctx.width, ctx.width);
		} else if (linearDir == "trbl") {
			gradient = ctx.createLinearGradient(ctx.width, 0, 0, ctx.width);
		} else if (linearDir == "bltr") {
			gradient = ctx.createLinearGradient(0, ctx.width, ctx.width, 0);
		} else if (linearDir == "brtl") {
			gradient = ctx.createLinearGradient(ctx.width, ctx.width, 0, 0);
		}
		for (let i = 0, len = value.length; i < len; i++) {
			let stop = i / len;
			if (i == 0) stop = 0;
			if (i == len - 1) stop = 1;
			gradient.addColorStop(stop, value[i]);
		}
		console.log(gradient)
		// gradient.addColorStop(1, value[1]);
		// #ifdef APP-NVUE || MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.strokeStyle = gradient;
		ctx.fillStyle = gradient;
		// #endif
		// #ifndef APP-NVUE || MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.setStrokeStyle(gradient)
		ctx.setFillStyle(gradient)
		// #endif
	} else {
		// #ifdef APP-NVUE || MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.fillStyle = value
		// #endif
		// #ifndef APP-NVUE || MP-WEIXIN || MP-ALIPAY || MP-QQ
		ctx.setFillStyle(value)
		// #endif
	}


}

function drawImage(canvas2d, ctx, opts) {
	uni.showLoading({
		title: "..."
	})
	let img = canvas2d.createImage()
	try{
		img.width = opts.width;
		img.height = opts.height;
	}catch(e){
		//TODO handle the exception
	}
	img.src = opts.src;
	return new Promise(res => {
		img.onload = function () {
			ctx.drawImage(img, opts.x, opts.y, opts.width, opts.height)
			uni.hideLoading()
			res(true)
		}
	})
}
