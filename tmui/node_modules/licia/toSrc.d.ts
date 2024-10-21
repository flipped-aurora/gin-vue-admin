import types = require('./types');

declare function toSrc(fn: types.AnyFn): string;

export = toSrc;
