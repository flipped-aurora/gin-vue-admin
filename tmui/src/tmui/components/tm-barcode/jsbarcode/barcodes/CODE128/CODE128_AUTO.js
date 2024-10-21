import CODE128 from './CODE128';
import autoSelectModes from './auto';

class CODE128AUTO extends CODE128 {
	constructor(data, options) {
		// ASCII value ranges 0-127, 200-211
		if (/^[\x00-\x7F\xC8-\xD3]+$/.test(data)) {
			super(autoSelectModes(data), options);
		} else {
			super(data, options);
		}
	}
}

export default CODE128AUTO;
