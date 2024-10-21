import Barcode from "../Barcode.js";
import { SHIFT, SET_A, SET_B, MODULO, STOP, FNC1, SET_BY_CODE, SWAP, BARS } from './constants';

// This is the master class,
// it does require the start code to be included in the string
class CODE128 extends Barcode {
	constructor(data, options) {
		super(data.substring(1), options);

		// Get array of ascii codes from data
		this.bytes = data.split('')
			.map(char => char.charCodeAt(0));
	}

	valid() {
		// ASCII value ranges 0-127, 200-211
		return /^[\x00-\x7F\xC8-\xD3]+$/.test(this.data);
	}

	// The public encoding function
	encode() {
		const bytes = this.bytes;
		// Remove the start code from the bytes and set its index
		const startIndex = bytes.shift() - 105;
		// Get start set by index
		const startSet = SET_BY_CODE[startIndex];

		if (startSet === undefined) {
			throw new RangeError('The encoding does not start with a start character.');
		}

		if (this.shouldEncodeAsEan128() === true) {
			bytes.unshift(FNC1);
		}

		// Start encode with the right type
		const encodingResult = CODE128.next(bytes, 1, startSet);

		return {
			text:
				this.text === this.data
					? this.text.replace(/[^\x20-\x7E]/g, '')
					: this.text,
			data:
				// Add the start bits
				CODE128.getBar(startIndex) +
				// Add the encoded bits
				encodingResult.result +
				// Add the checksum
				CODE128.getBar((encodingResult.checksum + startIndex) % MODULO) +
				// Add the end bits
				CODE128.getBar(STOP)
		};
	}

	// GS1-128/EAN-128
	shouldEncodeAsEan128() {
		let isEAN128 = this.options.ean128 || false;
		if (typeof isEAN128 === 'string') {
			isEAN128 = isEAN128.toLowerCase() === 'true';
		}
		return isEAN128;
	}

	// Get a bar symbol by index
	static getBar(index) {
		return BARS[index] ? BARS[index].toString() : '';
	}

	// Correct an index by a set and shift it from the bytes array
	static correctIndex(bytes, set) {
		if (set === SET_A) {
			const charCode = bytes.shift();
			return charCode < 32 ? charCode + 64 : charCode - 32;
		} else if (set === SET_B) {
			return bytes.shift() - 32;
		} else {
			return (bytes.shift() - 48) * 10 + bytes.shift() - 48;
		}
	}

	static next(bytes, pos, set) {
		if (!bytes.length) {
			return { result: '', checksum: 0 };
		}

		let nextCode, index;

		// Special characters
		if (bytes[0] >= 200) {
			index = bytes.shift() - 105;
			const nextSet = SWAP[index];

			// Swap to other set
			if (nextSet !== undefined) {
				nextCode = CODE128.next(bytes, pos + 1, nextSet);
			}
			// Continue on current set but encode a special character
			else {
				// Shift
				if ((set === SET_A || set === SET_B) && index === SHIFT) {
					// Convert the next character so that is encoded correctly
					bytes[0] = (set === SET_A)
						? bytes[0] > 95 ? bytes[0] - 96 : bytes[0]
						: bytes[0] < 32 ? bytes[0] + 96 : bytes[0];
				}
				nextCode = CODE128.next(bytes, pos + 1, set);
			}
		}
		// Continue encoding
		else {
			index = CODE128.correctIndex(bytes, set);
			nextCode = CODE128.next(bytes, pos + 1, set);
		}

		// Get the correct binary encoding and calculate the weight
		const enc = CODE128.getBar(index);
		const weight = index * pos;

		return {
			result: enc + nextCode.result,
			checksum: weight + nextCode.checksum
		};
	}
}

export default CODE128;
