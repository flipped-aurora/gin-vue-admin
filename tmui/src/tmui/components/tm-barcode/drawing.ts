/**
 * 条形码绘制
 * 移植自：https://github.com/lindell/JsBarcode/
 * @auth tmzdy | tmui 
 */
import JsBarcodeOptions from "./jsbarcode/options/defaults.js"
export interface Baroptions {
	width: number,//线条的宽度
	height: number,//线条的高度。
	format: string, //条码类型
	displayValue?: boolean, //是否显示下方文本
	fontOptions?: string,
	font?: string,
	text?: string,
	textAlign?: string,
	textPosition?: string,
	textMargin: number,//文本与条码的间距
	fontSize: number,
	background?: string,
	lineColor?: string,
	margin?: number,//间距
	marginTop: number,
	marginBottom?: number,
	marginLeft?: number,
	marginRight?: number,
	valid?: Function
}
export interface BarcodeObjType {
	text: string,
	data: string,
	options: Baroptions
}
let isAndroid = false
// #ifdef APP-NVUE
isAndroid = uni.getSystemInfoSync().osName == 'android';
// #endif
var ctx: CanvasRenderingContext2D;
function drawCanvasBarcode(options: Baroptions = JsBarcodeOptions, encoding: BarcodeObjType, is2d = false, canvasWidth = 300) {
	var binary = encoding.data;
	var yFrom;
	var height = options.height - 60
	if (options.textPosition == "top") {
		yFrom = options.marginTop + options.fontSize + options.textMargin;
	}
	else {
		yFrom = options.marginTop;
	}
	if (is2d) {
		ctx.fillStyle = options.lineColor;
	} else {
		ctx.setFillStyle(options.lineColor);
	}
	let MAR = (canvasWidth - binary.length * options.width) / 2

	for (var b = 0; b < binary.length; b++) {
		var x = b * options.width + Math.floor(MAR);
		if (binary[b] === "1") {
			ctx.fillRect(x, yFrom, options.width, options.height);
		}
		else if (binary[b]) {
			ctx.fillRect(x, yFrom, options.width, options.height * Number(binary[b]));
		}
	}
	if (!is2d) {
		ctx.draw()
	}

}

function drawCanvasText(options: Baroptions = JsBarcodeOptions, encoding: BarcodeObjType, is2d = false, canvasWidth = 300) {
	ctx.clearRect(0, 0, canvasWidth, options.height + 40)
	var font = ""
	// #ifdef APP-NVUE
	if (isAndroid) {
		font = ((options.fontSize)) + "px " + options.font;
	} else {
		font = ((options.fontSize) * uni.getSystemInfoSync().pixelRatio) + "px " + options.font;
	}
	// #endif
	// #ifndef APP-NVUE
	font = (options.fontSize) + "px " + options.font;
	// #endif
	if (options.displayValue) {
		var x, y;

		if (options.textPosition == "top") {
			y = options.marginTop + options.fontSize - options.textMargin;
		}
		else {
			y = options.height + options.textMargin + options.marginTop + options.fontSize;
		}
		ctx.font = font;
		// let textWidth = ctx.measureText(encoding.text+' ').width||0
		if (options.textAlign == "left") {
			x = 1;
			ctx.textAlign = 'left';
		}
		else if (options.textAlign == "right") {
			x = canvasWidth - 1;
			ctx.textAlign = 'right';
		}
		else {
			x = (canvasWidth) / 2;
			ctx.textAlign = 'center';
		}

		ctx.fillText(encoding.text, x, y);

	}
}

function drawBarCode(context: CanvasRenderingContext2D, options: Baroptions = JsBarcodeOptions, encoding: BarcodeObjType, is2d = false, canvasWidth = 300) {
	ctx = context;
	drawCanvasText(options, encoding, is2d, canvasWidth)
	drawCanvasBarcode(options, encoding, is2d, canvasWidth)
	// #ifdef APP-NVUE
	if (isAndroid) {
		// 最新的sdk3.6.0首次需要绘制两次在安卓上才会显示。
		setTimeout(function () {
			drawCanvasText(options, encoding, is2d, canvasWidth)
			drawCanvasBarcode(options, encoding, is2d, canvasWidth)
		}, 50);
	}
	// #endif
}


export { drawBarCode }





















