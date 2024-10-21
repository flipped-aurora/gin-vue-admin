import CODE128 from './CODE128.js';
import { C_START_CHAR, C_CHARS } from './constants';

class CODE128C extends CODE128 {
	constructor(string, options) {
		super(C_START_CHAR + string, options);
	}

	valid() {
		return (new RegExp(`^${C_CHARS}+$`)).test(this.data);
	}
}

export default CODE128C;
