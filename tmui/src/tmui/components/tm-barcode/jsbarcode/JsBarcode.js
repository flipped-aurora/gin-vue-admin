// Import all the barcodes
import barcodes from './barcodes/';
// Help functions
import merge from './help/merge.js';
import linearizeEncodings from './help/linearizeEncodings.js';
import fixOptions from './help/fixOptions.js';
import getRenderProperties from './help/getRenderProperties.js';
import optionsFromStrings from './help/optionsFromStrings.js';

// Exceptions
import ErrorHandler from './exceptions/ErrorHandler.js';
import { InvalidInputException, NoElementException } from './exceptions/exceptions.js';

// Default values
import defaults from './options/defaults.js';

// The protype of the object returned from the JsBarcode() call
let API = function () { };

// The first call of the library API
// Will return an object with all barcodes calls and the data that is used
// by the renderers
let JsBarcode = function (element, text, options) {
	var api = new API();

	if (typeof element === "undefined") {
		throw Error("No element to render on was provided.");
	}

	// Variables that will be pased through the API calls
	api._renderProperties = getRenderProperties(element);
	api._encodings = [];
	api._options = defaults;
	api._errorHandler = new ErrorHandler(api);

	// If text is set, use the simple syntax (render the barcode directly)
	if (typeof text !== "undefined") {
		options = options || {};

		if (!options.format || options.format == "auto") {
			options.format = autoSelectBarcode();
		}
		api.options(options)[options.format](text, options).render();
	}

	return api;
};

// To make tests work TODO: remove
JsBarcode.getModule = function (name) {
	return barcodes[name];
};

// Register all barcodes
for (var name in barcodes) {
	if (barcodes.hasOwnProperty(name)) { // Security check if the propery is a prototype property
		registerBarcode(barcodes, name);
	}
}
function registerBarcode(barcodes, name) {
	API.prototype[name] =
		API.prototype[name.toUpperCase()] =
		API.prototype[name.toLowerCase()] =
		function (text, options) {
			var api = this;
			return api._errorHandler.wrapBarcodeCall(function () {
				// Ensure text is options.text
				options.text = typeof options.text === 'undefined' ? undefined : '' + options.text;

				var newOptions = merge(api._options, options);
				newOptions = optionsFromStrings(newOptions);
				var Encoder = barcodes[name];
				var encoded = encode(text, Encoder, newOptions);
				api._encodings.push(encoded);

				return api;
			});
		};
}

// encode() handles the Encoder call and builds the binary string to be rendered
function encode(text, Encoder, options) {
	// Ensure that text is a string
	text = "" + text;

	var encoder = new Encoder(text, options);

	// If the input is not valid for the encoder, throw error.
	// If the valid callback option is set, call it instead of throwing error
	if (!encoder.valid()) {
		throw new InvalidInputException(encoder.constructor.name, text);
	}

	// Make a request for the binary data (and other infromation) that should be rendered
	var encoded = encoder.encode();

	// Encodings can be nestled like [[1-1, 1-2], 2, [3-1, 3-2]
	// Convert to [1-1, 1-2, 2, 3-1, 3-2]
	encoded = linearizeEncodings(encoded);

	// Merge
	for (let i = 0; i < encoded.length; i++) {
		encoded[i].options = merge(options, encoded[i].options);
	}

	return encoded;
}

function autoSelectBarcode() {
	// If CODE128 exists. Use it
	if (barcodes["CODE128"]) {
		return "CODE128";
	}

	// Else, take the first (probably only) barcode
	return Object.keys(barcodes)[0];
}

// Sets global encoder options
// Added to the api by the JsBarcode function
API.prototype.options = function (options) {
	this._options = merge(this._options, options);
	return this;
};

// Will create a blank space (usually in between barcodes)
API.prototype.blank = function (size) {
	const zeroes = new Array(size + 1).join("0");
	this._encodings.push({ data: zeroes });
	return this;
};

// Initialize JsBarcode on all HTML elements defined.
API.prototype.init = function () {
	// Should do nothing if no elements where found
	if (!this._renderProperties) {
		return;
	}

	// Make sure renderProperies is an array
	if (!Array.isArray(this._renderProperties)) {
		this._renderProperties = [this._renderProperties];
	}

	var renderProperty;
	for (let i in this._renderProperties) {
		renderProperty = this._renderProperties[i];
		var options = merge(this._options, renderProperty.options);

		if (options.format == "auto") {
			options.format = autoSelectBarcode();
		}

		this._errorHandler.wrapBarcodeCall(function () {
			var text = options.value;
			var Encoder = barcodes[options.format.toUpperCase()];
			var encoded = encode(text, Encoder, options);

			render(renderProperty, encoded, options);
		});
	}
};


// The render API call. Calls the real render function.
API.prototype.render = function () {
	if (!this._renderProperties) {
		throw new NoElementException();
	}

	if (Array.isArray(this._renderProperties)) {
		for (var i = 0; i < this._renderProperties.length; i++) {
			render(this._renderProperties[i], this._encodings, this._options);
		}
	}
	else {
		render(this._renderProperties, this._encodings, this._options);
	}

	return this;
};

API.prototype._defaults = defaults;

// Prepares the encodings and calls the renderer
function render(renderProperties, encodings, options) {
	encodings = linearizeEncodings(encodings);

	for (let i = 0; i < encodings.length; i++) {
		encodings[i].options = merge(options, encodings[i].options);
		fixOptions(encodings[i].options);
	}

	fixOptions(options);

	var Renderer = renderProperties.renderer;
	var renderer = new Renderer(renderProperties.element, encodings, options);
	renderer.render();

	if (renderProperties.afterRender) {
		renderProperties.afterRender();
	}
}

export { JsBarcode };