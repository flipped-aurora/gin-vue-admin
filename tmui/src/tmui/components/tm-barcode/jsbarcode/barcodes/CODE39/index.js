// Encoding documentation:
// https://en.wikipedia.org/wiki/Code_39#Encoding

import Barcode from "../Barcode.js";

class CODE39 extends Barcode {
	constructor(data, options) {
		data = data.toUpperCase();

		// Calculate mod43 checksum if enabled
		if (options.mod43) {
			data += getCharacter(mod43checksum(data));
		}

		super(data, options);
	}

	encode() {
		// First character is always a *
		var result = getEncoding("*");

		// Take every character and add the binary representation to the result
		for (let i = 0; i < this.data.length; i++) {
			result += getEncoding(this.data[i]) + "0";
		}

		// Last character is always a *
		result += getEncoding("*");

		return {
			data: result,
			text: this.text
		};
	}

	valid() {
		return this.data.search(/^[0-9A-Z\-\.\ \$\/\+\%]+$/) !== -1;
	}
}






// All characters. The position in the array is the (checksum) value
var characters = [
	"0", "1", "2", "3",
	"4", "5", "6", "7",
	"8", "9", "A", "B",
	"C", "D", "E", "F",
	"G", "H", "I", "J",
	"K", "L", "M", "N",
	"O", "P", "Q", "R",
	"S", "T", "U", "V",
	"W", "X", "Y", "Z",
	"-", ".", " ", "$",
	"/", "+", "%", "*"
];

// The decimal representation of the characters, is converted to the
// corresponding binary with the getEncoding function
var encodings = [
	20957, 29783, 23639, 30485,
	20951, 29813, 23669, 20855,
	29789, 23645, 29975, 23831,
	30533, 22295, 30149, 24005,
	21623, 29981, 23837, 22301,
	30023, 23879, 30545, 22343,
	30161, 24017, 21959, 30065,
	23921, 22385, 29015, 18263,
	29141, 17879, 29045, 18293,
	17783, 29021, 18269, 17477,
	17489, 17681, 20753, 35770
];

// Get the binary representation of a character by converting the encodings
// from decimal to binary
function getEncoding(character) {
	return getBinary(characterValue(character));
}

function getBinary(characterValue) {
	return encodings[characterValue].toString(2);
}

function getCharacter(characterValue) {
	return characters[characterValue];
}

function characterValue(character) {
	return characters.indexOf(character);
}

function mod43checksum(data) {
	var checksum = 0;
	for (let i = 0; i < data.length; i++) {
		checksum += characterValue(data[i]);
	}

	checksum = checksum % 43;
	return checksum;
}

export { CODE39 };
