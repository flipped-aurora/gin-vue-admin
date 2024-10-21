import merge from "../help/merge.js";
import { calculateEncodingAttributes, getTotalWidthOfEncodings, getMaximumHeightOfEncodings } from "./shared.js";

var svgns = "http://www.w3.org/2000/svg";

class SVGRenderer {
	constructor(svg, encodings, options) {
		this.svg = svg;
		this.encodings = encodings;
		this.options = options;
		this.document = options.xmlDocument || document;
	}

	render() {
		var currentX = this.options.marginLeft;

		this.prepareSVG();
		for (let i = 0; i < this.encodings.length; i++) {
			var encoding = this.encodings[i];
			var encodingOptions = merge(this.options, encoding.options);

			var group = this.createGroup(currentX, encodingOptions.marginTop, this.svg);

			this.setGroupOptions(group, encodingOptions);

			this.drawSvgBarcode(group, encodingOptions, encoding);
			this.drawSVGText(group, encodingOptions, encoding);

			currentX += encoding.width;
		}
	}

	prepareSVG() {
		// Clear the SVG
		while (this.svg.firstChild) {
			this.svg.removeChild(this.svg.firstChild);
		}

		calculateEncodingAttributes(this.encodings, this.options);
		var totalWidth = getTotalWidthOfEncodings(this.encodings);
		var maxHeight = getMaximumHeightOfEncodings(this.encodings);

		var width = totalWidth + this.options.marginLeft + this.options.marginRight;
		this.setSvgAttributes(width, maxHeight);

		if (this.options.background) {
			this.drawRect(0, 0, width, maxHeight, this.svg).setAttribute(
				"style", "fill:" + this.options.background + ";"
			);
		}
	}

	drawSvgBarcode(parent, options, encoding) {
		var binary = encoding.data;

		// Creates the barcode out of the encoded binary
		var yFrom;
		if (options.textPosition == "top") {
			yFrom = options.fontSize + options.textMargin;
		}
		else {
			yFrom = 0;
		}

		var barWidth = 0;
		var x = 0;
		for (var b = 0; b < binary.length; b++) {
			x = b * options.width + encoding.barcodePadding;

			if (binary[b] === "1") {
				barWidth++;
			}
			else if (barWidth > 0) {
				this.drawRect(x - options.width * barWidth, yFrom, options.width * barWidth, options.height, parent);
				barWidth = 0;
			}
		}

		// Last draw is needed since the barcode ends with 1
		if (barWidth > 0) {
			this.drawRect(x - options.width * (barWidth - 1), yFrom, options.width * barWidth, options.height, parent);
		}
	}

	drawSVGText(parent, options, encoding) {
		var textElem = this.document.createElementNS(svgns, 'text');

		// Draw the text if displayValue is set
		if (options.displayValue) {
			var x, y;

			textElem.setAttribute("style",
				"font:" + options.fontOptions + " " + options.fontSize + "px " + options.font
			);

			if (options.textPosition == "top") {
				y = options.fontSize - options.textMargin;
			}
			else {
				y = options.height + options.textMargin + options.fontSize;
			}

			// Draw the text in the correct X depending on the textAlign option
			if (options.textAlign == "left" || encoding.barcodePadding > 0) {
				x = 0;
				textElem.setAttribute("text-anchor", "start");
			}
			else if (options.textAlign == "right") {
				x = encoding.width - 1;
				textElem.setAttribute("text-anchor", "end");
			}
			// In all other cases, center the text
			else {
				x = encoding.width / 2;
				textElem.setAttribute("text-anchor", "middle");
			}

			textElem.setAttribute("x", x);
			textElem.setAttribute("y", y);

			textElem.appendChild(this.document.createTextNode(encoding.text));

			parent.appendChild(textElem);
		}
	}


	setSvgAttributes(width, height) {
		var svg = this.svg;
		svg.setAttribute("width", width + "px");
		svg.setAttribute("height", height + "px");
		svg.setAttribute("x", "0px");
		svg.setAttribute("y", "0px");
		svg.setAttribute("viewBox", "0 0 " + width + " " + height);

		svg.setAttribute("xmlns", svgns);
		svg.setAttribute("version", "1.1");

		svg.setAttribute("style", "transform: translate(0,0)");
	}

	createGroup(x, y, parent) {
		var group = this.document.createElementNS(svgns, 'g');
		group.setAttribute("transform", "translate(" + x + ", " + y + ")");

		parent.appendChild(group);

		return group;
	}

	setGroupOptions(group, options) {
		group.setAttribute("style",
			"fill:" + options.lineColor + ";"
		);
	}

	drawRect(x, y, width, height, parent) {
		var rect = this.document.createElementNS(svgns, 'rect');

		rect.setAttribute("x", x);
		rect.setAttribute("y", y);
		rect.setAttribute("width", width);
		rect.setAttribute("height", height);

		parent.appendChild(rect);

		return rect;
	}
}

export default SVGRenderer;
