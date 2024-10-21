import CODE128 from './CODE128.js';
import { A_START_CHAR, A_CHARS } from './constants';

class CODE128A extends CODE128 {
	constructor(string, options) {
		super(A_START_CHAR + string, options);
	}

	valid() {
		return (new RegExp(`^${A_CHARS}+$`)).test(this.data);
	}
}

export default CODE128A;
