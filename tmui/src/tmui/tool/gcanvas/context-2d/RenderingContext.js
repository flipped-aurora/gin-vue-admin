import FillStylePattern from './FillStylePattern';
import FillStyleLinearGradient from './FillStyleLinearGradient';
import FillStyleRadialGradient from './FillStyleRadialGradient';
import GImage from '../env/image.js';
import {
	ArrayBufferToBase64,
	Base64ToUint8ClampedArray
} from '../env/tool.js';

export default class CanvasRenderingContext2D {

	_drawCommands = '';

	_globalAlpha = 1.0;

	_fillStyle = 'rgb(0,0,0)';
	_strokeStyle = 'rgb(0,0,0)';

	_lineWidth = 1;
	_lineCap = 'butt';
	_lineJoin = 'miter';

	_miterLimit = 10;

	_globalCompositeOperation = 'source-over';

	_textAlign = 'start';
	_textBaseline = 'alphabetic';

	_font = '10px sans-serif';

	_savedGlobalAlpha = [];

	timer = null;
	componentId = null;

	_notCommitDrawImageCache = [];
	_needRedrawImageCache = [];
	_redrawCommands = '';
	_autoSaveContext = true;
	// _imageMap = new GHashMap();
	// _textureMap = new GHashMap();

	constructor() {
		this.className = 'CanvasRenderingContext2D';
		//this.save()
	}

	setFillStyle(value) {
		this.fillStyle = value;
	}

	set fillStyle(value) {
		this._fillStyle = value;

		if (typeof(value) == 'string') {
			this._drawCommands = this._drawCommands.concat("F" + value + ";");
		} else if (value instanceof FillStylePattern) {
			const image = value._img;
			if (!image.complete) {
				image.onload = () => {
					var index = this._needRedrawImageCache.indexOf(image);
					if (index > -1) {
						this._needRedrawImageCache.splice(index, 1);
						CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
						this._redrawflush(true);
					}
				}
				this._notCommitDrawImageCache.push(image);
			} else {
				CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
			}

			//CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
			this._drawCommands = this._drawCommands.concat("G" + image._id + "," + value._style + ";");
		} else if (value instanceof FillStyleLinearGradient) {
			var command = "D" + value._start_pos._x.toFixed(2) + "," + value._start_pos._y.toFixed(2) + "," +
				value._end_pos._x.toFixed(2) + "," + value._end_pos._y.toFixed(2) + "," +
				value._stop_count;
			for (var i = 0; i < value._stop_count; ++i) {
				command += ("," + value._stops[i]._pos + "," + value._stops[i]._color);
			}
			this._drawCommands = this._drawCommands.concat(command + ";");
		} else if (value instanceof FillStyleRadialGradient) {
			var command = "H" + value._start_pos._x.toFixed(2) + "," + value._start_pos._y.toFixed(2) + "," + value._start_pos._r
				.toFixed(2) + "," +
				value._end_pos._x.toFixed(2) + "," + value._end_pos._y.toFixed(2) + "," + value._end_pos._r.toFixed(2) + "," +
				value._stop_count;
			for (var i = 0; i < value._stop_count; ++i) {
				command += ("," + value._stops[i]._pos + "," + value._stops[i]._color);
			}
			this._drawCommands = this._drawCommands.concat(command + ";");
		}
	}

	get fillStyle() {
		return this._fillStyle;
	}

	get globalAlpha() {
		return this._globalAlpha;
	}

	setGlobalAlpha(value) {
		this.globalAlpha = value;
	}

	set globalAlpha(value) {
		this._globalAlpha = value;
		this._drawCommands = this._drawCommands.concat("a" + value.toFixed(2) + ";");
	}


	get strokeStyle() {
		return this._strokeStyle;
	}

	setStrokeStyle(value) {
		this.strokeStyle = value;
	}

