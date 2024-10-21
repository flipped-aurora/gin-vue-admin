import MSI from './MSI.js';
import { mod10 } from './checksums.js';

class MSI10 extends MSI {
	constructor(data, options) {
		super(data + mod10(data), options);
	}
}

export default MSI10;
