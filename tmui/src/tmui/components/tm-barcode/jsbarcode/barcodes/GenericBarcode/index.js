import Barcode from "../Barcode.js";

class GenericBarcode extends Barcode {
	constructor(data, options) {
		super(data, options); // Sets this.data and this.text
	}

	// Return the corresponding binary numbers for the data provided
	encode() {
		return {
			data: "10101010101010101010101010101010101010101",
			text: this.text
		};
	}

	// Resturn true/false if the string provided is valid for this encoder
	valid() {
		return true;
	}
}

export { GenericBarcode };