	set strokeStyle(value) {

		this._strokeStyle = value;

		if (typeof(value) == 'string') {
			this._drawCommands = this._drawCommands.concat("S" + value + ";");
		} else if (value instanceof FillStylePattern) {
			CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
			this._drawCommands = this._drawCommands.concat("G" + image._id + "," + value._style + ";");
		} else if (value instanceof FillStyleLinearGradient) {
			var command = "D" + value._start_pos._x.toFixed(2) + "," + value._start_pos._y.toFixed(2) + "," +
				value._end_pos._x.toFixed(2) + "," + value._end_pos._y.toFixed(2) + "," +
				value._stop_count;

			for (var i = 0; i < value._stop_count; ++i) {
				command += ("," + value._stops[i]._pos + "," + value._stops[i]._color);
			}
			this._drawCommands = this._drawCommands.concat(command + ";");
		} else if (value instanceof FillStyleRadialGradient) {
			var command = "H" + value._start_pos._x.toFixed(2) + "," + value._start_pos._y.toFixed(2) + "," + value._start_pos._r
				.toFixed(2) + "," +
				value._end_pos._x.toFixed(2) + "," + value._end_pos._y + ",".toFixed(2) + value._end_pos._r.toFixed(2) + "," +
				value._stop_count;

			for (var i = 0; i < value._stop_count; ++i) {
				command += ("," + value._stops[i]._pos + "," + value._stops[i]._color);
			}
			this._drawCommands = this._drawCommands.concat(command + ";");
		}
	}

	get lineWidth() {
		return this._lineWidth;
	}

	setLineWidth(value) {
		this.lineWidth = value;
	}

	set lineWidth(value) {
		this._lineWidth = value;
		this._drawCommands = this._drawCommands.concat("W" + value + ";");
	}

	get lineCap() {
		return this._lineCap;
	}

	setLineCap(value) {
		this.lineCap = value;
	}

	set lineCap(value) {
		this._lineCap = value;
		this._drawCommands = this._drawCommands.concat("C" + value + ";");
	}

	get lineJoin() {
		return this._lineJoin;
	}

	setLineJoin(value) {
		this.lineJoin = value
	}

	set lineJoin(value) {
		this._lineJoin = value;
		this._drawCommands = this._drawCommands.concat("J" + value + ";");
	}

	get miterLimit() {
		return this._miterLimit;
	}

	setMiterLimit(value) {
		this.miterLimit = value
	}

	set miterLimit(value) {
		this._miterLimit = value;
		this._drawCommands = this._drawCommands.concat("M" + value + ";");
	}

	get globalCompositeOperation() {
		return this._globalCompositeOperation;
	}

	set globalCompositeOperation(value) {

		this._globalCompositeOperation = value;
		let mode = 0;
		switch (value) {
			case "source-over":
				mode = 0;
				break;
			case "source-atop":
				mode = 5;
				break;
			case "source-in":
				mode = 0;
				break;
			case "source-out":
				mode = 2;
				break;
			case "destination-over":
				mode = 4;
				break;
			case "destination-atop":
				mode = 4;
				break;
			case "destination-in":
				mode = 4;
				break;
			case "destination-out":
				mode = 3;
				break;
			case "lighter":
				mode = 1;
				break;
			case "copy":
				mode = 2;
				break;
			case "xor":
				mode = 6;
				break;
			default:
				mode = 0;
		}

		this._drawCommands = this._drawCommands.concat("B" + mode + ";");
	}

	get textAlign() {
		return this._textAlign;
	}

	setTextAlign(value) {
		this.textAlign = value
	}

	set textAlign(value) {

		this._textAlign = value;
		let Align = 0;
		switch (value) {
			case "start":
				Align = 0;
				break;
			case "end":
				Align = 1;
				break;
			case "left":
				Align = 2;
				break;
			case "center":
				Align = 3;
				break;
			case "right":
				Align = 4;
				break;
			default:
				Align = 0;
		}

		this._drawCommands = this._drawCommands.concat("A" + Align + ";");
	}

	get textBaseline() {
		return this._textBaseline;
	}

	setTextBaseline(value) {
		this.textBaseline = value
	}

	set textBaseline(value) {
		this._textBaseline = value;
		let baseline = 0;
		switch (value) {
			case "alphabetic":
				baseline = 0;
				break;
			case "middle":
				baseline = 1;
				break;
			case "top":
				baseline = 2;
				break;
			case "hanging":
				baseline = 3;
				break;
			case "bottom":
				baseline = 4;
				break;
			case "ideographic":
				baseline = 5;
				break;
			default:
				baseline = 0;
				break;
		}

		this._drawCommands = this._drawCommands.concat("E" + baseline + ";");
	}

