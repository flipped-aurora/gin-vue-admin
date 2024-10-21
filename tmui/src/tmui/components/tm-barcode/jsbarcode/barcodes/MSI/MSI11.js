import MSI from './MSI.js';
import { mod11 } from './checksums.js';

class MSI11 extends MSI {
	constructor(data, options) {
		super(data + mod11(data), options);
	}
}

export default MSI11;
