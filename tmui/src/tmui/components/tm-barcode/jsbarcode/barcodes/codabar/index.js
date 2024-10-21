// Encoding specification:
// http://www.barcodeisland.com/codabar.phtml

import Barcode from "../Barcode.js";

class codabar extends Barcode {
	constructor(data, options) {
		if (data.search(/^[0-9\-\$\:\.\+\/]+$/) === 0) {
			data = "A" + data + "A";
		}

		super(data.toUpperCase(), options);

		this.text = this.options.text || this.text.replace(/[A-D]/g, '');
	}

	valid() {
		return this.data.search(/^[A-D][0-9\-\$\:\.\+\/]+[A-D]$/) !== -1;
	}

	encode() {
		var result = [];
		var encodings = this.getEncodings();
		for (var i = 0; i < this.data.length; i++) {
			result.push(encodings[this.data.charAt(i)]);
			// for all characters except the last, append a narrow-space ("0")
			if (i !== this.data.length - 1) {
				result.push("0");
			}
		}
		return {
			text: this.text,
			data: result.join('')
		};
	}

	getEncodings() {
		return {
			"0": "101010011",
			"1": "101011001",
			"2": "101001011",
			"3": "110010101",
			"4": "101101001",
			"5": "110101001",
			"6": "100101011",
			"7": "100101101",
			"8": "100110101",
			"9": "110100101",
			"-": "101001101",
			"$": "101100101",
			":": "1101011011",
			"/": "1101101011",
			".": "1101101101",
			"+": "1011011011",
			"A": "1011001001",
			"B": "1001001011",
			"C": "1010010011",
			"D": "1010011001"
		};
	}
}

export { codabar };
