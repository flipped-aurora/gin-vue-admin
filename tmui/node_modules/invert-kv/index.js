'use strict';

module.exports = object => {
	if (typeof object !== 'object' || object === null) {
		throw new TypeError('Expected an object');
	}

	const result = {};

	for (const [key, value] of Object.entries(object)) {
		result[value] = key;
	}

	for (const symbol of Object.getOwnPropertySymbols(object)) {
		const value = object[symbol];
		result[value] = symbol;
	}

	return result;
};
