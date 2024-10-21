// Encoding documentation:
// http://www.barcodeisland.com/ean8.phtml

import EAN from './EAN';

// Calculate the checksum digit
const checksum = (number) => {
	const res = number
		.substr(0, 7)
		.split('')
		.map((n) => +n)
		.reduce((sum, a, idx) => (
			idx % 2 ? sum + a : sum + a * 3
		), 0);

	return (10 - (res % 10)) % 10;
};

class EAN8 extends EAN {

	constructor(data, options) {
		// Add checksum if it does not exist
		if (data.search(/^[0-9]{7}$/) !== -1) {
			data += checksum(data);
		}

		super(data, options);
	}

	valid() {
		return (
			this.data.search(/^[0-9]{8}$/) !== -1 &&
			+this.data[7] === checksum(this.data)
		);
	}

	leftText() {
		return super.leftText(0, 4);
	}

	leftEncode() {
		const data = this.data.substr(0, 4);
		return super.leftEncode(data, 'LLLL');
	}

	rightText() {
		return super.rightText(4, 4);
	}

	rightEncode() {
		const data = this.data.substr(4, 4);
		return super.rightEncode(data, 'RRRR');
	}

}

export default EAN8;