	get font() {
		return this._font;
	}

	setFontSize(size) {
		var str = this._font;
		var strs = str.trim().split(/\s+/);
		for (var i = 0; i < strs.length; i++) {
			var values = ["normal", "italic", "oblique", "normal", "small-caps", "normal", "bold",
				"bolder", "lighter", "100", "200", "300", "400", "500", "600", "700", "800", "900",
				"normal", "ultra-condensed", "extra-condensed", "condensed", "semi-condensed",
				"semi-expanded", "expanded", "extra-expanded", "ultra-expanded"
			];

			if (-1 == values.indexOf(strs[i].trim())) {
				if (typeof size === 'string') {
					strs[i] = size;
				} else if (typeof size === 'number') {
					strs[i] = String(size) + 'px';
				}
				break;
			}
		}
		this.font = strs.join(" ");
	}

	set font(value) {
		this._font = value;
		this._drawCommands = this._drawCommands.concat("j" + value + ";");
	}

	setTransform(a, b, c, d, tx, ty) {
		this._drawCommands = this._drawCommands.concat("t" +
			(a === 1 ? "1" : a.toFixed(2)) + "," +
			(b === 0 ? "0" : b.toFixed(2)) + "," +
			(c === 0 ? "0" : c.toFixed(2)) + "," +
			(d === 1 ? "1" : d.toFixed(2)) + "," + tx.toFixed(2) + "," + ty.toFixed(2) + ";");
	}

	transform(a, b, c, d, tx, ty) {
		this._drawCommands = this._drawCommands.concat("f" +
			(a === 1 ? "1" : a.toFixed(2)) + "," +
			(b === 0 ? "0" : b.toFixed(2)) + "," +
			(c === 0 ? "0" : c.toFixed(2)) + "," +
			(d === 1 ? "1" : d.toFixed(2)) + "," + tx + "," + ty + ";");
	}

	resetTransform() {
		this._drawCommands = this._drawCommands.concat("m;");
	}

	scale(a, d) {
		this._drawCommands = this._drawCommands.concat("k" + a.toFixed(2) + "," +
			d.toFixed(2) + ";");
	}

	rotate(angle) {
		this._drawCommands = this._drawCommands
			.concat("r" + angle.toFixed(6) + ";");
	}

	translate(tx, ty) {
		this._drawCommands = this._drawCommands.concat("l" + tx.toFixed(2) + "," + ty.toFixed(2) + ";");
	}

	save() {
		this._savedGlobalAlpha.push(this._globalAlpha);
		this._drawCommands = this._drawCommands.concat("v;");
	}

	restore() {
		this._drawCommands = this._drawCommands.concat("e;");
		this._globalAlpha = this._savedGlobalAlpha.pop();
	}

	createPattern(img, pattern) {
		if (typeof img === 'string') {
			var imgObj = new GImage();
			imgObj.src = img;
			img = imgObj;
		}
		return new FillStylePattern(img, pattern);
	}

	createLinearGradient(x0, y0, x1, y1) {
		return new FillStyleLinearGradient(x0, y0, x1, y1);
	}

	createRadialGradient = function(x0, y0, r0, x1, y1, r1) {
		return new FillStyleRadialGradient(x0, y0, r0, x1, y1, r1);
	};

	createCircularGradient = function(x0, y0, r0) {
		return new FillStyleRadialGradient(x0, y0, 0, x0, y0, r0);
	};

	strokeRect(x, y, w, h) {
		this._drawCommands = this._drawCommands.concat("s" + x + "," + y + "," + w + "," + h + ";");
	}


	clearRect(x, y, w, h) {
		this._drawCommands = this._drawCommands.concat("c" + x + "," + y + "," + w +
			"," + h + ";");
	}

	clip() {
		this._drawCommands = this._drawCommands.concat("p;");
	}

	resetClip() {
		this._drawCommands = this._drawCommands.concat("q;");
	}

