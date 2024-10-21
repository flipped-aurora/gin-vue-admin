import ITF from './ITF';

// Calculate the checksum digit
const checksum = (data) => {
	const res = data
		.substr(0, 13)
		.split('')
		.map(num => parseInt(num, 10))
		.reduce((sum, n, idx) => sum + (n * (3 - (idx % 2) * 2)), 0);

	return Math.ceil(res / 10) * 10 - res;
};

class ITF14 extends ITF {

	constructor(data, options) {
		// Add checksum if it does not exist
		if (data.search(/^[0-9]{13}$/) !== -1) {
			data += checksum(data);
		}
		super(data, options);
	}

	valid() {
		return (
			this.data.search(/^[0-9]{14}$/) !== -1 &&
			+this.data[13] === checksum(this.data)
		);
	}

}

export default ITF14;
