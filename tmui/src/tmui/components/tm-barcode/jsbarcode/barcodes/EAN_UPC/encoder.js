import { BINARIES } from './constants';

// Encode data string
const encode = (data, structure, separator) => {
	let encoded = data
		.split('')
		.map((val, idx) => BINARIES[structure[idx]])
		.map((val, idx) => val ? val[data[idx]] : '');

	if (separator) {
		const last = data.length - 1;
		encoded = encoded.map((val, idx) => (
			idx < last ? val + separator : val
		));
	}

	return encoded.join('');
};

export default encode;
