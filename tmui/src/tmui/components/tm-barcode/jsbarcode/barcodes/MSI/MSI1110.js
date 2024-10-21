import MSI from './MSI.js';
import { mod10, mod11 } from './checksums.js';

class MSI1110 extends MSI {
	constructor(data, options) {
		data += mod11(data);
		data += mod10(data);
		super(data, options);
	}
}

export default MSI1110;
