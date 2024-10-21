// Encoding documentation:
// https://en.wikipedia.org/wiki/Universal_Product_Code#Encoding

import encode from './encoder';
import Barcode from "../Barcode.js";

class UPC extends Barcode {
	constructor(data, options) {
		// Add checksum if it does not exist
		if (data.search(/^[0-9]{11}$/) !== -1) {
			data += checksum(data);
		}

		super(data, options);

		this.displayValue = options.displayValue;

		// Make sure the font is not bigger than the space between the guard bars
		if (options.fontSize > options.width * 10) {
			this.fontSize = options.width * 10;
		}
		else {
			this.fontSize = options.fontSize;
		}

		// Make the guard bars go down half the way of the text
		this.guardHeight = options.height + this.fontSize / 2 + options.textMargin;
	}

	valid() {
		return this.data.search(/^[0-9]{12}$/) !== -1 &&
			this.data[11] == checksum(this.data);
	}

	encode() {
		if (this.options.flat) {
			return this.flatEncoding();
		}
		else {
			return this.guardedEncoding();
		}
	}

	flatEncoding() {
		var result = "";

		result += "101";
		result += encode(this.data.substr(0, 6), "LLLLLL");
		result += "01010";
		result += encode(this.data.substr(6, 6), "RRRRRR");
		result += "101";

		return {
			data: result,
			text: this.text
		};
	}

	guardedEncoding() {
		var result = [];

		// Add the first digit
		if (this.displayValue) {
			result.push({
				data: "00000000",
				text: this.text.substr(0, 1),
				options: { textAlign: "left", fontSize: this.fontSize }
			});
		}

		// Add the guard bars
		result.push({
			data: "101" + encode(this.data[0], "L"),
			options: { height: this.guardHeight }
		});

		// Add the left side
		result.push({
			data: encode(this.data.substr(1, 5), "LLLLL"),
			text: this.text.substr(1, 5),
			options: { fontSize: this.fontSize }
		});

		// Add the middle bits
		result.push({
			data: "01010",
			options: { height: this.guardHeight }
		});

		// Add the right side
		result.push({
			data: encode(this.data.substr(6, 5), "RRRRR"),
			text: this.text.substr(6, 5),
			options: { fontSize: this.fontSize }
		});

		// Add the end bits
		result.push({
			data: encode(this.data[11], "R") + "101",
			options: { height: this.guardHeight }
		});

		// Add the last digit
		if (this.displayValue) {
			result.push({
				data: "00000000",
				text: this.text.substr(11, 1),
				options: { textAlign: "right", fontSize: this.fontSize }
			});
		}

		return result;
	}
}

// Calulate the checksum digit
// https://en.wikipedia.org/wiki/International_Article_Number_(EAN)#Calculation_of_checksum_digit
export function checksum(number) {
	var result = 0;

	var i;
	for (i = 1; i < 11; i += 2) {
		result += parseInt(number[i]);
	}
	for (i = 0; i < 11; i += 2) {
		result += parseInt(number[i]) * 3;
	}

	return (10 - (result % 10)) % 10;
}

export default UPC;
