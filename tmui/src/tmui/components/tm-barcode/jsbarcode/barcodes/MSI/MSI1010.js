import MSI from './MSI.js';
import { mod10 } from './checksums.js';

class MSI1010 extends MSI {
	constructor(data, options) {
		data += mod10(data);
		data += mod10(data);
		super(data, options);
	}
}

export default MSI1010;