	closePath() {
		this._drawCommands = this._drawCommands.concat("o;");
	}

	moveTo(x, y) {
		this._drawCommands = this._drawCommands.concat("g" + x.toFixed(2) + "," + y.toFixed(2) + ";");
	}

	lineTo(x, y) {
		this._drawCommands = this._drawCommands.concat("i" + x.toFixed(2) + "," + y.toFixed(2) + ";");
	}

	quadraticCurveTo = function(cpx, cpy, x, y) {
		this._drawCommands = this._drawCommands.concat("u" + cpx + "," + cpy + "," + x + "," + y + ";");
	}

	bezierCurveTo(cp1x, cp1y, cp2x, cp2y, x, y, ) {
		this._drawCommands = this._drawCommands.concat(
			"z" + cp1x.toFixed(2) + "," + cp1y.toFixed(2) + "," + cp2x.toFixed(2) + "," + cp2y.toFixed(2) + "," +
			x.toFixed(2) + "," + y.toFixed(2) + ";");
	}

	arcTo(x1, y1, x2, y2, radius) {
		this._drawCommands = this._drawCommands.concat("h" + x1 + "," + y1 + "," + x2 + "," + y2 + "," + radius + ";");
	}

	beginPath() {
		this._drawCommands = this._drawCommands.concat("b;");
	}


	fillRect(x, y, w, h) {
		this._drawCommands = this._drawCommands.concat("n" + x + "," + y + "," + w +
			"," + h + ";");
	}

	rect(x, y, w, h) {
		this._drawCommands = this._drawCommands.concat("w" + x + "," + y + "," + w + "," + h + ";");
	}

	fill() {
		this._drawCommands = this._drawCommands.concat("L;");
	}

	stroke(path) {
		this._drawCommands = this._drawCommands.concat("x;");
	}

	arc(x, y, radius, startAngle, endAngle, anticlockwise) {

		let ianticlockwise = 0;
		if (anticlockwise) {
			ianticlockwise = 1;
		}

		this._drawCommands = this._drawCommands.concat(
			"y" + x.toFixed(2) + "," + y.toFixed(2) + "," +
			radius.toFixed(2) + "," + startAngle + "," + endAngle + "," + ianticlockwise +
			";"
		);
	}

	fillText(text, x, y) {
		let tmptext = text.replace(/!/g, "!!");
		tmptext = tmptext.replace(/,/g, "!,");
		tmptext = tmptext.replace(/;/g, "!;");
		this._drawCommands = this._drawCommands.concat("T" + tmptext + "," + x + "," + y + ",0.0;");
	}

	strokeText = function(text, x, y) {
		let tmptext = text.replace(/!/g, "!!");
		tmptext = tmptext.replace(/,/g, "!,");
		tmptext = tmptext.replace(/;/g, "!;");
		this._drawCommands = this._drawCommands.concat("U" + tmptext + "," + x + "," + y + ",0.0;");
	}

	measureText(text) {
		return CanvasRenderingContext2D.GBridge.measureText(text, this.font, this.componentId);
	}

	isPointInPath = function(x, y) {
		throw new Error('GCanvas not supported yet');
	}

	drawImage(image, sx, sy, sw, sh, dx, dy, dw, dh) {
		if (typeof image === 'string') {
			var imgObj = new GImage();
			imgObj.src = image;
			image = imgObj;
		}
		console.log(imgObj)
		if (image instanceof GImage) {
			if (!image.complete) {
				imgObj.onload = () => {
					
					var index = this._needRedrawImageCache.indexOf(image);
					if (index > -1) {
						this._needRedrawImageCache.splice(index, 1);
						CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
						this._redrawflush(true);
					}
				}
				this._notCommitDrawImageCache.push(image);
			} else {
				
				CanvasRenderingContext2D.GBridge.bindImageTexture(this.componentId, image.src, image._id);
			}
			var srcArgs = [image, sx, sy, sw, sh, dx, dy, dw, dh];
			var args = [];
			for (var arg in srcArgs) {
				if (typeof(srcArgs[arg]) != 'undefined') {
					args.push(srcArgs[arg]);
				}
			}
			this.__drawImage.apply(this, args);
			//this.__drawImage(image,sx, sy, sw, sh, dx, dy, dw, dh);
		}
	}

