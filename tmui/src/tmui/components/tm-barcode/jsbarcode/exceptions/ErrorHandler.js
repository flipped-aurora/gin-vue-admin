/*eslint no-console: 0 */

class ErrorHandler {
	constructor(api) {
		this.api = api;
	}

	handleCatch(e) {
		// If babel supported extending of Error in a correct way instanceof would be used here
		if (e.name === "InvalidInputException") {
			if (this.api._options.valid !== this.api._defaults.valid) {
				this.api._options.valid(false);
			}
			else {
				throw e.message;
			}
		}
		else {
			throw e;
		}

		this.api.render = function () { };
	}

	wrapBarcodeCall(func) {
		try {
			var result = func(...arguments);
			this.api._options.valid(true);
			return result;
		}
		catch (e) {
			this.handleCatch(e);

			return this.api;
		}
	}
}

export default ErrorHandler;
