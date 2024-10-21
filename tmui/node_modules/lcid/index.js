'use strict';
const invertKv = require('invert-kv');
const all = require('./lcid.json');

const inverted = invertKv(all);

exports.from = lcidCode => {
	if (typeof lcidCode !== 'number') {
		throw new TypeError('Expected a number');
	}

	return all[lcidCode];
};

exports.to = localeId => {
	if (typeof localeId !== 'string') {
		throw new TypeError('Expected a string');
	}

	const lcidCode = inverted[localeId];
	if (lcidCode) {
		return Number(inverted[localeId]);
	}
};

exports.all = new Proxy(
	inverted,
	{
		get(target, name) {
			const lcid = target[name];
			if (lcid) {
				return Number(lcid);
			}
		}
	}
);