	__drawImage(image, sx, sy, sw, sh, dx, dy, dw, dh) {
		const numArgs = arguments.length;

		function drawImageCommands() {

			if (numArgs === 3) {
				const x = parseFloat(sx) || 0.0;
				const y = parseFloat(sy) || 0.0;

				return ("d" + image._id + ",0,0," +
					image.width + "," + image.height + "," +
					x + "," + y + "," + image.width + "," + image.height + ";");
			} else if (numArgs === 5) {
				const x = parseFloat(sx) || 0.0;
				const y = parseFloat(sy) || 0.0;
				const width = parseInt(sw) || image.width;
				const height = parseInt(sh) || image.height;

				return ("d" + image._id + ",0,0," +
					image.width + "," + image.height + "," +
					x + "," + y + "," + width + "," + height + ";");
			} else if (numArgs === 9) {
				sx = parseFloat(sx) || 0.0;
				sy = parseFloat(sy) || 0.0;
				sw = parseInt(sw) || image.width;
				sh = parseInt(sh) || image.height;
				dx = parseFloat(dx) || 0.0;
				dy = parseFloat(dy) || 0.0;
				dw = parseInt(dw) || image.width;
				dh = parseInt(dh) || image.height;

				return ("d" + image._id + "," +
					sx + "," + sy + "," + sw + "," + sh + "," +
					dx + "," + dy + "," + dw + "," + dh + ";");
			}
		}
		this._drawCommands += drawImageCommands();
	}

	_flush(reserve, callback) {
		const commands = this._drawCommands;
		this._drawCommands = '';
		CanvasRenderingContext2D.GBridge.render2d(this.componentId, commands, callback);
		this._needRender = false;
	}

	_redrawflush(reserve, callback) {
		const commands = this._redrawCommands;
		CanvasRenderingContext2D.GBridge.render2d(this.componentId, commands, callback);
		if (this._needRedrawImageCache.length == 0) {
			this._redrawCommands = '';
		}
	}

	draw(reserve, callback) {
		if (!reserve) {
			this._globalAlpha = this._savedGlobalAlpha.pop();
			this._savedGlobalAlpha.push(this._globalAlpha);
			this._redrawCommands = this._drawCommands;
			this._needRedrawImageCache = this._notCommitDrawImageCache;
			if (this._autoSaveContext) {
				this._drawCommands = ("v;" + this._drawCommands);
				this._autoSaveContext = false;
			} else {
				this._drawCommands = ("e;X;v;" + this._drawCommands);
			}
		} else {
			this._needRedrawImageCache = this._needRedrawImageCache.concat(this._notCommitDrawImageCache);
			this._redrawCommands += this._drawCommands;
			if (this._autoSaveContext) {
				this._drawCommands = ("v;" + this._drawCommands);
				this._autoSaveContext = false;
			}
		}
		this._notCommitDrawImageCache = [];
		if (this._flush) {
			this._flush(reserve, callback);
		}
	}

	getImageData(x, y, w, h, callback) {
		CanvasRenderingContext2D.GBridge.getImageData(this.componentId, x, y, w, h, function(res) {
			res.data = Base64ToUint8ClampedArray(res.data);
			if (typeof(callback) == 'function') {
				callback(res);
			}
		});
	}

	putImageData(data, x, y, w, h, callback) {
		if (data instanceof Uint8ClampedArray) {
			data = ArrayBufferToBase64(data);
			CanvasRenderingContext2D.GBridge.putImageData(this.componentId, data, x, y, w, h, function(res) {
				if (typeof(callback) == 'function') {
					callback(res);
				}
			});
		}
	}

	toTempFilePath(x, y, width, height, destWidth, destHeight, fileType, quality, callback) {
		CanvasRenderingContext2D.GBridge.toTempFilePath(this.componentId, x, y, width, height, destWidth, destHeight,
			fileType, quality,
			function(res) {
				if (typeof(callback) == 'function') {
					callback(res);
				}
			});
	}
}
